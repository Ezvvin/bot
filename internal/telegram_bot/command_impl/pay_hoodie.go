package commandimpl

import (
	"bot/internal/domain"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PayHoodie(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg domain.BotConfig) {

	msg := tgbotapi.NewInvoice(update.Message.Chat.ID, "🤍WHITELOGO LÚQ HOODIE🤍", "Заказ номер №0001", "ЧТОКУДАЭТО", cfg.PayToken, "standart", "RUB", []tgbotapi.LabeledPrice{{Label: "БЕЛЫЙ LUQ ХУДИ", Amount: 350000}})
	msg.NeedName = true
	msg.NeedShippingAddress = false
	msg.NeedPhoneNumber = true
	msg.SuggestedTipAmounts = []int{} // ввод сумм чаевых
	// msg.MaxTipAmount = 100000 ввод максимальной суммы чаевых
	msg.PhotoURL = "https://sun9-10.userapi.com/impg/-JuLUtuq1ZI_9eubyAoe1RzpHqhoMAZJgNQoBg/IuWP_-PD9wU.jpg?size=1280x1280&quality=95&sign=dafb57631e851fd63af53f800674c99c&type=album"
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "pay hoodie")
	}
	userMap[update.Message.From.ID] = domain.Location_AcceptDelivery
}
