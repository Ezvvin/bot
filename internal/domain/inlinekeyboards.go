package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
// TODO: Переименуй файл "inline_keyboards"
// инлайн кнопки контактов
var ContactsMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Мы в Вконтакте", "https://vk.com/luq.wear"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Мы в Instagram", "https://instagram.com/luq.wear"),
	),
	
) 
