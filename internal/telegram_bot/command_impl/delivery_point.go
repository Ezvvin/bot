package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeliveryPoint(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	sendpoint := tgbotapi.NewLocation(update.Message.Chat.ID, 55.773398, 37.494530)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Заказ можете забрать по адресу:\nг.Москва, Ул. Мнёвники, д. 10к1")
	msg.ReplyMarkup = domain.AcceptKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliverypointhoodiebutton")
	}
	if _, err := bot.Send(sendpoint); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliverypointhoodiebutton")
	}
	userMap[update.Message.From.ID] = domain.Location_DeliveryPoint

}
