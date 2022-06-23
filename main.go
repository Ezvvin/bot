package main

import (
	"bot/internal/domain"
	"bot/internal/logger"
	tbot "bot/internal/telegram_bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

var cfg domain.Config

func init() {
	if err := configor.Load(&cfg, "config.yaml"); err != nil {
		log.WithError(err).Fatal(domain.LoggerConfigError)
	}
}

func main() {
	// Иницилизация логгера
	logger.InitLogger()
	// Установка уровня логгера
	log.SetLevel(logger.ParseLogLevel(cfg))

	// TODO: Переместить создание бота в отдельный пакет с собственной структурой
	// Подключаем бота
	bot, err := tgbotapi.NewBotAPI(cfg.Token) //TODO скрыть ебаный токен
	if err != nil {
		log.WithError(err).Fatal(domain.TelegramBotError_Init)
	}

	// Высставляем боту дебаг
	bot.Debug = true

	log.WithField("Bot name", bot.Self.UserName).Debug("Authorized on account")

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)
	tbot.InitHandler(bot, updates)
}
