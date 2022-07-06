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
	user := db_domain.User{
		Id:       int(update.Message.From.ID),
		Username: update.Message.From.UserName,
	}
	// TODO: create user cart

	log.Debug(user) // чтобы не жаловался на не использование

	// TODO: Add user in db
	// for example `dbu.AddUser(user)`
	// TODO: Add user cart
}
