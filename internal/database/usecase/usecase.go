package usecase

import "bot/internal/database/domain"

type DataBaseUsecase struct {
	Users []domain.User //база данных в которой находятся все юзеры
	Carts []domain.Cart //база данных в которой находятся все корзины
}

func InitDataBaseUsecase() *DataBaseUsecase { // создание бд
	dbu := new(DataBaseUsecase)
	return dbu
}
func (dbu *DataBaseUsecase) AddUser(u domain.User) { // создание нового юзера в бд
	for _, user := range dbu.Users {
		if u.Id == user.Id {
			return
		}
	}
	dbu.Users = append(dbu.Users, u)
}
func (dbu *DataBaseUsecase) AddCart(c domain.Cart) { // создание новой корзины в бд
	for _, cart := range dbu.Carts {
		if c.Id == cart.Id {
			return
		}
	}
	dbu.Carts = append(dbu.Carts, c)
}
func (dbu *DataBaseUsecase) UpdateUserCart(p domain.Product, user domain.User) { // апдейт  корзиныв бд
	for i, cart := range dbu.Carts {
		if cart.Id == user.UserCart.Id {
			cart.AddProduct(p)
			dbu.Carts[i] = cart
		}
	}
}
func (dbu *DataBaseUsecase) UpdateUser(u domain.User) { // апдейт корзины в  бд юбзеров
	for i, user := range dbu.Users {
		if user.Id == u.Id {
			user.UserCart = dbu.Carts[i]
			dbu.Users[i] = user
		}
	}
}

//TODO добавить апдейт юзера для записи данных для номера телефона
//TODO добавить метод получение корзины юзера из бд ()
//TODO func update user cart
