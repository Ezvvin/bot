package commandimpl

import (
	"bot/internal/domain"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Back(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch userMap[update.Message.From.ID] {
	case domain.Location_HoodyCatalogMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_HoodyMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в каталог!")
		msg.ReplyMarkup = domain.HoodyMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		userMap[update.Message.From.ID] = domain.Location_HoodyCatalogMenu
	}
}
