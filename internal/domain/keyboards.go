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
			tgbotapi.NewKeyboardButton("Наш Вконтакте"),
			tgbotapi.NewKeyboardButton("Наш Instagram"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Вопрос/Ответ❓❗️"),
			tgbotapi.NewKeyboardButton("⌨️"),
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
	// Меню покупки худи
	BuyHoodieKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Купить"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)

	CommendsKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Вконтакте"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("◀️Назад"),
		),
	)
)
