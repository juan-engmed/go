package controller

import (
	"fmt"
	"net/http"
	"pizzaria/pizza_project/internal/data"
	"pizzaria/pizza_project/internal/models"
	"pizzaria/pizza_project/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"pizzas": data.Pizzas})
}

func GetPizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for index, pizza := range data.Pizzas {
		fmt.Println(index)
		if pizza.ID == id {
			c.JSON(http.StatusOK, gin.H{"pizza": data.Pizzas})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Id não encontrado"})
}

func CreatePizza(c *gin.Context) {
	var newPizza models.Pizza

	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
	}

	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.WritePizza()
	c.JSON(http.StatusCreated, gin.H{"pizzas": data.Pizzas})
}

func UpdatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var updatePizzaById models.Pizza
	if err := c.ShouldBindBodyWithJSON(&updatePizzaById); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&updatePizzaById); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for index, pizza := range data.Pizzas {
		fmt.Println(index)
		if pizza.ID == id {
			data.Pizzas[index] = updatePizzaById
			data.Pizzas[index].ID = id
			data.WritePizza()
			c.JSON(http.StatusOK, gin.H{"pizza": data.Pizzas[index]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Id não encontrado"})
}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas = append(data.Pizzas[:index], data.Pizzas[index+1:]...)
			data.WritePizza()
			c.JSON(http.StatusOK, gin.H{"message": "pizza deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}
