package graph

type IEdge interface {
	// ID() returns unique ID of the edge
	ID() string

	// Data() returns free data of the edge
	Data() map[string]interface{}

	// Weight() returns the weight of the edge
	Weight() float64

	// From() returns the source vertice id
	From() string

	// To() returns the sink vertice id
	To() string
}

type defaultEdge struct {
	id     string
	data   map[string]interface{}
	weight float64
	from   string
	to     string
}

func NewDefaultEdge(id, from, to string, weight float64, data map[string]interface{}) IEdge {
	return &defaultEdge{
		id:     id,
		data:   data,
		weight: weight,
		from:   from,
		to:     to,
	}
}

func (e *defaultEdge) ID() string {
	return e.id
}

func (e *defaultEdge) Data() map[string]interface{} {
	return e.data
}

func (e *defaultEdge) Weight() float64 {
	return e.weight
}

func (e *defaultEdge) From() string {
	return e.from
}

func (e *defaultEdge) To() string {
	return e.to
}

type IVertice interface {
	// ID() returns unique ID of the vertice
	ID() string

	// Data() returns free data of the vertice
	Data() map[string]interface{}
}

type defaultVertice struct {
	id   string
	data map[string]interface{}
}

func NewDefaultVertice(id string, data map[string]interface{}) IVertice {
	return &defaultVertice{
		id:   id,
		data: data,
	}
}

func (v *defaultVertice) ID() string {
	return v.id
}

func (v *defaultVertice) Data() map[string]interface{} {
	return v.data
}

type SuccessorFunc func(verticeID string, edgeID string) bool

type IGraph interface {
	AddVertice(v IVertice)
	AddEdge(e IEdge)
	AllVertices() map[string]IVertice
	AllEdges() map[string]IEdge
	GetVertice(string) IVertice
	GetEdge(string) IEdge
	VerticeSuccessors(string) map[string]string
	VerticeSuccessorsTraversal(string, SuccessorFunc)
}

type defaultGraph struct {
	diGraph   bool
	vertices  map[string]IVertice
	edges     map[string]IEdge
	adjacency map[string]map[string]string
}

func NewDefaultGraph(diGraph bool) IGraph {
	return &defaultGraph{
		diGraph:   diGraph,
		vertices:  map[string]IVertice{},
		edges:     map[string]IEdge{},
		adjacency: map[string]map[string]string{},
	}
}

func (g *defaultGraph) AddVertice(v IVertice) {
	g.vertices[v.ID()] = v
}

func (g *defaultGraph) AddEdge(e IEdge) {
	vFromID := e.From()
	vToID := e.To()
	if _, found := g.vertices[vFromID]; !found {
		g.vertices[vFromID] = NewDefaultVertice(vFromID, nil)
	}
	if _, found := g.vertices[vToID]; !found {
		g.vertices[vToID] = NewDefaultVertice(vToID, nil)
	}
	g.edges[e.ID()] = e

	if _, found := g.adjacency[vFromID]; !found {
		g.adjacency[vFromID] = make(map[string]string)
	}
	g.adjacency[vFromID][vToID] = e.ID()

	if g.diGraph {
		return
	}

	if _, found := g.adjacency[vToID]; !found {
		g.adjacency[vToID] = make(map[string]string)
	}
	g.adjacency[vToID][vFromID] = e.ID()
}

func (g *defaultGraph) AllVertices() map[string]IVertice {
	return g.vertices
}

func (g *defaultGraph) AllEdges() map[string]IEdge {
	return g.edges
}

func (g *defaultGraph) GetVertice(id string) IVertice {
	if v, found := g.vertices[id]; found {
		return v
	}
	return nil
}

func (g *defaultGraph) GetEdge(id string) IEdge {
	if e, found := g.edges[id]; found {
		return e
	}
	return nil
}

func (g *defaultGraph) VerticeSuccessors(vID string) map[string]string {
	if successorMap, found := g.adjacency[vID]; found {
		return successorMap
	}
	return nil
}

func (g *defaultGraph) VerticeSuccessorsTraversal(vID string, f SuccessorFunc) {
	if successorMap, found := g.adjacency[vID]; found {
		for v, e := range successorMap {
			if !f(v, e) {
				break
			}
		}
	}
}
