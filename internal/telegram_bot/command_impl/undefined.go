package commandimpl

import (
	"bot/internal/domain"
	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Undefined(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мы вас не понимаем. Напишите /start для общения с ботом")
	msg.ReplyMarkup = domain.StartKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "undefined")
	}
	userMap[update.Message.From.ID] = domain.Location_StartMenu
}
