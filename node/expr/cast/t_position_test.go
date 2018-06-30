package cast_test

import (
	"testing"

	"github.com/z7zmey/php-parser/position"
)

func TestPosition(t *testing.T) {
	expected := position.NewPosition(1, 1, 1, 1)
	for _, n := range nodes {
		n.SetPosition(expected)
		actual := n.GetPosition()
		assertEqual(t, expected, actual)
	}
}
