package algo

import (
	"container/list"
	"fmt"

	gograph "github.com/tuantran1810/go-graph/graph"
	"github.com/tuantran1810/go-graph/utils"
)

// BFS() returns distance from start vertice to the others, nil means inf
func BFS(graph gograph.IGraph, startVerticeID string) (map[string]*int, error) {
	startVertice := graph.GetVertice(startVerticeID)
	if startVertice == nil {
		return nil, fmt.Errorf("start vertice not exist, id: %s", startVerticeID)
	}

	allVertices := graph.AllVertices()
	result := make(map[string]*int)
	for vID, _ := range allVertices {
		result[vID] = nil
	}

	distance := 0
	queue := list.New()
	queue.PushBack(startVerticeID)
	queue.PushBack("")
	visited := make(map[string]bool)

	for {
		if queue.Len() == 0 {
			break
		}

		vID := queue.Front().Value.(string)
		queue.Remove(queue.Front())
		if len(vID) == 0 {
			distance++
			if queue.Len() == 0 {
				break
			}
			lastID := queue.Back().Value.(string)
			if len(lastID) == 0 {
				continue
			}
			queue.PushBack("")
			continue
		}

		if found := visited[vID]; found {
			continue
		}
		visited[vID] = true
		result[vID] = utils.NewIntPtr(distance)
		graph.VerticeSuccessorsTraversal(
			vID,
			func(successorID, _ string) bool {
				if found := visited[successorID]; found {
					return true
				}
				queue.PushBack(successorID)
				return true
			},
		)
	}

	return result, nil
}

type BFSTraversalFunc func(vID string)

// BFSTraversal() travels from start vertice to the others
func BFSTraversal(graph gograph.IGraph, startVerticeID string, f BFSTraversalFunc) error {
	startVertice := graph.GetVertice(startVerticeID)
	if startVertice == nil {
		return fmt.Errorf("start vertice not exist, id: %s", startVerticeID)
	}

	queue := list.New()
	queue.PushBack(startVerticeID)
	visited := make(map[string]bool)

	for {
		if queue.Len() == 0 {
			break
		}
		vID := queue.Front().Value.(string)
		queue.Remove(queue.Front())
		if found := visited[vID]; found {
			continue
		}
		visited[vID] = true

		f(vID)
		graph.VerticeSuccessorsTraversal(
			vID,
			func(successorID, edgeID string) bool {
				if found := visited[successorID]; found {
					return true
				}
				queue.PushBack(successorID)
				return true
			},
		)
	}
	return nil
}
