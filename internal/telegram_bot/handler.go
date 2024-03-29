package telegrambot

import (
	db_domain "bot/internal/database/domain"
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
		u := db_domain.User{Id: int(update.Message.From.ID)}
		// проверяем, имеет ли сообщение формат адреса
		if userMap[update.Message.From.ID] == domain.Location_SendAdress {
			if update.Message.Text == "◀️Назад" {
				commandimpl.Back(userMap, bot.Bot, update)
				continue
			}
			//если отправлено сообщение, записывает в адрес юзера
			if update.Message.Text != "" {
				for i, user := range dbu.Users {
					if user.Id == u.Id {
						user.Adress = update.Message.Text
						dbu.Users[i] = user
					}
				}
			}
			commandimpl.AcceptDelivery(userMap, bot.Bot, update, cfg, dbu)
			continue
		}
		// проверяем , имеет ли сообщение формат контакта
		if update.Message.Contact != nil {
			// записываем номер в пользователя из доставки курьером
			if userMap[update.Message.From.ID] == domain.Location_SendContact {
				if update.Message.Contact.PhoneNumber != "" {
					for i, user := range dbu.Users {
						if user.Id == u.Id {
							user.Phone = update.Message.Contact.PhoneNumber
							dbu.Users[i] = user
						}
					}
				}
				commandimpl.SendLocation(userMap, bot.Bot, update, dbu)
				continue
			}
			// записываем номер в пользователя из доставки самовывозом
			if userMap[update.Message.From.ID] == domain.Location_DeliveryPoint {
				if update.Message.Contact.PhoneNumber != "" {
					for i, user := range dbu.Users {
						if user.Id == u.Id {
							user.Phone = update.Message.Contact.PhoneNumber
							dbu.Users[i] = user
						}
					}
				}
				commandimpl.AcceptDelivery(userMap, bot.Bot, update, cfg, dbu)
				continue
			}
		}
		// удаление продукта из корзины
		if userMap[update.Message.From.ID] == domain.Location_DeleteProduct {
			if update.Message.Text == "◀️Назад" {
				commandimpl.Back(userMap, bot.Bot, update)
				continue
			}
			// ждем индекс продукта (номер строки продукта у пользователя)
			if update.Message.Text != "" {
				//конвертируем номер строки продукта у пользователя в индекс продукта
				index, err := strconv.Atoi(update.Message.Text)
				if err != nil {
					log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "Index Product")
				}
				//ищем введенный индекс среди продуктов
				for i, user := range dbu.Users {
					if user.Id == u.Id {
						index -= 1
						if i != index {
							bot.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Такого товара нет в вашей корзине"))
						}
						//удаляем продукт из корзины и обновляем корзину юзера
						user.UserCart.RemoveProduct(index)
						dbu.Users[i] = user
						dbu.UpdateCart(user.UserCart)
					}
				}
				commandimpl.CartMenu(userMap, bot.Bot, update, cfg, dbu)
				continue
			}
		}
		//обратная связь, ждем сообщение от пользователя и отправляем в админ группу
		if userMap[update.Message.From.ID] == domain.Location_Support {
			if update.Message.Text == "◀️Назад" {
				commandimpl.Back(userMap, bot.Bot, update)
				continue
			}
			//оправляем вопрос в группу 
			msgSupport := tgbotapi.NewMessage(cfg.AdminChat, fmt.Sprintf(("ID: %d\nКлиент: %s\nВопрос: %s\n"), update.Message.From.ID, (fmt.Sprintf("<a href='tg://user?id=%v'>%s</a>", update.Message.From.ID, update.Message.From.FirstName)), update.Message.Text))
			//меняем парс мод для ссылки на профиль пользователя
			msgSupport.ParseMode = "HTML"
			if _, err := bot.Bot.Send(msgSupport); err != nil {
				log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "supportmsg")
			}
			//отправляем ответ пользователю
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В ближайшее время с вами свяжется наш менеджер!")
			msg.ReplyMarkup = domain.MainMenuKeyboard
			if _, err := bot.Bot.Send(msg); err != nil {
				log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "msgforusersupport")
			}
			userMap[update.Message.From.ID] = domain.Location_MainMenu
			continue

		}
		// команды

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

			case "Корзина":
				commandimpl.CartMenu(userMap, bot.Bot, update, cfg, dbu)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)

			}
		case domain.Location_CartMenu:
			switch update.Message.Text {

			case "Сделать заказ":
				commandimpl.Delivery(userMap, bot.Bot, update)

			case "Удалить товар":
				commandimpl.DeleteProduct(userMap, bot.Bot, update, dbu)

			case "Каталог одежды🥼":
				commandimpl.Catalog(userMap, bot.Bot, update)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)

			}
		case domain.Location_HoodyCatalogMenu:
			switch update.Message.Text {

			case "🖤Black Hoodie🖤":
				commandimpl.BlackHoodieCommand(userMap, bot.Bot, update, dbu)

			case "🤍White Hoodie🤍":
				commandimpl.WhiteHoodieCommand(userMap, bot.Bot, update)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_WhiteHoodieMenu:
			switch update.Message.Text {

			case "Добавить в корзину":
				commandimpl.SizeHoodie(userMap, bot.Bot, update)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_BlackHoodieMenu:
			switch update.Message.Text {

			case "Добавить в корзину":
				commandimpl.SizeHoodie(userMap, bot.Bot, update)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_BlackSizeHoodie:
			switch update.Message.Text {

			case "S-46 (EUR)", "M-48 (EUR)", "L-50 (EUR)", "XL-52 (EUR)":
				commandimpl.NewOrder(userMap, bot.Bot, update, dbu, update.Message.Text)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_WhiteSizeHoodie:
			switch update.Message.Text {

			case "S-46 (EUR)", "M-48 (EUR)", "L-50 (EUR)", "XL-52 (EUR)":
				commandimpl.NewOrder(userMap, bot.Bot, update, dbu, update.Message.Text)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_AddProduct:
			switch update.Message.Text {
			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)
			case "Корзина":
				commandimpl.CartMenu(userMap, bot.Bot, update, cfg, dbu)

			}
		case domain.Location_Delivery:
			switch update.Message.Text {

			case "Самовывоз":
				commandimpl.DeliveryPoint(userMap, bot.Bot, update, dbu)

			case "Курьером СДЭК":
				commandimpl.DeliveryCourier(userMap, bot.Bot, update, dbu)

			case "На пункт выдачи СДЭК":
				commandimpl.DeliveryPointSdek(userMap, bot.Bot, update, dbu)

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_SendContact:
			switch update.Message.Text {

			case "◀️Назад":
				commandimpl.Back(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		case domain.Location_DeliveryPoint:
			switch update.Message.Text {

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

			case "Главное меню":
				commandimpl.BackToMenu(userMap, bot.Bot, update)

			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}

		default:
			userMap[update.Message.From.ID] = domain.Location_StartMenu
			switch update.Message.Text {

			case "/start":
				commandimpl.Start(userMap, bot.Bot, update, dbu)
			default:
				commandimpl.Undefined(userMap, bot.Bot, update)
			}
		}
	}
}
