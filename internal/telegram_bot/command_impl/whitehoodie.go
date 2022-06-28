package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WhiteHoodieCommand(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🤍WHITELOGO LUQ HOODIE🤍\nМатериал: Футер 3х- нитка, 85% ХБ, 15% ПЭ🧵\n4300 рублей💰")
	msg.ReplyMarkup = domain.BuyHoodieKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Panic(domain.ErrCommand_Init)
	}

	image := tgbotapi.NewMediaGroup(msg.ChatID, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath("C:/Users/Pavel/bot/src/pictures/white_hoodie/whitehoodie.jpg")),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath("C:/Users/Pavel/bot/src/pictures/white_hoodie/whitehoodie2.jpg"))})
	_, err := bot.SendMediaGroup(image)

	if err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "whitehoodie")
	}
	userMap[update.Message.From.ID] = domain.Location_WhiteHoodyMenu

}
