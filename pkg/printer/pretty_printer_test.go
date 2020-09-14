package printer_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/printer"
)

func TestPrintFile(t *testing.T) {
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
	rootNode := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.NameName{
					Parts: []ast.Vertex{
						&ast.NameNamePart{Value: []byte("Foo")},
					},
				},
			},
			&ast.StmtClass{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
				ClassName: &ast.NameName{
					Parts: []ast.Vertex{
						&ast.NameNamePart{Value: []byte("Bar")},
					},
				},
				Extends: &ast.StmtClassExtends{
					ClassName: &ast.NameName{
						Parts: []ast.Vertex{
							&ast.NameNamePart{Value: []byte("Baz")},
						},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						MethodName: &ast.Identifier{Value: []byte("greet")},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtEcho{
									Exprs: []ast.Vertex{
										&ast.ScalarString{Value: []byte("'Hello world'")},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	o1 := bytes.NewBufferString("")
	p1 := printer.NewPrettyPrinter(o1, "\t")
	p1.Print(rootNode)
	if actual := o1.String(); expected != actual {
		t.Errorf("\nPrint the 1st time\nexpected: %s\ngot: %s\n", expected, actual)
	}

	o2 := bytes.NewBufferString("")
	p2 := printer.NewPrettyPrinter(o2, "\t")
	p2.Print(rootNode)
	if actual := o2.String(); expected != actual {
		t.Errorf("\nPrint the 2nd time\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintFileInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtExpression{
				Expr: &ast.ScalarHeredoc{
					Label: []byte("<<<\"LBL\"\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{Value: []byte("hello world\n")},
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
	p.Print(&ast.Identifier{Value: []byte("test")})

	if o.String() != `test` {
		t.Errorf("TestPrintIdentifier is failed\n")
	}
}

func TestPrintParameter(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.Parameter{
		Type: &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		Var: &ast.Variadic{
			Var: &ast.ExprVariable{
				VarName: &ast.Identifier{
					Value: []byte("var"),
				},
			},
		},
		DefaultValue: &ast.ScalarString{Value: []byte("'default'")},
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
	p.Print(&ast.Nullable{
		Expr: &ast.Parameter{
			Type:         &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
			Var:          &ast.Variadic{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}},
			DefaultValue: &ast.ScalarString{Value: []byte("'default'")},
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
	p.Print(&ast.Argument{
		IsReference: false,
		Variadic:    true,
		Expr:        &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.Argument{
		IsReference: true,
		Variadic:    false,
		Expr:        &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.NameNamePart{
		Value: []byte("foo"),
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
	p.Print(&ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{
				Value: []byte("Foo"),
			},
			&ast.NameNamePart{
				Value: []byte("Bar"),
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
	p.Print(&ast.NameFullyQualified{
		Parts: []ast.Vertex{
			&ast.NameNamePart{
				Value: []byte("Foo"),
			},
			&ast.NameNamePart{
				Value: []byte("Bar"),
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
	p.Print(&ast.NameRelative{
		Parts: []ast.Vertex{
			&ast.NameNamePart{
				Value: []byte("Foo"),
			},
			&ast.NameNamePart{
				Value: []byte("Bar"),
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
	p.Print(&ast.ScalarLnumber{Value: []byte("1")})

	if o.String() != `1` {
		t.Errorf("TestPrintScalarLNumber is failed\n")
	}
}

func TestPrintScalarDNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ScalarDnumber{Value: []byte(".1")})

	if o.String() != `.1` {
		t.Errorf("TestPrintScalarDNumber is failed\n")
	}
}

func TestPrintScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ScalarString{Value: []byte("'hello world'")})

	expected := `'hello world'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ScalarEncapsedStringPart{Value: []byte("hello world")})

	if o.String() != `hello world` {
		t.Errorf("TestPrintScalarEncapsedStringPart is failed\n")
	}
}

func TestPrintScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world")},
		},
	})

	if o.String() != `"hello {$var} world"` {
		t.Errorf("TestPrintScalarEncapsed is failed\n")
	}
}

func TestPrintScalarHeredoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ScalarHeredoc{
		Label: []byte("<<<LBL\n"),
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world\n")},
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
	p.Print(&ast.ScalarHeredoc{
		Label: []byte("<<<'LBL'\n"),
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello world\n")},
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
	p.Print(&ast.ScalarMagicConstant{Value: []byte("__DIR__")})

	if o.String() != `__DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprAssign{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignReference{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignBitwiseAnd{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignBitwiseOr{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignBitwiseXor{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignConcat{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignDiv{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignMinus{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignMod{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignMul{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignPlus{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignPow{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignShiftLeft{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprAssignShiftRight{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryBitwiseAnd{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryBitwiseOr{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryBitwiseXor{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryBooleanAnd{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryBooleanOr{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryCoalesce{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryConcat{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryDiv{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryEqual{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryGreaterOrEqual{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryGreater{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryIdentical{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryLogicalAnd{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryLogicalOr{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryLogicalXor{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryMinus{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryMod{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryMul{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryNotEqual{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryNotIdentical{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryPlus{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryPow{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryShiftLeft{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinaryShiftRight{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinarySmallerOrEqual{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinarySmaller{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprBinarySpaceship{
		Left:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Right: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprCastArray{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprCastBool{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprCastDouble{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprCastInt{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprCastObject{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprCastString{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprCastUnset{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprArrayDimFetch{
		Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
		Dim: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.ExprArrayItem{
		Key: &ast.ScalarString{Value: []byte("'Hello'")},
		Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("world")}},
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
	p.Print(&ast.ExprArrayItem{
		Val: &ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("world")}}},
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
	p.Print(&ast.ExprArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("world")}},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				Val: &ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprBitwiseNot{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprClassConstFetch{
		Class:        &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
		ConstantName: &ast.Identifier{Value: []byte("CONST")},
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
	p.Print(&ast.ExprClone{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprClosureUse{
		Uses: []ast.Vertex{
			&ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}}},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("bar")}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.ExprClosure{
				Static:     true,
				ReturnsRef: true,
				Params: []ast.Vertex{
					&ast.Parameter{
						Var: &ast.Reference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}},
					},
				},
				ClosureUse: &ast.ExprClosureUse{
					Uses: []ast.Vertex{
						&ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
						&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
					},
				},
				ReturnType: &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.ExprConstFetch{
		Const: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("null")}}},
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
	p.Print(&ast.ExprEmpty{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `empty($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrettyPrinterrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprErrorSuppress{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `@$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintEval(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprEval{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `eval($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintExit(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprExit{Die: false, Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `exit($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintDie(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprExit{Die: true, Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `die($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintFunctionCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprFunctionCall{
		Function: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					IsReference: true,
					Expr:        &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				},
				&ast.Argument{
					Variadic: true,
					Expr:     &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
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
	p.Print(&ast.ExprInclude{Expr: &ast.ScalarString{Value: []byte("'path'")}})

	expected := `include 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprIncludeOnce{Expr: &ast.ScalarString{Value: []byte("'path'")}})

	expected := `include_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprInstanceOf{
		Expr:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
		Class: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
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
	p.Print(&ast.ExprIsset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
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
	p.Print(&ast.ExprMethodCall{
		Var:    &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
		Method: &ast.Identifier{Value: []byte("bar")},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprNew{
		Class: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprPostDec{
		Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprPostInc{
		Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprPreDec{
		Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprPreInc{
		Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprPrint{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `print($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprPropertyFetch{
		Var:      &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
		Property: &ast.Identifier{Value: []byte("bar")},
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
	p.Print(&ast.ExprReference{
		Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
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
	p.Print(&ast.ExprRequire{Expr: &ast.ScalarString{Value: []byte("'path'")}})

	expected := `require 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprRequireOnce{Expr: &ast.ScalarString{Value: []byte("'path'")}})

	expected := `require_once 'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("world")}},
			&ast.ScalarEncapsedStringPart{Value: []byte("!")},
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
	p.Print(&ast.ExprShortArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("world")}},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				Val: &ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprShortList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
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
	p.Print(&ast.ExprStaticCall{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Call:  &ast.Identifier{Value: []byte("bar")},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprStaticPropertyFetch{
		Class:    &ast.Identifier{Value: []byte("Foo")},
		Property: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("bar")}},
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
	p.Print(&ast.ExprTernary{
		Condition: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		IfFalse:   &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.ExprTernary{
		Condition: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		IfTrue:    &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
		IfFalse:   &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
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
	p.Print(&ast.ExprUnaryMinus{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprUnaryPlus{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprVariable{VarName: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}})

	expected := `$$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.ExprYieldFrom{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprYield{
		Value: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.ExprYield{
		Key:   &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("k")}},
		Value: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.StmtElseIf{
		Alt:  true,
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtElseIf{
		Alt:  true,
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtStmtList{},
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
	p.Print(&ast.StmtElse{
		Alt: true,
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtElse{
		Alt:  true,
		Stmt: &ast.StmtStmtList{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtFor{
				Alt: true,
				Init: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				},
				Cond: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
				},
				Loop: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("d")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Alt:  true,
				Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
				Key:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("key")}},
				Var:  &ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("val")}}},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("d")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Alt:  true,
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("d")}}},
					},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Alt:  true,
						Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
							},
						},
					},
					&ast.StmtElseIf{
						Alt:  true,
						Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
						Stmt: &ast.StmtStmtList{},
					},
				},
				Else: &ast.StmtElse{
					Alt: true,
					Stmt: &ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtSwitch{
				Alt:  true,
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
				CaseList: []ast.Vertex{
					&ast.StmtCase{
						Cond: &ast.ScalarString{Value: []byte("'a'")},
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
						},
					},
					&ast.StmtCase{
						Cond: &ast.ScalarString{Value: []byte("'b'")},
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Alt:  true,
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtBreak{
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtCase{
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtCase{
		Cond:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmts: []ast.Vertex{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Exception")}}},
					&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("RuntimeException")}}},
				},
				Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("e")}},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtClassMethod{
		Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		ReturnsRef: true,
		MethodName: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type:         &ast.Nullable{Expr: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("int")}}}},
				Var:          &ast.Reference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				Var: &ast.Variadic{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
			},
		},
		ReturnType: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("void")}}},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtClassMethod{
		Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		ReturnsRef: true,
		MethodName: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type:         &ast.Nullable{Expr: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("int")}}}},
				Var:          &ast.Reference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				Var: &ast.Variadic{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
			},
		},
		ReturnType: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("void")}}},
		Stmt:       &ast.StmtNop{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
				ClassName: &ast.Identifier{Value: []byte("Foo")},
				Extends: &ast.StmtClassExtends{
					ClassName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
				},
				Implements: &ast.StmtClassImplements{
					InterfaceNames: []ast.Vertex{
						&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Baz")}}},
						&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Quuz")}}},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Name: &ast.Identifier{Value: []byte("FOO")},
								Expr: &ast.ScalarString{Value: []byte("'bar'")},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
				ArgumentList: &ast.ArgumentList{
					Arguments: []ast.Vertex{
						&ast.Argument{
							Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
						},
						&ast.Argument{
							Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
						},
					},
				},
				Extends: &ast.StmtClassExtends{
					ClassName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
				},
				Implements: &ast.StmtClassImplements{
					InterfaceNames: []ast.Vertex{
						&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Baz")}}},
						&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Quuz")}}},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Name: &ast.Identifier{Value: []byte("FOO")},
								Expr: &ast.ScalarString{Value: []byte("'bar'")},
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
	p.Print(&ast.StmtClassConstList{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("BAR")},
				Expr: &ast.ScalarString{Value: []byte("'b'")},
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
	p.Print(&ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("FOO")},
		Expr: &ast.ScalarString{Value: []byte("'BAR'")},
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
	p.Print(&ast.StmtContinue{
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtDeclare{
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Name: &ast.Identifier{Value: []byte("FOO")},
						Expr: &ast.ScalarString{Value: []byte("'bar'")},
					},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtDeclare{
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Name: &ast.Identifier{Value: []byte("FOO")},
						Expr: &ast.ScalarString{Value: []byte("'bar'")},
					},
				},
				Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtDefault{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtDefault{
		Stmts: []ast.Vertex{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtDo{
				Cond: &ast.ScalarLnumber{Value: []byte("1")},
				Stmt: &ast.StmtExpression{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtDo{
				Cond: &ast.ScalarLnumber{Value: []byte("1")},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtEcho{
		Exprs: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.StmtElseIf{
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtElseIf{
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtElseIf{
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtElse{
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtElse{
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}})

	expected := `$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtFinally{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtFor{
				Init: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
				},
				Cond: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("d")}},
				},
				Loop: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("e")}},
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("f")}},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtFor{
				Init: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				},
				Cond: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
				},
				Loop: []ast.Vertex{
					&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
				},
				Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
		},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Key:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("k")}},
				Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("v")}},
				Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtForeach{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Key:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("k")}},
		Var:  &ast.ExprReference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("v")}}},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				ReturnsRef:   true,
				FunctionName: &ast.Identifier{Value: []byte("foo")},
				Params: []ast.Vertex{
					&ast.Parameter{
						Var: &ast.Reference{Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}}},
					},
				},
				ReturnType: &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
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
	p.Print(&ast.StmtGlobal{
		Vars: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.StmtGoto{
		Label: &ast.Identifier{Value: []byte("FOO")},
	})

	expected := `goto FOO;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtHaltCompiler{})

	expected := `__halt_compiler();`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Stmt: &ast.StmtExpression{
					Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtExpression{
									Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("d")}},
								},
							},
						},
					},
					&ast.StmtElseIf{
						Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("e")}},
						Stmt: &ast.StmtNop{},
					},
				},
				Else: &ast.StmtElse{
					Stmt: &ast.StmtExpression{
						Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("f")}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{
							Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.StmtIf{
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtInlineHtml{
		Value: []byte("test"),
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtInterface{
				InterfaceName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Extends: &ast.StmtInterfaceExtends{
					InterfaceNames: []ast.Vertex{
						&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
						&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Baz")}}},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						MethodName: &ast.Identifier{Value: []byte("foo")},
						Params:     []ast.Vertex{},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtLabel{
		LabelName: &ast.Identifier{Value: []byte("FOO")},
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
	p.Print(&ast.StmtNamespace{
		Name: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtNop{})

	expected := `;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtPropertyList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		Properties: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			},
			&ast.StmtProperty{
				Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.StmtProperty{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtReturn{
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtStaticVar{
		Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtStatic{
		Vars: []ast.Vertex{
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			},
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
			&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
					&ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("c")}}},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtSwitch{
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
				CaseList: []ast.Vertex{
					&ast.StmtCase{
						Cond: &ast.ScalarString{Value: []byte("'a'")},
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
						},
					},
					&ast.StmtCase{
						Cond: &ast.ScalarString{Value: []byte("'b'")},
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
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
	p.Print(&ast.StmtThrow{
		Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("var")}},
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
	p.Print(&ast.StmtTraitMethodRef{
		Trait:  &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		Method: &ast.Identifier{Value: []byte("a")},
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
	p.Print(&ast.StmtTraitUseAlias{
		Ref: &ast.StmtTraitMethodRef{
			Trait:  &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
			Method: &ast.Identifier{Value: []byte("a")},
		},
		Modifier: &ast.Identifier{Value: []byte("public")},
		Alias:    &ast.Identifier{Value: []byte("b")},
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
	p.Print(&ast.StmtTraitUsePrecedence{
		Ref: &ast.StmtTraitMethodRef{
			Trait:  &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
			Method: &ast.Identifier{Value: []byte("a")},
		},
		Insteadof: []ast.Vertex{
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Baz")}}},
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
	p.Print(&ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtTraitUse{
				Traits: []ast.Vertex{
					&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
					&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
				},
				TraitAdaptationList: &ast.StmtTraitAdaptationList{
					Adaptations: []ast.Vertex{
						&ast.StmtTraitUseAlias{
							Ref: &ast.StmtTraitMethodRef{
								Trait:  &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
								Method: &ast.Identifier{Value: []byte("a")},
							},
							Alias: &ast.Identifier{Value: []byte("b")},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtTrait{
				TraitName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						MethodName: &ast.Identifier{Value: []byte("foo")},
						Params:     []ast.Vertex{},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
				},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Types: []ast.Vertex{
							&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Exception")}}},
							&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("RuntimeException")}}},
						},
						Var: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("e")}},
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}}},
						},
					},
				},
				Finally: &ast.StmtFinally{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
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
	p.Print(&ast.StmtUnset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
			&ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("b")}},
		},
	})

	expected := `unset($a, $b);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtUse{
		Type: &ast.Identifier{Value: []byte("function")},
		UseDeclarations: []ast.Vertex{
			&ast.StmtUseDeclaration{
				Use:   &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Alias: &ast.Identifier{Value: []byte("Bar")},
			},
			&ast.StmtUseDeclaration{
				Use: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Baz")}}},
			},
		},
	})

	expected := `use function Foo as Bar, Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintStmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtGroupUse{
		Type:   &ast.Identifier{Value: []byte("function")},
		Prefix: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		UseDeclarations: []ast.Vertex{
			&ast.StmtUseDeclaration{
				Use:   &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				Alias: &ast.Identifier{Value: []byte("Bar")},
			},
			&ast.StmtUseDeclaration{
				Use: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Baz")}}},
			},
		},
	})

	expected := `use function Foo\{Foo as Bar, Baz}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrintUseDeclaration(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrettyPrinter(o, "    ")
	p.Print(&ast.StmtUseDeclaration{
		Type:  &ast.Identifier{Value: []byte("function")},
		Use:   &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		Alias: &ast.Identifier{Value: []byte("Bar")},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
				Stmt: &ast.StmtExpression{Expr: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}}},
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
	p.Print(&ast.StmtWhile{
		Cond: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("a")}},
		Stmt: &ast.StmtNop{},
	})

	expected := `while ($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
