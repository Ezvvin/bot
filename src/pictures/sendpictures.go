package pictures

// import (
// 	"io/ioutil"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func Sendpictur(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	photoBytes, err := ioutil.ReadFile("/your/local/path/to/picture.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	photoFileBytes := tgbotapi.FileBytes{
// 		Name:  "picture",
// 		Bytes: photoBytes,
// 	}
// 	chatID := 12345678
// 	Sendpictur, err := bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))
// }
