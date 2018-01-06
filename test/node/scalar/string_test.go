package test

import (
	"bytes"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/parser"
)

func TestDoubleQuotedScalarString(t *testing.T) {
	src := `<? "test";`

	str := scalar.NewString("\"test\"").SetPosition(&node.Position{1, 1, 4, 9})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 1, 4, 10})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 10})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
func TestDoubleQuotedScalarStringWithEscapedVar(t *testing.T) {
	src := `<? "\$test";`

	str := scalar.NewString("\"\\$test\"").SetPosition(&node.Position{1, 1, 4, 11})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 1, 4, 12})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 12})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestMultilineDoubleQuotedScalarString(t *testing.T) {
	src := `<? "
	test
	";`

	str := scalar.NewString("\"\n\ttest\n\t\"").SetPosition(&node.Position{1, 3, 4, 13})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 3, 4, 14})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 3, 4, 14})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestSingleQuotedScalarString(t *testing.T) {
	src := `<? '$test';`

	str := scalar.NewString("'$test'").SetPosition(&node.Position{1, 1, 4, 10})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 1, 4, 11})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 11})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestMultilineSingleQuotedScalarString(t *testing.T) {
	src := `<? '
	$test
	';`

	str := scalar.NewString("'\n\t$test\n\t'").SetPosition(&node.Position{1, 3, 4, 14})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 3, 4, 15})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 3, 4, 15})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestPlainHeredocScalarString(t *testing.T) {
	src := `<? <<<CAD
	hello
CAD;
`

	str := scalar.NewString("\thello\n").SetPosition(&node.Position{1, 3, 4, 20})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 3, 4, 21})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 3, 4, 21})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestHeredocScalarString(t *testing.T) {
	src := `<? <<<"CAD"
	hello
CAD;
`

	str := scalar.NewString("\thello\n").SetPosition(&node.Position{1, 3, 4, 22})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 3, 4, 23})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 3, 4, 23})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestNowdocScalarString(t *testing.T) {
	src := `<? <<<'CAD'
	hello $world
CAD;
`

	str := scalar.NewString("\thello $world\n").SetPosition(&node.Position{1, 3, 4, 29})
	expr := stmt.NewExpression(str).SetPosition(&node.Position{1, 3, 4, 30})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 3, 4, 30})

	actual := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
