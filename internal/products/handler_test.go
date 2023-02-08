package products

import "github.com/gin-gonic/gin"


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

