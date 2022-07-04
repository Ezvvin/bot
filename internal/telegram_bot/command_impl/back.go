package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Back(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch userMap[update.Message.From.ID] {
	case domain.Location_HoodyCatalogMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "backbutton")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_HoodyColorMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в каталог!")
		msg.ReplyMarkup = domain.HoodyMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "backbutton")
		}
		userMap[update.Message.From.ID] = domain.Location_HoodyCatalogMenu

	}
}
