package domain

type User struct {
	Id        int
	UserCart  Cart
	Username  string
	FirstName string
	LastName  string
	Phone     string
	Adress    string
}
