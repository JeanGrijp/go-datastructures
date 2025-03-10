package graph

// Graph representa um grafo com um conjunto de vértices
type Graph struct {
	vertices map[string]*Vertex
}

// Vertex representa um vértice do grafo
type Vertex struct {
	id    string
	edges map[string]*Edge // arestas conectando a outros vértices
}

// Edge representa uma aresta que conecta dois vértices
type Edge struct {
	from   *Vertex
	to     *Vertex
	weight int // pode ser int ou float64, dependendo do caso
}

// NewGraph cria um novo grafo
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
	}
}

// AddVertex adiciona um vértice ao grafo
func (g *Graph) AddVertex(id string) {
	if _, ok := g.vertices[id]; !ok {
		g.vertices[id] = &Vertex{
			id:    id,
			edges: make(map[string]*Edge),
		}
	}
}

// AddEdge adiciona uma aresta ao grafo
func (g *Graph) AddEdge(from, to string, weight int) {
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
}
