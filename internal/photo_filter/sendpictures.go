package photo_filter

import (
	// "io/ioutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Sendpictur(bot *tgbotapi.BotAPI, update tgbotapi.Update, message *tgbotapi.Message) {
	image := tgbotapi.NewPhoto(5437936243, tgbotapi.FilePath("../../src/picturies/white_hoodie/whitehoodie.jpg"))
	_, err := bot.Send(image)

	if err != nil {
		panic(err)
	}
	// photoBytes, err := ioutil.ReadFile("../../src/picturies/white_hoodie/whitehoodie.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// photoFileBytes := tgbotapi.FileBytes{
	// 	Name:  "whitehoodie.jpg",
	// 	Bytes: photoBytes,
	// }

	// chatID := 5437936243
	// Message, err := bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))
}
