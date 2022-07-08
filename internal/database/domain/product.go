package domain

type Product struct {
	ProductId int
	Name      string
	Price     int
	Size      ProductSize
}

type ProductSize string

const (
	ProductSize_S ProductSize = "S"
	ProductSize_M ProductSize = "M"
	ProductSize_L ProductSize = "L"
)
