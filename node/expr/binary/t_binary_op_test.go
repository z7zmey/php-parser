package binary_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/binary"
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

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.BitwiseAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestBitwiseOr(t *testing.T) {
	src := `<? $a | $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.BitwiseOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestBitwiseXor(t *testing.T) {
	src := `<? $a ^ $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.BitwiseXor{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestBooleanAnd(t *testing.T) {
	src := `<? $a && $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.BooleanAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestBooleanOr(t *testing.T) {
	src := `<? $a || $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.BooleanOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestCoalesce(t *testing.T) {
	src := `<? $a ?? $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Coalesce{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestConcat(t *testing.T) {
	src := `<? $a . $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Concat{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestDiv(t *testing.T) {
	src := `<? $a / $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Div{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestEqual(t *testing.T) {
	src := `<? $a == $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Equal{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestGreaterOrEqual(t *testing.T) {
	src := `<? $a >= $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.GreaterOrEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestGreater(t *testing.T) {
	src := `<? $a > $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Greater{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestIdentical(t *testing.T) {
	src := `<? $a === $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Identical{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestLogicalAnd(t *testing.T) {
	src := `<? $a and $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.LogicalAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestLogicalOr(t *testing.T) {
	src := `<? $a or $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.LogicalOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestLogicalXor(t *testing.T) {
	src := `<? $a xor $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.LogicalXor{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestMinus(t *testing.T) {
	src := `<? $a - $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Minus{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestMod(t *testing.T) {
	src := `<? $a % $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Mod{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestMul(t *testing.T) {
	src := `<? $a * $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Mul{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestNotEqual(t *testing.T) {
	src := `<? $a != $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.NotEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestNotIdentical(t *testing.T) {
	src := `<? $a !== $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.NotIdentical{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestPlus(t *testing.T) {
	src := `<? $a + $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Plus{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestPow(t *testing.T) {
	src := `<? $a ** $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Pow{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestShiftLeft(t *testing.T) {
	src := `<? $a << $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.ShiftLeft{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestShiftRight(t *testing.T) {
	src := `<? $a >> $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.ShiftRight{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestSmallerOrEqual(t *testing.T) {
	src := `<? $a <= $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.SmallerOrEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestSmaller(t *testing.T) {
	src := `<? $a < $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Smaller{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestSpaceship(t *testing.T) {
	src := `<? $a <=> $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &binary.Spaceship{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
