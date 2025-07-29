package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("internal/data/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder((file))
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func WritePizza() {
	file, err := os.Create("internal/data/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&Pizzas); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}
}
