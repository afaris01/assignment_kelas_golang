package router

import (
	"assignment-2/controllers"
	"assignment-2/database"

	"github.com/gin-gonic/gin"
)

func MulaiApp(){
	db := database.AmbilDB()

	DBConn := &controllers.DBConn{db}
	router := gin.Default()

	router.GET("/orders", DBConn.GetOrders)
	router.GET("/order/:id", DBConn.GetOrder)
	router.DELETE("/order/:id", DBConn.DeleteOrder)
	router.POST("/order", DBConn.CreateOrder)
	router.PUT("/order", DBConn.UpdatOrder)

	router.Run(":8080")
}