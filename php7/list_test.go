package php7

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/syntaxtree/linkedtree"
)

func TestList(t *testing.T) {
	s := stackedNodeList{}

	s.push()

	s.add(linkedtree.NodeID(1))
	s.add(linkedtree.NodeID(2))
	s.add(linkedtree.NodeID(3))

	expected := []linkedtree.NodeID{linkedtree.NodeID(1), linkedtree.NodeID(2), linkedtree.NodeID(3)}
	actual := s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linkedtree.NodeID{}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)
}

func TestListNested(t *testing.T) {
	s := stackedNodeList{}

	s.push()

	s.add(linkedtree.NodeID(1))
	s.add(linkedtree.NodeID(2))
	s.add(linkedtree.NodeID(3))

	s.push()
	s.add(linkedtree.NodeID(4))
	s.add(linkedtree.NodeID(5))

	expected := []linkedtree.NodeID{linkedtree.NodeID(4), linkedtree.NodeID(5)}
	actual := s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linkedtree.NodeID{linkedtree.NodeID(1), linkedtree.NodeID(2), linkedtree.NodeID(3)}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linkedtree.NodeID{}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)
}
