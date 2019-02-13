package printer_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/freefloating"
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

func TestPrinterPrintFile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

	expected := `<?php namespace Foo;abstract class Bar extends Baz{public function greet(){echo 'Hello world';}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintFileInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Root{
		Stmts: []node.Node{
			&stmt.InlineHtml{Value: "<div>HTML</div>"},
			&stmt.Expression{
				Expr: &expr.Variable{
					FreeFloating: freefloating.Collection{
						freefloating.Start: []freefloating.String{
							{
								StringType: freefloating.TokenType,
								Value:      "$",
							},
						},
					},
					VarName: &node.Identifier{
						Value: "a",
					},
				},
			},
			&stmt.InlineHtml{Value: "<div>HTML</div>"},
			&stmt.Expression{
				Expr: &expr.Variable{
					FreeFloating: freefloating.Collection{
						freefloating.Start: []freefloating.String{
							{
								StringType: freefloating.TokenType,
								Value:      "$",
							},
						},
					},
					VarName: &node.Identifier{
						Value: "a",
					},
				},
			},
		},
	})

	expected := `<div>HTML</div><?php $a;?><div>HTML</div><?php $a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// node

func TestPrinterPrintIdentifier(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	n := &node.Identifier{
		Value: "test",
	}
	p.Print(n)

	expected := `test`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintParameter(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Parameter{
		ByRef:    false,
		Variadic: true,
		VariableType: &name.FullyQualified{
			Parts: []node.Node{
				&name.NamePart{
					Value: "Foo",
				},
			},
		},
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		DefaultValue: &scalar.String{
			Value: "'default'",
		},
	})

	expected := "\\Foo...$var='default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNullable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Nullable{
		Expr: &node.Parameter{
			ByRef:    true,
			Variadic: false,
			VariableType: &name.FullyQualified{
				Parts: []node.Node{
					&name.NamePart{
						Value: "Foo",
					},
				},
			},
			Variable: &expr.Variable{
				VarName: &node.Identifier{
					Value: "var",
				},
			},
			DefaultValue: &scalar.String{
				Value: "'default'",
			},
		},
	})

	expected := "?\\Foo&$var='default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintArgument(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Argument{
		IsReference: false,
		Variadic:    true,
		Expr: &expr.Variable{
			VarName: &node.Identifier{
				Value: "var",
			},
		},
	})

	expected := "...$var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
func TestPrinterPrintArgumentByRef(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Argument{
		IsReference: true,
		Variadic:    false,
		Expr: &expr.Variable{
			VarName: &node.Identifier{
				Value: "var",
			},
		},
	})

	expected := "&$var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// name

func TestPrinterPrintNameNamePart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&name.NamePart{
		Value: "foo",
	})

	expected := "foo"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameName(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintNameFullyQualified(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintNameRelative(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintScalarLNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Lnumber{
		Value: "1",
	})

	expected := "1"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarDNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Dnumber{
		Value: ".1",
	})

	expected := ".1"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.String{
		Value: "'hello world'",
	})

	expected := `'hello world'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.EncapsedStringPart{
		Value: "hello world",
	})

	expected := `hello world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Encapsed{
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{
				VarName: &node.Identifier{Value: "var"},
			},
			&scalar.EncapsedStringPart{Value: " world"},
		},
	})

	expected := `"hello $var world"`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarHeredoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Heredoc{
		Label: "LBL",
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{
				VarName: &node.Identifier{Value: "var"},
			},
			&scalar.EncapsedStringPart{Value: " world"},
		},
	})

	expected := `<<<LBL
hello $var world
LBL`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarNowdoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Heredoc{
		Label: "'LBL'",
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello world"},
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

func TestPrinterPrintScalarMagicConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.MagicConstant{
		Value: "__DIR__",
	})

	if o.String() != `__DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrinterPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Assign{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Reference{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a=&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.BitwiseAnd{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a&=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.BitwiseOr{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a|=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.BitwiseXor{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a^=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Concat{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a.=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Div{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a/=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Minus{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a-=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Mod{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a%=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMul(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Mul{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a*=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Plus{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a+=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Pow{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a**=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.ShiftLeft{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a<<=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.ShiftRight{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a>>=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// binary

func TestPrinterPrintBinaryBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BitwiseAnd{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BitwiseOr{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a|$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BitwiseXor{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a^$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BooleanAnd{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a&&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BooleanOr{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a||$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Coalesce{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a??$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Concat{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a.$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Div{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a/$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Equal{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a==$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreaterOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.GreaterOrEqual{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a>=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreater(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Greater{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Identical{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a===$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.LogicalAnd{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a and $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.LogicalOr{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a or $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.LogicalXor{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a xor $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Minus{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a-$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Mod{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a%$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMul(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Mul{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a*$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.NotEqual{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a!=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.NotIdentical{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a!==$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Plus{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a+$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Pow{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a**$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.ShiftLeft{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a<<$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.ShiftRight{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a>>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmallerOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.SmallerOrEqual{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a<=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmaller(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Smaller{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a<$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySpaceship(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Spaceship{
		Left: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a<=>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// cast

func TestPrinterPrintArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Array{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(array)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBool(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Bool{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(boolean)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDouble(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Double{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(float)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInt(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Int{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(integer)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintObject(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Object{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(object)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.String{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(string)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Unset{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `(unset)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// expr

func TestPrinterPrintExprArrayDimFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ArrayDimFetch{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		Dim: &scalar.Lnumber{Value: "1"},
	})

	expected := `$var[1]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemWithKey(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ArrayItem{
		Key: &scalar.String{Value: "'Hello'"},
		Val: &expr.Variable{
			VarName: &node.Identifier{Value: "world"},
		},
	})

	expected := `'Hello'=>$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItem(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ArrayItem{
		Val: &expr.Reference{Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "world"},
		}},
	})

	expected := `&$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Array{
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "'Hello'"},
				Val: &expr.Variable{
					VarName: &node.Identifier{Value: "world"},
				},
			},
			&expr.ArrayItem{
				Key: &scalar.Lnumber{Value: "2"},
				Val: &expr.Reference{Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "var"},
				}},
			},
			&expr.ArrayItem{
				Val: &expr.Variable{
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
	})

	expected := `array('Hello'=>$world,2=>&$var,$var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBitwiseNot(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.BitwiseNot{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `~$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBooleanNot(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.BooleanNot{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `!$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClassConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ClassConstFetch{
		Class: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		ConstantName: &node.Identifier{
			Value: "CONST",
		},
	})

	expected := `$var::CONST`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClone(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Clone{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `clone $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosureUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ClosureUse{
		Uses: []node.Node{
			&expr.Reference{Variable: &expr.Variable{
				VarName: &node.Identifier{Value: "foo"},
			}},
			&expr.Variable{
				VarName: &node.Identifier{Value: "bar"},
			},
		},
	})

	expected := `use(&$foo,$bar)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosure(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Closure{
		Static:     true,
		ReturnsRef: true,
		Params: []node.Node{
			&node.Parameter{
				ByRef:    true,
				Variadic: false,
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
		ClosureUse: &expr.ClosureUse{
			Uses: []node.Node{
				&expr.Reference{Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				}},
				&expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
		ReturnType: &name.FullyQualified{
			Parts: []node.Node{&name.NamePart{Value: "Foo"}},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := `static function&(&$var)use(&$a,$b):\Foo{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ConstFetch{
		Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}},
	})

	expected := "null"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Empty{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `empty($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrettyPrinterrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ErrorSuppress{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `@$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEval(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Eval{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `eval($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExit(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Exit{
		Die: false,
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `exit $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDie(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Exit{
		Die: true,
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `die $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintFunctionCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.FunctionCall{
		Function: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					IsReference: true,
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Variadic: true,
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "c"},
					},
				},
			},
		},
	})

	expected := `$var(&$a,...$b,$c)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInclude(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Include{
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := `include 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.IncludeOnce{
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := `include_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.InstanceOf{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
	})

	expected := `$var instanceof Foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIsset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Isset{
		Variables: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := `isset($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.List{
		Items: []node.Node{
			&expr.ArrayItem{
				Val: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&expr.ArrayItem{
				Val: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{
								VarName: &node.Identifier{Value: "b"},
							},
						},
						&expr.ArrayItem{
							Val: &expr.Variable{
								VarName: &node.Identifier{Value: "c"},
							},
						},
					},
				},
			},
		},
	})

	expected := `list($a,list($b,$c))`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintMethodCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.MethodCall{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "foo"},
		},
		Method: &node.Identifier{Value: "bar"},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `$foo->bar($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNew(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.New{
		Class: &name.Name{
			Parts: []node.Node{
				&name.NamePart{
					Value: "Foo",
				},
			},
		},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `new Foo($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostDec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PostDec{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `$var--`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostInc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PostInc{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `$var++`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreDec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PreDec{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `--$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreInc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PreInc{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `++$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPrint(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Print{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `print $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PropertyFetch{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "foo"},
		},
		Property: &node.Identifier{Value: "bar"},
	})

	expected := `$foo->bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Reference{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "foo"},
		},
	})

	expected := `&$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequire(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Require{
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := `require 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.RequireOnce{
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := `require_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ShellExec{
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{
				VarName: &node.Identifier{Value: "world"},
			},
			&scalar.EncapsedStringPart{Value: "!"},
		},
	})

	expected := "`hello $world!`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprShortArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ShortArray{
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "'Hello'"},
				Val: &expr.Variable{
					VarName: &node.Identifier{Value: "world"},
				},
			},
			&expr.ArrayItem{
				Key: &scalar.Lnumber{Value: "2"},
				Val: &expr.Reference{Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "var"},
				}},
			},
			&expr.ArrayItem{
				Val: &expr.Variable{
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
	})

	expected := `['Hello'=>$world,2=>&$var,$var]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShortList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ShortList{
		Items: []node.Node{
			&expr.ArrayItem{
				Val: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&expr.ArrayItem{
				Val: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{
								VarName: &node.Identifier{Value: "b"},
							},
						},
						&expr.ArrayItem{
							Val: &expr.Variable{
								VarName: &node.Identifier{Value: "c"},
							},
						},
					},
				},
			},
		},
	})

	expected := `[$a,list($b,$c)]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.StaticCall{
		Class: &node.Identifier{Value: "Foo"},
		Call:  &node.Identifier{Value: "bar"},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `Foo::bar($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.StaticPropertyFetch{
		Class: &node.Identifier{Value: "Foo"},
		Property: &expr.Variable{
			VarName: &node.Identifier{Value: "bar"},
		},
	})

	expected := `Foo::$bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernary(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Ternary{
		Condition: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		IfFalse: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a?:$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernaryFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Ternary{
		Condition: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		IfTrue: &expr.Variable{
			VarName: &node.Identifier{Value: "b"},
		},
		IfFalse: &expr.Variable{
			VarName: &node.Identifier{Value: "c"},
		},
	})

	expected := `$a?$b:$c`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.UnaryMinus{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `-$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.UnaryPlus{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `+$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintVariable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Variable{
		VarName: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `$$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.YieldFrom{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `yield from $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYield(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Yield{
		Value: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `yield $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Yield{
		Key: &expr.Variable{
			VarName: &node.Identifier{Value: "k"},
		},
		Value: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `yield $k=>$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// stmt

func TestPrinterPrintAltElseIf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElseIf{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				}},
			},
		},
	})

	expected := `elseif($a):$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseIfEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElseIf{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{},
	})

	expected := `elseif($a):`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElse{
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				}},
			},
		},
	})

	expected := `else:$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElse{
		Stmt: &stmt.StmtList{},
	})

	expected := `else:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltFor{
		Init: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
		},
		Cond: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
		Loop: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "c"},
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "d"},
				}},
			},
		},
	})

	expected := `for($a;$b;$c):$d;endfor;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltForeach{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		Key: &expr.Variable{
			VarName: &node.Identifier{Value: "key"},
		},
		Variable: &expr.Reference{Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "val"},
		}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "d"},
				}},
			},
		},
	})

	expected := `foreach($var as $key=>&$val):$d;endforeach;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltIf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltIf{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "d"},
				}},
			},
		},
		ElseIf: []node.Node{
			&stmt.AltElseIf{
				Cond: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "b"},
						}},
					},
				},
			},
			&stmt.AltElseIf{
				Cond: &expr.Variable{
					VarName: &node.Identifier{Value: "c"},
				},
				Stmt: &stmt.StmtList{},
			},
		},
		Else: &stmt.AltElse{
			Stmt: &stmt.StmtList{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					}},
				},
			},
		},
	})

	expected := `if($a):$d;elseif($b):$b;elseif($c):else:$b;endif;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAltSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltSwitch{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		CaseList: &stmt.CaseList{
			Cases: []node.Node{
				&stmt.Case{
					Cond: &scalar.String{Value: "'a'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
				&stmt.Case{
					Cond: &scalar.String{Value: "'b'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "b"},
						}},
					},
				},
			},
		},
	})

	expected := `switch($var):case 'a':$a;case 'b':$b;endswitch;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltWhile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltWhile{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				}},
			},
		},
	})

	expected := `while($a):$b;endwhile;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtBreak(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Break{
		Expr: &scalar.Lnumber{
			Value: "1",
		},
	})

	expected := "break 1;"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCase(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Case{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := `case $a:$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCaseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Case{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmts: []node.Node{},
	})

	expected := "case $a:"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCatch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Catch{
		Types: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Exception"}}},
			&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "RuntimeException"}}},
		},
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "e"},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := `catch(Exception|\RuntimeException$e){$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ClassMethod{
		Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
		ReturnsRef: true,
		MethodName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        true,
				VariableType: &node.Nullable{Expr: &name.Name{Parts: []node.Node{&name.NamePart{Value: "int"}}}},
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				},
				DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
			},
			&node.Parameter{
				Variadic: true,
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
		ReturnType: &name.Name{
			Parts: []node.Node{&name.NamePart{Value: "void"}},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				}},
			},
		},
	})

	expected := `public function &foo(?int&$a=null,...$b):void{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAbstractClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ClassMethod{
		Modifiers: []node.Node{
			&node.Identifier{Value: "public"},
			&node.Identifier{Value: "static"},
		},
		ReturnsRef: true,
		MethodName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        true,
				VariableType: &node.Nullable{Expr: &name.Name{Parts: []node.Node{&name.NamePart{Value: "int"}}}},
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				},
				DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
			},
			&node.Parameter{
				Variadic: true,
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
		ReturnType: &name.Name{
			Parts: []node.Node{&name.NamePart{Value: "void"}},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `public static function &foo(?int&$a=null,...$b):void;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Class{
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
				Modifiers: []node.Node{
					&node.Identifier{Value: "public"},
					&node.Identifier{Value: "static"},
				},
				Consts: []node.Node{
					&stmt.Constant{
						ConstantName: &node.Identifier{Value: "FOO"},
						Expr:         &scalar.String{Value: "'bar'"},
					},
				},
			},
		},
	})

	expected := `abstract class Foo extends Bar implements Baz,Quuz{public static const FOO='bar';}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAnonymousClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Class{
		Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
		ArgumentList: &node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
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
	})

	expected := `abstract class($a,$b) extends Bar implements Baz,Quuz{public const FOO='bar';}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

	expected := `public const FOO='a',BAR='b';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ConstList{
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

	expected := `const FOO='a',BAR='b';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Constant{
		ConstantName: &node.Identifier{Value: "FOO"},
		Expr:         &scalar.String{Value: "'BAR'"},
	})

	expected := "FOO='BAR'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtContinue(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Continue{
		Expr: &scalar.Lnumber{
			Value: "1",
		},
	})

	expected := `continue 1;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Declare{
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
	})

	expected := `declare(FOO='bar'){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Declare{
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "'bar'"},
			},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := `declare(FOO='bar')'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Declare{
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "'bar'"},
			},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `declare(FOO='bar');`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalut(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Default{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := `default:$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalutEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Default{
		Stmts: []node.Node{},
	})

	expected := `default:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_Expression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Do{
		Cond: &scalar.Lnumber{Value: "1"},
		Stmt: &stmt.Expression{
			Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
		},
	})

	expected := `do $a;while(1);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_StmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Do{
		Cond: &scalar.Lnumber{Value: "1"},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				}},
			},
		},
	})

	expected := `do{$a;}while(1);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoHtmlState(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Root{
		Stmts: []node.Node{
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
					&expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `<?=$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoPhpState(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Echo{
		Exprs: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := `echo $a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ElseIf{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `elseif($a){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ElseIf{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := `elseif($a)'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ElseIf{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `elseif($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Else{
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `else{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Else{
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := `else 'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Else{
		Stmt: &stmt.Nop{},
	})

	expected := `else ;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Expression{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
	})

	expected := `$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Finally{
		Stmts: []node.Node{
			&stmt.Nop{},
		},
	})

	expected := `finally{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.For{
		Init: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
		Cond: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "c"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "d"},
			},
		},
		Loop: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "e"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "f"},
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `for($a,$b;$c,$d;$e,$f){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Foreach{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Key: &expr.Variable{
			VarName: &node.Identifier{Value: "k"},
		},
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "v"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `foreach($a as $k=>$v){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Function{
		ReturnsRef:   true,
		FunctionName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:    true,
				Variadic: false,
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
		ReturnType: &name.FullyQualified{
			Parts: []node.Node{&name.NamePart{Value: "Foo"}},
		},
		Stmts: []node.Node{
			&stmt.Nop{},
		},
	})

	expected := `function &foo(&$var):\Foo{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGlobal(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Global{
		Vars: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := `global$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGoto(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Goto{
		Label: &node.Identifier{Value: "FOO"},
	})

	expected := `goto FOO;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

	expected := `use function Foo\{Bar as Baz,Quuz};`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.HaltCompiler{})

	expected := `__halt_compiler();`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.If{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Expression{
			Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
		ElseIf: []node.Node{
			&stmt.ElseIf{
				Cond: &expr.Variable{
					VarName: &node.Identifier{Value: "c"},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{
							Expr: &expr.Variable{
								VarName: &node.Identifier{Value: "d"},
							},
						},
					},
				},
			},
			&stmt.ElseIf{
				Cond: &expr.Variable{
					VarName: &node.Identifier{Value: "e"},
				},
				Stmt: &stmt.Nop{},
			},
		},
		Else: &stmt.Else{
			Stmt: &stmt.Expression{
				Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "f"},
				},
			},
		},
	})

	expected := `if($a)$b;elseif($c){$d;}elseif($e);else $f;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.If{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{
					Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `if($a){$b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.If{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `if($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Root{
		Stmts: []node.Node{
			&stmt.InlineHtml{
				Value: "test",
			},
		},
	})

	expected := `test`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInterface(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Interface{
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
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
			},
		},
	})

	expected := `interface Foo extends Bar,Baz{public function foo(){$a;}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintLabel(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Label{
		LabelName: &node.Identifier{Value: "FOO"},
	})

	expected := `FOO:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespace(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Namespace{
		NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
	})

	expected := `namespace Foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespaceWithStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Namespace{
		NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := `namespace Foo{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Nop{})

	expected := `;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.PropertyList{
		Modifiers: []node.Node{
			&node.Identifier{Value: "public"},
			&node.Identifier{Value: "static"},
		},
		Properties: []node.Node{
			&stmt.Property{
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				},
				Expr: &scalar.String{Value: "'a'"},
			},
			&stmt.Property{
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
	})

	expected := `public static $a='a',$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintProperty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Property{
		Variable: &expr.Variable{
			FreeFloating: freefloating.Collection{
				freefloating.Start: []freefloating.String{
					{
						StringType: freefloating.TokenType,
						Value:      "$",
					},
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `$a=1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReturn(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Return{
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `return 1;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticVar(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.StaticVar{
		Variable: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `$a=1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStatic(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Static{
		Vars: []node.Node{
			&stmt.StaticVar{
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&stmt.StaticVar{
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
	})

	expected := `static$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			}},
		},
	})

	expected := `{$a;$b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtListNested(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
			&stmt.StmtList{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					}},
					&stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Expression{Expr: &expr.Variable{
								VarName: &node.Identifier{Value: "c"},
							}},
						},
					},
				},
			},
		},
	})

	expected := `{$a;{$b;{$c;}}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Switch{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
		CaseList: &stmt.CaseList{
			Cases: []node.Node{
				&stmt.Case{
					Cond: &scalar.String{Value: "'a'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
				&stmt.Case{
					Cond: &scalar.String{Value: "'b'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "b"},
						}},
					},
				},
			},
		},
	})

	expected := `switch($var){case 'a':$a;case 'b':$b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtThrow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Throw{
		Expr: &expr.Variable{
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `throw $var;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitAdaptationList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitAdaptationList{
		Adaptations: []node.Node{
			&stmt.TraitUseAlias{
				Ref: &stmt.TraitMethodRef{
					Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
					Method: &node.Identifier{Value: "a"},
				},
				Alias: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := `{Foo::a as b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitMethodRef(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitMethodRef{
		Method: &node.Identifier{Value: "a"},
	})

	expected := `a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitMethodRefFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintStmtTraitUseAlias(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintStmtTraitUsePrecedence(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

	expected := `Foo::a insteadof Bar,Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitUse{
		Traits: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		},
		TraitAdaptationList: &stmt.Nop{},
	})

	expected := `use Foo,Bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitAdaptations(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitUse{
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
	})

	expected := `use Foo,Bar{Foo::a as b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTrait(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Trait{
		TraitName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Stmts: []node.Node{
			&stmt.ClassMethod{
				Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
				MethodName: &node.Identifier{Value: "foo"},
				Params:     []node.Node{},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
			},
		},
	})

	expected := `trait Foo{public function foo(){$a;}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTry(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Try{
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			}},
		},
		Catches: []node.Node{
			&stmt.Catch{
				Types: []node.Node{
					&name.Name{Parts: []node.Node{&name.NamePart{Value: "Exception"}}},
					&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "RuntimeException"}}},
				},
				Variable: &expr.Variable{
					VarName: &node.Identifier{Value: "e"},
				},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{
						VarName: &node.Identifier{Value: "b"},
					}},
				},
			},
		},
		Finally: &stmt.Finally{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := `try{$a;}catch(Exception|\RuntimeException$e){$b;}finally{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Unset{
		Vars: []node.Node{
			&expr.Variable{
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := `unset($a,$b);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUseList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

	expected := `use function Foo as Bar,Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintWhileStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.While{
		Cond: &expr.Variable{
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "a"},
				}},
			},
		},
	})

	expected := `while($a){$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
