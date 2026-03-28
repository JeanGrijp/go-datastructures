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
