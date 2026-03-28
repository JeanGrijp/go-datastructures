# Graph Package

The graph package provides a weighted graph data structure with support for both directed and undirected graphs, plus shortest-path algorithms.

## Features

- Directed and undirected graph modes
- Weighted edges with integer weights
- Vertex and edge CRUD operations
- Utility queries:
  - IsDirected
  - HasVertex
  - HasEdge
  - Degree
  - Neighbors
- Shortest path algorithms:
  - Dijkstra via ShortestPath
  - A* via AStar

## API Overview

### Constructors

- NewGraph(directed bool) *Graph

### Graph Operations

- AddVertex(id string) bool
- AddEdge(from, to string, weight int) bool
- RemoveVertex(id string) bool
- RemoveEdge(from, to string) bool

### Accessors

- GetVertex(id string) (*Vertex, bool)
- GetEdge(from, to string) (*Edge, bool)
- GetVertices() []*Vertex
- GetEdges() []*Edge

### Utilities

- IsDirected() bool
- HasVertex(id string) bool
- HasEdge(from, to string) bool
- Degree(id string) (int, bool)
- Neighbors(id string) ([]*Vertex, bool)

### Shortest Path Algorithms

- ShortestPath(start, goal string) ([]string, int, bool)
- AStar(start, goal string, heuristic func(current, goal *Vertex) int) ([]string, int, bool)

### Heuristic Helpers

- CoordinateExtractor
- ManhattanHeuristic(extract CoordinateExtractor) func(current, goal *Vertex) int
- EuclideanHeuristic(extract CoordinateExtractor) func(current, goal *Vertex) int

## Behavior Notes

### Graph Mode

- If directed is true, edges are one-way.
- If directed is false, AddEdge creates both directions and RemoveEdge removes both directions.

### AddEdge Rules

- Empty vertex IDs are rejected.
- Self-loops are rejected.
- Missing vertices are created automatically.
- If an edge already exists, its weight is overwritten.

### Degree Semantics

- Undirected graph: number of adjacent vertices.
- Directed graph: total degree = in-degree + out-degree.

### GetEdges in Undirected Graphs

- GetEdges returns each undirected pair only once (A-B and B-A are deduplicated).

## ShortestPath (Dijkstra)

ShortestPath uses Dijkstra and returns:

- path: ordered list of vertex IDs from start to goal
- cost: total path cost
- ok: true on success

Failure returns: []string{}, 0, false

Failure cases:

- Missing start or goal
- Start/goal not found
- No path between start and goal
- Any negative edge weight in the graph

### Complexity

With a binary heap priority queue, runtime is typically O((V + E) log V).

## AStar

AStar has the same return contract and failure contract as ShortestPath, with one extra failure case:

- heuristic is nil

If your heuristic always returns 0, AStar behaves like Dijkstra.

### Heuristic Guidance

To preserve optimality, use an admissible heuristic: it should never overestimate the remaining cost to goal.

Helper behavior notes:

- If the extractor is nil, helper constructors return nil.
- If coordinates are missing for either vertex, returned heuristics produce 0.
- Returning 0 is a safe fallback and keeps A* equivalent to Dijkstra for that comparison.

## Examples

### Directed Graph

```go
package main

import (
    "fmt"

    "github.com/JeanGrijp/go-datastructures/pkg/graph"
)

func main() {
    g := graph.NewGraph(true)

    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "C", 2)

    fmt.Println(g.HasEdge("A", "B")) // true
    fmt.Println(g.HasEdge("B", "A")) // false
}
```

### Undirected Graph

```go
package main

import (
    "fmt"

    "github.com/JeanGrijp/go-datastructures/pkg/graph"
)

func main() {
    g := graph.NewGraph(false)

    g.AddEdge("A", "B", 3)

    fmt.Println(g.HasEdge("A", "B")) // true
    fmt.Println(g.HasEdge("B", "A")) // true
}
```

### Dijkstra with ShortestPath

```go
package main

import (
    "fmt"

    "github.com/JeanGrijp/go-datastructures/pkg/graph"
)

func main() {
    g := graph.NewGraph(true)
    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "D", 2)
    g.AddEdge("A", "C", 2)
    g.AddEdge("C", "D", 5)

    path, cost, ok := g.ShortestPath("A", "D")
    if !ok {
        fmt.Println("path not found")
        return
    }

    fmt.Println(path) // [A B D]
    fmt.Println(cost) // 3
}
```

### A* with a Heuristic

```go
package main

import (
    "fmt"

    "github.com/JeanGrijp/go-datastructures/pkg/graph"
)

func main() {
    g := graph.NewGraph(true)
    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "D", 2)
    g.AddEdge("A", "C", 2)
    g.AddEdge("C", "D", 5)

    heuristic := func(current, goal *graph.Vertex) int {
        // Replace this with a domain heuristic (e.g., Manhattan distance on grids).
        return 0
    }

    path, cost, ok := g.AStar("A", "D", heuristic)
    if !ok {
        fmt.Println("path not found")
        return
    }

    fmt.Println(path) // [A B D]
    fmt.Println(cost) // 3
}
```

### A* with Built-in Manhattan Helper

```go
package main

import (
    "fmt"

    "github.com/JeanGrijp/go-datastructures/pkg/graph"
)

type Point struct {
    X int
    Y int
}

func main() {
    g := graph.NewGraph(true)
    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "D", 2)
    g.AddEdge("A", "C", 2)
    g.AddEdge("C", "D", 5)

    coords := map[string]Point{
        "A": {X: 0, Y: 0},
        "B": {X: 1, Y: 0},
        "C": {X: 0, Y: 1},
        "D": {X: 2, Y: 0},
    }

    heuristic := graph.ManhattanHeuristic(func(v *graph.Vertex) (x, y int, ok bool) {
        p, exists := coords[v.ID()]
        if !exists {
            return 0, 0, false
        }
        return p.X, p.Y, true
    })

    path, cost, ok := g.AStar("A", "D", heuristic)
    if !ok {
        fmt.Println("path not found")
        return
    }

    fmt.Println(path) // [A B D]
    fmt.Println(cost) // 3
}
```

## Testing

Run tests for this package:

```bash
go test ./pkg/graph
```

Run all repository tests:

```bash
go test ./...
```
