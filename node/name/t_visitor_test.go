package name_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/walker"
)

var nameNodesTests = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&name.Relative{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&name.NamePart{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
}

type visitorMock struct {
	visitChildren bool
	visitedKeys   []string
}

func (v *visitorMock) EnterNode(n walker.Walkable) bool { return v.visitChildren }
func (v *visitorMock) LeaveNode(n walker.Walkable)      {}
func (v *visitorMock) EnterChildNode(key string, w walker.Walkable) {
	v.visitedKeys = append(v.visitedKeys, key)
}
func (v *visitorMock) LeaveChildNode(key string, w walker.Walkable) {}
func (v *visitorMock) EnterChildList(key string, w walker.Walkable) {
	v.visitedKeys = append(v.visitedKeys, key)
}
func (v *visitorMock) LeaveChildList(key string, w walker.Walkable) {}

func TestNameVisitorDisableChildren(t *testing.T) {
	for _, tt := range nameNodesTests {
		v := &visitorMock{false, []string{}}
		tt.node.Walk(v)

		expected := []string{}
		actual := v.visitedKeys

		assert.DeepEqual(t, expected, actual)
	}
}

func TestNameVisitor(t *testing.T) {
	for _, tt := range nameNodesTests {
		v := &visitorMock{true, []string{}}
		tt.node.Walk(v)

		expected := tt.expectedVisitedKeys
		actual := v.visitedKeys

		assert.DeepEqual(t, expected, actual)
	}
}

// test Attributes()

func TestNameAttributes(t *testing.T) {
	for _, tt := range nameNodesTests {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		assert.DeepEqual(t, expected, actual)
	}
}
