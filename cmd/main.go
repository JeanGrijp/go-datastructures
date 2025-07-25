package main

import (
	"fmt"

	"github.com/JeanGrijp/go-datastructures/pkg/fatorial"
)

func main() {

	// Usando big.Int para números grandes
	lenFatorial := fatorial.Factorial(3000)
	stringLenFatorial := lenFatorial.String()
	fmt.Println("tamanho da string do fatorial de 3000:", len(stringLenFatorial))
}
