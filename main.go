package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ссылка"),
		tgbotapi.NewKeyboardButton("голый максим"),
		tgbotapi.NewKeyboardButton("худи"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
	),
)

//Создаем бота
func main() {
	bot, err := tgbotapi.NewBotAPI("5437936243:AAG8qSRApD7V90ZZUVDM4ze0bcRbUC1rbrE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		//Проверяем что от пользователья пришло именно текстовое сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm  bot.")
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard

		case "close":
			log.Println("скрыть мать")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
		}

	}
}
