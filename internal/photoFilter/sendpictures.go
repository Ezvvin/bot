package photoFilter

// TODO: В соурсе не должно быть исполнительных файлов, все го файлы должны быть в интернал
// import (
// 	"image"
// 	"io/ioutil"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func Sendpictur(bot *tgbotapi.BotAPI, update tgbotapi.Update, message *tgbotapi.Message) {
// 	image := tgbotapi.NewPhotoUpload(5437936243, "bot/src/pictures/whitehoodie.png")
// 	_, err = bot.Send(image)

// 	if err != nil {
// 		panic(err)
// 	}
// 	photoBytes, err := ioutil.ReadFile("bot/src/pictures/whitehoodie.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	photoFileBytes := tgbotapi.FileBytes{
// 		Name:  "whitehoodie",
// 		Bytes: photoBytes,
// 	}

// 	chatID := 5437936243
// 	Message, err := bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))
// }
