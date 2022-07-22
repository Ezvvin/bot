package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BackToMenu(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch userMap[update.Message.From.ID] {

	case domain.Location_BlackHoodieMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard

		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu
	case domain.Location_WhiteHoodieMenu:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard

		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_BlackSizeHoodie:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard

		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_WhiteSizeHoodie:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard

		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_Delivery:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_DeliveryCourier:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_DeliveryPoint:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_AcceptDelivery:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu

	case domain.Location_AddProduct:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись в главное меню!")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "back_To_menu_button")
		}
		userMap[update.Message.From.ID] = domain.Location_MainMenu
	}
	
}
