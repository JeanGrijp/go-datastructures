package graph

import "math"

// CoordinateExtractor returns integer coordinates for a vertex.
// ok must be true when coordinates are available for the given vertex.
// When ok is false, heuristic helpers return 0 to keep the estimate admissible.
type CoordinateExtractor func(v *Vertex) (x, y int, ok bool)

// ManhattanHeuristic builds a Manhattan-distance heuristic for A*.
// The returned function is compatible with Graph.AStar.
func ManhattanHeuristic(extract CoordinateExtractor) func(current, goal *Vertex) int {
	if extract == nil {
		return nil
	}

	return func(current, goal *Vertex) int {
		x1, y1, ok1 := extract(current)
		x2, y2, ok2 := extract(goal)
		if !ok1 || !ok2 {
			return 0
		}
		return absInt(x1-x2) + absInt(y1-y2)
	}
}

// EuclideanHeuristic builds a Euclidean-distance heuristic for A*.
// The returned distance is truncated to int.
func EuclideanHeuristic(extract CoordinateExtractor) func(current, goal *Vertex) int {
	if extract == nil {
		return nil
	}

	return func(current, goal *Vertex) int {
		x1, y1, ok1 := extract(current)
		x2, y2, ok2 := extract(goal)
		if !ok1 || !ok2 {
			return 0
		}

		dx := float64(x1 - x2)
		dy := float64(y1 - y2)
		distance := math.Sqrt(dx*dx + dy*dy)
		return int(distance)
	}
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func RomaniaBucharestHeuristic(current, goal *Vertex) int {
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

func RomaniaIasiHeuristic(current, goal *Vertex) int {
	if goal == nil || goal.ID() != "Iasi" {
		return 0
	}

	straightLineToIasi := map[string]int{
		"Arad":           418,
		"Bucharest":      226,
		"Craiova":        193,
		"Drobeta":        242,
		"Eforie":         161,
		"Fagaras":        291,
		"Giurgiu":        77,
		"Hirsova":        151,
		"Iasi":           0,
		"Lugoj":          244,
		"Mehadia":        241,
		"Neamt":          87,
		"Oradea":         380,
		"Pitesti":        98,
		"Rimnicu Vilcea": 193,
		"Sibiu":          253,
		"Timisoara":      329,
		"Urziceni":       80,
		"Vaslui":         199,
		"Zerind":         374,
	}

	distance, ok := straightLineToIasi[current.ID()]
	if !ok {
		return 0
	}

	return distance
}
