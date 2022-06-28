package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func VkontakteurlCommends(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	url := "https://vk.com/topic-200472710_49047107"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, url)
	msg.ReplyMarkup = domain.CommendsKeyboard
	msg.Entities = []tgbotapi.MessageEntity{tgbotapi.MessageEntity{Type: "url"}}
	if _, err := bot.Send(msg); err != nil {
		
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "vkcommendURL")
	}
	userMap[update.Message.From.ID] = domain.Location_Commends
}
