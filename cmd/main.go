package main

import (
	"fmt"

	btree "github.com/JeanGrijp/go-datastructures/pkg/btree"
	"github.com/JeanGrijp/go-datastructures/pkg/euclidean"
	"github.com/JeanGrijp/go-datastructures/pkg/factorial"
	"github.com/JeanGrijp/go-datastructures/pkg/fibonacci"
)

func main() {
	fmt.Println("=== Go Data Structures Demo ===")
	fmt.Println()

	// Factorial package demo
	fmt.Println("1. Factorial Package Demo:")
	factorialValue := factorial.Factorial(3000)
	factorialString := factorialValue.String()
	fmt.Printf("   3000! has %d digits\n", len(factorialString))
	fmt.Println()

	// Euclidean package demo
	fmt.Println("2. Euclidean Algorithm Package Demo:")

	// Basic GCD example
	fmt.Printf("   GCD(48, 18) = %d\n", euclidean.GCD(48, 18))

	// Farm problem from the book "Grokking Algorithms"
	fmt.Println("\n   📖 Example from the book 'Grokking Algorithms':")
	fmt.Println("   'If you find the largest square that divides this segment,")
	fmt.Println("    it will be the largest square that will divide the entire farm'")

	width, height := 1680, 1050
	squareSize := euclidean.LargestSquareSize(width, height)
	fmt.Printf("   Farm size: %dx%d meters\n", width, height)
	fmt.Printf("   Largest square that divides the farm: %dx%d meters\n", squareSize, squareSize)

	size, count := euclidean.SquareDivision(width, height)
	fmt.Printf("   The farm can be divided into %d squares of %dx%d meters\n", count, size, size)
	fmt.Println()

	// Visual algorithm walkthrough
	fmt.Println("   🔍 Euclidean Algorithm Visualization:")
	euclidean.Visualize(48, 18)

	// Coprime check
	fmt.Println()
	fmt.Println("   🔢 Coprime Number Check:")
	fmt.Printf("   Are 15 and 28 coprime? %t\n", euclidean.IsCoprime(15, 28))
	fmt.Printf("   Are 15 and 25 coprime? %t\n", euclidean.IsCoprime(15, 25))

	// Extended Euclidean algorithm
	fmt.Println()
	fmt.Println("   ➕ Extended Euclidean Algorithm:")
	gcd, x, y := euclidean.ExtendedGCD(30, 18)
	fmt.Printf("   30×%d + 18×%d = %d (GCD)\n", x, y, gcd)

	// LCM (Least Common Multiple)
	fmt.Println()
	fmt.Printf("   📊 LCM(12, 18) = %d\n", euclidean.LCM(12, 18))

	// GCD of multiple numbers
	numbers := []int{48, 18, 24, 30}
	fmt.Printf("   GCD of %v = %d\n", numbers, euclidean.GCDMultiple(numbers))
	fmt.Println()

	// Fibonacci package demo
	fmt.Println("3. Fibonacci Sequence Package Demo:")

	// Compare different implementations
	fmt.Println("   🚀 Different Fibonacci implementations:")
	n := 10
	fmt.Printf("   F(%d) Iterative: %d\n", n, fibonacci.Fibonacci(n))
	fmt.Printf("   F(%d) Recursive: %d\n", n, fibonacci.FibonacciRecursive(n))
	fmt.Printf("   F(%d) Memoized:  %d\n", n, fibonacci.FibonacciMemoized(n))
	fmt.Printf("   F(%d) Matrix:    %d\n", n, fibonacci.FibonacciMatrix(n))
	fmt.Println()

	// Fibonacci sequence
	fmt.Println("   📊 Fibonacci Sequence:")
	sequence := fibonacci.FibonacciSequence(15)
	fmt.Printf("   First 15: %v\n", sequence)
	fmt.Println()

	// Large numbers with big.Int
	fmt.Println("   🔢 Large numbers:")
	bigFib := fibonacci.FibonacciBig(100)
	fmt.Printf("   F(100) = %s\n", bigFib.String())
	fmt.Printf("   F(100) has %d digits\n", len(bigFib.String()))
	fmt.Println()

	// Educational helpers
	fmt.Println("   🧮 Educational helpers:")
	fmt.Printf("   Is 21 a Fibonacci number? %t\n", fibonacci.IsValidFibonacci(21))
	fmt.Printf("   Is 22 a Fibonacci number? %t\n", fibonacci.IsValidFibonacci(22))
	fmt.Printf("   Index of number 55: %d\n", fibonacci.FibonacciIndex(55))
	fmt.Printf("   Sum of first 10 numbers: %d\n", fibonacci.FibonacciSum(10))
	fmt.Printf("   Golden ratio (approx): %.6f\n", fibonacci.GoldenRatio(20))
	fmt.Println()

	// B-Tree package demo
	fmt.Println("4. B-Tree Package Demo:")
	fmt.Println()

	// Create a B-Tree with minimum degree 2 (2-3-4 tree)
	fmt.Println("   🌳 Creating B-Tree with minimum degree t=2:")
	bt := btree.NewBTree(2)

	// Insert keys
	keysToInsert := []int{10, 20, 5, 6, 12, 30, 7, 17, 3, 8}
	fmt.Printf("   Inserting keys: %v\n", keysToInsert)
	for _, k := range keysToInsert {
		bt.Insert(k)
	}
	fmt.Println("   ✅ All keys inserted!")
	fmt.Println()

	// Search keys
	fmt.Println("   🔍 Searching keys:")
	searchKeys := []int{6, 15, 30, 100}
	for _, k := range searchKeys {
		if bt.Search(k) {
			fmt.Printf("   Key %d: ✅ found\n", k)
		} else {
			fmt.Printf("   Key %d: ❌ not found\n", k)
		}
	}
	fmt.Println()

	// Remove keys
	fmt.Println("   🗑️  Removing keys:")
	keysToRemove := []int{6, 30, 3}
	for _, k := range keysToRemove {
		fmt.Printf("   Removing key %d...\n", k)
		bt.Remove(k)
	}
	fmt.Println()

	// Check after removal
	fmt.Println("   🔍 Verifying after removal:")
	checkKeys := []int{6, 30, 3, 10, 20}
	for _, k := range checkKeys {
		if bt.Search(k) {
			fmt.Printf("   Key %d: ✅ still present\n", k)
		} else {
			fmt.Printf("   Key %d: ❌ removed/not found\n", k)
		}
	}
	fmt.Println()

	// Example with larger minimum degree (more efficient for high volumes)
	fmt.Println("   📊 Example with minimum degree t=50 (typical for databases):")
	btLarge := btree.NewBTree(50)

	// Insert 1000 keys
	for i := 1; i <= 1000; i++ {
		btLarge.Insert(i)
	}
	fmt.Println("   ✅ 1000 keys inserted!")

	// Check a few keys
	testKeys := []int{1, 500, 1000, 1001}
	for _, k := range testKeys {
		if btLarge.Search(k) {
			fmt.Printf("   Key %d: ✅ found\n", k)
		} else {
			fmt.Printf("   Key %d: ❌ not found\n", k)
		}
	}
	fmt.Println()

	// Typical use case: database index
	fmt.Println("   💾 Database index simulation:")
	dbIndex := btree.NewBTree(100)

	// Simulate record IDs
	recordIDs := []int{1001, 2045, 3089, 4023, 5067, 6011, 7055, 8099}
	fmt.Printf("   Indexing records: %v\n", recordIDs)
	for _, id := range recordIDs {
		dbIndex.Insert(id)
	}

	// Search a record
	searchID := 3089
	if dbIndex.Search(searchID) {
		fmt.Printf("   Record #%d: ✅ found in index\n", searchID)
	}

	// Remove a deleted record
	deleteID := 2045
	dbIndex.Remove(deleteID)
	fmt.Printf("   Record #%d removed from index\n", deleteID)

	if !dbIndex.Search(deleteID) {
		fmt.Printf("   Record #%d: ❌ no longer in index\n", deleteID)
	}
	fmt.Println()

	fmt.Println("=== Demo Complete! ===")
}
