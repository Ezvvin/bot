package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Support(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Опишите вашу проблему, наш менеджер свяжется с вами и поможет вам!")
	msg.ReplyMarkup = domain.MainMenuKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "supportbutton")
	}
	userMap[update.Message.From.ID] = domain.Location_Support
	// *flag = true

}
