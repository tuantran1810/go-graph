package storage

import (
	"context"
	"fmt"

	gograph "github.com/tuantran1810/go-graph/graph"
	"go.uber.org/multierr"
)

type onmemVertice struct {
	id         string
	label      string
	data       map[string]interface{}
	successors map[string]float64
}

func NewOnmemVertice(id, label string, data map[string]interface{}, successors map[string]float64) IStorageVertice {
	if successors == nil {
		successors = map[string]float64{}
	}
	return &onmemVertice{
		id:         id,
		label:      label,
		data:       data,
		successors: successors,
	}
}

func (v *onmemVertice) ID() string {
	return v.id
}

func (v *onmemVertice) Data() map[string]interface{} {
	return v.data
}

func (v *onmemVertice) Label() string {
	return v.label
}

func (v *onmemVertice) UpdateData(data map[string]interface{}) {
	v.data = data
}

func (v *onmemVertice) UpdateLabel(label string) {
	v.label = label
}

func (v *onmemVertice) CreateSuccessorConnection(id string, w float64) {
	v.successors[id] = w
}

func (v *onmemVertice) AllSuccessors() map[string]float64 {
	return v.successors
}

func (v *onmemVertice) ToGraphVertice() gograph.IVertice {
	return gograph.NewDefaultVertice(v.id, v.data)
}

func (v *onmemVertice) ToGraphSuccessorEdges() []gograph.IEdge {
	edges := make([]gograph.IEdge, 0)
	for sv, w := range v.successors {
		eid := fmt.Sprintf("%s_%s", v.id, sv)
		e := gograph.NewDefaultEdge(eid, v.id, sv, w, nil)
		edges = append(edges, e)
	}
	return edges
}

type onmemGraphStorage struct {
	vertices map[string]*onmemVertice
	labels   map[string]map[string]bool
}

func NewOnmemGraphStorage() IGraphStorage {
	return &onmemGraphStorage{
		vertices: make(map[string]*onmemVertice),
		labels:   make(map[string]map[string]bool),
	}
}

func (s *onmemGraphStorage) GetVertice(_ context.Context, id string) (IStorageVertice, error) {
	if v, found := s.vertices[id]; found {
		return v, nil
	}
	return nil, fmt.Errorf("node with id %s does not exist", id)
}

func (s *onmemGraphStorage) GetVertices(_ context.Context, ids []string) ([]IStorageVertice, error) {
	var err error
	vertices := make([]IStorageVertice, 0)
	for _, id := range ids {
		v, e := s.GetVertice(nil, id)
		err = multierr.Append(err, e)
		vertices = append(vertices, v)
	}
	return vertices, err
}

func (s *onmemGraphStorage) GetVerticeswithLabel(_ context.Context, label string) ([]IStorageVertice, error) {
	mp, found := s.labels[label]
	if !found {
		return nil, fmt.Errorf("label %s not found")
	}
	ids := make([]string, 0)
	for id, _ := range mp {
		ids = append(ids, id)
	}
	return s.GetVertices(nil, ids)
}

func (s *onmemGraphStorage) UpsertVertice(_ context.Context, v IStorageVertice) error {
	newVertice := NewOnmemVertice(v.ID(), v.Label(), v.Data(), v.AllSuccessors()).(*onmemVertice)
	s.vertices[v.ID()] = newVertice
	if mp, found := s.labels[newVertice.Label()]; found {
		mp[newVertice.ID()] = true
	} else {
		mp = make(map[string]bool)
		mp[newVertice.ID()] = true
		s.labels[newVertice.Label()] = mp
	}
	return nil
}

func (s *onmemGraphStorage) UpsertVertices(_ context.Context, v []IStorageVertice) error {
	var err error
	for _, iv := range v {
		err = multierr.Append(err, s.UpsertVertice(nil, iv))
	}
	return err
}
