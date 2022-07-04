package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeliveryCourier(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Для оформления заказа требуется ваш телефон. Нажмите кнопку для отправки контакта.")
	msg.ReplyMarkup = domain.DeliveryCourierKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliveryCourierbutton")
	}
	userMap[update.Message.From.ID] = domain.Location_DeliveryCourier

}