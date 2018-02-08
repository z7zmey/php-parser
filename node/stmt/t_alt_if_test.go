package stmt_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

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
			t.Errorf("expected and actual are not equal\nexpectd: %+v\nactual: %+v\n", expected, actual)
		}

	}
}

func TestAltIf(t *testing.T) {
	src := `<?
		if ($a) :
		endif;
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)


	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestAltElseIf(t *testing.T) {
	src := `<?
		if ($a) :
		elseif ($b):
		endif;
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.AltElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
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

func TestAltElse(t *testing.T) {
	src := `<?
		if ($a) :
		else:
		endif;
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				Else: &stmt.AltElse{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)


	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestAltElseElseIf(t *testing.T) {
	src := `<?
		if ($a) :
		elseif ($b):
		elseif ($c):
		else:
		endif;
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.AltElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
					&stmt.AltElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.AltElse{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)


	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
