package domain

type Cart struct {
	Id         int 			// айди корзины (такой же как и айди юзера)
	Products   []Product  	// список продуктов
	TotalPrice int			// общая стоимость товаров в корзине
}

func (c *Cart) CalculateTP() int {
	for _, product := range c.Products {
		c.TotalPrice += product.Price
	}

	return c.TotalPrice
}

func (c *Cart) AddProduct(p Product) {
	c.Products = append(c.Products, p)
}
func (c *Cart) RemoveProduct(p Product) {
	for i, product := range c.Products {
		if p == product {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			return
		}
	}
}
