package printer_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/printer"
	"github.com/z7zmey/php-parser/walker"
)

type testNode struct{}

func (t *testNode) Attributes() map[string]interface{} {
	return nil
}

func (t *testNode) Walk(v walker.Visitor) {}

func TestPrintWrongNode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	o := bytes.NewBufferString("")

	printer.Print(o, &testNode{})
}

// node

func TestPrintIdentifier(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &node.Identifier{Value: "test"})

	if o.String() != `test` {
		t.Errorf("TestPrintIdentifier is failed\n")
	}
}

func TestPrintParameter(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &node.Parameter{
		ByRef:        false,
		Variadic:     true,
		VariableType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Variable:     &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		DefaultValue: &scalar.String{Value: "default"},
	})

	expected := "\\Foo ...$var = 'default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNullable(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &node.Nullable{
		Expr: &node.Parameter{
			ByRef:        false,
			Variadic:     true,
			VariableType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			DefaultValue: &scalar.String{Value: "default"},
		},
	})

	expected := "?\\Foo ...$var = 'default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintArgument(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &node.Argument{
		IsReference: false,
		Variadic:    true,
		Expr:        &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := "...$var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
func TestPrintArgumentByRef(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &node.Argument{
		IsReference: true,
		Variadic:    false,
		Expr:        &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := "&$var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// name

func TestPrintNameNamePart(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &name.NamePart{
		Value: "foo",
	})

	expected := "foo"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNameName(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &name.Name{
		Parts: []node.Node{
			&name.NamePart{
				Value: "Foo",
			},
			&name.NamePart{
				Value: "Bar",
			},
		},
	})

	expected := "Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNameFullyQualified(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &name.FullyQualified{
		Parts: []node.Node{
			&name.NamePart{
				Value: "Foo",
			},
			&name.NamePart{
				Value: "Bar",
			},
		},
	})

	expected := "\\Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNameRelative(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &name.Relative{
		Parts: []node.Node{
			&name.NamePart{
				Value: "Foo",
			},
			&name.NamePart{
				Value: "Bar",
			},
		},
	})

	expected := "namespace\\Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// scalar

func TestPrintScalarLNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &scalar.Lnumber{Value: "1"})

	if o.String() != `1` {
		t.Errorf("TestPrintScalarLNumber is failed\n")
	}
}

func TestPrintScalarDNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &scalar.Dnumber{Value: ".1"})

	if o.String() != `.1` {
		t.Errorf("TestPrintScalarDNumber is failed\n")
	}
}

func TestPrintScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &scalar.String{Value: "hello world"})

	if o.String() != `'hello world'` {
		t.Errorf("TestPrintScalarString is failed\n")
	}
}

func TestPrintScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &scalar.EncapsedStringPart{Value: "hello world"})

	if o.String() != `hello world` {
		t.Errorf("TestPrintScalarEncapsedStringPart is failed\n")
	}
}

func TestPrintScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &scalar.Encapsed{
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{VarName: &node.Identifier{Value: "var"}},
			&scalar.EncapsedStringPart{Value: " world"},
		},
	})

	if o.String() != `"hello $var world"` {
		t.Errorf("TestPrintScalarEncapsed is failed\n")
	}
}

func TestPrintScalarMagicConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &scalar.MagicConstant{Value: "__DIR__"})

	if o.String() != `__DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Assign{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a = $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignRef(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.AssignRef{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a =& $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.BitwiseAnd{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a &= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.BitwiseOr{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a |= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.BitwiseXor{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a ^= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Concat{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a .= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Div{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a /= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Minus{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a -= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignMod(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Mod{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a %= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignMul(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Mul{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a *= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Plus{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a += $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignPow(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.Pow{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a **= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.ShiftLeft{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a <<= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAssignShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &assign.ShiftRight{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a >>= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// binary

func TestPrintBinaryBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.BitwiseAnd{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a & $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.BitwiseOr{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a | $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.BitwiseXor{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a ^ $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryBooleanAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.BooleanAnd{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a && $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryBooleanOr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.BooleanOr{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a || $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Coalesce{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a ?? $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Concat{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a . $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Div{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a / $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Equal{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a == $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryGreaterOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.GreaterOrEqual{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a >= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryGreater(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Greater{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a > $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Identical{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a === $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryLogicalAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.LogicalAnd{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a and $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryLogicalOr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.LogicalOr{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a or $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryLogicalXor(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.LogicalXor{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a xor $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Minus{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a - $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryMod(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Mod{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a % $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryMul(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Mul{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a * $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryNotEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.NotEqual{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a != $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryNotIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.NotIdentical{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a !== $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Plus{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a + $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryPow(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Pow{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a ** $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.ShiftLeft{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a << $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinaryShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.ShiftRight{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a >> $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinarySmallerOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.SmallerOrEqual{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a <= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinarySmaller(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Smaller{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a < $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBinarySpaceship(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &binary.Spaceship{
		Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a <=> $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// cast

func TestPrintCastArray(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastArray{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(array)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintCastBool(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastBool{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(bool)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintCastDouble(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastDouble{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(float)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintCastInt(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastInt{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(int)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintCastObject(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastObject{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(object)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintCastString(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastString{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(string)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintCastUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &cast.CastUnset{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(unset)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// expr

func TestPrintExprArrayDimFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ArrayDimFetch{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		Dim:      &scalar.Lnumber{Value: "1"},
	})

	expected := `$var[1]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprArrayItemWithKey(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ArrayItem{
		ByRef: false,
		Key:   &scalar.String{Value: "Hello"},
		Val:   &expr.Variable{VarName: &node.Identifier{Value: "world"}},
	})

	expected := `'Hello' => $world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprArrayItem(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ArrayItem{
		ByRef: true,
		Val:   &expr.Variable{VarName: &node.Identifier{Value: "world"}},
	})

	expected := `&$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprArray(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Array{
		Items: []node.Node{
			&expr.ArrayItem{
				ByRef: false,
				Key:   &scalar.String{Value: "Hello"},
				Val:   &expr.Variable{VarName: &node.Identifier{Value: "world"}},
			},
			&expr.ArrayItem{
				ByRef: true,
				Key:   &scalar.Lnumber{Value: "2"},
				Val:   &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			},
			&expr.ArrayItem{
				ByRef: false,
				Val:   &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			},
		},
	})

	expected := `array('Hello' => $world, 2 => &$var, $var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprBitwiseNot(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.BitwiseNot{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `~$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprBooleanNot(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.BooleanNot{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `!$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprClassConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ClassConstFetch{
		Class:        &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		ConstantName: &node.Identifier{Value: "CONST"},
	})

	expected := `$var::CONST`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprClone(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Clone{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `clone $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprClosureUse(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ClosureUse{
		ByRef:    true,
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `&$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprClosure(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Closure{
		Static:     true,
		ReturnsRef: true,
		Params: []node.Node{
			&node.Parameter{
				ByRef:    true,
				Variadic: false,
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			},
		},
		Uses: []node.Node{
			&expr.ClosureUse{
				ByRef:    true,
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&expr.ClosureUse{
				ByRef:    false,
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
		ReturnType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Stmts:      []node.Node{},
	})

	expected := "static function &(&$var) use (&$a, $b): \\Foo {}"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ConstFetch{
		Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}},
	})

	expected := "null"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintDie(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Die{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `die($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Empty{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `empty($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintErrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ErrorSuppress{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `@$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintEval(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Eval{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `eval($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExit(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Exit{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `exit($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintFunctionCall(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.FunctionCall{
		Function: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		Arguments: []node.Node{
			&node.Argument{
				IsReference: true,
				Expr:        &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&node.Argument{
				Variadic: true,
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
			},
		},
	})

	expected := `$var(&$a, ...$b, $c)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInclude(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Include{Expr: &scalar.String{Value: "path"}})

	expected := `include 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.IncludeOnce{Expr: &scalar.String{Value: "path"}})

	expected := `include_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.InstanceOf{
		Expr:  &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
	})

	expected := `$var instanceof Foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIsset(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Isset{
		Variables: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
	})

	expected := `isset($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintList(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.List{
		Items: []node.Node{
			&expr.ArrayItem{
				Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&expr.ArrayItem{
				Val: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
						&expr.ArrayItem{
							Val: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						},
					},
				},
			},
		},
	})

	expected := `list($a, list($b, $c))`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintMethodCall(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.MethodCall{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
		Method:   &node.Identifier{Value: "bar"},
		Arguments: []node.Node{
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
	})

	expected := `$foo->bar($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNew(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.New{
		Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Arguments: []node.Node{
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
	})

	expected := `new Foo($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPostDec(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.PostDec{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `$var--`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPostInc(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.PostInc{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `$var++`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPreDec(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.PreDec{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `--$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPreInc(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.PreInc{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `++$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPrint(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Print{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `print($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.PropertyFetch{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
		Property: &node.Identifier{Value: "bar"},
	})

	expected := `$foo->bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintRequire(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Require{Expr: &scalar.String{Value: "path"}})

	expected := `require 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.RequireOnce{Expr: &scalar.String{Value: "path"}})

	expected := `require_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ShellExec{
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{VarName: &node.Identifier{Value: "world"}},
			&scalar.EncapsedStringPart{Value: "!"},
		},
	})

	expected := "`hello $world!`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprShortArray(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ShortArray{
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "Hello"},
				Val: &expr.Variable{VarName: &node.Identifier{Value: "world"}},
			},
			&expr.ArrayItem{
				ByRef: true,
				Key:   &scalar.Lnumber{Value: "2"},
				Val:   &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			},
			&expr.ArrayItem{
				Val: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			},
		},
	})

	expected := `['Hello' => $world, 2 => &$var, $var]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintShortList(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.ShortList{
		Items: []node.Node{
			&expr.ArrayItem{
				Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&expr.ArrayItem{
				Val: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
						&expr.ArrayItem{
							Val: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						},
					},
				},
			},
		},
	})

	expected := `[$a, list($b, $c)]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStaticCall(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.StaticCall{
		Class: &node.Identifier{Value: "Foo"},
		Call:  &node.Identifier{Value: "bar"},
		Arguments: []node.Node{
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
	})

	expected := `Foo::bar($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStaticPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.StaticPropertyFetch{
		Class:    &node.Identifier{Value: "Foo"},
		Property: &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
	})

	expected := `Foo::$bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintTernary(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Ternary{
		Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a ?: $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintTernaryFull(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Ternary{
		Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "b"}},
		IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "c"}},
	})

	expected := `$a ? $b : $c`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUnaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.UnaryMinus{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `-$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUnaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.UnaryPlus{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `+$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintVariable(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `$$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.YieldFrom{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `yield from $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintYield(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Yield{
		Value: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `yield $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintYieldFull(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &expr.Yield{
		Key:   &expr.Variable{VarName: &node.Identifier{Value: "k"}},
		Value: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `yield $k => $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// stmt

func TestPrintAltElseIf(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
			},
		},
	})

	expected := "elseif ($a) :\n$b;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltElse(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltElse{
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
			},
		},
	})

	expected := "else :\n$b;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltFor(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltFor{
		Init: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		Cond: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
		Loop: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "c"}},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "d"}}},
			},
		},
	})

	expected := "for ($a; $b; $c) :\n$d;\nendfor;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltForeach{
		ByRef:    true,
		Expr:     &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		Key:      &expr.Variable{VarName: &node.Identifier{Value: "key"}},
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "val"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "d"}}},
			},
		},
	})

	expected := "foreach ($var as $key => &$val) :\n$d;\nendforeach;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltIf(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "d"}}},
			},
		},
		ElseIf: []node.Node{
			&stmt.AltElseIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
					},
				},
			},
			&stmt.AltElseIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
				Stmt: &stmt.StmtList{},
			},
		},
		Else: &stmt.AltElse{
			Stmt: &stmt.StmtList{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
				},
			},
		},
	})

	expected := `if ($a) :
$d;
elseif ($b) :
$b;
elseif ($c) :
else :
$b;
endif;
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtAltSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltSwitch{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		Cases: []node.Node{
			&stmt.Case{
				Cond: &scalar.String{Value: "a"},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
				},
			},
			&stmt.Case{
				Cond: &scalar.String{Value: "b"},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
				},
			},
		},
	})

	expected := `switch ($var) :
case 'a':
$a;
case 'b':
$b;
endswitch;
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltWhile(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.AltWhile{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
			},
		},
	})

	expected := "while ($a) :\n$b;\nendwhile;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtBreak(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Break{
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := "break 1;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtCase(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Case{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
		},
	})

	expected := "case $a:\n$a;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtCatch(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Catch{
		Types: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Exception"}}},
			&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "RuntimeException"}}},
		},
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
		},
	})

	expected := `catch (Exception | \RuntimeException $e) {
$a;
}
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.ClassMethod{
		Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
		ReturnsRef: true,
		MethodName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        true,
				VariableType: &node.Nullable{Expr: &name.Name{Parts: []node.Node{&name.NamePart{Value: "int"}}}},
				Variable:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
			},
			&node.Parameter{
				Variadic: true,
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
		ReturnType: &name.Name{Parts: []node.Node{&name.NamePart{Value: "void"}}},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
		},
	})

	expected := `public function &foo(?int &$a = null, ...$b): void
{
$a;
}
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtClass(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Class{
		Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
		ClassName: &node.Identifier{Value: "Foo"},
		Extends:   &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		Implements: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Quuz"}}},
		},
		Stmts: []node.Node{
			&stmt.ClassConstList{
				Modifiers: []node.Node{&node.Identifier{Value: "public"}},
				Consts: []node.Node{
					&stmt.Constant{
						ConstantName: &node.Identifier{Value: "FOO"},
						Expr:         &scalar.String{Value: "bar"},
					},
				},
			},
		},
	})

	expected := `abstract class Foo extends Bar implements Baz, Quuz
{
public const FOO = 'bar';
}
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtAnonymousClass(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Class{
		Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
		Args: []node.Node{
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&node.Argument{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
		Extends: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		Implements: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Quuz"}}},
		},
		Stmts: []node.Node{
			&stmt.ClassConstList{
				Modifiers: []node.Node{&node.Identifier{Value: "public"}},
				Consts: []node.Node{
					&stmt.Constant{
						ConstantName: &node.Identifier{Value: "FOO"},
						Expr:         &scalar.String{Value: "bar"},
					},
				},
			},
		},
	})

	expected := `abstract class($a, $b) extends Bar implements Baz, Quuz
{
public const FOO = 'bar';
}
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtClassConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.ClassConstList{
		Modifiers: []node.Node{&node.Identifier{Value: "public"}},
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "a"},
			},
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "BAR"},
				Expr:         &scalar.String{Value: "b"},
			},
		},
	})

	expected := "public const FOO = 'a', BAR = 'b';\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Constant{
		ConstantName: &node.Identifier{Value: "FOO"},
		Expr:         &scalar.String{Value: "BAR"},
	})

	expected := "FOO = 'BAR'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtContinue(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Continue{
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := "continue 1;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDeclareStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Declare{
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "bar"},
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := "declare(FOO = 'bar') {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDeclareExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Declare{
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "bar"},
			},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "bar"}},
	})

	expected := "declare(FOO = 'bar')\n'bar';\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDeclareNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Declare{
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "bar"},
			},
		},
		Stmt: &stmt.Nop{},
	})

	expected := "declare(FOO = 'bar');\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDefalut(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Default{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
		},
	})

	expected := "default:\n$a;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDo_Expression(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Do{
		Cond: &scalar.Lnumber{Value: "1"},
		Stmt: &stmt.Expression{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
	})

	expected := "do\n$a;\nwhile (1);\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDo_StmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Do{
		Cond: &scalar.Lnumber{Value: "1"},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			},
		},
	})

	expected := "do {\n$a;\n} while (1);\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtEcho(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Echo{
		Exprs: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
	})

	expected := "echo $a, $b;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseIfStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.ElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := "elseif ($a) {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseIfExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.ElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "bar"}},
	})

	expected := "elseif ($a)\n'bar';\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.ElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Nop{},
	})

	expected := "elseif ($a);\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Else{
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := "else {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Else{
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "bar"}},
	})

	expected := "else\n'bar';\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Else{
		Stmt: &stmt.Nop{},
	})

	expected := "else;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}})

	expected := "$a;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Finally{
		Stmts: []node.Node{
			&stmt.Nop{},
		},
	})

	expected := "finally {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.For{
		Init: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
		Cond: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "c"}},
			&expr.Variable{VarName: &node.Identifier{Value: "d"}},
		},
		Loop: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "e"}},
			&expr.Variable{VarName: &node.Identifier{Value: "f"}},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := "for ($a, $b; $c, $d; $e, $f) {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.For{
		Init: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		Cond: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
		Loop: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "c"}},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "bar"}},
	})

	expected := "for ($a; $b; $c)\n'bar';\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.For{
		Init: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		Cond: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
		Loop: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "c"}},
		},
		Stmt: &stmt.Nop{},
	})

	expected := "for ($a; $b; $c);\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForeachStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Foreach{
		Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := "foreach ($a as $b) {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForeachExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Foreach{
		Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Key:      &expr.Variable{VarName: &node.Identifier{Value: "k"}},
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
		Stmt:     &stmt.Expression{Expr: &scalar.String{Value: "bar"}},
	})

	expected := "foreach ($a as $k => $v)\n'bar';\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForeachNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Foreach{
		ByRef:    true,
		Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Key:      &expr.Variable{VarName: &node.Identifier{Value: "k"}},
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
		Stmt:     &stmt.Nop{},
	})

	expected := "foreach ($a as $k => &$v);\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Function{
		ReturnsRef:   true,
		FunctionName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:    true,
				Variadic: false,
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			},
		},
		ReturnType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Stmts: []node.Node{
			&stmt.Nop{},
		},
	})

	expected := "function &foo(&$var): \\Foo {\n;\n}\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGlobal(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Global{
		Vars: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
	})

	expected := "global $a, $b;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGoto(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Goto{
		Label: &node.Identifier{Value: "FOO"},
	})

	expected := "goto FOO;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.GroupUse{
		UseType: &node.Identifier{Value: "function"},
		Prefix:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		UseList: []node.Node{
			&stmt.Use{
				Use:   &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
				Alias: &node.Identifier{Value: "Baz"},
			},
			&stmt.Use{
				Use: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Quuz"}}},
			},
		},
	})

	expected := "use function Foo\\{Bar as Baz, Quuz};\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.HaltCompiler{})

	expected := "__halt_compiler();\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.If{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Expression{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
		ElseIf: []node.Node{
			&stmt.ElseIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{
							Expr: &expr.Variable{VarName: &node.Identifier{Value: "d"}},
						},
					},
				},
			},
			&stmt.ElseIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
				Stmt: &stmt.Nop{},
			},
		},
		Else: &stmt.Else{
			Stmt: &stmt.Expression{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "f"}},
			},
		},
	})

	expected := `if ($a)
$b;
elseif ($c) {
$d;
}
elseif ($e);
else
$f;
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.If{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	})

	expected := `if ($a) {
$b;
}
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.If{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Nop{},
	})

	expected := `if ($a);
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.InlineHtml{
		Value: "test",
	})

	expected := `?>test<?php
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
		},
	})

	expected := "$a;\n$b;\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNop(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Nop{})

	expected := ";\n"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	printer.Print(o, &stmt.Use{
		UseType: &node.Identifier{Value: "function"},
		Use:     &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Alias:   &node.Identifier{Value: "Bar"},
	})

	expected := "function Foo as Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
