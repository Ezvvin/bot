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
func (dbu *DataBaseUsecase) AddUser(user domain.User) {
	dbu.Users = append(dbu.Users, user)
}
func (dbu *DataBaseUsecase) AddCart(cart domain.Cart) {
	dbu.Carts = append(dbu.Carts, cart)
}
