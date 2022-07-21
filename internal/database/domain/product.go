package domain

type Product struct {
	ProductId int         // айди продукта
	Name      string      // имя продукта
	Price     int         // стоимость продукта
	Size      ProductSize // размер продукта
}

type ProductSize string

const (
	ProductSize_S  ProductSize = "S - 46 (EUR)"
	ProductSize_M  ProductSize = "M - 48 (EUR)"
	ProductSize_L  ProductSize = "L - 50 (EUR)"
	ProductSize_XL ProductSize = "XL - 52 (EUR)"
)
