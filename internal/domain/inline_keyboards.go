package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// инлайн кнопки контактов
var (
	ContactsMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("LÚQ в Вконтакте", "https://vk.com/luq.wear"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("LÚQ в Instagram", "https://instagram.com/luq.wear"),
		),
	)
	// инлайн ссылка на отзывы
	CommendsInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Отзывы", "https://vk.com/topic-200472710_49047107"),
		),
	)
)
