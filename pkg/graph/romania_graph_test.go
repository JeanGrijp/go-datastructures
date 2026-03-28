package graph

import "testing"

type romaniaRoad struct {
	from   string
	to     string
	weight int
}

func buildRomaniaGraph() *Graph {
	g := NewGraph(false)

	roads := []romaniaRoad{
		{from: "Arad", to: "Zerind", weight: 75},
		{from: "Arad", to: "Sibiu", weight: 140},
		{from: "Arad", to: "Timisoara", weight: 118},
		{from: "Zerind", to: "Oradea", weight: 71},
		{from: "Oradea", to: "Sibiu", weight: 151},
		{from: "Sibiu", to: "Fagaras", weight: 99},
		{from: "Sibiu", to: "Rimnicu Vilcea", weight: 80},
		{from: "Fagaras", to: "Bucharest", weight: 211},
		{from: "Rimnicu Vilcea", to: "Pitesti", weight: 97},
		{from: "Rimnicu Vilcea", to: "Craiova", weight: 146},
		{from: "Pitesti", to: "Bucharest", weight: 101},
		{from: "Pitesti", to: "Craiova", weight: 138},
		{from: "Timisoara", to: "Lugoj", weight: 111},
		{from: "Lugoj", to: "Mehadia", weight: 70},
		{from: "Mehadia", to: "Drobeta", weight: 75},
		{from: "Drobeta", to: "Craiova", weight: 120},
		{from: "Bucharest", to: "Giurgiu", weight: 90},
		{from: "Bucharest", to: "Urziceni", weight: 85},
		{from: "Urziceni", to: "Hirsova", weight: 98},
		{from: "Hirsova", to: "Eforie", weight: 86},
		{from: "Urziceni", to: "Vaslui", weight: 142},
		{from: "Vaslui", to: "Iasi", weight: 92},
		{from: "Iasi", to: "Neamt", weight: 87},
	}

	for _, road := range roads {
		g.AddEdge(road.from, road.to, road.weight)
	}

	return g
}

func romaniaBucharestHeuristic(current, goal *Vertex) int {
	if goal == nil || goal.ID() != "Bucharest" {
		return 0
	}

	straightLineToBucharest := map[string]int{
		"Arad":           366,
		"Bucharest":      0,
		"Craiova":        160,
		"Drobeta":        242,
		"Eforie":         161,
		"Fagaras":        178,
		"Giurgiu":        77,
		"Hirsova":        151,
		"Iasi":           226,
		"Lugoj":          244,
		"Mehadia":        241,
		"Neamt":          234,
		"Oradea":         380,
		"Pitesti":        98,
		"Rimnicu Vilcea": 193,
		"Sibiu":          253,
		"Timisoara":      329,
		"Urziceni":       80,
		"Vaslui":         199,
		"Zerind":         374,
	}

	distance, ok := straightLineToBucharest[current.ID()]
	if !ok {
		return 0
	}

	return distance
}

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
	g := buildRomaniaGraph()

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
	g := buildRomaniaGraph()

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
	g := buildRomaniaGraph()

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
			path, cost, ok := g.AStar(tc.start, tc.goal, romaniaBucharestHeuristic)
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
