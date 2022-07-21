package commandimpl

import (
	db_domain "bot/internal/database/domain"
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func CartMenu(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg domain.Config, dbu *db_usecase.DataBaseUsecase) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyMarkup = domain.CartMenuKeyboard
	u := db_domain.User{Id: int(update.Message.From.ID)}
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			u = dbu.Users[i]
		}
	}
	if u.UserCart.Products == nil {
		msg.Text = "Ваша корзина пустая!"
	} else {
		msg.Text = fmt.Sprintf("Ваша корзина:%v", u.UserCart.Products)
	}
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "cartbutton")
	}
	userMap[update.Message.From.ID] = domain.Location_CartMenu
}
