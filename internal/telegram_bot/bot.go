package telegrambot

import (
	"bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Telegrambot struct {
	Bot       *tgbotapi.BotAPI
	UpdateCfg tgbotapi.UpdateConfig
}

func InitBot(cfg domain.BotConfig) (*Telegrambot, error) {
	// Создание бота
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	// Высставляем боту дебаг
	bot.Debug = cfg.Debug

	log.WithField("Bot name", bot.Self.UserName).Debug("Authorized on account")

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.Timeout
	return &Telegrambot{Bot: bot, UpdateCfg: u}, nil
}
func (bot *Telegrambot) SendPhoto() {
	image := tgbotapi.NewPhoto(5437936243, tgbotapi.FilePath("../../src/picturies/white_hoodie/whitehoodie.jpg"))
	_, err := bot.Bot.Send(image)

	if err != nil {
		panic(err)
	}
}
