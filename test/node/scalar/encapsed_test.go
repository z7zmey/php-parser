package test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/parser"
)

func TestSimpleVar(t *testing.T) {
	src := `<? "test $var";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "$var"}},
					},
				},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestSimpleVarPropertyFetch(t *testing.T) {
	src := `<? "test $foo->bar()";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.PropertyFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "$foo"}},
							Property: &node.Identifier{Value: "bar"},
						},
						&scalar.EncapsedStringPart{Value: "()"},
					},
				},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestDollarOpenCurlyBraces(t *testing.T) {
	src := `<? "test ${foo}";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					},
				},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestDollarOpenCurlyBracesDimNumber(t *testing.T) {
	src := `<? "test ${foo[0]}";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Dim:      &scalar.Lnumber{Value: "0"},
						},
					},
				},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestCurlyOpenMethodCall(t *testing.T) {
	src := `<? "test {$foo->bar()}";`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.MethodCall{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "$foo"}},
							Method:   &node.Identifier{Value: "bar"},
						},
					},
				},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
