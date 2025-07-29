package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	router := gin.Default()
    router.GET("/pizzas", getPizzas)
    router.GET("/pizzas/:id", getPizzasById)
    router.POST("/pizzas", createPizza)
	router.Run()

}

func loadPizzas(){
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder((file))
	if err := decoder.Decode(&pizzas); err != nil{
		fmt.Println("Error decoding JSON: ", err)
	}
}

func writePizza(){
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&pizzas); err != nil{
		fmt.Println("Error encoding JSON: ", err)
	}
}

func getPizzas(c *gin.Context) {
        c.JSON(200, gin.H{"pizzas": pizzas,
        })
}

func getPizzasById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	for index, pizza := range pizzas {
		fmt.Println(index)
		if pizza.ID == id{
			c.JSON(200, gin.H{"pizza": pizza})
			return
		}
	}

	c.JSON(404, gin.H{"message": "Id não encontrado"})
}


func createPizza(c *gin.Context){
	var newPizza models.Pizza

	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"erro": err.Error(),
	})}

	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	writePizza()
	c.JSON(201, gin.H{"pizzas": pizzas})
}