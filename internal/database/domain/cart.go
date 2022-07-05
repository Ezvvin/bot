package domain

type Cart struct {
	Id         int // for now use user id
	User       User
	Products   []Product
	TotalPrice int
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

// TODO: Add func for remove product from cart
