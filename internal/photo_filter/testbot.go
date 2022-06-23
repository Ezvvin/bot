package photo_filter

// import (
// 	"io/ioutil"
// 	"os"
// 	"testing"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func getBot(t *testing.T) (*tgbotapi.BotAPI, error) {
// 	bot, err := tgbotapi.NewBotAPI("")
// 	bot.Debug = true

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	return bot, err
// }

// func TestGetUpdates(t *testing.T) {
// 	bot, _ := getBot(t)

// 	u := tgbotapi.NewUpdate(0)

// 	_, err := bot.GetUpdates(u)

// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSendWithNewPhoto(t *testing.T) {
// 	bot, _ := getBot(t)

// 	msg := tgbotapi.NewPhoto(5437936243, tgbotapi.FilePath("src/pictures/whitehoodie.jpg"))
// 	msg.Caption = "Test"
// 	_, err := bot.Send(msg)

// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSendWithNewPhotoWithFileBytes(t *testing.T) {
// 	bot, _ := getBot(t)

// 	data, _ := ioutil.ReadFile("src/pictures/whitehoodie.jpg")
// 	b := tgbotapi.FileBytes{Name: "src/pictures/whitehoodie.jpg", Bytes: data}

// 	msg := tgbotapi.NewPhoto(5437936243, b)
// 	msg.Caption = "Test"
// 	_, err := bot.Send(msg)

// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSendWithNewPhotoWithFileReader(t *testing.T) {
// 	bot, _ := getBot(t)

// 	f, _ := os.Open("src/pictures/whitehoodie.jpg")
// 	reader := tgbotapi.FileReader{Name: "src/pictures/whitehoodie.jpg", Reader: f}

// 	msg := tgbotapi.NewPhoto(5437936243, reader)
// 	msg.Caption = "Test"
// 	_, err := bot.Send(msg)

// 	if err != nil {
// 		t.Error(err)
// 	}
// }
