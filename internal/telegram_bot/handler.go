package telegrambot

import (
	commandimpl "bot/internal/command_impl"
	"bot/internal/domain"

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
		case "/start":
			commandimpl.Start(userMap, bot, update)

		case "/close":
			commandimpl.Close(bot, update)

		case "Каталог одежды":
			commandimpl.Catalog(userMap, bot, update)

		case "Black Hoodie":
			commandimpl.BlackHoodieCommand(userMap, bot, update)

		case "White Hoodie":
			commandimpl.WhiteHoodieCommand(userMap, bot, update)

		case "Назад":
			commandimpl.Back(userMap, bot, update)

		default:
			commandimpl.Undefined(userMap, bot, update)
		}

	}
}
