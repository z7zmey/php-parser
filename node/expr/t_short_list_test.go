package expr_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
)

func TestShortList(t *testing.T) {
	src := `<? [$a] = $b;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    13,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    13,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    12,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  4,
							EndPos:    7,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    6,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  5,
										EndPos:    6,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  5,
											EndPos:    6,
										},
										Value: "a",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    12,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    12,
							},
							Value: "b",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestShortListArrayIndex(t *testing.T) {
	src := `<? [$a[]] = $b;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    15,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    15,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    14,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  4,
							EndPos:    9,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    8,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  5,
										EndPos:    8,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  5,
											EndPos:    6,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  5,
												EndPos:    6,
											},
											Value: "a",
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    14,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    14,
							},
							Value: "b",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestShortListList(t *testing.T) {
	src := `<? [list($a)] = $b;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    19,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    19,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    18,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  4,
							EndPos:    13,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    12,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  5,
										EndPos:    12,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  10,
												EndPos:    11,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  10,
													EndPos:    11,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  10,
														EndPos:    11,
													},
													Value: "a",
												},
											},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  17,
							EndPos:    18,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    18,
							},
							Value: "b",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
