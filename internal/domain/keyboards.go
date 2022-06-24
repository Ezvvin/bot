package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	// Главное меню
	MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог одежды"),
			tgbotapi.NewKeyboardButton("Отзывы(не работает)"),
		),
		
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Наш Вконтакте(не работает)"),
			tgbotapi.NewKeyboardButton("Наш Instagram(не работает)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Задать вопрос(не работает)"),
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
			tgbotapi.NewKeyboardButton("Black Hoodie"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("White Hoodie"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	// Меню покупки худи
	BuyHoodieKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Купить"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
)
