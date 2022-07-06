package telegrambot

import (
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	commandimpl "bot/internal/telegram_bot/command_impl"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func (bot *Telegrambot) InitHandler(cfg domain.Config, dbu *db_usecase.DataBaseUsecase) {
	log.Debug("Init bot handler")
	updates := bot.Bot.GetUpdatesChan(bot.UpdateCfg)
	// Создаем мапу для отслеживания локации пользователя
	userMap := map[int64]domain.Location{}
	// Проверяем каждое обновление
	for update := range updates {
		// Проверяем что сообщение не пустое
		if update.Message == nil {
			continue
		}
		// проверяем , имеет ли сообщение формат контакта
		if update.Message.Contact != nil {
			log.Debug("sos", update.Message.Contact)
			continue
		}
		if userMap[update.Message.From.ID] == domain.Location_Support {
			msgSupport := tgbotapi.NewMessage(cfg.AdminChat, fmt.Sprintf(strconv.FormatInt(update.Message.From.ID, 10), "Клиент", update.FromChat().UserName, "ждет ответа на вопрос:", update.Message.Text))
			if _, err := bot.Bot.Send(msgSupport); err != nil {
				log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "groupmessage")
			}
			userMap[update.Message.From.ID] = domain.Location_MainMenu
			continue
		}
		// TODO сделат ькнопку назад \ главное меню

		switch userMap[update.Message.From.ID] {
		case domain.Location_MainMenu:
			switch update.Message.Text {
			case "Каталог одежды🥼":
				commandimpl.Catalog(userMap, bot.Bot, update)

			case "Отзывы🔥":
				commandimpl.Commends(userMap, bot.Bot, update)

			case "Поддержка":
				commandimpl.Support(userMap, bot.Bot, update)

			case "Магазин LÚQ":
				commandimpl.Contacts(bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)

			}
		case domain.Location_HoodyCatalogMenu:
			switch update.Message.Text {
			case "🖤Black Hoodie🖤":
				commandimpl.BlackHoodieCommand(userMap, bot.Bot, update)

			case "🤍White Hoodie🤍":
				commandimpl.WhiteHoodieCommand(userMap, bot.Bot, update)
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_HoodyColorMenu:
			switch update.Message.Text {
			case "Выбрать размер📏":
				commandimpl.InfoHoodie(userMap, bot.Bot, update)
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_SizeHoodie:
			switch update.Message.Text {
			case "S-46 (EUR)", "M-48 (EUR)", "L-50 (EUR)":
				commandimpl.SizeHoodie(userMap, bot.Bot, update)
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_Delivery:
			switch update.Message.Text {
			case "Самовывоз":
				commandimpl.DeliveryPoint(userMap, bot.Bot, update)

			case "Доставка по адресу":
				commandimpl.DeliveryCourier(userMap, bot.Bot, update)
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_DeliveryCourier:
			switch update.Message.Text {
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_DeliveryPoint:
			switch update.Message.Text {
			case "Подтвердить заказ":
				commandimpl.AcceptDelivery(userMap, bot.Bot, update)
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_AcceptDelivery:
			switch update.Message.Text {
			case "Оплатить заказ":
				commandimpl.PayHoodie(userMap, bot.Bot, update, cfg.BotConfig)
			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)

			}
		default:
			switch update.Message.Text {
			case "/start":
				commandimpl.Start(userMap, bot.Bot, update, dbu)
				// TODO: add user in dbu

			}
		}
	}
}
