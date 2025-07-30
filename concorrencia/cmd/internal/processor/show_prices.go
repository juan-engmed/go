package processor

import (
	"fmt"
	"pizzaria/concorrencia/cmd/internal/models"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {

	var totalPrice float64
	countPrices := 0.0

	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("E-comm [%s] Preço recebido: R$ %.2f \n", price.StoreName, price.Value)
		fmt.Printf("Média de preços recebidos: R$ %.2f \n", avgPrice)
	}

	done <- true
}
