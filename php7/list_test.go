package php7

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/ast/linear"
)

func TestList(t *testing.T) {
	s := stackedNodeList{}

	s.push()

	s.add(linear.NodeID(1))
	s.add(linear.NodeID(2))
	s.add(linear.NodeID(3))

	expected := []linear.NodeID{linear.NodeID(1), linear.NodeID(2), linear.NodeID(3)}
	actual := s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linear.NodeID{}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)
}

func TestListNested(t *testing.T) {
	s := stackedNodeList{}

	s.push()

	s.add(linear.NodeID(1))
	s.add(linear.NodeID(2))
	s.add(linear.NodeID(3))

	s.push()
	s.add(linear.NodeID(4))
	s.add(linear.NodeID(5))

	expected := []linear.NodeID{linear.NodeID(4), linear.NodeID(5)}
	actual := s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linear.NodeID{linear.NodeID(1), linear.NodeID(2), linear.NodeID(3)}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linear.NodeID{}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)
}
