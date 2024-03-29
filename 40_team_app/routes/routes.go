package routes

import (
	"github.com/enstenr/golang-prog/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/",controllers.Welcome())
	incomingRoutes.GET("/signup",controllers.SignUpLoad())
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	//incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	//incomingRoutes.GET("/users/productview",controllers.searchProduct())
	//incomingRoutes.GET("/users/search",controllers.SearchProducByQuery())
}