package graph_test

import (
	"testing"

	"github.com/z7zmey/php-parser/internal/graph"
	"gotest.tools/assert"
)

func TestGraphTraverseDFS(t *testing.T) {
	g := new(graph.Graph)

	n1 := g.NewNode(1, 1)
	n2 := g.NewNode(2, 1)
	n3 := g.NewNode(3, 1)
	n4 := g.NewNode(4, 1)
	n5 := g.NewNode(5, 1)
	n6 := g.NewNode(6, 1)

	g.Link(n1, n2, nil)
	g.Link(n1, n3, nil)
	g.Link(n1, n4, nil)
	g.Link(n3, n5, nil)
	g.Link(n3, n6, nil)
	g.Link(n3, n1, nil)

	expected := []int{
		0, 1,
		1, 2,
		1, 3,
		2, 5,
		2, 6,
		1, 4,
	}
	actual := []int{}

	g.TraverseDFS(n1, func(node graph.Node, depth int) bool {
		actual = append(actual, depth)
		actual = append(actual, int(node.ID))

		return true
	})

	assert.DeepEqual(t, expected, actual)
}
