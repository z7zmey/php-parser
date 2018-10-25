package printer_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/meta"

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

	expected := `namespaceFooabstractclassBarextendsBaz{publicfunctiongreet(){'Hello world'}}`
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
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "<?php ",
						TokenName: meta.NodeStart,
					},
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "",
						TokenName: meta.SemiColonToken,
					},
				},
				Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{
						Value: "a",
					},
				},
			},
		},
	})

	expected := `<div>HTML</div><?php $a`
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     "  ",
				TokenName: meta.NodeStart,
			},
		},
		Value: "test",
	}
	p.Print(n)

	if o.String() != `  test` {
		t.Errorf("TestPrintIdentifier is failed\n")
	}
}

func TestPrinterPrintParameter(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Parameter{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EllipsisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EqualToken,
			},
		},
		ByRef:    false,
		Variadic: true,
		VariableType: &name.FullyQualified{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NsSeparatorToken,
				},
			},
			Parts: []node.Node{
				&name.NamePart{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.WhiteSpaceType,
							Value:     " ",
							TokenName: meta.StringToken,
						},
					},
					Value: "Foo",
				},
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		DefaultValue: &scalar.String{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ConstantEncapsedStringToken,
				},
			},
			Value: "'default'",
		},
	})

	expected := " \\ Foo ...$var = 'default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNullable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Nullable{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.QuestionMarkToken,
			},
		},
		Expr: &node.Parameter{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.AmpersandToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.EqualToken,
				},
			},
			ByRef:    true,
			Variadic: false,
			VariableType: &name.FullyQualified{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.WhiteSpaceType,
						Value:     " ",
						TokenName: meta.NsSeparatorToken,
					},
				},
				Parts: []node.Node{
					&name.NamePart{
						Value: "Foo",
					},
				},
			},
			Variable: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.WhiteSpaceType,
						Value:     " ",
						TokenName: meta.NodeStart,
					},
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{
					Value: "var",
				},
			},
			DefaultValue: &scalar.String{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.WhiteSpaceType,
						Value:     " ",
						TokenName: meta.ConstantEncapsedStringToken,
					},
				},
				Value: "'default'",
			},
		},
	})

	expected := " ? \\Foo & $var = 'default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintArgument(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Argument{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EllipsisToken,
			},
		},
		IsReference: false,
		Variadic:    true,
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{
				Value: "var",
			},
		},
	})

	expected := " ... $var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
func TestPrinterPrintArgumentByRef(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&node.Argument{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
		},
		IsReference: true,
		Variadic:    false,
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{
				Value: "var",
			},
		},
	})

	expected := " & $var"
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StringToken,
			},
			&meta.Data{
				Type:      meta.CommentType,
				Value:     "/*comment*/",
				TokenName: meta.StringToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StringToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NsSeparatorToken,
			},
		},
		Value: "foo",
	})

	expected := " /*comment*/ foo "
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameName(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&name.Name{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StringToken,
			},
		},
		Parts: []node.Node{
			&name.NamePart{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.CommentType,
						Value:     "/*comment*/",
						TokenName: meta.StringToken,
					},
					&meta.Data{
						Type:      meta.WhiteSpaceType,
						Value:     " ",
						TokenName: meta.StringToken,
					},
					&meta.Data{
						Type:      meta.WhiteSpaceType,
						Value:     " ",
						TokenName: meta.NsSeparatorToken,
					},
				},
				Value: "Foo",
			},
			&name.NamePart{
				Value: "Bar",
			},
		},
	})

	expected := " /*comment*/ Foo \\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameFullyQualified(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&name.FullyQualified{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NsSeparatorToken,
			},
		},
		Parts: []node.Node{
			&name.NamePart{
				Value: "Foo",
			},
			&name.NamePart{
				Value: "Bar",
			},
		},
	})

	expected := " \\Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameRelative(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&name.Relative{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NamespaceToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NsSeparatorToken,
			},
		},
		Parts: []node.Node{
			&name.NamePart{
				Value: "Foo",
			},
			&name.NamePart{
				Value: "Bar",
			},
		},
	})

	expected := " namespace \\Foo\\Bar"
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.LnumberToken,
			},
		},
		Value: "1",
	})

	expected := " 1"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarDNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Dnumber{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DnumberToken,
			},
		},
		Value: ".1",
	})

	expected := " .1"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.String{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ConstantEncapsedStringToken,
			},
		},
		Value: "'hello world'",
	})

	expected := ` 'hello world'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.EncapsedStringPart{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EncapsedAndWhitespaceToken,
			},
		},
		Value: "hello world",
	})

	if o.String() != ` hello world` {
		t.Errorf("TestPrintScalarEncapsedStringPart is failed\n")
	}
}

func TestPrinterPrintScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Encapsed{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoubleQuoteToken,
			},
		},
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "var"},
			},
			&scalar.EncapsedStringPart{Value: " world"},
		},
	})

	if o.String() != ` "hello $var world"` {
		t.Errorf("TestPrintScalarEncapsed is failed\n")
	}
}

func TestPrinterPrintScalarHeredoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&scalar.Heredoc{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StartHeredocToken,
			},
		},
		Label: "LBL",
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "var"},
			},
			&scalar.EncapsedStringPart{Value: " world"},
		},
	})

	expected := ` <<<LBL
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StartHeredocToken,
			},
		},
		Label: "'LBL'",
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello world"},
		},
	})

	expected := ` <<<'LBL'
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.MagicConstantToken,
			},
		},
		Value: "__DIR__",
	})

	if o.String() != ` __DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrinterPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Assign{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a = $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Reference{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EqualToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a = & $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.BitwiseAnd{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AndEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a &= $b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.BitwiseOr{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OrEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a |=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.BitwiseXor{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.XorEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ^=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Concat{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ConcatEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a .=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Div{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DivEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a /=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Minus{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.MinusEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a -=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Mod{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ModEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a %=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMul(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Mul{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.MulEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a *=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Plus{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PlusEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a +=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.Pow{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PowEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a **=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.ShiftLeft{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SlEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a <<=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&assign.ShiftRight{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SrEqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expression: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a >>=$b`
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a &$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BitwiseOr{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.VerticalBarToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a |$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BitwiseXor{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CaretToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ^$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BooleanAnd{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.BooleanAndToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a &&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.BooleanOr{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.BooleanOrToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ||$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Coalesce{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CoalesceToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ??$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Concat{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DotToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a .$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Div{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SlashToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a /$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Equal{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IsEqualToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ==$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreaterOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.GreaterOrEqual{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IsGreaterOrEqualToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a >=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreater(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Greater{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.GreaterToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a >$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Identical{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IsIdenticalToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ===$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.LogicalAnd{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.LogicalAndToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a and$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.LogicalOr{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.LogicalOrToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a or$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.LogicalXor{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.LogicalXorToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a xor$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Minus{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.MinusToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a -$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Mod{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PercentToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a %$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMul(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Mul{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AsteriskToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a *$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.NotEqual{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IsNotEqualToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "!=",
				TokenName: meta.IsNotEqualToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a !=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.NotIdentical{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IsNotIdenticalToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a !==$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Plus{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PlusToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a +$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Pow{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PowToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a **$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.ShiftLeft{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SlToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a <<$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.ShiftRight{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SrToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a >>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmallerOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.SmallerOrEqual{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IsSmallerOrEqualToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a <=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmaller(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Smaller{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.LessToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a <$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySpaceship(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&binary.Spaceship{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SpaceshipToken,
			},
		},
		Left: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Right: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a <=>$b`
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ArrayCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(array)",
				TokenName: meta.ArrayCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (array)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBool(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Bool{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.BoolCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(bool)",
				TokenName: meta.BoolCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (bool)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDouble(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Double{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoubleCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(float)",
				TokenName: meta.DoubleCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (float)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInt(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Int{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IntCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(int)",
				TokenName: meta.IntCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (int)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintObject(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Object{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ObjectCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(object)",
				TokenName: meta.ObjectCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (object)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.String{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StringCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(string)",
				TokenName: meta.StringCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (string)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&cast.Unset{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UnsetCastToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "(unset)",
				TokenName: meta.UnsetCastToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` (unset)$var`
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenSquareBracket,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "[",
				TokenName: meta.OpenSquareBracket,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseSquareBracket,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "]",
				TokenName: meta.CloseSquareBracket,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		Dim: &scalar.Lnumber{Value: "1"},
	})

	expected := `$var [1 ]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemWithKey(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ArrayItem{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoubleArrowToken,
			},
		},
		Key: &scalar.String{Value: "'Hello'"},
		Val: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "world"},
		},
	})

	expected := `'Hello' => $world`
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
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ArrayToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "'Hello'"},
				Val: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "world"},
				},
			},
			&expr.ArrayItem{
				Key: &scalar.Lnumber{Value: "2"},
				Val: &expr.Reference{Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "var"},
				}},
			},
			&expr.ArrayItem{
				Val: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
	})

	expected := ` array ('Hello'=>$world,2=>&$var,$var )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBitwiseNot(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.BitwiseNot{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.TildeToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` ~$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBooleanNot(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.BooleanNot{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ExclamationMarkToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` !$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClassConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ClassConstFetch{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PaamayimNekudotayimToken,
			},
		},
		Class: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		ConstantName: &node.Identifier{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
			},
			Value: "CONST",
		},
	})

	expected := `$var :: CONST`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClone(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Clone{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloneToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` clone$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosureUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ClosureUse{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UseToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Uses: []node.Node{
			&expr.Reference{Variable: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "foo"},
			}},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "bar"},
			},
		},
	})

	expected := ` use (&$foo,$bar )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosure(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Closure{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StaticToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.FunctionToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Static:     true,
		ReturnsRef: true,
		Params: []node.Node{
			&node.Parameter{
				ByRef:    true,
				Variadic: false,
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
		ClosureUse: &expr.ClosureUse{
			Uses: []node.Node{
				&expr.Reference{Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				}},
				&expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
		ReturnType: &name.FullyQualified{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ColonToken,
				},
			},
			Parts: []node.Node{&name.NamePart{Value: "Foo"}},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := ` static function & (&$var )use(&$a,$b) :\Foo {$a }`
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EmptyToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` empty ($var )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrettyPrinterrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ErrorSuppress{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AtToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` @$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEval(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Eval{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EvalToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` eval ($var )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExit(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Exit{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ExitToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Die: false,
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` exit $var `
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ExitToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` die $var `
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
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		ArgumentList: &node.ArgumentList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenParenthesisToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseParenthesisToken,
				},
			},
			Arguments: []node.Node{
				&node.Argument{
					IsReference: true,
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Variadic: true,
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "c"},
					},
				},
			},
		},
	})

	expected := `$var (&$a,...$b,$c )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInclude(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Include{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IncludeToken,
			},
		},
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := ` include'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.IncludeOnce{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IncludeOnceToken,
			},
		}, Expr: &scalar.String{Value: "'path'"},
	})

	expected := ` include_once'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.InstanceOf{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.InstanceofToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
	})

	expected := `$var instanceofFoo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIsset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Isset{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IssetToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Variables: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := ` isset ($a,$b )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.List{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ListToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Items: []node.Node{
			&expr.ArrayItem{
				Val: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&expr.ArrayItem{
				Val: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{
								Meta: meta.Collection{
									&meta.Data{
										Type:      meta.TokenType,
										Value:     "$",
										TokenName: meta.NodeStart,
									},
								},
								VarName: &node.Identifier{Value: "b"},
							},
						},
						&expr.ArrayItem{
							Val: &expr.Variable{
								Meta: meta.Collection{
									&meta.Data{
										Type:      meta.TokenType,
										Value:     "$",
										TokenName: meta.NodeStart,
									},
								},
								VarName: &node.Identifier{Value: "c"},
							},
						},
					},
				},
			},
		},
	})

	expected := ` list ($a,list($b,$c) )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintMethodCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.MethodCall{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ObjectOperatorToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "foo"},
		},
		Method: &node.Identifier{Value: "bar"},
		ArgumentList: &node.ArgumentList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenParenthesisToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseParenthesisToken,
				},
			},
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `$foo ->bar ($a,$b )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNew(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.New{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NewAnchor,
			},
		},
		Class: &name.Name{
			Parts: []node.Node{
				&name.NamePart{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.WhiteSpaceType,
							Value:     " ",
							TokenName: meta.StringToken,
						},
					},
					Value: "Foo",
				},
			},
		},
		ArgumentList: &node.ArgumentList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenParenthesisToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseParenthesisToken,
				},
			},
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := ` new Foo ($a,$b )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostDec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PostDec{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DecToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `$var --`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostInc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PostInc{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IncToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := `$var ++`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreDec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PreDec{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DecToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` --$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreInc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PreInc{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IncToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` ++$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPrint(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Print{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PrintToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` print $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.PropertyFetch{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ObjectOperatorToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "foo"},
		},
		Property: &node.Identifier{Value: "bar"},
	})

	expected := `$foo ->bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Reference{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "foo"},
		},
	})

	expected := ` &$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequire(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Require{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.RequireToken,
			},
		},
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := ` require'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.RequireOnce{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.RequireOnceToken,
			},
		},
		Expr: &scalar.String{Value: "'path'"},
	})

	expected := ` require_once'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ShellExec{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.BackquoteToken,
			},
		},
		Parts: []node.Node{
			&scalar.EncapsedStringPart{Value: "hello "},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "world"},
			},
			&scalar.EncapsedStringPart{Value: "!"},
		},
	})

	expected := " `hello $world!`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprShortArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ShortArray{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenSquareBracket,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseSquareBracket,
			},
		},
		Items: []node.Node{
			&expr.ArrayItem{
				Key: &scalar.String{Value: "'Hello'"},
				Val: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "world"},
				},
			},
			&expr.ArrayItem{
				Key: &scalar.Lnumber{Value: "2"},
				Val: &expr.Reference{Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "var"},
				}},
			},
			&expr.ArrayItem{
				Val: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
	})

	expected := ` ['Hello'=>$world,2=>&$var,$var ]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShortList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.ShortList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenSquareBracket,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseSquareBracket,
			},
		},
		Items: []node.Node{
			&expr.ArrayItem{
				Val: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&expr.ArrayItem{
				Val: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{
								Meta: meta.Collection{
									&meta.Data{
										Type:      meta.TokenType,
										Value:     "$",
										TokenName: meta.NodeStart,
									},
								},
								VarName: &node.Identifier{Value: "b"},
							},
						},
						&expr.ArrayItem{
							Val: &expr.Variable{
								Meta: meta.Collection{
									&meta.Data{
										Type:      meta.TokenType,
										Value:     "$",
										TokenName: meta.NodeStart,
									},
								},
								VarName: &node.Identifier{Value: "c"},
							},
						},
					},
				},
			},
		},
	})

	expected := ` [$a,list($b,$c) ]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.StaticCall{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PaamayimNekudotayimToken,
			},
		},
		Class: &node.Identifier{Value: "Foo"},
		Call:  &node.Identifier{Value: "bar"},
		ArgumentList: &node.ArgumentList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenParenthesisToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseParenthesisToken,
				},
			},
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := `Foo ::bar ($a,$b )`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.StaticPropertyFetch{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PaamayimNekudotayimToken,
			},
		},
		Class: &node.Identifier{Value: "Foo"},
		Property: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "bar"},
		},
	})

	expected := `Foo ::$bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernary(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Ternary{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.QuestionMarkToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		Condition: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		IfFalse: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
	})

	expected := `$a ? :$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernaryFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Ternary{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.QuestionMarkToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		Condition: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		IfTrue: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "b"},
		},
		IfFalse: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "c"},
		},
	})

	expected := `$a ?$b :$c`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.UnaryMinus{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.MinusToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` -$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.UnaryPlus{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PlusToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` +$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintVariable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Variable{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NodeStart,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "$",
				TokenName: meta.NodeStart,
			},
		},
		VarName: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` $$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.YieldFrom{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.YieldFromToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` yield from$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYield(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Yield{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.YieldToken,
			},
		},
		Value: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` yield$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&expr.Yield{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.YieldToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoubleArrowToken,
			},
		},
		Key: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "k"},
		},
		Value: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` yield$k =>$var`
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseifToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				}},
			},
		},
	})

	expected := ` elseif ($a ) :$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseIfEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElseIf{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseifToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{},
	})

	expected := ` elseif ($a ) :`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElse{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				}},
			},
		},
	})

	expected := ` else :$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltElse{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		Stmt: &stmt.StmtList{},
	})

	expected := ` else :`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltFor{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForInitSemicolonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForCondSemicolonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EndforToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Init: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
		},
		Cond: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
		Loop: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "c"},
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "d"},
				}},
			},
		},
	})

	expected := ` for ($a ;$b ;$c ) :$d endfor `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltForeach{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForeachToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AsToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoubleArrowToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EndforeachToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		Key: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "key"},
		},
		Variable: &expr.Reference{Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "val"},
		}},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "d"},
				}},
			},
		},
	})

	expected := ` foreach ($var as$key =>&$val ) :$d endforeach `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltIf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltIf{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IfToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EndifToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "d"},
				}},
			},
		},
		ElseIf: []node.Node{
			&stmt.AltElseIf{
				Cond: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "b"},
						}},
					},
				},
			},
			&stmt.AltElseIf{
				Cond: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "c"},
				},
				Stmt: &stmt.StmtList{},
			},
		},
		Else: &stmt.AltElse{
			Stmt: &stmt.StmtList{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					}},
				},
			},
		},
	})

	expected := ` if ($a ) :$delseif($b):$belseif($c):else:$b endif `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAltSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltSwitch{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SwitchToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EndswitchToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		CaseList: &stmt.CaseList{
			Cases: []node.Node{
				&stmt.Case{
					Cond: &scalar.String{Value: "'a'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
				&stmt.Case{
					Cond: &scalar.String{Value: "'b'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "b"},
						}},
					},
				},
			},
		},
	})

	expected := ` switch ($var ) :case'a':$acase'b':$b endswitch `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltWhile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.AltWhile{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.WhileToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EndwhileToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				}},
			},
		},
	})

	expected := ` while ($a ) :$b endwhile `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtBreak(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Break{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.BreakToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Expr: &scalar.Lnumber{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
			},
			Value: "1",
		},
	})

	expected := " break 1 "
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCase(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Case{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CaseToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := ` case$a:$a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCaseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Case{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CaseToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmts: []node.Node{},
	})

	expected := " case$a:"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCatch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Catch{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CatchToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Types: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Exception"}}},
			&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "RuntimeException"}}},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "e"},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := ` catch (Exception|\RuntimeException$e ) {$a }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ClassMethod{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.FunctionToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
		ReturnsRef: true,
		MethodName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        true,
				VariableType: &node.Nullable{Expr: &name.Name{Parts: []node.Node{&name.NamePart{Value: "int"}}}},
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				},
				DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
			},
			&node.Parameter{
				Variadic: true,
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
		ReturnType: &name.Name{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ColonToken,
				},
			},
			Parts: []node.Node{&name.NamePart{Value: "void"}},
		},
		Stmt: &stmt.StmtList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenCurlyBracesToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseCurlyBracesToken,
				},
			},
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				}},
			},
		},
	})

	expected := `public function &foo (?int&$a=null,...$b ) :void {$a }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAbstractClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ClassMethod{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.FunctionToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
		ReturnsRef: true,
		MethodName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        true,
				VariableType: &node.Nullable{Expr: &name.Name{Parts: []node.Node{&name.NamePart{Value: "int"}}}},
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				},
				DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
			},
			&node.Parameter{
				Variadic: true,
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
		ReturnType: &name.Name{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ColonToken,
				},
			},
			Parts: []node.Node{&name.NamePart{Value: "void"}},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `public function &foo (?int&$a=null,...$b ) :void`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Class{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ClassToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
		ClassName: &node.Identifier{Value: "Foo"},
		Extends: &stmt.ClassExtends{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ExtendsToken,
				},
			},
			ClassName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		},
		Implements: &stmt.ClassImplements{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ImplementsToken,
				},
			},
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

	expected := `abstract classFoo extendsBar implementsBaz,Quuz {publicconstFOO='bar' }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAnonymousClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Class{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ClassToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
		ArgumentList: &node.ArgumentList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenParenthesisToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseParenthesisToken,
				},
			},
			Arguments: []node.Node{
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "a"},
					},
				},
				&node.Argument{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
		Extends: &stmt.ClassExtends{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ExtendsToken,
				},
			},
			ClassName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		},
		Implements: &stmt.ClassImplements{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ImplementsToken,
				},
			},
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

	expected := `abstract class ($a,$b ) extendsBar implementsBaz,Quuz {publicconstFOO='bar' }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ClassConstList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ConstToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
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

	expected := `public constFOO='a',BAR='b' `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ConstList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ConstToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
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

	expected := ` constFOO='a',BAR='b' `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Constant{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EqualToken,
			},
		},
		ConstantName: &node.Identifier{Value: "FOO"},
		Expr:         &scalar.String{Value: "'BAR'"},
	})

	expected := "FOO ='BAR'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtContinue(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Continue{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ContinueToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Expr: &scalar.Lnumber{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.NodeStart,
				},
			},
			Value: "1",
		},
	})

	expected := ` continue 1 `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Declare{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DeclareToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
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

	expected := ` declare (FOO='bar' ){}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Declare{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DeclareToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Consts: []node.Node{
			&stmt.Constant{
				ConstantName: &node.Identifier{Value: "FOO"},
				Expr:         &scalar.String{Value: "'bar'"},
			},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := ` declare (FOO='bar' )'bar'`
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

	expected := `declare(FOO='bar')`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalut(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Default{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DefaultToken,
			},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := ` default:$a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalutEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Default{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DefaultToken,
			},
		},
		Stmts: []node.Node{},
	})

	expected := ` default:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_Expression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Do{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.WhileToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Cond: &scalar.Lnumber{Value: "1"},
		Stmt: &stmt.Expression{
			Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
		},
	})

	expected := ` do$a while (1 ) `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_StmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Do{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.WhileToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Cond: &scalar.Lnumber{Value: "1"},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				}},
			},
		},
	})

	expected := ` do{$a} while (1 ) `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEcho(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Echo{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "echo",
				TokenName: meta.EchoToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EchoToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Exprs: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := `echo $a,$b `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ElseIf{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseifToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := ` elseif ($a ){}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.ElseIf{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseifToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := ` elseif ($a )'bar'`
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
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `elseif($a)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Else{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseToken,
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := ` else{}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Else{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ElseToken,
			},
		},
		Stmt: &stmt.Expression{Expr: &scalar.String{Value: "'bar'"}},
	})

	expected := ` else'bar'`
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

	expected := `else`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Expression{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
	})

	expected := `$a `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Finally{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.FinallyToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Stmts: []node.Node{
			&stmt.Nop{},
		},
	})

	expected := ` finally { }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.For{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForInitSemicolonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForCondSemicolonToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Init: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
		Cond: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "c"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "d"},
			},
		},
		Loop: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "e"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "f"},
			},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := ` for ($a,$b ;$c,$d ;$e,$f ){}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Foreach{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ForeachToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AsToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.DoubleArrowToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Key: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "k"},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "v"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Nop{},
			},
		},
	})

	expected := ` foreach ($a as$k =>$v ){}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Function{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.FunctionToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AmpersandToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		ReturnsRef:   true,
		FunctionName: &node.Identifier{Value: "foo"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:    true,
				Variadic: false,
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "var"},
				},
			},
		},
		ReturnType: &name.FullyQualified{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ColonToken,
				},
			},
			Parts: []node.Node{&name.NamePart{Value: "Foo"}},
		},
		Stmts: []node.Node{
			&stmt.Nop{},
		},
	})

	expected := ` function &foo (&$var ) :\Foo { }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGlobal(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Global{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.GlobalToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Vars: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := ` global$a,$b `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGoto(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Goto{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.GotoToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Label: &node.Identifier{Value: "FOO"},
	})

	expected := ` gotoFOO `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.GroupUse{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UseToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NsSeparatorToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
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

	expected := ` usefunctionFoo \ {BarasBaz,Quuz } `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.HaltCompiler{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.HaltCompilerToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
	})

	expected := ` __halt_compiler ( ) `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.If{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IfToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Expression{
			Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
		ElseIf: []node.Node{
			&stmt.ElseIf{
				Cond: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "c"},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{
							Expr: &expr.Variable{
								Meta: meta.Collection{
									&meta.Data{
										Type:      meta.TokenType,
										Value:     "$",
										TokenName: meta.NodeStart,
									},
								},
								VarName: &node.Identifier{Value: "d"},
							},
						},
					},
				},
			},
			&stmt.ElseIf{
				Cond: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "e"},
				},
				Stmt: &stmt.Nop{},
			},
		},
		Else: &stmt.Else{
			Stmt: &stmt.Expression{
				Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "f"},
				},
			},
		},
	})

	expected := ` if ($a )$belseif($c){$d}elseif($e)else$f`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.If{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.IfToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{
					Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					},
				},
			},
		},
	})

	expected := ` if ($a ){$b}`
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
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.Nop{},
	})

	expected := `if($a)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.InlineHtml{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.InlineHTMLToken,
			},
		},
		Value: "test",
	})

	expected := ` test`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInterface(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Interface{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.InterfaceToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		InterfaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Extends: &stmt.InterfaceExtends{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.ExtendsToken,
				},
			},
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
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
			},
		},
	})

	expected := ` interfaceFoo extendsBar,Baz {publicfunctionfoo(){$a} }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintLabel(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Label{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ColonToken,
			},
		},
		LabelName: &node.Identifier{Value: "FOO"},
	})

	expected := `FOO :`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespace(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Namespace{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NamespaceToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
	})

	expected := ` namespaceFoo `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespaceWithStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Namespace{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NamespaceToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
		},
	})

	expected := ` namespaceFoo {$a }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Nop{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.NodeStart,
			},
		},
	})

	expected := ` `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.PropertyList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Modifiers: []node.Node{
			&node.Identifier{Value: "public"},
			&node.Identifier{Value: "static"},
		},
		Properties: []node.Node{
			&stmt.Property{
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&stmt.Property{
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
	})

	expected := `publicstatic$a,$b `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintProperty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Property{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `$a =1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReturn(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Return{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ReturnToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := ` return1 `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticVar(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.StaticVar{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.EqualToken,
			},
		},
		Variable: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Expr: &scalar.Lnumber{Value: "1"},
	})

	expected := `$a =1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStatic(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Static{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.StaticToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Vars: []node.Node{
			&stmt.StaticVar{
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				},
			},
			&stmt.StaticVar{
				Variable: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "b"},
				},
			},
		},
	})

	expected := ` static$a,$b `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.StmtList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			}},
		},
	})

	expected := ` {$a$b }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtListNested(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.StmtList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			}},
			&stmt.StmtList{
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
						VarName: &node.Identifier{Value: "b"},
					}},
					&stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Expression{Expr: &expr.Variable{
								Meta: meta.Collection{
									&meta.Data{
										Type:      meta.TokenType,
										Value:     "$",
										TokenName: meta.NodeStart,
									},
								},
								VarName: &node.Identifier{Value: "c"},
							}},
						},
					},
				},
			},
		},
	})

	expected := ` {$a{$b{$c}} }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Switch{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SwitchToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
		CaseList: &stmt.CaseList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenCurlyBracesToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseCurlyBracesToken,
				},
			},
			Cases: []node.Node{
				&stmt.Case{
					Cond: &scalar.String{Value: "'a'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
				&stmt.Case{
					Cond: &scalar.String{Value: "'b'"},
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "b"},
						}},
					},
				},
			},
		},
	})

	expected := ` switch ($var ) {case'a':$acase'b':$b }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtThrow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Throw{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.ThrowToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Expr: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "var"},
		},
	})

	expected := ` throw$var `
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
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.PaamayimNekudotayimToken,
			},
		},
		Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Method: &node.Identifier{Value: "a"},
	})

	expected := `Foo ::a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUseAlias(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitUseAlias{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AsToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Ref: &stmt.TraitMethodRef{
			Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			Method: &node.Identifier{Value: "a"},
		},
		Modifier: &node.Identifier{Value: "public"},
		Alias:    &node.Identifier{Value: "b"},
	})

	expected := `Foo::a aspublicb `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUsePrecedence(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitUsePrecedence{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.InsteadofToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Ref: &stmt.TraitMethodRef{
			Trait:  &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			Method: &node.Identifier{Value: "a"},
		},
		Insteadof: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
		},
	})

	expected := `Foo::a insteadofBar,Baz `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitUse{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UseToken,
			},
		},
		Traits: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		},
		TraitAdaptationList: &stmt.Nop{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.SemiColonToken,
				},
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "",
					TokenName: meta.SemiColonToken,
				},
			},
		},
	})

	expected := ` useFoo,Bar `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitAdaptations(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.TraitUse{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UseToken,
			},
		},
		Traits: []node.Node{
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
			&name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
		},
		TraitAdaptationList: &stmt.TraitAdaptationList{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.OpenCurlyBracesToken,
				},
				&meta.Data{
					Type:      meta.WhiteSpaceType,
					Value:     " ",
					TokenName: meta.CloseCurlyBracesToken,
				},
			},
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

	expected := ` useFoo,Bar {Foo::aasb }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTrait(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Trait{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.TraitToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		TraitName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Stmts: []node.Node{
			&stmt.ClassMethod{
				Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
				MethodName: &node.Identifier{Value: "foo"},
				Params:     []node.Node{},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Expression{Expr: &expr.Variable{
							Meta: meta.Collection{
								&meta.Data{
									Type:      meta.TokenType,
									Value:     "$",
									TokenName: meta.NodeStart,
								},
							},
							VarName: &node.Identifier{Value: "a"},
						}},
					},
				},
			},
		},
	})

	expected := ` traitFoo {publicfunctionfoo(){$a} }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTry(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Try{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.TryToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenCurlyBracesToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseCurlyBracesToken,
			},
		},
		Stmts: []node.Node{
			&stmt.Expression{Expr: &expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
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
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "e"},
				},
				Stmts: []node.Node{
					&stmt.Expression{Expr: &expr.Variable{
						Meta: meta.Collection{
							&meta.Data{
								Type:      meta.TokenType,
								Value:     "$",
								TokenName: meta.NodeStart,
							},
						},
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

	expected := ` try {$a }catch(Exception|\RuntimeException$e){$b}finally{}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Unset{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UnsetToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
		Vars: []node.Node{
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "a"},
			},
			&expr.Variable{
				Meta: meta.Collection{
					&meta.Data{
						Type:      meta.TokenType,
						Value:     "$",
						TokenName: meta.NodeStart,
					},
				},
				VarName: &node.Identifier{Value: "b"},
			},
		},
	})

	expected := ` unset ($a,$b ) `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUseList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.UseList{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.UseToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.SemiColonToken,
			},
			&meta.Data{
				Type:      meta.TokenType,
				Value:     "",
				TokenName: meta.SemiColonToken,
			},
		},
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

	expected := ` usefunctionFooasBar,Baz `
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.Use{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.AsToken,
			},
		},
		UseType: &node.Identifier{Value: "function"},
		Use:     &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		Alias:   &node.Identifier{Value: "Bar"},
	})

	expected := `functionFoo asBar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintWhileStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(&stmt.While{
		Meta: meta.Collection{
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.WhileToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.OpenParenthesisToken,
			},
			&meta.Data{
				Type:      meta.WhiteSpaceType,
				Value:     " ",
				TokenName: meta.CloseParenthesisToken,
			},
		},
		Cond: &expr.Variable{
			Meta: meta.Collection{
				&meta.Data{
					Type:      meta.TokenType,
					Value:     "$",
					TokenName: meta.NodeStart,
				},
			},
			VarName: &node.Identifier{Value: "a"},
		},
		Stmt: &stmt.StmtList{
			Stmts: []node.Node{
				&stmt.Expression{Expr: &expr.Variable{
					Meta: meta.Collection{
						&meta.Data{
							Type:      meta.TokenType,
							Value:     "$",
							TokenName: meta.NodeStart,
						},
					},
					VarName: &node.Identifier{Value: "a"},
				}},
			},
		},
	})

	expected := ` while ($a ){$a}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
