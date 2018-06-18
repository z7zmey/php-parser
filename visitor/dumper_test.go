// Package visitor contains walker.visitor implementations
package visitor_test

import (
	"bytes"
	"os"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func ExampleDumper() {
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
	php7parser.Parse()
	nodes := php7parser.GetRootNode()

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := &visitor.Dumper{
		Writer:     os.Stdout,
		Indent:     "| ",
		Comments:   php7parser.GetComments(),
		Positions:  php7parser.GetPositions(),
		NsResolver: nsResolver,
	}
	nodes.Walk(dumper)

	// Unordered output:
	//| [*node.Root]
	//|   "Position": Pos{Line: 3-11 Pos: 10-143};
	//|   "Stmts":
	//|     [*stmt.Namespace]
	//|       "Position": Pos{Line: 3-11 Pos: 10-143};
	//|       "NamespaceName":
	//|         [*name.Name]
	//|           "Position": Pos{Line: 3-3 Pos: 20-22};
	//|           "Parts":
	//|             [*name.NamePart]
	//|               "Position": Pos{Line: 3-3 Pos: 20-22};
	//|               "Value": Foo;
	//|       "Stmts":
	//|         [*stmt.Class]
	//|           "Position": Pos{Line: 4-10 Pos: 29-139};
	//|           "NamespacedName": Foo\Bar;
	//|           "PhpDocComment": ;
	//|           "ClassName":
	//|             [*node.Identifier]
	//|               "Position": Pos{Line: 4-4 Pos: 35-37};
	//|               "Value": Bar;
	//|           "Stmts":
	//|             [*stmt.ClassMethod]
	//|               "Position": Pos{Line: 5-9 Pos: 45-134};
	//|               "ReturnsRef": false;
	//|               "PhpDocComment": ;
	//|               "MethodName":
	//|                 [*node.Identifier]
	//|                   "Position": Pos{Line: 5-5 Pos: 61-72};
	//|                   "Value": FunctionName;
	//|               "Modifiers":
	//|                 [*node.Identifier]
	//|                   "Position": Pos{Line: 5-5 Pos: 45-50};
	//|                   "Value": public;
	//|               "Params":
	//|                 [*node.Parameter]
	//|                   "Position": Pos{Line: 5-5 Pos: 74-89};
	//|                   "ByRef": false;
	//|                   "Variadic": false;
	//|                   "VariableType":
	//|                     [*name.Name]
	//|                       "Position": Pos{Line: 5-5 Pos: 74-77};
	//|                       "NamespacedName": Foo\Type;
	//|                       "Parts":
	//|                         [*name.NamePart]
	//|                           "Position": Pos{Line: 5-5 Pos: 74-77};
	//|                           "Value": Type;
	//|                   "Variable":
	//|                     [*expr.Variable]
	//|                       "Position": Pos{Line: 5-5 Pos: 79-82};
	//|                       "VarName":
	//|                         [*node.Identifier]
	//|                           "Position": Pos{Line: 5-5 Pos: 79-82};
	//|                           "Value": var;
	//|                   "DefaultValue":
	//|                     [*expr.ConstFetch]
	//|                       "Position": Pos{Line: 5-5 Pos: 86-89};
	//|                       "Constant":
	//|                         [*name.Name]
	//|                           "Position": Pos{Line: 5-5 Pos: 86-89};
	//|                           "NamespacedName": null;
	//|                           "Parts":
	//|                             [*name.NamePart]
	//|                               "Position": Pos{Line: 5-5 Pos: 86-89};
	//|                               "Value": null;
	//|               "Stmt":
	//|                 [*stmt.StmtList]
	//|                   "Position": Pos{Line: 6-9 Pos: 96-134};
	//|                   "Stmts":
	//|                     [*stmt.Expression]
	//|                       "Position": Pos{Line: 8-8 Pos: 124-128};
	//|                       "Expr":
	//|                         [*expr.Variable]
	//|                           "Position": Pos{Line: 8-8 Pos: 124-127};
	//|                           "Comments":
	//|                             "// some comment\n" before "VariableToken"
	//|                           "VarName":
	//|                             [*node.Identifier]
	//|                               "Position": Pos{Line: 8-8 Pos: 124-127};
	//|                               "Value": var;
}
