package usecase

import "bot/internal/database/domain"

type DataBaseUsecase struct {
	Carts []domain.Cart
	Users []domain.User
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

//TODO добавить апдейт юзера для записи данных для номера телефона
