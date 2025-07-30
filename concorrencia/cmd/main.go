package main

import (
	"fmt"
	"pizzaria/concorrencia/cmd/internal/fetcher"
	"pizzaria/concorrencia/cmd/internal/models"
	"pizzaria/concorrencia/cmd/internal/processor"
	"time"
)

func main() {

	start := time.Now()
	//Buffer = 4
	priceChannel := make(chan models.PriceDetail, 4)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	//Atua como um mecanismo de sincronização e conclusão
	<-done

	endTime := time.Now()
	delta := endTime.Sub(start)

	fmt.Printf("Tempo total de Execução do Programa: %.2f", delta.Seconds())

}