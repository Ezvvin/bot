package handler

import (
	"bot/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BackButtonHandler(userMap map[int64]config.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch userMap[update.Message.From.ID] {
	case config.Location_HoodyCatalogMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = config.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		userMap[update.Message.From.ID] = config.Location_MainMenu

	case config.Location_HoodyMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в каталог!")
		msg.ReplyMarkup = config.HoodyMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		userMap[update.Message.From.ID] = config.Location_HoodyCatalogMenu
	}
}
