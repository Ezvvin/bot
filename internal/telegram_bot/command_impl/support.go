package commandimpl

// import (
// 	"bot/internal/domain"

// 	log "github.com/sirupsen/logrus"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func Support(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Опишите вашу проблему, наш менеджер свяжется с вами и поможет вам!")
// 	msg.ReplyMarkup = domain.MainMenuKeyboard
// 	copymessage := tgbotapi.MessageID{MessageID: 2}
// 	if _, err := bot.CopyMessage(tgbotapi.NewCopyMessage(5437936243, 430337954, copymessage.MessageID )); err != nil {
// 		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "supportbutton")
// 	}

// }
