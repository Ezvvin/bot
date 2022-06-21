package handler

import (
	commandimpl "bot/internal/command_impl"
	"bot/internal/domain"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	userMap := map[int64]domain.Location{}

	// Проверяем каждое обновление
	for update := range updates {
		// Проверяем что сообщение не пустое
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		//TODO упростить команды
		case "/start":
			commandimpl.Start(userMap, bot, update)

		case "/close":
			log.Println("ВНИМАНИЕ СКРЫВАЕТСЯ КЛАВИАТУРА ТГ")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bye! Have a nice day! If i need you again, send `/start` in the chat!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "Каталог одежды":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите худи")
			msg.ReplyMarkup = domain.HoodyMenuKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			userMap[update.Message.From.ID] = domain.Location_HoodyCatalogMenu
		case "Black Neega":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "4500 рублей на стол")
			msg.ReplyMarkup = domain.BlackNeegaKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			userMap[update.Message.From.ID] = domain.Location_HoodyMenu

		case "Назад":
			commandimpl.Back(userMap, bot, update)

		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "For start send `/start` in chat")
			msg.ReplyMarkup = domain.StartKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			userMap[update.Message.From.ID] = domain.Location_StartMenu
		}

	}
}
