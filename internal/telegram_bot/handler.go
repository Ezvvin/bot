package telegrambot

import (
	"bot/internal/domain"
	commandimpl "bot/internal/telegram_bot/command_impl"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func (bot *Telegrambot) InitHandler(cfg domain.Config) {
	log.Debug("Init bot handler")
	updates := bot.Bot.GetUpdatesChan(bot.UpdateCfg)
	// Создаем мапу для отслеживания локации пользователя
	userMap := map[int64]domain.Location{}
	// flag := false
	// Проверяем каждое обновление
	for update := range updates {
		// Проверяем что сообщение не пустое
		if update.Message == nil {
			continue
		}

		if update.Message.Contact != nil {
			log.Debug("sos", update.Message.Contact)
			continue
		}
		if userMap[update.Message.From.ID] == domain.Location_Support {
			copymessage := tgbotapi.NewCopyMessage(cfg.AdminChat, update.Message.Chat.ID, update.Message.MessageID)
			if _, err := bot.Bot.Send(copymessage); err != nil {
				log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "copymessage")
			}
			userMap[update.Message.From.ID] = domain.Location_MainMenu
			continue
		}

		switch update.Message.Text {
		// TODO: Сгруппируй по категориям(можно по локации например и как-то еще, подумай), иначе через 10 комманд ьуь юедт неразбериха...
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

		case "S-46 (EUR)", "M-48 (EUR)", "L-50 (EUR)":
			commandimpl.SizeHoodie(userMap, bot.Bot, update)

		case "Самовывоз":
			commandimpl.DeliveryPoint(userMap, bot.Bot, update)

		case "Доставка по адресу":
			commandimpl.DeliveryCourier(userMap, bot.Bot, update)

		case "Отзывы🔥":
			commandimpl.Commends(userMap, bot.Bot, update)

		case "Поддержка":
			commandimpl.Support(userMap, bot.Bot, update)

		case "Магазин LÚQ":
			commandimpl.Contacts(bot.Bot, update)

		case "Подтвердить заказ":
			commandimpl.AcceptDelivery(userMap, bot.Bot, update)

		case "Оплатить заказ":
			commandimpl.PayHoodie(userMap, bot.Bot, update)

		case "◀️Назад":
			commandimpl.Back(userMap, bot.Bot, update)

		case "phone_number":
			commandimpl.Catalog(userMap, bot.Bot, update)

		default:
			// ОСТОРОЖНО КОСТЫЛЬ
			// if flag {
			// 	copymessage := tgbotapi.NewCopyMessage(430337954, update.Message.Chat.ID, update.Message.MessageID)
			// 	if _, err := bot.Bot.Send(copymessage); err != nil {
			// 		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "copymessage")
			// 	}
			// 	continue
			// }
			commandimpl.Undefined(userMap, bot.Bot, update)
		}

	}
}
