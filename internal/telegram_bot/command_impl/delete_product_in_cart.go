package commandimpl

import (
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func DeleteProduct(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите номер строки товара, который желаете удалить из корзины:")
	msg.ReplyMarkup = domain.BackLocationKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "cartbutton")
	}
	userMap[update.Message.From.ID] = domain.Location_DeleteProduct
}
