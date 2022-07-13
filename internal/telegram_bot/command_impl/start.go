package commandimpl

import (
	db_domain "bot/internal/database/domain"
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func Start(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s, Добро пожаловать в наш магазин одежды  'LÚQ' 👋", update.Message.From.FirstName))
	msg.ReplyMarkup = domain.MainMenuKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "start")
	}
	userMap[update.Message.From.ID] = domain.Location_MainMenu
	cart := db_domain.Cart{Id: int(update.Message.From.ID)}
	user := db_domain.User{
		Id:        int(update.Message.From.ID),
		UserCart:  cart,
		Username:  update.Message.From.UserName,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
	}
	dbu.AddUser(user)
	dbu.AddCart(cart)
	log.WithField("cart", dbu.Carts).Debug("список корзин")
	log.WithField("user", dbu.Users).Debug("список юзеров")
}
