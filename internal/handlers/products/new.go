package products

import (
	"github.com/phuc-create/go-simple-crud/internal/controllers/products"
)

type ProductHandlers struct {
	prodCtrls products.ProductControllers
}

func New(prodCtrls products.ProductControllers) ProductHandlers {
	return ProductHandlers{prodCtrls}
}
