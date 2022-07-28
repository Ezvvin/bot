package domain

type User struct {
	Id        int    //Айди пользователя
	UserCart  Cart   //Личная корзина пользователя
	UserName  string //Никнейм пользователя
	FirstName string //Имя пользователя
	LastName  string //Фамилия пользователя
	Phone     string //Телефон пользователя
	Adress    string //Адрес пользователя
	Delivery  string //Вид доставки пользователя
}
