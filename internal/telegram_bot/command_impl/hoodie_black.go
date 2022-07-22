package commandimpl

import (
	"bot/internal/domain"

	db_usecase "bot/internal/database/usecase"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BlackHoodieCommand(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🖤BLACKLOGO LÚQ HOODIE🖤\nМатериал: Футер 3х- нитка, 85% ХБ, 15% ПЭ🧵\n4300 рублей💰")
	msg.ReplyMarkup = domain.AddProductKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "blackhoodiebutton")
	}
	userMap[update.Message.From.ID] = domain.Location_BlackHoodieMenu
}
