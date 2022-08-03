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
	sticker := tgbotapi.NewSticker(update.Message.Chat.ID, tgbotapi.FilePath("C:/Users/Pavel/bot/src/pictures/black_hoodie/11.png"))
	image := tgbotapi.NewMediaGroup(msg.ChatID, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath("C:/Users/Pavel/bot/src/pictures/size/size.jpg"))})
	if _, err := bot.Send(sticker); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "whitehoodie")
	}
	if _, err := bot.SendMediaGroup(image); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "whitehoodie")
	}
	userMap[update.Message.From.ID] = domain.Location_BlackHoodieMenu
}
