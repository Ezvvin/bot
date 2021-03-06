package commandimpl

import (
	db_domain "bot/internal/database/domain"
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func CartMenu(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg domain.Config, dbu *db_usecase.DataBaseUsecase) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	u := db_domain.User{Id: int(update.Message.From.ID)}
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			u = dbu.Users[i]
		}
	}
	if u.UserCart.Products == nil || len(u.UserCart.Products) == 0 {
		msg.Text = "Ваша корзина пустая!"
		msg.ReplyMarkup = domain.CartMenuKeyboardIfNil
	} else {
		msg.Text = "Ваша корзина:\n"
		for i, product := range u.UserCart.Products {
			i += 1
			msg.Text += strings.NewReplacer("{", "", "}", "").Replace((fmt.Sprintf("%d - %v\n", i, product)))
		}
		msg.ReplyMarkup = domain.CartMenuKeyboard
	}
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "cartbutton")
	}
	userMap[update.Message.From.ID] = domain.Location_CartMenu
}
