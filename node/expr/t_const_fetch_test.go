package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/php5"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
)

func TestConstFetch(t *testing.T) {
	src := `<? foo;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.ConstFetch{
					Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestConstFetchRelative(t *testing.T) {
	src := `<? namespace\foo;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.ConstFetch{
					Constant: &name.Relative{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestConstFetchFullyQualified(t *testing.T) {
	src := `<? \foo;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.ConstFetch{
					Constant: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
