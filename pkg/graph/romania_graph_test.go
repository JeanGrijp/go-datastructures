package graph

import "testing"

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRomaniaGraphRepresentation_TableDriven(t *testing.T) {
	g := BuildRomaniaGraph()

	if g.IsDirected() {
		t.Fatal("expected Romania graph to be undirected")
	}

	vertexCount := len(g.GetVertices())
	if vertexCount != 20 {
		t.Fatalf("expected 20 vertices, got %d", vertexCount)
	}

	edgeCount := len(g.GetEdges())
	if edgeCount != 23 {
		t.Fatalf("expected 23 undirected edges, got %d", edgeCount)
	}

	cases := []struct {
		name   string
		from   string
		to     string
		weight int
	}{
		{name: "Arad to Sibiu", from: "Arad", to: "Sibiu", weight: 140},
		{name: "Sibiu to Arad mirrored", from: "Sibiu", to: "Arad", weight: 140},
		{name: "Rimnicu Vilcea to Pitesti", from: "Rimnicu Vilcea", to: "Pitesti", weight: 97},
		{name: "Urziceni to Vaslui", from: "Urziceni", to: "Vaslui", weight: 142},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			edge, ok := g.GetEdge(tc.from, tc.to)
			if !ok {
				t.Fatalf("expected edge %s -> %s to exist", tc.from, tc.to)
			}
			if edge.Weight() != tc.weight {
				t.Fatalf("expected edge weight %d, got %d", tc.weight, edge.Weight())
			}
		})
	}
}

func TestRomaniaShortestPath_TableDriven(t *testing.T) {
	g := BuildRomaniaGraph()

	cases := []struct {
		name         string
		start        string
		goal         string
		expectedPath []string
		expectedCost int
	}{
		{
			name:         "Arad to Bucharest",
			start:        "Arad",
			goal:         "Bucharest",
			expectedPath: []string{"Arad", "Sibiu", "Rimnicu Vilcea", "Pitesti", "Bucharest"},
			expectedCost: 418,
		},
		{
			name:         "Timisoara to Bucharest",
			start:        "Timisoara",
			goal:         "Bucharest",
			expectedPath: []string{"Timisoara", "Arad", "Sibiu", "Rimnicu Vilcea", "Pitesti", "Bucharest"},
			expectedCost: 536,
		},
		{
			name:         "Neamt to Bucharest",
			start:        "Neamt",
			goal:         "Bucharest",
			expectedPath: []string{"Neamt", "Iasi", "Vaslui", "Urziceni", "Bucharest"},
			expectedCost: 406,
		},
		{
			name:         "Bucharest to Arad",
			start:        "Bucharest",
			goal:         "Arad",
			expectedPath: []string{"Bucharest", "Pitesti", "Rimnicu Vilcea", "Sibiu", "Arad"},
			expectedCost: 418,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path, cost, ok := g.ShortestPath(tc.start, tc.goal)
			if !ok {
				t.Fatalf("expected ShortestPath to find a route from %s to %s", tc.start, tc.goal)
			}
			if cost != tc.expectedCost {
				t.Fatalf("expected cost %d, got %d", tc.expectedCost, cost)
			}
			if !equalStringSlices(path, tc.expectedPath) {
				t.Fatalf("unexpected path. expected %v, got %v", tc.expectedPath, path)
			}
		})
	}
}

func TestRomaniaAStar_TableDriven(t *testing.T) {
	g := BuildRomaniaGraph()

	cases := []struct {
		name         string
		start        string
		goal         string
		expectedCost int
	}{
		{name: "Arad to Bucharest", start: "Arad", goal: "Bucharest", expectedCost: 418},
		{name: "Lugoj to Bucharest", start: "Lugoj", goal: "Bucharest", expectedCost: 504},
		{name: "Oradea to Bucharest", start: "Oradea", goal: "Bucharest", expectedCost: 429},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path, cost, ok := g.AStar(tc.start, tc.goal, RomaniaBucharestHeuristic)
			if !ok {
				t.Fatalf("expected AStar to find a route from %s to %s", tc.start, tc.goal)
			}
			if len(path) == 0 {
				t.Fatal("expected non-empty path")
			}
			if path[0] != tc.start || path[len(path)-1] != tc.goal {
				t.Fatalf("expected path to start at %s and end at %s, got %v", tc.start, tc.goal, path)
			}
			if cost != tc.expectedCost {
				t.Fatalf("expected cost %d, got %d", tc.expectedCost, cost)
			}

			// A* and Dijkstra should agree on optimal path cost for this map.
			_, dijkstraCost, dijkstraOK := g.ShortestPath(tc.start, tc.goal)
			if !dijkstraOK {
				t.Fatalf("expected ShortestPath to find a route from %s to %s", tc.start, tc.goal)
			}
			if cost != dijkstraCost {
				t.Fatalf("expected A* cost %d to match Dijkstra cost %d", cost, dijkstraCost)
			}
		})
	}
}
