package scalar_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

var nameNodesTests = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&scalar.String{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
	{
		&scalar.Lnumber{Value: "0"},
		[]string{},
		map[string]interface{}{"Value": "0"},
	},
	{
		&scalar.Dnumber{Value: "0.1"},
		[]string{},
		map[string]interface{}{"Value": "0.1"},
	},
	{
		&scalar.MagicConstant{Value: "__DIR__"},
		[]string{},
		map[string]interface{}{"Value": "__DIR__"},
	},
	{
		&scalar.EncapsedStringPart{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
	{
		&scalar.Encapsed{Parts: []node.Node{&scalar.EncapsedStringPart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&scalar.Heredoc{Label: "LBL", Parts: []node.Node{&scalar.EncapsedStringPart{Value: "foo"}}},
		[]string{"Parts"},
		map[string]interface{}{"Label": "LBL"},
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
