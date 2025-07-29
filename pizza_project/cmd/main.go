package main

import (
	"pizzaria/pizza_project/internal/data"
	"pizzaria/pizza_project/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", controller.GetPizzas)
	router.GET("/pizzas/:id", controller.GetPizzaById)
	router.POST("/pizzas", controller.CreatePizza)
	router.PUT("/pizzas/:id", controller.UpdatePizzaById)
	router.DELETE("/pizzas/:id", controller.DeletePizzaById)

	router.POST("pizzas/:id/reviews", controller.CreateReview)

	router.Run()
}
