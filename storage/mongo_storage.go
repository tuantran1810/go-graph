package storage

import (
	"fmt"

	gograph "github.com/tuantran1810/go-graph/graph"
)

type mongoVertice struct {
	id         string                 `bson:"id"`
	label      string                 `bson:"label"`
	data       map[string]interface{} `bson:"data"`
	successors map[string]float64     `bson:"successors"`
}

func NewMongoVertice(id, label string, data map[string]interface{}, successors map[string]float64) IStorageVertice {
	return &mongoVertice{
		id:         id,
		label:      label,
		data:       data,
		successors: successors,
	}
}

func (v *mongoVertice) ID() string {
	return v.id
}

func (v *mongoVertice) Data() map[string]interface{} {
	return v.data
}

func (v *mongoVertice) Label() string {
	return v.label
}

func (v *mongoVertice) UpdateData(data map[string]interface{}) {
	v.data = data
}

func (v *mongoVertice) UpdateLabel(label string) {
	v.label = label
}

func (v *mongoVertice) CreateSuccessorConnection(id string, w float64) {
	v.successors[id] = w
}

func (v *mongoVertice) AllSuccessors() map[string]float64 {
	return v.successors
}

func (v *mongoVertice) ToGraphVertice() gograph.IVertice {
	return gograph.NewDefaultVertice(v.id, v.data)
}

func (v *mongoVertice) ToGraphSuccessorEdges() []gograph.IEdge {
	edges := make([]gograph.IEdge, 0)
	for sv, w := range v.successors {
		eid := fmt.Sprintf("%s_%s", v.id, sv)
		e := gograph.NewDefaultEdge(eid, v.id, sv, w, nil)
		edges = append(edges, e)
	}
	return edges
}

type mongoGraphStorage struct {
}
