package stmt_test

import (
	"bytes"
	"testing"
	
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestGlobal(t *testing.T) {
	src := `<? global $a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Global{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestGlobalVars(t *testing.T) {
	src := `<? global $a, $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Global{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
