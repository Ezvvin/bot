package commandimpl

import (
	"bot/internal/domain"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commends(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	url := "https://vk.com/topic-200472710_49047107"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Вы можете прочитать или оставить отзыв о товаре в нашей группе Вконтакте: [VKontakte](%s)", url))
	msg.ReplyMarkup = domain.CommendsKeyboard
	// Меняю мод парсера для гипперссылок
	msg.ParseMode = "MarkdownV2"
	// Отключаю превью страницы
	msg.DisableWebPagePreview = true
	if _, err := bot.Send(msg); err != nil {
		// TODO: Сделай через нормальный логгер
		log.Panic(err)
	}
	userMap[update.Message.From.ID] = domain.Location_Commends
}
