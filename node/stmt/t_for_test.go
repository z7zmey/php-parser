package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/binary_op"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign_op"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestFor(t *testing.T) {
	src := `<? for($i = 0; $i < 10; $i++, $i++) {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.For{
				Init: []node.Node{
					&assign_op.Assign{
						Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
						Expression: &scalar.Lnumber{Value: "0"},
					},
				},
				Cond: []node.Node{
					&binary_op.Smaller{
						Left:  &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
						Right: &scalar.Lnumber{Value: "10"},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
					},
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
					},
				},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestAltFor(t *testing.T) {
	src := `<? for($i = 0; $i < 10; $i++, $i++) : endfor;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.For{
				Init: []node.Node{
					&assign_op.Assign{
						Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
						Expression: &scalar.Lnumber{Value: "0"},
					},
				},
				Cond: []node.Node{
					&binary_op.Smaller{
						Left:  &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
						Right: &scalar.Lnumber{Value: "10"},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
					},
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "$i"}},
					},
				},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
