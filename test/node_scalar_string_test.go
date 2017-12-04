package test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/token"
)

func TestNewNodeScalarString(t *testing.T) {
	src := `<? "test";`

	strToken := token.NewToken([]byte("\"test\""), 1, 1)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if !reflect.DeepEqual(expected, node) {
		t.Error("Not equal")
	}
}
