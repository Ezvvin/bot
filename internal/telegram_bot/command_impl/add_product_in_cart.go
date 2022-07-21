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
	log.WithField("productId", productId).Debug("ловим айди продукта")

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
	log.WithField("product", product).Debug("находим продукт по айди")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Товар добавлен в корзину")
	msg.ReplyMarkup = domain.MainMenuKeyboard
	user := db_domain.User{Id: int(update.Message.From.ID)}
	user.UserCart.Id = int(update.Message.From.ID)
	dbu.UpdateUserCart(product, user)
	dbu.UpdateUser(user)
	if _, err := bot.Send(msg); err != nil {
		log.WithError(err).Errorf(domain.ErrCommand_Init.Error(), "addproductbutton")
		return
	}

	userMap[update.Message.From.ID] = domain.Location_MainMenu

	log.WithField("user", user).Debug()
	log.WithField("бд юзеров", dbu.Users).Debug()
	log.WithField("бд корзин", dbu.Carts).Debug()

}
