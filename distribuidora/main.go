package main

import (
	"fmt"
)

func main() {
	faturamento := map[string]float64{
		"SP":     67836.43,
		"RJ":     36678.66,
		"MG":     29229.88,
		"ES":     27165.48,
		"Outros": 19849.53,
	}

	var totalFaturamento float64
	for _, valor := range faturamento {
		totalFaturamento += valor
	}

	fmt.Printf("Faturamento total: R$%.2f\n\n", totalFaturamento)
	for estado, valor := range faturamento {
		percentual := (valor / totalFaturamento) * 100
		fmt.Printf("Percentual de %s: %.2f%%\n", estado, percentual)
	}
}
