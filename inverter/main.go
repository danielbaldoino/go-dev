package main

import "fmt"

func inverterString(str string) string {
	// Converte a string para um slice de runes (para tratar caracteres especiais corretamente)
	runes := []rune(str)
	var resultado []rune

	// Itera sobre o slice de runes de trás para frente
	for i := len(runes) - 1; i >= 0; i-- {
		resultado = append(resultado, runes[i])
	}

	// Converte o slice de volta para string
	return string(resultado)
}

func main() {
	var entrada string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&entrada)

	entradaInvertida := inverterString(entrada)
	fmt.Printf("String invertida: %s\n", entradaInvertida)
}
