package models

type News struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Categories []int  `json:"categories"`
}

type NewsCategories struct {
	NewsID     int `json:"news_id"`
	CategoryID int `json:"category_id"`
}

type Response struct {
	Success bool   `json:"success"`
	News    []News `json:"news"`
}
