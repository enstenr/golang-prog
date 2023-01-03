package routes

import (
	"github.com/enstenr/golang-prog/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup",controllers.Signup())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.searchProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProducByQuery())
}