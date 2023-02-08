package products
/*
import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
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

type productsHandler struct{
	Data []Product
	Message string
}


func TestGetProducts(t *testing.T){

	repo := NewServiceStub()
	serv := CreateProductServer(repo)

	t.Run("GetProducts OK", func(t *testing.T) {
		code := http.StatusOK
		
	})

	t.Run("GetProducts empty", func(t *testing.T) {

	})	

	t.Run("GetProducts errorserver", func(t *testing.T) {

	})
}*/