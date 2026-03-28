package graph

import "container/heap"

// ShortestPath returns the shortest path between start and goal using Dijkstra.
// It returns ([]string{}, 0, false) when no path exists or when the graph
// contains negative-weight edges.
func (g *Graph) ShortestPath(start, goal string) ([]string, int, bool) {
	if start == "" || goal == "" {
		return []string{}, 0, false
	}

	if g.hasNegativeWeightEdge() {
		return []string{}, 0, false
	}

	if !g.HasVertex(start) || !g.HasVertex(goal) {
		return []string{}, 0, false
	}

	if start == goal {
		return []string{start}, 0, true
	}

	const inf = int(^uint(0) >> 1)
	dist := make(map[string]int, len(g.vertices))
	prev := make(map[string]string, len(g.vertices))
	for id := range g.vertices {
		dist[id] = inf
	}
	dist[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{vertexID: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*pqItem)
		if current.priority > dist[current.vertexID] {
			continue
		}

		if current.vertexID == goal {
			break
		}

		vertex := g.vertices[current.vertexID]
		for neighborID, edge := range vertex.edges {
			tentative := dist[current.vertexID] + edge.weight
			if tentative < dist[neighborID] {
				dist[neighborID] = tentative
				prev[neighborID] = current.vertexID
				heap.Push(pq, &pqItem{vertexID: neighborID, priority: tentative})
			}
		}
	}

	if dist[goal] == inf {
		return []string{}, 0, false
	}

	return buildPath(prev, start, goal), dist[goal], true
}

// AStar returns the shortest path between start and goal using A*.
// It returns ([]string{}, 0, false) when heuristic is nil, when no path exists,
// or when the graph contains negative-weight edges.
func (g *Graph) AStar(start, goal string, heuristic func(current, goal *Vertex) int) ([]string, int, bool) {
	if heuristic == nil {
		return []string{}, 0, false
	}

	if start == "" || goal == "" {
		return []string{}, 0, false
	}

	if g.hasNegativeWeightEdge() {
		return []string{}, 0, false
	}

	startVertex, startOK := g.GetVertex(start)
	goalVertex, goalOK := g.GetVertex(goal)
	if !startOK || !goalOK {
		return []string{}, 0, false
	}

	if start == goal {
		return []string{start}, 0, true
	}

	const inf = int(^uint(0) >> 1)
	gScore := make(map[string]int, len(g.vertices))
	prev := make(map[string]string, len(g.vertices))
	for id := range g.vertices {
		gScore[id] = inf
	}
	gScore[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{vertexID: start, priority: heuristic(startVertex, goalVertex)})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*pqItem)
		currentVertex := g.vertices[current.vertexID]

		expectedPriority := gScore[current.vertexID] + heuristic(currentVertex, goalVertex)
		if current.priority > expectedPriority {
			continue
		}

		if current.vertexID == goal {
			break
		}

		for neighborID, edge := range currentVertex.edges {
			tentativeG := gScore[current.vertexID] + edge.weight
			if tentativeG < gScore[neighborID] {
				gScore[neighborID] = tentativeG
				prev[neighborID] = current.vertexID
				neighborVertex := g.vertices[neighborID]
				fScore := tentativeG + heuristic(neighborVertex, goalVertex)
				heap.Push(pq, &pqItem{vertexID: neighborID, priority: fScore})
			}
		}
	}

	if gScore[goal] == inf {
		return []string{}, 0, false
	}

	return buildPath(prev, start, goal), gScore[goal], true
}

func (g *Graph) hasNegativeWeightEdge() bool {
	for _, vertex := range g.vertices {
		for _, edge := range vertex.edges {
			if edge.weight < 0 {
				return true
			}
		}
	}
	return false
}

func buildPath(prev map[string]string, start, goal string) []string {
	path := []string{goal}
	for current := goal; current != start; {
		parent := prev[current]
		path = append(path, parent)
		current = parent
	}
	reverseStrings(path)
	return path
}

func reverseStrings(values []string) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}

type pqItem struct {
	vertexID string
	priority int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(*pqItem))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
