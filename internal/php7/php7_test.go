package php7_test

import (
	"io/ioutil"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/internal/php7"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
)

func TestPhp7(t *testing.T) {
	src, err := ioutil.ReadFile("test.php")
	assert.NilError(t, err)

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   348,
				StartPos:  3,
				EndPos:    5706,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  3,
						EndPos:    18,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  3,
							EndPos:    17,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  3,
								EndPos:    6,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
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
								StartLine: 2,
								EndLine:   2,
								StartPos:  6,
								EndPos:    17,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  7,
										EndPos:    9,
									},
								},
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  7,
											EndPos:    9,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  7,
												EndPos:    9,
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
										StartPos:  11,
										EndPos:    16,
									},
								},
								Variadic: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  14,
											EndPos:    16,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
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
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  19,
						EndPos:    35,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  19,
							EndPos:    34,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  19,
								EndPos:    23,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  19,
									EndPos:    23,
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
								StartPos:  23,
								EndPos:    34,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  24,
										EndPos:    26,
									},
								},
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  24,
											EndPos:    26,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  24,
												EndPos:    26,
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
										StartPos:  28,
										EndPos:    33,
									},
								},
								Variadic: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  31,
											EndPos:    33,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  31,
												EndPos:    33,
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
						StartPos:  36,
						EndPos:    57,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  36,
							EndPos:    56,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  36,
								EndPos:    40,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  36,
									EndPos:    40,
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
								StartPos:  42,
								EndPos:    45,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  45,
								EndPos:    56,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  46,
										EndPos:    48,
									},
								},
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  46,
											EndPos:    48,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  46,
												EndPos:    48,
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
										StartPos:  50,
										EndPos:    55,
									},
								},
								Variadic: true,
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
						StartPos:  58,
						EndPos:    78,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  58,
							EndPos:    77,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  58,
								EndPos:    61,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  58,
										EndPos:    61,
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
								StartPos:  63,
								EndPos:    66,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  66,
								EndPos:    77,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  67,
										EndPos:    69,
									},
								},
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  67,
											EndPos:    69,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  67,
												EndPos:    69,
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
										StartPos:  71,
										EndPos:    76,
									},
								},
								Variadic: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  74,
											EndPos:    76,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  74,
												EndPos:    76,
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
						StartPos:  79,
						EndPos:    100,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  79,
							EndPos:    99,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  79,
								EndPos:    83,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  79,
									EndPos:    83,
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
								StartPos:  85,
								EndPos:    88,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  88,
								EndPos:    99,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  89,
										EndPos:    91,
									},
								},
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  89,
											EndPos:    91,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 6,
												EndLine:   6,
												StartPos:  89,
												EndPos:    91,
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
										StartPos:  93,
										EndPos:    98,
									},
								},
								Variadic: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  96,
											EndPos:    98,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 6,
												EndLine:   6,
												StartPos:  96,
												EndPos:    98,
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
						StartPos:  101,
						EndPos:    120,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  101,
							EndPos:    119,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 7,
								EndLine:   7,
								StartPos:  105,
								EndPos:    108,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  105,
										EndPos:    108,
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
								StartPos:  108,
								EndPos:    119,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  109,
										EndPos:    111,
									},
								},
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  109,
											EndPos:    111,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 7,
												EndLine:   7,
												StartPos:  109,
												EndPos:    111,
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
										StartPos:  113,
										EndPos:    118,
									},
								},
								Variadic: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  116,
											EndPos:    118,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 7,
												EndLine:   7,
												StartPos:  116,
												EndPos:    118,
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
						StartPos:  144,
						EndPos:    169,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  144,
							EndPos:    168,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  148,
								EndPos:    168,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  154,
									EndPos:    165,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  155,
											EndPos:    157,
										},
									},
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  155,
												EndPos:    157,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 9,
													EndLine:   9,
													StartPos:  155,
													EndPos:    157,
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
											StartPos:  159,
											EndPos:    164,
										},
									},
									Variadic: true,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  162,
												EndPos:    164,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 9,
													EndLine:   9,
													StartPos:  162,
													EndPos:    164,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  170,
						EndPos:    183,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  170,
							EndPos:    182,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  174,
								EndPos:    182,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  184,
						EndPos:    193,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  184,
							EndPos:    192,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  188,
								EndPos:    192,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  188,
									EndPos:    192,
								},
							},
							Value: []byte("$foo"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 12,
						EndLine:   12,
						StartPos:  194,
						EndPos:    206,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 12,
							EndLine:   12,
							StartPos:  194,
							EndPos:    205,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  198,
								EndPos:    205,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  198,
									EndPos:    202,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  198,
										EndPos:    202,
									},
								},
								Value: []byte("$foo"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  203,
									EndPos:    204,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   13,
						StartPos:  207,
						EndPos:    222,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 13,
							EndLine:   13,
							StartPos:  207,
							EndPos:    221,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 13,
								EndLine:   13,
								StartPos:  211,
								EndPos:    221,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  211,
									EndPos:    215,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 13,
										EndLine:   13,
										StartPos:  211,
										EndPos:    215,
									},
								},
								Value: []byte("$foo"),
							},
						},
						Dim: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  216,
									EndPos:    220,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 13,
										EndLine:   13,
										StartPos:  216,
										EndPos:    220,
									},
								},
								Value: []byte("$bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 14,
						EndLine:   14,
						StartPos:  223,
						EndPos:    237,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 14,
							EndLine:   14,
							StartPos:  223,
							EndPos:    236,
						},
					},
					Class: &ast.ExprPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  227,
								EndPos:    236,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  227,
									EndPos:    231,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 14,
										EndLine:   14,
										StartPos:  227,
										EndPos:    231,
									},
								},
								Value: []byte("$foo"),
							},
						},
						Property: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  233,
									EndPos:    236,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 15,
						EndLine:   15,
						StartPos:  238,
						EndPos:    253,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 15,
							EndLine:   15,
							StartPos:  238,
							EndPos:    252,
						},
					},
					Class: &ast.ExprStaticPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 15,
								EndLine:   15,
								StartPos:  242,
								EndPos:    252,
							},
						},
						Class: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  242,
									EndPos:    246,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 15,
										EndLine:   15,
										StartPos:  242,
										EndPos:    246,
									},
								},
								Value: []byte("$foo"),
							},
						},
						Property: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  248,
									EndPos:    252,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 15,
										EndLine:   15,
										StartPos:  248,
										EndPos:    252,
									},
								},
								Value: []byte("$bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 16,
						EndLine:   16,
						StartPos:  254,
						EndPos:    271,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 16,
							EndLine:   16,
							StartPos:  254,
							EndPos:    270,
						},
					},
					Class: &ast.ExprStaticPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 16,
								EndLine:   16,
								StartPos:  258,
								EndPos:    270,
							},
						},
						Class: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 16,
									EndLine:   16,
									StartPos:  258,
									EndPos:    264,
								},
							},
							Value: []byte("static"),
						},
						Property: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 16,
									EndLine:   16,
									StartPos:  266,
									EndPos:    270,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 16,
										EndLine:   16,
										StartPos:  266,
										EndPos:    270,
									},
								},
								Value: []byte("$bar"),
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  273,
						EndPos:    318,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  282,
							EndPos:    285,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  286,
								EndPos:    300,
							},
						},
						Type: &ast.Nullable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  286,
									EndPos:    290,
								},
							},
							Expr: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  287,
										EndPos:    290,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 18,
												EndLine:   18,
												StartPos:  287,
												EndPos:    290,
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
									StartLine: 18,
									EndLine:   18,
									StartPos:  291,
									EndPos:    295,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  291,
										EndPos:    295,
									},
								},
								Value: []byte("$bar"),
							},
						},
						DefaultValue: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  296,
									EndPos:    300,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  296,
										EndPos:    300,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 18,
												EndLine:   18,
												StartPos:  296,
												EndPos:    300,
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
								StartLine: 18,
								EndLine:   18,
								StartPos:  302,
								EndPos:    314,
							},
						},
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  302,
									EndPos:    305,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  302,
											EndPos:    305,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
						Var: &ast.Reference{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  306,
									EndPos:    314,
								},
							},
							Var: &ast.Variadic{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  307,
										EndPos:    314,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  310,
											EndPos:    314,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 18,
												EndLine:   18,
												StartPos:  310,
												EndPos:    314,
											},
										},
										Value: []byte("$baz"),
									},
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 19,
						EndLine:   19,
						StartPos:  319,
						EndPos:    383,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 19,
							EndLine:   19,
							StartPos:  325,
							EndPos:    328,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 19,
								EndLine:   19,
								StartPos:  330,
								EndPos:    382,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 19,
									EndLine:   19,
									StartPos:  346,
									EndPos:    349,
								},
							},
							Value: []byte("foo"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  330,
										EndPos:    336,
									},
								},
								Value: []byte("public"),
							},
						},
						Params: []ast.Vertex{
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  350,
										EndPos:    364,
									},
								},
								Type: &ast.Nullable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  350,
											EndPos:    354,
										},
									},
									Expr: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  351,
												EndPos:    354,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 19,
														EndLine:   19,
														StartPos:  351,
														EndPos:    354,
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
											StartLine: 19,
											EndLine:   19,
											StartPos:  355,
											EndPos:    359,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  355,
												EndPos:    359,
											},
										},
										Value: []byte("$bar"),
									},
								},
								DefaultValue: &ast.ExprConstFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  360,
											EndPos:    364,
										},
									},
									Const: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  360,
												EndPos:    364,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 19,
														EndLine:   19,
														StartPos:  360,
														EndPos:    364,
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
										StartLine: 19,
										EndLine:   19,
										StartPos:  366,
										EndPos:    378,
									},
								},
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  366,
											EndPos:    369,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  366,
													EndPos:    369,
												},
											},
											Value: []byte("baz"),
										},
									},
								},
								Var: &ast.Reference{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  370,
											EndPos:    378,
										},
									},
									Var: &ast.Variadic{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  371,
												EndPos:    378,
											},
										},
										Var: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  374,
													EndPos:    378,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 19,
														EndLine:   19,
														StartPos:  374,
														EndPos:    378,
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
									StartLine: 19,
									EndLine:   19,
									StartPos:  380,
									EndPos:    382,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 20,
						EndLine:   20,
						StartPos:  384,
						EndPos:    426,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 20,
							EndLine:   20,
							StartPos:  384,
							EndPos:    425,
						},
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  393,
									EndPos:    407,
								},
							},
							Type: &ast.Nullable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  393,
										EndPos:    397,
									},
								},
								Expr: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  394,
											EndPos:    397,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 20,
													EndLine:   20,
													StartPos:  394,
													EndPos:    397,
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
										StartLine: 20,
										EndLine:   20,
										StartPos:  398,
										EndPos:    402,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  398,
											EndPos:    402,
										},
									},
									Value: []byte("$bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  403,
										EndPos:    407,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  403,
											EndPos:    407,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 20,
													EndLine:   20,
													StartPos:  403,
													EndPos:    407,
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
									StartLine: 20,
									EndLine:   20,
									StartPos:  409,
									EndPos:    421,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  409,
										EndPos:    412,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  409,
												EndPos:    412,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.Reference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  413,
										EndPos:    421,
									},
								},
								Var: &ast.Variadic{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  414,
											EndPos:    421,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  417,
												EndPos:    421,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 20,
													EndLine:   20,
													StartPos:  417,
													EndPos:    421,
												},
											},
											Value: []byte("$baz"),
										},
									},
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 21,
						EndLine:   21,
						StartPos:  427,
						EndPos:    476,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 21,
							EndLine:   21,
							StartPos:  427,
							EndPos:    475,
						},
					},
					Static: true,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  443,
									EndPos:    457,
								},
							},
							Type: &ast.Nullable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  443,
										EndPos:    447,
									},
								},
								Expr: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  444,
											EndPos:    447,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 21,
													EndLine:   21,
													StartPos:  444,
													EndPos:    447,
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
										StartLine: 21,
										EndLine:   21,
										StartPos:  448,
										EndPos:    452,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  448,
											EndPos:    452,
										},
									},
									Value: []byte("$bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  453,
										EndPos:    457,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  453,
											EndPos:    457,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 21,
													EndLine:   21,
													StartPos:  453,
													EndPos:    457,
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
									StartLine: 21,
									EndLine:   21,
									StartPos:  459,
									EndPos:    471,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  459,
										EndPos:    462,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  459,
												EndPos:    462,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.Reference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  463,
										EndPos:    471,
									},
								},
								Var: &ast.Variadic{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  464,
											EndPos:    471,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  467,
												EndPos:    471,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 21,
													EndLine:   21,
													StartPos:  467,
													EndPos:    471,
												},
											},
											Value: []byte("$baz"),
										},
									},
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 23,
						EndLine:   23,
						StartPos:  478,
						EndPos:    498,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 23,
							EndLine:   23,
							StartPos:  478,
							EndPos:    497,
						},
					},
					Value: []byte("1234567890123456789"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  499,
						EndPos:    520,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 24,
							EndLine:   24,
							StartPos:  499,
							EndPos:    519,
						},
					},
					Value: []byte("12345678901234567890"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  521,
						EndPos:    524,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 25,
							EndLine:   25,
							StartPos:  521,
							EndPos:    523,
						},
					},
					Value: []byte("0."),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  525,
						EndPos:    592,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 26,
							EndLine:   26,
							StartPos:  525,
							EndPos:    591,
						},
					},
					Value: []byte("0b0111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  593,
						EndPos:    660,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 27,
							EndLine:   27,
							StartPos:  593,
							EndPos:    659,
						},
					},
					Value: []byte("0b1111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  661,
						EndPos:    682,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 28,
							EndLine:   28,
							StartPos:  661,
							EndPos:    681,
						},
					},
					Value: []byte("0x007111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 29,
						EndLine:   29,
						StartPos:  683,
						EndPos:    702,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 29,
							EndLine:   29,
							StartPos:  683,
							EndPos:    701,
						},
					},
					Value: []byte("0x8111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  703,
						EndPos:    713,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 30,
							EndLine:   30,
							StartPos:  703,
							EndPos:    712,
						},
					},
					Value: []byte("__CLASS__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 31,
						EndLine:   31,
						StartPos:  714,
						EndPos:    722,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 31,
							EndLine:   31,
							StartPos:  714,
							EndPos:    721,
						},
					},
					Value: []byte("__DIR__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 32,
						EndLine:   32,
						StartPos:  723,
						EndPos:    732,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 32,
							EndLine:   32,
							StartPos:  723,
							EndPos:    731,
						},
					},
					Value: []byte("__FILE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 33,
						EndLine:   33,
						StartPos:  733,
						EndPos:    746,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 33,
							EndLine:   33,
							StartPos:  733,
							EndPos:    745,
						},
					},
					Value: []byte("__FUNCTION__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 34,
						EndLine:   34,
						StartPos:  747,
						EndPos:    756,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 34,
							EndLine:   34,
							StartPos:  747,
							EndPos:    755,
						},
					},
					Value: []byte("__LINE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 35,
						EndLine:   35,
						StartPos:  757,
						EndPos:    771,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 35,
							EndLine:   35,
							StartPos:  757,
							EndPos:    770,
						},
					},
					Value: []byte("__NAMESPACE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 36,
						EndLine:   36,
						StartPos:  772,
						EndPos:    783,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 36,
							EndLine:   36,
							StartPos:  772,
							EndPos:    782,
						},
					},
					Value: []byte("__METHOD__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 37,
						EndLine:   37,
						StartPos:  784,
						EndPos:    794,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 37,
							EndLine:   37,
							StartPos:  784,
							EndPos:    793,
						},
					},
					Value: []byte("__TRAIT__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 39,
						EndLine:   39,
						StartPos:  796,
						EndPos:    808,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 39,
							EndLine:   39,
							StartPos:  796,
							EndPos:    807,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  797,
									EndPos:    802,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  802,
									EndPos:    806,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  802,
										EndPos:    806,
									},
								},
								Value: []byte("$var"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 40,
						EndLine:   40,
						StartPos:  809,
						EndPos:    824,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 40,
							EndLine:   40,
							StartPos:  809,
							EndPos:    823,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  810,
									EndPos:    815,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  815,
									EndPos:    822,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 40,
										EndLine:   40,
										StartPos:  815,
										EndPos:    819,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 40,
											EndLine:   40,
											StartPos:  815,
											EndPos:    819,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 40,
										EndLine:   40,
										StartPos:  820,
										EndPos:    821,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 41,
						EndLine:   41,
						StartPos:  825,
						EndPos:    841,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 41,
							EndLine:   41,
							StartPos:  825,
							EndPos:    840,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  826,
									EndPos:    831,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  831,
									EndPos:    839,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  831,
										EndPos:    835,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 41,
											EndLine:   41,
											StartPos:  831,
											EndPos:    835,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ExprUnaryMinus{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  836,
										EndPos:    838,
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 41,
											EndLine:   41,
											StartPos:  836,
											EndPos:    838,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 42,
						EndLine:   42,
						StartPos:  842,
						EndPos:    896,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 42,
							EndLine:   42,
							StartPos:  842,
							EndPos:    895,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  843,
									EndPos:    848,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  848,
									EndPos:    894,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 42,
										EndLine:   42,
										StartPos:  848,
										EndPos:    852,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 42,
											EndLine:   42,
											StartPos:  848,
											EndPos:    852,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 42,
										EndLine:   42,
										StartPos:  853,
										EndPos:    893,
									},
								},
								Value: []byte("1234567890123456789012345678901234567890"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 43,
						EndLine:   43,
						StartPos:  897,
						EndPos:    952,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 43,
							EndLine:   43,
							StartPos:  897,
							EndPos:    951,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  898,
									EndPos:    903,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  903,
									EndPos:    950,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 43,
										EndLine:   43,
										StartPos:  903,
										EndPos:    907,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 43,
											EndLine:   43,
											StartPos:  903,
											EndPos:    907,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 43,
										EndLine:   43,
										StartPos:  908,
										EndPos:    949,
									},
								},
								Value: []byte("-1234567890123456789012345678901234567890"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 44,
						EndLine:   44,
						StartPos:  953,
						EndPos:    970,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 44,
							EndLine:   44,
							StartPos:  953,
							EndPos:    969,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  954,
									EndPos:    959,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  959,
									EndPos:    968,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  959,
										EndPos:    963,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 44,
											EndLine:   44,
											StartPos:  959,
											EndPos:    963,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  964,
										EndPos:    967,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 45,
						EndLine:   45,
						StartPos:  971,
						EndPos:    989,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 45,
							EndLine:   45,
							StartPos:  971,
							EndPos:    988,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  972,
									EndPos:    977,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  977,
									EndPos:    987,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  977,
										EndPos:    981,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 45,
											EndLine:   45,
											StartPos:  977,
											EndPos:    981,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  982,
										EndPos:    986,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 45,
											EndLine:   45,
											StartPos:  982,
											EndPos:    986,
										},
									},
									Value: []byte("$bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 46,
						EndLine:   46,
						StartPos:  990,
						EndPos:    1002,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 46,
							EndLine:   46,
							StartPos:  990,
							EndPos:    1001,
						},
					},
					Parts: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  991,
									EndPos:    995,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 46,
										EndLine:   46,
										StartPos:  991,
										EndPos:    995,
									},
								},
								Value: []byte("$foo"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  995,
									EndPos:    996,
								},
							},
							Value: []byte(" "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  996,
									EndPos:    1000,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 46,
										EndLine:   46,
										StartPos:  996,
										EndPos:    1000,
									},
								},
								Value: []byte("$bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 47,
						EndLine:   47,
						StartPos:  1003,
						EndPos:    1022,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 47,
							EndLine:   47,
							StartPos:  1003,
							EndPos:    1021,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1004,
									EndPos:    1009,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1009,
									EndPos:    1018,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 47,
										EndLine:   47,
										StartPos:  1009,
										EndPos:    1013,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 47,
											EndLine:   47,
											StartPos:  1009,
											EndPos:    1013,
										},
									},
									Value: []byte("$foo"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 47,
										EndLine:   47,
										StartPos:  1015,
										EndPos:    1018,
									},
								},
								Value: []byte("bar"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1018,
									EndPos:    1020,
								},
							},
							Value: []byte("()"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 48,
						EndLine:   48,
						StartPos:  1023,
						EndPos:    1037,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 48,
							EndLine:   48,
							StartPos:  1023,
							EndPos:    1036,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 48,
									EndLine:   48,
									StartPos:  1024,
									EndPos:    1029,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 48,
									EndLine:   48,
									StartPos:  1029,
									EndPos:    1035,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 48,
										EndLine:   48,
										StartPos:  1031,
										EndPos:    1034,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 49,
						EndLine:   49,
						StartPos:  1038,
						EndPos:    1055,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 49,
							EndLine:   49,
							StartPos:  1038,
							EndPos:    1054,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1039,
									EndPos:    1044,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1044,
									EndPos:    1053,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 49,
										EndLine:   49,
										StartPos:  1046,
										EndPos:    1049,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 49,
											EndLine:   49,
											StartPos:  1046,
											EndPos:    1049,
										},
									},
									Value: []byte("foo"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 49,
										EndLine:   49,
										StartPos:  1050,
										EndPos:    1051,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 50,
						EndLine:   50,
						StartPos:  1056,
						EndPos:    1071,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 50,
							EndLine:   50,
							StartPos:  1056,
							EndPos:    1070,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1057,
									EndPos:    1062,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1062,
									EndPos:    1069,
								},
							},
							VarName: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 50,
										EndLine:   50,
										StartPos:  1064,
										EndPos:    1068,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 50,
											EndLine:   50,
											StartPos:  1064,
											EndPos:    1068,
										},
									},
									Value: []byte("$foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 51,
						EndLine:   51,
						StartPos:  1072,
						EndPos:    1093,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 51,
							EndLine:   51,
							StartPos:  1072,
							EndPos:    1092,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1073,
									EndPos:    1078,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1079,
									EndPos:    1090,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1079,
										EndPos:    1083,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 51,
											EndLine:   51,
											StartPos:  1079,
											EndPos:    1083,
										},
									},
									Value: []byte("$foo"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1085,
										EndPos:    1088,
									},
								},
								Value: []byte("bar"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1088,
										EndPos:    1090,
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 53,
						EndLine:   54,
						StartPos:  1095,
						EndPos:    1111,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 53,
							EndLine:   53,
							StartPos:  1099,
							EndPos:    1101,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 53,
								EndLine:   53,
								StartPos:  1099,
								EndPos:    1101,
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
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 55,
						EndLine:   57,
						StartPos:  1112,
						EndPos:    1141,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  1116,
							EndPos:    1118,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 55,
								EndLine:   55,
								StartPos:  1116,
								EndPos:    1118,
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
								StartLine: 56,
								EndLine:   -1,
								StartPos:  1122,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1130,
									EndPos:    1132,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 56,
										EndLine:   56,
										StartPos:  1130,
										EndPos:    1132,
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
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 58,
						EndLine:   60,
						StartPos:  1142,
						EndPos:    1164,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1146,
							EndPos:    1148,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 58,
								EndLine:   58,
								StartPos:  1146,
								EndPos:    1148,
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
							StartLine: 59,
							EndLine:   -1,
							StartPos:  1152,
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
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 61,
						EndLine:   65,
						StartPos:  1165,
						EndPos:    1213,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1169,
							EndPos:    1171,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1169,
								EndPos:    1171,
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
								StartLine: 62,
								EndLine:   -1,
								StartPos:  1175,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1183,
									EndPos:    1185,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 62,
										EndLine:   62,
										StartPos:  1183,
										EndPos:    1185,
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
								StartLine: 63,
								EndLine:   -1,
								StartPos:  1188,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 63,
									EndLine:   63,
									StartPos:  1196,
									EndPos:    1198,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 63,
										EndLine:   63,
										StartPos:  1196,
										EndPos:    1198,
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
							StartLine: 64,
							EndLine:   -1,
							StartPos:  1201,
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
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1215,
						EndPos:    1235,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 67,
							EndLine:   67,
							StartPos:  1222,
							EndPos:    1223,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 67,
							EndLine:   67,
							StartPos:  1225,
							EndPos:    1235,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1227,
									EndPos:    1233,
								},
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1236,
						EndPos:    1258,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1243,
							EndPos:    1244,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1246,
							EndPos:    1258,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 68,
									EndLine:   68,
									StartPos:  1248,
									EndPos:    1256,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 68,
										EndLine:   68,
										StartPos:  1254,
										EndPos:    1255,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtAltWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1259,
						EndPos:    1290,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1266,
							EndPos:    1267,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1271,
							EndPos:    1280,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1271,
									EndPos:    1280,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 69,
										EndLine:   69,
										StartPos:  1277,
										EndPos:    1278,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 70,
						EndLine:   70,
						StartPos:  1291,
						EndPos:    1334,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1297,
							EndPos:    1300,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 70,
								EndLine:   70,
								StartPos:  1302,
								EndPos:    1332,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1302,
										EndPos:    1308,
									},
								},
								Value: []byte("public"),
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1315,
										EndPos:    1322,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1315,
											EndPos:    1318,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1321,
											EndPos:    1322,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1324,
										EndPos:    1331,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1324,
											EndPos:    1327,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1330,
											EndPos:    1331,
										},
									},
									Value: []byte("2"),
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 71,
						EndLine:   71,
						StartPos:  1335,
						EndPos:    1371,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1341,
							EndPos:    1344,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 71,
								EndLine:   71,
								StartPos:  1346,
								EndPos:    1369,
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1352,
										EndPos:    1359,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1352,
											EndPos:    1355,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1358,
											EndPos:    1359,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1361,
										EndPos:    1368,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1361,
											EndPos:    1364,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1367,
											EndPos:    1368,
										},
									},
									Value: []byte("2"),
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 72,
						EndLine:   72,
						StartPos:  1372,
						EndPos:    1402,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1378,
							EndPos:    1381,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1383,
								EndPos:    1400,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1392,
									EndPos:    1395,
								},
							},
							Value: []byte("bar"),
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1398,
									EndPos:    1400,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 73,
						EndLine:   73,
						StartPos:  1403,
						EndPos:    1448,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 73,
							EndLine:   73,
							StartPos:  1409,
							EndPos:    1412,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1414,
								EndPos:    1446,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1438,
									EndPos:    1441,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 73,
										EndLine:   73,
										StartPos:  1414,
										EndPos:    1420,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 73,
										EndLine:   73,
										StartPos:  1421,
										EndPos:    1427,
									},
								},
								Value: []byte("static"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1444,
									EndPos:    1446,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 74,
						EndLine:   74,
						StartPos:  1449,
						EndPos:    1500,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1455,
							EndPos:    1458,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1460,
								EndPos:    1498,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1484,
									EndPos:    1487,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 74,
										EndLine:   74,
										StartPos:  1460,
										EndPos:    1466,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 74,
										EndLine:   74,
										StartPos:  1467,
										EndPos:    1473,
									},
								},
								Value: []byte("static"),
							},
						},
						ReturnType: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1491,
									EndPos:    1495,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 74,
											EndLine:   74,
											StartPos:  1491,
											EndPos:    1495,
										},
									},
									Value: []byte("void"),
								},
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1496,
									EndPos:    1498,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 75,
						EndLine:   75,
						StartPos:  1501,
						EndPos:    1522,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 75,
							EndLine:   75,
							StartPos:  1516,
							EndPos:    1519,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1501,
								EndPos:    1509,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1523,
						EndPos:    1554,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1535,
							EndPos:    1538,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1523,
								EndPos:    1528,
							},
						},
						Value: []byte("final"),
					},
				},
				Extends: &ast.StmtClassExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1539,
							EndPos:    1550,
						},
					},
					ClassName: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1547,
								EndPos:    1550,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 76,
										EndLine:   76,
										StartPos:  1547,
										EndPos:    1550,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1555,
						EndPos:    1589,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1567,
							EndPos:    1570,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1555,
								EndPos:    1560,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1571,
							EndPos:    1585,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1582,
									EndPos:    1585,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 77,
											EndLine:   77,
											StartPos:  1582,
											EndPos:    1585,
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
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1590,
						EndPos:    1629,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1602,
							EndPos:    1605,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1590,
								EndPos:    1595,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1606,
							EndPos:    1625,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1617,
									EndPos:    1620,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 78,
											EndLine:   78,
											StartPos:  1617,
											EndPos:    1620,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1622,
									EndPos:    1625,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 78,
											EndLine:   78,
											StartPos:  1622,
											EndPos:    1625,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1630,
						EndPos:    1678,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1630,
							EndPos:    1677,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1634,
								EndPos:    1677,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1639,
									EndPos:    1641,
								},
							},
						},
						Extends: &ast.StmtClassExtends{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1642,
									EndPos:    1653,
								},
							},
							ClassName: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 79,
										EndLine:   79,
										StartPos:  1650,
										EndPos:    1653,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 79,
												EndLine:   79,
												StartPos:  1650,
												EndPos:    1653,
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
									StartLine: 79,
									EndLine:   79,
									StartPos:  1654,
									EndPos:    1673,
								},
							},
							InterfaceNames: []ast.Vertex{
								&ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 79,
											EndLine:   79,
											StartPos:  1665,
											EndPos:    1668,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 79,
													EndLine:   79,
													StartPos:  1665,
													EndPos:    1668,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
								&ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 79,
											EndLine:   79,
											StartPos:  1670,
											EndPos:    1673,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 79,
													EndLine:   79,
													StartPos:  1670,
													EndPos:    1673,
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
			&ast.StmtConstList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 81,
						EndLine:   81,
						StartPos:  1680,
						EndPos:    1703,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1686,
								EndPos:    1693,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1686,
									EndPos:    1689,
								},
							},
							Value: []byte("FOO"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1692,
									EndPos:    1693,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1695,
								EndPos:    1702,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1695,
									EndPos:    1698,
								},
							},
							Value: []byte("BAR"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1701,
									EndPos:    1702,
								},
							},
							Value: []byte("2"),
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1704,
						EndPos:    1727,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1711,
							EndPos:    1712,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1714,
							EndPos:    1727,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 82,
									EndLine:   82,
									StartPos:  1716,
									EndPos:    1725,
								},
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1728,
						EndPos:    1753,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1735,
							EndPos:    1736,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1738,
							EndPos:    1753,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 83,
									EndLine:   83,
									StartPos:  1740,
									EndPos:    1751,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 83,
										EndLine:   83,
										StartPos:  1749,
										EndPos:    1750,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1754,
						EndPos:    1780,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1761,
							EndPos:    1762,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1764,
							EndPos:    1780,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 84,
									EndLine:   84,
									StartPos:  1766,
									EndPos:    1778,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 84,
										EndLine:   84,
										StartPos:  1775,
										EndPos:    1776,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  1781,
						EndPos:    1798,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1789,
								EndPos:    1796,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  1789,
									EndPos:    1794,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  1795,
									EndPos:    1796,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtNop{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1797,
							EndPos:    1798,
						},
					},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 86,
						EndLine:   86,
						StartPos:  1799,
						EndPos:    1818,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 86,
								EndLine:   86,
								StartPos:  1807,
								EndPos:    1814,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 86,
									EndLine:   86,
									StartPos:  1807,
									EndPos:    1812,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 86,
									EndLine:   86,
									StartPos:  1813,
									EndPos:    1814,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 86,
							EndLine:   86,
							StartPos:  1816,
							EndPos:    1818,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 87,
						EndLine:   87,
						StartPos:  1819,
						EndPos:    1848,
					},
				},
				Alt: true,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 87,
								EndLine:   87,
								StartPos:  1827,
								EndPos:    1834,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 87,
									EndLine:   87,
									StartPos:  1827,
									EndPos:    1832,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 87,
									EndLine:   87,
									StartPos:  1833,
									EndPos:    1834,
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
			},
			&ast.StmtDo{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   88,
						StartPos:  1849,
						EndPos:    1864,
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 88,
							EndLine:   88,
							StartPos:  1852,
							EndPos:    1854,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 88,
							EndLine:   88,
							StartPos:  1861,
							EndPos:    1862,
						},
					},
					Value: []byte("1"),
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 89,
						EndLine:   89,
						StartPos:  1865,
						EndPos:    1876,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  1870,
								EndPos:    1872,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 89,
									EndLine:   89,
									StartPos:  1870,
									EndPos:    1872,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  1874,
								EndPos:    1875,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 90,
						EndLine:   90,
						StartPos:  1877,
						EndPos:    1886,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  1882,
								EndPos:    1884,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 90,
									EndLine:   90,
									StartPos:  1882,
									EndPos:    1884,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 91,
						EndLine:   91,
						StartPos:  1887,
						EndPos:    1922,
					},
				},
				Init: []ast.Vertex{
					&ast.ExprAssign{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  1891,
								EndPos:    1897,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1891,
									EndPos:    1893,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  1891,
										EndPos:    1893,
									},
								},
								Value: []byte("$i"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1896,
									EndPos:    1897,
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
								StartLine: 91,
								EndLine:   91,
								StartPos:  1899,
								EndPos:    1906,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1899,
									EndPos:    1901,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  1899,
										EndPos:    1901,
									},
								},
								Value: []byte("$i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1904,
									EndPos:    1906,
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
								StartLine: 91,
								EndLine:   91,
								StartPos:  1908,
								EndPos:    1912,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1908,
									EndPos:    1910,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  1908,
										EndPos:    1910,
									},
								},
								Value: []byte("$i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  1914,
								EndPos:    1918,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1914,
									EndPos:    1916,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  1914,
										EndPos:    1916,
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
							StartLine: 91,
							EndLine:   91,
							StartPos:  1920,
							EndPos:    1922,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 92,
						EndLine:   92,
						StartPos:  1923,
						EndPos:    1959,
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  1929,
								EndPos:    1936,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  1929,
									EndPos:    1931,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 92,
										EndLine:   92,
										StartPos:  1929,
										EndPos:    1931,
									},
								},
								Value: []byte("$i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  1934,
									EndPos:    1936,
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
								StartLine: 92,
								EndLine:   92,
								StartPos:  1938,
								EndPos:    1942,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  1938,
									EndPos:    1940,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 92,
										EndLine:   92,
										StartPos:  1938,
										EndPos:    1940,
									},
								},
								Value: []byte("$i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  1944,
								EndPos:    1948,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  1944,
									EndPos:    1946,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 92,
										EndLine:   92,
										StartPos:  1944,
										EndPos:    1946,
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
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  1960,
						EndPos:    1981,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  1969,
							EndPos:    1971,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 93,
								EndLine:   93,
								StartPos:  1969,
								EndPos:    1971,
							},
						},
						Value: []byte("$a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  1975,
							EndPos:    1977,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 93,
								EndLine:   93,
								StartPos:  1975,
								EndPos:    1977,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  1979,
							EndPos:    1981,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  1982,
						EndPos:    2014,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  1991,
							EndPos:    1993,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  1991,
								EndPos:    1993,
							},
						},
						Value: []byte("$a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  1997,
							EndPos:    1999,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  1997,
								EndPos:    1999,
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
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2015,
						EndPos:    2042,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2024,
							EndPos:    2026,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2024,
								EndPos:    2026,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2030,
							EndPos:    2032,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2030,
								EndPos:    2032,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2036,
							EndPos:    2038,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2036,
								EndPos:    2038,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2040,
							EndPos:    2042,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2043,
						EndPos:    2071,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2052,
							EndPos:    2054,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2052,
								EndPos:    2054,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2058,
							EndPos:    2060,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2058,
								EndPos:    2060,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2064,
							EndPos:    2067,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2065,
								EndPos:    2067,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 96,
									EndLine:   96,
									StartPos:  2065,
									EndPos:    2067,
								},
							},
							Value: []byte("$v"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2069,
							EndPos:    2071,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2072,
						EndPos:    2105,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2081,
							EndPos:    2083,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2081,
								EndPos:    2083,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2087,
							EndPos:    2089,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2087,
								EndPos:    2089,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2093,
							EndPos:    2101,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2098,
									EndPos:    2100,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2098,
										EndPos:    2100,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 97,
											EndLine:   97,
											StartPos:  2098,
											EndPos:    2100,
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
							StartLine: 97,
							EndLine:   97,
							StartPos:  2103,
							EndPos:    2105,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2106,
						EndPos:    2135,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2115,
							EndPos:    2117,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2115,
								EndPos:    2117,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2121,
							EndPos:    2123,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2121,
								EndPos:    2123,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprShortList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2127,
							EndPos:    2131,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2128,
									EndPos:    2130,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 98,
										EndLine:   98,
										StartPos:  2128,
										EndPos:    2130,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 98,
											EndLine:   98,
											StartPos:  2128,
											EndPos:    2130,
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
							StartLine: 98,
							EndLine:   98,
							StartPos:  2133,
							EndPos:    2135,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2136,
						EndPos:    2153,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 99,
							EndLine:   99,
							StartPos:  2145,
							EndPos:    2148,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2154,
						EndPos:    2178,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2163,
							EndPos:    2166,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 100,
								EndLine:   100,
								StartPos:  2170,
								EndPos:    2177,
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2179,
						EndPos:    2206,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2189,
							EndPos:    2192,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2196,
								EndPos:    2205,
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2203,
									EndPos:    2204,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2207,
						EndPos:    2231,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2217,
							EndPos:    2220,
						},
					},
					Value: []byte("foo"),
				},
				ReturnType: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2224,
							EndPos:    2228,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 102,
									EndLine:   102,
									StartPos:  2224,
									EndPos:    2228,
								},
							},
							Value: []byte("void"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtGlobal{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2232,
						EndPos:    2246,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2239,
								EndPos:    2241,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2239,
									EndPos:    2241,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2243,
								EndPos:    2245,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2243,
									EndPos:    2245,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtLabel{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2247,
						EndPos:    2249,
					},
				},
				LabelName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2247,
							EndPos:    2248,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtGoto{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2250,
						EndPos:    2257,
					},
				},
				Label: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 105,
							EndLine:   105,
							StartPos:  2255,
							EndPos:    2256,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2258,
						EndPos:    2268,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2262,
							EndPos:    2264,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2262,
								EndPos:    2264,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2266,
							EndPos:    2268,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2269,
						EndPos:    2294,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2273,
							EndPos:    2275,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2273,
								EndPos:    2275,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2277,
							EndPos:    2279,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2280,
								EndPos:    2294,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2288,
									EndPos:    2290,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 107,
										EndLine:   107,
										StartPos:  2288,
										EndPos:    2290,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2292,
									EndPos:    2294,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2295,
						EndPos:    2313,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2299,
							EndPos:    2301,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2299,
								EndPos:    2301,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2303,
							EndPos:    2305,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2306,
							EndPos:    2313,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2311,
								EndPos:    2313,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2314,
						EndPos:    2362,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2318,
							EndPos:    2320,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2318,
								EndPos:    2320,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2322,
							EndPos:    2324,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2325,
								EndPos:    2339,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2333,
									EndPos:    2335,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 109,
										EndLine:   109,
										StartPos:  2333,
										EndPos:    2335,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2337,
									EndPos:    2339,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2340,
								EndPos:    2354,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2348,
									EndPos:    2350,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 109,
										EndLine:   109,
										StartPos:  2348,
										EndPos:    2350,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2352,
									EndPos:    2354,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2355,
							EndPos:    2362,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2360,
								EndPos:    2362,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2363,
						EndPos:    2412,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2367,
							EndPos:    2369,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2367,
								EndPos:    2369,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2371,
							EndPos:    2373,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2374,
								EndPos:    2388,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2382,
									EndPos:    2384,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2382,
										EndPos:    2384,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2386,
									EndPos:    2388,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2389,
							EndPos:    2412,
						},
					},
					Stmt: &ast.StmtIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2394,
								EndPos:    2412,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2398,
									EndPos:    2400,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2398,
										EndPos:    2400,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2402,
									EndPos:    2404,
								},
							},
							Stmts: []ast.Vertex{},
						},
						Else: &ast.StmtElse{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2405,
									EndPos:    2412,
								},
							},
							Stmt: &ast.StmtStmtList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2410,
										EndPos:    2412,
									},
								},
								Stmts: []ast.Vertex{},
							},
						},
					},
				},
			},
			&ast.StmtNop{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2413,
						EndPos:    2415,
					},
				},
			},
			&ast.StmtInlineHtml{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2415,
						EndPos:    2428,
					},
				},
				Value: []byte(" <div></div> "),
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2431,
						EndPos:    2447,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 112,
							EndLine:   112,
							StartPos:  2441,
							EndPos:    2444,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2448,
						EndPos:    2476,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2458,
							EndPos:    2461,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2462,
							EndPos:    2473,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 113,
									EndLine:   113,
									StartPos:  2470,
									EndPos:    2473,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2470,
											EndPos:    2473,
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
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2477,
						EndPos:    2510,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2487,
							EndPos:    2490,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2491,
							EndPos:    2507,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2499,
									EndPos:    2502,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2499,
											EndPos:    2502,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2504,
									EndPos:    2507,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2504,
											EndPos:    2507,
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
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2511,
						EndPos:    2525,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 115,
							EndLine:   115,
							StartPos:  2521,
							EndPos:    2524,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2521,
									EndPos:    2524,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 116,
						EndLine:   116,
						StartPos:  2526,
						EndPos:    2542,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 116,
							EndLine:   116,
							StartPos:  2536,
							EndPos:    2539,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2536,
									EndPos:    2539,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 117,
						EndLine:   117,
						StartPos:  2543,
						EndPos:    2555,
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 118,
						EndLine:   118,
						StartPos:  2556,
						EndPos:    2575,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2562,
							EndPos:    2565,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 118,
								EndLine:   118,
								StartPos:  2567,
								EndPos:    2574,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 118,
										EndLine:   118,
										StartPos:  2567,
										EndPos:    2570,
									},
								},
								Value: []byte("var"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 118,
										EndLine:   118,
										StartPos:  2571,
										EndPos:    2573,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 118,
											EndLine:   118,
											StartPos:  2571,
											EndPos:    2573,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 118,
												EndLine:   118,
												StartPos:  2571,
												EndPos:    2573,
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
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 119,
						EndLine:   119,
						StartPos:  2576,
						EndPos:    2613,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 119,
							EndLine:   119,
							StartPos:  2582,
							EndPos:    2585,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 119,
								EndLine:   119,
								StartPos:  2587,
								EndPos:    2612,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2587,
										EndPos:    2593,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2594,
										EndPos:    2600,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2601,
										EndPos:    2603,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2601,
											EndPos:    2603,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 119,
												EndLine:   119,
												StartPos:  2601,
												EndPos:    2603,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2605,
										EndPos:    2611,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2605,
											EndPos:    2607,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 119,
												EndLine:   119,
												StartPos:  2605,
												EndPos:    2607,
											},
										},
										Value: []byte("$b"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2610,
											EndPos:    2611,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 120,
						EndLine:   120,
						StartPos:  2614,
						EndPos:    2632,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2621,
								EndPos:    2623,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2621,
									EndPos:    2623,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2621,
										EndPos:    2623,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2625,
								EndPos:    2631,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2625,
									EndPos:    2627,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2625,
										EndPos:    2627,
									},
								},
								Value: []byte("$b"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2630,
									EndPos:    2631,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 122,
						EndLine:   126,
						StartPos:  2634,
						EndPos:    2694,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 122,
							EndLine:   122,
							StartPos:  2642,
							EndPos:    2643,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 123,
							EndLine:   -1,
							StartPos:  2651,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 123,
									EndLine:   -1,
									StartPos:  2651,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 123,
										EndLine:   123,
										StartPos:  2656,
										EndPos:    2657,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtDefault{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 124,
									EndLine:   -1,
									StartPos:  2663,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 125,
									EndLine:   -1,
									StartPos:  2676,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 125,
										EndLine:   125,
										StartPos:  2681,
										EndPos:    2682,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 128,
						EndLine:   131,
						StartPos:  2696,
						EndPos:    2744,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 128,
							EndLine:   128,
							StartPos:  2704,
							EndPos:    2705,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 129,
							EndLine:   -1,
							StartPos:  2714,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 129,
									EndLine:   -1,
									StartPos:  2714,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 129,
										EndLine:   129,
										StartPos:  2719,
										EndPos:    2720,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   -1,
									StartPos:  2726,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 130,
										EndLine:   130,
										StartPos:  2731,
										EndPos:    2732,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 133,
						EndLine:   136,
						StartPos:  2746,
						EndPos:    2798,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 133,
							EndLine:   133,
							StartPos:  2754,
							EndPos:    2755,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 133,
							EndLine:   136,
							StartPos:  2757,
							EndPos:    2798,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 134,
									EndLine:   134,
									StartPos:  2763,
									EndPos:    2777,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 134,
										EndLine:   134,
										StartPos:  2768,
										EndPos:    2769,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 134,
											EndLine:   134,
											StartPos:  2771,
											EndPos:    2777,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  2782,
									EndPos:    2796,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  2787,
										EndPos:    2788,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 135,
											EndLine:   135,
											StartPos:  2790,
											EndPos:    2796,
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 138,
						EndLine:   141,
						StartPos:  2800,
						EndPos:    2853,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   138,
							StartPos:  2808,
							EndPos:    2809,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   141,
							StartPos:  2811,
							EndPos:    2853,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 139,
									EndLine:   139,
									StartPos:  2818,
									EndPos:    2832,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 139,
										EndLine:   139,
										StartPos:  2823,
										EndPos:    2824,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 139,
											EndLine:   139,
											StartPos:  2826,
											EndPos:    2832,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 140,
									EndLine:   140,
									StartPos:  2837,
									EndPos:    2851,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 140,
										EndLine:   140,
										StartPos:  2842,
										EndPos:    2843,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 140,
											EndLine:   140,
											StartPos:  2845,
											EndPos:    2851,
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtThrow{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 143,
						EndLine:   143,
						StartPos:  2855,
						EndPos:    2864,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  2861,
							EndPos:    2863,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 143,
								EndLine:   143,
								StartPos:  2861,
								EndPos:    2863,
							},
						},
						Value: []byte("$e"),
					},
				},
			},
			&ast.StmtTrait{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 145,
						EndLine:   145,
						StartPos:  2866,
						EndPos:    2878,
					},
				},
				TraitName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 145,
							EndLine:   145,
							StartPos:  2872,
							EndPos:    2875,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 146,
						EndLine:   146,
						StartPos:  2879,
						EndPos:    2901,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 146,
							EndLine:   146,
							StartPos:  2885,
							EndPos:    2888,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 146,
								EndLine:   146,
								StartPos:  2891,
								EndPos:    2899,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 146,
										EndLine:   146,
										StartPos:  2895,
										EndPos:    2898,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  2895,
												EndPos:    2898,
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
									StartLine: 146,
									EndLine:   146,
									StartPos:  2898,
									EndPos:    2899,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 147,
						EndLine:   147,
						StartPos:  2902,
						EndPos:    2931,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 147,
							EndLine:   147,
							StartPos:  2908,
							EndPos:    2911,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  2914,
								EndPos:    2929,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  2918,
										EndPos:    2921,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 147,
												EndLine:   147,
												StartPos:  2918,
												EndPos:    2921,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  2923,
										EndPos:    2926,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 147,
												EndLine:   147,
												StartPos:  2923,
												EndPos:    2926,
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
									StartLine: 147,
									EndLine:   147,
									StartPos:  2927,
									EndPos:    2929,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 148,
						EndLine:   148,
						StartPos:  2932,
						EndPos:    2978,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 148,
							EndLine:   148,
							StartPos:  2938,
							EndPos:    2941,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  2944,
								EndPos:    2976,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  2948,
										EndPos:    2951,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  2948,
												EndPos:    2951,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  2953,
										EndPos:    2956,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  2953,
												EndPos:    2956,
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
									StartLine: 148,
									EndLine:   148,
									StartPos:  2957,
									EndPos:    2976,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  2959,
											EndPos:    2973,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  2959,
												EndPos:    2962,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 148,
													EndLine:   148,
													StartPos:  2959,
													EndPos:    2962,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  2966,
												EndPos:    2973,
											},
										},
										Value: []byte("include"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 149,
						EndLine:   149,
						StartPos:  2979,
						EndPos:    3024,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  2985,
							EndPos:    2988,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  2991,
								EndPos:    3022,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  2995,
										EndPos:    2998,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  2995,
												EndPos:    2998,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3000,
										EndPos:    3003,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3000,
												EndPos:    3003,
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
									StartLine: 149,
									EndLine:   149,
									StartPos:  3004,
									EndPos:    3022,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3006,
											EndPos:    3019,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3006,
												EndPos:    3009,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 149,
													EndLine:   149,
													StartPos:  3006,
													EndPos:    3009,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3013,
												EndPos:    3019,
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
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 150,
						EndLine:   150,
						StartPos:  3025,
						EndPos:    3074,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3031,
							EndPos:    3034,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3037,
								EndPos:    3072,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3041,
										EndPos:    3044,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3041,
												EndPos:    3044,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3046,
										EndPos:    3049,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3046,
												EndPos:    3049,
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
									StartLine: 150,
									EndLine:   150,
									StartPos:  3050,
									EndPos:    3072,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3052,
											EndPos:    3069,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3052,
												EndPos:    3055,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 150,
													EndLine:   150,
													StartPos:  3052,
													EndPos:    3055,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3059,
												EndPos:    3065,
											},
										},
										Value: []byte("public"),
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3066,
												EndPos:    3069,
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
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 151,
						EndLine:   151,
						StartPos:  3075,
						EndPos:    3152,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3081,
							EndPos:    3084,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 151,
								EndLine:   151,
								StartPos:  3087,
								EndPos:    3150,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3091,
										EndPos:    3094,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3091,
												EndPos:    3094,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3096,
										EndPos:    3099,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3096,
												EndPos:    3099,
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
									StartLine: 151,
									EndLine:   151,
									StartPos:  3100,
									EndPos:    3150,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUsePrecedence{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3102,
											EndPos:    3130,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3102,
												EndPos:    3110,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3102,
													EndPos:    3105,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3102,
															EndPos:    3105,
														},
													},
													Value: []byte("Bar"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3107,
													EndPos:    3110,
												},
											},
											Value: []byte("one"),
										},
									},
									Insteadof: []ast.Vertex{
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3121,
													EndPos:    3124,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3121,
															EndPos:    3124,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3126,
													EndPos:    3130,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3126,
															EndPos:    3130,
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
											StartLine: 151,
											EndLine:   151,
											StartPos:  3132,
											EndPos:    3147,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3132,
												EndPos:    3140,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3132,
													EndPos:    3135,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3132,
															EndPos:    3135,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3137,
													EndPos:    3140,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3144,
												EndPos:    3147,
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
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 153,
						EndLine:   -1,
						StartPos:  3154,
						EndPos:    -1,
					},
				},
				Stmts:   []ast.Vertex{},
				Catches: []ast.Vertex{},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 154,
						EndLine:   154,
						StartPos:  3161,
						EndPos:    3191,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3168,
								EndPos:    3191,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 154,
										EndLine:   154,
										StartPos:  3175,
										EndPos:    3184,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 154,
												EndLine:   154,
												StartPos:  3175,
												EndPos:    3184,
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
									StartLine: 154,
									EndLine:   154,
									StartPos:  3185,
									EndPos:    3187,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 154,
										EndLine:   154,
										StartPos:  3185,
										EndPos:    3187,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 155,
						EndLine:   155,
						StartPos:  3192,
						EndPos:    3239,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3199,
								EndPos:    3239,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3206,
										EndPos:    3215,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 155,
												EndLine:   155,
												StartPos:  3206,
												EndPos:    3215,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3216,
										EndPos:    3232,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 155,
												EndLine:   155,
												StartPos:  3216,
												EndPos:    3232,
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
									StartLine: 155,
									EndLine:   155,
									StartPos:  3233,
									EndPos:    3235,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3233,
										EndPos:    3235,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 156,
						EndLine:   156,
						StartPos:  3240,
						EndPos:    3301,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3247,
								EndPos:    3270,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3254,
										EndPos:    3263,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 156,
												EndLine:   156,
												StartPos:  3254,
												EndPos:    3263,
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
									StartLine: 156,
									EndLine:   156,
									StartPos:  3264,
									EndPos:    3266,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3264,
										EndPos:    3266,
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
								StartLine: 156,
								EndLine:   156,
								StartPos:  3271,
								EndPos:    3301,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3278,
										EndPos:    3294,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 156,
												EndLine:   156,
												StartPos:  3278,
												EndPos:    3294,
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
									StartLine: 156,
									EndLine:   156,
									StartPos:  3295,
									EndPos:    3297,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3295,
										EndPos:    3297,
									},
								},
								Value: []byte("$e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 157,
						EndLine:   157,
						StartPos:  3302,
						EndPos:    3343,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3309,
								EndPos:    3332,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3316,
										EndPos:    3325,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 157,
												EndLine:   157,
												StartPos:  3316,
												EndPos:    3325,
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
									StartLine: 157,
									EndLine:   157,
									StartPos:  3326,
									EndPos:    3328,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3326,
										EndPos:    3328,
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
							StartLine: 157,
							EndLine:   157,
							StartPos:  3333,
							EndPos:    3343,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 159,
						EndLine:   159,
						StartPos:  3345,
						EndPos:    3360,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3351,
								EndPos:    3353,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3351,
									EndPos:    3353,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3355,
								EndPos:    3357,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3355,
									EndPos:    3357,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 161,
						EndLine:   161,
						StartPos:  3362,
						EndPos:    3370,
					},
				},
				UseList: &ast.StmtUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3366,
							EndPos:    3369,
						},
					},
					UseDeclarations: []ast.Vertex{
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 161,
									EndLine:   161,
									StartPos:  3366,
									EndPos:    3369,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 161,
										EndLine:   161,
										StartPos:  3366,
										EndPos:    3369,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 161,
												EndLine:   161,
												StartPos:  3366,
												EndPos:    3369,
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
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 162,
						EndLine:   162,
						StartPos:  3371,
						EndPos:    3380,
					},
				},
				UseList: &ast.StmtUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3376,
							EndPos:    3379,
						},
					},
					UseDeclarations: []ast.Vertex{
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3376,
									EndPos:    3379,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3376,
										EndPos:    3379,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 162,
												EndLine:   162,
												StartPos:  3376,
												EndPos:    3379,
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
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 163,
						EndLine:   163,
						StartPos:  3381,
						EndPos:    3397,
					},
				},
				UseList: &ast.StmtUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 163,
							EndLine:   163,
							StartPos:  3386,
							EndPos:    3396,
						},
					},
					UseDeclarations: []ast.Vertex{
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 163,
									EndLine:   163,
									StartPos:  3386,
									EndPos:    3396,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 163,
										EndLine:   163,
										StartPos:  3386,
										EndPos:    3389,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 163,
												EndLine:   163,
												StartPos:  3386,
												EndPos:    3389,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							Alias: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 163,
										EndLine:   163,
										StartPos:  3393,
										EndPos:    3396,
									},
								},
								Value: []byte("Bar"),
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 164,
						EndLine:   164,
						StartPos:  3398,
						EndPos:    3411,
					},
				},
				UseList: &ast.StmtUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3402,
							EndPos:    3410,
						},
					},
					UseDeclarations: []ast.Vertex{
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 164,
									EndLine:   164,
									StartPos:  3402,
									EndPos:    3405,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3402,
										EndPos:    3405,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 164,
												EndLine:   164,
												StartPos:  3402,
												EndPos:    3405,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
						},
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 164,
									EndLine:   164,
									StartPos:  3407,
									EndPos:    3410,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3407,
										EndPos:    3410,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 164,
												EndLine:   164,
												StartPos:  3407,
												EndPos:    3410,
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
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 165,
						EndLine:   165,
						StartPos:  3412,
						EndPos:    3432,
					},
				},
				UseList: &ast.StmtUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3416,
							EndPos:    3431,
						},
					},
					UseDeclarations: []ast.Vertex{
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3416,
									EndPos:    3419,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3416,
										EndPos:    3419,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 165,
												EndLine:   165,
												StartPos:  3416,
												EndPos:    3419,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
						},
						&ast.StmtUseDeclaration{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3421,
									EndPos:    3431,
								},
							},
							Use: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3421,
										EndPos:    3424,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 165,
												EndLine:   165,
												StartPos:  3421,
												EndPos:    3424,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							Alias: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3428,
										EndPos:    3431,
									},
								},
								Value: []byte("Baz"),
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 166,
						EndLine:   166,
						StartPos:  3433,
						EndPos:    3456,
					},
				},
				UseList: &ast.StmtUseType{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3437,
							EndPos:    3455,
						},
					},
					Type: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3437,
								EndPos:    3445,
							},
						},
						Value: []byte("function"),
					},
					Use: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3446,
								EndPos:    3455,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 166,
										EndLine:   166,
										StartPos:  3446,
										EndPos:    3449,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 166,
											EndLine:   166,
											StartPos:  3446,
											EndPos:    3449,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 166,
													EndLine:   166,
													StartPos:  3446,
													EndPos:    3449,
												},
											},
											Value: []byte("Foo"),
										},
									},
								},
							},
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 166,
										EndLine:   166,
										StartPos:  3452,
										EndPos:    3455,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 166,
											EndLine:   166,
											StartPos:  3452,
											EndPos:    3455,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 166,
													EndLine:   166,
													StartPos:  3452,
													EndPos:    3455,
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
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3457,
						EndPos:    3494,
					},
				},
				UseList: &ast.StmtUseType{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3461,
							EndPos:    3493,
						},
					},
					Type: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3461,
								EndPos:    3469,
							},
						},
						Value: []byte("function"),
					},
					Use: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3470,
								EndPos:    3493,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3470,
										EndPos:    3480,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 167,
											EndLine:   167,
											StartPos:  3470,
											EndPos:    3473,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 167,
													EndLine:   167,
													StartPos:  3470,
													EndPos:    3473,
												},
											},
											Value: []byte("Foo"),
										},
									},
								},
								Alias: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 167,
											EndLine:   167,
											StartPos:  3477,
											EndPos:    3480,
										},
									},
									Value: []byte("foo"),
								},
							},
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3483,
										EndPos:    3493,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 167,
											EndLine:   167,
											StartPos:  3483,
											EndPos:    3486,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 167,
													EndLine:   167,
													StartPos:  3483,
													EndPos:    3486,
												},
											},
											Value: []byte("Bar"),
										},
									},
								},
								Alias: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 167,
											EndLine:   167,
											StartPos:  3490,
											EndPos:    3493,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 168,
						EndLine:   168,
						StartPos:  3495,
						EndPos:    3515,
					},
				},
				UseList: &ast.StmtUseType{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3499,
							EndPos:    3514,
						},
					},
					Type: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3499,
								EndPos:    3504,
							},
						},
						Value: []byte("const"),
					},
					Use: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3505,
								EndPos:    3514,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3505,
										EndPos:    3508,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3505,
											EndPos:    3508,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 168,
													EndLine:   168,
													StartPos:  3505,
													EndPos:    3508,
												},
											},
											Value: []byte("Foo"),
										},
									},
								},
							},
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3511,
										EndPos:    3514,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3511,
											EndPos:    3514,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 168,
													EndLine:   168,
													StartPos:  3511,
													EndPos:    3514,
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
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 169,
						EndLine:   169,
						StartPos:  3516,
						EndPos:    3550,
					},
				},
				UseList: &ast.StmtUseType{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3520,
							EndPos:    3549,
						},
					},
					Type: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3520,
								EndPos:    3525,
							},
						},
						Value: []byte("const"),
					},
					Use: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3526,
								EndPos:    3549,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3526,
										EndPos:    3536,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3526,
											EndPos:    3529,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 169,
													EndLine:   169,
													StartPos:  3526,
													EndPos:    3529,
												},
											},
											Value: []byte("Foo"),
										},
									},
								},
								Alias: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3533,
											EndPos:    3536,
										},
									},
									Value: []byte("foo"),
								},
							},
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3539,
										EndPos:    3549,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3539,
											EndPos:    3542,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 169,
													EndLine:   169,
													StartPos:  3539,
													EndPos:    3542,
												},
											},
											Value: []byte("Bar"),
										},
									},
								},
								Alias: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3546,
											EndPos:    3549,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 171,
						EndLine:   171,
						StartPos:  3552,
						EndPos:    3572,
					},
				},
				UseList: &ast.StmtGroupUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3556,
							EndPos:    3571,
						},
					},
					Prefix: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3557,
								EndPos:    3560,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 171,
										EndLine:   171,
										StartPos:  3557,
										EndPos:    3560,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					UseList: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3562,
								EndPos:    3570,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 171,
										EndLine:   171,
										StartPos:  3562,
										EndPos:    3565,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 171,
											EndLine:   171,
											StartPos:  3562,
											EndPos:    3565,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 171,
													EndLine:   171,
													StartPos:  3562,
													EndPos:    3565,
												},
											},
											Value: []byte("Bar"),
										},
									},
								},
							},
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 171,
										EndLine:   171,
										StartPos:  3567,
										EndPos:    3570,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 171,
											EndLine:   171,
											StartPos:  3567,
											EndPos:    3570,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 171,
													EndLine:   171,
													StartPos:  3567,
													EndPos:    3570,
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
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3573,
						EndPos:    3600,
					},
				},
				UseList: &ast.StmtGroupUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3577,
							EndPos:    3599,
						},
					},
					Prefix: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3577,
								EndPos:    3580,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3577,
										EndPos:    3580,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					UseList: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3582,
								EndPos:    3598,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3582,
										EndPos:    3585,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3582,
											EndPos:    3585,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 172,
													EndLine:   172,
													StartPos:  3582,
													EndPos:    3585,
												},
											},
											Value: []byte("Bar"),
										},
									},
								},
							},
							&ast.StmtUseDeclaration{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3587,
										EndPos:    3598,
									},
								},
								Use: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3587,
											EndPos:    3590,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 172,
													EndLine:   172,
													StartPos:  3587,
													EndPos:    3590,
												},
											},
											Value: []byte("Baz"),
										},
									},
								},
								Alias: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3594,
											EndPos:    3598,
										},
									},
									Value: []byte("quux"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 173,
						EndLine:   173,
						StartPos:  3601,
						EndPos:    3629,
					},
				},
				UseList: &ast.StmtUseType{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 173,
							EndLine:   173,
							StartPos:  3605,
							EndPos:    3628,
						},
					},
					Type: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3605,
								EndPos:    3613,
							},
						},
						Value: []byte("function"),
					},
					Use: &ast.StmtGroupUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3614,
								EndPos:    3628,
							},
						},
						Prefix: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 173,
									EndLine:   173,
									StartPos:  3614,
									EndPos:    3617,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 173,
											EndLine:   173,
											StartPos:  3614,
											EndPos:    3617,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						UseList: &ast.StmtUseList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 173,
									EndLine:   173,
									StartPos:  3619,
									EndPos:    3627,
								},
							},
							UseDeclarations: []ast.Vertex{
								&ast.StmtUseDeclaration{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 173,
											EndLine:   173,
											StartPos:  3619,
											EndPos:    3622,
										},
									},
									Use: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 173,
												EndLine:   173,
												StartPos:  3619,
												EndPos:    3622,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 173,
														EndLine:   173,
														StartPos:  3619,
														EndPos:    3622,
													},
												},
												Value: []byte("Bar"),
											},
										},
									},
								},
								&ast.StmtUseDeclaration{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 173,
											EndLine:   173,
											StartPos:  3624,
											EndPos:    3627,
										},
									},
									Use: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 173,
												EndLine:   173,
												StartPos:  3624,
												EndPos:    3627,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 173,
														EndLine:   173,
														StartPos:  3624,
														EndPos:    3627,
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
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3630,
						EndPos:    3656,
					},
				},
				UseList: &ast.StmtUseType{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3634,
							EndPos:    3655,
						},
					},
					Type: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3634,
								EndPos:    3639,
							},
						},
						Value: []byte("const"),
					},
					Use: &ast.StmtGroupUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3640,
								EndPos:    3655,
							},
						},
						Prefix: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 174,
									EndLine:   174,
									StartPos:  3641,
									EndPos:    3644,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 174,
											EndLine:   174,
											StartPos:  3641,
											EndPos:    3644,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						UseList: &ast.StmtUseList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 174,
									EndLine:   174,
									StartPos:  3646,
									EndPos:    3654,
								},
							},
							UseDeclarations: []ast.Vertex{
								&ast.StmtUseDeclaration{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 174,
											EndLine:   174,
											StartPos:  3646,
											EndPos:    3649,
										},
									},
									Use: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 174,
												EndLine:   174,
												StartPos:  3646,
												EndPos:    3649,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 174,
														EndLine:   174,
														StartPos:  3646,
														EndPos:    3649,
													},
												},
												Value: []byte("Bar"),
											},
										},
									},
								},
								&ast.StmtUseDeclaration{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 174,
											EndLine:   174,
											StartPos:  3651,
											EndPos:    3654,
										},
									},
									Use: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 174,
												EndLine:   174,
												StartPos:  3651,
												EndPos:    3654,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 174,
														EndLine:   174,
														StartPos:  3651,
														EndPos:    3654,
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
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3657,
						EndPos:    3691,
					},
				},
				UseList: &ast.StmtGroupUseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3661,
							EndPos:    3690,
						},
					},
					Prefix: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3661,
								EndPos:    3664,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3661,
										EndPos:    3664,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					UseList: &ast.StmtUseList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3666,
								EndPos:    3689,
							},
						},
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseType{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3666,
										EndPos:    3675,
									},
								},
								Type: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 175,
											EndLine:   175,
											StartPos:  3666,
											EndPos:    3671,
										},
									},
									Value: []byte("const"),
								},
								Use: &ast.StmtUseDeclaration{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 175,
											EndLine:   175,
											StartPos:  3672,
											EndPos:    3675,
										},
									},
									Use: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 175,
												EndLine:   175,
												StartPos:  3672,
												EndPos:    3675,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 175,
														EndLine:   175,
														StartPos:  3672,
														EndPos:    3675,
													},
												},
												Value: []byte("Bar"),
											},
										},
									},
								},
							},
							&ast.StmtUseType{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3677,
										EndPos:    3689,
									},
								},
								Type: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 175,
											EndLine:   175,
											StartPos:  3677,
											EndPos:    3685,
										},
									},
									Value: []byte("function"),
								},
								Use: &ast.StmtUseDeclaration{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 175,
											EndLine:   175,
											StartPos:  3686,
											EndPos:    3689,
										},
									},
									Use: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 175,
												EndLine:   175,
												StartPos:  3686,
												EndPos:    3689,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 175,
														EndLine:   175,
														StartPos:  3686,
														EndPos:    3689,
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
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 177,
						EndLine:   177,
						StartPos:  3693,
						EndPos:    3699,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3693,
							EndPos:    3698,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 177,
								EndLine:   177,
								StartPos:  3693,
								EndPos:    3695,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 177,
									EndLine:   177,
									StartPos:  3693,
									EndPos:    3695,
								},
							},
							Value: []byte("$a"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 177,
								EndLine:   177,
								StartPos:  3696,
								EndPos:    3697,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3700,
						EndPos:    3709,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3700,
							EndPos:    3708,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3700,
								EndPos:    3705,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3700,
									EndPos:    3702,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3700,
										EndPos:    3702,
									},
								},
								Value: []byte("$a"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3703,
									EndPos:    3704,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3706,
								EndPos:    3707,
							},
						},
						Value: []byte("2"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  3710,
						EndPos:    3718,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  3710,
							EndPos:    3717,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  3719,
						EndPos:    3728,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 180,
							EndLine:   180,
							StartPos:  3719,
							EndPos:    3727,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  3725,
									EndPos:    3726,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 180,
										EndLine:   180,
										StartPos:  3725,
										EndPos:    3726,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  3729,
						EndPos:    3747,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 181,
							EndLine:   181,
							StartPos:  3729,
							EndPos:    3746,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  3735,
									EndPos:    3739,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  3735,
										EndPos:    3736,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  3738,
										EndPos:    3739,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  3741,
									EndPos:    3744,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  3741,
										EndPos:    3744,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 181,
											EndLine:   181,
											StartPos:  3742,
											EndPos:    3744,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 181,
												EndLine:   181,
												StartPos:  3742,
												EndPos:    3744,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 182,
						EndLine:   182,
						StartPos:  3748,
						EndPos:    3752,
					},
				},
				Expr: &ast.ExprBitwiseNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 182,
							EndLine:   182,
							StartPos:  3748,
							EndPos:    3751,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  3749,
								EndPos:    3751,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 182,
									EndLine:   182,
									StartPos:  3749,
									EndPos:    3751,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 183,
						EndLine:   183,
						StartPos:  3753,
						EndPos:    3757,
					},
				},
				Expr: &ast.ExprBooleanNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  3753,
							EndPos:    3756,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 183,
								EndLine:   183,
								StartPos:  3754,
								EndPos:    3756,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 183,
									EndLine:   183,
									StartPos:  3754,
									EndPos:    3756,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 185,
						EndLine:   185,
						StartPos:  3759,
						EndPos:    3768,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  3759,
							EndPos:    3767,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  3759,
								EndPos:    3762,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 185,
										EndLine:   185,
										StartPos:  3759,
										EndPos:    3762,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  3764,
								EndPos:    3767,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 186,
						EndLine:   186,
						StartPos:  3769,
						EndPos:    3779,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  3769,
							EndPos:    3778,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  3769,
								EndPos:    3773,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 186,
									EndLine:   186,
									StartPos:  3769,
									EndPos:    3773,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  3775,
								EndPos:    3778,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  3780,
						EndPos:    3790,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  3780,
							EndPos:    3788,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  3786,
								EndPos:    3788,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 187,
									EndLine:   187,
									StartPos:  3786,
									EndPos:    3788,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 188,
						EndLine:   188,
						StartPos:  3791,
						EndPos:    3800,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  3791,
							EndPos:    3799,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  3797,
								EndPos:    3799,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 188,
									EndLine:   188,
									StartPos:  3797,
									EndPos:    3799,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 189,
						EndLine:   189,
						StartPos:  3801,
						EndPos:    3814,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 189,
							EndLine:   189,
							StartPos:  3801,
							EndPos:    3813,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  3815,
						EndPos:    3849,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  3815,
							EndPos:    3848,
						},
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  3824,
									EndPos:    3826,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  3824,
										EndPos:    3826,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  3824,
											EndPos:    3826,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  3828,
									EndPos:    3830,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  3828,
										EndPos:    3830,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  3828,
											EndPos:    3830,
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
								StartLine: 190,
								EndLine:   190,
								StartPos:  3832,
								EndPos:    3845,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  3837,
										EndPos:    3839,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  3837,
											EndPos:    3839,
										},
									},
									Value: []byte("$c"),
								},
							},
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  3841,
										EndPos:    3844,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  3842,
											EndPos:    3844,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 190,
												EndLine:   190,
												StartPos:  3842,
												EndPos:    3844,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 191,
						EndLine:   191,
						StartPos:  3850,
						EndPos:    3870,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 191,
							EndLine:   191,
							StartPos:  3850,
							EndPos:    3869,
						},
					},
					ReturnType: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 191,
								EndLine:   191,
								StartPos:  3862,
								EndPos:    3866,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 191,
										EndLine:   191,
										StartPos:  3862,
										EndPos:    3866,
									},
								},
								Value: []byte("void"),
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  3871,
						EndPos:    3875,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  3871,
							EndPos:    3874,
						},
					},
					Const: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 192,
								EndLine:   192,
								StartPos:  3871,
								EndPos:    3874,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 192,
										EndLine:   192,
										StartPos:  3871,
										EndPos:    3874,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 193,
						EndLine:   193,
						StartPos:  3876,
						EndPos:    3890,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  3876,
							EndPos:    3889,
						},
					},
					Const: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 193,
								EndLine:   193,
								StartPos:  3876,
								EndPos:    3889,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 193,
										EndLine:   193,
										StartPos:  3886,
										EndPos:    3889,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 194,
						EndLine:   194,
						StartPos:  3891,
						EndPos:    3896,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  3891,
							EndPos:    3895,
						},
					},
					Const: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  3891,
								EndPos:    3895,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  3892,
										EndPos:    3895,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 196,
						EndLine:   196,
						StartPos:  3898,
						EndPos:    3908,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  3898,
							EndPos:    3907,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  3904,
								EndPos:    3906,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  3904,
									EndPos:    3906,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 197,
						EndLine:   197,
						StartPos:  3909,
						EndPos:    3913,
					},
				},
				Expr: &ast.ExprErrorSuppress{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 197,
							EndLine:   197,
							StartPos:  3909,
							EndPos:    3912,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 197,
								EndLine:   197,
								StartPos:  3910,
								EndPos:    3912,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 197,
									EndLine:   197,
									StartPos:  3910,
									EndPos:    3912,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  3914,
						EndPos:    3923,
					},
				},
				Expr: &ast.ExprEval{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  3914,
							EndPos:    3922,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  3919,
								EndPos:    3921,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 198,
									EndLine:   198,
									StartPos:  3919,
									EndPos:    3921,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 199,
						EndLine:   199,
						StartPos:  3924,
						EndPos:    3929,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 199,
							EndLine:   199,
							StartPos:  3924,
							EndPos:    3928,
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  3930,
						EndPos:    3939,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  3930,
							EndPos:    3938,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  3935,
								EndPos:    3937,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 200,
									EndLine:   200,
									StartPos:  3935,
									EndPos:    3937,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 201,
						EndLine:   201,
						StartPos:  3940,
						EndPos:    3944,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  3940,
							EndPos:    3943,
						},
					},
					Die: true,
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 202,
						EndLine:   202,
						StartPos:  3945,
						EndPos:    3953,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 202,
							EndLine:   202,
							StartPos:  3945,
							EndPos:    3952,
						},
					},
					Die: true,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 202,
								EndLine:   202,
								StartPos:  3949,
								EndPos:    3951,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 202,
									EndLine:   202,
									StartPos:  3949,
									EndPos:    3951,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 203,
						EndLine:   203,
						StartPos:  3954,
						EndPos:    3960,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  3954,
							EndPos:    3959,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  3954,
								EndPos:    3957,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 203,
										EndLine:   203,
										StartPos:  3954,
										EndPos:    3957,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  3957,
								EndPos:    3959,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 204,
						EndLine:   204,
						StartPos:  3961,
						EndPos:    3977,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  3961,
							EndPos:    3976,
						},
					},
					Function: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  3961,
								EndPos:    3974,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 204,
										EndLine:   204,
										StartPos:  3971,
										EndPos:    3974,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  3974,
								EndPos:    3976,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  3978,
						EndPos:    3985,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  3978,
							EndPos:    3984,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  3978,
								EndPos:    3982,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 205,
										EndLine:   205,
										StartPos:  3979,
										EndPos:    3982,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  3982,
								EndPos:    3984,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  3986,
						EndPos:    3993,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  3986,
							EndPos:    3992,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  3986,
								EndPos:    3990,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 206,
									EndLine:   206,
									StartPos:  3986,
									EndPos:    3990,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  3990,
								EndPos:    3992,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 208,
						EndLine:   208,
						StartPos:  3995,
						EndPos:    4000,
					},
				},
				Expr: &ast.ExprPostDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  3995,
							EndPos:    3999,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  3995,
								EndPos:    3997,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 208,
									EndLine:   208,
									StartPos:  3995,
									EndPos:    3997,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 209,
						EndLine:   209,
						StartPos:  4001,
						EndPos:    4006,
					},
				},
				Expr: &ast.ExprPostInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4001,
							EndPos:    4005,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4001,
								EndPos:    4003,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 209,
									EndLine:   209,
									StartPos:  4001,
									EndPos:    4003,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 210,
						EndLine:   210,
						StartPos:  4007,
						EndPos:    4012,
					},
				},
				Expr: &ast.ExprPreDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4007,
							EndPos:    4011,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4009,
								EndPos:    4011,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 210,
									EndLine:   210,
									StartPos:  4009,
									EndPos:    4011,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 211,
						EndLine:   211,
						StartPos:  4013,
						EndPos:    4018,
					},
				},
				Expr: &ast.ExprPreInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 211,
							EndLine:   211,
							StartPos:  4013,
							EndPos:    4017,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 211,
								EndLine:   211,
								StartPos:  4015,
								EndPos:    4017,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 211,
									EndLine:   211,
									StartPos:  4015,
									EndPos:    4017,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 213,
						EndLine:   213,
						StartPos:  4020,
						EndPos:    4031,
					},
				},
				Expr: &ast.ExprInclude{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 213,
							EndLine:   213,
							StartPos:  4020,
							EndPos:    4030,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 213,
								EndLine:   213,
								StartPos:  4028,
								EndPos:    4030,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 213,
									EndLine:   213,
									StartPos:  4028,
									EndPos:    4030,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4032,
						EndPos:    4048,
					},
				},
				Expr: &ast.ExprIncludeOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4032,
							EndPos:    4047,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4045,
								EndPos:    4047,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 214,
									EndLine:   214,
									StartPos:  4045,
									EndPos:    4047,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 215,
						EndLine:   215,
						StartPos:  4049,
						EndPos:    4060,
					},
				},
				Expr: &ast.ExprRequire{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4049,
							EndPos:    4059,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4057,
								EndPos:    4059,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 215,
									EndLine:   215,
									StartPos:  4057,
									EndPos:    4059,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4061,
						EndPos:    4077,
					},
				},
				Expr: &ast.ExprRequireOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4061,
							EndPos:    4076,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4074,
								EndPos:    4076,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 216,
									EndLine:   216,
									StartPos:  4074,
									EndPos:    4076,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 218,
						EndLine:   218,
						StartPos:  4079,
						EndPos:    4097,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 218,
							EndLine:   218,
							StartPos:  4079,
							EndPos:    4096,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 218,
								EndLine:   218,
								StartPos:  4079,
								EndPos:    4081,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 218,
									EndLine:   218,
									StartPos:  4079,
									EndPos:    4081,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 218,
								EndLine:   218,
								StartPos:  4093,
								EndPos:    4096,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 218,
										EndLine:   218,
										StartPos:  4093,
										EndPos:    4096,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 219,
						EndLine:   219,
						StartPos:  4098,
						EndPos:    4126,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4098,
							EndPos:    4125,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4098,
								EndPos:    4100,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4098,
									EndPos:    4100,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4112,
								EndPos:    4125,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 219,
										EndLine:   219,
										StartPos:  4122,
										EndPos:    4125,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 220,
						EndLine:   220,
						StartPos:  4127,
						EndPos:    4146,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4127,
							EndPos:    4145,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4127,
								EndPos:    4129,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 220,
									EndLine:   220,
									StartPos:  4127,
									EndPos:    4129,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4141,
								EndPos:    4145,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 220,
										EndLine:   220,
										StartPos:  4142,
										EndPos:    4145,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 222,
						EndLine:   222,
						StartPos:  4148,
						EndPos:    4162,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4148,
							EndPos:    4161,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4154,
									EndPos:    4156,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 222,
										EndLine:   222,
										StartPos:  4154,
										EndPos:    4156,
									},
								},
								Value: []byte("$a"),
							},
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4158,
									EndPos:    4160,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 222,
										EndLine:   222,
										StartPos:  4158,
										EndPos:    4160,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 223,
						EndLine:   223,
						StartPos:  4163,
						EndPos:    4177,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4163,
							EndPos:    4176,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4163,
								EndPos:    4171,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 223,
										EndLine:   223,
										StartPos:  4168,
										EndPos:    4170,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 223,
											EndLine:   223,
											StartPos:  4168,
											EndPos:    4170,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 223,
												EndLine:   223,
												StartPos:  4168,
												EndPos:    4170,
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
								StartLine: 223,
								EndLine:   223,
								StartPos:  4174,
								EndPos:    4176,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4174,
									EndPos:    4176,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4178,
						EndPos:    4194,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4178,
							EndPos:    4193,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4178,
								EndPos:    4188,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 224,
										EndLine:   224,
										StartPos:  4183,
										EndPos:    4187,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 224,
											EndLine:   224,
											StartPos:  4183,
											EndPos:    4187,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 224,
												EndLine:   224,
												StartPos:  4183,
												EndPos:    4185,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 224,
													EndLine:   224,
													StartPos:  4183,
													EndPos:    4185,
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
								StartLine: 224,
								EndLine:   224,
								StartPos:  4191,
								EndPos:    4193,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4191,
									EndPos:    4193,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4195,
						EndPos:    4215,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4195,
							EndPos:    4214,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4195,
								EndPos:    4209,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 225,
										EndLine:   225,
										StartPos:  4200,
										EndPos:    4208,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 225,
											EndLine:   225,
											StartPos:  4200,
											EndPos:    4208,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 225,
													EndLine:   225,
													StartPos:  4205,
													EndPos:    4207,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 225,
														EndLine:   225,
														StartPos:  4205,
														EndPos:    4207,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 225,
															EndLine:   225,
															StartPos:  4205,
															EndPos:    4207,
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
								StartLine: 225,
								EndLine:   225,
								StartPos:  4212,
								EndPos:    4214,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4212,
									EndPos:    4214,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 227,
						EndLine:   227,
						StartPos:  4217,
						EndPos:    4227,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4217,
							EndPos:    4226,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4217,
								EndPos:    4219,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4217,
									EndPos:    4219,
								},
							},
							Value: []byte("$a"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4221,
								EndPos:    4224,
							},
						},
						Value: []byte("foo"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4224,
								EndPos:    4226,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4228,
						EndPos:    4238,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4228,
							EndPos:    4237,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4232,
								EndPos:    4235,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 228,
										EndLine:   228,
										StartPos:  4232,
										EndPos:    4235,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4235,
								EndPos:    4237,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 229,
						EndLine:   229,
						StartPos:  4239,
						EndPos:    4259,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4239,
							EndPos:    4258,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4243,
								EndPos:    4256,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 229,
										EndLine:   229,
										StartPos:  4253,
										EndPos:    4256,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4256,
								EndPos:    4258,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4260,
						EndPos:    4271,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4260,
							EndPos:    4270,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 230,
								EndLine:   230,
								StartPos:  4264,
								EndPos:    4268,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 230,
										EndLine:   230,
										StartPos:  4265,
										EndPos:    4268,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 230,
								EndLine:   230,
								StartPos:  4268,
								EndPos:    4270,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4272,
						EndPos:    4297,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4272,
							EndPos:    4296,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 231,
								EndLine:   231,
								StartPos:  4276,
								EndPos:    4296,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 231,
									EndLine:   231,
									StartPos:  4282,
									EndPos:    4293,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 231,
											EndLine:   231,
											StartPos:  4283,
											EndPos:    4285,
										},
									},
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 231,
												EndLine:   231,
												StartPos:  4283,
												EndPos:    4285,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 231,
													EndLine:   231,
													StartPos:  4283,
													EndPos:    4285,
												},
											},
											Value: []byte("$a"),
										},
									},
								},
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 231,
											EndLine:   231,
											StartPos:  4287,
											EndPos:    4292,
										},
									},
									Variadic: true,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 231,
												EndLine:   231,
												StartPos:  4290,
												EndPos:    4292,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 231,
													EndLine:   231,
													StartPos:  4290,
													EndPos:    4292,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4298,
						EndPos:    4308,
					},
				},
				Expr: &ast.ExprPrint{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 232,
							EndLine:   232,
							StartPos:  4298,
							EndPos:    4306,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 232,
								EndLine:   232,
								StartPos:  4304,
								EndPos:    4306,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 232,
									EndLine:   232,
									StartPos:  4304,
									EndPos:    4306,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 233,
						EndLine:   233,
						StartPos:  4309,
						EndPos:    4317,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4309,
							EndPos:    4316,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4309,
								EndPos:    4311,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4309,
									EndPos:    4311,
								},
							},
							Value: []byte("$a"),
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4313,
								EndPos:    4316,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 234,
						EndLine:   234,
						StartPos:  4318,
						EndPos:    4327,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 234,
							EndLine:   234,
							StartPos:  4318,
							EndPos:    4326,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 234,
									EndLine:   234,
									StartPos:  4319,
									EndPos:    4323,
								},
							},
							Value: []byte("cmd "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 234,
									EndLine:   234,
									StartPos:  4323,
									EndPos:    4325,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 234,
										EndLine:   234,
										StartPos:  4323,
										EndPos:    4325,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 235,
						EndLine:   235,
						StartPos:  4328,
						EndPos:    4334,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4328,
							EndPos:    4333,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4329,
									EndPos:    4332,
								},
							},
							Value: []byte("cmd"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 236,
						EndLine:   236,
						StartPos:  4335,
						EndPos:    4338,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4335,
							EndPos:    4337,
						},
					},
					Parts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4339,
						EndPos:    4342,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4339,
							EndPos:    4341,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4343,
						EndPos:    4347,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4343,
							EndPos:    4346,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4344,
									EndPos:    4345,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 238,
										EndLine:   238,
										StartPos:  4344,
										EndPos:    4345,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 239,
						EndLine:   239,
						StartPos:  4348,
						EndPos:    4361,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4348,
							EndPos:    4360,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4349,
									EndPos:    4353,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4349,
										EndPos:    4350,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4352,
										EndPos:    4353,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4355,
									EndPos:    4358,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4355,
										EndPos:    4358,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 239,
											EndLine:   239,
											StartPos:  4356,
											EndPos:    4358,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 239,
												EndLine:   239,
												StartPos:  4356,
												EndPos:    4358,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 241,
						EndLine:   241,
						StartPos:  4363,
						EndPos:    4373,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4363,
							EndPos:    4372,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4363,
								EndPos:    4367,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 241,
										EndLine:   241,
										StartPos:  4364,
										EndPos:    4366,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 241,
											EndLine:   241,
											StartPos:  4364,
											EndPos:    4366,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 241,
												EndLine:   241,
												StartPos:  4364,
												EndPos:    4366,
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
								StartLine: 241,
								EndLine:   241,
								StartPos:  4370,
								EndPos:    4372,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 241,
									EndLine:   241,
									StartPos:  4370,
									EndPos:    4372,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 242,
						EndLine:   242,
						StartPos:  4374,
						EndPos:    4386,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4374,
							EndPos:    4385,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4374,
								EndPos:    4380,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 242,
										EndLine:   242,
										StartPos:  4375,
										EndPos:    4379,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 242,
											EndLine:   242,
											StartPos:  4375,
											EndPos:    4379,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 242,
												EndLine:   242,
												StartPos:  4375,
												EndPos:    4377,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 242,
													EndLine:   242,
													StartPos:  4375,
													EndPos:    4377,
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
								StartLine: 242,
								EndLine:   242,
								StartPos:  4383,
								EndPos:    4385,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4383,
									EndPos:    4385,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 243,
						EndLine:   243,
						StartPos:  4387,
						EndPos:    4403,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4387,
							EndPos:    4402,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4387,
								EndPos:    4397,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 243,
										EndLine:   243,
										StartPos:  4388,
										EndPos:    4396,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 243,
											EndLine:   243,
											StartPos:  4388,
											EndPos:    4396,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 243,
													EndLine:   243,
													StartPos:  4393,
													EndPos:    4395,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 243,
														EndLine:   243,
														StartPos:  4393,
														EndPos:    4395,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 243,
															EndLine:   243,
															StartPos:  4393,
															EndPos:    4395,
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
								StartLine: 243,
								EndLine:   243,
								StartPos:  4400,
								EndPos:    4402,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4400,
									EndPos:    4402,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 244,
						EndLine:   244,
						StartPos:  4404,
						EndPos:    4415,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4404,
							EndPos:    4414,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4404,
								EndPos:    4407,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 244,
										EndLine:   244,
										StartPos:  4404,
										EndPos:    4407,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4409,
								EndPos:    4412,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4412,
								EndPos:    4414,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4416,
						EndPos:    4437,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4416,
							EndPos:    4436,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4416,
								EndPos:    4429,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4426,
										EndPos:    4429,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4431,
								EndPos:    4434,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4434,
								EndPos:    4436,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4438,
						EndPos:    4450,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4438,
							EndPos:    4449,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4438,
								EndPos:    4442,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4439,
										EndPos:    4442,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4444,
								EndPos:    4447,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4447,
								EndPos:    4449,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4451,
						EndPos:    4461,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4451,
							EndPos:    4460,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4451,
								EndPos:    4454,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 247,
										EndLine:   247,
										StartPos:  4451,
										EndPos:    4454,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4456,
								EndPos:    4460,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4456,
									EndPos:    4460,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 248,
						EndLine:   248,
						StartPos:  4462,
						EndPos:    4473,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4462,
							EndPos:    4472,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4462,
								EndPos:    4466,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4462,
									EndPos:    4466,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4468,
								EndPos:    4472,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4468,
									EndPos:    4472,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 249,
						EndLine:   249,
						StartPos:  4474,
						EndPos:    4494,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4474,
							EndPos:    4493,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4474,
								EndPos:    4487,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 249,
										EndLine:   249,
										StartPos:  4484,
										EndPos:    4487,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4489,
								EndPos:    4493,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 249,
									EndLine:   249,
									StartPos:  4489,
									EndPos:    4493,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 250,
						EndLine:   250,
						StartPos:  4495,
						EndPos:    4506,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4495,
							EndPos:    4505,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4495,
								EndPos:    4499,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 250,
										EndLine:   250,
										StartPos:  4496,
										EndPos:    4499,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4501,
								EndPos:    4505,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4501,
									EndPos:    4505,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 251,
						EndLine:   251,
						StartPos:  4507,
						EndPos:    4520,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4507,
							EndPos:    4519,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4507,
								EndPos:    4509,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4507,
									EndPos:    4509,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4512,
								EndPos:    4514,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4512,
									EndPos:    4514,
								},
							},
							Value: []byte("$b"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4517,
								EndPos:    4519,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4517,
									EndPos:    4519,
								},
							},
							Value: []byte("$c"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4521,
						EndPos:    4531,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4521,
							EndPos:    4530,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4521,
								EndPos:    4523,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4521,
									EndPos:    4523,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4528,
								EndPos:    4530,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4528,
									EndPos:    4530,
								},
							},
							Value: []byte("$c"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 253,
						EndLine:   253,
						StartPos:  4532,
						EndPos:    4555,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4532,
							EndPos:    4554,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4532,
								EndPos:    4534,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4532,
									EndPos:    4534,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfTrue: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4537,
								EndPos:    4549,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4537,
									EndPos:    4539,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 253,
										EndLine:   253,
										StartPos:  4537,
										EndPos:    4539,
									},
								},
								Value: []byte("$b"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4542,
									EndPos:    4544,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 253,
										EndLine:   253,
										StartPos:  4542,
										EndPos:    4544,
									},
								},
								Value: []byte("$c"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4547,
									EndPos:    4549,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 253,
										EndLine:   253,
										StartPos:  4547,
										EndPos:    4549,
									},
								},
								Value: []byte("$d"),
							},
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4552,
								EndPos:    4554,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4552,
									EndPos:    4554,
								},
							},
							Value: []byte("$e"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 254,
						EndLine:   254,
						StartPos:  4556,
						EndPos:    4579,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4556,
							EndPos:    4578,
						},
					},
					Condition: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4556,
								EndPos:    4568,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4556,
									EndPos:    4558,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  4556,
										EndPos:    4558,
									},
								},
								Value: []byte("$a"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4561,
									EndPos:    4563,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  4561,
										EndPos:    4563,
									},
								},
								Value: []byte("$b"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4566,
									EndPos:    4568,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  4566,
										EndPos:    4568,
									},
								},
								Value: []byte("$c"),
							},
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4571,
								EndPos:    4573,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4571,
									EndPos:    4573,
								},
							},
							Value: []byte("$d"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4576,
								EndPos:    4578,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4576,
									EndPos:    4578,
								},
							},
							Value: []byte("$e"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  4580,
						EndPos:    4584,
					},
				},
				Expr: &ast.ExprUnaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  4580,
							EndPos:    4583,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4581,
								EndPos:    4583,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  4581,
									EndPos:    4583,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 256,
						EndLine:   256,
						StartPos:  4585,
						EndPos:    4589,
					},
				},
				Expr: &ast.ExprUnaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 256,
							EndLine:   256,
							StartPos:  4585,
							EndPos:    4588,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 256,
								EndLine:   256,
								StartPos:  4586,
								EndPos:    4588,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 256,
									EndLine:   256,
									StartPos:  4586,
									EndPos:    4588,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  4590,
						EndPos:    4594,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  4590,
							EndPos:    4593,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  4591,
								EndPos:    4593,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 257,
									EndLine:   257,
									StartPos:  4591,
									EndPos:    4593,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 258,
						EndLine:   258,
						StartPos:  4595,
						EndPos:    4601,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 258,
							EndLine:   258,
							StartPos:  4595,
							EndPos:    4600,
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  4602,
						EndPos:    4611,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  4602,
							EndPos:    4610,
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  4608,
								EndPos:    4610,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 259,
									EndLine:   259,
									StartPos:  4608,
									EndPos:    4610,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 260,
						EndLine:   260,
						StartPos:  4612,
						EndPos:    4627,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  4612,
							EndPos:    4626,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  4618,
								EndPos:    4620,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 260,
									EndLine:   260,
									StartPos:  4618,
									EndPos:    4620,
								},
							},
							Value: []byte("$a"),
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  4624,
								EndPos:    4626,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 260,
									EndLine:   260,
									StartPos:  4624,
									EndPos:    4626,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 261,
						EndLine:   261,
						StartPos:  4628,
						EndPos:    4642,
					},
				},
				Expr: &ast.ExprYieldFrom{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  4628,
							EndPos:    4641,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  4639,
								EndPos:    4641,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 261,
									EndLine:   261,
									StartPos:  4639,
									EndPos:    4641,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 263,
						EndLine:   263,
						StartPos:  4644,
						EndPos:    4654,
					},
				},
				Expr: &ast.ExprCastArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  4644,
							EndPos:    4653,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  4651,
								EndPos:    4653,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 263,
									EndLine:   263,
									StartPos:  4651,
									EndPos:    4653,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 264,
						EndLine:   264,
						StartPos:  4655,
						EndPos:    4667,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  4655,
							EndPos:    4666,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  4664,
								EndPos:    4666,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 264,
									EndLine:   264,
									StartPos:  4664,
									EndPos:    4666,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 265,
						EndLine:   265,
						StartPos:  4668,
						EndPos:    4677,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  4668,
							EndPos:    4676,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  4674,
								EndPos:    4676,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 265,
									EndLine:   265,
									StartPos:  4674,
									EndPos:    4676,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 266,
						EndLine:   266,
						StartPos:  4678,
						EndPos:    4689,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  4678,
							EndPos:    4688,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  4686,
								EndPos:    4688,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 266,
									EndLine:   266,
									StartPos:  4686,
									EndPos:    4688,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 267,
						EndLine:   267,
						StartPos:  4690,
						EndPos:    4700,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 267,
							EndLine:   267,
							StartPos:  4690,
							EndPos:    4699,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 267,
								EndLine:   267,
								StartPos:  4697,
								EndPos:    4699,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 267,
									EndLine:   267,
									StartPos:  4697,
									EndPos:    4699,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 268,
						EndLine:   268,
						StartPos:  4701,
						EndPos:    4713,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  4701,
							EndPos:    4712,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  4710,
								EndPos:    4712,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 268,
									EndLine:   268,
									StartPos:  4710,
									EndPos:    4712,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 269,
						EndLine:   269,
						StartPos:  4714,
						EndPos:    4722,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  4714,
							EndPos:    4721,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  4719,
								EndPos:    4721,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 269,
									EndLine:   269,
									StartPos:  4719,
									EndPos:    4721,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 270,
						EndLine:   270,
						StartPos:  4723,
						EndPos:    4734,
					},
				},
				Expr: &ast.ExprCastObject{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  4723,
							EndPos:    4733,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  4731,
								EndPos:    4733,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 270,
									EndLine:   270,
									StartPos:  4731,
									EndPos:    4733,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 271,
						EndLine:   271,
						StartPos:  4735,
						EndPos:    4746,
					},
				},
				Expr: &ast.ExprCastString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  4735,
							EndPos:    4745,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  4743,
								EndPos:    4745,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 271,
									EndLine:   271,
									StartPos:  4743,
									EndPos:    4745,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 272,
						EndLine:   272,
						StartPos:  4747,
						EndPos:    4757,
					},
				},
				Expr: &ast.ExprCastUnset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  4747,
							EndPos:    4756,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  4754,
								EndPos:    4756,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 272,
									EndLine:   272,
									StartPos:  4754,
									EndPos:    4756,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 274,
						EndLine:   274,
						StartPos:  4759,
						EndPos:    4767,
					},
				},
				Expr: &ast.ExprBinaryBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  4759,
							EndPos:    4766,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  4759,
								EndPos:    4761,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  4759,
									EndPos:    4761,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  4764,
								EndPos:    4766,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  4764,
									EndPos:    4766,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 275,
						EndLine:   275,
						StartPos:  4768,
						EndPos:    4776,
					},
				},
				Expr: &ast.ExprBinaryBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  4768,
							EndPos:    4775,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  4768,
								EndPos:    4770,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  4768,
									EndPos:    4770,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  4773,
								EndPos:    4775,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  4773,
									EndPos:    4775,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 276,
						EndLine:   276,
						StartPos:  4777,
						EndPos:    4785,
					},
				},
				Expr: &ast.ExprBinaryBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  4777,
							EndPos:    4784,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  4777,
								EndPos:    4779,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
									StartPos:  4777,
									EndPos:    4779,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  4782,
								EndPos:    4784,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
									StartPos:  4782,
									EndPos:    4784,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 277,
						EndLine:   277,
						StartPos:  4786,
						EndPos:    4795,
					},
				},
				Expr: &ast.ExprBinaryBooleanAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  4786,
							EndPos:    4794,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  4786,
								EndPos:    4788,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  4786,
									EndPos:    4788,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  4792,
								EndPos:    4794,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  4792,
									EndPos:    4794,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 278,
						EndLine:   278,
						StartPos:  4796,
						EndPos:    4805,
					},
				},
				Expr: &ast.ExprBinaryBooleanOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  4796,
							EndPos:    4804,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  4796,
								EndPos:    4798,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  4796,
									EndPos:    4798,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  4802,
								EndPos:    4804,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  4802,
									EndPos:    4804,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 279,
						EndLine:   279,
						StartPos:  4806,
						EndPos:    4815,
					},
				},
				Expr: &ast.ExprBinaryCoalesce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  4806,
							EndPos:    4814,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  4806,
								EndPos:    4808,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  4806,
									EndPos:    4808,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  4812,
								EndPos:    4814,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  4812,
									EndPos:    4814,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 280,
						EndLine:   280,
						StartPos:  4816,
						EndPos:    4824,
					},
				},
				Expr: &ast.ExprBinaryConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  4816,
							EndPos:    4823,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  4816,
								EndPos:    4818,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  4816,
									EndPos:    4818,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  4821,
								EndPos:    4823,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  4821,
									EndPos:    4823,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 281,
						EndLine:   281,
						StartPos:  4825,
						EndPos:    4833,
					},
				},
				Expr: &ast.ExprBinaryDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  4825,
							EndPos:    4832,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  4825,
								EndPos:    4827,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  4825,
									EndPos:    4827,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  4830,
								EndPos:    4832,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  4830,
									EndPos:    4832,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 282,
						EndLine:   282,
						StartPos:  4834,
						EndPos:    4843,
					},
				},
				Expr: &ast.ExprBinaryEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  4834,
							EndPos:    4842,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  4834,
								EndPos:    4836,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  4834,
									EndPos:    4836,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  4840,
								EndPos:    4842,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  4840,
									EndPos:    4842,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 283,
						EndLine:   283,
						StartPos:  4844,
						EndPos:    4853,
					},
				},
				Expr: &ast.ExprBinaryGreaterOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  4844,
							EndPos:    4852,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  4844,
								EndPos:    4846,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  4844,
									EndPos:    4846,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  4850,
								EndPos:    4852,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  4850,
									EndPos:    4852,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 284,
						EndLine:   284,
						StartPos:  4854,
						EndPos:    4862,
					},
				},
				Expr: &ast.ExprBinaryGreater{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  4854,
							EndPos:    4861,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  4854,
								EndPos:    4856,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  4854,
									EndPos:    4856,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  4859,
								EndPos:    4861,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  4859,
									EndPos:    4861,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 285,
						EndLine:   285,
						StartPos:  4863,
						EndPos:    4873,
					},
				},
				Expr: &ast.ExprBinaryIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  4863,
							EndPos:    4872,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  4863,
								EndPos:    4865,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  4863,
									EndPos:    4865,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  4870,
								EndPos:    4872,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  4870,
									EndPos:    4872,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 286,
						EndLine:   286,
						StartPos:  4874,
						EndPos:    4884,
					},
				},
				Expr: &ast.ExprBinaryLogicalAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  4874,
							EndPos:    4883,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  4874,
								EndPos:    4876,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  4874,
									EndPos:    4876,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  4881,
								EndPos:    4883,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  4881,
									EndPos:    4883,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 287,
						EndLine:   287,
						StartPos:  4885,
						EndPos:    4894,
					},
				},
				Expr: &ast.ExprBinaryLogicalOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  4885,
							EndPos:    4893,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  4885,
								EndPos:    4887,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  4885,
									EndPos:    4887,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  4891,
								EndPos:    4893,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  4891,
									EndPos:    4893,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 288,
						EndLine:   288,
						StartPos:  4895,
						EndPos:    4905,
					},
				},
				Expr: &ast.ExprBinaryLogicalXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  4895,
							EndPos:    4904,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  4895,
								EndPos:    4897,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
									StartPos:  4895,
									EndPos:    4897,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  4902,
								EndPos:    4904,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
									StartPos:  4902,
									EndPos:    4904,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 289,
						EndLine:   289,
						StartPos:  4906,
						EndPos:    4914,
					},
				},
				Expr: &ast.ExprBinaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  4906,
							EndPos:    4913,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  4906,
								EndPos:    4908,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  4906,
									EndPos:    4908,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  4911,
								EndPos:    4913,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  4911,
									EndPos:    4913,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 290,
						EndLine:   290,
						StartPos:  4915,
						EndPos:    4923,
					},
				},
				Expr: &ast.ExprBinaryMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  4915,
							EndPos:    4922,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  4915,
								EndPos:    4917,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  4915,
									EndPos:    4917,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  4920,
								EndPos:    4922,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  4920,
									EndPos:    4922,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 291,
						EndLine:   291,
						StartPos:  4924,
						EndPos:    4932,
					},
				},
				Expr: &ast.ExprBinaryMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  4924,
							EndPos:    4931,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  4924,
								EndPos:    4926,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  4924,
									EndPos:    4926,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  4929,
								EndPos:    4931,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  4929,
									EndPos:    4931,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 292,
						EndLine:   292,
						StartPos:  4933,
						EndPos:    4942,
					},
				},
				Expr: &ast.ExprBinaryNotEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  4933,
							EndPos:    4941,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  4933,
								EndPos:    4935,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  4933,
									EndPos:    4935,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  4939,
								EndPos:    4941,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  4939,
									EndPos:    4941,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 293,
						EndLine:   293,
						StartPos:  4943,
						EndPos:    4953,
					},
				},
				Expr: &ast.ExprBinaryNotIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  4943,
							EndPos:    4952,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  4943,
								EndPos:    4945,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 293,
									EndLine:   293,
									StartPos:  4943,
									EndPos:    4945,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  4950,
								EndPos:    4952,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 293,
									EndLine:   293,
									StartPos:  4950,
									EndPos:    4952,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 294,
						EndLine:   294,
						StartPos:  4954,
						EndPos:    4962,
					},
				},
				Expr: &ast.ExprBinaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  4954,
							EndPos:    4961,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  4954,
								EndPos:    4956,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  4954,
									EndPos:    4956,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  4959,
								EndPos:    4961,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  4959,
									EndPos:    4961,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 295,
						EndLine:   295,
						StartPos:  4963,
						EndPos:    4972,
					},
				},
				Expr: &ast.ExprBinaryPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  4963,
							EndPos:    4971,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  4963,
								EndPos:    4965,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  4963,
									EndPos:    4965,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  4969,
								EndPos:    4971,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  4969,
									EndPos:    4971,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  4973,
						EndPos:    4982,
					},
				},
				Expr: &ast.ExprBinaryShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  4973,
							EndPos:    4981,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  4973,
								EndPos:    4975,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  4973,
									EndPos:    4975,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  4979,
								EndPos:    4981,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  4979,
									EndPos:    4981,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  4983,
						EndPos:    4992,
					},
				},
				Expr: &ast.ExprBinaryShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  4983,
							EndPos:    4991,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  4983,
								EndPos:    4985,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  4983,
									EndPos:    4985,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  4989,
								EndPos:    4991,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  4989,
									EndPos:    4991,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 298,
						EndLine:   298,
						StartPos:  4993,
						EndPos:    5002,
					},
				},
				Expr: &ast.ExprBinarySmallerOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  4993,
							EndPos:    5001,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  4993,
								EndPos:    4995,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  4993,
									EndPos:    4995,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  4999,
								EndPos:    5001,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  4999,
									EndPos:    5001,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 299,
						EndLine:   299,
						StartPos:  5003,
						EndPos:    5011,
					},
				},
				Expr: &ast.ExprBinarySmaller{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5003,
							EndPos:    5010,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5003,
								EndPos:    5005,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5003,
									EndPos:    5005,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5008,
								EndPos:    5010,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5008,
									EndPos:    5010,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 300,
						EndLine:   300,
						StartPos:  5012,
						EndPos:    5022,
					},
				},
				Expr: &ast.ExprBinarySpaceship{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5012,
							EndPos:    5021,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5012,
								EndPos:    5014,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5012,
									EndPos:    5014,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5019,
								EndPos:    5021,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5019,
									EndPos:    5021,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 302,
						EndLine:   302,
						StartPos:  5024,
						EndPos:    5033,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5024,
							EndPos:    5032,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5024,
								EndPos:    5026,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5024,
									EndPos:    5026,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5030,
								EndPos:    5032,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5030,
									EndPos:    5032,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 303,
						EndLine:   303,
						StartPos:  5034,
						EndPos:    5042,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5034,
							EndPos:    5041,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5034,
								EndPos:    5036,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5034,
									EndPos:    5036,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5039,
								EndPos:    5041,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5039,
									EndPos:    5041,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 304,
						EndLine:   304,
						StartPos:  5043,
						EndPos:    5052,
					},
				},
				Expr: &ast.ExprAssignBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5043,
							EndPos:    5051,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5043,
								EndPos:    5045,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5043,
									EndPos:    5045,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5049,
								EndPos:    5051,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5049,
									EndPos:    5051,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 305,
						EndLine:   305,
						StartPos:  5053,
						EndPos:    5062,
					},
				},
				Expr: &ast.ExprAssignBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5053,
							EndPos:    5061,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5053,
								EndPos:    5055,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5053,
									EndPos:    5055,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5059,
								EndPos:    5061,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5059,
									EndPos:    5061,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 306,
						EndLine:   306,
						StartPos:  5063,
						EndPos:    5072,
					},
				},
				Expr: &ast.ExprAssignBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5063,
							EndPos:    5071,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5063,
								EndPos:    5065,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5063,
									EndPos:    5065,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5069,
								EndPos:    5071,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5069,
									EndPos:    5071,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 307,
						EndLine:   307,
						StartPos:  5073,
						EndPos:    5082,
					},
				},
				Expr: &ast.ExprAssignConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5073,
							EndPos:    5081,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5073,
								EndPos:    5075,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5073,
									EndPos:    5075,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5079,
								EndPos:    5081,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5079,
									EndPos:    5081,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 308,
						EndLine:   308,
						StartPos:  5083,
						EndPos:    5092,
					},
				},
				Expr: &ast.ExprAssignDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5083,
							EndPos:    5091,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5083,
								EndPos:    5085,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5083,
									EndPos:    5085,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5089,
								EndPos:    5091,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5089,
									EndPos:    5091,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 309,
						EndLine:   309,
						StartPos:  5093,
						EndPos:    5102,
					},
				},
				Expr: &ast.ExprAssignMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5093,
							EndPos:    5101,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5093,
								EndPos:    5095,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5093,
									EndPos:    5095,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5099,
								EndPos:    5101,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5099,
									EndPos:    5101,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 310,
						EndLine:   310,
						StartPos:  5103,
						EndPos:    5112,
					},
				},
				Expr: &ast.ExprAssignMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5103,
							EndPos:    5111,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5103,
								EndPos:    5105,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 310,
									EndLine:   310,
									StartPos:  5103,
									EndPos:    5105,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5109,
								EndPos:    5111,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 310,
									EndLine:   310,
									StartPos:  5109,
									EndPos:    5111,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 311,
						EndLine:   311,
						StartPos:  5113,
						EndPos:    5122,
					},
				},
				Expr: &ast.ExprAssignMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5113,
							EndPos:    5121,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5113,
								EndPos:    5115,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 311,
									EndLine:   311,
									StartPos:  5113,
									EndPos:    5115,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5119,
								EndPos:    5121,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 311,
									EndLine:   311,
									StartPos:  5119,
									EndPos:    5121,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 312,
						EndLine:   312,
						StartPos:  5123,
						EndPos:    5132,
					},
				},
				Expr: &ast.ExprAssignPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5123,
							EndPos:    5131,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5123,
								EndPos:    5125,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 312,
									EndLine:   312,
									StartPos:  5123,
									EndPos:    5125,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5129,
								EndPos:    5131,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 312,
									EndLine:   312,
									StartPos:  5129,
									EndPos:    5131,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 313,
						EndLine:   313,
						StartPos:  5133,
						EndPos:    5143,
					},
				},
				Expr: &ast.ExprAssignPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5133,
							EndPos:    5142,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5133,
								EndPos:    5135,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5133,
									EndPos:    5135,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5140,
								EndPos:    5142,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5140,
									EndPos:    5142,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5144,
						EndPos:    5154,
					},
				},
				Expr: &ast.ExprAssignShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5144,
							EndPos:    5153,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5144,
								EndPos:    5146,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5144,
									EndPos:    5146,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5151,
								EndPos:    5153,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5151,
									EndPos:    5153,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5155,
						EndPos:    5165,
					},
				},
				Expr: &ast.ExprAssignShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5155,
							EndPos:    5164,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5155,
								EndPos:    5157,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5155,
									EndPos:    5157,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5162,
								EndPos:    5164,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5162,
									EndPos:    5164,
								},
							},
							Value: []byte("$b"),
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5167,
						EndPos:    5206,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5173,
							EndPos:    5176,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5178,
								EndPos:    5204,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5194,
									EndPos:    5199,
								},
							},
							Value: []byte("class"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 317,
										EndLine:   317,
										StartPos:  5178,
										EndPos:    5184,
									},
								},
								Value: []byte("public"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5202,
									EndPos:    5204,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 318,
						EndLine:   318,
						StartPos:  5207,
						EndPos:    5218,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5207,
							EndPos:    5217,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5207,
								EndPos:    5215,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 318,
										EndLine:   318,
										StartPos:  5208,
										EndPos:    5211,
									},
								},
								Value: []byte("foo"),
							},
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 318,
										EndLine:   318,
										StartPos:  5212,
										EndPos:    5215,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5215,
								EndPos:    5217,
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   326,
						StartPos:  5220,
						EndPos:    5328,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5229,
							EndPos:    5232,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5233,
								EndPos:    5236,
							},
						},
						Var: &ast.Reference{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 320,
									EndLine:   320,
									StartPos:  5233,
									EndPos:    5236,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 320,
										EndLine:   320,
										StartPos:  5234,
										EndPos:    5236,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 320,
											EndLine:   320,
											StartPos:  5234,
											EndPos:    5236,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5238,
								EndPos:    5243,
							},
						},
						Var: &ast.Variadic{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 320,
									EndLine:   320,
									StartPos:  5238,
									EndPos:    5243,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 320,
										EndLine:   320,
										StartPos:  5241,
										EndPos:    5243,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 320,
											EndLine:   320,
											StartPos:  5241,
											EndPos:    5243,
										},
									},
									Value: []byte("$b"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtFunction{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 322,
								EndLine:   322,
								StartPos:  5252,
								EndPos:    5269,
							},
						},
						FunctionName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 322,
									EndLine:   322,
									StartPos:  5261,
									EndPos:    5264,
								},
							},
							Value: []byte("bar"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5274,
								EndPos:    5286,
							},
						},
						ClassName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5280,
									EndPos:    5283,
								},
							},
							Value: []byte("Baz"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtTrait{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5291,
								EndPos:    5303,
							},
						},
						TraitName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5297,
									EndPos:    5301,
								},
							},
							Value: []byte("Quux"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtInterface{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5308,
								EndPos:    5326,
							},
						},
						InterfaceName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5318,
									EndPos:    5323,
								},
							},
							Value: []byte("Quuux"),
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 328,
						EndLine:   328,
						StartPos:  5330,
						EndPos:    5373,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 328,
							EndLine:   328,
							StartPos:  5339,
							EndPos:    5342,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5343,
								EndPos:    5350,
							},
						},
						Var: &ast.Reference{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5343,
									EndPos:    5346,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5344,
										EndPos:    5346,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 328,
											EndLine:   328,
											StartPos:  5344,
											EndPos:    5346,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
						DefaultValue: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5349,
									EndPos:    5350,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5352,
								EndPos:    5361,
							},
						},
						Var: &ast.Variadic{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5352,
									EndPos:    5357,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5355,
										EndPos:    5357,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 328,
											EndLine:   328,
											StartPos:  5355,
											EndPos:    5357,
										},
									},
									Value: []byte("$b"),
								},
							},
						},
						DefaultValue: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5360,
									EndPos:    5361,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5363,
								EndPos:    5369,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5363,
									EndPos:    5365,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5363,
										EndPos:    5365,
									},
								},
								Value: []byte("$c"),
							},
						},
						DefaultValue: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5368,
									EndPos:    5369,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 329,
						EndLine:   329,
						StartPos:  5374,
						EndPos:    5412,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5383,
							EndPos:    5386,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5387,
								EndPos:    5395,
							},
						},
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5387,
									EndPos:    5392,
								},
							},
							Value: []byte("array"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5393,
									EndPos:    5395,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5393,
										EndPos:    5395,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5397,
								EndPos:    5408,
							},
						},
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5397,
									EndPos:    5405,
								},
							},
							Value: []byte("callable"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5406,
									EndPos:    5408,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5406,
										EndPos:    5408,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 330,
						EndLine:   330,
						StartPos:  5413,
						EndPos:    5515,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  5434,
							EndPos:    5437,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5413,
								EndPos:    5421,
							},
						},
						Value: []byte("abstract"),
					},
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5422,
								EndPos:    5427,
							},
						},
						Value: []byte("final"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5440,
								EndPos:    5481,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5475,
									EndPos:    5478,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5440,
										EndPos:    5448,
									},
								},
								Value: []byte("abstract"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5449,
										EndPos:    5458,
									},
								},
								Value: []byte("protected"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5459,
										EndPos:    5465,
									},
								},
								Value: []byte("static"),
							},
						},
						Stmt: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5480,
									EndPos:    5481,
								},
							},
						},
					},
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5482,
								EndPos:    5513,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5505,
									EndPos:    5508,
								},
							},
							Value: []byte("baz"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5482,
										EndPos:    5487,
									},
								},
								Value: []byte("final"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5488,
										EndPos:    5495,
									},
								},
								Value: []byte("private"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5511,
									EndPos:    5513,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 332,
						EndLine:   332,
						StartPos:  5518,
						EndPos:    5532,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 332,
							EndLine:   332,
							StartPos:  5518,
							EndPos:    5531,
						},
					},
					Var: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  5518,
								EndPos:    5525,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  5522,
									EndPos:    5525,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 332,
											EndLine:   332,
											StartPos:  5522,
											EndPos:    5525,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  5528,
								EndPos:    5531,
							},
						},
						Value: []byte("bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 333,
						EndLine:   333,
						StartPos:  5534,
						EndPos:    5545,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  5534,
							EndPos:    5544,
						},
					},
					Function: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  5534,
								EndPos:    5541,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  5538,
									EndPos:    5541,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 333,
											EndLine:   333,
											StartPos:  5538,
											EndPos:    5541,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  5542,
								EndPos:    5544,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 334,
						EndLine:   334,
						StartPos:  5546,
						EndPos:    5558,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  5546,
							EndPos:    5557,
						},
					},
					Function: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  5546,
								EndPos:    5555,
							},
						},
						Var: &ast.ExprShortArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  5546,
									EndPos:    5552,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 334,
											EndLine:   334,
											StartPos:  5547,
											EndPos:    5551,
										},
									},
									Val: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 334,
												EndLine:   334,
												StartPos:  5547,
												EndPos:    5551,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 334,
													EndLine:   334,
													StartPos:  5547,
													EndPos:    5551,
												},
											},
											Value: []byte("$foo"),
										},
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  5553,
									EndPos:    5554,
								},
							},
							Value: []byte("0"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  5555,
								EndPos:    5557,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 335,
						EndLine:   335,
						StartPos:  5559,
						EndPos:    5568,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  5559,
							EndPos:    5567,
						},
					},
					Function: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  5559,
								EndPos:    5565,
							},
						},
						Var: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  5559,
									EndPos:    5562,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  5559,
										EndPos:    5562,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 335,
												EndLine:   335,
												StartPos:  5559,
												EndPos:    5562,
											},
										},
										Value: []byte("foo"),
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  5563,
									EndPos:    5564,
								},
							},
							Value: []byte("1"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  5565,
								EndPos:    5567,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 336,
						EndLine:   336,
						StartPos:  5569,
						EndPos:    5577,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  5569,
							EndPos:    5576,
						},
					},
					Function: &ast.ScalarString{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  5569,
								EndPos:    5574,
							},
						},
						Value: []byte("\"foo\""),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  5574,
								EndPos:    5576,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 337,
						EndLine:   337,
						StartPos:  5578,
						EndPos:    5590,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  5578,
							EndPos:    5589,
						},
					},
					Function: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  5578,
								EndPos:    5587,
							},
						},
						Var: &ast.ExprShortArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  5578,
									EndPos:    5581,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 337,
											EndLine:   337,
											StartPos:  5579,
											EndPos:    5580,
										},
									},
									Val: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 337,
												EndLine:   337,
												StartPos:  5579,
												EndPos:    5580,
											},
										},
										Value: []byte("1"),
									},
								},
							},
						},
						Dim: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  5582,
									EndPos:    5586,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  5582,
										EndPos:    5586,
									},
								},
								Value: []byte("$foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  5587,
								EndPos:    5589,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 338,
						EndLine:   338,
						StartPos:  5591,
						EndPos:    5600,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  5591,
							EndPos:    5599,
						},
					},
					VarName: &ast.ExprFunctionCall{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  5593,
								EndPos:    5598,
							},
						},
						Function: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  5593,
									EndPos:    5596,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 338,
											EndLine:   338,
											StartPos:  5593,
											EndPos:    5596,
										},
									},
									Value: []byte("foo"),
								},
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  5596,
									EndPos:    5598,
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 340,
						EndLine:   340,
						StartPos:  5602,
						EndPos:    5614,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  5602,
							EndPos:    5613,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  5602,
								EndPos:    5605,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  5602,
										EndPos:    5605,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  5607,
								EndPos:    5611,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  5607,
									EndPos:    5611,
								},
							},
							Value: []byte("$bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  5611,
								EndPos:    5613,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 341,
						EndLine:   341,
						StartPos:  5615,
						EndPos:    5632,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  5615,
							EndPos:    5631,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  5615,
								EndPos:    5618,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  5615,
										EndPos:    5618,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  5621,
								EndPos:    5628,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  5621,
									EndPos:    5625,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  5621,
										EndPos:    5625,
									},
								},
								Value: []byte("$bar"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  5626,
									EndPos:    5627,
								},
							},
							Value: []byte("0"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  5629,
								EndPos:    5631,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 343,
						EndLine:   343,
						StartPos:  5634,
						EndPos:    5645,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 343,
							EndLine:   343,
							StartPos:  5634,
							EndPos:    5644,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  5634,
								EndPos:    5638,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  5634,
									EndPos:    5638,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  5640,
								EndPos:    5644,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  5640,
									EndPos:    5644,
								},
							},
							Value: []byte("$bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 344,
						EndLine:   344,
						StartPos:  5646,
						EndPos:    5662,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  5646,
							EndPos:    5660,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  5646,
								EndPos:    5650,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  5646,
									EndPos:    5650,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Property: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  5653,
								EndPos:    5660,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  5653,
									EndPos:    5657,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  5653,
										EndPos:    5657,
									},
								},
								Value: []byte("$bar"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  5658,
									EndPos:    5659,
								},
							},
							Value: []byte("0"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 346,
						EndLine:   346,
						StartPos:  5664,
						EndPos:    5686,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 346,
							EndLine:   346,
							StartPos:  5664,
							EndPos:    5685,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  5665,
									EndPos:    5671,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5665,
										EndPos:    5666,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5668,
										EndPos:    5671,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 346,
											EndLine:   346,
											StartPos:  5669,
											EndPos:    5671,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 346,
												EndLine:   346,
												StartPos:  5669,
												EndPos:    5671,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  5673,
									EndPos:    5684,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5673,
										EndPos:    5674,
									},
								},
								Value: []byte("2"),
							},
							Val: &ast.ExprList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5676,
										EndPos:    5684,
									},
								},
								Items: []ast.Vertex{
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 346,
												EndLine:   346,
												StartPos:  5681,
												EndPos:    5683,
											},
										},
										Val: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 346,
													EndLine:   346,
													StartPos:  5681,
													EndPos:    5683,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 346,
														EndLine:   346,
														StartPos:  5681,
														EndPos:    5683,
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
			&ast.StmtHaltCompiler{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 348,
						EndLine:   348,
						StartPos:  5688,
						EndPos:    5706,
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer(src, "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp5Strings(t *testing.T) {
	src := `<?
		"test";
		"\$test";
		"
			test
		";
		'$test';
		'
			$test
		';
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   10,
				StartPos:  5,
				EndPos:    70,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  5,
						EndPos:    12,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  5,
							EndPos:    11,
						},
					},
					Value: []byte("\"test\""),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  15,
						EndPos:    24,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  15,
							EndPos:    23,
						},
					},
					Value: []byte("\"\\$test\""),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  27,
						EndPos:    41,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   6,
							StartPos:  27,
							EndPos:    40,
						},
					},
					Value: []byte("\"\n\t\t\ttest\n\t\t\""),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  44,
						EndPos:    52,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  44,
							EndPos:    51,
						},
					},
					Value: []byte("'$test'"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 8,
						EndLine:   10,
						StartPos:  55,
						EndPos:    70,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 8,
							EndLine:   10,
							StartPos:  55,
							EndPos:    69,
						},
					},
					Value: []byte("'\n\t\t\t$test\n\t\t'"),
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, nil)
	php7parser := php7.NewParser(lexer, nil)
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp5Heredoc(t *testing.T) {
	src := `<?
		<<<CAD
CAD;
		<<<CAD
	hello
CAD;
		<<<"CAD"
	hello
CAD;
		<<<"CAD"
	hello $world
CAD;
		<<<'CAD'
	hello $world
CAD;
	`

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   15,
				StartPos:  5,
				EndPos:    120,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   3,
						StartPos:  5,
						EndPos:    16,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   3,
							StartPos:  5,
							EndPos:    15,
						},
					},
					Label: []byte("<<<CAD\n"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  19,
						EndPos:    37,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   6,
							StartPos:  19,
							EndPos:    36,
						},
					},
					Label: []byte("<<<CAD\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  26,
									EndPos:    33,
								},
							},
							Value: []byte("\thello\n"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   9,
						StartPos:  40,
						EndPos:    60,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   9,
							StartPos:  40,
							EndPos:    59,
						},
					},
					Label: []byte("<<<\"CAD\"\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 8,
									EndLine:   8,
									StartPos:  49,
									EndPos:    56,
								},
							},
							Value: []byte("\thello\n"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   12,
						StartPos:  63,
						EndPos:    90,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   12,
							StartPos:  63,
							EndPos:    89,
						},
					},
					Label: []byte("<<<\"CAD\"\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  72,
									EndPos:    79,
								},
							},
							Value: []byte("\thello "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  79,
									EndPos:    85,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  79,
										EndPos:    85,
									},
								},
								Value: []byte("$world"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  85,
									EndPos:    86,
								},
							},
							Value: []byte("\n"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   15,
						StartPos:  93,
						EndPos:    120,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 13,
							EndLine:   15,
							StartPos:  93,
							EndPos:    119,
						},
					},
					Label: []byte("<<<'CAD'\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  102,
									EndPos:    116,
								},
							},
							Value: []byte("\thello $world\n"),
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

func TestPhp7ControlCharsErrors(t *testing.T) {
	src := "<?php \004 echo $b; \"$a[\005test]\";"

	expected := []*errors.Error{
		{
			Msg: "WARNING: Unexpected character in input: '\004' (ASCII=4)",
			Pos: &position.Position{StartLine: 1, EndLine: 1, StartPos: 6, EndPos: 7},
		},
		{
			Msg: "WARNING: Unexpected character in input: '\005' (ASCII=5)",
			Pos: &position.Position{StartLine: 1, EndLine: 1, StartPos: 21, EndPos: 22},
		},
	}

	parserErrors := []*errors.Error{}
	errorHandlerFunc := func(e *errors.Error) {
		parserErrors = append(parserErrors, e)
	}

	lexer := scanner.NewLexer([]byte(src), "7.4", false, errorHandlerFunc)
	php7parser := php7.NewParser(lexer, errorHandlerFunc)
	php7parser.Parse()
	assert.DeepEqual(t, expected, parserErrors)
}
