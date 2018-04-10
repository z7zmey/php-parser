package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestEmptyList(t *testing.T) {
	src := `<? list() = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestList(t *testing.T) {
	src := `<? list($a) = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestListArrayIndex(t *testing.T) {
	src := `<? list($a[]) = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val: &expr.ArrayDimFetch{
									Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestListList(t *testing.T) {
	src := `<? list(list($a)) = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val: &expr.List{
									Items: []node.Node{
										&expr.ArrayItem{
											ByRef: false,
											Val:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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
