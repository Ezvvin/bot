package telegrambot

import (
	"bot/internal/domain"
	commandimpl "bot/internal/telegram_bot/command_impl"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func InitHandler(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	logrus.Debug("Init bot handler")
	// Создаем мапу для отслеживания локации пользователя
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
