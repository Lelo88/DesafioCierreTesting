package products

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func CreateProductServer(rp Repository) *gin.Engine{
	ser := NewService(rp)

	server := gin.Default()

	router := server.Group("/api/v1")
	{
		h:= NewHandler(ser)
		group := router.Group("/products")
		group.GET("", h.GetProducts)
	}

	return server
}

func NewRequest(method, path, body string) (req *http.Request, res *httptest.ResponseRecorder) {
	// request
	req = httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	// response
	res = httptest.NewRecorder()

	return
}

/*type productsHandler struct{
	ID          string  `json:"id"`
	SellerID    string  `json:"seller_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}*/


func TestGetProducts(t *testing.T){

	/*repo := NewServiceStub()
	serv := CreateProductServer(repo)*/
	repo:= NewFakeRepository()
	serv := CreateProductServer(repo)
	t.Run("GetProducts OK", func(t *testing.T) {
		repo.Reset()
		code := http.StatusOK
		repo.ProductMem = []Product{{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
		},}

		expResult := []Product{{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},}

		req, res := NewRequest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC", "")
		serv.ServeHTTP(res, req)

		var resp []Product
		err := json.Unmarshal(res.Body.Bytes(), &resp)

		assert.NoError(t,err)
		assert.Equal(t, code, res.Code)
		assert.Equal(t, expResult, resp)

	})

	t.Run("GetProducts empty", func(t *testing.T) {
		repo.Reset()
		code := http.StatusBadRequest

		req, res := NewRequest(http.MethodGet, "/api/v1/products?seller_id=", "")
		
		serv.ServeHTTP(res, req)
		var resp = []Product{}
		err := json.Unmarshal(res.Body.Bytes(), &resp)
		assert.Error(t, err)
		assert.Equal(t, code, res.Code)
	})	

	t.Run("GetProducts errorserver", func(t *testing.T) {
		repo.Reset()
		code := http.StatusInternalServerError
		repo.Error = errors.New("Error")

		req, res := NewRequest(http.MethodGet, "/api/v1/products?seller_id=adsasd", "")
		
		serv.ServeHTTP(res, req)
		var resp []Product
		err := json.Unmarshal(res.Body.Bytes(), &resp)
		assert.Error(t, err)
		assert.Equal(t, code, res.Code)
	})
}