package commandimpl

import (
	"bot/internal/domain"
	"log"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commends(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы можете прочитать или оставить отзыв о товаре в нашей группе Вконтакте:
	 url := 'https://vk.com/topic-200472710_49047107'")
			msg.ReplyMarkup = domain.CommendsKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			userMap[update.Message.From.ID] = domain.Location_Commends
}