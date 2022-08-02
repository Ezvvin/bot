package domain

type Cart struct {
	Id         int       // айди корзины (такой же как и айди юзера)
	Products   []Product // список продуктов
	TotalPrice int       // общая стоимость товаров в корзине
}

// Общая сумма корзины
func (c *Cart) CalculateTP() int {
	for _, product := range c.Products {
		c.TotalPrice += product.Price
	}

	return c.TotalPrice
}

// добавление продукта в корзину
func (c *Cart) AddProduct(p Product) {
	c.Products = append(c.Products, p)
}

// удаление продукта из корзины
func (c *Cart) RemoveProduct(index int) {
	for i := range c.Products {
		if index == i {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
		}
	}
}
