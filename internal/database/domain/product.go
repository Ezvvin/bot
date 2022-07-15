package domain

type Product struct {
	ProductId int			// айди продукта
	Name      string		// имя продукта
	Price     int			// стоимость продукта
	Size      ProductSize	// размер продукта
}

type ProductSize string

const (
	ProductSize_S ProductSize = "S"
	ProductSize_M ProductSize = "M"
	ProductSize_L ProductSize = "L"
)
