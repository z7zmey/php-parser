package stmt_test

import (
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/name"
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestSimpleFunction(t *testing.T) {
	src := `<? function foo() {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef: false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestFunctionReturn(t *testing.T) {
	src := `<? function foo() {return;}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef: false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.Return{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestFunctionReturnVar(t *testing.T) {
	src := `<? function foo() {return $a;}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef: false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.Return{
						Expr: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestRefFunction(t *testing.T) {
	src := `<? function &foo() {return 1;}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef: true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.Return{
						Expr: &scalar.Lnumber{Value: "1"},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestReturnTypeFunction(t *testing.T) {
	src := `<? function &foo(): void {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef: true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{Value: "foo"},
				ReturnType: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "void"},
					},
				},
				Stmts: []node.Node{},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
