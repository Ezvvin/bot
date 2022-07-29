package commandimpl

import (
	db_domain "bot/internal/database/domain"
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeliveryPoint(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase) {
	u := db_domain.User{Id: int(update.Message.From.ID)}
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			user.Delivery = "Самовывоз"
			dbu.Users[i] = user
		}
	}
	sendpoint := tgbotapi.NewLocation(update.Message.Chat.ID, 55.773398, 37.494530)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Заказ можете забрать по адресу:\nг.Москва, Ул. Мнёвники, д. 10к1")
	msg.ReplyMarkup = domain.SendContactKeyboard

	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliverypointhoodiebutton")
	}
	if _, err := bot.Send(sendpoint); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliverypointhoodiebutton")
	}
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте пожалуйста ваш контакт для связи! ")
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliverypointhoodiebutton")
	}
	userMap[update.Message.From.ID] = domain.Location_DeliveryPoint
}
