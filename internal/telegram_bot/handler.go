package telegrambot

import (
	"bot/internal/domain"
	commandimpl "bot/internal/telegram_bot/command_impl"

	"github.com/sirupsen/logrus"
)

func (bot *Telegrambot) InitHandler() {
	logrus.Debug("Init bot handler")
	updates := bot.Bot.GetUpdatesChan(bot.UpdateCfg)
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
			commandimpl.Start(userMap, bot.Bot, update)

		case "Каталог одежды🥼":
			commandimpl.Catalog(userMap, bot.Bot, update)

		case "🖤Black Hoodie🖤":
			commandimpl.BlackHoodieCommand(userMap, bot.Bot, update)

		case "🤍White Hoodie🤍":
			commandimpl.WhiteHoodieCommand(userMap, bot.Bot, update)

		case "Выбрать размер📏":
			commandimpl.InfoHoodie(userMap, bot.Bot, update)

		case "S-46 (EUR)":
			commandimpl.SizeHoodie(userMap, bot.Bot, update)

		case "Отзывы🔥":
			commandimpl.Commends(userMap, bot.Bot, update)

		case "◀️Назад":
			commandimpl.Back(userMap, bot.Bot, update)

		case "⌨️":
			commandimpl.Close(bot.Bot, update)

		default:
			commandimpl.Undefined(userMap, bot.Bot, update)
		}

	}
}
