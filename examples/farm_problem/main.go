// Package main demonstrates the practical application of Euclid's algorithm
// to solve the geometric problem from "Understanding Algorithms" book.
//
// The problem: "If you find the largest square that divides this segment,
// it will be the largest square that will divide the entire farm"
package main

import (
	"fmt"
	"strings"

	"github.com/JeanGrijp/go-datastructures/pkg/euclidean"
)

func main() {
	fmt.Println("🚜 Farm Division Problem Solver")
	fmt.Println("================================")
	fmt.Println()

	// Problem from "Understanding Algorithms" book
	fmt.Println("📖 From 'Understanding Algorithms':")
	fmt.Println("\"If you find the largest square that divides this segment,")
	fmt.Println(" it will be the largest square that will divide the entire farm\"")
	fmt.Println()

	// Example farms with different dimensions
	farms := [][]int{
		{1680, 1050}, // Original example from the book
		{100, 60},    // Smaller farm
		{24, 16},     // Even smaller
		{15, 10},     // Simple case
		{17, 13},     // Prime dimensions (only 1x1 squares)
	}

	for i, farm := range farms {
		width, height := farm[0], farm[1]
		fmt.Printf("Farm #%d: %d x %d meters\n", i+1, width, height)
		fmt.Println(strings.Repeat("-", 30))

		// Find the largest square
		squareSize := euclidean.LargestSquareSize(width, height)
		fmt.Printf("Largest square size: %d x %d meters\n", squareSize, squareSize)

		// Calculate how many squares are needed
		_, count := euclidean.SquareDivision(width, height)
		fmt.Printf("Total squares needed: %d\n", count)

		// Show the area calculations
		totalArea := width * height
		squareArea := squareSize * squareSize
		fmt.Printf("Farm area: %d m²\n", totalArea)
		fmt.Printf("Each square area: %d m²\n", squareArea)
		fmt.Printf("Verification: %d squares × %d m² = %d m²\n", count, squareArea, count*squareArea)

		// Visual representation for smaller farms
		if width <= 24 && height <= 24 {
			fmt.Println("\nVisual representation:")
			showFarmDivision(width, height, squareSize)
		}

		// Show the step-by-step algorithm
		fmt.Println("\nEuclid's Algorithm steps:")
		euclidean.Visualize(width, height)

		fmt.Println()
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println()
	}

	// Special case: What if we want to find optimal rectangular plots?
	fmt.Println("🌾 Bonus: Finding optimal rectangular plots")
	fmt.Println("==========================================")

	// Let's say we want to divide a farm into rectangular plots
	// that are as square as possible but not necessarily squares
	width, height := 60, 40
	fmt.Printf("Farm: %d x %d meters\n", width, height)

	gcd := euclidean.GCD(width, height)
	fmt.Printf("GCD: %d\n", gcd)

	// Different subdivision options
	divisors := findDivisors(gcd)
	fmt.Println("\nPossible plot sizes:")
	for _, divisor := range divisors {
		plotWidth := width / divisor
		plotHeight := height / divisor
		numPlots := divisor * divisor
		fmt.Printf("- %d plots of %d x %d meters each\n", numPlots, plotWidth, plotHeight)
	}
}

// showFarmDivision creates a visual representation of how squares divide the farm
func showFarmDivision(width, height, squareSize int) {
	// Calculate how many squares fit in each dimension
	squaresX := width / squareSize
	squaresY := height / squareSize

	// Create visual grid
	for y := 0; y < squaresY; y++ {
		// Top border of squares
		for x := 0; x < squaresX; x++ {
			fmt.Print("+")
			for i := 0; i < squareSize; i++ {
				fmt.Print("-")
			}
		}
		fmt.Println("+")

		// Square content
		for row := 0; row < squareSize; row++ {
			for x := 0; x < squaresX; x++ {
				fmt.Print("|")
				for i := 0; i < squareSize; i++ {
					fmt.Print(" ")
				}
			}
			fmt.Println("|")
		}
	}

	// Bottom border
	for x := 0; x < squaresX; x++ {
		fmt.Print("+")
		for i := 0; i < squareSize; i++ {
			fmt.Print("-")
		}
	}
	fmt.Println("+")
	fmt.Printf("Grid: %d x %d squares of size %d x %d\n", squaresX, squaresY, squareSize, squareSize)
}

// findDivisors finds all divisors of a number
func findDivisors(n int) []int {
	var divisors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
