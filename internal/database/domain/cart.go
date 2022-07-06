package domain

type Cart struct {
	Id         int // for now use user id
	User       User
	Products   []Product //whiye hoodie black hoodie white
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
func (c *Cart) RemoveProduct(p Product, i int) {
	c.Products[i] = c.Products[len(c.Products)-1]
	c.Products = c.Products[:len(c.Products)-1]
}

// complete?
// TODO: Add func for remove product from cart
