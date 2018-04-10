package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestUnset(t *testing.T) {
	src := `<? unset($a);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
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

func TestUnsetVars(t *testing.T) {
	src := `<? unset($a, $b);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestUnsetTrailingComma(t *testing.T) {
	src := `<? unset($a, $b,);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
