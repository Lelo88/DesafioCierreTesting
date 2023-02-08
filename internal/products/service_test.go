package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type serviceStub struct{
	ID string
	ProductMem []Product
	Error error
}

func (s *serviceStub) GetAllBySeller(id string) ([]Product, error) {
	return s.ProductMem, s.Error
}

func NewServiceStub(s *serviceStub) Service{
	return &service{repo: s}
}

//stubs
func TestGetAllBySeller(t *testing.T) {
	
	product := Product{
		ID: "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}

	products := []Product{product}

	t.Run("GetProductsBySeller", func(t *testing.T) {
		ser := NewServiceStub(&serviceStub{
			ProductMem: products,
			Error: nil,
		})

		listproducts, err := ser.GetAllBySeller(product.ID)
		assert.NoError(t,err)
		assert.Equal(t, 1,len(listproducts))
		assert.Equal(t, products, listproducts)
	})

	t.Run("GetError", func(t *testing.T) {

	})
}


