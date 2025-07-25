# Fibonacci Package

Esta implementação fornece várias abordagens para calcular números de Fibonacci, cada uma com suas próprias características de performance e casos de uso.

## Índice

- [Sobre a Sequência de Fibonacci](#sobre-a-sequência-de-fibonacci)
- [Implementações Disponíveis](#implementações-disponíveis)
- [Exemplos de Uso](#exemplos-de-uso)
- [Ferramentas Educacionais](#ferramentas-educacionais)
- [Performance](#performance)
- [Casos de Uso](#casos-de-uso)

## Sobre a Sequência de Fibonacci

A sequência de Fibonacci é uma série de números onde cada número é a soma dos dois números anteriores. A sequência tradicionalmente começa com 0 e 1:

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2) para n > 1
```

**Sequência:** 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377...

### Propriedades Matemáticas

- **Razão Áurea**: À medida que n aumenta, F(n+1)/F(n) aproxima-se da razão áurea φ ≈ 1.618033988749
- **Fórmula de Binet**: F(n) = (φⁿ - ψⁿ)/√5, onde ψ = (1-√5)/2
- **Identidade de Cassini**: F(n-1)×F(n+1) - F(n)² = (-1)ⁿ

## Implementações Disponíveis

### 1. Implementação Iterativa - `Fibonacci(n int) int`
- **Complexidade**: O(n) tempo, O(1) espaço
- **Uso recomendado**: Valores pequenos a médios (n < 50)
- **Características**: Mais eficiente para a maioria dos casos de uso

```go
result := fibonacci.Fibonacci(10) // Retorna 55
```

### 2. Implementação Recursiva - `FibonacciRecursive(n int) int`
- **Complexidade**: O(2ⁿ) tempo, O(n) espaço
- **Uso recomendado**: Apenas para demonstração educacional (n < 15)
- **Características**: Implementação mais intuitiva, mas extremamente ineficiente

```go
result := fibonacci.FibonacciRecursive(8) // Retorna 21
```

### 3. Implementação com Memoização - `FibonacciMemoized(n int) int`
- **Complexidade**: O(n) tempo, O(n) espaço
- **Uso recomendado**: Múltiplas chamadas para diferentes valores
- **Características**: Combina clareza recursiva com eficiência

```go
result := fibonacci.FibonacciMemoized(30) // Retorna 832040
```

### 4. Implementação com Números Grandes - `FibonacciBig(n int) *big.Int`
- **Complexidade**: O(n) tempo, O(1) espaço
- **Uso recomendado**: Valores muito grandes (n > 50)
- **Características**: Suporta números arbitrariamente grandes

```go
result := fibonacci.FibonacciBig(100) // Retorna 354224848179261915075
```

### 5. Implementação com Matriz - `FibonacciMatrix(n int) int`
- **Complexidade**: O(log n) tempo, O(1) espaço
- **Uso recomendado**: Valores grandes quando precisão int é suficiente
- **Características**: Algoritmo mais rápido para valores únicos grandes

```go
result := fibonacci.FibonacciMatrix(50) // Retorna 12586269025
```

## Exemplos de Uso

### Exemplo Básico
```go
package main

import (
    "fmt"
    "yourmodule/pkg/fibonacci"
)

func main() {
    // Calcular o 10º número de Fibonacci
    fmt.Printf("F(10) = %d\n", fibonacci.Fibonacci(10))
    
    // Gerar sequência dos primeiros 10 números
    sequence := fibonacci.FibonacciSequence(10)
    fmt.Printf("Primeiros 10: %v\n", sequence)
    
    // Verificar se um número é Fibonacci
    fmt.Printf("8 é Fibonacci? %t\n", fibonacci.IsValidFibonacci(8))
    
    // Calcular número muito grande
    big := fibonacci.FibonacciBig(200)
    fmt.Printf("F(200) = %s\n", big.String())
}
```

### Exemplo de Performance
```go
package main

import (
    "fmt"
    "time"
    "yourmodule/pkg/fibonacci"
)

func main() {
    n := 40
    
    // Comparar diferentes implementações
    start := time.Now()
    result1 := fibonacci.Fibonacci(n)
    fmt.Printf("Iterativo: F(%d) = %d, Tempo: %v\n", 
        n, result1, time.Since(start))
    
    start = time.Now()
    result2 := fibonacci.FibonacciMatrix(n)
    fmt.Printf("Matriz: F(%d) = %d, Tempo: %v\n", 
        n, result2, time.Since(start))
    
    start = time.Now()
    result3 := fibonacci.FibonacciMemoized(n)
    fmt.Printf("Memoizado: F(%d) = %d, Tempo: %v\n", 
        n, result3, time.Since(start))
}
```

## Ferramentas Educacionais

### Geração de Sequência - `FibonacciSequence(n int) []int`
Gera os primeiros n números da sequência de Fibonacci.

```go
sequence := fibonacci.FibonacciSequence(8)
// Retorna: [0, 1, 1, 2, 3, 5, 8, 13]
```

### Validação - `IsValidFibonacci(num int) bool`
Verifica se um número pertence à sequência de Fibonacci.

```go
fmt.Println(fibonacci.IsValidFibonacci(13)) // true
fmt.Println(fibonacci.IsValidFibonacci(14)) // false
```

### Busca de Índice - `FibonacciIndex(num int) int`
Encontra o índice de um número de Fibonacci na sequência.

```go
index := fibonacci.FibonacciIndex(55) // Retorna 10
```

### Razão Áurea - `GoldenRatio(n int) float64`
Calcula a aproximação da razão áurea usando F(n+1)/F(n).

```go
ratio := fibonacci.GoldenRatio(20) // Aproximadamente 1.618034
```

### Soma da Sequência - `FibonacciSum(n int) int`
Calcula a soma dos primeiros n números de Fibonacci.

```go
sum := fibonacci.FibonacciSum(6) // 0+1+1+2+3+5 = 12
```

### Visualização - `Visualize(n int) int`
Calcula e exibe o processo de cálculo (útil para debugging e educação).

```go
result := fibonacci.Visualize(5) // Mostra o processo de cálculo
```

## Performance

### Complexidade Temporal
| Implementação | Complexidade | Melhor Para |
|---------------|--------------|-------------|
| Iterativa | O(n) | Uso geral (n < 50) |
| Recursiva | O(2ⁿ) | Educação apenas (n < 15) |
| Memoizada | O(n) | Múltiplas chamadas |
| Matriz | O(log n) | Valores únicos grandes |
| Big Int | O(n) | Números muito grandes |

### Benchmarks (aproximados)
```
BenchmarkFibonacci-8               30000000        50.0 ns/op
BenchmarkFibonacciMemoized-8      10000000       150.0 ns/op  
BenchmarkFibonacciMatrix-8         5000000       300.0 ns/op
BenchmarkFibonacciBig-8            1000000      1500.0 ns/op
BenchmarkFibonacciRecursive-8           100  15000000.0 ns/op
```

## Casos de Uso

### 1. Educação e Demonstração
- Use `FibonacciRecursive()` para mostrar o conceito básico
- Use `Visualize()` para demonstrar o processo de cálculo
- Use `FibonacciSequence()` para mostrar a progressão

### 2. Aplicações de Performance
- Use `Fibonacci()` para valores pequenos a médios
- Use `FibonacciMatrix()` para cálculos únicos de valores grandes
- Use `FibonacciMemoized()` quando calculando múltiplos valores

### 3. Computação Científica
- Use `FibonacciBig()` para pesquisa matemática com números grandes
- Use `GoldenRatio()` para aproximações da razão áurea
- Use `FibonacciSum()` para análises estatísticas

### 4. Validação e Análise
- Use `IsValidFibonacci()` para verificar se dados seguem padrões Fibonacci
- Use `FibonacciIndex()` para encontrar posições na sequência

## Executando os Testes

```bash
# Executar todos os testes
go test ./pkg/fibonacci

# Executar testes com verbose
go test -v ./pkg/fibonacci

# Executar benchmarks
go test -bench=. ./pkg/fibonacci

# Executar testes de coverage
go test -cover ./pkg/fibonacci
```

## Notas sobre Limitações

1. **Overflow de Inteiros**: Implementações com `int` têm overflow após F(46) = 1836311903
2. **Performance Recursiva**: A implementação recursiva é impraticável para n > 15
3. **Memória**: A implementação memoizada usa O(n) memória adicional
4. **Precisão**: Para números muito grandes, use sempre `FibonacciBig()`

## Referências Matemáticas

- [Sequência de Fibonacci - Wikipedia](https://pt.wikipedia.org/wiki/Sequ%C3%AAncia_de_Fibonacci)
- [The Golden Ratio](https://en.wikipedia.org/wiki/Golden_ratio)
- [Matrix Form of Fibonacci Sequence](https://www.mathsisfun.com/algebra/matrix-fibonacci.html)
