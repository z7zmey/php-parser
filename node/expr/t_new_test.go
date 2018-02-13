package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestNew(t *testing.T) {
	src := `<? new Foo;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
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

func TestNewRelative(t *testing.T) {
	src := `<? new namespace\Foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Arguments: []node.Node{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestNewFullyQualified(t *testing.T) {
	src := `<? new \Foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Arguments: []node.Node{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestNewAnonymous(t *testing.T) {
	src := `<? new class ($a, ...$b) {};`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.New{
					Class: &stmt.Class{
						PhpDocComment: "",
						Args: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "$a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "$b"}}},
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
