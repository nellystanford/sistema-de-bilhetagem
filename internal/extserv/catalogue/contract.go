package extserv

type Catalogue struct {
	Tenant     string  `json:"tenant"`
	ProductSKU string  `json:"product"`
	Price      float64 `json:"price"`
	UseUnit    string  `json:"use_unit"`
}
