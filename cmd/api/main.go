package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/prerec/news-webserver-v1/internal/app/api"
)

var (
	configPath string
)

func init() {
	// Приложение на этапе запуска получает путь до конфиг файла
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file .toml")
}

func main() {
	// В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}
	server := api.New(config)

	// API server start
	log.Fatal(server.Start())
}
