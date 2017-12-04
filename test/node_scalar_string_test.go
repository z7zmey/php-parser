package test

import (
	"bytes"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/token"
)

func TestDoubleQuotedScalarString(t *testing.T) {
	src := `<? "test";`

	strToken := token.NewToken([]byte("\"test\""), 1, 1)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestMultilineDoubleQuotedScalarString(t *testing.T) {
	src := `<? "
	test
	";`

	strToken := token.NewToken([]byte("\"\n\ttest\n\t\""), 1, 3)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestSingleQuotedScalarString(t *testing.T) {
	src := `<? '$test';`

	strToken := token.NewToken([]byte("'$test'"), 1, 1)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestMultilineSingleQuotedScalarString(t *testing.T) {
	src := `<? '
	$test
	';`

	strToken := token.NewToken([]byte("'\n\t$test\n\t'"), 1, 3)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestPlainHeredocScalarString(t *testing.T) {
	src := `<? <<<CAD
	hello
CAD;
`

	strToken := token.NewToken([]byte("\thello\n"), 2, 3)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestHeredocScalarString(t *testing.T) {
	src := `<? <<<"CAD"
	hello
CAD;
`

	strToken := token.NewToken([]byte("\thello\n"), 2, 3)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestNowdocScalarString(t *testing.T) {
	src := `<? <<<'CAD'
	hello $world
CAD;
`

	strToken := token.NewToken([]byte("\thello $world\n"), 2, 3)
	strNode := node.NewNodeScalarString(strToken)
	expected := node.SimpleNode("Statements").Append(strNode)

	node := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, node); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
