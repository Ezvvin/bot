package main

import (
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	"bot/internal/logger"
	telegrambot "bot/internal/telegram_bot"

	"github.com/jasonwinn/geocoder"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

var cfg domain.Config

func init() {
	if err := configor.Load(&cfg, "config.yaml"); err != nil {
		log.WithError(err).Fatal(domain.ErrLogger_Config)
	}
}

func main() {
	// Иницилизация логгера
	logger.InitLogger()
	// Установка уровня логгера
	log.SetLevel(logger.ParseLogLevel(cfg))
	// apikey яндекс карт
	geocoder.SetAPIKey(cfg.ApiKeyMap)
	// Подключаем бота
	bot, err := telegrambot.InitBot(cfg.BotConfig)
	if err != nil {
		log.WithError(err).Fatal(domain.ErrTelegramBot_Init)
	}

	// Создаем юзкейз "базы данных"
	dbUsecase := db_usecase.InitDataBaseUsecase()

	//Получаем обновления от бота
	bot.InitHandler(cfg, dbUsecase)

}
