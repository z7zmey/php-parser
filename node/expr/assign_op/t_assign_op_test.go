package assign_op_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign_op"

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
				Expr: &assign_op.AssignRef{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.AssignRef{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
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
				Expr: &assign_op.AssignRef{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.New{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Arguments: []node.Node{
							&node.Argument{
								Variadic: false,
								IsReference: false,
								Expr:  &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Assign{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.BitwiseAnd{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.BitwiseOr{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.BitwiseXor{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Concat{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Div{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Minus{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Mod{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Mul{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Plus{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.Pow{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.ShiftLeft{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
				Expr: &assign_op.ShiftRight{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
