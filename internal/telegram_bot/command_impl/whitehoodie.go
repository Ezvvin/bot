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
	image := tgbotapi.NewPhoto(430337954, tgbotapi.FilePath("D:/TelBot/bot/src/pictures/white_hoodie/whitehoodie.jpg"))
	_, err := bot.Send(image)

	if err != nil {
		panic(err)
	}
	// photoBytes, err := ioutil.ReadFile("D:/TelBot/bot/src/pictures/white_hoodie/whitehoodie.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// photoFileBytes := tgbotapi.FileBytes{
	// 	Name:  "whitehoodie.jpg",
	// 	Bytes: photoBytes,
	// }

	// chatID := 430337954
	// _, err = bot.Send(tgbotapi.NewPhoto(int64(chatID), photoFileBytes))
	// if err != nil {
	// 	panic(err)
	// }
	userMap[update.Message.From.ID] = domain.Location_WhiteHoodyMenu

}
