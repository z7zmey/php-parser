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
	php7parser.WithMeta()
	php7parser.Parse()
	nodes := php7parser.GetRootNode()

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := &visitor.Dumper{
		Writer:     os.Stdout,
		Indent:     "| ",
		NsResolver: nsResolver,
	}
	nodes.Walk(dumper)

	// Unordered output:
	// | [*node.Root]
	// |   "Position": Pos{Line: 3-11 Pos: 10-143}
	// |   "Stmts":
	// |     [*stmt.Namespace]
	// |       "Position": Pos{Line: 3-11 Pos: 10-143}
	// |       "Meta":
	// |         "<?php" before "NodeStart"
	// |         "\n\n\t\t" before "NodeStart"
	// |         " " before "OpenCurlyBracesToken"
	// |         "\n\t\t" before "CloseCurlyBracesToken"
	// |       "NamespaceName":
	// |         [*name.Name]
	// |           "Position": Pos{Line: 3-3 Pos: 20-22}
	// |           "Parts":
	// |             [*name.NamePart]
	// |               "Position": Pos{Line: 3-3 Pos: 20-22}
	// |               "Meta":
	// |                 " " before "NodeStart"
	// |               "Value": "Foo"
	// |       "Stmts":
	// |         [*stmt.Class]
	// |           "Position": Pos{Line: 4-10 Pos: 29-139}
	// |           "NamespacedName": "Foo\\Bar"
	// |           "Meta":
	// |             "\n\t\t\t" before "NodeStart"
	// |             " " before "OpenCurlyBracesToken"
	// |             "\n\t\t\t" before "CloseCurlyBracesToken"
	// |           "PhpDocComment": ""
	// |           "ClassName":
	// |             [*node.Identifier]
	// |               "Position": Pos{Line: 4-4 Pos: 35-37}
	// |               "Meta":
	// |                 " " before "NodeStart"
	// |               "Value": "Bar"
	// |           "Stmts":
	// |             [*stmt.ClassMethod]
	// |               "Position": Pos{Line: 5-9 Pos: 45-134}
	// |               "Meta":
	// |                 "\n\t\t\t\t" before "NodeStart"
	// |                 " " before "FunctionToken"
	// |               "ReturnsRef": false
	// |               "PhpDocComment": ""
	// |               "MethodName":
	// |                 [*node.Identifier]
	// |                   "Position": Pos{Line: 5-5 Pos: 61-72}
	// |                   "Meta":
	// |                     " " before "NodeStart"
	// |                   "Value": "FunctionName"
	// |               "Modifiers":
	// |                 [*node.Identifier]
	// |                   "Position": Pos{Line: 5-5 Pos: 45-50}
	// |                   "Value": "public"
	// |               "Params":
	// |                 [*node.Parameter]
	// |                   "Position": Pos{Line: 5-5 Pos: 74-89}
	// |                   "Meta":
	// |                     " " before "EqualToken"
	// |                   "ByRef": false
	// |                   "Variadic": false
	// |                   "VariableType":
	// |                     [*name.Name]
	// |                       "Position": Pos{Line: 5-5 Pos: 74-77}
	// |                       "NamespacedName": "Foo\\Type"
	// |                       "Parts":
	// |                         [*name.NamePart]
	// |                           "Position": Pos{Line: 5-5 Pos: 74-77}
	// |                           "Value": "Type"
	// |                   "Variable":
	// |                     [*expr.Variable]
	// |                       "Position": Pos{Line: 5-5 Pos: 79-82}
	// |                       "Meta":
	// |                         " " before "NodeStart"
	// |                         "$" before "NodeStart"
	// |                       "VarName":
	// |                         [*node.Identifier]
	// |                           "Position": Pos{Line: 5-5 Pos: 79-82}
	// |                           "Value": "var"
	// |                   "DefaultValue":
	// |                     [*expr.ConstFetch]
	// |                       "Position": Pos{Line: 5-5 Pos: 86-89}
	// |                       "Meta":
	// |                         " " before "NodeStart"
	// |                       "Constant":
	// |                         [*name.Name]
	// |                           "Position": Pos{Line: 5-5 Pos: 86-89}
	// |                           "NamespacedName": "null"
	// |                           "Parts":
	// |                             [*name.NamePart]
	// |                               "Position": Pos{Line: 5-5 Pos: 86-89}
	// |                               "Value": "null"
	// |               "Stmt":
	// |                 [*stmt.StmtList]
	// |                   "Position": Pos{Line: 6-9 Pos: 96-134}
	// |                   "Meta":
	// |                     "\n\t\t\t\t" before "NodeStart"
	// |                     "\n\t\t\t\t" before "CloseCurlyBracesToken"
	// |                   "Stmts":
	// |                     [*stmt.Expression]
	// |                       "Position": Pos{Line: 8-8 Pos: 124-128}
	// |                       "Meta":
	// |                         "\n\t\t\t\t\t" before "NodeStart"
	// |                         "// some comment\n" before "NodeStart"
	// |                         "\t\t\t\t\t" before "NodeStart"
	// |                         ";" before "SemiColonToken"
	// |                       "Expr":
	// |                         [*expr.Variable]
	// |                           "Position": Pos{Line: 8-8 Pos: 124-127}
	// |                           "Meta":
	// |                             "$" before "NodeStart"
	// |                           "VarName":
	// |                             [*node.Identifier]
	// |                               "Position": Pos{Line: 8-8 Pos: 124-127}
	// |                               "Value": "var"
}
