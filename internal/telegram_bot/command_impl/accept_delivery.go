package commandimpl

import (
	db_domain "bot/internal/database/domain"
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func AcceptDelivery(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg domain.Config, dbu *db_usecase.DataBaseUsecase) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Спасибо за заказ! В течении 5 минут наш менеджер свяжется для подтверждения заказа!")
	msg.ReplyMarkup = domain.MainMenuKeyboard
	u := db_domain.User{Id: int(update.Message.From.ID)}
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			u = dbu.Users[i]
		}
	}
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "accepthoodiebutton")
	}
	msg = tgbotapi.NewMessage(cfg.AdminChat, fmt.Sprintf("%s\nНомер телефона:\nЗаказ: %v", update.Message.From.FirstName, u.UserCart.Products))
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "accepthoodiebutton")
	}
	userMap[update.Message.From.ID] = domain.Location_MainMenu
}
