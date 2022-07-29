package domain

type User struct {
	Id        int    //Айди пользователя
	UserCart  Cart   //Личная корзина пользователя
	UserName  string //Никнейм пользователя
	FirstName string //Имя пользователя
	LastName  string //Фамилия пользователя
	Phone     string //Телефон пользователя
	Adress    Adress //Адрес пользователя
	Delivery  string //Вид доставки пользователя
}
type Adress struct {
	Latitude  float64
	Longitude float64
}
