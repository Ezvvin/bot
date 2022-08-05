package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SizeHoodie(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Для добавления товара в корзину, выберите пожалуйста, нужный вам размер:")
	msg.ReplyMarkup = domain.SizeHoodieKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "Sizehoodiebutton")
	}
	if userMap[update.Message.From.ID] == domain.Location_BlackHoodieMenu {
		userMap[update.Message.From.ID] = domain.Location_BlackSizeHoodie
	}
	if userMap[update.Message.From.ID] == domain.Location_WhiteHoodieMenu {
		userMap[update.Message.From.ID] = domain.Location_WhiteSizeHoodie
	}
}
