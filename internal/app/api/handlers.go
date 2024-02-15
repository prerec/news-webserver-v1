package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prerec/news-webserver-v1/internal/app/models"
)

// Message вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(c *fiber.Ctx) {
	// Установка заголовка Content-Type в "application/json"
	c.Set("Content-Type", "application/json")
}

func (api *API) GetAllNews(c *fiber.Ctx) error {
	// Инициализируем заголовки
	initHeaders(c)

	// Логируем момент начала обработки запроса
	api.logger.Info("Get All News GET /list")
	// Пытаемся что-то получить от БД
	news, err := api.storage.News().SelectAll()
	if err != nil {
		// Ошибка на этапе подключения
		api.logger.Info("Error while News.SelectAll:", err)
		msg := Message{
			StatusCode: fiber.StatusNotImplemented,
			Message:    "Troubles to accessing database. Try again later.",
			IsError:    true,
		}
		return c.Status(fiber.StatusInternalServerError).SendString(msg.Message)
	}
	return c.Status(fiber.StatusOK).JSON(news)
}

func (api *API) UpdateNewsByID(c *fiber.Ctx) error {
	// Инициализируем заголовки
	initHeaders(c)

	// Логируем момент начала обработки запроса
	api.logger.Info("Get update news POST /edit/:id")
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Ошибка на этапе получения URL параметра
		api.logger.Info("Error while News.EditByID:", err)
		msg := Message{
			StatusCode: fiber.StatusNotImplemented,
			Message:    "Troubles to convert id to int",
			IsError:    true,
		}
		return c.Status(fiber.StatusNotImplemented).SendString(msg.Message)
	}
	var req models.News
	if err := c.BodyParser(&req); err != nil {
		// Ошибка на этапе извлечения JSON в объект News
		api.logger.Info("Error while News.EditByID:", err)
		msg := Message{
			StatusCode: fiber.StatusNotImplemented,
			Message:    "Troubles to parse JSON",
			IsError:    true,
		}
		return c.Status(fiber.StatusNotImplemented).SendString(msg.Message)
	}
	err = api.storage.News().UpdateNewsByID(id, req)
	if err != nil {
		// Ошибка на этапе подключения
		api.logger.Info("Error while News.EditByID:", err)
		msg := Message{
			StatusCode: fiber.StatusNotImplemented,
			Message:    "Troubles to accessing database. Try again later.",
			IsError:    true,
		}
		return c.Status(fiber.StatusInternalServerError).SendString(msg.Message)
	}
	return c.Status(fiber.StatusOK).SendString("Edited!")
}
