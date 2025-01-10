package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DiaFaturamento struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func main() {
	file, err := os.Open("faturamento.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	var faturamento []DiaFaturamento
	if err := json.Unmarshal(bytes, &faturamento); err != nil {
		log.Fatalf("Erro ao parsear o JSON: %v", err)
	}

	var menor, maior float64 = -1, -1
	var soma, diasComFaturamento float64
	var diasAcimaMedia int

	for _, dia := range faturamento {
		if dia.Valor > 0 {
			if menor == -1 || dia.Valor < menor {
				menor = dia.Valor
			}
			if maior == -1 || dia.Valor > maior {
				maior = dia.Valor
			}
			soma += dia.Valor
			diasComFaturamento++
		}
	}

	media := soma / diasComFaturamento

	for _, dia := range faturamento {
		if dia.Valor > media {
			diasAcimaMedia++
		}
	}

	fmt.Printf("Menor valor de faturamento: %.2f\n", menor)
	fmt.Printf("Maior valor de faturamento: %.2f\n", maior)
	fmt.Printf("Número de dias com faturamento acima da média mensal: %d\n", diasAcimaMedia)
}
