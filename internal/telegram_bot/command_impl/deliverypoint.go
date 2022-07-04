package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeliveryPoint(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Заказ можете забрать по адресу: улица Мнёвники, д. 10к1")
	msg.ReplyMarkup = domain.DeliveryKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliverypointhoodiebutton")
	}
	userMap[update.Message.From.ID] = domain.Location_DeliveryPoint

}
