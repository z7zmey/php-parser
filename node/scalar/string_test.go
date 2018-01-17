package scalar_test

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

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"test\""},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
func TestDoubleQuotedScalarStringWithEscapedVar(t *testing.T) {
	src := `<? "\$test";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"\\$test\""},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestMultilineDoubleQuotedScalarString(t *testing.T) {
	src := `<? "
	test
	";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"\n\ttest\n\t\""},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestSingleQuotedScalarString(t *testing.T) {
	src := `<? '$test';`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "'$test'"},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestMultilineSingleQuotedScalarString(t *testing.T) {
	src := `<? '
	$test
	';`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "'\n\t$test\n\t'"},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestPlainHeredocScalarString(t *testing.T) {
	src := `<? <<<CAD
	hello
CAD;
`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\thello\n"},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestHeredocScalarString(t *testing.T) {
	src := `<? <<<"CAD"
	hello
CAD;
`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\thello\n"},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestNowdocScalarString(t *testing.T) {
	src := `<? <<<'CAD'
	hello $world
CAD;
`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\thello $world\n"},
			},
		},
	}

	actual, _, _ := parser.ParsePhp7(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
