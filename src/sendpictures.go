package pictures

import (
	"io/ioutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Sendpictur(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	photoBytes, err := ioutil.ReadFile("bot/src/pictures/whitehoodie.png")
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "whitehoodie",
		Bytes: photoBytes,
	}
	chatID := 12345678
	Sendpictur, err := bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))
}
