package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AcceptDelivery(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш заказ:\nЦвет:\nРазмер:\nТип доставки:\nВаш контакт:")
	msg.ReplyMarkup = domain.PayKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "accepthoodiebutton")
	}
	userMap[update.Message.From.ID] = domain.Location_AcceptDelivery
}
