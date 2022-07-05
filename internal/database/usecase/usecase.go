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

// TODO: Create functions for add user and user cart into database
