package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"

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
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts:         []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestFunctionReturn(t *testing.T) {
	src := `<? function foo() {return;}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.Return{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestFunctionReturnVar(t *testing.T) {
	src := `<? function foo(array $a, callable $b) {return $a;}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Params: []node.Node{
					&node.Parameter{
						ByRef:        false,
						Variadic:     false,
						VariableType: &node.Identifier{Value: "array"},
						Variable:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
					&node.Parameter{
						ByRef:        false,
						Variadic:     false,
						VariableType: &node.Identifier{Value: "callable"},
						Variable:     &expr.Variable{VarName: &node.Identifier{Value: "b"}},
					},
				},
				Stmts: []node.Node{
					&stmt.Return{
						Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestRefFunction(t *testing.T) {
	src := `<? function &foo() {return 1;}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.Return{
						Expr: &scalar.Lnumber{Value: "1"},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestReturnTypeFunction(t *testing.T) {
	src := `<? function &foo(): void {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				ReturnType: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "void"},
					},
				},
				Stmts: []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
