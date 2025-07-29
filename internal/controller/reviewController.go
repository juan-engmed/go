package controller

import (
	"fmt"
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newReview models.Review
	if err := c.ShouldBindBodyWithJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidateRating(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for index, pizza := range data.Pizzas {
		fmt.Println(index)
		if pizza.ID == id {
			pizza.Review = append(pizza.Review, newReview)

			data.Pizzas[index] = pizza
			data.WritePizza()
			c.JSON(http.StatusCreated, gin.H{"pizza": data.Pizzas})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}
