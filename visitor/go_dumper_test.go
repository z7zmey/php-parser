package visitor_test

import (
	"os"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func ExampleGoDumper() {
	src := `<?php

		namespace Foo {
			class Bar {
				public function FunctionName(Type $var = null)
				{
					// some comment
					$var;
				}
			}
		}`

	php7parser := php7.NewParser([]byte(src))
	php7parser.WithFreeFloating()
	php7parser.Parse()
	nodes := php7parser.GetRootNode()

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := &visitor.GoDumper{
		Writer: os.Stdout,
	}
	nodes.Walk(dumper)

	// Unordered output:
	// &node.Root{
	// 	Position: &position.Position{
	// 		StartLine: 3,
	// 		EndLine: 11,
	// 		StartPos: 9,
	// 		EndPos: 143,
	// 	},
	// 	Stmts: []node.Node{
	// 		&stmt.Namespace{
	// 			Position: &position.Position{
	// 				StartLine: 3,
	// 				EndLine: 11,
	// 				StartPos: 9,
	// 				EndPos: 143,
	// 			},
	// 			FreeFloating: freefloating.Collection{
	// 				"Stmts": []freefloating.String{
	// 					freefloating.String{
	// 						Type: freefloating.WhiteSpaceType,
	// 						Position: &position.Position{
	// 							StartLine: 10,
	// 							EndLine: 11,
	// 							StartPos: 139,
	// 							EndPos: 142,
	// 						},
	// 						Value: "\n\t\t",
	// 					},
	// 				},
	// 				"Start": []freefloating.String{
	// 					freefloating.String{
	// 						Type: freefloating.TokenType,
	// 						Position: &position.Position{
	// 							StartLine: 1,
	// 							EndLine: 1,
	// 							StartPos: 0,
	// 							EndPos: 5,
	// 						},
	// 						Value: "<?php",
	// 					},
	// 					freefloating.String{
	// 						Type: freefloating.WhiteSpaceType,
	// 						Position: &position.Position{
	// 							StartLine: 1,
	// 							EndLine: 3,
	// 							StartPos: 5,
	// 							EndPos: 9,
	// 						},
	// 						Value: "\n\n\t\t",
	// 					},
	// 				},
	// 			},
	// 			NamespaceName: &name.Name{
	// 				Position: &position.Position{
	// 					StartLine: 3,
	// 					EndLine: 3,
	// 					StartPos: 19,
	// 					EndPos: 22,
	// 				},
	// 				FreeFloating: freefloating.Collection{
	// 					"Start": []freefloating.String{
	// 						freefloating.String{
	// 							Type: freefloating.WhiteSpaceType,
	// 							Position: &position.Position{
	// 								StartLine: 3,
	// 								EndLine: 3,
	// 								StartPos: 18,
	// 								EndPos: 19,
	// 							},
	// 							Value: " ",
	// 						},
	// 					},
	// 					"End": []freefloating.String{
	// 						freefloating.String{
	// 							Type: freefloating.WhiteSpaceType,
	// 							Position: &position.Position{
	// 								StartLine: 3,
	// 								EndLine: 3,
	// 								StartPos: 22,
	// 								EndPos: 23,
	// 							},
	// 							Value: " ",
	// 						},
	// 					},
	// 				},
	// 				Parts: []node.Node{
	// 					&name.NamePart{
	// 						Position: &position.Position{
	// 							StartLine: 3,
	// 							EndLine: 3,
	// 							StartPos: 19,
	// 							EndPos: 22,
	// 						},
	// 						Value: "Foo",
	// 					},
	// 				},
	// 			},
	// 			Stmts: []node.Node{
	// 				&stmt.Class{
	// 					Position: &position.Position{
	// 						StartLine: 4,
	// 						EndLine: 10,
	// 						StartPos: 28,
	// 						EndPos: 139,
	// 					},
	// 					FreeFloating: freefloating.Collection{
	// 						"Start": []freefloating.String{
	// 							freefloating.String{
	// 								Type: freefloating.WhiteSpaceType,
	// 								Position: &position.Position{
	// 									StartLine: 3,
	// 									EndLine: 4,
	// 									StartPos: 24,
	// 									EndPos: 28,
	// 								},
	// 								Value: "\n\t\t\t",
	// 							},
	// 						},
	// 						"Name": []freefloating.String{
	// 							freefloating.String{
	// 								Type: freefloating.WhiteSpaceType,
	// 								Position: &position.Position{
	// 									StartLine: 4,
	// 									EndLine: 4,
	// 									StartPos: 37,
	// 									EndPos: 38,
	// 								},
	// 								Value: " ",
	// 							},
	// 						},
	// 						"Stmts": []freefloating.String{
	// 							freefloating.String{
	// 								Type: freefloating.WhiteSpaceType,
	// 								Position: &position.Position{
	// 									StartLine: 9,
	// 									EndLine: 10,
	// 									StartPos: 134,
	// 									EndPos: 138,
	// 								},
	// 								Value: "\n\t\t\t",
	// 							},
	// 						},
	// 					},
	// 					PhpDocComment: "",
	// 					ClassName: &node.Identifier{
	// 						Position: &position.Position{
	// 							StartLine: 4,
	// 							EndLine: 4,
	// 							StartPos: 34,
	// 							EndPos: 37,
	// 						},
	// 						FreeFloating: freefloating.Collection{
	// 							"Start": []freefloating.String{
	// 								freefloating.String{
	// 									Type: freefloating.WhiteSpaceType,
	// 									Position: &position.Position{
	// 										StartLine: 4,
	// 										EndLine: 4,
	// 										StartPos: 33,
	// 										EndPos: 34,
	// 									},
	// 									Value: " ",
	// 								},
	// 							},
	// 						},
	// 						Value: "Bar",
	// 					},
	// 					Stmts: []node.Node{
	// 						&stmt.ClassMethod{
	// 							Position: &position.Position{
	// 								StartLine: 5,
	// 								EndLine: 9,
	// 								StartPos: 44,
	// 								EndPos: 134,
	// 							},
	// 							FreeFloating: freefloating.Collection{
	// 								"Start": []freefloating.String{
	// 									freefloating.String{
	// 										Type: freefloating.WhiteSpaceType,
	// 										Position: &position.Position{
	// 											StartLine: 4,
	// 											EndLine: 5,
	// 											StartPos: 39,
	// 											EndPos: 44,
	// 										},
	// 										Value: "\n\t\t\t\t",
	// 									},
	// 								},
	// 								"ModifierList": []freefloating.String{
	// 									freefloating.String{
	// 										Type: freefloating.WhiteSpaceType,
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 50,
	// 											EndPos: 51,
	// 										},
	// 										Value: " ",
	// 									},
	// 								},
	// 								"Function": []freefloating.String{
	// 									freefloating.String{
	// 										Type: freefloating.WhiteSpaceType,
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 59,
	// 											EndPos: 60,
	// 										},
	// 										Value: " ",
	// 									},
	// 								},
	// 							},
	// 							ReturnsRef: false,
	// 							PhpDocComment: "",
	// 							MethodName: &node.Identifier{
	// 								Position: &position.Position{
	// 									StartLine: 5,
	// 									EndLine: 5,
	// 									StartPos: 60,
	// 									EndPos: 72,
	// 								},
	// 								Value: "FunctionName",
	// 							},
	// 							Modifiers: []node.Node{
	// 								&node.Identifier{
	// 									Position: &position.Position{
	// 										StartLine: 5,
	// 										EndLine: 5,
	// 										StartPos: 44,
	// 										EndPos: 50,
	// 									},
	// 									Value: "public",
	// 								},
	// 							},
	// 							Params: []node.Node{
	// 								&node.Parameter{
	// 									Position: &position.Position{
	// 										StartLine: 5,
	// 										EndLine: 5,
	// 										StartPos: 73,
	// 										EndPos: 89,
	// 									},
	// 									FreeFloating: freefloating.Collection{
	// 										"OptionalType": []freefloating.String{
	// 											freefloating.String{
	// 												Type: freefloating.WhiteSpaceType,
	// 												Position: &position.Position{
	// 													StartLine: 5,
	// 													EndLine: 5,
	// 													StartPos: 77,
	// 													EndPos: 78,
	// 												},
	// 												Value: " ",
	// 											},
	// 										},
	// 										"Var": []freefloating.String{
	// 											freefloating.String{
	// 												Type: freefloating.WhiteSpaceType,
	// 												Position: &position.Position{
	// 													StartLine: 5,
	// 													EndLine: 5,
	// 													StartPos: 82,
	// 													EndPos: 83,
	// 												},
	// 												Value: " ",
	// 											},
	// 										},
	// 									},
	// 									ByRef: false,
	// 									Variadic: false,
	// 									VariableType: &name.Name{
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 73,
	// 											EndPos: 77,
	// 										},
	// 										Parts: []node.Node{
	// 											&name.NamePart{
	// 												Position: &position.Position{
	// 													StartLine: 5,
	// 													EndLine: 5,
	// 													StartPos: 73,
	// 													EndPos: 77,
	// 												},
	// 												Value: "Type",
	// 											},
	// 										},
	// 									},
	// 									Variable: &expr.Variable{
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 78,
	// 											EndPos: 82,
	// 										},
	// 										FreeFloating: freefloating.Collection{
	// 											"Dollar": []freefloating.String{
	// 												freefloating.String{
	// 													Type: freefloating.TokenType,
	// 													Position: &position.Position{
	// 														StartLine: 5,
	// 														EndLine: 5,
	// 														StartPos: 78,
	// 														EndPos: 79,
	// 													},
	// 													Value: "$",
	// 												},
	// 											},
	// 										},
	// 										VarName: &node.Identifier{
	// 											Position: &position.Position{
	// 												StartLine: 5,
	// 												EndLine: 5,
	// 												StartPos: 78,
	// 												EndPos: 82,
	// 											},
	// 											Value: "var",
	// 										},
	// 									},
	// 									DefaultValue: &expr.ConstFetch{
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 85,
	// 											EndPos: 89,
	// 										},
	// 										FreeFloating: freefloating.Collection{
	// 											"Start": []freefloating.String{
	// 												freefloating.String{
	// 													Type: freefloating.WhiteSpaceType,
	// 													Position: &position.Position{
	// 														StartLine: 5,
	// 														EndLine: 5,
	// 														StartPos: 84,
	// 														EndPos: 85,
	// 													},
	// 													Value: " ",
	// 												},
	// 											},
	// 										},
	// 										Constant: &name.Name{
	// 											Position: &position.Position{
	// 												StartLine: 5,
	// 												EndLine: 5,
	// 												StartPos: 85,
	// 												EndPos: 89,
	// 											},
	// 											Parts: []node.Node{
	// 												&name.NamePart{
	// 													Position: &position.Position{
	// 														StartLine: 5,
	// 														EndLine: 5,
	// 														StartPos: 85,
	// 														EndPos: 89,
	// 													},
	// 													Value: "null",
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 							Stmt: &stmt.StmtList{
	// 								Position: &position.Position{
	// 									StartLine: 6,
	// 									EndLine: 9,
	// 									StartPos: 95,
	// 									EndPos: 134,
	// 								},
	// 								FreeFloating: freefloating.Collection{
	// 									"Start": []freefloating.String{
	// 										freefloating.String{
	// 											Type: freefloating.WhiteSpaceType,
	// 											Position: &position.Position{
	// 												StartLine: 5,
	// 												EndLine: 6,
	// 												StartPos: 90,
	// 												EndPos: 95,
	// 											},
	// 											Value: "\n\t\t\t\t",
	// 										},
	// 									},
	// 									"Stmts": []freefloating.String{
	// 										freefloating.String{
	// 											Type: freefloating.WhiteSpaceType,
	// 											Position: &position.Position{
	// 												StartLine: 8,
	// 												EndLine: 9,
	// 												StartPos: 128,
	// 												EndPos: 133,
	// 											},
	// 											Value: "\n\t\t\t\t",
	// 										},
	// 									},
	// 								},
	// 								Stmts: []node.Node{
	// 									&stmt.Expression{
	// 										Position: &position.Position{
	// 											StartLine: 8,
	// 											EndLine: 8,
	// 											StartPos: 123,
	// 											EndPos: 128,
	// 										},
	// 										FreeFloating: freefloating.Collection{
	// 											"Start": []freefloating.String{
	// 												freefloating.String{
	// 													Type: freefloating.WhiteSpaceType,
	// 													Position: &position.Position{
	// 														StartLine: 6,
	// 														EndLine: 7,
	// 														StartPos: 96,
	// 														EndPos: 102,
	// 													},
	// 													Value: "\n\t\t\t\t\t",
	// 												},
	// 												freefloating.String{
	// 													Type: freefloating.CommentType,
	// 													Position: &position.Position{
	// 														StartLine: 7,
	// 														EndLine: 7,
	// 														StartPos: 102,
	// 														EndPos: 118,
	// 													},
	// 													Value: "// some comment\n",
	// 												},
	// 												freefloating.String{
	// 													Type: freefloating.WhiteSpaceType,
	// 													Position: &position.Position{
	// 														StartLine: 8,
	// 														EndLine: 8,
	// 														StartPos: 118,
	// 														EndPos: 123,
	// 													},
	// 													Value: "\t\t\t\t\t",
	// 												},
	// 											},
	// 											"SemiColon": []freefloating.String{
	// 												freefloating.String{
	// 													Type: freefloating.TokenType,
	// 													Position: &position.Position{
	// 														StartLine: 8,
	// 														EndLine: 8,
	// 														StartPos: 127,
	// 														EndPos: 128,
	// 													},
	// 													Value: ";",
	// 												},
	// 											},
	// 										},
	// 										Expr: &expr.Variable{
	// 											Position: &position.Position{
	// 												StartLine: 8,
	// 												EndLine: 8,
	// 												StartPos: 123,
	// 												EndPos: 127,
	// 											},
	// 											FreeFloating: freefloating.Collection{
	// 												"Dollar": []freefloating.String{
	// 													freefloating.String{
	// 														Type: freefloating.TokenType,
	// 														Position: &position.Position{
	// 															StartLine: 8,
	// 															EndLine: 8,
	// 															StartPos: 123,
	// 															EndPos: 124,
	// 														},
	// 														Value: "$",
	// 													},
	// 												},
	// 											},
	// 											VarName: &node.Identifier{
	// 												Position: &position.Position{
	// 													StartLine: 8,
	// 													EndLine: 8,
	// 													StartPos: 123,
	// 													EndPos: 127,
	// 												},
	// 												Value: "var",
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
