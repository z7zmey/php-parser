package binary_op_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/binary_op"
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

func TestBitwiseAnd(t *testing.T) {
	src := `<? $a & $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.BitwiseAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a | $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.BitwiseOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a ^ $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.BitwiseXor{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBooleanAnd(t *testing.T) {
	src := `<? $a && $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.BooleanAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBooleanOr(t *testing.T) {
	src := `<? $a || $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.BooleanOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestCoalesce(t *testing.T) {
	src := `<? $a ?? $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Coalesce{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestConcat(t *testing.T) {
	src := `<? $a . $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Concat{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a / $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Div{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestEqual(t *testing.T) {
	src := `<? $a == $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Equal{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestGreaterOrEqual(t *testing.T) {
	src := `<? $a >= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.GreaterOrEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestGreater(t *testing.T) {
	src := `<? $a > $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Greater{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestIdentical(t *testing.T) {
	src := `<? $a === $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Identical{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestLogicalAnd(t *testing.T) {
	src := `<? $a and $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.LogicalAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestLogicalOr(t *testing.T) {
	src := `<? $a or $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.LogicalOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestLogicalXor(t *testing.T) {
	src := `<? $a xor $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.LogicalXor{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a - $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Minus{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a % $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Mod{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a * $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Mul{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestNotEqual(t *testing.T) {
	src := `<? $a != $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.NotEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestNotIdentical(t *testing.T) {
	src := `<? $a !== $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.NotIdentical{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a + $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Plus{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a ** $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Pow{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a << $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.ShiftLeft{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
	src := `<? $a >> $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.ShiftRight{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestSmallerOrEqual(t *testing.T) {
	src := `<? $a <= $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.SmallerOrEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestSmaller(t *testing.T) {
	src := `<? $a < $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Smaller{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestSpaceship(t *testing.T) {
	src := `<? $a <=> $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary_op.Spaceship{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
