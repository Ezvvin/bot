package handler

import (
	"bot/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message == nil {
			continue
		}
		//Проверяем что от пользователья пришло именно текстовое сообщение

		switch update.Message.Text {
		//TODO добавить команды
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm  bot. Choose option:")
			msg.ReplyMarkup = config.MainMenuKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

		case "/close":
			log.Println("ВНИМАНИЕ СКРЫВАЕТСЯ КЛАВИАТУРА ТГ")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bye! Have a nice day! If i need you again, send `/start` in the chat!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "Каталог одежды":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите худи")
			msg.ReplyMarkup = config.HoodyMenuKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "For start send `/start` in chat")
			msg.ReplyMarkup = config.StartKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}
}
