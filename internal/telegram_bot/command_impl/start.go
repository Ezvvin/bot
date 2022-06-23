package commandimpl

import (
	"bot/internal/domain"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s, Добро пожаловать в наш магазин одежды  'LÚQ' 👋", update.Message.From.UserName))
	msg.ReplyMarkup = domain.MainMenuKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	userMap[update.Message.From.ID] = domain.Location_MainMenu
}
