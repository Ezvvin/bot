package main

import (
	"bot/internal/handler"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Создаем бота
func main() {
	bot, err := tgbotapi.NewBotAPI("5437936243:AAG8qSRApD7V90ZZUVDM4ze0bcRbUC1rbrE") //TODO скрыть ебаный токен

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)
	handler.Handler(bot, updates)
}
