package commandimpl

import (
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"

	db_domain "bot/internal/database/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func NewOrder(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase, sizeProduct string) {
	switch userMap[update.Message.From.ID] {
	case domain.Location_WhiteSizeHoodie:

		product := db_domain.Product{}
		flag := false
		for _, productFromList := range db_domain.ProductList {
			if productFromList.Size == sizeProduct && productFromList.Name == "White Hoodie" {
				product = productFromList
				flag = true
			}
		}

		if !flag {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товара нет в наличии")
			msg.ReplyMarkup = domain.MainMenuKeyboard
			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товар: добавлен в вашу корзину")
		msg.ReplyMarkup = domain.BackFromAddProductKeyboard
		user := db_domain.User{Id: int(update.Message.From.ID)}
		user.UserCart.Id = int(update.Message.From.ID)
		dbu.UpdateUserCart(product, user)
		dbu.UpdateUser(user)
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "orderbuttom")
		}
		userMap[update.Message.From.ID] = domain.Location_AddProduct

	case domain.Location_BlackSizeHoodie:

		product := db_domain.Product{}
		flag := false
		for _, productFromList := range db_domain.ProductList {
			if productFromList.Size == sizeProduct && productFromList.Name == "Black Hoodie" {
				product = productFromList
				flag = true
			}
		}

		if !flag {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товара нет в наличии")
			msg.ReplyMarkup = domain.MainMenuKeyboard
			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товар: добавлен в вашу корзину")
		msg.ReplyMarkup = domain.BackFromAddProductKeyboard
		user := db_domain.User{Id: int(update.Message.From.ID)}
		user.UserCart.Id = int(update.Message.From.ID)
		dbu.UpdateUserCart(product, user)
		dbu.UpdateUser(user)
		if _, err := bot.Send(msg); err != nil {
			log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "black order buttom")
		}
		userMap[update.Message.From.ID] = domain.Location_AddProduct

	}
}
