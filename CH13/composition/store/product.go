package store

type Product struct {
	Name, Category string
	price float64
}

func (p *Product) Price(taxRate float64) float64{
	retur p.price + p.price * taxRate
}

func NewProduct(name, category string, price float64){
	return &Product{name, category, price}
}