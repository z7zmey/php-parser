package scalar_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
