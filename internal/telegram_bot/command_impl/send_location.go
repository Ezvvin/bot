package commandimpl

import (
	db_domain "bot/internal/database/domain"
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendLocation(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase) {
	u := db_domain.User{Id: int(update.Message.From.ID)}
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			dbu.Users[i] = user
		}
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте пожалуйста ваш адрес для доставки!")
	msg.ReplyMarkup = domain.BackLocationKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "deliveryCourierbutton")
	}
	userMap[update.Message.From.ID] = domain.Location_SendAdress

}
