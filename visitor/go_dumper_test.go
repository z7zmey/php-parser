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
	//	Stmts: []node.Node{
	//		&stmt.Namespace{
	//			NamespaceName: &name.Name{
	//				Parts: []node.Node{
	//					&name.NamePart{
	//						Value: "Foo",
	//					},
	//				},
	//			},
	//			Stmts: []node.Node{
	//				&stmt.Class{
	//					PhpDocComment: "",
	//					ClassName: &node.Identifier{
	//						Value: "Bar",
	//					},
	//					Stmts: []node.Node{
	//						&stmt.ClassMethod{
	//							ReturnsRef: false,
	//							PhpDocComment: "",
	//							MethodName: &node.Identifier{
	//								Value: "FunctionName",
	//							},
	//							Modifiers: []node.Node{
	//								&node.Identifier{
	//									Value: "public",
	//								},
	//							},
	//							Params: []node.Node{
	//								&node.Parameter{
	//									ByRef: false,
	//									Variadic: false,
	//									VariableType: &name.Name{
	//										Parts: []node.Node{
	//											&name.NamePart{
	//												Value: "Type",
	//											},
	//										},
	//									},
	//									Variable: &expr.Variable{
	//										VarName: &node.Identifier{
	//											Value: "var",
	//										},
	//									},
	//									DefaultValue: &expr.ConstFetch{
	//										Constant: &name.Name{
	//											Parts: []node.Node{
	//												&name.NamePart{
	//													Value: "null",
	//												},
	//											},
	//										},
	//									},
	//								},
	//							},
	//							Stmt: &stmt.StmtList{
	//								Stmts: []node.Node{
	//									&stmt.Expression{
	//										Expr: &expr.Variable{
	//											VarName: &node.Identifier{
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
