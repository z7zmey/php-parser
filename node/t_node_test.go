package node_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestIdentifier(t *testing.T) {
	src := `<? $foo;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    8,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    8,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    7,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  4,
							EndPos:    7,
						},
						Value: "foo",
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp7ArgumentNode(t *testing.T) {
	src := `<? 
		foo($a, ...$b);
		$foo($a, ...$b);
		$foo->bar($a, ...$b);
		foo::bar($a, ...$b);
		$foo::bar($a, ...$b);
		new foo($a, ...$b);
		/** anonymous class */
		new class ($a, ...$b) {};
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   9,
			StartPos:  7,
			EndPos:    186,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  7,
					EndPos:    21,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  7,
						EndPos:    20,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  7,
							EndPos:    9,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  7,
									EndPos:    9,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  10,
							EndPos:    20,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  11,
									EndPos:    12,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  11,
										EndPos:    12,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  11,
											EndPos:    12,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  15,
									EndPos:    19,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  18,
										EndPos:    19,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  18,
											EndPos:    19,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  25,
					EndPos:    40,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  25,
						EndPos:    39,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  25,
							EndPos:    28,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  25,
								EndPos:    28,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  29,
							EndPos:    39,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  30,
									EndPos:    31,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  30,
										EndPos:    31,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  30,
											EndPos:    31,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  34,
									EndPos:    38,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  37,
										EndPos:    38,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  37,
											EndPos:    38,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   4,
					StartPos:  44,
					EndPos:    64,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  44,
						EndPos:    63,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  44,
							EndPos:    47,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  44,
								EndPos:    47,
							},
							Value: "foo",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  50,
							EndPos:    52,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  53,
							EndPos:    63,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  54,
									EndPos:    55,
								},
								IsReference: false,
								Variadic:    false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  54,
										EndPos:    55,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  54,
											EndPos:    55,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  58,
									EndPos:    62,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  61,
										EndPos:    62,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  61,
											EndPos:    62,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 5,
					EndLine:   5,
					StartPos:  68,
					EndPos:    87,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  68,
						EndPos:    86,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  68,
							EndPos:    70,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  68,
									EndPos:    70,
								},
								Value: "foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  73,
							EndPos:    75,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  76,
							EndPos:    86,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  77,
									EndPos:    78,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  77,
										EndPos:    78,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  77,
											EndPos:    78,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  81,
									EndPos:    85,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  84,
										EndPos:    85,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  84,
											EndPos:    85,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 6,
					EndLine:   6,
					StartPos:  91,
					EndPos:    111,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  91,
						EndPos:    110,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  91,
							EndPos:    94,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  91,
								EndPos:    94,
							},
							Value: "foo",
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  97,
							EndPos:    99,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  100,
							EndPos:    110,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  101,
									EndPos:    102,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  101,
										EndPos:    102,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  101,
											EndPos:    102,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  105,
									EndPos:    109,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  108,
										EndPos:    109,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  108,
											EndPos:    109,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   7,
					StartPos:  115,
					EndPos:    133,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  115,
						EndPos:    132,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  119,
							EndPos:    121,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  119,
									EndPos:    121,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  122,
							EndPos:    132,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  123,
									EndPos:    124,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  123,
										EndPos:    124,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  123,
											EndPos:    124,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  127,
									EndPos:    131,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  130,
										EndPos:    131,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  130,
											EndPos:    131,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 9,
					EndLine:   9,
					StartPos:  162,
					EndPos:    186,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  162,
						EndPos:    185,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  166,
							EndPos:    185,
						},
						PhpDocComment: "/** anonymous class */",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  172,
								EndPos:    182,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  173,
										EndPos:    174,
									},
									Variadic:    false,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  173,
											EndPos:    174,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  173,
												EndPos:    174,
											},
											Value: "a",
										},
									},
								},
								&node.Argument{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  177,
										EndPos:    181,
									},
									Variadic:    true,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  180,
											EndPos:    181,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  180,
												EndPos:    181,
											},
											Value: "b",
										},
									},
								},
							},
						},
						Stmts: []node.Node{},
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

func TestPhp5ArgumentNode(t *testing.T) {
	src := `<? 
		foo($a, ...$b);
		$foo($a, ...$b);
		$foo->bar($a, ...$b);
		foo::bar($a, ...$b);
		$foo::bar($a, ...$b);
		new foo($a, ...$b);
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   7,
			StartPos:  7,
			EndPos:    133,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  7,
					EndPos:    21,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  7,
						EndPos:    20,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  7,
							EndPos:    9,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  7,
									EndPos:    9,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  10,
							EndPos:    20,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  11,
									EndPos:    12,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  11,
										EndPos:    12,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  11,
											EndPos:    12,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  15,
									EndPos:    19,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  18,
										EndPos:    19,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  18,
											EndPos:    19,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  25,
					EndPos:    40,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  25,
						EndPos:    39,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  25,
							EndPos:    28,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  25,
								EndPos:    28,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  29,
							EndPos:    39,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  30,
									EndPos:    31,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  30,
										EndPos:    31,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  30,
											EndPos:    31,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  34,
									EndPos:    38,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  37,
										EndPos:    38,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  37,
											EndPos:    38,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   4,
					StartPos:  44,
					EndPos:    64,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  44,
						EndPos:    63,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  44,
							EndPos:    47,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  44,
								EndPos:    47,
							},
							Value: "foo",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  50,
							EndPos:    52,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  53,
							EndPos:    63,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  54,
									EndPos:    55,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  54,
										EndPos:    55,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  54,
											EndPos:    55,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  58,
									EndPos:    62,
								},
								IsReference: false,
								Variadic:    true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  61,
										EndPos:    62,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  61,
											EndPos:    62,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 5,
					EndLine:   5,
					StartPos:  68,
					EndPos:    87,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  68,
						EndPos:    86,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  68,
							EndPos:    70,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  68,
									EndPos:    70,
								},
								Value: "foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  73,
							EndPos:    75,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  76,
							EndPos:    86,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  77,
									EndPos:    78,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  77,
										EndPos:    78,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  77,
											EndPos:    78,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  81,
									EndPos:    85,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  84,
										EndPos:    85,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  84,
											EndPos:    85,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 6,
					EndLine:   6,
					StartPos:  91,
					EndPos:    111,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  91,
						EndPos:    110,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  91,
							EndPos:    94,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  91,
								EndPos:    94,
							},
							Value: "foo",
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  97,
							EndPos:    99,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  100,
							EndPos:    110,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  101,
									EndPos:    102,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  101,
										EndPos:    102,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  101,
											EndPos:    102,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  105,
									EndPos:    109,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  108,
										EndPos:    109,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  108,
											EndPos:    109,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   7,
					StartPos:  115,
					EndPos:    133,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  115,
						EndPos:    132,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  119,
							EndPos:    121,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  119,
									EndPos:    121,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  122,
							EndPos:    132,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  123,
									EndPos:    124,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  123,
										EndPos:    124,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  123,
											EndPos:    124,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  127,
									EndPos:    131,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  130,
										EndPos:    131,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  130,
											EndPos:    131,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp7ParameterNode(t *testing.T) {
	src := `<? 
		function foo(?bar $bar=null, baz &...$baz) {}
		class foo {public function foo(?bar $bar=null, baz &...$baz) {}}
		function(?bar $bar=null, baz &...$baz) {};
		static function(?bar $bar=null, baz &...$baz) {};
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   5,
			StartPos:  7,
			EndPos:    215,
		},
		Stmts: []node.Node{
			&stmt.Function{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  7,
					EndPos:    51,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  16,
						EndPos:    18,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  20,
							EndPos:    33,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Nullable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  20,
								EndPos:    23,
							},
							Expr: &name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  21,
									EndPos:    23,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  21,
											EndPos:    23,
										},
										Value: "bar",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  25,
								EndPos:    28,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  25,
									EndPos:    28,
								},
								Value: "bar",
							},
						},
						DefaultValue: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  30,
								EndPos:    33,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    33,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  30,
											EndPos:    33,
										},
										Value: "null",
									},
								},
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  36,
							EndPos:    47,
						},
						ByRef:    true,
						Variadic: true,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  36,
								EndPos:    38,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  36,
										EndPos:    38,
									},
									Value: "baz",
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  44,
								EndPos:    47,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  44,
									EndPos:    47,
								},
								Value: "baz",
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  55,
					EndPos:    118,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  61,
						EndPos:    63,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  66,
							EndPos:    117,
						},
						PhpDocComment: "",
						ReturnsRef:    false,
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  82,
								EndPos:    84,
							},
							Value: "foo",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  66,
									EndPos:    71,
								},
								Value: "public",
							},
						},
						Params: []node.Node{
							&node.Parameter{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  86,
									EndPos:    99,
								},
								ByRef:    false,
								Variadic: false,
								VariableType: &node.Nullable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  86,
										EndPos:    89,
									},
									Expr: &name.Name{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  87,
											EndPos:    89,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 3,
													EndLine:   3,
													StartPos:  87,
													EndPos:    89,
												},
												Value: "bar",
											},
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  91,
										EndPos:    94,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  91,
											EndPos:    94,
										},
										Value: "bar",
									},
								},
								DefaultValue: &expr.ConstFetch{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  96,
										EndPos:    99,
									},
									Constant: &name.Name{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  96,
											EndPos:    99,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 3,
													EndLine:   3,
													StartPos:  96,
													EndPos:    99,
												},
												Value: "null",
											},
										},
									},
								},
							},
							&node.Parameter{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  102,
									EndPos:    113,
								},
								ByRef:    true,
								Variadic: true,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  102,
										EndPos:    104,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  102,
												EndPos:    104,
											},
											Value: "baz",
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  110,
										EndPos:    113,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  110,
											EndPos:    113,
										},
										Value: "baz",
									},
								},
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  116,
								EndPos:    117,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   4,
					StartPos:  122,
					EndPos:    163,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  122,
						EndPos:    162,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  131,
								EndPos:    144,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &node.Nullable{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  131,
									EndPos:    134,
								},
								Expr: &name.Name{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  132,
										EndPos:    134,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  132,
												EndPos:    134,
											},
											Value: "bar",
										},
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  136,
									EndPos:    139,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  136,
										EndPos:    139,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  141,
									EndPos:    144,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  141,
										EndPos:    144,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  141,
												EndPos:    144,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  147,
								EndPos:    158,
							},
							Variadic: true,
							ByRef:    true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  147,
									EndPos:    149,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  147,
											EndPos:    149,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  155,
									EndPos:    158,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  155,
										EndPos:    158,
									},
									Value: "baz",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 5,
					EndLine:   5,
					StartPos:  167,
					EndPos:    215,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  167,
						EndPos:    214,
					},
					Static:        true,
					PhpDocComment: "",
					ReturnsRef:    false,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  183,
								EndPos:    196,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &node.Nullable{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  183,
									EndPos:    186,
								},
								Expr: &name.Name{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  184,
										EndPos:    186,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  184,
												EndPos:    186,
											},
											Value: "bar",
										},
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  188,
									EndPos:    191,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  188,
										EndPos:    191,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  193,
									EndPos:    196,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  193,
										EndPos:    196,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  193,
												EndPos:    196,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  199,
								EndPos:    210,
							},
							Variadic: true,
							ByRef:    true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  199,
									EndPos:    201,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  199,
											EndPos:    201,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  207,
									EndPos:    210,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  207,
										EndPos:    210,
									},
									Value: "baz",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp5ParameterNode(t *testing.T) {
	src := `<? 
		function foo(bar $bar=null, baz &...$baz) {}
		class foo {public function foo(bar $bar=null, baz &...$baz) {}}
		function(bar $bar=null, baz &...$baz) {};
		static function(bar $bar=null, baz &...$baz) {};
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   5,
			StartPos:  7,
			EndPos:    211,
		},
		Stmts: []node.Node{
			&stmt.Function{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  7,
					EndPos:    50,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  16,
						EndPos:    18,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  20,
							EndPos:    32,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  20,
								EndPos:    22,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  20,
										EndPos:    22,
									},
									Value: "bar",
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  24,
								EndPos:    27,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  24,
									EndPos:    27,
								},
								Value: "bar",
							},
						},
						DefaultValue: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  29,
								EndPos:    32,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  29,
									EndPos:    32,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  29,
											EndPos:    32,
										},
										Value: "null",
									},
								},
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  35,
							EndPos:    46,
						},
						ByRef:    true,
						Variadic: true,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  35,
								EndPos:    37,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  35,
										EndPos:    37,
									},
									Value: "baz",
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  43,
								EndPos:    46,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  43,
									EndPos:    46,
								},
								Value: "baz",
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  54,
					EndPos:    116,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  60,
						EndPos:    62,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  65,
							EndPos:    115,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  81,
								EndPos:    83,
							},
							Value: "foo",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  65,
									EndPos:    70,
								},
								Value: "public",
							},
						},
						Params: []node.Node{
							&node.Parameter{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  85,
									EndPos:    97,
								},
								ByRef:    false,
								Variadic: false,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  85,
										EndPos:    87,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  85,
												EndPos:    87,
											},
											Value: "bar",
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  89,
										EndPos:    92,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  89,
											EndPos:    92,
										},
										Value: "bar",
									},
								},
								DefaultValue: &expr.ConstFetch{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  94,
										EndPos:    97,
									},
									Constant: &name.Name{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  94,
											EndPos:    97,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 3,
													EndLine:   3,
													StartPos:  94,
													EndPos:    97,
												},
												Value: "null",
											},
										},
									},
								},
							},
							&node.Parameter{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  100,
									EndPos:    111,
								},
								ByRef:    true,
								Variadic: true,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  100,
										EndPos:    102,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  100,
												EndPos:    102,
											},
											Value: "baz",
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  108,
										EndPos:    111,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  108,
											EndPos:    111,
										},
										Value: "baz",
									},
								},
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  114,
								EndPos:    115,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   4,
					StartPos:  120,
					EndPos:    160,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  120,
						EndPos:    159,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  129,
								EndPos:    141,
							},
							Variadic: false,
							ByRef:    false,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  129,
									EndPos:    131,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  129,
											EndPos:    131,
										},
										Value: "bar",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  133,
									EndPos:    136,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  133,
										EndPos:    136,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  138,
									EndPos:    141,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  138,
										EndPos:    141,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  138,
												EndPos:    141,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  144,
								EndPos:    155,
							},
							ByRef:    true,
							Variadic: true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  144,
									EndPos:    146,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  144,
											EndPos:    146,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  152,
									EndPos:    155,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  152,
										EndPos:    155,
									},
									Value: "baz",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 5,
					EndLine:   5,
					StartPos:  164,
					EndPos:    211,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  164,
						EndPos:    210,
					},
					PhpDocComment: "",
					ReturnsRef:    false,
					Static:        true,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  180,
								EndPos:    192,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  180,
									EndPos:    182,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  180,
											EndPos:    182,
										},
										Value: "bar",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  184,
									EndPos:    187,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  184,
										EndPos:    187,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  189,
									EndPos:    192,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  189,
										EndPos:    192,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  189,
												EndPos:    192,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  195,
								EndPos:    206,
							},
							ByRef:    true,
							Variadic: true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  195,
									EndPos:    197,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  195,
											EndPos:    197,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  203,
									EndPos:    206,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  203,
										EndPos:    206,
									},
									Value: "baz",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCommentEndFile(t *testing.T) {
	src := `<? //comment at the end)`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: -1,
			EndLine:   -1,
			StartPos:  -1,
			EndPos:    -1,
		},
		Stmts: []node.Node{},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
