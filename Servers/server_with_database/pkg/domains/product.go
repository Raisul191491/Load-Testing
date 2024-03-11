package domains

import "github.com/gin-gonic/gin"

type ProductController interface {
	AddProduct(g *gin.Context)
	GetProduct(g *gin.Context)
}
