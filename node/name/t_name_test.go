package name_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
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

func TestName(t *testing.T) {
	src := `<? foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{
						Parts: []node.Node{&name.NamePart{Value: "foo"}},
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

func TestFullyQualified(t *testing.T) {
	src := `<? \foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.FullyQualified{
						Parts: []node.Node{&name.NamePart{Value: "foo"}},
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

func TestRelative(t *testing.T) {
	src := `<? namespace\foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Relative{
						Parts: []node.Node{&name.NamePart{Value: "foo"}},
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
