package commandimpl

import (
	"bot/internal/domain"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WhiteHoodieCommand(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "4300 рублей")
	msg.ReplyMarkup = domain.BuyHoodieKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	image := tgbotapi.NewMediaGroup(430337954, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath("D:/TelBot/bot/src/pictures/white_hoodie/whitehoodie.jpg")),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath("D:/TelBot/bot/src/pictures/white_hoodie/whitehoodie2.jpg"))})
	_, err := bot.SendMediaGroup(image)

	if err != nil {
		panic(err)
	}
	userMap[update.Message.From.ID] = domain.Location_WhiteHoodyMenu

}
