package main

import (
	"fmt"

	btree "github.com/JeanGrijp/go-datastructures/pkg/b-tree"
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
	sequence := fibonacci.FibonacciSequence(15)
	fmt.Printf("   Primeiros 15: %v\n", sequence)
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
	fmt.Println()

	// Demo do pacote B-Tree
	fmt.Println("4. B-Tree Package Demo:")
	fmt.Println()

	// Criando uma B-Tree com grau mínimo 2 (2-3-4 tree)
	fmt.Println("   🌳 Criando B-Tree com grau mínimo t=2:")
	bt := btree.NewBTree(2)

	// Inserindo chaves
	keysToInsert := []int{10, 20, 5, 6, 12, 30, 7, 17, 3, 8}
	fmt.Printf("   Inserindo chaves: %v\n", keysToInsert)
	for _, k := range keysToInsert {
		bt.Insert(k)
	}
	fmt.Println("   ✅ Todas as chaves inseridas!")
	fmt.Println()

	// Buscando chaves
	fmt.Println("   🔍 Buscando chaves:")
	searchKeys := []int{6, 15, 30, 100}
	for _, k := range searchKeys {
		if bt.Search(k) {
			fmt.Printf("   Chave %d: ✅ encontrada\n", k)
		} else {
			fmt.Printf("   Chave %d: ❌ não encontrada\n", k)
		}
	}
	fmt.Println()

	// Removendo chaves
	fmt.Println("   🗑️  Removendo chaves:")
	keysToRemove := []int{6, 30, 3}
	for _, k := range keysToRemove {
		fmt.Printf("   Removendo chave %d...\n", k)
		bt.Remove(k)
	}
	fmt.Println()

	// Verificando após remoção
	fmt.Println("   🔍 Verificando após remoção:")
	checkKeys := []int{6, 30, 3, 10, 20}
	for _, k := range checkKeys {
		if bt.Search(k) {
			fmt.Printf("   Chave %d: ✅ ainda presente\n", k)
		} else {
			fmt.Printf("   Chave %d: ❌ removida/não existe\n", k)
		}
	}
	fmt.Println()

	// Exemplo com grau mínimo maior (mais eficiente para grandes volumes)
	fmt.Println("   📊 Exemplo com grau mínimo t=50 (típico para databases):")
	btLarge := btree.NewBTree(50)

	// Inserindo 1000 chaves
	for i := 1; i <= 1000; i++ {
		btLarge.Insert(i)
	}
	fmt.Println("   ✅ 1000 chaves inseridas!")

	// Verificando algumas chaves
	testKeys := []int{1, 500, 1000, 1001}
	for _, k := range testKeys {
		if btLarge.Search(k) {
			fmt.Printf("   Chave %d: ✅ encontrada\n", k)
		} else {
			fmt.Printf("   Chave %d: ❌ não encontrada\n", k)
		}
	}
	fmt.Println()

	// Demonstração de uso típico: índice de banco de dados
	fmt.Println("   💾 Simulação de índice de banco de dados:")
	dbIndex := btree.NewBTree(100)

	// Simulando IDs de registros
	recordIDs := []int{1001, 2045, 3089, 4023, 5067, 6011, 7055, 8099}
	fmt.Printf("   Indexando registros: %v\n", recordIDs)
	for _, id := range recordIDs {
		dbIndex.Insert(id)
	}

	// Buscando um registro
	searchID := 3089
	if dbIndex.Search(searchID) {
		fmt.Printf("   Registro #%d: ✅ encontrado no índice\n", searchID)
	}

	// Removendo um registro deletado
	deleteID := 2045
	dbIndex.Remove(deleteID)
	fmt.Printf("   Registro #%d removido do índice\n", deleteID)

	if !dbIndex.Search(deleteID) {
		fmt.Printf("   Registro #%d: ❌ não está mais no índice\n", deleteID)
	}
	fmt.Println()

	fmt.Println("=== Demo Completo! ===")
}
