package domain

type User struct {
	Id        int
	UserCart  Cart			//личная корзина юзера
	Username  string
	FirstName string
	LastName  string
	Phone     string
	Adress    string
}
