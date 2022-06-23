package main

import (
	"bot/internal/config"
	"bot/internal/handler"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/configor"
)

var cfg config.Config

func init() {
	if err := configor.Load(&cfg, "config.yaml"); err != nil {
		log.Fatalf("Couldn't load config, err: '%s'", err.Error())
	}
}

func main() {
	//TODO: Добавить адекватный логгер
	log.Println(cfg)

	// TODO: Переместить создание бота в отдельный пакет с собственной структурой
	// Подключаем бота
	bot, err := tgbotapi.NewBotAPI(cfg.Token) //TODO скрыть ебаный токен
	if err != nil {
		log.Panic(err)
	}

	// Высставляем боту дебаг
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)
	handler.Handler(bot, updates)
}
