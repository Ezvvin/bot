package commandimpl

import (
	"bot/internal/domain"
	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BlackHoodieCommand(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "4500 рублей")
	msg.ReplyMarkup = domain.BuyHoodieKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Panic(domain.ErrCommand_Init)
	}
	userMap[update.Message.From.ID] = domain.Location_BlackHoodyMenu

}