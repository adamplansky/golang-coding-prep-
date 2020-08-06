package exercise

import (
	"fmt"
	"testing"
)

func Test_getMaxWidth(t *testing.T) {
	nodes := []*Node{
		&Node{Val: 0},
		&Node{Val: 1},
		&Node{Val: 2},
		&Node{Val: 3},
		&Node{Val: 4},
		&Node{Val: 5},
		&Node{Val: 6},
		&Node{Val: 7},
		&Node{Val: 8},
		&Node{Val: 9},
		&Node{Val: 10},
		&Node{Val: 11},
	}

	root := nodes[0]
	root.Left = nodes[1]
	root.Right = nodes[2]

	nodes[1].Left = nodes[3]
	nodes[1].Right = nodes[6]

	nodes[3].Left = nodes[4]
	nodes[4].Right = nodes[5]

	nodes[2].Left = nodes[7]
	nodes[2].Right = nodes[8]

	nodes[8].Right = nodes[9]
	nodes[9].Left = nodes[10]
	nodes[10].Left = nodes[11]

	fmt.Println(root)

	tests := []struct {
		name string
		root *Node
		want int
	}{
		{name: "empty root", root: nil, want: 0},
		{name: "only root node", root: nodes[0], want: 4},
		{name: "two siblings", root: nodes[1], want: 2.

		.},
		{name: "node 10", root: nodes[10], want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxWidth(tt.root); got != tt.want {
				t.Errorf("getMaxWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}
