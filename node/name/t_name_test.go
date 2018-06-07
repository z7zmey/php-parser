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

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{
						Parts: []node.Node{&name.NamePart{Value: "foo"}},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestFullyQualified(t *testing.T) {
	src := `<? \foo();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.FullyQualified{
						Parts: []node.Node{&name.NamePart{Value: "foo"}},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestRelative(t *testing.T) {
	src := `<? namespace\foo();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Relative{
						Parts: []node.Node{&name.NamePart{Value: "foo"}},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestNamePartsGetter(t *testing.T) {
	expected := []node.Node{
		&name.NamePart{Value: "a"},
		&name.NamePart{Value: "b"},
	}

	plainName := &name.Name{Parts: expected}
	relativeName := &name.Relative{Parts: expected}
	fullyQualifiedName := &name.FullyQualified{Parts: expected}

	assertEqual(t, expected, plainName.GetParts())
	assertEqual(t, expected, relativeName.GetParts())
	assertEqual(t, expected, fullyQualifiedName.GetParts())
}
