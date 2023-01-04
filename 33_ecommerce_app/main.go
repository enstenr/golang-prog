package main

import (
	"log"
	"os"

	"github.com/enstenr/golang-prog/controllers"
	"github.com/enstenr/golang-prog/database"
	"github.com/enstenr/golang-prog/middleware"
	"github.com/enstenr/golang-prog/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port:=os.Getenv("PORT")
	if port!=""{
		port="8001"
	}
	app:=controllers.NewApplication(database.ProductData(database.Client,"Product"),
									database.UserData(database.Client,"Users"),
								)

	router:=gin.New()
	router.Use(gin.Logger())

	routes.userRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart",app.AddToCart())
	router.GET("/removeitem",app.RemoveItem())
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("/instantbuy",app.InstantBuy())

	log.Fatal(router.Run(":"+port))
}