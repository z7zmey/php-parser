package test

import (
	"bytes"
	"testing"

	"github.com/kylelemons/godebug/pretty"
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

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("post-AddCrew diff: (-expected +actual)\n%s", diff)
	}
}
