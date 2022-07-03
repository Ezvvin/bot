package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	// Главное меню
	MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог одежды🥼"),
			tgbotapi.NewKeyboardButton("Отзывы🔥"),
		),

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Магазин LUQ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Вопрос/Ответ❓❗️"),
		),
	)
	// Стартовая меню
	StartKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
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
	InfoHoodieKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Выбрать размер📏"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// Меню выбора размера худи
	SizeHoodieKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("S-46 (EUR)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("M-48 (EUR)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("L-50 (EUR)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// TODO: Почему коммент просто "меню доставки", а в названии переменной фигурирует слово худи?
	// меню доставки
	DeliveryHoodieKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Доставка по адресу"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Самовывоз"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
	// Отзывы
	CommendsKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
)
