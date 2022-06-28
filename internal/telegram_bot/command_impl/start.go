package commandimpl

import (
	"bot/internal/domain"
	"fmt"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s, Добро пожаловать в наш магазин одежды  'LÚQ' 👋", update.Message.From.FirstName))
	msg.ReplyMarkup = domain.MainMenuKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "start")
	}
	userMap[update.Message.From.ID] = domain.Location_MainMenu
}
