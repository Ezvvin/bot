package usecase

import (
	"bot/internal/database/domain"
)
// база данных
type DataBaseUsecase struct {
	Users []domain.User //база данных в которой находятся все юзеры
	Carts []domain.Cart //база данных в которой находятся все корзины
}

// создание бд
func InitDataBaseUsecase() *DataBaseUsecase {
	dbu := new(DataBaseUsecase)
	return dbu
}

// создание нового юзера в бд
func (dbu *DataBaseUsecase) AddUser(u domain.User) {
	for _, user := range dbu.Users {
		if u.Id == user.Id {
			return
		}
	}
	dbu.Users = append(dbu.Users, u)
}

// создание новой корзины в бд
func (dbu *DataBaseUsecase) AddCart(c domain.Cart) {
	for _, cart := range dbu.Carts {
		if c.Id == cart.Id {
			return
		}
	}
	dbu.Carts = append(dbu.Carts, c)
}

// обавление продукта в корзину в бд
func (dbu *DataBaseUsecase) UpdateUserCart(p domain.Product, user domain.User) {
	for i, cart := range dbu.Carts {
		if cart.Id == user.UserCart.Id {
			cart.AddProduct(p)
			dbu.Carts[i] = cart
		}
	}
}

// апдейт корзины в  бд юбзеров
func (dbu *DataBaseUsecase) UpdateUser(u domain.User) {
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			user.UserCart = dbu.Carts[i]
			dbu.Users[i] = user
		}
	}
}

// апдейт корзины в бд корзин
func (dbu *DataBaseUsecase) UpdateCart(c domain.Cart) {
	for i, cart := range dbu.Carts {
		if cart.Id == c.Id {
			dbu.Carts[i] = c
		}
	}
}

// Очистка корзины юзера в бд юзеров
func (dbu *DataBaseUsecase) ClearUserCart(u domain.User) {
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			user.UserCart.Products = nil
			dbu.Users[i] = user
		}
	}
}

// Очистка корзины  в бд корзин
func (dbu *DataBaseUsecase) ClearCart(u domain.User) {
	for i, cart := range dbu.Carts {
		if cart.Id == u.Id {
			cart.Products = nil
			dbu.Carts[i] = cart
		}
	}
}
