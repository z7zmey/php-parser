//Package visitor contains walker.visitor implementations
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
					//some comment
					$var;
				}
			}
		}`

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	nodes := php7parser.GetRootNode()

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := &visitor.GoDumper{
		Writer: os.Stdout,
	}
	nodes.Walk(dumper)

	// Unordered output:
	//&node.Root{
	//	Position: &position.Position{
	//		StartLine: 3,
	//		EndLine: 11,
	//		StartPos: 10,
	//		EndPos: 142,
	//	},
	//	Stmts: []node.Node{
	//		&stmt.Namespace{
	//			Position: &position.Position{
	//				StartLine: 3,
	//				EndLine: 11,
	//				StartPos: 10,
	//				EndPos: 142,
	//			},
	//			NamespaceName: &name.Name{
	//				Position: &position.Position{
	//					StartLine: 3,
	//					EndLine: 3,
	//					StartPos: 20,
	//					EndPos: 22,
	//				},
	//				Parts: []node.Node{
	//					&name.NamePart{
	//						Position: &position.Position{
	//							StartLine: 3,
	//							EndLine: 3,
	//							StartPos: 20,
	//							EndPos: 22,
	//						},
	//						Value: "Foo",
	//					},
	//				},
	//			},
	//			Stmts: []node.Node{
	//				&stmt.Class{
	//					Position: &position.Position{
	//						StartLine: 4,
	//						EndLine: 10,
	//						StartPos: 29,
	//						EndPos: 138,
	//					},
	//					PhpDocComment: "",
	//					ClassName: &node.Identifier{
	//						Position: &position.Position{
	//							StartLine: 4,
	//							EndLine: 4,
	//							StartPos: 35,
	//							EndPos: 37,
	//						},
	//						Value: "Bar",
	//					},
	//					Stmts: []node.Node{
	//						&stmt.ClassMethod{
	//							Position: &position.Position{
	//								StartLine: 5,
	//								EndLine: 9,
	//								StartPos: 45,
	//								EndPos: 133,
	//							},
	//							ReturnsRef: false,
	//							PhpDocComment: "",
	//							MethodName: &node.Identifier{
	//								Position: &position.Position{
	//									StartLine: 5,
	//									EndLine: 5,
	//									StartPos: 61,
	//									EndPos: 72,
	//								},
	//								Value: "FunctionName",
	//							},
	//							Modifiers: []node.Node{
	//								&node.Identifier{
	//									Position: &position.Position{
	//										StartLine: 5,
	//										EndLine: 5,
	//										StartPos: 45,
	//										EndPos: 50,
	//									},
	//									Value: "public",
	//								},
	//							},
	//							Params: []node.Node{
	//								&node.Parameter{
	//									Position: &position.Position{
	//										StartLine: 5,
	//										EndLine: 5,
	//										StartPos: 74,
	//										EndPos: 89,
	//									},
	//									Variadic: false,
	//									ByRef: false,
	//									VariableType: &name.Name{
	//										Position: &position.Position{
	//											StartLine: 5,
	//											EndLine: 5,
	//											StartPos: 74,
	//											EndPos: 77,
	//										},
	//										Parts: []node.Node{
	//											&name.NamePart{
	//												Position: &position.Position{
	//													StartLine: 5,
	//													EndLine: 5,
	//													StartPos: 74,
	//													EndPos: 77,
	//												},
	//												Value: "Type",
	//											},
	//										},
	//									},
	//									Variable: &expr.Variable{
	//										Position: &position.Position{
	//											StartLine: 5,
	//											EndLine: 5,
	//											StartPos: 79,
	//											EndPos: 82,
	//										},
	//										VarName: &node.Identifier{
	//											Position: &position.Position{
	//												StartLine: 5,
	//												EndLine: 5,
	//												StartPos: 79,
	//												EndPos: 82,
	//											},
	//											Value: "var",
	//										},
	//									},
	//									DefaultValue: &expr.ConstFetch{
	//										Position: &position.Position{
	//											StartLine: 5,
	//											EndLine: 5,
	//											StartPos: 86,
	//											EndPos: 89,
	//										},
	//										Constant: &name.Name{
	//											Position: &position.Position{
	//												StartLine: 5,
	//												EndLine: 5,
	//												StartPos: 86,
	//												EndPos: 89,
	//											},
	//											Parts: []node.Node{
	//												&name.NamePart{
	//													Position: &position.Position{
	//														StartLine: 5,
	//														EndLine: 5,
	//														StartPos: 86,
	//														EndPos: 89,
	//													},
	//													Value: "null",
	//												},
	//											},
	//										},
	//									},
	//								},
	//							},
	//							Stmt: &stmt.StmtList{
	//								Position: &position.Position{
	//									StartLine: 6,
	//									EndLine: 9,
	//									StartPos: 96,
	//									EndPos: 133,
	//								},
	//								Stmts: []node.Node{
	//									&stmt.Expression{
	//										Position: &position.Position{
	//											StartLine: 8,
	//											EndLine: 8,
	//											StartPos: 123,
	//											EndPos: 127,
	//										},
	//										Expr: &expr.Variable{
	//											Position: &position.Position{
	//												StartLine: 8,
	//												EndLine: 8,
	//												StartPos: 123,
	//												EndPos: 126,
	//											},
	//											VarName: &node.Identifier{
	//												Position: &position.Position{
	//													StartLine: 8,
	//													EndLine: 8,
	//													StartPos: 123,
	//													EndPos: 126,
	//												},
	//												Value: "var",
	//											},
	//										},
	//									},
	//								},
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
}
