package commandimpl

import (
	db_usecase "bot/internal/database/usecase"
	"bot/internal/domain"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	db_domain "bot/internal/database/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddProductInCart(userMap map[int64]domain.Location, bot *tgbotapi.BotAPI, update tgbotapi.Update, dbu *db_usecase.DataBaseUsecase, productInfo string) {
	productInfoArr := strings.Split(productInfo, " ")
	productId, err := strconv.Atoi(productInfoArr[3])
	if err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "addproductbutton")
		return
	}
	log.WithField("productId", productId).Debug()

	product := db_domain.Product{}
	flag := false
	for _, productFromList := range db_domain.ProductList {
		if productFromList.ProductId == productId {
			product = productFromList
			flag = true
		}
	}
	
	if !flag {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товара нет в наличии")
		msg.ReplyMarkup = domain.MainMenuKeyboard
		return
	}
	log.WithField("product in Flag", product).Debug()
	
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товар добавлен в корзину")
	msg.ReplyMarkup = domain.MainMenuKeyboard
	dbu.UpdateUserCart(product, int(update.Message.From.ID))
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "addproductbutton")
		return
	}
	userMap[update.Message.From.ID] = domain.Location_AddProductInCart

	log.WithField("product cart", dbu.Carts).Debug()

}
