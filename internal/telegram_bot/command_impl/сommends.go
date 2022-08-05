package commandimpl

import (
	"bot/internal/domain"
	"fmt"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commends(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	//присваиваем ссылку
	url := "https://vk.com/topic-200472710_49047107"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Вы можете прочитать или оставить отзыв о товаре в нашей группе Вконтакте: [Отзывы](%s)", url))
	msg.ReplyMarkup = domain.CommendsInline
	// Меняю мод парсера для гипперссылок
	msg.ParseMode = "MarkdownV2"
	// Отключаю превью страницы
	msg.DisableWebPagePreview = true

	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "commends")
	}
	userMap[update.Message.From.ID] = domain.Location_MainMenu
}
