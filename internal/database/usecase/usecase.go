package usecase

import "bot/internal/database/domain"

type DataBaseUsecase struct {
	Users []domain.User //база данных в которой находятся все юзеры
	Carts []domain.Cart //база данных в которой находятся все корзины
}

func InitDataBaseUsecase() *DataBaseUsecase {
	dbu := new(DataBaseUsecase)
	return dbu
}
func (dbu *DataBaseUsecase) AddUser(u domain.User) {
	for _, user := range dbu.Users {
		if u.Id == user.Id {
			return
		}
	}
	dbu.Users = append(dbu.Users, u)
}
func (dbu *DataBaseUsecase) AddCart(c domain.Cart) {
	for _, cart := range dbu.Carts {
		if c.Id == cart.Id {
			return
		}
	}
	dbu.Carts = append(dbu.Carts, c)
}
func (dbu *DataBaseUsecase) UpdateCart(c domain.Cart) {
	for i, cart := range dbu.Carts {
		if c.Id == cart.Id {
			cart = c
			dbu.Carts = append(dbu.Carts, dbu.Carts[i:]...)
			return
		}

	}
}
func (dbu *DataBaseUsecase) UpdateUser(u domain.User, id int) {
	for _, user := range dbu.Users {
		if user.Id == id {
			user = u
		}
	}
}

//TODO добавить апдейт юзера для записи данных для номера телефона
//TODO добавить метод получение корзины юзера из бд
//TODO func update user cart
