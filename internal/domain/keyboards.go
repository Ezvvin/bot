package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	// Стартовая меню
	StartKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
		),
	)
	// Главное меню
	MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог одежды🥼"),
			tgbotapi.NewKeyboardButton("Отзывы🔥"),
		),

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Магазин LÚQ"),
			tgbotapi.NewKeyboardButton("Корзина"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Поддержка"),
		),
	)
	// Меню корзины с товарами
	CartMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Сделать заказ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Удалить товар"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог одежды🥼"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// Меню корзины без товаров
	CartMenuKeyboardIfNil = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог одежды🥼"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// Меню выбора худи
	HoodyMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🖤Black Hoodie🖤"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🤍White Hoodie🤍"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// Меню информации о худи
	AddProductKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Добавить в корзину"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
			tgbotapi.NewKeyboardButton("Главное меню"),
		),
	)
	// Меню выбора размера худи
	SizeHoodieKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("S-46 (EUR)"),
			tgbotapi.NewKeyboardButton("M-48 (EUR)"),
			tgbotapi.NewKeyboardButton("L-50 (EUR)"),
			tgbotapi.NewKeyboardButton("XL-52 (EUR)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
			tgbotapi.NewKeyboardButton("Главное меню"),
		),
	)
	// Возвращение после добавление товара в заказ
	BackFromAddProductKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Главное меню"),
			tgbotapi.NewKeyboardButton("Корзина"),
		),
	)
	// меню выбора доставки
	DeliveryKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Самовывоз"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Курьером СДЭК"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("На пункт выдачи СДЭК"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
			tgbotapi.NewKeyboardButton("Главное меню"),
		),
	)
	// Запрос контактов
	SendContactKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Отправить контакт"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// кнопка назад
	BackLocationKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// Оплата заказа
	PayKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Оплатить заказ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
			tgbotapi.NewKeyboardButton("Главное меню"),
		),
	)
)
