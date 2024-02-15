package api

import (
	"github.com/prerec/news-webserver-v1/storage"
	"github.com/sirupsen/logrus"
)

// Конфигурируем логгер
func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(logLevel)
	return nil
}

// Конфигурируем маршрутизатор
func (api *API) configureRouterField() {
	api.router.Get("/list", api.GetAllNews)
	api.router.Post("/edit/:id", api.UpdateNewsByID)
}

// Конфигурируем хранилище (storage API)
func (api *API) configureStorageField() error {
	storage := storage.New(api.config.Storage)
	// Пытаемся установить соединение, если невозможно - возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	api.storage = storage
	return nil
}
