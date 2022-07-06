package main

import (
	"encoding/json"
	"fmt"

	"github.com/tuantran1810/go-graph/algo"
	"github.com/tuantran1810/go-graph/graph"
	"github.com/tuantran1810/go-graph/storage"
)

func main() {
	graphStorage := storage.NewOnmemGraphStorage()
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u1", "c1", nil, map[string]float64{"d1": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u2", "c1", nil, map[string]float64{"d1": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u3", "c1", nil, map[string]float64{"d1": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u4", "c2", nil, map[string]float64{"d2": 1.0, "d4": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u5", "c2", nil, map[string]float64{"d2": 1.0, "d3": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u6", "c2", nil, map[string]float64{"d3": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u7", "c2", nil, map[string]float64{"d3": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u8", "c2", nil, map[string]float64{"d4": 1.0}))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("u9", "c2", nil, map[string]float64{"d4": 1.0}))

	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("d1", "c1", nil, nil))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("d2", "c2", nil, nil))
	graphStorage.UpsertVertice(nil, storage.NewOnmemVertice("d3", "c2", nil, nil))

	// nếu có cạnh mới u3_d2
	u3, _ := graphStorage.GetVertice(nil, "u3")
	d2, _ := graphStorage.GetVertice(nil, "d2")
	u3.CreateSuccessorConnection("d2", 1.0)
	graphStorage.UpsertVertice(nil, u3)

	if u3.Label() == d2.Label() {
		// cùng label thì thôi
	}
	fmt.Printf("khác label nè: %s - %s\n", u3.Label(), d2.Label())
	c1, _ := graphStorage.GetVerticeswithLabel(nil, u3.Label())
	c2, _ := graphStorage.GetVerticeswithLabel(nil, d2.Label())
	allVertices := append(c1, c2...)
	fmt.Printf("join 2 clusters lại, có %d vertices\n", len(allVertices))
	for _, v := range allVertices {
		v.UpdateLabel("c3")
	}
	graphStorage.UpsertVertices(nil, allVertices)

	// giờ là query u1 nè
	u1, _ := graphStorage.GetVertice(nil, "u1")
	clusterVertices, _ := graphStorage.GetVerticeswithLabel(nil, u1.Label())
	graphContainer := graph.NewDefaultGraph(false)
	for _, v := range clusterVertices {
		for _, e := range v.ToGraphSuccessorEdges() {
			graphContainer.AddEdge(e)
		}
	}
	result, _ := algo.BFS(graphContainer, "u1")
	str, _ := json.Marshal(result)
	print("%s\n", string(str))
}
