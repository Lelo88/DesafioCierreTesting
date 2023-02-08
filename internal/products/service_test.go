package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*type serviceStub struct{
	ID string
	ProductMem []Product
	Error error
}

func (s *serviceStub) GetAllBySeller(id string) ([]Product, error) {
	return s.ProductMem, s.Error
}

func NewServiceStub(s *serviceStub) Service{
	return &service{repo: s}
}*/

//stubs
func TestGetAllBySeller(t *testing.T) {
	
	t.Run("ok", func(t *testing.T) {
	//arrange
	repo := NewFakeRepository()
	repo.ProductMem = []Product{{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
	},}

	ser := NewService(repo)
	
	expResult := []Product{{
		ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
	},}

	newResult,err := ser.GetAllBySeller("FEX112AC")
	
	assert.NoError(t, err)
	assert.Equal(t, expResult, newResult)
		/*ser := NewServiceStub(&serviceStub{
			ProductMem: products,
			Error: nil,
		})

		listproducts, err := ser.GetAllBySeller(product.ID)
		assert.NoError(t,err)
		assert.Equal(t, 1,len(listproducts))
		assert.Equal(t, products, listproducts)*/
	})

	t.Run("GetError", func(t *testing.T) {
		
		repo := NewFakeRepository()
		repo.Error = ErrInternal

		ser := NewService(repo)


		_,err :=ser.GetAllBySeller("FEX112AC")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInternal))
		//errInternal:=errors.New("error in repository sellerId: hola")
		//productEmpty := Product{}
		/*		ser:=NewServiceStub(&serviceStub{
			ProductMem: products,
			Error: ErrInternal,
		})

		_, err := ser.GetAllBySeller("hola")
		
		assert.Error(t, err)
		assert.True(t,errors.Is(ErrInternal,err))
		//assert.NotEqual(t, []Product{}, listproducts)*/
	})
}


