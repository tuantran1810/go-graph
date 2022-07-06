package storage

import (
	"context"

	gograph "github.com/tuantran1810/go-graph/graph"
)

type IStorageVertice interface {
	ID() string
	Data() map[string]interface{}
	Label() string
	UpdateData(map[string]interface{})
	UpdateLabel(string)
	CreateSuccessorConnection(string, float64)
	AllSuccessors() map[string]float64
	ToGraphVertice() gograph.IVertice
	ToGraphSuccessorEdges() []gograph.IEdge
}

type IGraphStorage interface {
	GetVertice(ctx context.Context, id string) (IStorageVertice, error)
	GetVertices(ctx context.Context, ids []string) ([]IStorageVertice, error)
	GetVerticeswithLabel(ctx context.Context, label string) ([]IStorageVertice, error)
	UpsertVertice(ctx context.Context, v IStorageVertice) error
	UpsertVertices(ctx context.Context, v []IStorageVertice) error
}
