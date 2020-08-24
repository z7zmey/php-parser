package php5_test

import (
	"io/ioutil"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/internal/php5"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/ast/traverser"
	"github.com/z7zmey/php-parser/pkg/ast/visitor"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
)

func TestPhp5(t *testing.T) {
	src, err := ioutil.ReadFile("test.php")
	assert.NilError(t, err)

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   379,
				StartPos:  3,
				EndPos:    6285,
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
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  122,
						EndPos:    166,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  131,
							EndPos:    134,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  135,
								EndPos:    148,
							},
						},
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  135,
									EndPos:    138,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  135,
											EndPos:    138,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  139,
									EndPos:    143,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  139,
										EndPos:    143,
									},
								},
								Value: []byte("$bar"),
							},
						},
						DefaultValue: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  144,
									EndPos:    148,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  144,
										EndPos:    148,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  144,
												EndPos:    148,
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
								StartLine: 9,
								EndLine:   9,
								StartPos:  150,
								EndPos:    162,
							},
						},
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  150,
									EndPos:    153,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  150,
											EndPos:    153,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
						Var: &ast.Reference{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  154,
									EndPos:    162,
								},
							},
							Var: &ast.Variadic{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  155,
										EndPos:    162,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  158,
											EndPos:    162,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  158,
												EndPos:    162,
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
						StartLine: 10,
						EndLine:   10,
						StartPos:  167,
						EndPos:    230,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  173,
							EndPos:    176,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  178,
								EndPos:    229,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  194,
									EndPos:    197,
								},
							},
							Value: []byte("foo"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  178,
										EndPos:    184,
									},
								},
								Value: []byte("public"),
							},
						},
						Params: []ast.Vertex{
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  198,
										EndPos:    211,
									},
								},
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  198,
											EndPos:    201,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 10,
													EndLine:   10,
													StartPos:  198,
													EndPos:    201,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  202,
											EndPos:    206,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  202,
												EndPos:    206,
											},
										},
										Value: []byte("$bar"),
									},
								},
								DefaultValue: &ast.ExprConstFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  207,
											EndPos:    211,
										},
									},
									Const: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  207,
												EndPos:    211,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 10,
														EndLine:   10,
														StartPos:  207,
														EndPos:    211,
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
										StartLine: 10,
										EndLine:   10,
										StartPos:  213,
										EndPos:    225,
									},
								},
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  213,
											EndPos:    216,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 10,
													EndLine:   10,
													StartPos:  213,
													EndPos:    216,
												},
											},
											Value: []byte("baz"),
										},
									},
								},
								Var: &ast.Reference{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  217,
											EndPos:    225,
										},
									},
									Var: &ast.Variadic{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  218,
												EndPos:    225,
											},
										},
										Var: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 10,
													EndLine:   10,
													StartPos:  221,
													EndPos:    225,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 10,
														EndLine:   10,
														StartPos:  221,
														EndPos:    225,
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
									StartLine: 10,
									EndLine:   10,
									StartPos:  227,
									EndPos:    229,
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
						StartLine: 11,
						EndLine:   11,
						StartPos:  231,
						EndPos:    272,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  231,
							EndPos:    271,
						},
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  240,
									EndPos:    253,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  240,
										EndPos:    243,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 11,
												EndLine:   11,
												StartPos:  240,
												EndPos:    243,
											},
										},
										Value: []byte("bar"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  244,
										EndPos:    248,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  244,
											EndPos:    248,
										},
									},
									Value: []byte("$bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  249,
										EndPos:    253,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  249,
											EndPos:    253,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 11,
													EndLine:   11,
													StartPos:  249,
													EndPos:    253,
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
									StartLine: 11,
									EndLine:   11,
									StartPos:  255,
									EndPos:    267,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  255,
										EndPos:    258,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 11,
												EndLine:   11,
												StartPos:  255,
												EndPos:    258,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.Reference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  259,
										EndPos:    267,
									},
								},
								Var: &ast.Variadic{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  260,
											EndPos:    267,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 11,
												EndLine:   11,
												StartPos:  263,
												EndPos:    267,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 11,
													EndLine:   11,
													StartPos:  263,
													EndPos:    267,
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
						StartLine: 12,
						EndLine:   12,
						StartPos:  273,
						EndPos:    321,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 12,
							EndLine:   12,
							StartPos:  273,
							EndPos:    320,
						},
					},
					Static: true,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  289,
									EndPos:    302,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  289,
										EndPos:    292,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 12,
												EndLine:   12,
												StartPos:  289,
												EndPos:    292,
											},
										},
										Value: []byte("bar"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  293,
										EndPos:    297,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  293,
											EndPos:    297,
										},
									},
									Value: []byte("$bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  298,
										EndPos:    302,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  298,
											EndPos:    302,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 12,
													EndLine:   12,
													StartPos:  298,
													EndPos:    302,
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
									StartLine: 12,
									EndLine:   12,
									StartPos:  304,
									EndPos:    316,
								},
							},
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  304,
										EndPos:    307,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 12,
												EndLine:   12,
												StartPos:  304,
												EndPos:    307,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.Reference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  308,
										EndPos:    316,
									},
								},
								Var: &ast.Variadic{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  309,
											EndPos:    316,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 12,
												EndLine:   12,
												StartPos:  312,
												EndPos:    316,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 12,
													EndLine:   12,
													StartPos:  312,
													EndPos:    316,
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
						StartLine: 14,
						EndLine:   14,
						StartPos:  323,
						EndPos:    343,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 14,
							EndLine:   14,
							StartPos:  323,
							EndPos:    342,
						},
					},
					Value: []byte("1234567890123456789"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 15,
						EndLine:   15,
						StartPos:  344,
						EndPos:    365,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 15,
							EndLine:   15,
							StartPos:  344,
							EndPos:    364,
						},
					},
					Value: []byte("12345678901234567890"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 16,
						EndLine:   16,
						StartPos:  366,
						EndPos:    369,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 16,
							EndLine:   16,
							StartPos:  366,
							EndPos:    368,
						},
					},
					Value: []byte("0."),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 17,
						EndLine:   17,
						StartPos:  370,
						EndPos:    437,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 17,
							EndLine:   17,
							StartPos:  370,
							EndPos:    436,
						},
					},
					Value: []byte("0b0111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  438,
						EndPos:    505,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  438,
							EndPos:    504,
						},
					},
					Value: []byte("0b1111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 19,
						EndLine:   19,
						StartPos:  506,
						EndPos:    527,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 19,
							EndLine:   19,
							StartPos:  506,
							EndPos:    526,
						},
					},
					Value: []byte("0x007111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 20,
						EndLine:   20,
						StartPos:  528,
						EndPos:    547,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 20,
							EndLine:   20,
							StartPos:  528,
							EndPos:    546,
						},
					},
					Value: []byte("0x8111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 21,
						EndLine:   21,
						StartPos:  548,
						EndPos:    558,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 21,
							EndLine:   21,
							StartPos:  548,
							EndPos:    557,
						},
					},
					Value: []byte("__CLASS__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 22,
						EndLine:   22,
						StartPos:  559,
						EndPos:    567,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 22,
							EndLine:   22,
							StartPos:  559,
							EndPos:    566,
						},
					},
					Value: []byte("__DIR__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 23,
						EndLine:   23,
						StartPos:  568,
						EndPos:    577,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 23,
							EndLine:   23,
							StartPos:  568,
							EndPos:    576,
						},
					},
					Value: []byte("__FILE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  578,
						EndPos:    591,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 24,
							EndLine:   24,
							StartPos:  578,
							EndPos:    590,
						},
					},
					Value: []byte("__FUNCTION__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  592,
						EndPos:    601,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 25,
							EndLine:   25,
							StartPos:  592,
							EndPos:    600,
						},
					},
					Value: []byte("__LINE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  602,
						EndPos:    616,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 26,
							EndLine:   26,
							StartPos:  602,
							EndPos:    615,
						},
					},
					Value: []byte("__NAMESPACE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  617,
						EndPos:    628,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 27,
							EndLine:   27,
							StartPos:  617,
							EndPos:    627,
						},
					},
					Value: []byte("__METHOD__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  629,
						EndPos:    639,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 28,
							EndLine:   28,
							StartPos:  629,
							EndPos:    638,
						},
					},
					Value: []byte("__TRAIT__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  641,
						EndPos:    653,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 30,
							EndLine:   30,
							StartPos:  641,
							EndPos:    652,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 30,
									EndLine:   30,
									StartPos:  642,
									EndPos:    647,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 30,
									EndLine:   30,
									StartPos:  647,
									EndPos:    651,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 30,
										EndLine:   30,
										StartPos:  647,
										EndPos:    651,
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
						StartLine: 31,
						EndLine:   31,
						StartPos:  654,
						EndPos:    669,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 31,
							EndLine:   31,
							StartPos:  654,
							EndPos:    668,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 31,
									EndLine:   31,
									StartPos:  655,
									EndPos:    660,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 31,
									EndLine:   31,
									StartPos:  660,
									EndPos:    667,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 31,
										EndLine:   31,
										StartPos:  660,
										EndPos:    664,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 31,
											EndLine:   31,
											StartPos:  660,
											EndPos:    664,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 31,
										EndLine:   31,
										StartPos:  665,
										EndPos:    666,
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
						StartLine: 32,
						EndLine:   32,
						StartPos:  670,
						EndPos:    724,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 32,
							EndLine:   32,
							StartPos:  670,
							EndPos:    723,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 32,
									EndLine:   32,
									StartPos:  671,
									EndPos:    676,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 32,
									EndLine:   32,
									StartPos:  676,
									EndPos:    722,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 32,
										EndLine:   32,
										StartPos:  676,
										EndPos:    680,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 32,
											EndLine:   32,
											StartPos:  676,
											EndPos:    680,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 32,
										EndLine:   32,
										StartPos:  681,
										EndPos:    721,
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
						StartLine: 33,
						EndLine:   33,
						StartPos:  725,
						EndPos:    742,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 33,
							EndLine:   33,
							StartPos:  725,
							EndPos:    741,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 33,
									EndLine:   33,
									StartPos:  726,
									EndPos:    731,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 33,
									EndLine:   33,
									StartPos:  731,
									EndPos:    740,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 33,
										EndLine:   33,
										StartPos:  731,
										EndPos:    735,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 33,
											EndLine:   33,
											StartPos:  731,
											EndPos:    735,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 33,
										EndLine:   33,
										StartPos:  736,
										EndPos:    739,
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
						StartLine: 34,
						EndLine:   34,
						StartPos:  743,
						EndPos:    761,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 34,
							EndLine:   34,
							StartPos:  743,
							EndPos:    760,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 34,
									EndLine:   34,
									StartPos:  744,
									EndPos:    749,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 34,
									EndLine:   34,
									StartPos:  749,
									EndPos:    759,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 34,
										EndLine:   34,
										StartPos:  749,
										EndPos:    753,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 34,
											EndLine:   34,
											StartPos:  749,
											EndPos:    753,
										},
									},
									Value: []byte("$var"),
								},
							},
							Dim: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 34,
										EndLine:   34,
										StartPos:  754,
										EndPos:    758,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 34,
											EndLine:   34,
											StartPos:  754,
											EndPos:    758,
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
						StartLine: 35,
						EndLine:   35,
						StartPos:  762,
						EndPos:    774,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 35,
							EndLine:   35,
							StartPos:  762,
							EndPos:    773,
						},
					},
					Parts: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  763,
									EndPos:    767,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 35,
										EndLine:   35,
										StartPos:  763,
										EndPos:    767,
									},
								},
								Value: []byte("$foo"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  767,
									EndPos:    768,
								},
							},
							Value: []byte(" "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  768,
									EndPos:    772,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 35,
										EndLine:   35,
										StartPos:  768,
										EndPos:    772,
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
						StartLine: 36,
						EndLine:   36,
						StartPos:  775,
						EndPos:    794,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 36,
							EndLine:   36,
							StartPos:  775,
							EndPos:    793,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  776,
									EndPos:    781,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  781,
									EndPos:    790,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 36,
										EndLine:   36,
										StartPos:  781,
										EndPos:    785,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 36,
											EndLine:   36,
											StartPos:  781,
											EndPos:    785,
										},
									},
									Value: []byte("$foo"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 36,
										EndLine:   36,
										StartPos:  787,
										EndPos:    790,
									},
								},
								Value: []byte("bar"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  790,
									EndPos:    792,
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
						StartLine: 37,
						EndLine:   37,
						StartPos:  795,
						EndPos:    809,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 37,
							EndLine:   37,
							StartPos:  795,
							EndPos:    808,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 37,
									EndLine:   37,
									StartPos:  796,
									EndPos:    801,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 37,
									EndLine:   37,
									StartPos:  801,
									EndPos:    807,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 37,
										EndLine:   37,
										StartPos:  803,
										EndPos:    806,
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
						StartLine: 38,
						EndLine:   38,
						StartPos:  810,
						EndPos:    827,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 38,
							EndLine:   38,
							StartPos:  810,
							EndPos:    826,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 38,
									EndLine:   38,
									StartPos:  811,
									EndPos:    816,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 38,
									EndLine:   38,
									StartPos:  816,
									EndPos:    825,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 38,
										EndLine:   38,
										StartPos:  818,
										EndPos:    821,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 38,
											EndLine:   38,
											StartPos:  818,
											EndPos:    821,
										},
									},
									Value: []byte("foo"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 38,
										EndLine:   38,
										StartPos:  822,
										EndPos:    823,
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
						StartLine: 39,
						EndLine:   39,
						StartPos:  828,
						EndPos:    849,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 39,
							EndLine:   39,
							StartPos:  828,
							EndPos:    848,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  829,
									EndPos:    834,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  835,
									EndPos:    846,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  835,
										EndPos:    839,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 39,
											EndLine:   39,
											StartPos:  835,
											EndPos:    839,
										},
									},
									Value: []byte("$foo"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  841,
										EndPos:    844,
									},
								},
								Value: []byte("bar"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  844,
										EndPos:    846,
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
						StartLine: 41,
						EndLine:   42,
						StartPos:  851,
						EndPos:    867,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 41,
							EndLine:   41,
							StartPos:  855,
							EndPos:    857,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 41,
								EndLine:   41,
								StartPos:  855,
								EndPos:    857,
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
						StartLine: 43,
						EndLine:   45,
						StartPos:  868,
						EndPos:    897,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 43,
							EndLine:   43,
							StartPos:  872,
							EndPos:    874,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 43,
								EndLine:   43,
								StartPos:  872,
								EndPos:    874,
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
								StartLine: 44,
								EndLine:   -1,
								StartPos:  878,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  886,
									EndPos:    888,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  886,
										EndPos:    888,
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
						StartLine: 46,
						EndLine:   48,
						StartPos:  898,
						EndPos:    920,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 46,
							EndLine:   46,
							StartPos:  902,
							EndPos:    904,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  902,
								EndPos:    904,
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
							StartLine: 47,
							EndLine:   -1,
							StartPos:  908,
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
						StartLine: 49,
						EndLine:   53,
						StartPos:  921,
						EndPos:    969,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 49,
							EndLine:   49,
							StartPos:  925,
							EndPos:    927,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 49,
								EndLine:   49,
								StartPos:  925,
								EndPos:    927,
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
								StartLine: 50,
								EndLine:   -1,
								StartPos:  931,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  939,
									EndPos:    941,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 50,
										EndLine:   50,
										StartPos:  939,
										EndPos:    941,
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
								StartLine: 51,
								EndLine:   -1,
								StartPos:  944,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  952,
									EndPos:    954,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  952,
										EndPos:    954,
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
							StartLine: 52,
							EndLine:   -1,
							StartPos:  957,
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
						StartLine: 55,
						EndLine:   55,
						StartPos:  971,
						EndPos:    991,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  978,
							EndPos:    979,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  981,
							EndPos:    991,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 55,
									EndLine:   55,
									StartPos:  983,
									EndPos:    989,
								},
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 56,
						EndLine:   56,
						StartPos:  992,
						EndPos:    1014,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 56,
							EndLine:   56,
							StartPos:  999,
							EndPos:    1000,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 56,
							EndLine:   56,
							StartPos:  1002,
							EndPos:    1014,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1004,
									EndPos:    1012,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 56,
										EndLine:   56,
										StartPos:  1010,
										EndPos:    1011,
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
						StartLine: 57,
						EndLine:   57,
						StartPos:  1015,
						EndPos:    1046,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 57,
							EndLine:   57,
							StartPos:  1022,
							EndPos:    1023,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 57,
							EndLine:   57,
							StartPos:  1027,
							EndPos:    1036,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 57,
									EndLine:   57,
									StartPos:  1027,
									EndPos:    1036,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 57,
										EndLine:   57,
										StartPos:  1033,
										EndPos:    1034,
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
						StartLine: 58,
						EndLine:   58,
						StartPos:  1047,
						EndPos:    1083,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1053,
							EndPos:    1056,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 58,
								EndLine:   58,
								StartPos:  1058,
								EndPos:    1081,
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1064,
										EndPos:    1071,
									},
								},
								Name: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1064,
											EndPos:    1067,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1070,
											EndPos:    1071,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1073,
										EndPos:    1080,
									},
								},
								Name: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1073,
											EndPos:    1076,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1079,
											EndPos:    1080,
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
						StartLine: 59,
						EndLine:   59,
						StartPos:  1084,
						EndPos:    1114,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 59,
							EndLine:   59,
							StartPos:  1090,
							EndPos:    1093,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 59,
								EndLine:   59,
								StartPos:  1095,
								EndPos:    1112,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 59,
									EndLine:   59,
									StartPos:  1104,
									EndPos:    1107,
								},
							},
							Value: []byte("bar"),
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 59,
									EndLine:   59,
									StartPos:  1110,
									EndPos:    1112,
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
						StartLine: 60,
						EndLine:   60,
						StartPos:  1115,
						EndPos:    1160,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 60,
							EndLine:   60,
							StartPos:  1121,
							EndPos:    1124,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 60,
								EndLine:   60,
								StartPos:  1126,
								EndPos:    1158,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 60,
									EndLine:   60,
									StartPos:  1150,
									EndPos:    1153,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 60,
										EndLine:   60,
										StartPos:  1126,
										EndPos:    1132,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 60,
										EndLine:   60,
										StartPos:  1133,
										EndPos:    1139,
									},
								},
								Value: []byte("static"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 60,
									EndLine:   60,
									StartPos:  1156,
									EndPos:    1158,
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
						StartLine: 61,
						EndLine:   61,
						StartPos:  1161,
						EndPos:    1233,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1167,
							EndPos:    1170,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1172,
								EndPos:    1203,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1195,
									EndPos:    1198,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 61,
										EndLine:   61,
										StartPos:  1172,
										EndPos:    1177,
									},
								},
								Value: []byte("final"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 61,
										EndLine:   61,
										StartPos:  1178,
										EndPos:    1185,
									},
								},
								Value: []byte("private"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1201,
									EndPos:    1203,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1204,
								EndPos:    1231,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1223,
									EndPos:    1226,
								},
							},
							Value: []byte("baz"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 61,
										EndLine:   61,
										StartPos:  1204,
										EndPos:    1213,
									},
								},
								Value: []byte("protected"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1229,
									EndPos:    1231,
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
						StartLine: 62,
						EndLine:   62,
						StartPos:  1234,
						EndPos:    1287,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 62,
							EndLine:   62,
							StartPos:  1249,
							EndPos:    1252,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1234,
								EndPos:    1242,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1254,
								EndPos:    1285,
							},
						},
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1279,
									EndPos:    1282,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 62,
										EndLine:   62,
										StartPos:  1254,
										EndPos:    1262,
									},
								},
								Value: []byte("abstract"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 62,
										EndLine:   62,
										StartPos:  1263,
										EndPos:    1269,
									},
								},
								Value: []byte("public"),
							},
						},
						Stmt: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1284,
									EndPos:    1285,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 63,
						EndLine:   63,
						StartPos:  1288,
						EndPos:    1319,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   63,
							StartPos:  1300,
							EndPos:    1303,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 63,
								EndLine:   63,
								StartPos:  1288,
								EndPos:    1293,
							},
						},
						Value: []byte("final"),
					},
				},
				Extends: &ast.StmtClassExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   63,
							StartPos:  1304,
							EndPos:    1315,
						},
					},
					ClassName: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 63,
								EndLine:   63,
								StartPos:  1312,
								EndPos:    1315,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 63,
										EndLine:   63,
										StartPos:  1312,
										EndPos:    1315,
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
						StartLine: 64,
						EndLine:   64,
						StartPos:  1320,
						EndPos:    1354,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 64,
							EndLine:   64,
							StartPos:  1332,
							EndPos:    1335,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 64,
								EndLine:   64,
								StartPos:  1320,
								EndPos:    1325,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 64,
							EndLine:   64,
							StartPos:  1336,
							EndPos:    1350,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 64,
									EndLine:   64,
									StartPos:  1347,
									EndPos:    1350,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 64,
											EndLine:   64,
											StartPos:  1347,
											EndPos:    1350,
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
						StartLine: 65,
						EndLine:   65,
						StartPos:  1355,
						EndPos:    1394,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 65,
							EndLine:   65,
							StartPos:  1367,
							EndPos:    1370,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 65,
								EndLine:   65,
								StartPos:  1355,
								EndPos:    1360,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 65,
							EndLine:   65,
							StartPos:  1371,
							EndPos:    1390,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 65,
									EndLine:   65,
									StartPos:  1382,
									EndPos:    1385,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 65,
											EndLine:   65,
											StartPos:  1382,
											EndPos:    1385,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 65,
									EndLine:   65,
									StartPos:  1387,
									EndPos:    1390,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 65,
											EndLine:   65,
											StartPos:  1387,
											EndPos:    1390,
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
			&ast.StmtConstList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1396,
						EndPos:    1419,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1402,
								EndPos:    1409,
							},
						},
						Name: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1402,
									EndPos:    1405,
								},
							},
							Value: []byte("FOO"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1408,
									EndPos:    1409,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1411,
								EndPos:    1418,
							},
						},
						Name: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1411,
									EndPos:    1414,
								},
							},
							Value: []byte("BAR"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1417,
									EndPos:    1418,
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
						StartLine: 68,
						EndLine:   68,
						StartPos:  1420,
						EndPos:    1443,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1427,
							EndPos:    1428,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1430,
							EndPos:    1443,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 68,
									EndLine:   68,
									StartPos:  1432,
									EndPos:    1441,
								},
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1444,
						EndPos:    1469,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1451,
							EndPos:    1452,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1454,
							EndPos:    1469,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1456,
									EndPos:    1467,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 69,
										EndLine:   69,
										StartPos:  1465,
										EndPos:    1466,
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
						StartLine: 70,
						EndLine:   70,
						StartPos:  1470,
						EndPos:    1496,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1477,
							EndPos:    1478,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1480,
							EndPos:    1496,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1482,
									EndPos:    1494,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1491,
										EndPos:    1492,
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
						StartLine: 71,
						EndLine:   71,
						StartPos:  1497,
						EndPos:    1514,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 71,
								EndLine:   71,
								StartPos:  1505,
								EndPos:    1512,
							},
						},
						Name: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1505,
									EndPos:    1510,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1511,
									EndPos:    1512,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtNop{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1513,
							EndPos:    1514,
						},
					},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 72,
						EndLine:   72,
						StartPos:  1515,
						EndPos:    1550,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1523,
								EndPos:    1530,
							},
						},
						Name: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1523,
									EndPos:    1528,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1529,
									EndPos:    1530,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1532,
								EndPos:    1546,
							},
						},
						Name: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1532,
									EndPos:    1544,
								},
							},
							Value: []byte("strict_types"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1545,
									EndPos:    1546,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1548,
							EndPos:    1550,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 73,
						EndLine:   73,
						StartPos:  1551,
						EndPos:    1580,
					},
				},
				Alt: true,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1559,
								EndPos:    1566,
							},
						},
						Name: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1559,
									EndPos:    1564,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1565,
									EndPos:    1566,
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
						StartLine: 74,
						EndLine:   74,
						StartPos:  1581,
						EndPos:    1596,
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1584,
							EndPos:    1586,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1593,
							EndPos:    1594,
						},
					},
					Value: []byte("1"),
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 75,
						EndLine:   75,
						StartPos:  1597,
						EndPos:    1608,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1602,
								EndPos:    1604,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 75,
									EndLine:   75,
									StartPos:  1602,
									EndPos:    1604,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1606,
								EndPos:    1607,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1609,
						EndPos:    1618,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1614,
								EndPos:    1616,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 76,
									EndLine:   76,
									StartPos:  1614,
									EndPos:    1616,
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
						StartLine: 77,
						EndLine:   77,
						StartPos:  1619,
						EndPos:    1654,
					},
				},
				Init: []ast.Vertex{
					&ast.ExprAssign{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1623,
								EndPos:    1629,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1623,
									EndPos:    1625,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1623,
										EndPos:    1625,
									},
								},
								Value: []byte("$i"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1628,
									EndPos:    1629,
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
								StartLine: 77,
								EndLine:   77,
								StartPos:  1631,
								EndPos:    1638,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1631,
									EndPos:    1633,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1631,
										EndPos:    1633,
									},
								},
								Value: []byte("$i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1636,
									EndPos:    1638,
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
								StartLine: 77,
								EndLine:   77,
								StartPos:  1640,
								EndPos:    1644,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1640,
									EndPos:    1642,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1640,
										EndPos:    1642,
									},
								},
								Value: []byte("$i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1646,
								EndPos:    1650,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1646,
									EndPos:    1648,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1646,
										EndPos:    1648,
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
							StartLine: 77,
							EndLine:   77,
							StartPos:  1652,
							EndPos:    1654,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1655,
						EndPos:    1685,
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1661,
								EndPos:    1668,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1661,
									EndPos:    1663,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1661,
										EndPos:    1663,
									},
								},
								Value: []byte("$i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1666,
									EndPos:    1668,
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
								StartLine: 78,
								EndLine:   78,
								StartPos:  1670,
								EndPos:    1674,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1670,
									EndPos:    1672,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1670,
										EndPos:    1672,
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
						StartLine: 79,
						EndLine:   79,
						StartPos:  1686,
						EndPos:    1707,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1695,
							EndPos:    1697,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1695,
								EndPos:    1697,
							},
						},
						Value: []byte("$a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1701,
							EndPos:    1703,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1701,
								EndPos:    1703,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1705,
							EndPos:    1707,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 80,
						EndLine:   80,
						StartPos:  1708,
						EndPos:    1729,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1717,
							EndPos:    1719,
						},
					},
					Items: []ast.Vertex{},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1723,
							EndPos:    1725,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 80,
								EndLine:   80,
								StartPos:  1723,
								EndPos:    1725,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1727,
							EndPos:    1729,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 81,
						EndLine:   81,
						StartPos:  1730,
						EndPos:    1762,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1739,
							EndPos:    1741,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1739,
								EndPos:    1741,
							},
						},
						Value: []byte("$a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1745,
							EndPos:    1747,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1745,
								EndPos:    1747,
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
						StartLine: 82,
						EndLine:   82,
						StartPos:  1763,
						EndPos:    1790,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1772,
							EndPos:    1774,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1772,
								EndPos:    1774,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1778,
							EndPos:    1780,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1778,
								EndPos:    1780,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1784,
							EndPos:    1786,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1784,
								EndPos:    1786,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1788,
							EndPos:    1790,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1791,
						EndPos:    1818,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1800,
							EndPos:    1802,
						},
					},
					Items: []ast.Vertex{},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1806,
							EndPos:    1808,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 83,
								EndLine:   83,
								StartPos:  1806,
								EndPos:    1808,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1812,
							EndPos:    1814,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 83,
								EndLine:   83,
								StartPos:  1812,
								EndPos:    1814,
							},
						},
						Value: []byte("$v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1816,
							EndPos:    1818,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1819,
						EndPos:    1847,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1828,
							EndPos:    1830,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1828,
								EndPos:    1830,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1834,
							EndPos:    1836,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1834,
								EndPos:    1836,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1840,
							EndPos:    1843,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1841,
								EndPos:    1843,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 84,
									EndLine:   84,
									StartPos:  1841,
									EndPos:    1843,
								},
							},
							Value: []byte("$v"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1845,
							EndPos:    1847,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  1848,
						EndPos:    1881,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1857,
							EndPos:    1859,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1857,
								EndPos:    1859,
							},
						},
						Value: []byte("$a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1863,
							EndPos:    1865,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1863,
								EndPos:    1865,
							},
						},
						Value: []byte("$k"),
					},
				},
				Var: &ast.ExprList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1869,
							EndPos:    1877,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  1874,
									EndPos:    1876,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 85,
										EndLine:   85,
										StartPos:  1874,
										EndPos:    1876,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 85,
											EndLine:   85,
											StartPos:  1874,
											EndPos:    1876,
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
							StartLine: 85,
							EndLine:   85,
							StartPos:  1879,
							EndPos:    1881,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 86,
						EndLine:   86,
						StartPos:  1882,
						EndPos:    1899,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 86,
							EndLine:   86,
							StartPos:  1891,
							EndPos:    1894,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   92,
						StartPos:  1901,
						EndPos:    1973,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 88,
							EndLine:   88,
							StartPos:  1910,
							EndPos:    1913,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtFunction{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  1922,
								EndPos:    1939,
							},
						},
						FunctionName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 89,
									EndLine:   89,
									StartPos:  1931,
									EndPos:    1934,
								},
							},
							Value: []byte("bar"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  1944,
								EndPos:    1956,
							},
						},
						ClassName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 90,
									EndLine:   90,
									StartPos:  1950,
									EndPos:    1953,
								},
							},
							Value: []byte("Baz"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  1961,
								EndPos:    1971,
							},
						},
						Expr: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  1968,
									EndPos:    1970,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  1968,
										EndPos:    1970,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  1975,
						EndPos:    2020,
					},
				},
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  1984,
							EndPos:    1987,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  1988,
								EndPos:    1996,
							},
						},
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  1988,
									EndPos:    1993,
								},
							},
							Value: []byte("array"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  1994,
									EndPos:    1996,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 94,
										EndLine:   94,
										StartPos:  1994,
										EndPos:    1996,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  1998,
								EndPos:    2009,
							},
						},
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  1998,
									EndPos:    2006,
								},
							},
							Value: []byte("callable"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2007,
									EndPos:    2009,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 94,
										EndLine:   94,
										StartPos:  2007,
										EndPos:    2009,
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
								StartLine: 94,
								EndLine:   94,
								StartPos:  2012,
								EndPos:    2019,
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2021,
						EndPos:    2048,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2031,
							EndPos:    2034,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2038,
								EndPos:    2047,
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 95,
									EndLine:   95,
									StartPos:  2045,
									EndPos:    2046,
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
						StartLine: 96,
						EndLine:   96,
						StartPos:  2049,
						EndPos:    2067,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2059,
							EndPos:    2062,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtGlobal{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2068,
						EndPos:    2097,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2075,
								EndPos:    2077,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2075,
									EndPos:    2077,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2079,
								EndPos:    2081,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2079,
									EndPos:    2081,
								},
							},
							Value: []byte("$b"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2083,
								EndPos:    2086,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2084,
									EndPos:    2086,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2084,
										EndPos:    2086,
									},
								},
								Value: []byte("$c"),
							},
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2088,
								EndPos:    2096,
							},
						},
						VarName: &ast.ExprFunctionCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2090,
									EndPos:    2095,
								},
							},
							Function: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2090,
										EndPos:    2093,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 97,
												EndLine:   97,
												StartPos:  2090,
												EndPos:    2093,
											},
										},
										Value: []byte("foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2093,
										EndPos:    2095,
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtLabel{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2098,
						EndPos:    2100,
					},
				},
				LabelName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2098,
							EndPos:    2099,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtGoto{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2101,
						EndPos:    2108,
					},
				},
				Label: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 99,
							EndLine:   99,
							StartPos:  2106,
							EndPos:    2107,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2109,
						EndPos:    2119,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2113,
							EndPos:    2115,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 100,
								EndLine:   100,
								StartPos:  2113,
								EndPos:    2115,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2117,
							EndPos:    2119,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2120,
						EndPos:    2145,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2124,
							EndPos:    2126,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2124,
								EndPos:    2126,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2128,
							EndPos:    2130,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2131,
								EndPos:    2145,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2139,
									EndPos:    2141,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 101,
										EndLine:   101,
										StartPos:  2139,
										EndPos:    2141,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2143,
									EndPos:    2145,
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
						StartLine: 102,
						EndLine:   102,
						StartPos:  2146,
						EndPos:    2164,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2150,
							EndPos:    2152,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 102,
								EndLine:   102,
								StartPos:  2150,
								EndPos:    2152,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2154,
							EndPos:    2156,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2157,
							EndPos:    2164,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 102,
								EndLine:   102,
								StartPos:  2162,
								EndPos:    2164,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2165,
						EndPos:    2213,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2169,
							EndPos:    2171,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2169,
								EndPos:    2171,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2173,
							EndPos:    2175,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2176,
								EndPos:    2190,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2184,
									EndPos:    2186,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 103,
										EndLine:   103,
										StartPos:  2184,
										EndPos:    2186,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2188,
									EndPos:    2190,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2191,
								EndPos:    2205,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2199,
									EndPos:    2201,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 103,
										EndLine:   103,
										StartPos:  2199,
										EndPos:    2201,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2203,
									EndPos:    2205,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2206,
							EndPos:    2213,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2211,
								EndPos:    2213,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2214,
						EndPos:    2263,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2218,
							EndPos:    2220,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2218,
								EndPos:    2220,
							},
						},
						Value: []byte("$a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2222,
							EndPos:    2224,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2225,
								EndPos:    2239,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2233,
									EndPos:    2235,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 104,
										EndLine:   104,
										StartPos:  2233,
										EndPos:    2235,
									},
								},
								Value: []byte("$b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2237,
									EndPos:    2239,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2240,
							EndPos:    2263,
						},
					},
					Stmt: &ast.StmtIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2245,
								EndPos:    2263,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2249,
									EndPos:    2251,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 104,
										EndLine:   104,
										StartPos:  2249,
										EndPos:    2251,
									},
								},
								Value: []byte("$c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2253,
									EndPos:    2255,
								},
							},
							Stmts: []ast.Vertex{},
						},
						Else: &ast.StmtElse{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2256,
									EndPos:    2263,
								},
							},
							Stmt: &ast.StmtStmtList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 104,
										EndLine:   104,
										StartPos:  2261,
										EndPos:    2263,
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
						StartLine: 105,
						EndLine:   105,
						StartPos:  2264,
						EndPos:    2266,
					},
				},
			},
			&ast.StmtInlineHtml{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2266,
						EndPos:    2279,
					},
				},
				Value: []byte(" <div></div> "),
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2282,
						EndPos:    2298,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2292,
							EndPos:    2295,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2299,
						EndPos:    2327,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2309,
							EndPos:    2312,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2313,
							EndPos:    2324,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2321,
									EndPos:    2324,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 107,
											EndLine:   107,
											StartPos:  2321,
											EndPos:    2324,
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
						StartLine: 108,
						EndLine:   108,
						StartPos:  2328,
						EndPos:    2361,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2338,
							EndPos:    2341,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2342,
							EndPos:    2358,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 108,
									EndLine:   108,
									StartPos:  2350,
									EndPos:    2353,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 108,
											EndLine:   108,
											StartPos:  2350,
											EndPos:    2353,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 108,
									EndLine:   108,
									StartPos:  2355,
									EndPos:    2358,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 108,
											EndLine:   108,
											StartPos:  2355,
											EndPos:    2358,
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
						StartLine: 109,
						EndLine:   109,
						StartPos:  2362,
						EndPos:    2376,
					},
				},
				Name: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2372,
							EndPos:    2375,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2372,
									EndPos:    2375,
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
						StartLine: 110,
						EndLine:   110,
						StartPos:  2377,
						EndPos:    2397,
					},
				},
				Name: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2387,
							EndPos:    2394,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2387,
									EndPos:    2390,
								},
							},
							Value: []byte("Foo"),
						},
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2390,
									EndPos:    2394,
								},
							},
							Value: []byte("Bar"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2398,
						EndPos:    2410,
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2411,
						EndPos:    2430,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 112,
							EndLine:   112,
							StartPos:  2417,
							EndPos:    2420,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 112,
								EndLine:   112,
								StartPos:  2422,
								EndPos:    2429,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 112,
										EndLine:   112,
										StartPos:  2422,
										EndPos:    2425,
									},
								},
								Value: []byte("var"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 112,
										EndLine:   112,
										StartPos:  2426,
										EndPos:    2428,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 112,
											EndLine:   112,
											StartPos:  2426,
											EndPos:    2428,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 112,
												EndLine:   112,
												StartPos:  2426,
												EndPos:    2428,
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
						StartLine: 113,
						EndLine:   113,
						StartPos:  2431,
						EndPos:    2468,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2437,
							EndPos:    2440,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 113,
								EndLine:   113,
								StartPos:  2442,
								EndPos:    2467,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2442,
										EndPos:    2448,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2449,
										EndPos:    2455,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2456,
										EndPos:    2458,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2456,
											EndPos:    2458,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 113,
												EndLine:   113,
												StartPos:  2456,
												EndPos:    2458,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2460,
										EndPos:    2466,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2460,
											EndPos:    2462,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 113,
												EndLine:   113,
												StartPos:  2460,
												EndPos:    2462,
											},
										},
										Value: []byte("$b"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2465,
											EndPos:    2466,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2469,
						EndPos:    2506,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2475,
							EndPos:    2478,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 114,
								EndLine:   114,
								StartPos:  2480,
								EndPos:    2505,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2480,
										EndPos:    2486,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2487,
										EndPos:    2493,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2494,
										EndPos:    2500,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2494,
											EndPos:    2496,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 114,
												EndLine:   114,
												StartPos:  2494,
												EndPos:    2496,
											},
										},
										Value: []byte("$a"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2499,
											EndPos:    2500,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2502,
										EndPos:    2504,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2502,
											EndPos:    2504,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 114,
												EndLine:   114,
												StartPos:  2502,
												EndPos:    2504,
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
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2507,
						EndPos:    2525,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2514,
								EndPos:    2516,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2514,
									EndPos:    2516,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2514,
										EndPos:    2516,
									},
								},
								Value: []byte("$a"),
							},
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2518,
								EndPos:    2524,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2518,
									EndPos:    2520,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2518,
										EndPos:    2520,
									},
								},
								Value: []byte("$b"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2523,
									EndPos:    2524,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 116,
						EndLine:   116,
						StartPos:  2526,
						EndPos:    2544,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2533,
								EndPos:    2539,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2533,
									EndPos:    2535,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 116,
										EndLine:   116,
										StartPos:  2533,
										EndPos:    2535,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2538,
									EndPos:    2539,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2541,
								EndPos:    2543,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2541,
									EndPos:    2543,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 116,
										EndLine:   116,
										StartPos:  2541,
										EndPos:    2543,
									},
								},
								Value: []byte("$b"),
							},
						},
					},
				},
			},
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 118,
						EndLine:   122,
						StartPos:  2546,
						EndPos:    2606,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2554,
							EndPos:    2555,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 119,
							EndLine:   -1,
							StartPos:  2563,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   -1,
									StartPos:  2563,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2568,
										EndPos:    2569,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtDefault{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   -1,
									StartPos:  2575,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 121,
									EndLine:   -1,
									StartPos:  2588,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 121,
										EndLine:   121,
										StartPos:  2593,
										EndPos:    2594,
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
						StartLine: 124,
						EndLine:   127,
						StartPos:  2608,
						EndPos:    2656,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 124,
							EndLine:   124,
							StartPos:  2616,
							EndPos:    2617,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 125,
							EndLine:   -1,
							StartPos:  2626,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 125,
									EndLine:   -1,
									StartPos:  2626,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 125,
										EndLine:   125,
										StartPos:  2631,
										EndPos:    2632,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 126,
									EndLine:   -1,
									StartPos:  2638,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 126,
										EndLine:   126,
										StartPos:  2643,
										EndPos:    2644,
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
						StartLine: 129,
						EndLine:   132,
						StartPos:  2658,
						EndPos:    2710,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 129,
							EndLine:   129,
							StartPos:  2666,
							EndPos:    2667,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 129,
							EndLine:   132,
							StartPos:  2669,
							EndPos:    2710,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   130,
									StartPos:  2675,
									EndPos:    2689,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 130,
										EndLine:   130,
										StartPos:  2680,
										EndPos:    2681,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 130,
											EndLine:   130,
											StartPos:  2683,
											EndPos:    2689,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 131,
									EndLine:   131,
									StartPos:  2694,
									EndPos:    2708,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 131,
										EndLine:   131,
										StartPos:  2699,
										EndPos:    2700,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 131,
											EndLine:   131,
											StartPos:  2702,
											EndPos:    2708,
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
						StartLine: 134,
						EndLine:   137,
						StartPos:  2712,
						EndPos:    2765,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 134,
							EndLine:   134,
							StartPos:  2720,
							EndPos:    2721,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 134,
							EndLine:   137,
							StartPos:  2723,
							EndPos:    2765,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  2730,
									EndPos:    2744,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  2735,
										EndPos:    2736,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 135,
											EndLine:   135,
											StartPos:  2738,
											EndPos:    2744,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 136,
									EndLine:   136,
									StartPos:  2749,
									EndPos:    2763,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 136,
										EndLine:   136,
										StartPos:  2754,
										EndPos:    2755,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 136,
											EndLine:   136,
											StartPos:  2757,
											EndPos:    2763,
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
						StartLine: 138,
						EndLine:   138,
						StartPos:  2766,
						EndPos:    2775,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   138,
							StartPos:  2772,
							EndPos:    2774,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 138,
								EndLine:   138,
								StartPos:  2772,
								EndPos:    2774,
							},
						},
						Value: []byte("$e"),
					},
				},
			},
			&ast.StmtTrait{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 139,
						EndLine:   139,
						StartPos:  2776,
						EndPos:    2788,
					},
				},
				TraitName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 139,
							EndLine:   139,
							StartPos:  2782,
							EndPos:    2785,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 140,
						EndLine:   140,
						StartPos:  2789,
						EndPos:    2811,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 140,
							EndLine:   140,
							StartPos:  2795,
							EndPos:    2798,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 140,
								EndLine:   140,
								StartPos:  2801,
								EndPos:    2809,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 140,
										EndLine:   140,
										StartPos:  2805,
										EndPos:    2808,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 140,
												EndLine:   140,
												StartPos:  2805,
												EndPos:    2808,
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
									StartLine: 140,
									EndLine:   140,
									StartPos:  2808,
									EndPos:    2809,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 141,
						EndLine:   141,
						StartPos:  2812,
						EndPos:    2841,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 141,
							EndLine:   141,
							StartPos:  2818,
							EndPos:    2821,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 141,
								EndLine:   141,
								StartPos:  2824,
								EndPos:    2839,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 141,
										EndLine:   141,
										StartPos:  2828,
										EndPos:    2831,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 141,
												EndLine:   141,
												StartPos:  2828,
												EndPos:    2831,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 141,
										EndLine:   141,
										StartPos:  2833,
										EndPos:    2836,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 141,
												EndLine:   141,
												StartPos:  2833,
												EndPos:    2836,
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
									StartLine: 141,
									EndLine:   141,
									StartPos:  2837,
									EndPos:    2839,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 142,
						EndLine:   142,
						StartPos:  2842,
						EndPos:    2887,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 142,
							EndLine:   142,
							StartPos:  2848,
							EndPos:    2851,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 142,
								EndLine:   142,
								StartPos:  2854,
								EndPos:    2885,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 142,
										EndLine:   142,
										StartPos:  2858,
										EndPos:    2861,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  2858,
												EndPos:    2861,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 142,
										EndLine:   142,
										StartPos:  2863,
										EndPos:    2866,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  2863,
												EndPos:    2866,
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
									StartLine: 142,
									EndLine:   142,
									StartPos:  2867,
									EndPos:    2885,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  2869,
											EndPos:    2882,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  2869,
												EndPos:    2872,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 142,
													EndLine:   142,
													StartPos:  2869,
													EndPos:    2872,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  2876,
												EndPos:    2882,
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
						StartLine: 143,
						EndLine:   143,
						StartPos:  2888,
						EndPos:    2937,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  2894,
							EndPos:    2897,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 143,
								EndLine:   143,
								StartPos:  2900,
								EndPos:    2935,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 143,
										EndLine:   143,
										StartPos:  2904,
										EndPos:    2907,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  2904,
												EndPos:    2907,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 143,
										EndLine:   143,
										StartPos:  2909,
										EndPos:    2912,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  2909,
												EndPos:    2912,
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
									StartLine: 143,
									EndLine:   143,
									StartPos:  2913,
									EndPos:    2935,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  2915,
											EndPos:    2932,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  2915,
												EndPos:    2918,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 143,
													EndLine:   143,
													StartPos:  2915,
													EndPos:    2918,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  2922,
												EndPos:    2928,
											},
										},
										Value: []byte("public"),
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  2929,
												EndPos:    2932,
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
						StartLine: 144,
						EndLine:   144,
						StartPos:  2938,
						EndPos:    3015,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 144,
							EndLine:   144,
							StartPos:  2944,
							EndPos:    2947,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 144,
								EndLine:   144,
								StartPos:  2950,
								EndPos:    3013,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  2954,
										EndPos:    2957,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  2954,
												EndPos:    2957,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  2959,
										EndPos:    2962,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  2959,
												EndPos:    2962,
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
									StartLine: 144,
									EndLine:   144,
									StartPos:  2963,
									EndPos:    3013,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUsePrecedence{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  2965,
											EndPos:    2993,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  2965,
												EndPos:    2973,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  2965,
													EndPos:    2968,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  2965,
															EndPos:    2968,
														},
													},
													Value: []byte("Bar"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  2970,
													EndPos:    2973,
												},
											},
											Value: []byte("one"),
										},
									},
									Insteadof: []ast.Vertex{
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  2984,
													EndPos:    2987,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  2984,
															EndPos:    2987,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  2989,
													EndPos:    2993,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  2989,
															EndPos:    2993,
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
											StartLine: 144,
											EndLine:   144,
											StartPos:  2995,
											EndPos:    3010,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  2995,
												EndPos:    3003,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  2995,
													EndPos:    2998,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  2995,
															EndPos:    2998,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3000,
													EndPos:    3003,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3007,
												EndPos:    3010,
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
						StartLine: 146,
						EndLine:   -1,
						StartPos:  3017,
						EndPos:    -1,
					},
				},
				Stmts:   []ast.Vertex{},
				Catches: []ast.Vertex{},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 147,
						EndLine:   147,
						StartPos:  3024,
						EndPos:    3054,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  3031,
								EndPos:    3054,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  3038,
										EndPos:    3047,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 147,
												EndLine:   147,
												StartPos:  3038,
												EndPos:    3047,
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
									StartLine: 147,
									EndLine:   147,
									StartPos:  3048,
									EndPos:    3050,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  3048,
										EndPos:    3050,
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
						StartLine: 148,
						EndLine:   148,
						StartPos:  3055,
						EndPos:    3116,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3062,
								EndPos:    3085,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3069,
										EndPos:    3078,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3069,
												EndPos:    3078,
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
									StartLine: 148,
									EndLine:   148,
									StartPos:  3079,
									EndPos:    3081,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3079,
										EndPos:    3081,
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
								StartLine: 148,
								EndLine:   148,
								StartPos:  3086,
								EndPos:    3116,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3093,
										EndPos:    3109,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3093,
												EndPos:    3109,
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
									StartLine: 148,
									EndLine:   148,
									StartPos:  3110,
									EndPos:    3112,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3110,
										EndPos:    3112,
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
						StartLine: 149,
						EndLine:   149,
						StartPos:  3117,
						EndPos:    3221,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3124,
								EndPos:    3147,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3131,
										EndPos:    3140,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3131,
												EndPos:    3140,
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
									StartLine: 149,
									EndLine:   149,
									StartPos:  3141,
									EndPos:    3143,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3141,
										EndPos:    3143,
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
								StartLine: 149,
								EndLine:   149,
								StartPos:  3148,
								EndPos:    3179,
							},
						},
						Types: []ast.Vertex{
							&ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3155,
										EndPos:    3172,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3156,
												EndPos:    3172,
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
									StartLine: 149,
									EndLine:   149,
									StartPos:  3173,
									EndPos:    3175,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3173,
										EndPos:    3175,
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
								StartLine: 149,
								EndLine:   149,
								StartPos:  3180,
								EndPos:    3221,
							},
						},
						Types: []ast.Vertex{
							&ast.NameRelative{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3187,
										EndPos:    3214,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3197,
												EndPos:    3214,
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
									StartLine: 149,
									EndLine:   149,
									StartPos:  3215,
									EndPos:    3217,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3215,
										EndPos:    3217,
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
						StartLine: 150,
						EndLine:   150,
						StartPos:  3222,
						EndPos:    3263,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3229,
								EndPos:    3252,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3236,
										EndPos:    3245,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3236,
												EndPos:    3245,
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
									StartLine: 150,
									EndLine:   150,
									StartPos:  3246,
									EndPos:    3248,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3246,
										EndPos:    3248,
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
							StartLine: 150,
							EndLine:   150,
							StartPos:  3253,
							EndPos:    3263,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 152,
						EndLine:   152,
						StartPos:  3265,
						EndPos:    3279,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3271,
								EndPos:    3273,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3271,
									EndPos:    3273,
								},
							},
							Value: []byte("$a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3275,
								EndPos:    3277,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3275,
									EndPos:    3277,
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
						StartLine: 154,
						EndLine:   154,
						StartPos:  3281,
						EndPos:    3289,
					},
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3285,
								EndPos:    3288,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 154,
									EndLine:   154,
									StartPos:  3285,
									EndPos:    3288,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 154,
											EndLine:   154,
											StartPos:  3285,
											EndPos:    3288,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 155,
						EndLine:   155,
						StartPos:  3290,
						EndPos:    3299,
					},
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3294,
								EndPos:    3298,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3295,
									EndPos:    3298,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 155,
											EndLine:   155,
											StartPos:  3295,
											EndPos:    3298,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 156,
						EndLine:   156,
						StartPos:  3300,
						EndPos:    3316,
					},
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3304,
								EndPos:    3315,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3305,
									EndPos:    3308,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 156,
											EndLine:   156,
											StartPos:  3305,
											EndPos:    3308,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3312,
									EndPos:    3315,
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
						StartLine: 157,
						EndLine:   157,
						StartPos:  3317,
						EndPos:    3330,
					},
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3321,
								EndPos:    3324,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3321,
									EndPos:    3324,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3321,
											EndPos:    3324,
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
								StartLine: 157,
								EndLine:   157,
								StartPos:  3326,
								EndPos:    3329,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3326,
									EndPos:    3329,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3326,
											EndPos:    3329,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 158,
						EndLine:   158,
						StartPos:  3331,
						EndPos:    3351,
					},
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3335,
								EndPos:    3338,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3335,
									EndPos:    3338,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 158,
											EndLine:   158,
											StartPos:  3335,
											EndPos:    3338,
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
								StartLine: 158,
								EndLine:   158,
								StartPos:  3340,
								EndPos:    3350,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3340,
									EndPos:    3343,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 158,
											EndLine:   158,
											StartPos:  3340,
											EndPos:    3343,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3347,
									EndPos:    3350,
								},
							},
							Value: []byte("Baz"),
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 159,
						EndLine:   159,
						StartPos:  3352,
						EndPos:    3375,
					},
				},
				Type: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3356,
							EndPos:    3364,
						},
					},
					Value: []byte("function"),
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3365,
								EndPos:    3368,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3365,
									EndPos:    3368,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 159,
											EndLine:   159,
											StartPos:  3365,
											EndPos:    3368,
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
								StartLine: 159,
								EndLine:   159,
								StartPos:  3370,
								EndPos:    3374,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3371,
									EndPos:    3374,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 159,
											EndLine:   159,
											StartPos:  3371,
											EndPos:    3374,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 160,
						EndLine:   160,
						StartPos:  3376,
						EndPos:    3413,
					},
				},
				Type: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3380,
							EndPos:    3388,
						},
					},
					Value: []byte("function"),
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3389,
								EndPos:    3399,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3389,
									EndPos:    3392,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 160,
											EndLine:   160,
											StartPos:  3389,
											EndPos:    3392,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3396,
									EndPos:    3399,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3401,
								EndPos:    3412,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3402,
									EndPos:    3405,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 160,
											EndLine:   160,
											StartPos:  3402,
											EndPos:    3405,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3409,
									EndPos:    3412,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 161,
						EndLine:   161,
						StartPos:  3414,
						EndPos:    3434,
					},
				},
				Type: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3418,
							EndPos:    3423,
						},
					},
					Value: []byte("const"),
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3424,
								EndPos:    3427,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 161,
									EndLine:   161,
									StartPos:  3424,
									EndPos:    3427,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 161,
											EndLine:   161,
											StartPos:  3424,
											EndPos:    3427,
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
								StartLine: 161,
								EndLine:   161,
								StartPos:  3429,
								EndPos:    3433,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 161,
									EndLine:   161,
									StartPos:  3430,
									EndPos:    3433,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 161,
											EndLine:   161,
											StartPos:  3430,
											EndPos:    3433,
										},
									},
									Value: []byte("Bar"),
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
						StartPos:  3435,
						EndPos:    3469,
					},
				},
				Type: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3439,
							EndPos:    3444,
						},
					},
					Value: []byte("const"),
				},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3445,
								EndPos:    3455,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3445,
									EndPos:    3448,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 162,
											EndLine:   162,
											StartPos:  3445,
											EndPos:    3448,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3452,
									EndPos:    3455,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUseDeclaration{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3457,
								EndPos:    3468,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3458,
									EndPos:    3461,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 162,
											EndLine:   162,
											StartPos:  3458,
											EndPos:    3461,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3465,
									EndPos:    3468,
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
						StartLine: 164,
						EndLine:   164,
						StartPos:  3471,
						EndPos:    3477,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3471,
							EndPos:    3476,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3471,
								EndPos:    3473,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 164,
									EndLine:   164,
									StartPos:  3471,
									EndPos:    3473,
								},
							},
							Value: []byte("$a"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3474,
								EndPos:    3475,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 165,
						EndLine:   165,
						StartPos:  3478,
						EndPos:    3487,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3478,
							EndPos:    3486,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3478,
								EndPos:    3483,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3478,
									EndPos:    3480,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3478,
										EndPos:    3480,
									},
								},
								Value: []byte("$a"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3481,
									EndPos:    3482,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3484,
								EndPos:    3485,
							},
						},
						Value: []byte("2"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 166,
						EndLine:   166,
						StartPos:  3488,
						EndPos:    3496,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3488,
							EndPos:    3495,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3497,
						EndPos:    3506,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3497,
							EndPos:    3505,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3503,
									EndPos:    3504,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3503,
										EndPos:    3504,
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
						StartLine: 168,
						EndLine:   168,
						StartPos:  3507,
						EndPos:    3525,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3507,
							EndPos:    3524,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3513,
									EndPos:    3517,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3513,
										EndPos:    3514,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3516,
										EndPos:    3517,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3519,
									EndPos:    3522,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3519,
										EndPos:    3522,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3520,
											EndPos:    3522,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 168,
												EndLine:   168,
												StartPos:  3520,
												EndPos:    3522,
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
						StartLine: 169,
						EndLine:   169,
						StartPos:  3526,
						EndPos:    3541,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3526,
							EndPos:    3540,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3532,
									EndPos:    3539,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3532,
										EndPos:    3533,
									},
								},
								Value: []byte("3"),
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3536,
										EndPos:    3539,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3537,
											EndPos:    3539,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 169,
												EndLine:   169,
												StartPos:  3537,
												EndPos:    3539,
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
						StartLine: 170,
						EndLine:   170,
						StartPos:  3542,
						EndPos:    3571,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 170,
							EndLine:   170,
							StartPos:  3542,
							EndPos:    3570,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3548,
									EndPos:    3551,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3548,
										EndPos:    3551,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3549,
											EndPos:    3551,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 170,
												EndLine:   170,
												StartPos:  3549,
												EndPos:    3551,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3553,
									EndPos:    3557,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3553,
										EndPos:    3554,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3556,
										EndPos:    3557,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3559,
									EndPos:    3560,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3559,
										EndPos:    3560,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3562,
									EndPos:    3569,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3562,
										EndPos:    3563,
									},
								},
								Value: []byte("3"),
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3566,
										EndPos:    3569,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3567,
											EndPos:    3569,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 170,
												EndLine:   170,
												StartPos:  3567,
												EndPos:    3569,
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
						StartLine: 171,
						EndLine:   171,
						StartPos:  3572,
						EndPos:    3576,
					},
				},
				Expr: &ast.ExprBitwiseNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3572,
							EndPos:    3575,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3573,
								EndPos:    3575,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3573,
									EndPos:    3575,
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
						StartLine: 172,
						EndLine:   172,
						StartPos:  3577,
						EndPos:    3581,
					},
				},
				Expr: &ast.ExprBooleanNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3577,
							EndPos:    3580,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3578,
								EndPos:    3580,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3578,
									EndPos:    3580,
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
						StartLine: 174,
						EndLine:   174,
						StartPos:  3583,
						EndPos:    3592,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3583,
							EndPos:    3591,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3583,
								EndPos:    3586,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 174,
										EndLine:   174,
										StartPos:  3583,
										EndPos:    3586,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3588,
								EndPos:    3591,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3593,
						EndPos:    3603,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3593,
							EndPos:    3602,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3599,
								EndPos:    3601,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3599,
									EndPos:    3601,
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
						StartLine: 176,
						EndLine:   176,
						StartPos:  3604,
						EndPos:    3613,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  3604,
							EndPos:    3612,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3610,
								EndPos:    3612,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 176,
									EndLine:   176,
									StartPos:  3610,
									EndPos:    3612,
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
						StartLine: 177,
						EndLine:   177,
						StartPos:  3614,
						EndPos:    3627,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3614,
							EndPos:    3626,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3628,
						EndPos:    3662,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3628,
							EndPos:    3661,
						},
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3637,
									EndPos:    3639,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3637,
										EndPos:    3639,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3637,
											EndPos:    3639,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3641,
									EndPos:    3643,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3641,
										EndPos:    3643,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3641,
											EndPos:    3643,
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
								StartLine: 178,
								EndLine:   178,
								StartPos:  3645,
								EndPos:    3658,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3650,
										EndPos:    3652,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3650,
											EndPos:    3652,
										},
									},
									Value: []byte("$c"),
								},
							},
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3654,
										EndPos:    3657,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3655,
											EndPos:    3657,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 178,
												EndLine:   178,
												StartPos:  3655,
												EndPos:    3657,
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
						StartLine: 179,
						EndLine:   179,
						StartPos:  3663,
						EndPos:    3697,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  3663,
							EndPos:    3696,
						},
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3672,
									EndPos:    3674,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3672,
										EndPos:    3674,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3672,
											EndPos:    3674,
										},
									},
									Value: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3676,
									EndPos:    3678,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3676,
										EndPos:    3678,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3676,
											EndPos:    3678,
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
								StartLine: 179,
								EndLine:   179,
								StartPos:  3680,
								EndPos:    3693,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3685,
										EndPos:    3688,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3686,
											EndPos:    3688,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 179,
												EndLine:   179,
												StartPos:  3686,
												EndPos:    3688,
											},
										},
										Value: []byte("$c"),
									},
								},
							},
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3690,
										EndPos:    3692,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3690,
											EndPos:    3692,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  3698,
						EndPos:    3712,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 180,
							EndLine:   180,
							StartPos:  3698,
							EndPos:    3711,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  3713,
						EndPos:    3717,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 181,
							EndLine:   181,
							StartPos:  3713,
							EndPos:    3716,
						},
					},
					Const: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  3713,
								EndPos:    3716,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  3713,
										EndPos:    3716,
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
						StartLine: 182,
						EndLine:   182,
						StartPos:  3718,
						EndPos:    3732,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 182,
							EndLine:   182,
							StartPos:  3718,
							EndPos:    3731,
						},
					},
					Const: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  3718,
								EndPos:    3731,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 182,
										EndLine:   182,
										StartPos:  3728,
										EndPos:    3731,
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
						StartLine: 183,
						EndLine:   183,
						StartPos:  3733,
						EndPos:    3738,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  3733,
							EndPos:    3737,
						},
					},
					Const: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 183,
								EndLine:   183,
								StartPos:  3733,
								EndPos:    3737,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 183,
										EndLine:   183,
										StartPos:  3734,
										EndPos:    3737,
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
						StartLine: 185,
						EndLine:   185,
						StartPos:  3740,
						EndPos:    3750,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  3740,
							EndPos:    3749,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  3746,
								EndPos:    3748,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 185,
									EndLine:   185,
									StartPos:  3746,
									EndPos:    3748,
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
						StartLine: 186,
						EndLine:   186,
						StartPos:  3751,
						EndPos:    3762,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  3751,
							EndPos:    3761,
						},
					},
					Expr: &ast.ExprConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  3757,
								EndPos:    3760,
							},
						},
						Const: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 186,
									EndLine:   186,
									StartPos:  3757,
									EndPos:    3760,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 186,
											EndLine:   186,
											StartPos:  3757,
											EndPos:    3760,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  3763,
						EndPos:    3767,
					},
				},
				Expr: &ast.ExprErrorSuppress{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  3763,
							EndPos:    3766,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  3764,
								EndPos:    3766,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 187,
									EndLine:   187,
									StartPos:  3764,
									EndPos:    3766,
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
						StartPos:  3768,
						EndPos:    3777,
					},
				},
				Expr: &ast.ExprEval{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  3768,
							EndPos:    3776,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  3773,
								EndPos:    3775,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 188,
									EndLine:   188,
									StartPos:  3773,
									EndPos:    3775,
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
						StartPos:  3778,
						EndPos:    3783,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 189,
							EndLine:   189,
							StartPos:  3778,
							EndPos:    3782,
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  3784,
						EndPos:    3793,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  3784,
							EndPos:    3792,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 190,
								EndLine:   190,
								StartPos:  3789,
								EndPos:    3791,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  3789,
									EndPos:    3791,
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
						StartLine: 191,
						EndLine:   191,
						StartPos:  3794,
						EndPos:    3800,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 191,
							EndLine:   191,
							StartPos:  3794,
							EndPos:    3799,
						},
					},
					Die: true,
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  3801,
						EndPos:    3809,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  3801,
							EndPos:    3808,
						},
					},
					Die: true,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 192,
								EndLine:   192,
								StartPos:  3805,
								EndPos:    3807,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 192,
									EndLine:   192,
									StartPos:  3805,
									EndPos:    3807,
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
						StartLine: 193,
						EndLine:   193,
						StartPos:  3810,
						EndPos:    3816,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  3810,
							EndPos:    3815,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 193,
								EndLine:   193,
								StartPos:  3810,
								EndPos:    3813,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 193,
										EndLine:   193,
										StartPos:  3810,
										EndPos:    3813,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 193,
								EndLine:   193,
								StartPos:  3813,
								EndPos:    3815,
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
						StartPos:  3817,
						EndPos:    3836,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  3817,
							EndPos:    3835,
						},
					},
					Function: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  3817,
								EndPos:    3830,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  3827,
										EndPos:    3830,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  3830,
								EndPos:    3835,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  3832,
										EndPos:    3834,
									},
								},
								IsReference: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 194,
											EndLine:   194,
											StartPos:  3832,
											EndPos:    3834,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 194,
												EndLine:   194,
												StartPos:  3832,
												EndPos:    3834,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 195,
						EndLine:   195,
						StartPos:  3837,
						EndPos:    3846,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  3837,
							EndPos:    3845,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 195,
								EndLine:   195,
								StartPos:  3837,
								EndPos:    3841,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 195,
										EndLine:   195,
										StartPos:  3838,
										EndPos:    3841,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 195,
								EndLine:   195,
								StartPos:  3841,
								EndPos:    3845,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 195,
										EndLine:   195,
										StartPos:  3842,
										EndPos:    3844,
									},
								},
								Expr: &ast.ExprShortArray{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 195,
											EndLine:   195,
											StartPos:  3842,
											EndPos:    3844,
										},
									},
									Items: []ast.Vertex{},
								},
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
						StartPos:  3847,
						EndPos:    3862,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  3847,
							EndPos:    3861,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  3847,
								EndPos:    3851,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  3847,
									EndPos:    3851,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  3851,
								EndPos:    3861,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 196,
										EndLine:   196,
										StartPos:  3852,
										EndPos:    3860,
									},
								},
								Expr: &ast.ExprYield{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 196,
											EndLine:   196,
											StartPos:  3852,
											EndPos:    3860,
										},
									},
									Value: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 196,
												EndLine:   196,
												StartPos:  3858,
												EndPos:    3860,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 196,
													EndLine:   196,
													StartPos:  3858,
													EndPos:    3860,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  3864,
						EndPos:    3869,
					},
				},
				Expr: &ast.ExprPostDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  3864,
							EndPos:    3868,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  3864,
								EndPos:    3866,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 198,
									EndLine:   198,
									StartPos:  3864,
									EndPos:    3866,
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
						StartPos:  3870,
						EndPos:    3875,
					},
				},
				Expr: &ast.ExprPostInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 199,
							EndLine:   199,
							StartPos:  3870,
							EndPos:    3874,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 199,
								EndLine:   199,
								StartPos:  3870,
								EndPos:    3872,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 199,
									EndLine:   199,
									StartPos:  3870,
									EndPos:    3872,
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
						StartLine: 200,
						EndLine:   200,
						StartPos:  3876,
						EndPos:    3881,
					},
				},
				Expr: &ast.ExprPreDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  3876,
							EndPos:    3880,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  3878,
								EndPos:    3880,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 200,
									EndLine:   200,
									StartPos:  3878,
									EndPos:    3880,
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
						StartPos:  3882,
						EndPos:    3887,
					},
				},
				Expr: &ast.ExprPreInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  3882,
							EndPos:    3886,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 201,
								EndLine:   201,
								StartPos:  3884,
								EndPos:    3886,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 201,
									EndLine:   201,
									StartPos:  3884,
									EndPos:    3886,
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
						StartPos:  3889,
						EndPos:    3900,
					},
				},
				Expr: &ast.ExprInclude{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  3889,
							EndPos:    3899,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  3897,
								EndPos:    3899,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 203,
									EndLine:   203,
									StartPos:  3897,
									EndPos:    3899,
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
						StartLine: 204,
						EndLine:   204,
						StartPos:  3901,
						EndPos:    3917,
					},
				},
				Expr: &ast.ExprIncludeOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  3901,
							EndPos:    3916,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  3914,
								EndPos:    3916,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 204,
									EndLine:   204,
									StartPos:  3914,
									EndPos:    3916,
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
						StartLine: 205,
						EndLine:   205,
						StartPos:  3918,
						EndPos:    3929,
					},
				},
				Expr: &ast.ExprRequire{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  3918,
							EndPos:    3928,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  3926,
								EndPos:    3928,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 205,
									EndLine:   205,
									StartPos:  3926,
									EndPos:    3928,
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
						StartLine: 206,
						EndLine:   206,
						StartPos:  3930,
						EndPos:    3946,
					},
				},
				Expr: &ast.ExprRequireOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  3930,
							EndPos:    3945,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  3943,
								EndPos:    3945,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 206,
									EndLine:   206,
									StartPos:  3943,
									EndPos:    3945,
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
						StartLine: 208,
						EndLine:   208,
						StartPos:  3948,
						EndPos:    3966,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  3948,
							EndPos:    3965,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  3948,
								EndPos:    3950,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 208,
									EndLine:   208,
									StartPos:  3948,
									EndPos:    3950,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  3962,
								EndPos:    3965,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 208,
										EndLine:   208,
										StartPos:  3962,
										EndPos:    3965,
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
						StartLine: 209,
						EndLine:   209,
						StartPos:  3967,
						EndPos:    3995,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  3967,
							EndPos:    3994,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  3967,
								EndPos:    3969,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 209,
									EndLine:   209,
									StartPos:  3967,
									EndPos:    3969,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  3981,
								EndPos:    3994,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 209,
										EndLine:   209,
										StartPos:  3991,
										EndPos:    3994,
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
						StartLine: 210,
						EndLine:   210,
						StartPos:  3996,
						EndPos:    4015,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  3996,
							EndPos:    4014,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  3996,
								EndPos:    3998,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 210,
									EndLine:   210,
									StartPos:  3996,
									EndPos:    3998,
								},
							},
							Value: []byte("$a"),
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4010,
								EndPos:    4014,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 210,
										EndLine:   210,
										StartPos:  4011,
										EndPos:    4014,
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
						StartLine: 212,
						EndLine:   212,
						StartPos:  4017,
						EndPos:    4031,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 212,
							EndLine:   212,
							StartPos:  4017,
							EndPos:    4030,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4023,
									EndPos:    4025,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 212,
										EndLine:   212,
										StartPos:  4023,
										EndPos:    4025,
									},
								},
								Value: []byte("$a"),
							},
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4027,
									EndPos:    4029,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 212,
										EndLine:   212,
										StartPos:  4027,
										EndPos:    4029,
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
						StartLine: 213,
						EndLine:   213,
						StartPos:  4032,
						EndPos:    4043,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 213,
							EndLine:   213,
							StartPos:  4032,
							EndPos:    4042,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 213,
									EndLine:   213,
									StartPos:  4038,
									EndPos:    4041,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 213,
										EndLine:   213,
										StartPos:  4038,
										EndPos:    4041,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 213,
												EndLine:   213,
												StartPos:  4038,
												EndPos:    4041,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4044,
						EndPos:    4056,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4044,
							EndPos:    4055,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4044,
								EndPos:    4050,
							},
						},
						Items: []ast.Vertex{},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4053,
								EndPos:    4055,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 214,
									EndLine:   214,
									StartPos:  4053,
									EndPos:    4055,
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
						StartLine: 215,
						EndLine:   215,
						StartPos:  4057,
						EndPos:    4075,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4057,
							EndPos:    4074,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4057,
								EndPos:    4069,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 215,
										EndLine:   215,
										StartPos:  4062,
										EndPos:    4064,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4062,
											EndPos:    4064,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 215,
												EndLine:   215,
												StartPos:  4062,
												EndPos:    4064,
											},
										},
										Value: []byte("$a"),
									},
								},
							},
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 215,
										EndLine:   215,
										StartPos:  4066,
										EndPos:    4068,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4066,
											EndPos:    4068,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 215,
												EndLine:   215,
												StartPos:  4066,
												EndPos:    4068,
											},
										},
										Value: []byte("$b"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4072,
								EndPos:    4074,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 215,
									EndLine:   215,
									StartPos:  4072,
									EndPos:    4074,
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
						StartLine: 216,
						EndLine:   216,
						StartPos:  4076,
						EndPos:    4092,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4076,
							EndPos:    4091,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4076,
								EndPos:    4086,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 216,
										EndLine:   216,
										StartPos:  4081,
										EndPos:    4085,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 216,
											EndLine:   216,
											StartPos:  4081,
											EndPos:    4085,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 216,
												EndLine:   216,
												StartPos:  4081,
												EndPos:    4083,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 216,
													EndLine:   216,
													StartPos:  4081,
													EndPos:    4083,
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
								StartLine: 216,
								EndLine:   216,
								StartPos:  4089,
								EndPos:    4091,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 216,
									EndLine:   216,
									StartPos:  4089,
									EndPos:    4091,
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
						StartLine: 217,
						EndLine:   217,
						StartPos:  4093,
						EndPos:    4113,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 217,
							EndLine:   217,
							StartPos:  4093,
							EndPos:    4112,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 217,
								EndLine:   217,
								StartPos:  4093,
								EndPos:    4107,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 217,
										EndLine:   217,
										StartPos:  4098,
										EndPos:    4106,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 217,
											EndLine:   217,
											StartPos:  4098,
											EndPos:    4106,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 217,
													EndLine:   217,
													StartPos:  4103,
													EndPos:    4105,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 217,
														EndLine:   217,
														StartPos:  4103,
														EndPos:    4105,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 217,
															EndLine:   217,
															StartPos:  4103,
															EndPos:    4105,
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
								StartLine: 217,
								EndLine:   217,
								StartPos:  4110,
								EndPos:    4112,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 217,
									EndLine:   217,
									StartPos:  4110,
									EndPos:    4112,
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
						StartLine: 219,
						EndLine:   219,
						StartPos:  4115,
						EndPos:    4125,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4115,
							EndPos:    4124,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4115,
								EndPos:    4117,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4115,
									EndPos:    4117,
								},
							},
							Value: []byte("$a"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4119,
								EndPos:    4122,
							},
						},
						Value: []byte("foo"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4122,
								EndPos:    4124,
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
						StartPos:  4126,
						EndPos:    4134,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4126,
							EndPos:    4133,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4130,
								EndPos:    4133,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 220,
										EndLine:   220,
										StartPos:  4130,
										EndPos:    4133,
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
						StartLine: 221,
						EndLine:   221,
						StartPos:  4135,
						EndPos:    4155,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4135,
							EndPos:    4154,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 221,
								EndLine:   221,
								StartPos:  4139,
								EndPos:    4152,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 221,
										EndLine:   221,
										StartPos:  4149,
										EndPos:    4152,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 221,
								EndLine:   221,
								StartPos:  4152,
								EndPos:    4154,
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
						StartPos:  4156,
						EndPos:    4167,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4156,
							EndPos:    4166,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 222,
								EndLine:   222,
								StartPos:  4160,
								EndPos:    4164,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 222,
										EndLine:   222,
										StartPos:  4161,
										EndPos:    4164,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 222,
								EndLine:   222,
								StartPos:  4164,
								EndPos:    4166,
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
						StartPos:  4168,
						EndPos:    4178,
					},
				},
				Expr: &ast.ExprPrint{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4168,
							EndPos:    4177,
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
							Value: []byte("$a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4179,
						EndPos:    4187,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4179,
							EndPos:    4186,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4179,
								EndPos:    4181,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4179,
									EndPos:    4181,
								},
							},
							Value: []byte("$a"),
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4183,
								EndPos:    4186,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4188,
						EndPos:    4199,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4188,
							EndPos:    4197,
						},
					},
					Var: &ast.ExprPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4188,
								EndPos:    4195,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4188,
									EndPos:    4190,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 225,
										EndLine:   225,
										StartPos:  4188,
										EndPos:    4190,
									},
								},
								Value: []byte("$a"),
							},
						},
						Property: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4192,
									EndPos:    4195,
								},
							},
							Value: []byte("foo"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4196,
								EndPos:    4197,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 226,
						EndLine:   226,
						StartPos:  4200,
						EndPos:    4229,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4200,
							EndPos:    4227,
						},
					},
					Var: &ast.ExprPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4200,
								EndPos:    4225,
							},
						},
						Var: &ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4200,
									EndPos:    4219,
								},
							},
							Var: &ast.ExprPropertyFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4200,
										EndPos:    4212,
									},
								},
								Var: &ast.ExprPropertyFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 226,
											EndLine:   226,
											StartPos:  4200,
											EndPos:    4207,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 226,
												EndLine:   226,
												StartPos:  4200,
												EndPos:    4202,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 226,
													EndLine:   226,
													StartPos:  4200,
													EndPos:    4202,
												},
											},
											Value: []byte("$a"),
										},
									},
									Property: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 226,
												EndLine:   226,
												StartPos:  4204,
												EndPos:    4207,
											},
										},
										Value: []byte("foo"),
									},
								},
								Property: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 226,
											EndLine:   226,
											StartPos:  4209,
											EndPos:    4212,
										},
									},
									Value: []byte("bar"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4214,
										EndPos:    4217,
									},
								},
								Value: []byte("baz"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4217,
										EndPos:    4219,
									},
								},
							},
						},
						Property: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4221,
									EndPos:    4225,
								},
							},
							Value: []byte("quux"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4226,
								EndPos:    4227,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 227,
						EndLine:   227,
						StartPos:  4230,
						EndPos:    4246,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4230,
							EndPos:    4244,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4230,
								EndPos:    4241,
							},
						},
						Var: &ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4230,
									EndPos:    4239,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 227,
										EndLine:   227,
										StartPos:  4230,
										EndPos:    4232,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 227,
											EndLine:   227,
											StartPos:  4230,
											EndPos:    4232,
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
										StartPos:  4234,
										EndPos:    4237,
									},
								},
								Value: []byte("foo"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 227,
										EndLine:   227,
										StartPos:  4237,
										EndPos:    4239,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4240,
									EndPos:    4241,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4243,
								EndPos:    4244,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4247,
						EndPos:    4256,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4247,
							EndPos:    4255,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4248,
									EndPos:    4252,
								},
							},
							Value: []byte("cmd "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4252,
									EndPos:    4254,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 228,
										EndLine:   228,
										StartPos:  4252,
										EndPos:    4254,
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
						StartLine: 229,
						EndLine:   229,
						StartPos:  4257,
						EndPos:    4263,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4257,
							EndPos:    4262,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4258,
									EndPos:    4261,
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
						StartLine: 230,
						EndLine:   230,
						StartPos:  4264,
						EndPos:    4267,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4264,
							EndPos:    4266,
						},
					},
					Parts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4268,
						EndPos:    4271,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4268,
							EndPos:    4270,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4272,
						EndPos:    4276,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 232,
							EndLine:   232,
							StartPos:  4272,
							EndPos:    4275,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 232,
									EndLine:   232,
									StartPos:  4273,
									EndPos:    4274,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 232,
										EndLine:   232,
										StartPos:  4273,
										EndPos:    4274,
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
						StartLine: 233,
						EndLine:   233,
						StartPos:  4277,
						EndPos:    4290,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4277,
							EndPos:    4289,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4278,
									EndPos:    4282,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4278,
										EndPos:    4279,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4281,
										EndPos:    4282,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4284,
									EndPos:    4287,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4284,
										EndPos:    4287,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 233,
											EndLine:   233,
											StartPos:  4285,
											EndPos:    4287,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 233,
												EndLine:   233,
												StartPos:  4285,
												EndPos:    4287,
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
						StartLine: 235,
						EndLine:   235,
						StartPos:  4292,
						EndPos:    4303,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4292,
							EndPos:    4302,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4292,
								EndPos:    4295,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 235,
										EndLine:   235,
										StartPos:  4292,
										EndPos:    4295,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4297,
								EndPos:    4300,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4300,
								EndPos:    4302,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 236,
						EndLine:   236,
						StartPos:  4304,
						EndPos:    4325,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4304,
							EndPos:    4324,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 236,
								EndLine:   236,
								StartPos:  4304,
								EndPos:    4317,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 236,
										EndLine:   236,
										StartPos:  4314,
										EndPos:    4317,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 236,
								EndLine:   236,
								StartPos:  4319,
								EndPos:    4322,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 236,
								EndLine:   236,
								StartPos:  4322,
								EndPos:    4324,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4326,
						EndPos:    4338,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4326,
							EndPos:    4337,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 237,
								EndLine:   237,
								StartPos:  4326,
								EndPos:    4330,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 237,
										EndLine:   237,
										StartPos:  4327,
										EndPos:    4330,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 237,
								EndLine:   237,
								StartPos:  4332,
								EndPos:    4335,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 237,
								EndLine:   237,
								StartPos:  4335,
								EndPos:    4337,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4339,
						EndPos:    4351,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4339,
							EndPos:    4350,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4339,
								EndPos:    4342,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 238,
										EndLine:   238,
										StartPos:  4339,
										EndPos:    4342,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4344,
								EndPos:    4348,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4344,
									EndPos:    4348,
								},
							},
							Value: []byte("$bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4348,
								EndPos:    4350,
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
						StartPos:  4352,
						EndPos:    4365,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4352,
							EndPos:    4364,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4352,
								EndPos:    4356,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4352,
									EndPos:    4356,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4358,
								EndPos:    4362,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4358,
									EndPos:    4362,
								},
							},
							Value: []byte("$bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4362,
								EndPos:    4364,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 240,
						EndLine:   240,
						StartPos:  4366,
						EndPos:    4376,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4366,
							EndPos:    4375,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4366,
								EndPos:    4369,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 240,
										EndLine:   240,
										StartPos:  4366,
										EndPos:    4369,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4371,
								EndPos:    4375,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4371,
									EndPos:    4375,
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
						StartLine: 241,
						EndLine:   241,
						StartPos:  4377,
						EndPos:    4397,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4377,
							EndPos:    4396,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4377,
								EndPos:    4390,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 241,
										EndLine:   241,
										StartPos:  4387,
										EndPos:    4390,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4392,
								EndPos:    4396,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 241,
									EndLine:   241,
									StartPos:  4392,
									EndPos:    4396,
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
						StartLine: 242,
						EndLine:   242,
						StartPos:  4398,
						EndPos:    4409,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4398,
							EndPos:    4408,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4398,
								EndPos:    4402,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 242,
										EndLine:   242,
										StartPos:  4399,
										EndPos:    4402,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4404,
								EndPos:    4408,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4404,
									EndPos:    4408,
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
						StartLine: 243,
						EndLine:   243,
						StartPos:  4410,
						EndPos:    4423,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4410,
							EndPos:    4422,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4410,
								EndPos:    4412,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4410,
									EndPos:    4412,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4415,
								EndPos:    4417,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4415,
									EndPos:    4417,
								},
							},
							Value: []byte("$b"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4420,
								EndPos:    4422,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4420,
									EndPos:    4422,
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
						StartLine: 244,
						EndLine:   244,
						StartPos:  4424,
						EndPos:    4434,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4424,
							EndPos:    4433,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4424,
								EndPos:    4426,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4424,
									EndPos:    4426,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4431,
								EndPos:    4433,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4431,
									EndPos:    4433,
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
						StartLine: 245,
						EndLine:   245,
						StartPos:  4435,
						EndPos:    4458,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4435,
							EndPos:    4457,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4435,
								EndPos:    4437,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4435,
									EndPos:    4437,
								},
							},
							Value: []byte("$a"),
						},
					},
					IfTrue: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4440,
								EndPos:    4452,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4440,
									EndPos:    4442,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4440,
										EndPos:    4442,
									},
								},
								Value: []byte("$b"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4445,
									EndPos:    4447,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4445,
										EndPos:    4447,
									},
								},
								Value: []byte("$c"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4450,
									EndPos:    4452,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4450,
										EndPos:    4452,
									},
								},
								Value: []byte("$d"),
							},
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4455,
								EndPos:    4457,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4455,
									EndPos:    4457,
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
						StartLine: 246,
						EndLine:   246,
						StartPos:  4459,
						EndPos:    4482,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4459,
							EndPos:    4481,
						},
					},
					Condition: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4459,
								EndPos:    4471,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4459,
									EndPos:    4461,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4459,
										EndPos:    4461,
									},
								},
								Value: []byte("$a"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4464,
									EndPos:    4466,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4464,
										EndPos:    4466,
									},
								},
								Value: []byte("$b"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4469,
									EndPos:    4471,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4469,
										EndPos:    4471,
									},
								},
								Value: []byte("$c"),
							},
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4474,
								EndPos:    4476,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4474,
									EndPos:    4476,
								},
							},
							Value: []byte("$d"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4479,
								EndPos:    4481,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4479,
									EndPos:    4481,
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
						StartLine: 247,
						EndLine:   247,
						StartPos:  4483,
						EndPos:    4487,
					},
				},
				Expr: &ast.ExprUnaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4483,
							EndPos:    4486,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4484,
								EndPos:    4486,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4484,
									EndPos:    4486,
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
						StartLine: 248,
						EndLine:   248,
						StartPos:  4488,
						EndPos:    4492,
					},
				},
				Expr: &ast.ExprUnaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4488,
							EndPos:    4491,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4489,
								EndPos:    4491,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4489,
									EndPos:    4491,
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
						StartLine: 249,
						EndLine:   249,
						StartPos:  4493,
						EndPos:    4497,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4493,
							EndPos:    4496,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4494,
								EndPos:    4496,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 249,
									EndLine:   249,
									StartPos:  4494,
									EndPos:    4496,
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
						StartLine: 250,
						EndLine:   250,
						StartPos:  4498,
						EndPos:    4503,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4498,
							EndPos:    4502,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4499,
								EndPos:    4502,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4500,
									EndPos:    4502,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 250,
										EndLine:   250,
										StartPos:  4500,
										EndPos:    4502,
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
						StartLine: 251,
						EndLine:   251,
						StartPos:  4504,
						EndPos:    4510,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4504,
							EndPos:    4509,
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4511,
						EndPos:    4520,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4511,
							EndPos:    4519,
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4517,
								EndPos:    4519,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4517,
									EndPos:    4519,
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
						StartLine: 253,
						EndLine:   253,
						StartPos:  4521,
						EndPos:    4536,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4521,
							EndPos:    4535,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4527,
								EndPos:    4529,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4527,
									EndPos:    4529,
								},
							},
							Value: []byte("$a"),
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4533,
								EndPos:    4535,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4533,
									EndPos:    4535,
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
						StartLine: 254,
						EndLine:   254,
						StartPos:  4537,
						EndPos:    4554,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4537,
							EndPos:    4553,
						},
					},
					Value: &ast.ExprClassConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4543,
								EndPos:    4553,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4543,
									EndPos:    4546,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 254,
											EndLine:   254,
											StartPos:  4543,
											EndPos:    4546,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4548,
									EndPos:    4553,
								},
							},
							Value: []byte("class"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  4555,
						EndPos:    4578,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  4555,
							EndPos:    4577,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4561,
								EndPos:    4563,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  4561,
									EndPos:    4563,
								},
							},
							Value: []byte("$a"),
						},
					},
					Value: &ast.ExprClassConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4567,
								EndPos:    4577,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  4567,
									EndPos:    4570,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 255,
											EndLine:   255,
											StartPos:  4567,
											EndPos:    4570,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  4572,
									EndPos:    4577,
								},
							},
							Value: []byte("class"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  4580,
						EndPos:    4590,
					},
				},
				Expr: &ast.ExprCastArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  4580,
							EndPos:    4589,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  4587,
								EndPos:    4589,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 257,
									EndLine:   257,
									StartPos:  4587,
									EndPos:    4589,
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
						StartPos:  4591,
						EndPos:    4603,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 258,
							EndLine:   258,
							StartPos:  4591,
							EndPos:    4602,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 258,
								EndLine:   258,
								StartPos:  4600,
								EndPos:    4602,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 258,
									EndLine:   258,
									StartPos:  4600,
									EndPos:    4602,
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
						StartLine: 259,
						EndLine:   259,
						StartPos:  4604,
						EndPos:    4613,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  4604,
							EndPos:    4612,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  4610,
								EndPos:    4612,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 259,
									EndLine:   259,
									StartPos:  4610,
									EndPos:    4612,
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
						StartPos:  4614,
						EndPos:    4625,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  4614,
							EndPos:    4624,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  4622,
								EndPos:    4624,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 260,
									EndLine:   260,
									StartPos:  4622,
									EndPos:    4624,
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
						StartLine: 261,
						EndLine:   261,
						StartPos:  4626,
						EndPos:    4636,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  4626,
							EndPos:    4635,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  4633,
								EndPos:    4635,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 261,
									EndLine:   261,
									StartPos:  4633,
									EndPos:    4635,
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
						StartLine: 262,
						EndLine:   262,
						StartPos:  4637,
						EndPos:    4649,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 262,
							EndLine:   262,
							StartPos:  4637,
							EndPos:    4648,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 262,
								EndLine:   262,
								StartPos:  4646,
								EndPos:    4648,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 262,
									EndLine:   262,
									StartPos:  4646,
									EndPos:    4648,
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
						StartPos:  4650,
						EndPos:    4658,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  4650,
							EndPos:    4657,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  4655,
								EndPos:    4657,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 263,
									EndLine:   263,
									StartPos:  4655,
									EndPos:    4657,
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
						StartPos:  4659,
						EndPos:    4670,
					},
				},
				Expr: &ast.ExprCastObject{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  4659,
							EndPos:    4669,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  4667,
								EndPos:    4669,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 264,
									EndLine:   264,
									StartPos:  4667,
									EndPos:    4669,
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
						StartPos:  4671,
						EndPos:    4682,
					},
				},
				Expr: &ast.ExprCastString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  4671,
							EndPos:    4681,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  4679,
								EndPos:    4681,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 265,
									EndLine:   265,
									StartPos:  4679,
									EndPos:    4681,
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
						StartPos:  4683,
						EndPos:    4693,
					},
				},
				Expr: &ast.ExprCastUnset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  4683,
							EndPos:    4692,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  4690,
								EndPos:    4692,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 266,
									EndLine:   266,
									StartPos:  4690,
									EndPos:    4692,
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
						StartPos:  4695,
						EndPos:    4703,
					},
				},
				Expr: &ast.ExprBinaryBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  4695,
							EndPos:    4702,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  4695,
								EndPos:    4697,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 268,
									EndLine:   268,
									StartPos:  4695,
									EndPos:    4697,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  4700,
								EndPos:    4702,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 268,
									EndLine:   268,
									StartPos:  4700,
									EndPos:    4702,
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
						StartLine: 269,
						EndLine:   269,
						StartPos:  4704,
						EndPos:    4712,
					},
				},
				Expr: &ast.ExprBinaryBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  4704,
							EndPos:    4711,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  4704,
								EndPos:    4706,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 269,
									EndLine:   269,
									StartPos:  4704,
									EndPos:    4706,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  4709,
								EndPos:    4711,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 269,
									EndLine:   269,
									StartPos:  4709,
									EndPos:    4711,
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
						StartLine: 270,
						EndLine:   270,
						StartPos:  4713,
						EndPos:    4721,
					},
				},
				Expr: &ast.ExprBinaryBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  4713,
							EndPos:    4720,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  4713,
								EndPos:    4715,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 270,
									EndLine:   270,
									StartPos:  4713,
									EndPos:    4715,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  4718,
								EndPos:    4720,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 270,
									EndLine:   270,
									StartPos:  4718,
									EndPos:    4720,
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
						StartLine: 271,
						EndLine:   271,
						StartPos:  4722,
						EndPos:    4731,
					},
				},
				Expr: &ast.ExprBinaryBooleanAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  4722,
							EndPos:    4730,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  4722,
								EndPos:    4724,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 271,
									EndLine:   271,
									StartPos:  4722,
									EndPos:    4724,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  4728,
								EndPos:    4730,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 271,
									EndLine:   271,
									StartPos:  4728,
									EndPos:    4730,
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
						StartLine: 272,
						EndLine:   272,
						StartPos:  4732,
						EndPos:    4741,
					},
				},
				Expr: &ast.ExprBinaryBooleanOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  4732,
							EndPos:    4740,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  4732,
								EndPos:    4734,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 272,
									EndLine:   272,
									StartPos:  4732,
									EndPos:    4734,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  4738,
								EndPos:    4740,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 272,
									EndLine:   272,
									StartPos:  4738,
									EndPos:    4740,
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
						StartLine: 273,
						EndLine:   273,
						StartPos:  4742,
						EndPos:    4750,
					},
				},
				Expr: &ast.ExprBinaryConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  4742,
							EndPos:    4749,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  4742,
								EndPos:    4744,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 273,
									EndLine:   273,
									StartPos:  4742,
									EndPos:    4744,
								},
							},
							Value: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  4747,
								EndPos:    4749,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 273,
									EndLine:   273,
									StartPos:  4747,
									EndPos:    4749,
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
						StartLine: 274,
						EndLine:   274,
						StartPos:  4751,
						EndPos:    4759,
					},
				},
				Expr: &ast.ExprBinaryDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  4751,
							EndPos:    4758,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  4751,
								EndPos:    4753,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  4751,
									EndPos:    4753,
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
								StartPos:  4756,
								EndPos:    4758,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  4756,
									EndPos:    4758,
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
						StartPos:  4760,
						EndPos:    4769,
					},
				},
				Expr: &ast.ExprBinaryEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  4760,
							EndPos:    4768,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  4760,
								EndPos:    4762,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  4760,
									EndPos:    4762,
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
								StartPos:  4766,
								EndPos:    4768,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  4766,
									EndPos:    4768,
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
						StartPos:  4770,
						EndPos:    4779,
					},
				},
				Expr: &ast.ExprBinaryGreaterOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  4770,
							EndPos:    4778,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  4770,
								EndPos:    4772,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
									StartPos:  4770,
									EndPos:    4772,
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
								StartPos:  4776,
								EndPos:    4778,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
									StartPos:  4776,
									EndPos:    4778,
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
						StartPos:  4780,
						EndPos:    4788,
					},
				},
				Expr: &ast.ExprBinaryGreater{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  4780,
							EndPos:    4787,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  4780,
								EndPos:    4782,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  4780,
									EndPos:    4782,
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
								StartPos:  4785,
								EndPos:    4787,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  4785,
									EndPos:    4787,
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
						StartPos:  4789,
						EndPos:    4799,
					},
				},
				Expr: &ast.ExprBinaryIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  4789,
							EndPos:    4798,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  4789,
								EndPos:    4791,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  4789,
									EndPos:    4791,
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
						StartPos:  4800,
						EndPos:    4810,
					},
				},
				Expr: &ast.ExprBinaryLogicalAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  4800,
							EndPos:    4809,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  4800,
								EndPos:    4802,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  4800,
									EndPos:    4802,
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
								StartPos:  4807,
								EndPos:    4809,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  4807,
									EndPos:    4809,
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
						StartPos:  4811,
						EndPos:    4820,
					},
				},
				Expr: &ast.ExprBinaryLogicalOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  4811,
							EndPos:    4819,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  4811,
								EndPos:    4813,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  4811,
									EndPos:    4813,
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
								StartPos:  4817,
								EndPos:    4819,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  4817,
									EndPos:    4819,
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
						StartPos:  4821,
						EndPos:    4831,
					},
				},
				Expr: &ast.ExprBinaryLogicalXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  4821,
							EndPos:    4830,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  4821,
								EndPos:    4823,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  4821,
									EndPos:    4823,
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
								StartPos:  4828,
								EndPos:    4830,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  4828,
									EndPos:    4830,
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
						StartPos:  4832,
						EndPos:    4840,
					},
				},
				Expr: &ast.ExprBinaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  4832,
							EndPos:    4839,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  4832,
								EndPos:    4834,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  4832,
									EndPos:    4834,
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
								StartPos:  4837,
								EndPos:    4839,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  4837,
									EndPos:    4839,
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
						StartPos:  4841,
						EndPos:    4849,
					},
				},
				Expr: &ast.ExprBinaryMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  4841,
							EndPos:    4848,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  4841,
								EndPos:    4843,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  4841,
									EndPos:    4843,
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
								StartPos:  4846,
								EndPos:    4848,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  4846,
									EndPos:    4848,
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
						StartPos:  4850,
						EndPos:    4858,
					},
				},
				Expr: &ast.ExprBinaryMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  4850,
							EndPos:    4857,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  4850,
								EndPos:    4852,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  4850,
									EndPos:    4852,
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
								StartPos:  4855,
								EndPos:    4857,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  4855,
									EndPos:    4857,
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
						StartPos:  4859,
						EndPos:    4868,
					},
				},
				Expr: &ast.ExprBinaryNotEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  4859,
							EndPos:    4867,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  4859,
								EndPos:    4861,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  4859,
									EndPos:    4861,
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
								StartPos:  4865,
								EndPos:    4867,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  4865,
									EndPos:    4867,
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
						StartPos:  4869,
						EndPos:    4879,
					},
				},
				Expr: &ast.ExprBinaryNotIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  4869,
							EndPos:    4878,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  4869,
								EndPos:    4871,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  4869,
									EndPos:    4871,
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
								StartPos:  4876,
								EndPos:    4878,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  4876,
									EndPos:    4878,
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
						StartPos:  4880,
						EndPos:    4888,
					},
				},
				Expr: &ast.ExprBinaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  4880,
							EndPos:    4887,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  4880,
								EndPos:    4882,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  4880,
									EndPos:    4882,
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
						StartPos:  4889,
						EndPos:    4898,
					},
				},
				Expr: &ast.ExprBinaryPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  4889,
							EndPos:    4897,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  4889,
								EndPos:    4891,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
									StartPos:  4889,
									EndPos:    4891,
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
						StartPos:  4899,
						EndPos:    4908,
					},
				},
				Expr: &ast.ExprBinaryShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  4899,
							EndPos:    4907,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  4899,
								EndPos:    4901,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  4899,
									EndPos:    4901,
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
								StartPos:  4905,
								EndPos:    4907,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  4905,
									EndPos:    4907,
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
						StartPos:  4909,
						EndPos:    4918,
					},
				},
				Expr: &ast.ExprBinaryShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  4909,
							EndPos:    4917,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  4909,
								EndPos:    4911,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  4909,
									EndPos:    4911,
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
						StartPos:  4919,
						EndPos:    4928,
					},
				},
				Expr: &ast.ExprBinarySmallerOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  4919,
							EndPos:    4927,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  4919,
								EndPos:    4921,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  4919,
									EndPos:    4921,
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
								StartPos:  4925,
								EndPos:    4927,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  4925,
									EndPos:    4927,
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
						StartPos:  4929,
						EndPos:    4937,
					},
				},
				Expr: &ast.ExprBinarySmaller{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  4929,
							EndPos:    4936,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  4929,
								EndPos:    4931,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  4929,
									EndPos:    4931,
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
								StartPos:  4934,
								EndPos:    4936,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  4934,
									EndPos:    4936,
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
						StartPos:  4939,
						EndPos:    4948,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  4939,
							EndPos:    4947,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  4939,
								EndPos:    4941,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  4939,
									EndPos:    4941,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  4945,
								EndPos:    4947,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  4945,
									EndPos:    4947,
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
						StartPos:  4949,
						EndPos:    4963,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  4949,
							EndPos:    4962,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  4949,
								EndPos:    4951,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  4949,
									EndPos:    4951,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  4955,
								EndPos:    4962,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  4959,
									EndPos:    4962,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 295,
											EndLine:   295,
											StartPos:  4959,
											EndPos:    4962,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  4964,
						EndPos:    4982,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  4964,
							EndPos:    4981,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  4964,
								EndPos:    4966,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  4964,
									EndPos:    4966,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  4970,
								EndPos:    4981,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  4974,
									EndPos:    4977,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 296,
											EndLine:   296,
											StartPos:  4974,
											EndPos:    4977,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  4977,
									EndPos:    4981,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 296,
											EndLine:   296,
											StartPos:  4978,
											EndPos:    4980,
										},
									},
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 296,
												EndLine:   296,
												StartPos:  4978,
												EndPos:    4980,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 296,
													EndLine:   296,
													StartPos:  4978,
													EndPos:    4980,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  4983,
						EndPos:    4991,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  4983,
							EndPos:    4990,
						},
					},
					Var: &ast.ExprVariable{
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
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  4988,
								EndPos:    4990,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  4988,
									EndPos:    4990,
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
						StartPos:  4992,
						EndPos:    5001,
					},
				},
				Expr: &ast.ExprAssignBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  4992,
							EndPos:    5000,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  4992,
								EndPos:    4994,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  4992,
									EndPos:    4994,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  4998,
								EndPos:    5000,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  4998,
									EndPos:    5000,
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
						StartPos:  5002,
						EndPos:    5011,
					},
				},
				Expr: &ast.ExprAssignBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5002,
							EndPos:    5010,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5002,
								EndPos:    5004,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5002,
									EndPos:    5004,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
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
						EndPos:    5021,
					},
				},
				Expr: &ast.ExprAssignBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5012,
							EndPos:    5020,
						},
					},
					Var: &ast.ExprVariable{
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
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5018,
								EndPos:    5020,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5018,
									EndPos:    5020,
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
						StartLine: 301,
						EndLine:   301,
						StartPos:  5022,
						EndPos:    5031,
					},
				},
				Expr: &ast.ExprAssignConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5022,
							EndPos:    5030,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5022,
								EndPos:    5024,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 301,
									EndLine:   301,
									StartPos:  5022,
									EndPos:    5024,
								},
							},
							Value: []byte("$a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5028,
								EndPos:    5030,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 301,
									EndLine:   301,
									StartPos:  5028,
									EndPos:    5030,
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
						StartPos:  5032,
						EndPos:    5041,
					},
				},
				Expr: &ast.ExprAssignDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5032,
							EndPos:    5040,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5032,
								EndPos:    5034,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5032,
									EndPos:    5034,
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
								StartPos:  5038,
								EndPos:    5040,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5038,
									EndPos:    5040,
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
						StartPos:  5042,
						EndPos:    5051,
					},
				},
				Expr: &ast.ExprAssignMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5042,
							EndPos:    5050,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5042,
								EndPos:    5044,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5042,
									EndPos:    5044,
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
								StartPos:  5048,
								EndPos:    5050,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5048,
									EndPos:    5050,
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
						StartPos:  5052,
						EndPos:    5061,
					},
				},
				Expr: &ast.ExprAssignMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5052,
							EndPos:    5060,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5052,
								EndPos:    5054,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5052,
									EndPos:    5054,
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
								StartPos:  5058,
								EndPos:    5060,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5058,
									EndPos:    5060,
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
						StartPos:  5062,
						EndPos:    5071,
					},
				},
				Expr: &ast.ExprAssignMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5062,
							EndPos:    5070,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5062,
								EndPos:    5064,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5062,
									EndPos:    5064,
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
								StartPos:  5068,
								EndPos:    5070,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5068,
									EndPos:    5070,
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
						StartPos:  5072,
						EndPos:    5081,
					},
				},
				Expr: &ast.ExprAssignPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5072,
							EndPos:    5080,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5072,
								EndPos:    5074,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5072,
									EndPos:    5074,
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
								StartPos:  5078,
								EndPos:    5080,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5078,
									EndPos:    5080,
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
						StartPos:  5082,
						EndPos:    5092,
					},
				},
				Expr: &ast.ExprAssignPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5082,
							EndPos:    5091,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5082,
								EndPos:    5084,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5082,
									EndPos:    5084,
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
								StartPos:  5089,
								EndPos:    5091,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
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
						StartLine: 308,
						EndLine:   308,
						StartPos:  5093,
						EndPos:    5103,
					},
				},
				Expr: &ast.ExprAssignShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5093,
							EndPos:    5102,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5093,
								EndPos:    5095,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
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
								StartLine: 308,
								EndLine:   308,
								StartPos:  5100,
								EndPos:    5102,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5100,
									EndPos:    5102,
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
						StartPos:  5104,
						EndPos:    5114,
					},
				},
				Expr: &ast.ExprAssignShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5104,
							EndPos:    5113,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5104,
								EndPos:    5106,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5104,
									EndPos:    5106,
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
								StartPos:  5111,
								EndPos:    5113,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5111,
									EndPos:    5113,
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
						StartPos:  5118,
						EndPos:    5130,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5118,
							EndPos:    5128,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5122,
								EndPos:    5126,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 312,
										EndLine:   312,
										StartPos:  5123,
										EndPos:    5126,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5126,
								EndPos:    5128,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 313,
						EndLine:   313,
						StartPos:  5152,
						EndPos:    5156,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5152,
							EndPos:    5155,
						},
					},
					Var: &ast.ExprMethodCall{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5148,
								EndPos:    5150,
							},
						},
						Var: &ast.ExprNew{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5132,
									EndPos:    5142,
								},
							},
							Class: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 313,
										EndLine:   313,
										StartPos:  5136,
										EndPos:    5140,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 313,
												EndLine:   313,
												StartPos:  5137,
												EndPos:    5140,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 313,
										EndLine:   313,
										StartPos:  5140,
										EndPos:    5142,
									},
								},
							},
						},
						Method: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5145,
									EndPos:    5148,
								},
							},
							Value: []byte("bar"),
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5148,
									EndPos:    5150,
								},
							},
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5152,
								EndPos:    5155,
							},
						},
						Value: []byte("baz"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5173,
						EndPos:    5176,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5173,
							EndPos:    5174,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5170,
								EndPos:    5171,
							},
						},
						Var: &ast.ExprNew{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5158,
									EndPos:    5168,
								},
							},
							Class: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 314,
										EndLine:   314,
										StartPos:  5162,
										EndPos:    5166,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 314,
												EndLine:   314,
												StartPos:  5163,
												EndPos:    5166,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 314,
										EndLine:   314,
										StartPos:  5166,
										EndPos:    5168,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5170,
									EndPos:    5171,
								},
							},
							Value: []byte("0"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5173,
								EndPos:    5174,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5197,
						EndPos:    5200,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5197,
							EndPos:    5199,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5190,
								EndPos:    5191,
							},
						},
						Var: &ast.ExprNew{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5178,
									EndPos:    5188,
								},
							},
							Class: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 315,
										EndLine:   315,
										StartPos:  5182,
										EndPos:    5186,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 315,
												EndLine:   315,
												StartPos:  5183,
												EndPos:    5186,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 315,
										EndLine:   315,
										StartPos:  5186,
										EndPos:    5188,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5190,
									EndPos:    5191,
								},
							},
							Value: []byte("0"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5194,
								EndPos:    5197,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5197,
								EndPos:    5199,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5202,
						EndPos:    5219,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5202,
							EndPos:    5218,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5202,
								EndPos:    5215,
							},
						},
						Var: &ast.ExprArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5202,
									EndPos:    5212,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 317,
											EndLine:   317,
											StartPos:  5208,
											EndPos:    5211,
										},
									},
									Val: &ast.ExprShortArray{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 317,
												EndLine:   317,
												StartPos:  5208,
												EndPos:    5211,
											},
										},
										Items: []ast.Vertex{
											&ast.ExprArrayItem{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 317,
														EndLine:   317,
														StartPos:  5209,
														EndPos:    5210,
													},
												},
												Val: &ast.ScalarLnumber{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 317,
															EndLine:   317,
															StartPos:  5209,
															EndPos:    5210,
														},
													},
													Value: []byte("0"),
												},
											},
										},
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5213,
									EndPos:    5214,
								},
							},
							Value: []byte("0"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5216,
								EndPos:    5217,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 318,
						EndLine:   318,
						StartPos:  5220,
						EndPos:    5229,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5220,
							EndPos:    5228,
						},
					},
					Var: &ast.ScalarString{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5220,
								EndPos:    5225,
							},
						},
						Value: []byte("\"foo\""),
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5226,
								EndPos:    5227,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 319,
						EndLine:   319,
						StartPos:  5230,
						EndPos:    5237,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5230,
							EndPos:    5236,
						},
					},
					Var: &ast.ExprConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 319,
								EndLine:   319,
								StartPos:  5230,
								EndPos:    5233,
							},
						},
						Const: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 319,
									EndLine:   319,
									StartPos:  5230,
									EndPos:    5233,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 319,
											EndLine:   319,
											StartPos:  5230,
											EndPos:    5233,
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
								StartLine: 319,
								EndLine:   319,
								StartPos:  5234,
								EndPos:    5235,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   320,
						StartPos:  5238,
						EndPos:    5250,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5238,
							EndPos:    5249,
						},
					},
					Class: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5238,
								EndPos:    5244,
							},
						},
						Value: []byte("static"),
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5246,
								EndPos:    5249,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 322,
						EndLine:   322,
						StartPos:  5252,
						EndPos:    5261,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5252,
							EndPos:    5260,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 322,
								EndLine:   322,
								StartPos:  5256,
								EndPos:    5260,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 322,
									EndLine:   322,
									StartPos:  5256,
									EndPos:    5260,
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
						StartLine: 323,
						EndLine:   323,
						StartPos:  5262,
						EndPos:    5277,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 323,
							EndLine:   323,
							StartPos:  5262,
							EndPos:    5276,
						},
					},
					Class: &ast.ExprStaticPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5266,
								EndPos:    5276,
							},
						},
						Class: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5266,
									EndPos:    5270,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 323,
										EndLine:   323,
										StartPos:  5266,
										EndPos:    5270,
									},
								},
								Value: []byte("$foo"),
							},
						},
						Property: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5272,
									EndPos:    5276,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 323,
										EndLine:   323,
										StartPos:  5272,
										EndPos:    5276,
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
						StartLine: 324,
						EndLine:   324,
						StartPos:  5278,
						EndPos:    5291,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 324,
							EndLine:   324,
							StartPos:  5278,
							EndPos:    5289,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5288,
								EndPos:    5289,
							},
						},
						Var: &ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5286,
									EndPos:    5289,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 324,
										EndLine:   324,
										StartPos:  5282,
										EndPos:    5287,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 324,
											EndLine:   324,
											StartPos:  5282,
											EndPos:    5284,
										},
									},
									Value: []byte("$a"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 324,
										EndLine:   324,
										StartPos:  5286,
										EndPos:    5287,
									},
								},
								Value: []byte("b"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5288,
									EndPos:    5289,
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
						StartLine: 325,
						EndLine:   325,
						StartPos:  5292,
						EndPos:    5324,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5292,
							EndPos:    5322,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5321,
								EndPos:    5322,
							},
						},
						Var: &ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5319,
									EndPos:    5322,
								},
							},
							Var: &ast.ExprPropertyFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5315,
										EndPos:    5320,
									},
								},
								Var: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5302,
											EndPos:    5317,
										},
									},
									Var: &ast.ExprPropertyFetch{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5300,
												EndPos:    5312,
											},
										},
										Var: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5296,
													EndPos:    5301,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 325,
														EndLine:   325,
														StartPos:  5296,
														EndPos:    5298,
													},
												},
												Value: []byte("$a"),
											},
										},
										Property: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5300,
													EndPos:    5301,
												},
											},
											Value: []byte("b"),
										},
									},
									Dim: &ast.ExprTernary{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5302,
												EndPos:    5312,
											},
										},
										Condition: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5302,
													EndPos:    5304,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 325,
														EndLine:   325,
														StartPos:  5302,
														EndPos:    5304,
													},
												},
												Value: []byte("$b"),
											},
										},
										IfFalse: &ast.ExprConstFetch{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5308,
													EndPos:    5312,
												},
											},
											Const: &ast.NameName{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 325,
														EndLine:   325,
														StartPos:  5308,
														EndPos:    5312,
													},
												},
												Parts: []ast.Vertex{
													&ast.NameNamePart{
														Node: ast.Node{
															Position: &position.Position{
																StartLine: 325,
																EndLine:   325,
																StartPos:  5308,
																EndPos:    5312,
															},
														},
														Value: []byte("null"),
													},
												},
											},
										},
									},
								},
								Property: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5315,
											EndPos:    5317,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5315,
												EndPos:    5317,
											},
										},
										Value: []byte("$c"),
									},
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5319,
										EndPos:    5320,
									},
								},
								Value: []byte("d"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5321,
									EndPos:    5322,
								},
							},
							Value: []byte("0"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 325,
						EndLine:   325,
						StartPos:  5324,
						EndPos:    5343,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5331,
								EndPos:    5342,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5331,
									EndPos:    5333,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5331,
										EndPos:    5333,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5336,
									EndPos:    5342,
								},
							},
							Var: &ast.ExprShortArray{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5336,
										EndPos:    5339,
									},
								},
								Items: []ast.Vertex{
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5337,
												EndPos:    5338,
											},
										},
										Val: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5337,
													EndPos:    5338,
												},
											},
											Value: []byte("1"),
										},
									},
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5340,
										EndPos:    5341,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 327,
						EndLine:   327,
						StartPos:  5345,
						EndPos:    5360,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5352,
								EndPos:    5359,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5352,
									EndPos:    5354,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 327,
										EndLine:   327,
										StartPos:  5352,
										EndPos:    5354,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBooleanNot{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5357,
									EndPos:    5359,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 327,
										EndLine:   327,
										StartPos:  5358,
										EndPos:    5359,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 328,
						EndLine:   328,
						StartPos:  5361,
						EndPos:    5376,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5368,
								EndPos:    5375,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5368,
									EndPos:    5370,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5368,
										EndPos:    5370,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBitwiseNot{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5373,
									EndPos:    5375,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5374,
										EndPos:    5375,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 329,
						EndLine:   329,
						StartPos:  5377,
						EndPos:    5392,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5384,
								EndPos:    5391,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5384,
									EndPos:    5386,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5384,
										EndPos:    5386,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprUnaryPlus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5389,
									EndPos:    5391,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5390,
										EndPos:    5391,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 330,
						EndLine:   330,
						StartPos:  5393,
						EndPos:    5408,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5400,
								EndPos:    5407,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5400,
									EndPos:    5402,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5400,
										EndPos:    5402,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprUnaryMinus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5405,
									EndPos:    5407,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5406,
										EndPos:    5407,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 331,
						EndLine:   331,
						StartPos:  5409,
						EndPos:    5425,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  5416,
								EndPos:    5423,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5416,
									EndPos:    5418,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 331,
										EndLine:   331,
										StartPos:  5416,
										EndPos:    5418,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5422,
									EndPos:    5423,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 332,
						EndLine:   332,
						StartPos:  5426,
						EndPos:    5445,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  5433,
								EndPos:    5444,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  5433,
									EndPos:    5435,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  5433,
										EndPos:    5435,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprTernary{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  5438,
									EndPos:    5444,
								},
							},
							Condition: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  5438,
										EndPos:    5439,
									},
								},
								Value: []byte("1"),
							},
							IfFalse: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  5443,
										EndPos:    5444,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 333,
						EndLine:   333,
						StartPos:  5446,
						EndPos:    5468,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  5453,
								EndPos:    5467,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  5453,
									EndPos:    5455,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  5453,
										EndPos:    5455,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprTernary{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  5458,
									EndPos:    5467,
								},
							},
							Condition: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  5458,
										EndPos:    5459,
									},
								},
								Value: []byte("1"),
							},
							IfTrue: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  5462,
										EndPos:    5463,
									},
								},
								Value: []byte("2"),
							},
							IfFalse: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  5466,
										EndPos:    5467,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 334,
						EndLine:   334,
						StartPos:  5469,
						EndPos:    5487,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  5476,
								EndPos:    5486,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  5476,
									EndPos:    5478,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  5476,
										EndPos:    5478,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryBitwiseAnd{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  5481,
									EndPos:    5486,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  5481,
										EndPos:    5482,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  5485,
										EndPos:    5486,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 335,
						EndLine:   335,
						StartPos:  5488,
						EndPos:    5506,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  5495,
								EndPos:    5505,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  5495,
									EndPos:    5497,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  5495,
										EndPos:    5497,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryBitwiseOr{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  5500,
									EndPos:    5505,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  5500,
										EndPos:    5501,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  5504,
										EndPos:    5505,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 336,
						EndLine:   336,
						StartPos:  5507,
						EndPos:    5525,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  5514,
								EndPos:    5524,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  5514,
									EndPos:    5516,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 336,
										EndLine:   336,
										StartPos:  5514,
										EndPos:    5516,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryBitwiseXor{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  5519,
									EndPos:    5524,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 336,
										EndLine:   336,
										StartPos:  5519,
										EndPos:    5520,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 336,
										EndLine:   336,
										StartPos:  5523,
										EndPos:    5524,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 337,
						EndLine:   337,
						StartPos:  5526,
						EndPos:    5545,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  5533,
								EndPos:    5544,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  5533,
									EndPos:    5535,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  5533,
										EndPos:    5535,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryBooleanAnd{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  5538,
									EndPos:    5544,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  5538,
										EndPos:    5539,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  5543,
										EndPos:    5544,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 338,
						EndLine:   338,
						StartPos:  5546,
						EndPos:    5565,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  5553,
								EndPos:    5564,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  5553,
									EndPos:    5555,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  5553,
										EndPos:    5555,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryBooleanOr{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  5558,
									EndPos:    5564,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  5558,
										EndPos:    5559,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  5563,
										EndPos:    5564,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 339,
						EndLine:   339,
						StartPos:  5566,
						EndPos:    5584,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  5573,
								EndPos:    5583,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  5573,
									EndPos:    5575,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  5573,
										EndPos:    5575,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryConcat{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  5578,
									EndPos:    5583,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  5578,
										EndPos:    5579,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  5582,
										EndPos:    5583,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 340,
						EndLine:   340,
						StartPos:  5585,
						EndPos:    5603,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  5592,
								EndPos:    5602,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  5592,
									EndPos:    5594,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  5592,
										EndPos:    5594,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryDiv{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  5597,
									EndPos:    5602,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  5597,
										EndPos:    5598,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  5601,
										EndPos:    5602,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 341,
						EndLine:   341,
						StartPos:  5604,
						EndPos:    5623,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  5611,
								EndPos:    5622,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  5611,
									EndPos:    5613,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  5611,
										EndPos:    5613,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  5616,
									EndPos:    5622,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  5616,
										EndPos:    5617,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  5621,
										EndPos:    5622,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 342,
						EndLine:   342,
						StartPos:  5624,
						EndPos:    5643,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  5631,
								EndPos:    5642,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  5631,
									EndPos:    5633,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 342,
										EndLine:   342,
										StartPos:  5631,
										EndPos:    5633,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryGreaterOrEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  5636,
									EndPos:    5642,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 342,
										EndLine:   342,
										StartPos:  5636,
										EndPos:    5637,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 342,
										EndLine:   342,
										StartPos:  5641,
										EndPos:    5642,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 343,
						EndLine:   343,
						StartPos:  5644,
						EndPos:    5662,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  5651,
								EndPos:    5661,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  5651,
									EndPos:    5653,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 343,
										EndLine:   343,
										StartPos:  5651,
										EndPos:    5653,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryGreater{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  5656,
									EndPos:    5661,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 343,
										EndLine:   343,
										StartPos:  5656,
										EndPos:    5657,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 343,
										EndLine:   343,
										StartPos:  5660,
										EndPos:    5661,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 344,
						EndLine:   344,
						StartPos:  5663,
						EndPos:    5683,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  5670,
								EndPos:    5682,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  5670,
									EndPos:    5672,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  5670,
										EndPos:    5672,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryIdentical{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  5675,
									EndPos:    5682,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  5675,
										EndPos:    5676,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  5681,
										EndPos:    5682,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 345,
						EndLine:   345,
						StartPos:  5684,
						EndPos:    5704,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  5691,
								EndPos:    5703,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  5691,
									EndPos:    5693,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 345,
										EndLine:   345,
										StartPos:  5691,
										EndPos:    5693,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryLogicalAnd{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  5696,
									EndPos:    5703,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 345,
										EndLine:   345,
										StartPos:  5696,
										EndPos:    5697,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 345,
										EndLine:   345,
										StartPos:  5702,
										EndPos:    5703,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 346,
						EndLine:   346,
						StartPos:  5705,
						EndPos:    5724,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  5712,
								EndPos:    5723,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  5712,
									EndPos:    5714,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5712,
										EndPos:    5714,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryLogicalOr{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  5717,
									EndPos:    5723,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5717,
										EndPos:    5718,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  5722,
										EndPos:    5723,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 347,
						EndLine:   347,
						StartPos:  5725,
						EndPos:    5745,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  5732,
								EndPos:    5744,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  5732,
									EndPos:    5734,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  5732,
										EndPos:    5734,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryLogicalXor{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  5737,
									EndPos:    5744,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  5737,
										EndPos:    5738,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  5743,
										EndPos:    5744,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 348,
						EndLine:   348,
						StartPos:  5746,
						EndPos:    5764,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 348,
								EndLine:   348,
								StartPos:  5753,
								EndPos:    5763,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  5753,
									EndPos:    5755,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 348,
										EndLine:   348,
										StartPos:  5753,
										EndPos:    5755,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryMinus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  5758,
									EndPos:    5763,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 348,
										EndLine:   348,
										StartPos:  5758,
										EndPos:    5759,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 348,
										EndLine:   348,
										StartPos:  5762,
										EndPos:    5763,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 349,
						EndLine:   349,
						StartPos:  5765,
						EndPos:    5783,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 349,
								EndLine:   349,
								StartPos:  5772,
								EndPos:    5782,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  5772,
									EndPos:    5774,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 349,
										EndLine:   349,
										StartPos:  5772,
										EndPos:    5774,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryMod{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  5777,
									EndPos:    5782,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 349,
										EndLine:   349,
										StartPos:  5777,
										EndPos:    5778,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 349,
										EndLine:   349,
										StartPos:  5781,
										EndPos:    5782,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 350,
						EndLine:   350,
						StartPos:  5784,
						EndPos:    5802,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 350,
								EndLine:   350,
								StartPos:  5791,
								EndPos:    5801,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  5791,
									EndPos:    5793,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 350,
										EndLine:   350,
										StartPos:  5791,
										EndPos:    5793,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryMul{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  5796,
									EndPos:    5801,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 350,
										EndLine:   350,
										StartPos:  5796,
										EndPos:    5797,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 350,
										EndLine:   350,
										StartPos:  5800,
										EndPos:    5801,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 351,
						EndLine:   351,
						StartPos:  5803,
						EndPos:    5822,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 351,
								EndLine:   351,
								StartPos:  5810,
								EndPos:    5821,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  5810,
									EndPos:    5812,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 351,
										EndLine:   351,
										StartPos:  5810,
										EndPos:    5812,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryNotEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  5815,
									EndPos:    5821,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 351,
										EndLine:   351,
										StartPos:  5815,
										EndPos:    5816,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 351,
										EndLine:   351,
										StartPos:  5820,
										EndPos:    5821,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 352,
						EndLine:   352,
						StartPos:  5823,
						EndPos:    5843,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 352,
								EndLine:   352,
								StartPos:  5830,
								EndPos:    5842,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  5830,
									EndPos:    5832,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 352,
										EndLine:   352,
										StartPos:  5830,
										EndPos:    5832,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryNotIdentical{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  5835,
									EndPos:    5842,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 352,
										EndLine:   352,
										StartPos:  5835,
										EndPos:    5836,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 352,
										EndLine:   352,
										StartPos:  5841,
										EndPos:    5842,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 353,
						EndLine:   353,
						StartPos:  5844,
						EndPos:    5862,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 353,
								EndLine:   353,
								StartPos:  5851,
								EndPos:    5861,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  5851,
									EndPos:    5853,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 353,
										EndLine:   353,
										StartPos:  5851,
										EndPos:    5853,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryPlus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  5856,
									EndPos:    5861,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 353,
										EndLine:   353,
										StartPos:  5856,
										EndPos:    5857,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 353,
										EndLine:   353,
										StartPos:  5860,
										EndPos:    5861,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 354,
						EndLine:   354,
						StartPos:  5863,
						EndPos:    5882,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 354,
								EndLine:   354,
								StartPos:  5870,
								EndPos:    5881,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  5870,
									EndPos:    5872,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 354,
										EndLine:   354,
										StartPos:  5870,
										EndPos:    5872,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryPow{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  5875,
									EndPos:    5881,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 354,
										EndLine:   354,
										StartPos:  5875,
										EndPos:    5876,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 354,
										EndLine:   354,
										StartPos:  5880,
										EndPos:    5881,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 355,
						EndLine:   355,
						StartPos:  5883,
						EndPos:    5902,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 355,
								EndLine:   355,
								StartPos:  5890,
								EndPos:    5901,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  5890,
									EndPos:    5892,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 355,
										EndLine:   355,
										StartPos:  5890,
										EndPos:    5892,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryShiftLeft{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  5895,
									EndPos:    5901,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 355,
										EndLine:   355,
										StartPos:  5895,
										EndPos:    5896,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 355,
										EndLine:   355,
										StartPos:  5900,
										EndPos:    5901,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 356,
						EndLine:   356,
						StartPos:  5903,
						EndPos:    5922,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 356,
								EndLine:   356,
								StartPos:  5910,
								EndPos:    5921,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  5910,
									EndPos:    5912,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 356,
										EndLine:   356,
										StartPos:  5910,
										EndPos:    5912,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinaryShiftRight{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  5915,
									EndPos:    5921,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 356,
										EndLine:   356,
										StartPos:  5915,
										EndPos:    5916,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 356,
										EndLine:   356,
										StartPos:  5920,
										EndPos:    5921,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 357,
						EndLine:   357,
						StartPos:  5923,
						EndPos:    5942,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 357,
								EndLine:   357,
								StartPos:  5930,
								EndPos:    5941,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  5930,
									EndPos:    5932,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 357,
										EndLine:   357,
										StartPos:  5930,
										EndPos:    5932,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinarySmallerOrEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  5935,
									EndPos:    5941,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 357,
										EndLine:   357,
										StartPos:  5935,
										EndPos:    5936,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 357,
										EndLine:   357,
										StartPos:  5940,
										EndPos:    5941,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 358,
						EndLine:   358,
						StartPos:  5943,
						EndPos:    5961,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 358,
								EndLine:   358,
								StartPos:  5950,
								EndPos:    5960,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  5950,
									EndPos:    5952,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 358,
										EndLine:   358,
										StartPos:  5950,
										EndPos:    5952,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprBinarySmaller{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  5955,
									EndPos:    5960,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 358,
										EndLine:   358,
										StartPos:  5955,
										EndPos:    5956,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 358,
										EndLine:   358,
										StartPos:  5959,
										EndPos:    5960,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 359,
						EndLine:   359,
						StartPos:  5962,
						EndPos:    5983,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 359,
								EndLine:   359,
								StartPos:  5969,
								EndPos:    5982,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  5969,
									EndPos:    5971,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 359,
										EndLine:   359,
										StartPos:  5969,
										EndPos:    5971,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprClassConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  5974,
									EndPos:    5982,
								},
							},
							Class: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 359,
										EndLine:   359,
										StartPos:  5974,
										EndPos:    5977,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 359,
												EndLine:   359,
												StartPos:  5974,
												EndPos:    5977,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ConstantName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 359,
										EndLine:   359,
										StartPos:  5979,
										EndPos:    5982,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 360,
						EndLine:   360,
						StartPos:  5984,
						EndPos:    6007,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 360,
								EndLine:   360,
								StartPos:  5991,
								EndPos:    6006,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  5991,
									EndPos:    5993,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 360,
										EndLine:   360,
										StartPos:  5991,
										EndPos:    5993,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprClassConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  5996,
									EndPos:    6006,
								},
							},
							Class: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 360,
										EndLine:   360,
										StartPos:  5996,
										EndPos:    5999,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 360,
												EndLine:   360,
												StartPos:  5996,
												EndPos:    5999,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ConstantName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 360,
										EndLine:   360,
										StartPos:  6001,
										EndPos:    6006,
									},
								},
								Value: []byte("class"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 361,
						EndLine:   361,
						StartPos:  6008,
						EndPos:    6030,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 361,
								EndLine:   361,
								StartPos:  6015,
								EndPos:    6029,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6015,
									EndPos:    6017,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 361,
										EndLine:   361,
										StartPos:  6015,
										EndPos:    6017,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ScalarMagicConstant{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6020,
									EndPos:    6029,
								},
							},
							Value: []byte("__CLASS__"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 362,
						EndLine:   362,
						StartPos:  6031,
						EndPos:    6047,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 362,
								EndLine:   362,
								StartPos:  6038,
								EndPos:    6046,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6038,
									EndPos:    6040,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 362,
										EndLine:   362,
										StartPos:  6038,
										EndPos:    6040,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6043,
									EndPos:    6046,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 362,
										EndLine:   362,
										StartPos:  6043,
										EndPos:    6046,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 362,
												EndLine:   362,
												StartPos:  6043,
												EndPos:    6046,
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
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 363,
						EndLine:   363,
						StartPos:  6048,
						EndPos:    6074,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 363,
								EndLine:   363,
								StartPos:  6055,
								EndPos:    6073,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6055,
									EndPos:    6057,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 363,
										EndLine:   363,
										StartPos:  6055,
										EndPos:    6057,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6060,
									EndPos:    6073,
								},
							},
							Const: &ast.NameRelative{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 363,
										EndLine:   363,
										StartPos:  6060,
										EndPos:    6073,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 363,
												EndLine:   363,
												StartPos:  6070,
												EndPos:    6073,
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
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 364,
						EndLine:   364,
						StartPos:  6075,
						EndPos:    6092,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 364,
								EndLine:   364,
								StartPos:  6082,
								EndPos:    6091,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6082,
									EndPos:    6084,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 364,
										EndLine:   364,
										StartPos:  6082,
										EndPos:    6084,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6087,
									EndPos:    6091,
								},
							},
							Const: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 364,
										EndLine:   364,
										StartPos:  6087,
										EndPos:    6091,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 364,
												EndLine:   364,
												StartPos:  6088,
												EndPos:    6091,
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
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 365,
						EndLine:   365,
						StartPos:  6093,
						EndPos:    6113,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 365,
								EndLine:   365,
								StartPos:  6100,
								EndPos:    6112,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6100,
									EndPos:    6102,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 365,
										EndLine:   365,
										StartPos:  6100,
										EndPos:    6102,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6105,
									EndPos:    6112,
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 366,
						EndLine:   366,
						StartPos:  6114,
						EndPos:    6143,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 366,
								EndLine:   366,
								StartPos:  6121,
								EndPos:    6142,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 366,
									EndLine:   366,
									StartPos:  6121,
									EndPos:    6123,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 366,
										EndLine:   366,
										StartPos:  6121,
										EndPos:    6123,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 366,
									EndLine:   366,
									StartPos:  6126,
									EndPos:    6142,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6132,
											EndPos:    6138,
										},
									},
									Key: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 366,
												EndLine:   366,
												StartPos:  6132,
												EndPos:    6133,
											},
										},
										Value: []byte("1"),
									},
									Val: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 366,
												EndLine:   366,
												StartPos:  6137,
												EndPos:    6138,
											},
										},
										Value: []byte("1"),
									},
								},
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6140,
											EndPos:    6141,
										},
									},
									Val: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 366,
												EndLine:   366,
												StartPos:  6140,
												EndPos:    6141,
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
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 367,
						EndLine:   367,
						StartPos:  6144,
						EndPos:    6171,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 367,
								EndLine:   367,
								StartPos:  6151,
								EndPos:    6170,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6151,
									EndPos:    6153,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 367,
										EndLine:   367,
										StartPos:  6151,
										EndPos:    6153,
									},
								},
								Value: []byte("$a"),
							},
						},
						Expr: &ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6156,
									EndPos:    6170,
								},
							},
							Var: &ast.ExprShortArray{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 367,
										EndLine:   367,
										StartPos:  6156,
										EndPos:    6167,
									},
								},
								Items: []ast.Vertex{
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6157,
												EndPos:    6158,
											},
										},
										Val: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 367,
													EndLine:   367,
													StartPos:  6157,
													EndPos:    6158,
												},
											},
											Value: []byte("1"),
										},
									},
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6160,
												EndPos:    6166,
											},
										},
										Key: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 367,
													EndLine:   367,
													StartPos:  6160,
													EndPos:    6161,
												},
											},
											Value: []byte("2"),
										},
										Val: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 367,
													EndLine:   367,
													StartPos:  6165,
													EndPos:    6166,
												},
											},
											Value: []byte("2"),
										},
									},
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 367,
										EndLine:   367,
										StartPos:  6168,
										EndPos:    6169,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 369,
						EndLine:   369,
						StartPos:  6173,
						EndPos:    6188,
					},
				},
				Cond: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 369,
							EndLine:   369,
							StartPos:  6177,
							EndPos:    6184,
						},
					},
					Value: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 369,
								EndLine:   369,
								StartPos:  6183,
								EndPos:    6184,
							},
						},
						Value: []byte("1"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 369,
							EndLine:   369,
							StartPos:  6186,
							EndPos:    6188,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 370,
						EndLine:   370,
						StartPos:  6189,
						EndPos:    6200,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 370,
							EndLine:   370,
							StartPos:  6189,
							EndPos:    6199,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 370,
								EndLine:   370,
								StartPos:  6189,
								EndPos:    6192,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 370,
										EndLine:   370,
										StartPos:  6189,
										EndPos:    6192,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 370,
								EndLine:   370,
								StartPos:  6194,
								EndPos:    6199,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 370,
									EndLine:   370,
									StartPos:  6195,
									EndPos:    6199,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 370,
										EndLine:   370,
										StartPos:  6195,
										EndPos:    6199,
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
						StartLine: 372,
						EndLine:   372,
						StartPos:  6202,
						EndPos:    6209,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 372,
							EndLine:   372,
							StartPos:  6202,
							EndPos:    6208,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 372,
								EndLine:   372,
								StartPos:  6202,
								EndPos:    6206,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 372,
									EndLine:   372,
									StartPos:  6202,
									EndPos:    6206,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 372,
								EndLine:   372,
								StartPos:  6206,
								EndPos:    6208,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 373,
						EndLine:   373,
						StartPos:  6210,
						EndPos:    6223,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 373,
							EndLine:   373,
							StartPos:  6210,
							EndPos:    6222,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 373,
								EndLine:   373,
								StartPos:  6210,
								EndPos:    6219,
							},
						},
						Var: &ast.ExprFunctionCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 373,
									EndLine:   373,
									StartPos:  6210,
									EndPos:    6216,
								},
							},
							Function: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 373,
										EndLine:   373,
										StartPos:  6210,
										EndPos:    6214,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 373,
											EndLine:   373,
											StartPos:  6210,
											EndPos:    6214,
										},
									},
									Value: []byte("$foo"),
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 373,
										EndLine:   373,
										StartPos:  6214,
										EndPos:    6216,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 373,
									EndLine:   373,
									StartPos:  6217,
									EndPos:    6218,
								},
							},
							Value: []byte("0"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 373,
								EndLine:   373,
								StartPos:  6220,
								EndPos:    6221,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 374,
						EndLine:   374,
						StartPos:  6224,
						EndPos:    6231,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 374,
							EndLine:   374,
							StartPos:  6224,
							EndPos:    6230,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6224,
								EndPos:    6226,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 374,
									EndLine:   374,
									StartPos:  6224,
									EndPos:    6226,
								},
							},
							Value: []byte("$a"),
						},
					},
					Dim: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6227,
								EndPos:    6229,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 374,
									EndLine:   374,
									StartPos:  6227,
									EndPos:    6229,
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
						StartLine: 375,
						EndLine:   375,
						StartPos:  6232,
						EndPos:    6238,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 375,
							EndLine:   375,
							StartPos:  6232,
							EndPos:    6237,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 375,
								EndLine:   375,
								StartPos:  6234,
								EndPos:    6236,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 375,
									EndLine:   375,
									StartPos:  6234,
									EndPos:    6236,
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
						StartLine: 376,
						EndLine:   376,
						StartPos:  6239,
						EndPos:    6254,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6239,
							EndPos:    6253,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6239,
								EndPos:    6243,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 376,
									EndLine:   376,
									StartPos:  6239,
									EndPos:    6243,
								},
							},
							Value: []byte("$foo"),
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6245,
								EndPos:    6251,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 376,
									EndLine:   376,
									StartPos:  6246,
									EndPos:    6250,
								},
							},
							Value: []byte("$bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6251,
								EndPos:    6253,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 377,
						EndLine:   377,
						StartPos:  6255,
						EndPos:    6265,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 377,
							EndLine:   377,
							StartPos:  6255,
							EndPos:    6264,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 377,
								EndLine:   377,
								StartPos:  6255,
								EndPos:    6259,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 377,
									EndLine:   377,
									StartPos:  6255,
									EndPos:    6259,
								},
							},
							Value: []byte("$foo"),
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 377,
								EndLine:   377,
								StartPos:  6261,
								EndPos:    6264,
							},
						},
						Value: []byte("bar"),
					},
				},
			},
			&ast.StmtHaltCompiler{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 379,
						EndLine:   379,
						StartPos:  6267,
						EndPos:    6285,
					},
				},
			},
		},
	}

	lexer := scanner.NewLexer([]byte(src), "5.6", nil)
	php5parser := php5.NewParser(lexer, nil)
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	traverser.NewDFS(new(visitor.FilterParserNodes)).Traverse(actual)
	traverser.NewDFS(new(visitor.FilterTokens)).Traverse(actual)
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

	lexer := scanner.NewLexer([]byte(src), "5.6", nil)
	php5parser := php5.NewParser(lexer, nil)
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	traverser.NewDFS(new(visitor.FilterParserNodes)).Traverse(actual)
	traverser.NewDFS(new(visitor.FilterTokens)).Traverse(actual)
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

	lexer := scanner.NewLexer([]byte(src), "5.6", nil)
	php5parser := php5.NewParser(lexer, nil)
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	traverser.NewDFS(new(visitor.FilterParserNodes)).Traverse(actual)
	traverser.NewDFS(new(visitor.FilterTokens)).Traverse(actual)
	assert.DeepEqual(t, expected, actual)
}

func TestPhp5ControlCharsErrors(t *testing.T) {
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

	lexer := scanner.NewLexer([]byte(src), "5.6", errorHandlerFunc)
	php5parser := php5.NewParser(lexer, errorHandlerFunc)
	php5parser.Parse()
	assert.DeepEqual(t, expected, parserErrors)
}
