package main

import (
	"fmt"

	"github.com/JeanGrijp/go-datastructures/pkg/euclidean"
	"github.com/JeanGrijp/go-datastructures/pkg/fatorial"
	"github.com/JeanGrijp/go-datastructures/pkg/fibonacci"
)

func main() {
	fmt.Println("=== Go Data Structures Demo ===")
	fmt.Println()

	// Demo do pacote fatorial
	fmt.Println("1. Fatorial Package Demo:")
	lenFatorial := fatorial.Factorial(3000)
	stringLenFatorial := lenFatorial.String()
	fmt.Printf("   Fatorial de 3000 tem %d dígitos\n", len(stringLenFatorial))
	fmt.Println()

	// Demo do pacote euclidean
	fmt.Println("2. Euclidean Algorithm Package Demo:")

	// Exemplo básico de GCD
	fmt.Printf("   GCD(48, 18) = %d\n", euclidean.GCD(48, 18))

	// Problema da fazenda do livro "Understanding Algorithms"
	fmt.Println("\n   📖 Exemplo do livro 'Understanding Algorithms':")
	fmt.Println("   'If you find the largest square that divides this segment,")
	fmt.Println("    it will be the largest square that will divide the entire farm'")

	width, height := 1680, 1050
	squareSize := euclidean.LargestSquareSize(width, height)
	fmt.Printf("   Fazenda de %dx%d metros\n", width, height)
	fmt.Printf("   Maior quadrado que divide a fazenda: %dx%d metros\n", squareSize, squareSize)

	size, count := euclidean.SquareDivision(width, height)
	fmt.Printf("   A fazenda pode ser dividida em %d quadrados de %dx%d metros\n", count, size, size)
	fmt.Println()

	// Demonstração visual do algoritmo
	fmt.Println("   🔍 Visualização do Algoritmo de Euclides:")
	euclidean.Visualize(48, 18)

	// Teste de números coprimos
	fmt.Println()
	fmt.Println("   🔢 Teste de números coprimos:")
	fmt.Printf("   15 e 28 são coprimos? %t\n", euclidean.IsCoprime(15, 28))
	fmt.Printf("   15 e 25 são coprimos? %t\n", euclidean.IsCoprime(15, 25))

	// Algoritmo Euclidiano Estendido
	fmt.Println()
	fmt.Println("   ➕ Algoritmo Euclidiano Estendido:")
	gcd, x, y := euclidean.ExtendedGCD(30, 18)
	fmt.Printf("   30×%d + 18×%d = %d (GCD)\n", x, y, gcd)

	// LCM (Least Common Multiple)
	fmt.Println()
	fmt.Printf("   📊 MMC(12, 18) = %d\n", euclidean.LCM(12, 18))

	// GCD de múltiplos números
	numbers := []int{48, 18, 24, 30}
	fmt.Printf("   GCD de %v = %d\n", numbers, euclidean.GCDMultiple(numbers))
	fmt.Println()

	// Demo do pacote fibonacci
	fmt.Println("3. Fibonacci Sequence Package Demo:")

	// Comparação de diferentes implementações
	fmt.Println("   🚀 Diferentes implementações do Fibonacci:")
	n := 10
	fmt.Printf("   F(%d) Iterativo: %d\n", n, fibonacci.Fibonacci(n))
	fmt.Printf("   F(%d) Recursivo: %d\n", n, fibonacci.FibonacciRecursive(n))
	fmt.Printf("   F(%d) Memoizado: %d\n", n, fibonacci.FibonacciMemoized(n))
	fmt.Printf("   F(%d) Matriz:    %d\n", n, fibonacci.FibonacciMatrix(n))
	fmt.Println()

	// Sequência de Fibonacci
	fmt.Println("   📊 Sequência de Fibonacci:")
	sequence := fibonacci.FibonacciSequence(100)
	fmt.Printf("   Primeiros 100: %v\n", sequence)
	fmt.Println()

	// Números grandes com big.Int
	fmt.Println("   🔢 Números grandes:")
	bigFib := fibonacci.FibonacciBig(100)
	fmt.Printf("   F(100) = %s\n", bigFib.String())
	fmt.Printf("   F(100) tem %d dígitos\n", len(bigFib.String()))
	fmt.Println()

	// Ferramentas educacionais
	fmt.Println("   🧮 Ferramentas educacionais:")
	fmt.Printf("   21 é um número de Fibonacci? %t\n", fibonacci.IsValidFibonacci(21))
	fmt.Printf("   22 é um número de Fibonacci? %t\n", fibonacci.IsValidFibonacci(22))
	fmt.Printf("   Índice do número 55: %d\n", fibonacci.FibonacciIndex(55))
	fmt.Printf("   Soma dos primeiros 10: %d\n", fibonacci.FibonacciSum(10))
	fmt.Printf("   Razão áurea (aprox): %.6f\n", fibonacci.GoldenRatio(20))

	// Teste de Fibonacci Matriz
	fmt.Println()
	fmt.Println("   📊 Teste de Fibonacci Matriz:")
	n = 50
	fmt.Printf("   F(%d) Matriz: %d\n", n, fibonacci.FibonacciMatrix(n))

}
