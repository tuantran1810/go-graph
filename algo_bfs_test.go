package graph

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	graph := NewDefaultGraph(false)
	graph.AddEdge(NewDefaultEdge("0_1", "0", "1", 0, nil))
	graph.AddEdge(NewDefaultEdge("0_4", "0", "4", 0, nil))
	graph.AddEdge(NewDefaultEdge("0_7", "0", "7", 0, nil))
	graph.AddEdge(NewDefaultEdge("0_9", "0", "9", 0, nil))
	graph.AddEdge(NewDefaultEdge("4_2", "4", "2", 0, nil))
	graph.AddEdge(NewDefaultEdge("7_5", "7", "5", 0, nil))
	graph.AddEdge(NewDefaultEdge("7_8", "7", "8", 0, nil))
	graph.AddEdge(NewDefaultEdge("2_3", "2", "3", 0, nil))
	graph.AddEdge(NewDefaultEdge("5_6", "5", "6", 0, nil))
	graph.AddEdge(NewDefaultEdge("3_6", "3", "6", 0, nil))
	graph.AddEdge(NewDefaultEdge("8_9", "8", "9", 0, nil))
	graph.AddEdge(NewDefaultEdge("4_4", "4", "4", 0, nil))

	tests := []struct {
		graph          IGraph
		startVerticeID string
		want           map[string]*int
		wantErr        bool
	}{
		{
			graph:          graph,
			startVerticeID: "0",
			want: map[string]*int{
				"0": newIntPtr(0),
				"1": newIntPtr(1),
				"4": newIntPtr(1),
				"7": newIntPtr(1),
				"9": newIntPtr(1),
				"2": newIntPtr(2),
				"5": newIntPtr(2),
				"8": newIntPtr(2),
				"3": newIntPtr(3),
				"6": newIntPtr(3),
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
	graph := NewDefaultGraph(false)
	graph.AddEdge(NewDefaultEdge("0_1", "0", "1", 0, nil))
	graph.AddEdge(NewDefaultEdge("0_4", "0", "4", 0, nil))
	graph.AddEdge(NewDefaultEdge("0_7", "0", "7", 0, nil))
	graph.AddEdge(NewDefaultEdge("0_9", "0", "9", 0, nil))
	graph.AddEdge(NewDefaultEdge("4_2", "4", "2", 0, nil))
	graph.AddEdge(NewDefaultEdge("7_5", "7", "5", 0, nil))
	graph.AddEdge(NewDefaultEdge("7_8", "7", "8", 0, nil))
	graph.AddEdge(NewDefaultEdge("2_3", "2", "3", 0, nil))
	graph.AddEdge(NewDefaultEdge("5_6", "5", "6", 0, nil))
	graph.AddEdge(NewDefaultEdge("3_6", "3", "6", 0, nil))
	graph.AddEdge(NewDefaultEdge("8_9", "8", "9", 0, nil))
	graph.AddEdge(NewDefaultEdge("4_4", "4", "4", 0, nil))

	text := ""

	type args struct {
		graph          IGraph
		startVerticeID string
		f              BFSTraversalFunc
	}
	tests := []struct {
		graph          IGraph
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
