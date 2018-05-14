package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
)

func TestShortList(t *testing.T) {
	src := `<? [$a] = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.ShortList{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
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
}

func TestShortListArrayIndex(t *testing.T) {
	src := `<? [$a[]] = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.ShortList{
						Items: []node.Node{
							&expr.ArrayItem{
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
}

func TestShortListList(t *testing.T) {
	src := `<? [list($a)] = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.ShortList{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.List{
									Items: []node.Node{
										&expr.ArrayItem{
											Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
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
}
