package visitor_test

import (
	"bytes"
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

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.WithMeta()
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
	// 		StartPos: 10,
	// 		EndPos: 143,
	// 	},
	// 	Stmts: []node.Node{
	// 		&stmt.Namespace{
	// 			Position: &position.Position{
	// 				StartLine: 3,
	// 				EndLine: 11,
	// 				StartPos: 10,
	// 				EndPos: 143,
	// 			},
	// 			Meta: meta.Collection{
	// 				&meta.Data{
	// 					Type: meta.TokenType,
	// 					Position: &position.Position{
	// 						StartLine: 1,
	// 						EndLine: 1,
	// 						StartPos: 1,
	// 						EndPos: 5,
	// 					},
	// 					Value: "<?php",
	// 					TokenName: meta.NamespaceToken,
	// 				},
	// 				&meta.Data{
	// 					Type: meta.WhiteSpaceType,
	// 					Position: &position.Position{
	// 						StartLine: 1,
	// 						EndLine: 3,
	// 						StartPos: 6,
	// 						EndPos: 9,
	// 					},
	// 					Value: "\n\n\t\t",
	// 					TokenName: meta.NamespaceToken,
	// 				},
	// 				&meta.Data{
	// 					Type: meta.WhiteSpaceType,
	// 					Position: &position.Position{
	// 						StartLine: 3,
	// 						EndLine: 3,
	// 						StartPos: 23,
	// 						EndPos: 23,
	// 					},
	// 					Value: " ",
	// 					TokenName: meta.OpenCurlyBracesToken,
	// 				},
	// 				&meta.Data{
	// 					Type: meta.WhiteSpaceType,
	// 					Position: &position.Position{
	// 						StartLine: 10,
	// 						EndLine: 11,
	// 						StartPos: 140,
	// 						EndPos: 142,
	// 					},
	// 					Value: "\n\t\t",
	// 					TokenName: meta.CloseCurlyBracesToken,
	// 				},
	// 			},
	// 			NamespaceName: &name.Name{
	// 				Position: &position.Position{
	// 					StartLine: 3,
	// 					EndLine: 3,
	// 					StartPos: 20,
	// 					EndPos: 22,
	// 				},
	// 				Parts: []node.Node{
	// 					&name.NamePart{
	// 						Position: &position.Position{
	// 							StartLine: 3,
	// 							EndLine: 3,
	// 							StartPos: 20,
	// 							EndPos: 22,
	// 						},
	// 						Meta: meta.Collection{
	// 							&meta.Data{
	// 								Type: meta.WhiteSpaceType,
	// 								Position: &position.Position{
	// 									StartLine: 3,
	// 									EndLine: 3,
	// 									StartPos: 19,
	// 									EndPos: 19,
	// 								},
	// 								Value: " ",
	// 								TokenName: meta.NodeStart,
	// 							},
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
	// 						StartPos: 29,
	// 						EndPos: 139,
	// 					},
	// 					Meta: meta.Collection{
	// 						&meta.Data{
	// 							Type: meta.WhiteSpaceType,
	// 							Position: &position.Position{
	// 								StartLine: 3,
	// 								EndLine: 4,
	// 								StartPos: 25,
	// 								EndPos: 28,
	// 							},
	// 							Value: "\n\t\t\t",
	// 							TokenName: meta.ClassToken,
	// 						},
	// 						&meta.Data{
	// 							Type: meta.WhiteSpaceType,
	// 							Position: &position.Position{
	// 								StartLine: 4,
	// 								EndLine: 4,
	// 								StartPos: 38,
	// 								EndPos: 38,
	// 							},
	// 							Value: " ",
	// 							TokenName: meta.OpenCurlyBracesToken,
	// 						},
	// 						&meta.Data{
	// 							Type: meta.WhiteSpaceType,
	// 							Position: &position.Position{
	// 								StartLine: 9,
	// 								EndLine: 10,
	// 								StartPos: 135,
	// 								EndPos: 138,
	// 							},
	// 							Value: "\n\t\t\t",
	// 							TokenName: meta.CloseCurlyBracesToken,
	// 						},
	// 					},
	// 					PhpDocComment: "",
	// 					ClassName: &node.Identifier{
	// 						Position: &position.Position{
	// 							StartLine: 4,
	// 							EndLine: 4,
	// 							StartPos: 35,
	// 							EndPos: 37,
	// 						},
	// 						Meta: meta.Collection{
	// 							&meta.Data{
	// 								Type: meta.WhiteSpaceType,
	// 								Position: &position.Position{
	// 									StartLine: 4,
	// 									EndLine: 4,
	// 									StartPos: 34,
	// 									EndPos: 34,
	// 								},
	// 								Value: " ",
	// 								TokenName: meta.NodeStart,
	// 							},
	// 						},
	// 						Value: "Bar",
	// 					},
	// 					Stmts: []node.Node{
	// 						&stmt.ClassMethod{
	// 							Position: &position.Position{
	// 								StartLine: 5,
	// 								EndLine: 9,
	// 								StartPos: 45,
	// 								EndPos: 134,
	// 							},
	// 							Meta: meta.Collection{
	// 								&meta.Data{
	// 									Type: meta.WhiteSpaceType,
	// 									Position: &position.Position{
	// 										StartLine: 5,
	// 										EndLine: 5,
	// 										StartPos: 51,
	// 										EndPos: 51,
	// 									},
	// 									Value: " ",
	// 									TokenName: meta.FunctionToken,
	// 								},
	// 							},
	// 							ReturnsRef: false,
	// 							PhpDocComment: "",
	// 							MethodName: &node.Identifier{
	// 								Position: &position.Position{
	// 									StartLine: 5,
	// 									EndLine: 5,
	// 									StartPos: 61,
	// 									EndPos: 72,
	// 								},
	// 								Meta: meta.Collection{
	// 									&meta.Data{
	// 										Type: meta.WhiteSpaceType,
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 60,
	// 											EndPos: 60,
	// 										},
	// 										Value: " ",
	// 										TokenName: meta.NodeStart,
	// 									},
	// 								},
	// 								Value: "FunctionName",
	// 							},
	// 							Modifiers: []node.Node{
	// 								&node.Identifier{
	// 									Position: &position.Position{
	// 										StartLine: 5,
	// 										EndLine: 5,
	// 										StartPos: 45,
	// 										EndPos: 50,
	// 									},
	// 									Meta: meta.Collection{
	// 										&meta.Data{
	// 											Type: meta.WhiteSpaceType,
	// 											Position: &position.Position{
	// 												StartLine: 4,
	// 												EndLine: 5,
	// 												StartPos: 40,
	// 												EndPos: 44,
	// 											},
	// 											Value: "\n\t\t\t\t",
	// 											TokenName: meta.NodeStart,
	// 										},
	// 									},
	// 									Value: "public",
	// 								},
	// 							},
	// 							Params: []node.Node{
	// 								&node.Parameter{
	// 									Position: &position.Position{
	// 										StartLine: 5,
	// 										EndLine: 5,
	// 										StartPos: 74,
	// 										EndPos: 89,
	// 									},
	// 									Meta: meta.Collection{
	// 										&meta.Data{
	// 											Type: meta.WhiteSpaceType,
	// 											Position: &position.Position{
	// 												StartLine: 5,
	// 												EndLine: 5,
	// 												StartPos: 83,
	// 												EndPos: 83,
	// 											},
	// 											Value: " ",
	// 											TokenName: meta.EqualToken,
	// 										},
	// 									},
	// 									ByRef: false,
	// 									Variadic: false,
	// 									VariableType: &name.Name{
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 74,
	// 											EndPos: 77,
	// 										},
	// 										Parts: []node.Node{
	// 											&name.NamePart{
	// 												Position: &position.Position{
	// 													StartLine: 5,
	// 													EndLine: 5,
	// 													StartPos: 74,
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
	// 											StartPos: 79,
	// 											EndPos: 82,
	// 										},
	// 										Meta: meta.Collection{
	// 											&meta.Data{
	// 												Type: meta.WhiteSpaceType,
	// 												Position: &position.Position{
	// 													StartLine: 5,
	// 													EndLine: 5,
	// 													StartPos: 78,
	// 													EndPos: 78,
	// 												},
	// 												Value: " ",
	// 												TokenName: meta.NodeStart,
	// 											},
	// 											&meta.Data{
	// 												Type: meta.TokenType,
	// 												Position: nil,
	// 												Value: "$",
	// 												TokenName: meta.NodeStart,
	// 											},
	// 										},
	// 										VarName: &node.Identifier{
	// 											Position: &position.Position{
	// 												StartLine: 5,
	// 												EndLine: 5,
	// 												StartPos: 79,
	// 												EndPos: 82,
	// 											},
	// 											Value: "var",
	// 										},
	// 									},
	// 									DefaultValue: &expr.ConstFetch{
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 5,
	// 											StartPos: 86,
	// 											EndPos: 89,
	// 										},
	// 										Constant: &name.Name{
	// 											Position: &position.Position{
	// 												StartLine: 5,
	// 												EndLine: 5,
	// 												StartPos: 86,
	// 												EndPos: 89,
	// 											},
	// 											Parts: []node.Node{
	// 												&name.NamePart{
	// 													Position: &position.Position{
	// 														StartLine: 5,
	// 														EndLine: 5,
	// 														StartPos: 86,
	// 														EndPos: 89,
	// 													},
	// 													Meta: meta.Collection{
	// 														&meta.Data{
	// 															Type: meta.WhiteSpaceType,
	// 															Position: &position.Position{
	// 																StartLine: 5,
	// 																EndLine: 5,
	// 																StartPos: 85,
	// 																EndPos: 85,
	// 															},
	// 															Value: " ",
	// 															TokenName: meta.NodeStart,
	// 														},
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
	// 									StartPos: 96,
	// 									EndPos: 134,
	// 								},
	// 								Meta: meta.Collection{
	// 									&meta.Data{
	// 										Type: meta.WhiteSpaceType,
	// 										Position: &position.Position{
	// 											StartLine: 5,
	// 											EndLine: 6,
	// 											StartPos: 91,
	// 											EndPos: 95,
	// 										},
	// 										Value: "\n\t\t\t\t",
	// 										TokenName: meta.OpenCurlyBracesToken,
	// 									},
	// 									&meta.Data{
	// 										Type: meta.WhiteSpaceType,
	// 										Position: &position.Position{
	// 											StartLine: 8,
	// 											EndLine: 9,
	// 											StartPos: 129,
	// 											EndPos: 133,
	// 										},
	// 										Value: "\n\t\t\t\t",
	// 										TokenName: meta.CloseCurlyBracesToken,
	// 									},
	// 								},
	// 								Stmts: []node.Node{
	// 									&stmt.Expression{
	// 										Position: &position.Position{
	// 											StartLine: 8,
	// 											EndLine: 8,
	// 											StartPos: 124,
	// 											EndPos: 128,
	// 										},
	// 										Meta: meta.Collection{
	// 											&meta.Data{
	// 												Type: meta.TokenType,
	// 												Position: &position.Position{
	// 													StartLine: 8,
	// 													EndLine: 8,
	// 													StartPos: 128,
	// 													EndPos: 128,
	// 												},
	// 												Value: ";",
	// 												TokenName: meta.SemiColonToken,
	// 											},
	// 										},
	// 										Expr: &expr.Variable{
	// 											Position: &position.Position{
	// 												StartLine: 8,
	// 												EndLine: 8,
	// 												StartPos: 124,
	// 												EndPos: 127,
	// 											},
	// 											Meta: meta.Collection{
	// 												&meta.Data{
	// 													Type: meta.WhiteSpaceType,
	// 													Position: &position.Position{
	// 														StartLine: 6,
	// 														EndLine: 7,
	// 														StartPos: 97,
	// 														EndPos: 102,
	// 													},
	// 													Value: "\n\t\t\t\t\t",
	// 													TokenName: meta.NodeStart,
	// 												},
	// 												&meta.Data{
	// 													Type: meta.CommentType,
	// 													Position: &position.Position{
	// 														StartLine: 7,
	// 														EndLine: 7,
	// 														StartPos: 103,
	// 														EndPos: 118,
	// 													},
	// 													Value: "// some comment\n",
	// 													TokenName: meta.NodeStart,
	// 												},
	// 												&meta.Data{
	// 													Type: meta.WhiteSpaceType,
	// 													Position: &position.Position{
	// 														StartLine: 8,
	// 														EndLine: 8,
	// 														StartPos: 119,
	// 														EndPos: 123,
	// 													},
	// 													Value: "\t\t\t\t\t",
	// 													TokenName: meta.NodeStart,
	// 												},
	// 												&meta.Data{
	// 													Type: meta.TokenType,
	// 													Position: nil,
	// 													Value: "$",
	// 													TokenName: meta.NodeStart,
	// 												},
	// 											},
	// 											VarName: &node.Identifier{
	// 												Position: &position.Position{
	// 													StartLine: 8,
	// 													EndLine: 8,
	// 													StartPos: 124,
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
