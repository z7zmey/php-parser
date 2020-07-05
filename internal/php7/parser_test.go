package php7_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/internal/php7"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/position"
)

func TestIdentifier(t *testing.T) {
	src := `<? $foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Value: []byte("$foo"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
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

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   9,
				StartPos:  6,
				EndPos:    186,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    21,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  6,
							EndPos:    20,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  6,
								EndPos:    9,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  6,
										EndPos:    9,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  9,
								EndPos:    20,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  10,
										EndPos:    12,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  10,
											EndPos:    12,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  10,
												EndPos:    12,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  14,
										EndPos:    19,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  17,
											EndPos:    19,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  17,
												EndPos:    19,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  24,
						EndPos:    40,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  24,
							EndPos:    39,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  24,
								EndPos:    28,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  24,
									EndPos:    28,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  28,
								EndPos:    39,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  29,
										EndPos:    31,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  29,
											EndPos:    31,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  29,
												EndPos:    31,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  33,
										EndPos:    38,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  36,
											EndPos:    38,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  36,
												EndPos:    38,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  43,
						EndPos:    64,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  43,
							EndPos:    63,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  43,
								EndPos:    47,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  43,
									EndPos:    47,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  49,
								EndPos:    52,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  52,
								EndPos:    63,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  53,
										EndPos:    55,
									},
								},
								IsReference: false,
								Variadic:    false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  53,
											EndPos:    55,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  53,
												EndPos:    55,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  57,
										EndPos:    62,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  60,
											EndPos:    62,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  60,
												EndPos:    62,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  67,
						EndPos:    87,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  67,
							EndPos:    86,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  67,
								EndPos:    70,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  67,
										EndPos:    70,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  72,
								EndPos:    75,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  75,
								EndPos:    86,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  76,
										EndPos:    78,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  76,
											EndPos:    78,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  76,
												EndPos:    78,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  80,
										EndPos:    85,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  83,
											EndPos:    85,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  83,
												EndPos:    85,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  90,
						EndPos:    111,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  90,
							EndPos:    110,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  90,
								EndPos:    94,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  90,
									EndPos:    94,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  96,
								EndPos:    99,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  99,
								EndPos:    110,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  100,
										EndPos:    102,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  100,
											EndPos:    102,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 6,
												EndLine:   6,
												StartPos:  100,
												EndPos:    102,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  104,
										EndPos:    109,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  107,
											EndPos:    109,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 6,
												EndLine:   6,
												StartPos:  107,
												EndPos:    109,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  114,
						EndPos:    133,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  114,
							EndPos:    132,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 7,
								EndLine:   7,
								StartPos:  118,
								EndPos:    121,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  118,
										EndPos:    121,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 7,
								EndLine:   7,
								StartPos:  121,
								EndPos:    132,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  122,
										EndPos:    124,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  122,
											EndPos:    124,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 7,
												EndLine:   7,
												StartPos:  122,
												EndPos:    124,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  126,
										EndPos:    131,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  129,
											EndPos:    131,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 7,
												EndLine:   7,
												StartPos:  129,
												EndPos:    131,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  161,
						EndPos:    186,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  161,
							EndPos:    185,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  165,
								EndPos:    185,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  171,
									EndPos:    182,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  172,
											EndPos:    174,
										},
									},
									Variadic:    false,
									IsReference: false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  172,
												EndPos:    174,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 9,
													EndLine:   9,
													StartPos:  172,
													EndPos:    174,
												},
											},
											Value: []byte("$a"),
										},
									},
								},
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  176,
											EndPos:    181,
										},
									},
									Variadic:    true,
									IsReference: false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  179,
												EndPos:    181,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 9,
													EndLine:   9,
													StartPos:  179,
													EndPos:    181,
												},
											},
											Value: []byte("$b"),
										},
									},
								},
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp7ParameterNode(t *testing.T) {
	src := `<?
		function foo(?bar $bar=null, baz &...$baz) {}
		class foo {public function foo(?bar $bar=null, baz &...$baz) {}}
		function(?bar $bar=null, baz &...$baz) {};
		static function(?bar $bar=null, baz &...$baz) {};
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   5,
				StartPos:  5,
				EndPos:    214,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  5,
						EndPos:    50,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  14,
							EndPos:    17,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  18,
								EndPos:    32,
							},
						},
						Type: &ast.Nullable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  18,
									EndPos:    22,
								},
							},
							Expr: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("bar"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  23,
									EndPos:    27,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  23,
										EndPos:    27,
									},
								},
								Value: []byte("$bar"),
							},
						},
						DefaultValue: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  28,
									EndPos:    32,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  28,
										EndPos:    32,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  28,
												EndPos:    32,
											},
										},
										Value: []byte("null"),
									},
								},
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  34,
								EndPos:    46,
							},
						},
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  34,
									EndPos:    37,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  34,
											EndPos:    37,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
						Var: &ast.Reference{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  38,
									EndPos:    46,
								},
							},
							Var: &ast.Variadic{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  39,
										EndPos:    46,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  42,
											EndPos:    46,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  42,
												EndPos:    46,
											},
										},
										Value: []byte("$baz"),
									},
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  53,
						EndPos:    117,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  59,
							EndPos:    62,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  64,
								EndPos:    116,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  80,
									EndPos:    83,
								},
							},
							Value: []byte("foo"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  64,
										EndPos:    70,
									},
								},
								Value: []byte("public"),
							},
						},
						Params: []ast.Vertex{
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  84,
										EndPos:    98,
									},
								},
								Type: &ast.Nullable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  84,
											EndPos:    88,
										},
									},
									Expr: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  85,
												EndPos:    88,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 3,
														EndLine:   3,
														StartPos:  85,
														EndPos:    88,
													},
												},
												Value: []byte("bar"),
											},
										},
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  89,
											EndPos:    93,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  89,
												EndPos:    93,
											},
										},
										Value: []byte("$bar"),
									},
								},
								DefaultValue: &ast.ExprConstFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  94,
											EndPos:    98,
										},
									},
									Const: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  94,
												EndPos:    98,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 3,
														EndLine:   3,
														StartPos:  94,
														EndPos:    98,
													},
												},
												Value: []byte("null"),
											},
										},
									},
								},
							},
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  100,
										EndPos:    112,
									},
								},
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  100,
											EndPos:    103,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 3,
													EndLine:   3,
													StartPos:  100,
													EndPos:    103,
												},
											},
											Value: []byte("baz"),
										},
									},
								},
								Var: &ast.Reference{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  104,
											EndPos:    112,
										},
									},
									Var: &ast.Variadic{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  105,
												EndPos:    112,
											},
										},
										Var: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 3,
													EndLine:   3,
													StartPos:  108,
													EndPos:    112,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 3,
														EndLine:   3,
														StartPos:  108,
														EndPos:    112,
													},
												},
												Value: []byte("$baz"),
											},
										},
									},
								},
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  114,
									EndPos:    116,
								},
							},
							Stmts: []ast.Vertex{
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  120,
						EndPos:    162,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  120,
							EndPos:    161,
						},
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  129,
									EndPos:    143,
								},
							},
							Type: &ast.Nullable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  129,
										EndPos:    133,
									},
								},
								Expr: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  130,
											EndPos:    133,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 4,
													EndLine:   4,
													StartPos:  130,
													EndPos:    133,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  134,
										EndPos:    138,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  134,
											EndPos:    138,
										},
									},
									Value: []byte("$bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  139,
										EndPos:    143,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  139,
											EndPos:    143,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 4,
													EndLine:   4,
													StartPos:  139,
													EndPos:    143,
												},
											},
											Value: []byte("null"),
										},
									},
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  145,
									EndPos:    157,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  145,
										EndPos:    148,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  145,
												EndPos:    148,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.Reference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  149,
										EndPos:    157,
									},
								},
								Var: &ast.Variadic{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  150,
											EndPos:    157,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  153,
												EndPos:    157,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 4,
													EndLine:   4,
													StartPos:  153,
													EndPos:    157,
												},
											},
											Value: []byte("$baz"),
										},
									},
								},
							},
						},
					},
					Stmts: []ast.Vertex{
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  165,
						EndPos:    214,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  165,
							EndPos:    213,
						},
					},
					Static: true,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  181,
									EndPos:    195,
								},
							},
							Type: &ast.Nullable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  181,
										EndPos:    185,
									},
								},
								Expr: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  182,
											EndPos:    185,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 5,
													EndLine:   5,
													StartPos:  182,
													EndPos:    185,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  186,
										EndPos:    190,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  186,
											EndPos:    190,
										},
									},
									Value: []byte("$bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  191,
										EndPos:    195,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  191,
											EndPos:    195,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 5,
													EndLine:   5,
													StartPos:  191,
													EndPos:    195,
												},
											},
											Value: []byte("null"),
										},
									},
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  197,
									EndPos:    209,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  197,
										EndPos:    200,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  197,
												EndPos:    200,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.Reference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  201,
										EndPos:    209,
									},
								},
								Var: &ast.Variadic{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  202,
											EndPos:    209,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  205,
												EndPos:    209,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 5,
													EndLine:   5,
													StartPos:  205,
													EndPos:    209,
												},
											},
											Value: []byte("$baz"),
										},
									},
								},
							},
						},
					},
					Stmts: []ast.Vertex{
					},
				},
			},
		},
	}


	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCommentEndFile(t *testing.T) {
	src := `<? //comment at the end)`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: -1,
				EndLine:   -1,
				StartPos:  -1,
				EndPos:    -1,
			},
		},
		Stmts: []ast.Vertex{},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// name

func TestName(t *testing.T) {
	src := `<? foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    9,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    9,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    8,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFullyQualified(t *testing.T) {
	src := `<? \foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    10,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    10,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    9,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    7,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestRelative(t *testing.T) {
	src := `<? namespace\foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Function: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    16,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    16,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// scalar

func TestScalarEncapsed_SimpleVar(t *testing.T) {
	src := `<? "test $var";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    13,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    13,
									},
								},
								Value: []byte("$var"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_SimpleVarOneChar(t *testing.T) {
	src := `<? "test $a";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    11,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_SimpleVarEndsEcapsed(t *testing.T) {
	src := `<? "test $var\"";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    13,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    13,
									},
								},
								Value: []byte("$var"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							Value: []byte("\\\""),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_StringVarCurveOpen(t *testing.T) {
	src := `<? "=$a{$b}";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    5,
								},
							},
							Value: []byte("="),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    7,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  5,
										EndPos:    7,
									},
								},
								Value: []byte("$a"),
							},
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_SimpleVarPropertyFetch(t *testing.T) {
	src := `<? "test $foo->bar()";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    22,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    22,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    21,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    18,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    13,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  9,
											EndPos:    13,
										},
									},
									Value: []byte("$foo"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  15,
										EndPos:    18,
									},
								},
								Value: []byte("bar"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
							},
							Value: []byte("()"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_DollarOpenCurlyBraces(t *testing.T) {
	src := `<? "test ${foo}";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    15,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  11,
										EndPos:    14,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_DollarOpenCurlyBracesDimNumber(t *testing.T) {
	src := `<? "test ${foo[0]}";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    20,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    20,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    19,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    18,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  11,
										EndPos:    14,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  11,
											EndPos:    14,
										},
									},
									Value: []byte("foo"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  15,
										EndPos:    16,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarEncapsed_CurlyOpenMethodCall(t *testing.T) {
	src := `<? "test {$foo->bar()}";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    23,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    9,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    21,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    14,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  10,
											EndPos:    14,
										},
									},
									Value: []byte("$foo"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  16,
										EndPos:    19,
									},
								},
								Value: []byte("bar"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    21,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarHeredoc_HeredocSimpleLabel(t *testing.T) {
	src := `<? <<<LBL
test $var
LBL;
`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   3,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   3,
							StartPos:  3,
							EndPos:    23,
						},
					},
					Label: []byte("<<<LBL\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  10,
									EndPos:    15,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  15,
									EndPos:    19,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  15,
										EndPos:    19,
									},
								},
								Value: []byte("$var"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  19,
									EndPos:    20,
								},
							},
							Value: []byte("\n"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarHeredoc_SimpleHeredocLabel(t *testing.T) {
	src := `<? <<<"LBL"
test $var
LBL;
`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   3,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    26,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   3,
							StartPos:  3,
							EndPos:    25,
						},
					},
					Label: []byte("<<<\"LBL\"\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  12,
									EndPos:    17,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  17,
									EndPos:    21,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  17,
										EndPos:    21,
									},
								},
								Value: []byte("$var"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  21,
									EndPos:    22,
								},
							},
							Value: []byte("\n"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarHeredoc_SimpleNowdocLabel(t *testing.T) {
	src := `<? <<<'LBL'
test $var
LBL;
`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   3,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    26,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   3,
							StartPos:  3,
							EndPos:    25,
						},
					},
					Label: []byte("<<<'LBL'\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  12,
									EndPos:    22,
								},
							},
							Value: []byte("test $var\n"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarHeredoc_EmptyHeredoc(t *testing.T) {
	src := `<? <<<CAD
CAD;
`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   2,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   2,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   2,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Label: []byte("<<<CAD\n"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarHeredoc_HeredocScalarString(t *testing.T) {
	src := `<? <<<CAD
	hello
CAD;
`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   3,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   3,
							StartPos:  3,
							EndPos:    20,
						},
					},
					Label: []byte("<<<CAD\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  10,
									EndPos:    17,
								},
							},
							Value: []byte("\thello\n"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarMagicConstant(t *testing.T) {
	// TODO: test all magic constants
	src := `<? __DIR__;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Value: []byte("__DIR__"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_LNumber(t *testing.T) {
	src := `<? 1234567890123456789;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					Value: []byte("1234567890123456789"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_DNumber(t *testing.T) {
	src := `<? 12345678901234567890;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    23,
						},
					},
					Value: []byte("12345678901234567890"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_Float(t *testing.T) {
	src := `<? 0.;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    6,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    6,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
					},
					Value: []byte("0."),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_BinaryLNumber(t *testing.T) {
	src := `<? 0b0111111111111111111111111111111111111111111111111111111111111111;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    70,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    70,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    69,
						},
					},
					Value: []byte("0b0111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_BinaryDNumber(t *testing.T) {
	src := `<? 0b1111111111111111111111111111111111111111111111111111111111111111;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    70,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    70,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    69,
						},
					},
					Value: []byte("0b1111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_HLNumber(t *testing.T) {
	src := `<? 0x007111111111111111;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    23,
						},
					},
					Value: []byte("0x007111111111111111"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarNumber_HDNumber(t *testing.T) {
	src := `<? 0x8111111111111111;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    22,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    22,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    21,
						},
					},
					Value: []byte("0x8111111111111111"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarString_DoubleQuotedScalarString(t *testing.T) {
	src := `<? "test";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    10,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    10,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    9,
						},
					},
					Value: []byte("\"test\""),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarString_DoubleQuotedScalarStringWithEscapedVar(t *testing.T) {
	src := `<? "\$test";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Value: []byte("\"\\$test\""),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarString_MultilineDoubleQuotedScalarString(t *testing.T) {
	src := `<? "
	test
	";`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   3,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   3,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Value: []byte("\"\n\ttest\n\t\""),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarString_SingleQuotedScalarString(t *testing.T) {
	src := `<? '$test';`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Value: []byte("'$test'"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestScalarString_MultilineSingleQuotedScalarString(t *testing.T) {
	src := `<? '
	$test
	';`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   3,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   3,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Value: []byte("'\n\t$test\n\t'"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// stmt

func TestStmtAltIf_AltIf(t *testing.T) {
	src := `<?
		if ($a) :
		endif;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   3,
				StartPos:  5,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   3,
						StartPos:  5,
						EndPos:    23,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  9,
							EndPos:    11,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  9,
								EndPos:    11,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtAltIf_AltElseIf(t *testing.T) {
	src := `<?
		if ($a) :
		elseif ($b):
		endif;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   4,
				StartPos:  5,
				EndPos:    38,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   4,
						StartPos:  5,
						EndPos:    38,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  9,
							EndPos:    11,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  9,
								EndPos:    11,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtAltElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   -1,
								StartPos:  17,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  25,
									EndPos:    27,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  25,
										EndPos:    27,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: -1,
									EndLine:   -1,
									StartPos:  -1,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtAltIf_AltElse(t *testing.T) {
	src := `<?
		if ($a) :
		else:
		endif;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   4,
				StartPos:  5,
				EndPos:    31,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   4,
						StartPos:  5,
						EndPos:    31,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  9,
							EndPos:    11,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  9,
								EndPos:    11,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtAltElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   -1,
							StartPos:  17,
							EndPos:    -1,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtAltIf_AltElseElseIf(t *testing.T) {
	src := `<?
		if ($a) :
		elseif ($b):
		elseif ($c):
		else:
		endif;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   6,
				StartPos:  5,
				EndPos:    61,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   6,
						StartPos:  5,
						EndPos:    61,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  9,
							EndPos:    11,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  9,
								EndPos:    11,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtAltElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   -1,
								StartPos:  17,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  25,
									EndPos:    27,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  25,
										EndPos:    27,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: -1,
									EndLine:   -1,
									StartPos:  -1,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtAltElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   -1,
								StartPos:  32,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  40,
									EndPos:    42,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  40,
										EndPos:    42,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: -1,
									EndLine:   -1,
									StartPos:  -1,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtAltElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   -1,
							StartPos:  47,
							EndPos:    -1,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassConstList(t *testing.T) {
	src := `<? class foo{ public const FOO = 1, BAR = 2; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    46,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    46,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    44,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    20,
									},
								},
								Value: []byte("public"),
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  27,
										EndPos:    34,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  27,
											EndPos:    30,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  33,
											EndPos:    34,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  36,
										EndPos:    43,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  36,
											EndPos:    39,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  42,
											EndPos:    43,
										},
									},
									Value: []byte("2"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassConstList_WithoutModifiers(t *testing.T) {
	src := `<? class foo{ const FOO = 1, BAR = 2; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    39,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    39,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    37,
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  20,
										EndPos:    27,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  20,
											EndPos:    23,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  26,
											EndPos:    27,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  29,
										EndPos:    36,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  29,
											EndPos:    32,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  35,
											EndPos:    36,
										},
									},
									Value: []byte("2"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassMethod_SimpleClassMethod(t *testing.T) {
	src := `<? class foo{ function bar() {} }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    33,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    33,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    31,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    26,
								},
							},
							Value: []byte("bar"),
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  29,
									EndPos:    31,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassMethod_PrivateProtectedClassMethod(t *testing.T) {
	src := `<? class foo{ final private function bar() {} protected function baz() {} }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    75,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    75,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    45,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  37,
									EndPos:    40,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    19,
									},
								},
								Value: []byte("final"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  20,
										EndPos:    27,
									},
								},
								Value: []byte("private"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  43,
									EndPos:    45,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  46,
								EndPos:    73,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  65,
									EndPos:    68,
								},
							},
							Value: []byte("baz"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  46,
										EndPos:    55,
									},
								},
								Value: []byte("protected"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  71,
									EndPos:    73,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassMethod_Php7ClassMethod(t *testing.T) {
	src := `<? class foo{ public static function &bar(): void {} }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    54,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    54,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    52,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  38,
									EndPos:    41,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    20,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  21,
										EndPos:    27,
									},
								},
								Value: []byte("static"),
							},
						},
						ReturnType: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  45,
									EndPos:    49,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  45,
											EndPos:    49,
										},
									},
									Value: []byte("void"),
								},
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  50,
									EndPos:    52,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassMethod_AbstractClassMethod(t *testing.T) {
	src := `<? abstract class Foo{ abstract public function bar(); }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    56,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    56,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    21,
						},
					},
					Value: []byte("Foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    11,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    54,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  48,
									EndPos:    51,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  23,
										EndPos:    31,
									},
								},
								Value: []byte("abstract"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  32,
										EndPos:    38,
									},
								},
								Value: []byte("public"),
							},
						},
						Stmt: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  53,
									EndPos:    54,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClassMethod_Php7AbstractClassMethod(t *testing.T) {
	src := `<? abstract class Foo{ public function bar(): void; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    53,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    53,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    21,
						},
					},
					Value: []byte("Foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    11,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    51,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  39,
									EndPos:    42,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  23,
										EndPos:    29,
									},
								},
								Value: []byte("public"),
							},
						},
						ReturnType: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  46,
									EndPos:    50,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  46,
											EndPos:    50,
										},
									},
									Value: []byte("void"),
								},
							},
						},
						Stmt: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  50,
									EndPos:    51,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClass_SimpleClass(t *testing.T) {
	src := `<? class foo{ }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClass_AbstractClass(t *testing.T) {
	src := `<? abstract class foo{ }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    21,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    11,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClass_ClassExtends(t *testing.T) {
	src := `<? final class foo extends bar { }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    34,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    34,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    18,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    8,
							},
						},
						Value: []byte("final"),
					},
				},
				Extends: &ast.StmtClassExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    30,
						},
					},
					ClassName: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  27,
								EndPos:    30,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  27,
										EndPos:    30,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClass_ClassImplement(t *testing.T) {
	src := `<? final class foo implements bar { }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    37,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    37,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    18,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    8,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    33,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  30,
									EndPos:    33,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    33,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClass_ClassImplements(t *testing.T) {
	src := `<? final class foo implements bar, baz { }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    42,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    42,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    18,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    8,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    38,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  30,
									EndPos:    33,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    33,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  35,
									EndPos:    38,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  35,
											EndPos:    38,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtClass_AnonimousClass(t *testing.T) {
	src := `<? new class() extends foo implements bar, baz { };`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    51,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    51,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    50,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    50,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
						},
						Extends: &ast.StmtClassExtends{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    26,
								},
							},
							ClassName: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  23,
										EndPos:    26,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  23,
												EndPos:    26,
											},
										},
										Value: []byte("foo"),
									},
								},
							},
						},
						Implements: &ast.StmtClassImplements{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  27,
									EndPos:    46,
								},
							},
							InterfaceNames: []ast.Vertex{
								&ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  38,
											EndPos:    41,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  38,
													EndPos:    41,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
								&ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  43,
											EndPos:    46,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  43,
													EndPos:    46,
												},
											},
											Value: []byte("baz"),
										},
									},
								},
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtConstList(t *testing.T) {
	src := `<? const FOO = 1, BAR = 2;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtConstList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    26,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    16,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    12,
								},
							},
							Value: []byte("FOO"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    16,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    25,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    21,
								},
							},
							Value: []byte("BAR"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    25,
								},
							},
							Value: []byte("2"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtContinue_Empty(t *testing.T) {
	src := `<? while (1) { continue; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    26,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    11,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    26,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    24,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtContinue_Light(t *testing.T) {
	src := `<? while (1) { continue 2; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    28,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    28,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    11,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    28,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    26,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    25,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtContinue(t *testing.T) {
	src := `<? while (1) { continue(3); }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    29,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    29,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    11,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    29,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    27,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    25,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtDeclare(t *testing.T) {
	src := `<? declare(ticks=1);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    20,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    20,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    18,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    16,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    18,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtNop{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    20,
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtDeclare_Stmts(t *testing.T) {
	src := `<? declare(ticks=1, strict_types=1) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    38,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    38,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    18,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    16,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    18,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    34,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    32,
								},
							},
							Value: []byte("strict_types"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  33,
									EndPos:    34,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  36,
							EndPos:    38,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtDeclare_Alt(t *testing.T) {
	src := `<? declare(ticks=1): enddeclare;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    32,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    32,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    18,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    16,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    18,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Alt: true,
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtDo(t *testing.T) {
	src := `<? do {} while(1);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    18,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtDo{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    18,
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    8,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    16,
						},
					},
					Value: []byte("1"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtEcho(t *testing.T) {
	src := `<? echo $a, 1;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    13,
							},
						},
						Value: []byte("1"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtEcho_Parenthesis(t *testing.T) {
	src := `<? echo($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtExpression(t *testing.T) {
	src := `<? 1;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    5,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    5,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    4,
						},
					},
					Value: []byte("1"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFor(t *testing.T) {
	src := `<? for($i = 0; $i < 10; $i++, $i++) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    38,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    38,
					},
				},
				Init: []ast.Vertex{
					&ast.ExprAssign{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    13,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    9,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  7,
										EndPos:    9,
									},
								},
								Value: []byte("$i"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    13,
								},
							},
							Value: []byte("0"),
						},
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    22,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    17,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  15,
										EndPos:    17,
									},
								},
								Value: []byte("$i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    22,
								},
							},
							Value: []byte("10"),
						},
					},
				},
				Loop: []ast.Vertex{
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    28,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    26,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    26,
									},
								},
								Value: []byte("$i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  30,
								EndPos:    34,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  30,
									EndPos:    32,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  30,
										EndPos:    32,
									},
								},
								Value: []byte("$i"),
							},
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  36,
							EndPos:    38,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFor_Alt(t *testing.T) {
	src := `<? for(; $i < 10; $i++) : endfor;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    33,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    33,
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    16,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    11,
									},
								},
								Value: []byte("$i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    16,
								},
							},
							Value: []byte("10"),
						},
					},
				},
				Loop: []ast.Vertex{
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    22,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
								},
								Value: []byte("$i"),
							},
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach(t *testing.T) {
	src := `<? foreach ($a as $v) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Value: []byte("$a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  22,
							EndPos:    24,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach_Expr(t *testing.T) {
	src := `<? foreach ([] as $v) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					Items: []ast.Vertex{},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  22,
							EndPos:    24,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach_Alt(t *testing.T) {
	src := `<? foreach ($a as $v) : endforeach;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    35,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    35,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Value: []byte("$a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach_WithKey(t *testing.T) {
	src := `<? foreach ($a as $k => $v) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    30,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    30,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  24,
							EndPos:    26,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    26,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  28,
							EndPos:    30,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach_ExprWithKey(t *testing.T) {
	src := `<? foreach ([] as $k => $v) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    30,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    30,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					Items: []ast.Vertex{},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  24,
							EndPos:    26,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    26,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  28,
							EndPos:    30,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach_WithRef(t *testing.T) {
	src := `<? foreach ($a as $k => &$v) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    31,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    31,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  24,
							EndPos:    27,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  25,
								EndPos:    27,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  25,
									EndPos:    27,
								},
							},
							Value: []byte("$v"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  29,
							EndPos:    31,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtForeach_WithList(t *testing.T) {
	src := `<? foreach ($a as $k => list($v)) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    36,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    36,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  24,
							EndPos:    32,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  29,
									EndPos:    31,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  29,
										EndPos:    31,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  29,
											EndPos:    31,
										},
									},
									Value: []byte("$v"),
								},
							},
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  34,
							EndPos:    36,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFunction(t *testing.T) {
	src := `<? function foo() {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    20,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    20,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    15,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFunction_Return(t *testing.T) {
	src := `<? function foo() {return;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    27,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    27,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    15,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    26,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFunction_ReturnVar(t *testing.T) {
	src := `<? function foo(array $a, callable $b) {return $a;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    51,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    51,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    15,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    24,
							},
						},
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    21,
								},
							},
							Value: []byte("array"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  26,
								EndPos:    37,
							},
						},
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  26,
									EndPos:    34,
								},
							},
							Value: []byte("callable"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  35,
									EndPos:    37,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  35,
										EndPos:    37,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  40,
								EndPos:    50,
							},
						},
						Expr: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  47,
									EndPos:    49,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  47,
										EndPos:    49,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFunction_Ref(t *testing.T) {
	src := `<? function &foo() {return 1;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    30,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    30,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    29,
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  27,
									EndPos:    28,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtFunction_ReturnType(t *testing.T) {
	src := `<? function &foo(): void {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    27,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    27,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Value: []byte("foo"),
				},
				ReturnType: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  20,
							EndPos:    24,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    24,
								},
							},
							Value: []byte("void"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtGlobal(t *testing.T) {
	src := `<? global $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGlobal{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtGlobal_Vars(t *testing.T) {
	src := `<? global $a, $b, $$c, ${foo()};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    32,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGlobal{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    32,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    16,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    16,
								},
							},
							Value: []byte("$b"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    21,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    21,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    21,
									},
								},
								Value: []byte("$c"),
							},
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    31,
							},
						},
						VarName: &ast.ExprFunctionCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  25,
									EndPos:    30,
								},
							},
							Function: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  25,
										EndPos:    28,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  25,
												EndPos:    28,
											},
										},
										Value: []byte("foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  28,
										EndPos:    30,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtGotoLabel(t *testing.T) {
	src := `<? a: goto a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtLabel{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    5,
					},
				},
				LabelName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    4,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtGoto{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    13,
					},
				},
				Label: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    12,
						},
					},
					Value: []byte("a"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtHaltCompiler(t *testing.T) {
	src := `<? __halt_compiler();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtHaltCompiler{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtIf(t *testing.T) {
	src := `<? if ($a) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    13,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtIf_ElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    28,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    28,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    13,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    28,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  26,
									EndPos:    28,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtIf_Else(t *testing.T) {
	src := `<? if ($a) {} else {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    13,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    21,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    21,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtIf_ElseElseIf(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} elseif ($c) {} else {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    51,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    51,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    13,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    28,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  26,
									EndPos:    28,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  29,
								EndPos:    43,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  37,
									EndPos:    39,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  37,
										EndPos:    39,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  41,
									EndPos:    43,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  44,
							EndPos:    51,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  49,
								EndPos:    51,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtIf_ElseIfElseIfElse(t *testing.T) {
	src := `<? if ($a) {} elseif ($b) {} else if ($c) {} else {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    52,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    52,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    9,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    13,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    28,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  26,
									EndPos:    28,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  29,
							EndPos:    52,
						},
					},
					Stmt: &ast.StmtIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  34,
								EndPos:    52,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  38,
									EndPos:    40,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  38,
										EndPos:    40,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  42,
									EndPos:    44,
								},
							},
							Stmts: []ast.Vertex{},
						},
						Else: &ast.StmtElse{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  45,
									EndPos:    52,
								},
							},
							Stmt: &ast.StmtStmtList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  50,
										EndPos:    52,
									},
								},
								Stmts: []ast.Vertex{},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtInlineHtml(t *testing.T) {
	src := `<? ?> <div></div>`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    5,
					},
				},
			},
			&ast.StmtInlineHtml{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  5,
						EndPos:    17,
					},
				},
				Value: []byte(" <div></div>"),
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtInterface(t *testing.T) {
	src := `<? interface Foo {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtInterface_Extend(t *testing.T) {
	src := `<? interface Foo extends Bar {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    31,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    31,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  17,
							EndPos:    28,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  25,
									EndPos:    28,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  25,
											EndPos:    28,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtInterface_Extends(t *testing.T) {
	src := `<? interface Foo extends Bar, Baz {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    36,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    36,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  17,
							EndPos:    33,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  25,
									EndPos:    28,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  25,
											EndPos:    28,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  30,
									EndPos:    33,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    33,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtNamespace(t *testing.T) {
	src := `<? namespace Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtNamespace_Stmts(t *testing.T) {
	src := `<? namespace Foo {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtNamespace_Anonymous(t *testing.T) {
	src := `<? namespace {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtProperty(t *testing.T) {
	src := `<? class foo {var $a;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    22,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    22,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    21,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    17,
									},
								},
								Value: []byte("var"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    20,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  18,
												EndPos:    20,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtProperty_Properties(t *testing.T) {
	src := `<? class foo {public static $a, $b = 1;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    40,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    40,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    39,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    20,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  21,
										EndPos:    27,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  28,
										EndPos:    30,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  28,
											EndPos:    30,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  28,
												EndPos:    30,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  32,
										EndPos:    38,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  32,
											EndPos:    34,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  32,
												EndPos:    34,
											},
										},
										Value: []byte("$b"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  37,
											EndPos:    38,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtProperty_Properties2(t *testing.T) {
	src := `<? class foo {public static $a = 1, $b;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    40,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    40,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    39,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    20,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  21,
										EndPos:    27,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  28,
										EndPos:    34,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  28,
											EndPos:    30,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  28,
												EndPos:    30,
											},
										},
										Value: []byte("$a"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  33,
											EndPos:    34,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  36,
										EndPos:    38,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  36,
											EndPos:    38,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  36,
												EndPos:    38,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtProperty_PropertyType(t *testing.T) {
	src := `<? class foo {var bar $a;}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    26,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    25,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    17,
									},
								},
								Value: []byte("var"),
							},
						},
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    21,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    21,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  22,
											EndPos:    24,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  22,
												EndPos:    24,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtStaticVar(t *testing.T) {
	src := `<? static $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    12,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtStaticVar_Vars(t *testing.T) {
	src := `<? static $a, $b = 1;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    12,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    20,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    16,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    16,
									},
								},
								Value: []byte("$b"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    20,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtStaticVar_Vars2(t *testing.T) {
	src := `<? static $a = 1, $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    16,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    12,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    16,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtSwitch(t *testing.T) {
	src := `<? 
		switch (1) {
			case 1: break;
			case 2: break;
		}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   5,
				StartPos:  6,
				EndPos:    58,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   5,
						StartPos:  6,
						EndPos:    58,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  14,
							EndPos:    15,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   5,
							StartPos:  17,
							EndPos:    58,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  22,
									EndPos:    36,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  27,
										EndPos:    28,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  30,
											EndPos:    36,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  40,
									EndPos:    54,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  45,
										EndPos:    46,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  48,
											EndPos:    54,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtSwitch_Semicolon(t *testing.T) {
	src := `<? 
		switch (1) {;
			case 1; break;
			case 2; break;
		}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   5,
				StartPos:  6,
				EndPos:    59,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   5,
						StartPos:  6,
						EndPos:    59,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  14,
							EndPos:    15,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   5,
							StartPos:  17,
							EndPos:    59,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  23,
									EndPos:    37,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  28,
										EndPos:    29,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  31,
											EndPos:    37,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  41,
									EndPos:    55,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  46,
										EndPos:    47,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  49,
											EndPos:    55,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtSwitch_Alt(t *testing.T) {
	src := `<? 
		switch (1) :
			case 1:
			default:
			case 2;
		endswitch;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   6,
				StartPos:  6,
				EndPos:    65,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   6,
						StartPos:  6,
						EndPos:    65,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  14,
							EndPos:    15,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   -1,
							StartPos:  22,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   -1,
									StartPos:  22,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  27,
										EndPos:    28,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtDefault{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   -1,
									StartPos:  33,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   -1,
									StartPos:  45,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  50,
										EndPos:    51,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtSwitch_AltSemicolon(t *testing.T) {
	src := `<? 
		switch (1) :;
			case 1;
			case 2;
		endswitch;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   5,
				StartPos:  6,
				EndPos:    54,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   5,
						StartPos:  6,
						EndPos:    54,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  14,
							EndPos:    15,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   -1,
							StartPos:  23,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   -1,
									StartPos:  23,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  28,
										EndPos:    29,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   -1,
									StartPos:  34,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  39,
										EndPos:    40,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtThrow(t *testing.T) {
	src := `<? throw $e;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtThrow{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    11,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						Value: []byte("$e"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTrait(t *testing.T) {
	src := `<? trait Foo {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTrait{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				TraitName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTraitUse(t *testing.T) {
	src := `<? class Foo { use Bar; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    25,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    25,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    23,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    23,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTraitUse_Uses(t *testing.T) {
	src := `<? class Foo { use Bar, Baz; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    30,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    30,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    28,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    27,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  24,
												EndPos:    27,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  27,
									EndPos:    28,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTraitUse_EmptyAdaptations(t *testing.T) {
	src := `<? class Foo { use Bar, Baz {} }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    32,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    32,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    30,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    27,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  24,
												EndPos:    27,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    30,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTraitUse_Modifier(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { one as public; } }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    48,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    48,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    46,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    27,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  24,
												EndPos:    27,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    46,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    43,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    33,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  30,
													EndPos:    33,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  37,
												EndPos:    43,
											},
										},
										Value: []byte("public"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTraitUse_AliasModifier(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { one as public two; } }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    52,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    52,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    50,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    27,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  24,
												EndPos:    27,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    50,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    47,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    33,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  30,
													EndPos:    33,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  37,
												EndPos:    43,
											},
										},
										Value: []byte("public"),
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  44,
												EndPos:    47,
											},
										},
										Value: []byte("two"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTraitUse_Adaptions(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { Bar::one insteadof Baz, Quux; Baz::one as two; } }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    80,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    80,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    12,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    78,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    22,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  19,
												EndPos:    22,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    27,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  24,
												EndPos:    27,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    78,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUsePrecedence{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    58,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    38,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  30,
													EndPos:    33,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 1,
															EndLine:   1,
															StartPos:  30,
															EndPos:    33,
														},
													},
													Value: []byte("Bar"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  35,
													EndPos:    38,
												},
											},
											Value: []byte("one"),
										},
									},
									Insteadof: []ast.Vertex{
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  49,
													EndPos:    52,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 1,
															EndLine:   1,
															StartPos:  49,
															EndPos:    52,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  54,
													EndPos:    58,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 1,
															EndLine:   1,
															StartPos:  54,
															EndPos:    58,
														},
													},
													Value: []byte("Quux"),
												},
											},
										},
									},
								},
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  60,
											EndPos:    75,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  60,
												EndPos:    68,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  60,
													EndPos:    63,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 1,
															EndLine:   1,
															StartPos:  60,
															EndPos:    63,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  65,
													EndPos:    68,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  72,
												EndPos:    75,
											},
										},
										Value: []byte("two"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTry_Try(t *testing.T) {
	src := `<? 
		try {}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   -1,
				StartPos:  6,
				EndPos:    -1,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   -1,
						StartPos:  6,
						EndPos:    -1,
					},
				},
				Stmts:   []ast.Vertex{},
				Catches: []ast.Vertex{},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTry_TryCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  6,
				EndPos:    36,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    36,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    36,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  20,
										EndPos:    29,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  20,
												EndPos:    29,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    32,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  30,
										EndPos:    32,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTry_Php7TryCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception|RuntimeException $e) {}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  6,
				EndPos:    53,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    53,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    53,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  20,
										EndPos:    29,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  20,
												EndPos:    29,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  30,
										EndPos:    46,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  30,
												EndPos:    46,
											},
										},
										Value: []byte("RuntimeException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  47,
									EndPos:    49,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  47,
										EndPos:    49,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTry_TryCatchCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {} catch (RuntimeException $e) {}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  6,
				EndPos:    67,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    67,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    36,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  20,
										EndPos:    29,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  20,
												EndPos:    29,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    32,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  30,
										EndPos:    32,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  37,
								EndPos:    67,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  44,
										EndPos:    60,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  44,
												EndPos:    60,
											},
										},
										Value: []byte("RuntimeException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  61,
									EndPos:    63,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  61,
										EndPos:    63,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTry_TryCatchFinally(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {} finally {}
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  6,
				EndPos:    47,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    47,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    36,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  20,
										EndPos:    29,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  20,
												EndPos:    29,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    32,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  30,
										EndPos:    32,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
				Finally: &ast.StmtFinally{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  37,
							EndPos:    47,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtTry_TryCatchCatchCatch(t *testing.T) {
	src := `<? try {} catch (Exception $e) {} catch (\RuntimeException $e) {} catch (namespace\AdditionException $e) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    107,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    107,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    33,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    26,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  17,
												EndPos:    26,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  27,
									EndPos:    29,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  27,
										EndPos:    29,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  34,
								EndPos:    65,
							},
						},
						Types: []ast.Vertex{
							&ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  41,
										EndPos:    58,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  42,
												EndPos:    58,
											},
										},
										Value: []byte("RuntimeException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  59,
									EndPos:    61,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  59,
										EndPos:    61,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  66,
								EndPos:    107,
							},
						},
						Types: []ast.Vertex{
							&ast.NameRelative{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  73,
										EndPos:    100,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  83,
												EndPos:    100,
											},
										},
										Value: []byte("AdditionException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  101,
									EndPos:    103,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  101,
										EndPos:    103,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUnset(t *testing.T) {
	src := `<? unset($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUnset_Vars(t *testing.T) {
	src := `<? unset($a, $b);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUnset_TrailingComma(t *testing.T) {
	src := `<? unset($a, $b,);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    18,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    18,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse(t *testing.T) {
	src := `<? use Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    10,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  7,
											EndPos:    10,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_FullyQualified(t *testing.T) {
	src := `<? use \Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    11,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    11,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    11,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_FullyQualifiedAlias(t *testing.T) {
	src := `<? use \Foo as Bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    18,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    11,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    11,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    18,
								},
							},
							Value: []byte("Bar"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_List(t *testing.T) {
	src := `<? use Foo, Bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    10,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  7,
											EndPos:    10,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    15,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    15,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    15,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_ListAlias(t *testing.T) {
	src := `<? use Foo, Bar as Baz;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    10,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  7,
											EndPos:    10,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    22,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    15,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    15,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
							},
							Value: []byte("Baz"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_ListFunctionType(t *testing.T) {
	src := `<? use function Foo, \Bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    26,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    15,
						},
					},
					Value: []byte("function"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    19,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    19,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  16,
											EndPos:    19,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  22,
								EndPos:    25,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    25,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  22,
											EndPos:    25,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_ListFunctionTypeAliases(t *testing.T) {
	src := `<? use function Foo as foo, \Bar as bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    40,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    40,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    15,
						},
					},
					Value: []byte("function"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    26,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    19,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  16,
											EndPos:    19,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    26,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  29,
								EndPos:    39,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  29,
									EndPos:    32,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  29,
											EndPos:    32,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  36,
									EndPos:    39,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_ListConstType(t *testing.T) {
	src := `<? use const Foo, \Bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    12,
						},
					},
					Value: []byte("const"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    16,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  13,
											EndPos:    16,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    22,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_ListConstTypeAliases(t *testing.T) {
	src := `<? use const Foo as foo, \Bar as bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    37,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    37,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    12,
						},
					},
					Value: []byte("const"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    23,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  13,
											EndPos:    16,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    23,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  26,
								EndPos:    36,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  26,
									EndPos:    29,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  26,
											EndPos:    29,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  33,
									EndPos:    36,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_GroupUse(t *testing.T) {
	src := `<? use Foo\{Bar, Baz};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    22,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    22,
					},
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    10,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    10,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    15,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    15,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    15,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    20,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    20,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  17,
											EndPos:    20,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_GroupUseAlias(t *testing.T) {
	src := `<? use Foo\{Bar, Baz as quux};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    30,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    30,
					},
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    10,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    10,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    15,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    15,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    15,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    28,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    20,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  17,
											EndPos:    20,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    28,
								},
							},
							Value: []byte("quux"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_FunctionGroupUse(t *testing.T) {
	src := `<? use function Foo\{Bar, Baz};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    31,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    31,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    15,
						},
					},
					Value: []byte("function"),
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  16,
							EndPos:    19,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    19,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  21,
								EndPos:    24,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  21,
									EndPos:    24,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  21,
											EndPos:    24,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  26,
								EndPos:    29,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  26,
									EndPos:    29,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  26,
											EndPos:    29,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_ConstGroupUse(t *testing.T) {
	src := `<? use const Foo\{Bar, Baz};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    28,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    28,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    12,
						},
					},
					Value: []byte("const"),
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    21,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    21,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    21,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    26,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    26,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  23,
											EndPos:    26,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtUse_MixedGroupUse(t *testing.T) {
	src := `<? use Foo\{const Bar, function Baz};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    37,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    37,
					},
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    10,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    10,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    21,
							},
						},
						UseType: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    17,
								},
							},
							Value: []byte("const"),
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    21,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    21,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  32,
								EndPos:    35,
							},
						},
						UseType: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    31,
								},
							},
							Value: []byte("function"),
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  32,
									EndPos:    35,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  32,
											EndPos:    35,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtBreak_Empty(t *testing.T) {
	src := `<? while (1) { break; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    11,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    23,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    21,
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtBreak_Light(t *testing.T) {
	src := `<? while (1) { break 2; }`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    25,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    25,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    11,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    25,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    23,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  21,
										EndPos:    22,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStmtBreak(t *testing.T) {
	src := `<? while (1) : break(3); endwhile;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    34,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtAltWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    34,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    11,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    24,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    24,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  21,
										EndPos:    22,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// expr

func TestExprArrayDimFetch(t *testing.T) {
	src := `<? $a[1];`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    9,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    9,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    7,
							},
						},
						Value: []byte("1"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArrayDimFetch_Nested(t *testing.T) {
	src := `<? $a[1][2];`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    8,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    5,
									},
								},
								Value: []byte("$a"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  6,
									EndPos:    7,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    10,
							},
						},
						Value: []byte("2"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArray(t *testing.T) {
	src := `<? array();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Items: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArray_Item(t *testing.T) {
	src := `<? array(1);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    10,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    10,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArray_Items(t *testing.T) {
	src := `<? array(1=>1, &$b,);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    20,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    13,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    10,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    13,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    18,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  15,
										EndPos:    18,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  16,
											EndPos:    18,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  16,
												EndPos:    18,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
						&ast.ExprArrayItem{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArray_ItemUnpack(t *testing.T) {
	src := `<? array(...$b);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    15,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    14,
								},
							},
							Unpack: true,
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    14,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    14,
										},
									},
									Value: []byte("$b"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArrowFunction(t *testing.T) {
	src := `<? fn() => $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprArrowFunction{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprArrowFunction_ReturnType(t *testing.T) {
	src := `<? fn & () : foo => $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ExprArrowFunction{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					Static:     false,
					ReturnsRef: true,
					ReturnType: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    16,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    16,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    22,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBitwiseNot(t *testing.T) {
	src := `<? ~$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprBitwiseNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    6,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBooleanNot(t *testing.T) {
	src := `<? !$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprBooleanNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    6,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClassConstFetch(t *testing.T) {
	src := `<? Foo::Bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    11,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClassConstFetch_Static(t *testing.T) {
	src := `<? static::bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Class: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    9,
							},
						},
						Value: []byte("static"),
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    14,
							},
						},
						Value: []byte("bar"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClone_Brackets(t *testing.T) {
	src := `<? clone($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClone(t *testing.T) {
	src := `<? clone $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClosure(t *testing.T) {
	src := `<? function(){};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    15,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Stmts:      []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClosure_Use(t *testing.T) {
	src := `<? function($a, $b) use ($c, &$d) {};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    37,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    37,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    36,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    14,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    14,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  16,
										EndPos:    18,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  16,
											EndPos:    18,
										},
									},
									Value: []byte("$b"),
								},
							},
						},
					},
					ClosureUse: &ast.ExprClosureUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    33,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  25,
										EndPos:    27,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  25,
											EndPos:    27,
										},
									},
									Value: []byte("$c"),
								},
							},
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  29,
										EndPos:    32,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    32,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    32,
											},
										},
										Value: []byte("$d"),
									},
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClosure_Use2(t *testing.T) {
	src := `<? function($a, $b) use (&$c, $d) {};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    37,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    37,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    36,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    14,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    14,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  16,
										EndPos:    18,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  16,
											EndPos:    18,
										},
									},
									Value: []byte("$b"),
								},
							},
						},
					},
					ClosureUse: &ast.ExprClosureUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    33,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  25,
										EndPos:    28,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  26,
											EndPos:    28,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  26,
												EndPos:    28,
											},
										},
										Value: []byte("$c"),
									},
								},
							},
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  30,
										EndPos:    32,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    32,
										},
									},
									Value: []byte("$d"),
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprClosure_ReturnType(t *testing.T) {
	src := `<? function(): void {};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					ReturnsRef: false,
					Static:     false,
					ReturnType: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    19,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  15,
										EndPos:    19,
									},
								},
								Value: []byte("void"),
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprConstFetch(t *testing.T) {
	src := `<? foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Const: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprConstFetch_Relative(t *testing.T) {
	src := `<? namespace\foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Const: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    16,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    16,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprConstFetch_FullyQualified(t *testing.T) {
	src := `<? \foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
					Const: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    7,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprEmpty(t *testing.T) {
	src := `<? empty($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprErrorSuppress(t *testing.T) {
	src := `<? @$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprErrorSuppress{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    6,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprEval(t *testing.T) {
	src := `<? eval($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprEval{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprExit(t *testing.T) {
	src := `<? exit;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprExit{
					Die: false,
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprExit_Empty(t *testing.T) {
	src := `<? exit();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    10,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    10,
					},
				},
				Expr: &ast.ExprExit{
					Die: false,
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    9,
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprExit_Expr(t *testing.T) {
	src := `<? exit($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprExit{
					Die: false,
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprDie(t *testing.T) {
	src := `<? die;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprExit{
					Die: true,
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprDie_Empty(t *testing.T) {
	src := `<? die();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    9,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    9,
					},
				},
				Expr: &ast.ExprExit{
					Die: true,
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprDie_Expr(t *testing.T) {
	src := `<? die($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprExit{
					Die: true,
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    9,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    9,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprFunctionCall(t *testing.T) {
	src := `<? foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    9,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    9,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    8,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprFunctionCall_Relative(t *testing.T) {
	src := `<? namespace\foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Function: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    16,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    16,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprFunctionCall_FullyQualified(t *testing.T) {
	src := `<? \foo([]);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    7,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    11,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprShortArray{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    10,
										},
									},
									Items: []ast.Vertex{},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprFunctionCall_Var(t *testing.T) {
	src := `<? $foo(yield $a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    18,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    18,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    17,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    7,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    17,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    16,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprYield{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    16,
										},
									},
									Value: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  14,
												EndPos:    16,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  14,
													EndPos:    16,
												},
											},
											Value: []byte("$a"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprFunctionCall_ExprArg(t *testing.T) {
	src := `<? ceil($foo/3);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    15,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    7,
									},
								},
								Value: []byte("ceil"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    15,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    14,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprBinaryDiv{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    14,
										},
									},
									Left: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  8,
												EndPos:    12,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  8,
													EndPos:    12,
												},
											},
											Value: []byte("$foo"),
										},
									},
									Right: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  13,
												EndPos:    14,
											},
										},
										Value: []byte("3"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprPostDec(t *testing.T) {
	src := `<? $a--;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprPostDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprPostInc(t *testing.T) {
	src := `<? $a++;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprPostInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprPreDec(t *testing.T) {
	src := `<? --$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprPreDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    7,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    7,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprPreInc(t *testing.T) {
	src := `<? ++$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    8,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
				},
				Expr: &ast.ExprPreInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    7,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    7,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprInclude(t *testing.T) {
	src := `<? include $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprInclude{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprInclude_Once(t *testing.T) {
	src := `<? include_once $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprIncludeOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprRequire(t *testing.T) {
	src := `<? require $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprRequire{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprRequire_Once(t *testing.T) {
	src := `<? require_once $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprRequireOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprInstanceOf(t *testing.T) {
	src := `<? $a instanceof Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    20,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    20,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    20,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprInstanceOf_Relative(t *testing.T) {
	src := `<? $a instanceof namespace\Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    31,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    31,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    30,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    30,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  27,
										EndPos:    30,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprInstanceOf_FullyQualified(t *testing.T) {
	src := `<? $a instanceof \Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    22,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    22,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    21,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    21,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    21,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprIsset(t *testing.T) {
	src := `<? isset($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    11,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprIsset_Variables(t *testing.T) {
	src := `<? isset($a, $b);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    11,
									},
								},
								Value: []byte("$a"),
							},
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprList_Empty(t *testing.T) {
	src := `<? list() = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    9,
							},
						},
						Items: []ast.Vertex{},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprList(t *testing.T) {
	src := `<? list($a) = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    11,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    10,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  8,
												EndPos:    10,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    16,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    16,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprList_ArrayIndex(t *testing.T) {
	src := `<? list($a[]) = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    13,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    12,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    12,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  8,
												EndPos:    10,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  8,
													EndPos:    10,
												},
											},
											Value: []byte("$a"),
										},
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprList_List(t *testing.T) {
	src := `<? list(list($a)) = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    17,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    16,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    16,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  13,
													EndPos:    15,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  13,
														EndPos:    15,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 1,
															EndLine:   1,
															StartPos:  13,
															EndPos:    15,
														},
													},
													Value: []byte("$a"),
												},
											},
										},
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    22,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprList_EmptyItem(t *testing.T) {
	src := `<? list(, $a) = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    13,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{},
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    12,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  10,
											EndPos:    12,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  10,
												EndPos:    12,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprList_EmptyItems(t *testing.T) {
	src := `<? list(, , $a, ) = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    17,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{},
							&ast.ExprArrayItem{},
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    14,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  12,
											EndPos:    14,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  12,
												EndPos:    14,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.ExprArrayItem{},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    22,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprMethodCall(t *testing.T) {
	src := `<? $a->foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
						},
						Value: []byte("foo"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprNew(t *testing.T) {
	src := `<? new Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  7,
										EndPos:    10,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprNew_Relative(t *testing.T) {
	src := `<? new namespace\Foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    20,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    20,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprNew_FullyQualified(t *testing.T) {
	src := `<? new \Foo();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    11,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    11,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprNew_Anonymous(t *testing.T) {
	src := `<? new class ($a, ...$b) {};`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    28,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    28,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    27,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    27,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    24,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  14,
											EndPos:    16,
										},
									},
									Variadic:    false,
									IsReference: false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  14,
												EndPos:    16,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  14,
													EndPos:    16,
												},
											},
											Value: []byte("$a"),
										},
									},
								},
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    23,
										},
									},
									IsReference: false,
									Variadic:    true,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  21,
												EndPos:    23,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  21,
													EndPos:    23,
												},
											},
											Value: []byte("$b"),
										},
									},
								},
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprPrint(t *testing.T) {
	src := `<? print($a);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprPrint{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprPropertyFetch(t *testing.T) {
	src := `<? $a->foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprReference_ForeachWithRef(t *testing.T) {
	t.Helper()
	src := `<? foreach ($a as $k => &$v) {}`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    31,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    31,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  24,
							EndPos:    27,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  25,
								EndPos:    27,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  25,
									EndPos:    27,
								},
							},
							Value: []byte("$v"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  29,
							EndPos:    31,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShellExec(t *testing.T) {
	src := "<? `cmd $a`;"

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    8,
								},
							},
							Value: []byte("cmd "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShortArray(t *testing.T) {
	src := `<? [];`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    6,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    6,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
					},
					Items: []ast.Vertex{},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShortArray_Item(t *testing.T) {
	src := `<? [1];`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    5,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    5,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShortArray_Items(t *testing.T) {
	src := `<? [1=>1, &$b,];`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    15,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    8,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    5,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  7,
										EndPos:    8,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    13,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    13,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  11,
											EndPos:    13,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  11,
												EndPos:    13,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
						&ast.ExprArrayItem{},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShortList(t *testing.T) {
	src := `<? [$a] = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    6,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  4,
											EndPos:    6,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  4,
												EndPos:    6,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShortList_ArrayIndex(t *testing.T) {
	src := `<? [$a[]] = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    9,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    8,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  4,
											EndPos:    8,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  4,
												EndPos:    6,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  4,
													EndPos:    6,
												},
											},
											Value: []byte("$a"),
										},
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprShortList_List(t *testing.T) {
	src := `<? [list($a)] = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    19,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    19,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    18,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    13,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    12,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  4,
											EndPos:    12,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  9,
													EndPos:    11,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  9,
														EndPos:    11,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 1,
															EndLine:   1,
															StartPos:  9,
															EndPos:    11,
														},
													},
													Value: []byte("$a"),
												},
											},
										},
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticCall(t *testing.T) {
	src := `<? Foo::bar();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    11,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticCall_Relative(t *testing.T) {
	src := `<? namespace\Foo::bar();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    24,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    24,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    23,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    16,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    16,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    21,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  21,
								EndPos:    23,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticCall_FullyQualified(t *testing.T) {
	src := `<? \Foo::bar();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    7,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    12,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticCall_Var(t *testing.T) {
	src := `<? Foo::$bar();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    12,
								},
							},
							Value: []byte("$bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticCall_VarVar(t *testing.T) {
	src := `<? $foo::$bar();`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    15,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    7,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    13,
								},
							},
							Value: []byte("$bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticPropertyFetch(t *testing.T) {
	src := `<? Foo::$bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    6,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    12,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticPropertyFetch_Relative(t *testing.T) {
	src := `<? namespace\Foo::$bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    23,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    23,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    22,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    16,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    16,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    22,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    22,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprStaticPropertyFetch_FullyQualified(t *testing.T) {
	src := `<? \Foo::$bar;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    7,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    7,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    13,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprTernary(t *testing.T) {
	src := `<? $a ? $b : $c;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    16,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    16,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    15,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							Value: []byte("$c"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprTernary_Simple(t *testing.T) {
	src := `<? $a ? : $c;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$c"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprTernary_NestedTrue(t *testing.T) {
	src := `<? $a ? $b ? $c : $d : $e;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    26,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    25,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfTrue: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    20,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
								},
								Value: []byte("$b"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
								},
								Value: []byte("$c"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
								},
								Value: []byte("$d"),
							},
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    25,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    25,
								},
							},
							Value: []byte("$e"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprTernary_NestedCond(t *testing.T) {
	src := `<? $a ? $b : $c ? $d : $e;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    26,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    26,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    25,
						},
					},
					Condition: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    15,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  3,
										EndPos:    5,
									},
								},
								Value: []byte("$a"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
								},
								Value: []byte("$b"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    15,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
								},
								Value: []byte("$c"),
							},
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
							},
							Value: []byte("$d"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    25,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    25,
								},
							},
							Value: []byte("$e"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprUnaryMinus(t *testing.T) {
	src := `<? -$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprUnaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    6,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprUnaryPlus(t *testing.T) {
	src := `<? +$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprUnaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    6,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprVariable(t *testing.T) {
	src := `<? $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    6,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    6,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						Value: []byte("$a"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprVariable_Variable(t *testing.T) {
	src := `<? $$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    7,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    6,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprYield(t *testing.T) {
	src := `<? yield;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    9,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    9,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprYield_Val(t *testing.T) {
	src := `<? yield $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprYield_KeyVal(t *testing.T) {
	src := `<? yield $a => $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    18,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    18,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    17,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    17,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  15,
									EndPos:    17,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprYield_Expr(t *testing.T) {
	src := `<? yield 1;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Value: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    10,
							},
						},
						Value: []byte("1"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprYield_KeyExpr(t *testing.T) {
	src := `<? yield $a => 1;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
					Value: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    16,
							},
						},
						Value: []byte("1"),
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprYieldFrom(t *testing.T) {
	src := `<? yield from $a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ExprYieldFrom{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    16,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    16,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// expr assign

func TestExprAssign_Assign(t *testing.T) {
	src := `<? $a = $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Reference(t *testing.T) {
	src := `<? $a =& $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_ReferenceNew(t *testing.T) {
	src := `<? $a =& new Foo;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    17,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    17,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    16,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    16,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  13,
											EndPos:    16,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_ReferenceArgs(t *testing.T) {
	src := `<? $a =& new Foo($b);`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    21,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    21,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    20,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    20,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  13,
											EndPos:    16,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    20,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  17,
											EndPos:    19,
										},
									},
									Variadic:    false,
									IsReference: false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  17,
												EndPos:    19,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  17,
													EndPos:    19,
												},
											},
											Value: []byte("$b"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_BitwiseAnd(t *testing.T) {
	src := `<? $a &= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_BitwiseOr(t *testing.T) {
	src := `<? $a |= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_BitwiseXor(t *testing.T) {
	src := `<? $a ^= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Concat(t *testing.T) {
	src := `<? $a .= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Div(t *testing.T) {
	src := `<? $a /= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Minus(t *testing.T) {
	src := `<? $a -= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Mod(t *testing.T) {
	src := `<? $a %= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Mul(t *testing.T) {
	src := `<? $a *= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Plus(t *testing.T) {
	src := `<? $a += $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprAssignPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Pow(t *testing.T) {
	src := `<? $a **= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprAssignPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_ShiftLeft(t *testing.T) {
	src := `<? $a <<= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprAssignShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_ShiftRight(t *testing.T) {
	src := `<? $a >>= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprAssignShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprAssign_Coalesce(t *testing.T) {
	src := `<? $a ??= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprAssignCoalesce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// expr binary

func TestExprBinary_BitwiseAnd(t *testing.T) {
	src := `<? $a & $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_BitwiseOr(t *testing.T) {
	src := `<? $a | $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_BitwiseXor(t *testing.T) {
	src := `<? $a ^ $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_BooleanAnd(t *testing.T) {
	src := `<? $a && $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryBooleanAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_BooleanOr(t *testing.T) {
	src := `<? $a || $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryBooleanOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Coalesce(t *testing.T) {
	src := `<? $a ?? $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryCoalesce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Concat(t *testing.T) {
	src := `<? $a . $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Div(t *testing.T) {
	src := `<? $a / $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Equal(t *testing.T) {
	src := `<? $a == $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_GreaterOrEqual(t *testing.T) {
	src := `<? $a >= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryGreaterOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Greater(t *testing.T) {
	src := `<? $a > $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryGreater{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Identical(t *testing.T) {
	src := `<? $a === $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprBinaryIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_LogicalAnd(t *testing.T) {
	src := `<? $a and $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprBinaryLogicalAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_LogicalOr(t *testing.T) {
	src := `<? $a or $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryLogicalOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_LogicalXor(t *testing.T) {
	src := `<? $a xor $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprBinaryLogicalXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Minus(t *testing.T) {
	src := `<? $a - $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Mod(t *testing.T) {
	src := `<? $a % $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Mul(t *testing.T) {
	src := `<? $a * $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_NotEqual(t *testing.T) {
	src := `<? $a != $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryNotEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_NotIdentical(t *testing.T) {
	src := `<? $a !== $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprBinaryNotIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Plus(t *testing.T) {
	src := `<? $a + $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Pow(t *testing.T) {
	src := `<? $a ** $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_ShiftLeft(t *testing.T) {
	src := `<? $a << $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_ShiftRight(t *testing.T) {
	src := `<? $a >> $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinaryShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_SmallerOrEqual(t *testing.T) {
	src := `<? $a <= $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprBinarySmallerOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Smaller(t *testing.T) {
	src := `<? $a < $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprBinarySmaller{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprBinary_Spaceship(t *testing.T) {
	src := `<? $a <=> $b;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprBinarySpaceship{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

// expr cast

func TestExprCast_Array(t *testing.T) {
	src := `<? (array)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprCastArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_Bool(t *testing.T) {
	src := `<? (boolean)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_BoolShort(t *testing.T) {
	src := `<? (bool)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    12,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    11,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_Double(t *testing.T) {
	src := `<? (double)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_CastFloat(t *testing.T) {
	src := `<? (float)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_Int(t *testing.T) {
	src := `<? (integer)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    15,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    15,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    14,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_IntShort(t *testing.T) {
	src := `<? (int)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    11,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    10,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_Object(t *testing.T) {
	src := `<? (object)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprCastObject{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_String(t *testing.T) {
	src := `<? (string)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprCastString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_BinaryString(t *testing.T) {
	src := `<? (binary)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    14,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
				},
				Expr: &ast.ExprCastString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    13,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExprCast_Unset(t *testing.T) {
	src := `<? (unset)$a;`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  3,
				EndPos:    13,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
				},
				Expr: &ast.ExprCastUnset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    12,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    12,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
