package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	})

	t.Run("GetError", func(t *testing.T) {
		
		repo := NewFakeRepository()
		repo.Error = ErrInternal

		ser := NewService(repo)


		_,err :=ser.GetAllBySeller("FEX112AC")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInternal))

	})
}


