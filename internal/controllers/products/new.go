package products

import "github.com/phuc-create/go-simple-crud/internal/repository"

type ProductControllers interface {
	GetAllProducts() error
}

func New(repo repository.Registry) ProductControllers {
	return implement{repo}
}

type implement struct {
	repo repository.Registry
}
