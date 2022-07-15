package domain

type User struct {
	Id        int    		//айди юзера
	UserCart  Cart   		//личная корзина юзера
	Username  string 		//никнейм юзера
	FirstName string 		//Имя юзера
	LastName  string 		//Фамилия юзера
	Phone     string 		//Телефон юзера
	Adress    string 		//Адрес юзера
}
