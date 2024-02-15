package storage

import (
	"log"

	"github.com/lib/pq"
	"github.com/prerec/news-webserver-v1/internal/app/models"
)

type NewsResponseRepository struct {
	storage *Storage
}

func (a *NewsResponseRepository) SelectAll() (models.Response, error) {
	queryNews := `
        SELECT u.Id, u.Title, u.Content, array_agg(c.CategoryId) AS Categories
        FROM news AS u
        LEFT JOIN newscategories AS c ON c.newsid = u.id
        GROUP BY u.id, u.Title, u.Content
    `
	rowsNews, err := a.storage.db.Query(queryNews)
	if err != nil {
		return models.Response{}, err
	}
	defer rowsNews.Close()

	response := models.Response{}
	news := make([]models.News, 0)
	for rowsNews.Next() {
		response.Success = true
		n := models.News{}
		var categories pq.Int64Array
		err := rowsNews.Scan(&n.ID, &n.Title, &n.Content, &categories)
		if err != nil {
			log.Println(err)
			continue
		}
		n.Categories = make([]int, len(categories))
		for i, v := range categories {
			n.Categories[i] = int(v)
		}
		news = append(news, n)
	}

	response.News = news
	return response, nil
}

func (a *NewsResponseRepository) UpdateNewsByID(id int, updatedNews models.News) error {

	var query string

	if updatedNews.Title != "" {
		query = `
        UPDATE news
        SET title = $2
        WHERE id = $1
    `
	}
	_, err := a.storage.db.Exec(query, id, updatedNews.Title)
	if err != nil {
		return err
	}

	if updatedNews.Content != "" {
		query = `
        UPDATE news
        SET content = $2
        WHERE id = $1
    `
	}
	_, err = a.storage.db.Exec(query, id, updatedNews.Content)
	if err != nil {
		return err
	}

	if len(updatedNews.Categories) != 0 {
		query = `
        UPDATE news
        SET id = $2
        WHERE Id = $1
    `
	}
	// Удаляем существующие категории
	_, err = a.storage.db.Exec("DELETE FROM newscategories WHERE NewsId = $1", id)
	if err != nil {
		return err
	}
	// Вставляем новые категории
	for _, categoryID := range updatedNews.Categories {
		_, err := a.storage.db.Exec("INSERT INTO newscategories (newsid, categoryid) VALUES ($1, $2)", id, categoryID)
		if err != nil {
			return err
		}
	}

	return nil
}
