package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prerec/news-webserver-v1/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *fiber.App
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: fiber.New(),
	}
}

func (api *API) Start() error {
	// Конфигурируем логгер
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// Подтверждение того, что логгер сконфигурирован
	api.logger.Info("Starting API server at port:", api.config.BindAddr)
	// Конфигурируем маршрутизатор
	api.configureRouterField()
	// Конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}
	// На этапе валидного завершения стартуем http сервер
	return api.router.Listen(api.config.BindAddr)
}
