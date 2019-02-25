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
)

func TestPrintFile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "\t")
	p.Print(&node.Root{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
			},
			&stmt.Class{
				Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
				ClassName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Bar"},
					},
				},
				Extends: &stmt.ClassExtends{
					ClassName: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Baz"},
						},
					},
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
						MethodName: &node.Identifier{Value: "greet"},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{
								&stmt.Echo{
									Exprs: []node.Node{
										&scalar.String{Value: "'Hello world'"},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	expected := `<?php
namespace Foo;
abstract class Bar extends Baz
{
	public function greet()
	{
		echo 'Hello world';
	}
}
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintFileInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&node.Root{
		Stmts: []node.Node{
			&stmt.InlineHtml{Value: "<div>HTML</div>"},
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "\"LBL\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "hello world\n"},
					},
				},
			},
		},
	})

	expected := `<div>HTML</div><?php
<<<"LBL"
hello world
LBL;
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// node

func TestPrintIdentifier(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&node.Identifier{Value: "test"})

	if o.String() != `test` {
		t.Errorf("TestPrintIdentifier is failed\n")
	}
}

func TestPrintParameter(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&node.Parameter{
		ByRef:        false,
		Variadic:     true,
		VariableType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Variable:     &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		DefaultValue: &scalar.String{Value: "'default'"},
	})

	expected := "\\Foo ...$var = 'default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNullable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&node.Nullable{
		Expr: &node.Parameter{
			ByRef:        false,
			Variadic:     true,
			VariableType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "var"}},
			DefaultValue: &scalar.String{Value: "'default'"},
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&node.Argument{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&node.Argument{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&name.NamePart{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&name.Name{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&name.FullyQualified{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&name.Relative{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.Lnumber{Value: "1"})

	if o.String() != `1` {
		t.Errorf("TestPrintScalarLNumber is failed\n")
	}
}

func TestPrintScalarDNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.Dnumber{Value: ".1"})

	if o.String() != `.1` {
		t.Errorf("TestPrintScalarDNumber is failed\n")
	}
}

func TestPrintScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.String{Value: "'hello world'"})

	expected := `'hello world'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.EncapsedStringPart{Value: "hello world"})

	if o.String() != `hello world` {
		t.Errorf("TestPrintScalarEncapsedStringPart is failed\n")
	}
}

func TestPrintScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.Encapsed{
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{VarName: &node.Identifier{Value: "var"}},
			&scalar.EncapsedStringPart{Value: " world"},
		},
	})

	if o.String() != `"hello {$var} world"` {
		t.Errorf("TestPrintScalarEncapsed is failed\n")
	}
}

func TestPrintScalarHeredoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.Heredoc{
		Label: "LBL",
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{VarName: &node.Identifier{Value: "var"}},
			&scalar.EncapsedStringPart{Value: " world\n"},
		},
	})

	expected := `<<<LBL
hello {$var} world
LBL`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintScalarNowdoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.Heredoc{
		Label: "'LBL'",
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello world\n"},
		},
	})

	expected := `<<<'LBL'
hello world
LBL`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintScalarMagicConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&scalar.MagicConstant{Value: "__DIR__"})

	if o.String() != `__DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Assign{
		Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a = $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Reference{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.BitwiseAnd{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.BitwiseOr{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.BitwiseXor{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Concat{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Div{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Minus{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Mod{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Mul{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Plus{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.Pow{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.ShiftLeft{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&assign.ShiftRight{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.BitwiseAnd{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.BitwiseOr{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.BitwiseXor{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.BooleanAnd{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.BooleanOr{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Coalesce{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Concat{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Div{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Equal{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.GreaterOrEqual{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Greater{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Identical{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.LogicalAnd{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.LogicalOr{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.LogicalXor{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Minus{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Mod{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Mul{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.NotEqual{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.NotIdentical{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Plus{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Pow{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.ShiftLeft{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.ShiftRight{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.SmallerOrEqual{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Smaller{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&binary.Spaceship{
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

func TestPrintArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.Array{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(array)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintBool(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.Bool{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(bool)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintDouble(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.Double{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(float)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInt(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.Int{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(int)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintObject(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.Object{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(object)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.String{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `(string)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&cast.Unset{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ArrayDimFetch{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ArrayItem{
		Key: &scalar.String{Value: "'Hello'"},
		Val: &expr.Variable{VarName: &node.Identifier{Value: "world"}},
	})

	expected := `'Hello' => $world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprArrayItem(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ArrayItem{
		Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "world"}}},
	})

	expected := `&$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Array{
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "'Hello'"},
				Val: &expr.Variable{VarName: &node.Identifier{Value: "world"}},
			},
			&expr.ArrayItem{
				Key: &scalar.Lnumber{Value: "2"},
				Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}}},
			},
			&expr.ArrayItem{
				Val: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.BitwiseNot{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.BooleanNot{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ClassConstFetch{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Clone{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ClosureUse{
		Uses: []node.Node{
			&expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}}},
			&expr.Variable{VarName: &node.Identifier{Value: "bar"}},
		},
	})

	expected := `use (&$foo, $bar)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprClosure(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&expr.Closure{
				Static:     true,
				ReturnsRef: true,
				Params: []node.Node{
					&node.Parameter{
						ByRef:    true,
						Variadic: false,
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
					},
				},
				ClosureUse: &expr.ClosureUse{
					Uses: []node.Node{
						&expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
						&expr.Variable{VarName: &node.Identifier{Value: "b"}},
					},
				},
				ReturnType: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
				},
			},
		},
	})

	expected := `namespace {
    static function &(&$var) use (&$a, $b): \Foo {
        $a;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ConstFetch{
		Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}},
	})

	expected := "null"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Empty{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `empty($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrettyPrinterrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ErrorSuppress{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `@$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintEval(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Eval{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `eval($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExit(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Exit{Die: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `exit($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintDie(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Exit{Die: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `die($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintFunctionCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.FunctionCall{
		Function: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
		ArgumentList: &node.ArgumentList{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Include{Expr: &scalar.String{Value: "'path'"}})

	expected := `include 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.IncludeOnce{Expr: &scalar.String{Value: "'path'"}})

	expected := `include_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.InstanceOf{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Isset{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.List{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.MethodCall{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
		Method:   &node.Identifier{Value: "bar"},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
				&node.Argument{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.New{
		Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
				&node.Argument{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.PostDec{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.PostInc{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.PreDec{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.PreInc{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Print{Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `print($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.PropertyFetch{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
		Property: &node.Identifier{Value: "bar"},
	})

	expected := `$foo->bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Reference{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
	})

	expected := `&$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintRequire(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Require{Expr: &scalar.String{Value: "'path'"}})

	expected := `require 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.RequireOnce{Expr: &scalar.String{Value: "'path'"}})

	expected := `require_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ShellExec{
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{VarName: &node.Identifier{Value: "world"}},
			&scalar.EncapsedStringPart{Value: "!"},
		},
	})

	expected := "`hello {$world}!`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExprShortArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ShortArray{
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "'Hello'"},
				Val: &expr.Variable{VarName: &node.Identifier{Value: "world"}},
			},
			&expr.ArrayItem{
				Key: &scalar.Lnumber{Value: "2"},
				Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}}},
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.ShortList{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.StaticCall{
		Class: &node.Identifier{Value: "Foo"},
		Call:  &node.Identifier{Value: "bar"},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
				&node.Argument{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.StaticPropertyFetch{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Ternary{
		Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "b"}},
	})

	expected := `$a ?: $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintTernaryFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Ternary{
		Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "b"}},
		IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "c"}},
	})

	expected := `$a ? $b : $c`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUnaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.UnaryMinus{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.UnaryPlus{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "var"}}})

	expected := `$$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.YieldFrom{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Yield{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&expr.Yield{
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

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.AltElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
			},
		},
	})

	expected := `elseif ($a) :
    $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltElseIfEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.AltElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{},
	})

	expected := `elseif ($a) :`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltElse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.AltElse{
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
			},
		},
	})

	expected := `else :
    $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltElseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.AltElse{
		Stmt: &stmt.StmtList{},
	})

	expected := `else :`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.AltFor{
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
			},
		},
	})

	expected := `namespace {
    for ($a; $b; $c) :
        $d;
    endfor;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.AltForeach{
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "var"}},
				Key:      &expr.Variable{VarName: &node.Identifier{Value: "key"}},
				Variable: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "val"}}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "d"}}},
					},
				},
			},
		},
	})

	expected := `namespace {
    foreach ($var as $key => &$val) :
        $d;
    endforeach;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltIf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.AltIf{
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
			},
		},
	})

	expected := `namespace {
    if ($a) :
        $d;
    elseif ($b) :
        $b;
    elseif ($c) :
    else :
        $b;
    endif;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtAltSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
				CaseList: &stmt.CaseList{
					Cases: []node.Node{
						&stmt.Case{
							Cond: &scalar.String{Value: "'a'"},
							Stmts: []node.Node{
								&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							},
						},
						&stmt.Case{
							Cond: &scalar.String{Value: "'b'"},
							Stmts: []node.Node{
								&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
							},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    switch ($var) :
        case 'a':
            $a;
        case 'b':
            $b;
    endswitch;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintAltWhile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.AltWhile{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
					},
				},
			},
		},
	})

	expected := `namespace {
    while ($a) :
        $b;
    endwhile;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtBreak(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Break{
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := "break 1;"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtCase(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Case{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
		},
	})

	expected := `case $a:
    $a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtCaseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Case{
		Cond:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmts: []node.Node{},
	})

	expected := "case $a:"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtCatch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Catch{
				Types: []node.Node{
					&name.Name{Parts: []node.Node{&name.NamePart{Value: "Exception"}}},
					&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "RuntimeException"}}},
				},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
				},
			},
		},
	})

	expected := `namespace {
    catch (Exception | \RuntimeException $e) {
        $a;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.ClassMethod{
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
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			},
		},
	})

	expected := `public function &foo(?int &$a = null, ...$b): void
{
    $a;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
func TestPrintStmtAbstractClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.ClassMethod{
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
		Stmt:       &stmt.Nop{},
	})

	expected := `public function &foo(?int &$a = null, ...$b): void;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Class{
				Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
				ClassName: &node.Identifier{Value: "Foo"},
				Extends: &stmt.ClassExtends{
					ClassName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
				},
				Implements: &stmt.ClassImplements{
					InterfaceNames: []node.Node{
						&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
						&name.Name{Parts: []node.Node{&name.NamePart{Value: "Quuz"}}},
					},
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Modifiers: []node.Node{&node.Identifier{Value: "public"}},
						Consts: []node.Node{
							&stmt.Constant{
								ConstantName: &node.Identifier{Value: "FOO"},
								Expr:         &scalar.String{Value: "'bar'"},
							},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    abstract class Foo extends Bar implements Baz, Quuz
    {
        public const FOO = 'bar';
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtAnonymousClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Class{
				Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
				ArgumentList: &node.ArgumentList{
					Arguments: []node.Node{
						&node.Argument{
							Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						},
						&node.Argument{
							Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
					},
				},
				Extends: &stmt.ClassExtends{
					ClassName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
				},
				Implements: &stmt.ClassImplements{
					InterfaceNames: []node.Node{
						&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
						&name.Name{Parts: []node.Node{&name.NamePart{Value: "Quuz"}}},
					},
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Modifiers: []node.Node{&node.Identifier{Value: "public"}},
						Consts: []node.Node{
							&stmt.Constant{
								ConstantName: &node.Identifier{Value: "FOO"},
								Expr:         &scalar.String{Value: "'bar'"},
							},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    abstract class($a, $b) extends Bar implements Baz, Quuz
    {
        public const FOO = 'bar';
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtClassConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.ClassConstList{
		Modifiers: []node.Node{&node.Identifier{Value: "public"}},
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "'a'"},
			},
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "BAR"},
				Expr:         &scalar.String{Value: "'b'"},
			},
		},
	})

	expected := `public const FOO = 'a', BAR = 'b';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Constant{
		ConstantName: &node.Identifier{Value: "FOO"},
		Expr:         &scalar.String{Value: "'BAR'"},
	})

	expected := "FOO = 'BAR'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtContinue(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Continue{
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `continue 1;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDeclareStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						ConstantName: &node.Identifier{Value: "FOO"},
						Expr:         &scalar.String{Value: "'bar'"},
					},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Nop{},
					},
				},
			},
		},
	})

	expected := `{
    declare(FOO = 'bar') {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDeclareExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						ConstantName: &node.Identifier{Value: "FOO"},
						Expr:         &scalar.String{Value: "'bar'"},
					},
				},
				Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
			},
		},
	})

	expected := `{
    declare(FOO = 'bar')
        'bar';
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDeclareNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Declare{
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "'bar'"},
			},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `declare(FOO = 'bar');`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDefalut(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Default{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
		},
	})

	expected := `default:
    $a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDefalutEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Default{
		Stmts: []node.Node{},
	})

	expected := `default:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDo_Expression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Do{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.Expression{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	})

	expected := `namespace {
    do
        $a;
    while (1);
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtDo_StmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Do{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
					},
				},
			},
		},
	})

	expected := `namespace {
    do {
        $a;
    } while (1);
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtEcho(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Echo{
		Exprs: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
	})

	expected := `echo $a, $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseIfStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.ElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `elseif ($a) {
    ;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseIfExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.ElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := `elseif ($a)
    'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.ElseIf{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Nop{},
	})

	expected := `elseif ($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Else{
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `else {
    ;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Else{
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := `else
    'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtElseNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Else{
		Stmt: &stmt.Nop{},
	})

	expected := `else;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}})

	expected := `$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Finally{
				Stmts: []node.Node{
					&stmt.Nop{},
				},
			},
		},
	})

	expected := `namespace {
    finally {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.For{
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
			},
		},
	})

	expected := `namespace {
    for ($a, $b; $c, $d; $e, $f) {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.For{
				Init: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
				Cond: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
				Loop: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "c"}},
				},
				Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
			},
		},
	})

	expected := `namespace {
    for ($a; $b; $c)
        'bar';
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.For{
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

	expected := `for ($a; $b; $c);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForeachStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Foreach{
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Nop{},
					},
				},
			},
		},
	})

	expected := `namespace {
    foreach ($a as $b) {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForeachExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Foreach{
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Key:      &expr.Variable{VarName: &node.Identifier{Value: "k"}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				Stmt:     &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
			},
		},
	})

	expected := `namespace {
    foreach ($a as $k => $v)
        'bar';
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtForeachNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Foreach{
		Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Key:      &expr.Variable{VarName: &node.Identifier{Value: "k"}},
		Variable: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}}},
		Stmt:     &stmt.Nop{},
	})

	expected := `foreach ($a as $k => &$v);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Function{
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
			},
		},
	})

	expected := `namespace {
    function &foo(&$var): \Foo {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGlobal(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Global{
		Vars: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
	})

	expected := `global $a, $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGoto(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Goto{
		Label: &node.Identifier{Value: "FOO"},
	})

	expected := `goto FOO;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.GroupUse{
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

	expected := `use function Foo\{Bar as Baz, Quuz};`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.HaltCompiler{})

	expected := `__halt_compiler();`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.If{
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
			},
		},
	})

	expected := `namespace {
    if ($a)
        $b;
    elseif ($c) {
        $d;
    }
    elseif ($e);
    else
        $f;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{
							Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    if ($a) {
        $b;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.If{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Nop{},
	})

	expected := `if ($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.InlineHtml{
		Value: "test",
	})

	expected := `?>test<?php`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInterface(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Interface{
				InterfaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
				Extends: &stmt.InterfaceExtends{
					InterfaceNames: []node.Node{
						&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
						&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
					},
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
						MethodName: &node.Identifier{Value: "foo"},
						Params:     []node.Node{},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{
								&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    interface Foo extends Bar, Baz
    {
        public function foo()
        {
            $a;
        }
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintLabel(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Label{
		LabelName: &node.Identifier{Value: "FOO"},
	})

	expected := `FOO:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNamespace(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
	})

	expected := `namespace Foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNamespaceWithStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
				},
			},
		},
	})

	expected := `{
    namespace Foo {
        $a;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Nop{})

	expected := `;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.PropertyList{
		Modifiers: []node.Node{
			&node.Identifier{Value: "public"},
			&node.Identifier{Value: "static"},
		},
		Properties: []node.Node{
			&stmt.Property{
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&stmt.Property{
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
	})

	expected := `public static $a, $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintProperty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Property{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expr:     &scalar.Lnumber{Value: "1"},
	})

	expected := `$a = 1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintReturn(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Return{
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `return 1;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStaticVar(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StaticVar{
		Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Expr:     &scalar.Lnumber{Value: "1"},
	})

	expected := `$a = 1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStatic(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Static{
		Vars: []node.Node{
			&stmt.StaticVar{
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
			&stmt.StaticVar{
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			},
		},
	})

	expected := `static $a, $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
		},
	})

	expected := `{
    $a;
    $b;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtListNested(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			&stmt.StmtList{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
					&stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "c"}}},
						},
					},
				},
			},
		},
	})

	expected := `{
    $a;
    {
        $b;
        {
            $c;
        }
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Switch{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
				CaseList: &stmt.CaseList{
					Cases: []node.Node{
						&stmt.Case{
							Cond: &scalar.String{Value: "'a'"},
							Stmts: []node.Node{
								&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							},
						},
						&stmt.Case{
							Cond: &scalar.String{Value: "'b'"},
							Stmts: []node.Node{
								&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
							},
						},
					},
				},
			},
		},
	})

	expected := `{
    switch ($var) {
        case 'a':
            $a;
        case 'b':
            $b;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtThrow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Throw{
		Expr: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
	})

	expected := `throw $var;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtTraitMethodRef(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.TraitMethodRef{
		Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Method: &node.Identifier{Value: "a"},
	})

	expected := `Foo::a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtTraitUseAlias(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.TraitUseAlias{
		Ref: &stmt.TraitMethodRef{
			Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			Method: &node.Identifier{Value: "a"},
		},
		Modifier: &node.Identifier{Value: "public"},
		Alias:    &node.Identifier{Value: "b"},
	})

	expected := `Foo::a as public b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtTraitUsePrecedence(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.TraitUsePrecedence{
		Ref: &stmt.TraitMethodRef{
			Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			Method: &node.Identifier{Value: "a"},
		},
		Insteadof: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
		},
	})

	expected := `Foo::a insteadof Bar, Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtTraitUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.TraitUse{
		Traits: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		},
	})

	expected := `use Foo, Bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtTraitAdaptations(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.TraitUse{
				Traits: []node.Node{
					&name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
					&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
				},
				TraitAdaptationList: &stmt.TraitAdaptationList{
					Adaptations: []node.Node{
						&stmt.TraitUseAlias{
							Ref: &stmt.TraitMethodRef{
								Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
								Method: &node.Identifier{Value: "a"},
							},
							Alias: &node.Identifier{Value: "b"},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    use Foo, Bar {
        Foo::a as b;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintTrait(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Trait{
				TraitName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
						MethodName: &node.Identifier{Value: "foo"},
						Params:     []node.Node{},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{
								&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							},
						},
					},
				},
			},
		},
	})

	expected := `namespace {
    trait Foo
    {
        public function foo()
        {
            $a;
        }
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtTry(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
				},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{Parts: []node.Node{&name.NamePart{Value: "Exception"}}},
							&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "RuntimeException"}}},
						},
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
						Stmts: []node.Node{
							&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
				Finally: &stmt.Finally{
					Stmts: []node.Node{
						&stmt.Nop{},
					},
				},
			},
		},
	})

	expected := `namespace {
    try {
        $a;
    }
    catch (Exception | \RuntimeException $e) {
        $b;
    }
    finally {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtUset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Unset{
		Vars: []node.Node{
			&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			&expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
	})

	expected := `unset($a, $b);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtUseList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.UseList{
		UseType: &node.Identifier{Value: "function"},
		Uses: []node.Node{
			&stmt.Use{
				Use:   &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
				Alias: &node.Identifier{Value: "Bar"},
			},
			&stmt.Use{
				Use: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
			},
		},
	})

	expected := `use function Foo as Bar, Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Use{
		UseType: &node.Identifier{Value: "function"},
		Use:     &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Alias:   &node.Identifier{Value: "Bar"},
	})

	expected := `function Foo as Bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintWhileStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
					},
				},
			},
		},
	})

	expected := `namespace {
    while ($a) {
        $a;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintWhileExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.Namespace{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.Expression{Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			},
		},
	})

	expected := `namespace {
    while ($a)
        $a;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintWhileNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&stmt.While{
		Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		Stmt: &stmt.Nop{},
	})

	expected := `while ($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
