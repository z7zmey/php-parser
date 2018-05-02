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

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
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

func TestElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
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

func TestElse(t *testing.T) {
	src := `<? if ($a) {} else {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				Else: &stmt.Else{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
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

func TestElseElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} elseif ($c) {} else {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.Else{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
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

func TestElseIfElseIfElse(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} else if ($c) {} else {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.Else{
					Stmt: &stmt.If{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
						Else: &stmt.Else{
							Stmt: &stmt.StmtList{Stmts: []node.Node{}},
						},
					},
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
