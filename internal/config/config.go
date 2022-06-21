package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог одежды"),
		),

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/close"),
		),
	)
	StartKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
		),
	)
	HoodyMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Black Neega"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("White Hodak"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	BlackNeegaKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Купить"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
)

type Location string

var (
	Location_HoodyCatalogMenu Location = "HoodyCatalogMenu"
	Location_MainMenu         Location = "MainMenu"
	Location_StartMenu        Location = "StartMenu"
	Location_HoodyMenu        Location = "HoodyMenu"
)
