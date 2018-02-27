// Package visitor contains walker.visitor implementations
package visitor_test

import (
	"bytes"

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

	nodes, comments, positions := php7.Parse(bytes.NewBufferString(src), "test.php")

	dumper := visitor.Dumper{
		Indent:    "| ",
		Comments:  comments,
		Positions: positions,
	}
	nodes.Walk(dumper)

	// Unordered output:
	//| *stmt.StmtList Pos{Line: 3-11 Pos: 10-143}
	//|   "Stmts":
	//|     *stmt.Namespace Pos{Line: 3-11 Pos: 10-143}
	//|       "NamespaceName":
	//|         *name.Name Pos{Line: 3-3 Pos: 20-22}
	//|           "Parts":
	//|             *name.NamePart Pos{Line: 3-3 Pos: 20-22}
	//|               "Value": Foo;
	//|       "Stmts":
	//|         *stmt.Class Pos{Line: 4-10 Pos: 29-139}
	//|           "PhpDocComment": ;
	//|           "ClassName":
	//|             *node.Identifier Pos{Line: 4-4 Pos: 35-37}
	//|               "Value": Bar;
	//|           "Stmts":
	//|             *stmt.ClassMethod Pos{Line: 5-9 Pos: 45-134}
	//|               "ReturnsRef": false;
	//|               "PhpDocComment": ;
	//|               "MethodName":
	//|                 *node.Identifier Pos{Line: 5-5 Pos: 61-72}
	//|                   "Value": FunctionName;
	//|               "Modifiers":
	//|                 *node.Identifier Pos{Line: 5-5 Pos: 45-50}
	//|                   "Value": public;
	//|               "Params":
	//|                 *node.Parameter Pos{Line: 5-5 Pos: 74-89}
	//|                   "ByRef": false;
	//|                   "Variadic": false;
	//|                   "VariableType":
	//|                     *name.Name Pos{Line: 5-5 Pos: 74-77}
	//|                       "Parts":
	//|                         *name.NamePart Pos{Line: 5-5 Pos: 74-77}
	//|                           "Value": Type;
	//|                   "Variable":
	//|                     *expr.Variable Pos{Line: 5-5 Pos: 79-82}
	//|                       "VarName":
	//|                         *node.Identifier Pos{Line: 5-5 Pos: 79-82}
	//|                           "Value": $var;
	//|                   "DefaultValue":
	//|                     *expr.ConstFetch Pos{Line: 5-5 Pos: 86-89}
	//|                       "Constant":
	//|                         *name.Name Pos{Line: 5-5 Pos: 86-89}
	//|                           "Parts":
	//|                             *name.NamePart Pos{Line: 5-5 Pos: 86-89}
	//|                               "Value": null;
	//|               "Stmts":
	//|                 *stmt.Expression Pos{Line: 8-8 Pos: 124-128}
	//|                   "Comments":
	//|                     "// some comment\n"
	//|                   "Expr":
	//|                     *expr.Variable Pos{Line: 8-8 Pos: 124-127}
	//|                       "Comments":
	//|                         "// some comment\n"
	//|                       "VarName":
	//|                         *node.Identifier Pos{Line: 8-8 Pos: 124-127}
	//|                           "Value": $var;
	//|                           "Comments":
	//|                             "// some comment\n"
}
