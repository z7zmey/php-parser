package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestIf(t *testing.T) {
	src := `<? if ($a) {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.If{
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

func TestElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {}
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
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

func TestElse(t *testing.T) {
	src := `<? if ($a) {} else {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				Else: &stmt.Else{
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

func TestElseElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} elseif ($c) {} else {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.Else{
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

func TestElseIfElseIfElse(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} else if ($c) {} else {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.Else{
					Stmt: &stmt.If{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
						Else: &stmt.Else{
							Stmt: &stmt.StmtList{Stmts: []node.Node{}},
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
