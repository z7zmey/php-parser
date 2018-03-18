package assign_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		diff := pretty.Compare(expected, actual)

		if diff != "" {
			t.Errorf("diff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("expected and actual are not equal\n")
		}
	}
}

func TestAssignRef(t *testing.T) {
	src := `<? $a =& $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.AssignRef{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestAssignRefNew(t *testing.T) {
	src := `<? $a =& new Foo;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.AssignRef{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.New{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
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

func TestAssignRefArgs(t *testing.T) {
	src := `<? $a =& new Foo($b);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.AssignRef{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.New{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Arguments: []node.Node{
							&node.Argument{
								Variadic:    false,
								IsReference: false,
								Expr:        &expr.Variable{VarName: &node.Identifier{Value: "b"}},
							},
						},
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

func TestAssign(t *testing.T) {
	src := `<? $a = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBitwiseAnd(t *testing.T) {
	src := `<? $a &= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.BitwiseAnd{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBitwiseOr(t *testing.T) {
	src := `<? $a |= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.BitwiseOr{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBitwiseXor(t *testing.T) {
	src := `<? $a ^= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.BitwiseXor{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestConcat(t *testing.T) {
	src := `<? $a .= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Concat{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestDiv(t *testing.T) {
	src := `<? $a /= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Div{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestMinus(t *testing.T) {
	src := `<? $a -= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Minus{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestMod(t *testing.T) {
	src := `<? $a %= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Mod{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestMul(t *testing.T) {
	src := `<? $a *= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Mul{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestPlus(t *testing.T) {
	src := `<? $a += $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Plus{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestPow(t *testing.T) {
	src := `<? $a **= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Pow{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestShiftLeft(t *testing.T) {
	src := `<? $a <<= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.ShiftLeft{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestShiftRight(t *testing.T) {
	src := `<? $a >>= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.ShiftRight{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
