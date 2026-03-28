package graph

import "testing"

func TestNewGraph(t *testing.T) {
	directed := NewGraph(true)
	if !directed.IsDirected() {
		t.Error("expected directed graph")
	}

	undirected := NewGraph(false)
	if undirected.IsDirected() {
		t.Error("expected undirected graph")
	}
}

func TestAddVertex(t *testing.T) {
	g := NewGraph(true)

	if !g.AddVertex("A") {
		t.Error("expected AddVertex to return true for new vertex")
	}
	if g.AddVertex("A") {
		t.Error("expected AddVertex to return false for duplicate vertex")
	}
	if g.AddVertex("") {
		t.Error("expected AddVertex to return false for empty id")
	}

	if !g.HasVertex("A") {
		t.Error("expected graph to contain vertex A")
	}
	if g.HasVertex("B") {
		t.Error("expected graph to not contain vertex B")
	}
}

func TestDirectedGraphAddEdgeAndOverwriteWeight(t *testing.T) {
	g := NewGraph(true)

	if !g.AddEdge("A", "B", 10) {
		t.Fatal("expected AddEdge to return true")
	}
	if !g.HasVertex("A") || !g.HasVertex("B") {
		t.Fatal("expected AddEdge to auto-create missing vertices")
	}
	if !g.HasEdge("A", "B") {
		t.Fatal("expected edge A->B to exist")
	}
	if g.HasEdge("B", "A") {
		t.Fatal("expected edge B->A to not exist in directed graph")
	}

	edge, ok := g.GetEdge("A", "B")
	if !ok {
		t.Fatal("expected to get edge A->B")
	}
	if edge.Weight() != 10 {
		t.Fatalf("expected weight 10, got %d", edge.Weight())
	}

	if !g.AddEdge("A", "B", 20) {
		t.Fatal("expected AddEdge to overwrite existing edge weight")
	}

	edge, ok = g.GetEdge("A", "B")
	if !ok {
		t.Fatal("expected to get edge A->B after overwrite")
	}
	if edge.Weight() != 20 {
		t.Fatalf("expected overwritten weight 20, got %d", edge.Weight())
	}

	edges := g.GetEdges()
	if len(edges) != 1 {
		t.Fatalf("expected 1 edge, got %d", len(edges))
	}
}

func TestUndirectedGraphAddEdgeMirroredAndDeduplicatedGetEdges(t *testing.T) {
	g := NewGraph(false)

	if !g.AddEdge("A", "B", 7) {
		t.Fatal("expected AddEdge to return true")
	}

	if !g.HasEdge("A", "B") {
		t.Fatal("expected edge A->B to exist")
	}
	if !g.HasEdge("B", "A") {
		t.Fatal("expected mirrored edge B->A to exist in undirected graph")
	}

	edges := g.GetEdges()
	if len(edges) != 1 {
		t.Fatalf("expected deduplicated GetEdges to return 1, got %d", len(edges))
	}
}

func TestAddEdgeInvalidInputs(t *testing.T) {
	g := NewGraph(true)

	if g.AddEdge("", "B", 1) {
		t.Error("expected AddEdge to return false for empty from")
	}
	if g.AddEdge("A", "", 1) {
		t.Error("expected AddEdge to return false for empty to")
	}
	if g.AddEdge("A", "A", 1) {
		t.Error("expected AddEdge to return false for self-loop")
	}
}

func TestRemoveEdge(t *testing.T) {
	t.Run("directed", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", 1)

		if !g.RemoveEdge("A", "B") {
			t.Fatal("expected RemoveEdge to return true")
		}
		if g.HasEdge("A", "B") {
			t.Fatal("expected edge A->B to be removed")
		}
		if g.RemoveEdge("A", "B") {
			t.Fatal("expected RemoveEdge to return false for missing edge")
		}
	})

	t.Run("undirected", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge("A", "B", 1)

		if !g.RemoveEdge("A", "B") {
			t.Fatal("expected RemoveEdge to return true")
		}
		if g.HasEdge("A", "B") || g.HasEdge("B", "A") {
			t.Fatal("expected both mirrored edges to be removed")
		}
	})
}

func TestRemoveVertex(t *testing.T) {
	g := NewGraph(true)
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 1)
	g.AddEdge("C", "B", 1)

	if !g.RemoveVertex("B") {
		t.Fatal("expected RemoveVertex to return true")
	}
	if g.HasVertex("B") {
		t.Fatal("expected vertex B to be removed")
	}
	if g.HasEdge("A", "B") || g.HasEdge("B", "C") || g.HasEdge("C", "B") {
		t.Fatal("expected all incoming and outgoing edges of B to be removed")
	}
	if g.RemoveVertex("B") {
		t.Fatal("expected RemoveVertex to return false for missing vertex")
	}
}

func TestDegree(t *testing.T) {
	t.Run("directed_total_degree", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", 1)
		g.AddEdge("C", "B", 1)
		g.AddEdge("B", "D", 1)

		degree, ok := g.Degree("B")
		if !ok {
			t.Fatal("expected Degree to return true for existing vertex")
		}
		if degree != 3 {
			t.Fatalf("expected total degree 3, got %d", degree)
		}
	})

	t.Run("undirected_degree", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge("A", "B", 1)
		g.AddEdge("B", "C", 1)

		degree, ok := g.Degree("B")
		if !ok {
			t.Fatal("expected Degree to return true for existing vertex")
		}
		if degree != 2 {
			t.Fatalf("expected degree 2, got %d", degree)
		}
	})

	if degree, ok := NewGraph(true).Degree("Z"); ok || degree != 0 {
		t.Fatal("expected Degree to return (0, false) for missing vertex")
	}
}

func TestNeighbors(t *testing.T) {
	g := NewGraph(true)
	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 1)

	neighbors, ok := g.Neighbors("A")
	if !ok {
		t.Fatal("expected Neighbors to return true for existing vertex")
	}
	if len(neighbors) != 2 {
		t.Fatalf("expected 2 neighbors, got %d", len(neighbors))
	}

	neighborSet := make(map[string]bool)
	for _, v := range neighbors {
		neighborSet[v.ID()] = true
	}

	if !neighborSet["B"] || !neighborSet["C"] {
		t.Fatal("expected neighbors to contain B and C")
	}

	missingNeighbors, ok := g.Neighbors("Z")
	if ok {
		t.Fatal("expected Neighbors to return false for missing vertex")
	}
	if missingNeighbors != nil {
		t.Fatal("expected nil neighbors for missing vertex")
	}
}

func TestShortestPathDirected(t *testing.T) {
	g := NewGraph(true)
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "D", 2)
	g.AddEdge("A", "C", 2)
	g.AddEdge("C", "D", 5)

	path, cost, ok := g.ShortestPath("A", "D")
	if !ok {
		t.Fatal("expected ShortestPath to find a route")
	}
	if cost != 3 {
		t.Fatalf("expected cost 3, got %d", cost)
	}
	if len(path) != 3 || path[0] != "A" || path[1] != "B" || path[2] != "D" {
		t.Fatalf("unexpected path: %v", path)
	}
}

func TestShortestPathUndirected(t *testing.T) {
	g := NewGraph(false)
	g.AddEdge("A", "B", 4)
	g.AddEdge("B", "C", 1)
	g.AddEdge("A", "C", 10)

	path, cost, ok := g.ShortestPath("A", "C")
	if !ok {
		t.Fatal("expected ShortestPath to find a route")
	}
	if cost != 5 {
		t.Fatalf("expected cost 5, got %d", cost)
	}
	if len(path) != 3 || path[0] != "A" || path[1] != "B" || path[2] != "C" {
		t.Fatalf("unexpected path: %v", path)
	}
}

func TestShortestPathStartEqualsGoal(t *testing.T) {
	g := NewGraph(true)
	g.AddVertex("A")

	path, cost, ok := g.ShortestPath("A", "A")
	if !ok {
		t.Fatal("expected ShortestPath to succeed when start equals goal")
	}
	if cost != 0 {
		t.Fatalf("expected cost 0, got %d", cost)
	}
	if len(path) != 1 || path[0] != "A" {
		t.Fatalf("unexpected path: %v", path)
	}
}

func TestShortestPathFailureCases(t *testing.T) {
	t.Run("no_path", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", 1)
		g.AddVertex("D")

		path, cost, ok := g.ShortestPath("A", "D")
		if ok {
			t.Fatal("expected ShortestPath to fail when no route exists")
		}
		if len(path) != 0 || cost != 0 {
			t.Fatalf("expected empty path and zero cost, got path=%v cost=%d", path, cost)
		}
	})

	t.Run("missing_vertex", func(t *testing.T) {
		g := NewGraph(true)

		path, cost, ok := g.ShortestPath("X", "Y")
		if ok {
			t.Fatal("expected ShortestPath to fail for missing vertices")
		}
		if len(path) != 0 || cost != 0 {
			t.Fatalf("expected empty path and zero cost, got path=%v cost=%d", path, cost)
		}
	})

	t.Run("negative_weight", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", -1)

		path, cost, ok := g.ShortestPath("A", "B")
		if ok {
			t.Fatal("expected ShortestPath to fail when graph has negative weights")
		}
		if len(path) != 0 || cost != 0 {
			t.Fatalf("expected empty path and zero cost, got path=%v cost=%d", path, cost)
		}
	})
}

func TestAStarDirected(t *testing.T) {
	g := NewGraph(true)
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "D", 2)
	g.AddEdge("A", "C", 2)
	g.AddEdge("C", "D", 5)

	heuristic := func(current, goal *Vertex) int {
		return 0
	}

	path, cost, ok := g.AStar("A", "D", heuristic)
	if !ok {
		t.Fatal("expected AStar to find a route")
	}
	if cost != 3 {
		t.Fatalf("expected cost 3, got %d", cost)
	}
	if len(path) != 3 || path[0] != "A" || path[1] != "B" || path[2] != "D" {
		t.Fatalf("unexpected path: %v", path)
	}
}

func TestAStarUndirected(t *testing.T) {
	g := NewGraph(false)
	g.AddEdge("A", "B", 4)
	g.AddEdge("B", "C", 1)
	g.AddEdge("A", "C", 10)

	heuristic := func(current, goal *Vertex) int {
		return 0
	}

	path, cost, ok := g.AStar("A", "C", heuristic)
	if !ok {
		t.Fatal("expected AStar to find a route")
	}
	if cost != 5 {
		t.Fatalf("expected cost 5, got %d", cost)
	}
	if len(path) != 3 || path[0] != "A" || path[1] != "B" || path[2] != "C" {
		t.Fatalf("unexpected path: %v", path)
	}
}

func TestAStarStartEqualsGoal(t *testing.T) {
	g := NewGraph(true)
	g.AddVertex("A")

	heuristic := func(current, goal *Vertex) int {
		return 0
	}

	path, cost, ok := g.AStar("A", "A", heuristic)
	if !ok {
		t.Fatal("expected AStar to succeed when start equals goal")
	}
	if cost != 0 {
		t.Fatalf("expected cost 0, got %d", cost)
	}
	if len(path) != 1 || path[0] != "A" {
		t.Fatalf("unexpected path: %v", path)
	}
}

func TestAStarFailureCases(t *testing.T) {
	t.Run("no_path", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", 1)
		g.AddVertex("D")

		heuristic := func(current, goal *Vertex) int {
			return 0
		}

		path, cost, ok := g.AStar("A", "D", heuristic)
		if ok {
			t.Fatal("expected AStar to fail when no route exists")
		}
		if len(path) != 0 || cost != 0 {
			t.Fatalf("expected empty path and zero cost, got path=%v cost=%d", path, cost)
		}
	})

	t.Run("nil_heuristic", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", 1)

		path, cost, ok := g.AStar("A", "B", nil)
		if ok {
			t.Fatal("expected AStar to fail when heuristic is nil")
		}
		if len(path) != 0 || cost != 0 {
			t.Fatalf("expected empty path and zero cost, got path=%v cost=%d", path, cost)
		}
	})

	t.Run("negative_weight", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge("A", "B", -1)

		heuristic := func(current, goal *Vertex) int {
			return 0
		}

		path, cost, ok := g.AStar("A", "B", heuristic)
		if ok {
			t.Fatal("expected AStar to fail when graph has negative weights")
		}
		if len(path) != 0 || cost != 0 {
			t.Fatalf("expected empty path and zero cost, got path=%v cost=%d", path, cost)
		}
	})
}

func TestHeuristicHelpers(t *testing.T) {
	coords := map[string][2]int{
		"A": {0, 0},
		"B": {3, 4},
	}

	extract := func(v *Vertex) (x, y int, ok bool) {
		c, exists := coords[v.ID()]
		if !exists {
			return 0, 0, false
		}
		return c[0], c[1], true
	}

	g := NewGraph(true)
	g.AddVertex("A")
	g.AddVertex("B")
	a, _ := g.GetVertex("A")
	b, _ := g.GetVertex("B")

	manhattan := ManhattanHeuristic(extract)
	if manhattan == nil {
		t.Fatal("expected ManhattanHeuristic to return a function")
	}
	if got := manhattan(a, b); got != 7 {
		t.Fatalf("expected Manhattan distance 7, got %d", got)
	}

	euclidean := EuclideanHeuristic(extract)
	if euclidean == nil {
		t.Fatal("expected EuclideanHeuristic to return a function")
	}
	if got := euclidean(a, b); got != 5 {
		t.Fatalf("expected Euclidean distance 5, got %d", got)
	}

	if ManhattanHeuristic(nil) != nil {
		t.Fatal("expected ManhattanHeuristic(nil) to return nil")
	}
	if EuclideanHeuristic(nil) != nil {
		t.Fatal("expected EuclideanHeuristic(nil) to return nil")
	}
}

func TestHeuristicHelpersMissingCoordinates(t *testing.T) {
	g := NewGraph(true)
	g.AddVertex("A")
	g.AddVertex("B")
	a, _ := g.GetVertex("A")
	b, _ := g.GetVertex("B")

	extract := func(v *Vertex) (x, y int, ok bool) {
		return 0, 0, false
	}

	if got := ManhattanHeuristic(extract)(a, b); got != 0 {
		t.Fatalf("expected Manhattan fallback 0, got %d", got)
	}
	if got := EuclideanHeuristic(extract)(a, b); got != 0 {
		t.Fatalf("expected Euclidean fallback 0, got %d", got)
	}
}

func TestAStarWithManhattanHelper(t *testing.T) {
	g := NewGraph(true)
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "D", 2)
	g.AddEdge("A", "C", 2)
	g.AddEdge("C", "D", 5)

	coords := map[string][2]int{
		"A": {0, 0},
		"B": {1, 0},
		"C": {0, 1},
		"D": {2, 0},
	}

	heuristic := ManhattanHeuristic(func(v *Vertex) (x, y int, ok bool) {
		c, exists := coords[v.ID()]
		if !exists {
			return 0, 0, false
		}
		return c[0], c[1], true
	})

	path, cost, ok := g.AStar("A", "D", heuristic)
	if !ok {
		t.Fatal("expected AStar to find a route with helper heuristic")
	}
	if cost != 3 {
		t.Fatalf("expected cost 3, got %d", cost)
	}
	if len(path) != 3 || path[0] != "A" || path[1] != "B" || path[2] != "D" {
		t.Fatalf("unexpected path: %v", path)
	}
}
