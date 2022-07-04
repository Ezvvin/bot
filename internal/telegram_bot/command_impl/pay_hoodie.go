package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PayHoodie(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewInvoice(update.Message.Chat.ID, "ЗАГОЛОВОК", "описание", "загрузка оплаты", "провайдертокен", "старт параметры", "USD", []tgbotapi.LabeledPrice{{Label: "USD", Amount: 200}})
	msg.NeedName = true
	msg.NeedShippingAddress = true
	msg.NeedPhoneNumber = true
	// msg.SuggestedTipAmounts = [3750, 37508026]
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "pay hoodie")
	}
	userMap[update.Message.From.ID] = domain.Location_AcceptDelivery
}
