package printer_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/token"

	"github.com/z7zmey/php-parser/pkg/printer"
)

func TestPrinterPrintFile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				NamespaceName: &ast.NameName{
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
	p.Print(&ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtExpression{
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Tokens: token.Collection{
							token.Start: []*token.Token{
								{
									ID:    token.ID('$'),
									Value: []byte("$"),
								},
							},
						},
					},
					VarName: &ast.Identifier{
						Value: []byte("a"),
					},
				},
			},
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtExpression{
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Tokens: token.Collection{
							token.Start: []*token.Token{
								{
									ID:    token.ID('$'),
									Value: []byte("$"),
								},
							},
						},
					},
					VarName: &ast.Identifier{
						Value: []byte("a"),
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
	n := &ast.Identifier{
		Value: []byte("test"),
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
	p.Print(&ast.Parameter{
		Type: &ast.NameFullyQualified{
			Parts: []ast.Vertex{
				&ast.NameNamePart{
					Value: []byte("Foo"),
				},
			},
		},
		Var: &ast.Variadic{
			Var: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$var")},
			},
		},
		DefaultValue: &ast.ScalarString{
			Value: []byte("'default'"),
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
	p.Print(&ast.Nullable{
		Expr: &ast.Parameter{
			Type: &ast.NameFullyQualified{
				Parts: []ast.Vertex{
					&ast.NameNamePart{
						Value: []byte("Foo"),
					},
				},
			},
			Var: &ast.Reference{
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$var")},
				},
			},
			DefaultValue: &ast.ScalarString{
				Value: []byte("'default'"),
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
	p.Print(&ast.Argument{
		IsReference: false,
		Variadic:    true,
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.Argument{
		IsReference: true,
		Variadic:    false,
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.NameNamePart{
		Value: []byte("foo"),
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

func TestPrinterPrintNameFullyQualified(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintNameRelative(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintScalarLNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ScalarLnumber{
		Value: []byte("1"),
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
	p.Print(&ast.ScalarDnumber{
		Value: []byte(".1"),
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
	p.Print(&ast.ScalarString{
		Value: []byte("'hello world'"),
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
	p.Print(&ast.ScalarEncapsedStringPart{
		Value: []byte("hello world"),
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
	p.Print(&ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$var")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world")},
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
	p.Print(&ast.ScalarHeredoc{
		Label: []byte("<<<LBL\n"),
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$var")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world\n")},
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

func TestPrinterPrintScalarMagicConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ScalarMagicConstant{
		Value: []byte("__DIR__"),
	})

	if o.String() != `__DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrinterPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ExprAssign{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignReference{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignBitwiseAnd{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignBitwiseOr{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignBitwiseXor{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
		},
	})

	expected := `$a^=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ExprAssignCoalesce{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
		},
	})

	expected := `$a??=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ExprAssignConcat{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignDiv{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignMinus{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignMod{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignMul{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignPlus{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignPow{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignShiftLeft{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprAssignShiftRight{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryBitwiseAnd{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryBitwiseOr{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryBitwiseXor{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryBooleanAnd{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryBooleanOr{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryCoalesce{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryConcat{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryDiv{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryEqual{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryGreaterOrEqual{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryGreater{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryIdentical{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryLogicalAnd{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryLogicalOr{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryLogicalXor{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryMinus{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryMod{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryMul{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryNotEqual{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryNotIdentical{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryPlus{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryPow{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryShiftLeft{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinaryShiftRight{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinarySmallerOrEqual{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinarySmaller{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprBinarySpaceship{
		Left: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprCastArray{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprCastBool{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprCastDouble{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprCastInt{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprCastObject{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprCastString{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprCastUnset{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprArrayDimFetch{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		Dim: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.ExprArrayItem{
		Key: &ast.ScalarString{Value: []byte("'Hello'")},
		Val: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$world")},
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
	p.Print(&ast.ExprArrayItem{
		Val: &ast.ExprReference{Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$world")},
		}},
	})

	expected := `&$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemUnpack(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ExprArrayItem{
		Unpack: true,
		Val: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$world")},
		},
	})

	expected := `...$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ExprArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$world")},
				},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				Val: &ast.ExprReference{Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$var")},
				}},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprBitwiseNot{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprClassConstFetch{
		Class: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		ConstantName: &ast.Identifier{
			Value: []byte("CONST"),
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
	p.Print(&ast.ExprClone{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprClosureUse{
		Uses: []ast.Vertex{
			&ast.ExprReference{Var: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$foo")},
			}},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$bar")},
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
	p.Print(&ast.ExprClosure{
		Static:     true,
		ReturnsRef: true,
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.Reference{
					Var: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$var")},
					},
				},
			},
		},
		ClosureUse: &ast.ExprClosureUse{
			Uses: []ast.Vertex{
				&ast.ExprReference{Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
				}},
				&ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.NameFullyQualified{
			Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	})

	expected := `static function&(&$var)use(&$a,$b):\Foo{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrowFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtExpression{
		Expr: &ast.ExprArrowFunction{
			Static:     true,
			ReturnsRef: true,
			Params: []ast.Vertex{
				&ast.Parameter{
					Var: &ast.Reference{
						Var: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$var")},
						},
					},
				},
			},
			ReturnType: &ast.NameFullyQualified{
				Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}},
			},
			Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
		},
	})

	expected := `static fn&(&$var):\Foo=>$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.ExprConstFetch{
		Const: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("null")}}},
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
	p.Print(&ast.ExprEmpty{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprErrorSuppress{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprEval{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprExit{
		Die: false,
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprExit{
		Die: true,
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprFunctionCall{
		Function: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					IsReference: true,
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				&ast.Argument{
					Variadic: true,
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$c")},
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
	p.Print(&ast.ExprInclude{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
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
	p.Print(&ast.ExprIncludeOnce{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
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
	p.Print(&ast.ExprInstanceOf{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		Class: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
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
	p.Print(&ast.ExprIsset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								VarName: &ast.Identifier{Value: []byte("$b")},
							},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								VarName: &ast.Identifier{Value: []byte("$c")},
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
	p.Print(&ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$foo")},
		},
		Method: &ast.Identifier{Value: []byte("bar")},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprNew{
		Class: &ast.NameName{
			Parts: []ast.Vertex{
				&ast.NameNamePart{
					Value: []byte("Foo"),
				},
			},
		},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprPostDec{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprPostInc{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprPreDec{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprPreInc{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprPrint{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$foo")},
		},
		Property: &ast.Identifier{Value: []byte("bar")},
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
	p.Print(&ast.ExprReference{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$foo")},
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
	p.Print(&ast.ExprRequire{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
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
	p.Print(&ast.ExprRequireOnce{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
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
	p.Print(&ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$world")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte("!")},
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
	p.Print(&ast.ExprShortArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$world")},
				},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				Val: &ast.ExprReference{Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$var")},
				}},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprShortList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								VarName: &ast.Identifier{Value: []byte("$b")},
							},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								VarName: &ast.Identifier{Value: []byte("$c")},
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
	p.Print(&ast.ExprStaticCall{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Call:  &ast.Identifier{Value: []byte("bar")},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprStaticPropertyFetch{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Property: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$bar")},
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
	p.Print(&ast.ExprTernary{
		Condition: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		IfFalse: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.ExprTernary{
		Condition: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		IfTrue: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$b")},
		},
		IfFalse: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$c")},
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
	p.Print(&ast.ExprUnaryMinus{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprUnaryPlus{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprVariable{
		VarName: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprYieldFrom{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprYield{
		Value: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.ExprYield{
		Key: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$k")},
		},
		Value: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.StmtAltElseIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtAltElseIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{},
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
	p.Print(&ast.StmtAltElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtAltElse{
		Stmt: &ast.StmtStmtList{},
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
	p.Print(&ast.StmtAltFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$c")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$d")},
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
	p.Print(&ast.StmtAltForeach{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		Key: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$key")},
		},
		Var: &ast.ExprReference{Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$val")},
		}},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$d")},
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
	p.Print(&ast.StmtAltIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtAltElseIf{
				Cond: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$b")},
						}},
					},
				},
			},
			&ast.StmtAltElseIf{
				Cond: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$c")},
				},
				Stmt: &ast.StmtStmtList{},
			},
		},
		Else: &ast.StmtAltElse{
			Stmt: &ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtAltSwitch{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		CaseList: &ast.StmtCaseList{
			Cases: []ast.Vertex{
				&ast.StmtCase{
					Cond: &ast.ScalarString{Value: []byte("'a'")},
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$a")},
						}},
					},
				},
				&ast.StmtCase{
					Cond: &ast.ScalarString{Value: []byte("'b'")},
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtAltWhile{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtBreak{
		Expr: &ast.ScalarLnumber{
			Value: []byte("1"),
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
	p.Print(&ast.StmtCase{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtCase{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmts: []ast.Vertex{},
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
	p.Print(&ast.StmtCatch{
		Types: []ast.Vertex{
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Exception")}}},
			&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("RuntimeException")}}},
		},
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$e")},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtClassMethod{
		Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		ReturnsRef: true,
		MethodName: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: &ast.Nullable{Expr: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("int")}}}},
				Var: &ast.Reference{
					Var: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				Var: &ast.Variadic{
					Var: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
		ReturnType: &ast.NameName{
			Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("void")}},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtClassMethod{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		ReturnsRef: true,
		MethodName: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: &ast.Nullable{Expr: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("int")}}}},
				Var: &ast.Reference{
					Var: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				Var: &ast.Variadic{
					Var: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
		ReturnType: &ast.NameName{
			Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("void")}},
		},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtClass{
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
				Modifiers: []ast.Vertex{
					&ast.Identifier{Value: []byte("public")},
					&ast.Identifier{Value: []byte("static")},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						ConstantName: &ast.Identifier{Value: []byte("FOO")},
						Expr:         &ast.ScalarString{Value: []byte("'bar'")},
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
	p.Print(&ast.StmtClass{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
		ArgumentList: &ast.ArgumentList{
			Arguments: []ast.Vertex{
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
				},
				&ast.Argument{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					},
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
						ConstantName: &ast.Identifier{Value: []byte("FOO")},
						Expr:         &ast.ScalarString{Value: []byte("'bar'")},
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
	p.Print(&ast.StmtClassConstList{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("FOO")},
				Expr:         &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("BAR")},
				Expr:         &ast.ScalarString{Value: []byte("'b'")},
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
	p.Print(&ast.StmtConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("FOO")},
				Expr:         &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("BAR")},
				Expr:         &ast.ScalarString{Value: []byte("'b'")},
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
	p.Print(&ast.StmtConstant{
		ConstantName: &ast.Identifier{Value: []byte("FOO")},
		Expr:         &ast.ScalarString{Value: []byte("'BAR'")},
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
	p.Print(&ast.StmtContinue{
		Expr: &ast.ScalarLnumber{
			Value: []byte("1"),
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
	p.Print(&ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("FOO")},
				Expr:         &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("FOO")},
				Expr:         &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				ConstantName: &ast.Identifier{Value: []byte("FOO")},
				Expr:         &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtDefault{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtDefault{
		Stmts: []ast.Vertex{},
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
	p.Print(&ast.StmtDo{
		Cond: &ast.ScalarLnumber{Value: []byte("1")},
		Stmt: &ast.StmtExpression{
			Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtDo{
		Cond: &ast.ScalarLnumber{Value: []byte("1")},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$a")},
					},
					&ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
	})

	expected := `<?php echo $a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoPhpState(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtEcho{
		Exprs: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtElse{
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
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
	p.Print(&ast.StmtElse{
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.StmtExpression{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtFinally{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
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
	p.Print(&ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$c")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$d")},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$e")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$f")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtForeach{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Key: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$k")},
		},
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$v")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtFunction{
		ReturnsRef:   true,
		FunctionName: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.Reference{
					Var: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$var")},
					},
				},
			},
		},
		ReturnType: &ast.NameFullyQualified{
			Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
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
	p.Print(&ast.StmtGlobal{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtGoto{
		Label: &ast.Identifier{Value: []byte("FOO")},
	})

	expected := `goto FOO;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtHaltCompiler{})

	expected := `__halt_compiler();`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtExpression{
			Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$c")},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{
							Expr: &ast.ExprVariable{
								VarName: &ast.Identifier{Value: []byte("$d")},
							},
						},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$e")},
				},
				Stmt: &ast.StmtNop{},
			},
		},
		Else: &ast.StmtElse{
			Stmt: &ast.StmtExpression{
				Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$f")},
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
	p.Print(&ast.StmtIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{
					Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtIf{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtNop{},
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
	p.Print(&ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{
				Value: []byte("test"),
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
	p.Print(&ast.StmtInterface{
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
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtLabel{
		LabelName: &ast.Identifier{Value: []byte("FOO")},
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
	p.Print(&ast.StmtNamespace{
		NamespaceName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
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
	p.Print(&ast.StmtNamespace{
		NamespaceName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtNop{})

	expected := `;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtPropertyList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		Type: &ast.NameName{
			Parts: []ast.Vertex{
				&ast.NameNamePart{
					Value: []byte("Foo"),
				},
			},
		},
		Properties: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
				},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtProperty{
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	})

	expected := `public static Foo $a='a',$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintProperty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtProperty{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtReturn{
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
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
	p.Print(&ast.StmtStatic{
		Vars: []ast.Vertex{
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			}},
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			}},
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					}},
					&ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{
								VarName: &ast.Identifier{Value: []byte("$c")},
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
	p.Print(&ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
		},
		CaseList: &ast.StmtCaseList{
			Cases: []ast.Vertex{
				&ast.StmtCase{
					Cond: &ast.ScalarString{Value: []byte("'a'")},
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$a")},
						}},
					},
				},
				&ast.StmtCase{
					Cond: &ast.ScalarString{Value: []byte("'b'")},
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$b")},
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
	p.Print(&ast.StmtThrow{
		Expr: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$var")},
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
	p.Print(&ast.StmtTraitAdaptationList{
		Adaptations: []ast.Vertex{
			&ast.StmtTraitUseAlias{
				Ref: &ast.StmtTraitMethodRef{
					Trait:  &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
					Method: &ast.Identifier{Value: []byte("a")},
				},
				Alias: &ast.Identifier{Value: []byte("b")},
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
	p.Print(&ast.StmtTraitMethodRef{
		Method: &ast.Identifier{Value: []byte("a")},
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

func TestPrinterPrintStmtTraitUseAlias(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

func TestPrinterPrintStmtTraitUsePrecedence(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
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

	expected := `Foo::a insteadof Bar,Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
			&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Bar")}}},
		},
		TraitAdaptationList: &ast.StmtNop{},
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
	p.Print(&ast.StmtTraitUse{
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
	p.Print(&ast.StmtTrait{
		TraitName: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		Stmts: []ast.Vertex{
			&ast.StmtClassMethod{
				Modifiers:  []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				MethodName: &ast.Identifier{Value: []byte("foo")},
				Params:     []ast.Vertex{},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							VarName: &ast.Identifier{Value: []byte("$a")},
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
	p.Print(&ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			}},
		},
		Catches: []ast.Vertex{
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Exception")}}},
					&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("RuntimeException")}}},
				},
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$e")},
				},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						VarName: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
		Finally: &ast.StmtFinally{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
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
	p.Print(&ast.StmtUnset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				VarName: &ast.Identifier{Value: []byte("$b")},
			},
		},
	})

	expected := `unset($a,$b);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtUse{
		UseList: &ast.StmtUseList{
			UseDeclarations: []ast.Vertex{
				&ast.StmtUseDeclaration{
					Use: &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
				},
			},
		},
	})

	expected := `use Foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGroupUseList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtGroupUseList{
		Prefix:  &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		UseList: &ast.StmtUseList{},
	})

	expected := `Foo\{}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUseList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtUseList{
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

	expected := `Foo as Bar,Baz`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUseDeclaration(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtUseDeclaration{
		Use:   &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
		Alias: &ast.Identifier{Value: []byte("Bar")},
	})

	expected := `Foo as Bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUseType(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&ast.StmtUseType{
		Type: &ast.Identifier{Value: []byte("function")},
		Use: &ast.StmtUseDeclaration{
			Use:   &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}},
			Alias: &ast.Identifier{Value: []byte("Bar")},
		},
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
	p.Print(&ast.StmtWhile{
		Cond: &ast.ExprVariable{
			VarName: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("$a")},
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
