package commandimpl

// import (
// 	"bot/internal/domain"

// 	log "github.com/sirupsen/logrus"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func Close(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "`/start` in the chat!")
// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
// 	if _, err := bot.Send(msg); err != nil {
// 		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "hidekeyboard")
// 	}

// }
