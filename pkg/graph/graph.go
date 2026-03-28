// Package graph provides data structures and algorithms for working with graphs.
package graph

// Graph represents a graph with a set of vertices.
type Graph struct {
	vertices map[string]*Vertex
	directed bool
}

// Vertex represents a graph vertex.
type Vertex struct {
	id    string
	edges map[string]*Edge // outbound edges keyed by destination vertex ID
}

// Edge represents a weighted connection between two vertices.
type Edge struct {
	from   *Vertex
	to     *Vertex
	weight int
}

// NewGraph creates a new graph.
// When directed is true, the graph is directed.
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
		directed: directed,
	}
}

// IsDirected reports whether the graph is directed.
func (g *Graph) IsDirected() bool {
	return g.directed
}

// AddVertex adds a new vertex to the graph. It returns false if the vertex ID is empty or already exists.
func (g *Graph) AddVertex(id string) bool {
	if id == "" {
		return false
	}

	if _, ok := g.vertices[id]; ok {
		return false
	}

	g.vertices[id] = &Vertex{
		id:    id,
		edges: make(map[string]*Edge),
	}

	return true
}

// AddEdge adds or updates an edge in the graph.
// In an undirected graph, the reverse edge is also created/updated.
func (g *Graph) AddEdge(from, to string, weight int) bool {
	if from == "" || to == "" || from == to {
		return false
	}

	if _, ok := g.vertices[from]; !ok {
		g.AddVertex(from)
	}

	if _, ok := g.vertices[to]; !ok {
		g.AddVertex(to)
	}

	edge := &Edge{
		from:   g.vertices[from],
		to:     g.vertices[to],
		weight: weight,
	}
	g.vertices[from].edges[to] = edge

	if !g.directed {
		reverse := &Edge{
			from:   g.vertices[to],
			to:     g.vertices[from],
			weight: weight,
		}
		g.vertices[to].edges[from] = reverse
	}

	return true
}

// GetVertex returns a vertex by its ID.
func (g *Graph) GetVertex(id string) (*Vertex, bool) {
	vertex, ok := g.vertices[id]
	return vertex, ok
}

// GetEdge returns an edge between two vertices.
func (g *Graph) GetEdge(from, to string) (*Edge, bool) {
	vertex, ok := g.GetVertex(from)
	if !ok {
		return nil, false
	}
	edge, ok := vertex.edges[to]
	return edge, ok
}

// GetVertices returns all graph vertices.
func (g *Graph) GetVertices() []*Vertex {
	vertices := make([]*Vertex, 0, len(g.vertices))
	for _, vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// GetEdges returns all graph edges.
// In an undirected graph, each edge pair is returned once.
func (g *Graph) GetEdges() []*Edge {
	edges := make([]*Edge, 0)
	seenUndirected := make(map[string]struct{})

	for _, vertex := range g.vertices {
		for _, edge := range vertex.edges {
			if !g.directed {
				key := makeUndirectedEdgeKey(edge.from.id, edge.to.id)
				if _, ok := seenUndirected[key]; ok {
					continue
				}
				seenUndirected[key] = struct{}{}
			}

			edges = append(edges, edge)
		}
	}
	return edges
}

// RemoveVertex removes a vertex and all incident edges.
func (g *Graph) RemoveVertex(id string) bool {
	if _, ok := g.vertices[id]; !ok {
		return false
	}

	for _, vertex := range g.vertices {
		delete(vertex.edges, id)
	}

	delete(g.vertices, id)
	return true
}

// RemoveEdge removes an edge from the graph.
// In an undirected graph, the reverse edge is also removed.
func (g *Graph) RemoveEdge(from, to string) bool {
	vertex, ok := g.GetVertex(from)
	if !ok {
		return false
	}

	if _, ok := vertex.edges[to]; !ok {
		return false
	}

	delete(vertex.edges, to)

	if !g.directed {
		if reverseVertex, ok := g.GetVertex(to); ok {
			delete(reverseVertex.edges, from)
		}
	}

	return true
}

// HasVertex reports whether a vertex exists.
func (g *Graph) HasVertex(id string) bool {
	_, ok := g.vertices[id]
	return ok
}

// HasEdge reports whether an edge exists.
func (g *Graph) HasEdge(from, to string) bool {
	_, ok := g.GetEdge(from, to)
	return ok
}

// Degree returns the vertex degree.
// In directed graphs, this is the total degree (in-degree + out-degree).
func (g *Graph) Degree(id string) (int, bool) {
	vertex, ok := g.GetVertex(id)
	if !ok {
		return 0, false
	}

	if !g.directed {
		return len(vertex.edges), true
	}

	outDegree := len(vertex.edges)
	inDegree := 0
	for otherID, other := range g.vertices {
		if otherID == id {
			continue
		}
		if _, ok := other.edges[id]; ok {
			inDegree++
		}
	}

	return inDegree + outDegree, true
}

// Neighbors returns outbound neighbors reachable from the given vertex.
func (g *Graph) Neighbors(id string) ([]*Vertex, bool) {
	vertex, ok := g.GetVertex(id)
	if !ok {
		return nil, false
	}

	neighbors := make([]*Vertex, 0, len(vertex.edges))
	for _, edge := range vertex.edges {
		neighbors = append(neighbors, edge.to)
	}

	return neighbors, true
}

// ID returns the vertex identifier.
func (v *Vertex) ID() string {
	return v.id
}

// From returns the source vertex of the edge.
func (e *Edge) From() *Vertex {
	return e.from
}

// To returns the destination vertex of the edge.
func (e *Edge) To() *Vertex {
	return e.to
}

// Weight returns the edge weight.
func (e *Edge) Weight() int {
	return e.weight
}

func makeUndirectedEdgeKey(a, b string) string {
	if a < b {
		return a + "|" + b
	}
	return b + "|" + a
}
