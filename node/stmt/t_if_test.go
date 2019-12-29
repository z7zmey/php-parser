package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestIf(t *testing.T) {
	src := `<? if ($a) {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    13,
		},
		Stmts: []node.Node{
			&stmt.If{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    13,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  7,
						EndPos:    9,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  11,
						EndPos:    13,
					},
					Stmts: []node.Node{},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    28,
		},
		Stmts: []node.Node{
			&stmt.If{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    28,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  7,
						EndPos:    9,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  11,
						EndPos:    13,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    28,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  22,
								EndPos:    24,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  26,
								EndPos:    28,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestElse(t *testing.T) {
	src := `<? if ($a) {} else {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    21,
		},
		Stmts: []node.Node{
			&stmt.If{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    21,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  7,
						EndPos:    9,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  11,
						EndPos:    13,
					},
					Stmts: []node.Node{},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  14,
						EndPos:    21,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    21,
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestElseElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} elseif ($c) {} else {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    51,
		},
		Stmts: []node.Node{
			&stmt.If{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    51,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  7,
						EndPos:    9,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  11,
						EndPos:    13,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    28,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  22,
								EndPos:    24,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  26,
								EndPos:    28,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  29,
							EndPos:    43,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  37,
								EndPos:    39,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  37,
									EndPos:    39,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  41,
								EndPos:    43,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  44,
						EndPos:    51,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  49,
							EndPos:    51,
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestElseIfElseIfElse(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} else if ($c) {} else {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    52,
		},
		Stmts: []node.Node{
			&stmt.If{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    52,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  7,
						EndPos:    9,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  11,
						EndPos:    13,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    28,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  22,
								EndPos:    24,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  26,
								EndPos:    28,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  29,
						EndPos:    52,
					},
					Stmt: &stmt.If{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  34,
							EndPos:    52,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  38,
								EndPos:    40,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  38,
									EndPos:    40,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  42,
								EndPos:    44,
							},
							Stmts: []node.Node{},
						},
						Else: &stmt.Else{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  45,
								EndPos:    52,
							},
							Stmt: &stmt.StmtList{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  50,
									EndPos:    52,
								},
								Stmts: []node.Node{},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
