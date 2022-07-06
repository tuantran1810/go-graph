package algo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	gograph "github.com/tuantran1810/go-graph/graph"
	"github.com/tuantran1810/go-graph/utils"
)

func TestBFS(t *testing.T) {
	graph := gograph.NewDefaultGraph(false)
	graph.AddEdge(gograph.NewDefaultEdge("0_1", "0", "1", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("0_4", "0", "4", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("0_7", "0", "7", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("0_9", "0", "9", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("4_2", "4", "2", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("7_5", "7", "5", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("7_8", "7", "8", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("2_3", "2", "3", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("5_6", "5", "6", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("3_6", "3", "6", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("8_9", "8", "9", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("4_4", "4", "4", 0, nil))

	tests := []struct {
		graph          gograph.IGraph
		startVerticeID string
		want           map[string]*int
		wantErr        bool
	}{
		{
			graph:          graph,
			startVerticeID: "0",
			want: map[string]*int{
				"0": utils.NewIntPtr(0),
				"1": utils.NewIntPtr(1),
				"4": utils.NewIntPtr(1),
				"7": utils.NewIntPtr(1),
				"9": utils.NewIntPtr(1),
				"2": utils.NewIntPtr(2),
				"5": utils.NewIntPtr(2),
				"8": utils.NewIntPtr(2),
				"3": utils.NewIntPtr(3),
				"6": utils.NewIntPtr(3),
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestBFS_%i", i), func(t *testing.T) {
			got, err := BFS(tt.graph, tt.startVerticeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				gotJson, _ := json.Marshal(got)
				wantJson, _ := json.Marshal(tt.want)
				t.Errorf("BFS() = %s, want %s", gotJson, wantJson)
			}
		})
	}
}

func TestBFSTraversal(t *testing.T) {
	graph := gograph.NewDefaultGraph(false)
	graph.AddEdge(gograph.NewDefaultEdge("0_1", "0", "1", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("0_4", "0", "4", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("0_7", "0", "7", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("0_9", "0", "9", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("4_2", "4", "2", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("7_5", "7", "5", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("7_8", "7", "8", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("2_3", "2", "3", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("5_6", "5", "6", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("3_6", "3", "6", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("8_9", "8", "9", 0, nil))
	graph.AddEdge(gograph.NewDefaultEdge("4_4", "4", "4", 0, nil))

	text := ""

	type args struct {
		graph          gograph.IGraph
		startVerticeID string
		f              BFSTraversalFunc
	}
	tests := []struct {
		graph          gograph.IGraph
		startVerticeID string
		f              BFSTraversalFunc
		wantErr        bool
	}{
		{
			graph:          graph,
			startVerticeID: "0",
			f: func(vID string) {
				text = fmt.Sprintf("%s%s", text, vID)
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestBFSTraversal_%d", i), func(t *testing.T) {
			if err := BFSTraversal(tt.graph, tt.startVerticeID, tt.f); (err != nil) != tt.wantErr {
				t.Errorf("BFSTraversal() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("result text = %s", text)
			}
		})
	}
}
