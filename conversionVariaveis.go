package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Definindo uma variável string com um valor numérico
	str := "123"

	// Convertendo a string para int usando strconv.Atoi
	num, err := strconv.Atoi(str)

	// Verificando se houve erro na conversão
	if err != nil {
		fmt.Println("Erro na conversão:", err)
	} else {
		fmt.Println("Número convertido:", num)
	}
}
