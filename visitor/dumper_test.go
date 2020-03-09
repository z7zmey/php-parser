package visitor_test

import (
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

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.WithFreeFloating()
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
	// |   "Position": Pos{Line: 3-11 Pos: 9-144}
	// |   "Stmts":
	// |     [*stmt.Namespace]
	// |       "Position": Pos{Line: 3-11 Pos: 9-144}
	// |       "freefloating":
	// |         "Start": "<?php"
	// |         "Start": "\n\n\t\t"
	// |         "Stmts": "\n\t\t"
	// |       "NamespaceName":
	// |         [*name.Name]
	// |           "Position": Pos{Line: 3-3 Pos: 19-22}
	// |           "freefloating":
	// |             "Start": " "
	// |             "End": " "
	// |           "Parts":
	// |             [*name.NamePart]
	// |               "Position": Pos{Line: 3-3 Pos: 19-22}
	// |               "Value": "Foo"
	// |       "Stmts":
	// |         [*stmt.Class]
	// |           "Position": Pos{Line: 4-10 Pos: 28-140}
	// |           "NamespacedName": "Foo\\Bar"
	// |           "freefloating":
	// |             "Start": "\n\t\t\t"
	// |             "Name": " "
	// |             "Stmts": "\n\t\t\t"
	// |           "PhpDocComment": ""
	// |           "ClassName":
	// |             [*node.Identifier]
	// |               "Position": Pos{Line: 4-4 Pos: 34-37}
	// |               "freefloating":
	// |                 "Start": " "
	// |               "Value": "Bar"
	// |           "Stmts":
	// |             [*stmt.ClassMethod]
	// |               "Position": Pos{Line: 5-9 Pos: 45-135}
	// |               "freefloating":
	// |                 "Start": " \n\t\t\t\t"
	// |                 "ModifierList": " "
	// |                 "Function": " "
	// |               "ReturnsRef": false
	// |               "PhpDocComment": ""
	// |               "MethodName":
	// |                 [*node.Identifier]
	// |                   "Position": Pos{Line: 5-5 Pos: 61-73}
	// |                   "Value": "FunctionName"
	// |               "Modifiers":
	// |                 [*node.Identifier]
	// |                   "Position": Pos{Line: 5-5 Pos: 45-51}
	// |                   "Value": "public"
	// |               "Params":
	// |                 [*node.Parameter]
	// |                   "Position": Pos{Line: 5-5 Pos: 74-90}
	// |                   "freefloating":
	// |                     "OptionalType": " "
	// |                     "Var": " "
	// |                   "Variadic": false
	// |                   "ByRef": false
	// |                   "VariableType":
	// |                     [*name.Name]
	// |                       "Position": Pos{Line: 5-5 Pos: 74-78}
	// |                       "NamespacedName": "Foo\\Type"
	// |                       "Parts":
	// |                         [*name.NamePart]
	// |                           "Position": Pos{Line: 5-5 Pos: 74-78}
	// |                           "Value": "Type"
	// |                   "Variable":
	// |                     [*expr.Variable]
	// |                       "Position": Pos{Line: 5-5 Pos: 79-83}
	// |                       "freefloating":
	// |                         "Dollar": "$"
	// |                       "VarName":
	// |                         [*node.Identifier]
	// |                           "Position": Pos{Line: 5-5 Pos: 79-83}
	// |                           "Value": "var"
	// |                   "DefaultValue":
	// |                     [*expr.ConstFetch]
	// |                       "Position": Pos{Line: 5-5 Pos: 86-90}
	// |                       "freefloating":
	// |                         "Start": " "
	// |                       "Constant":
	// |                         [*name.Name]
	// |                           "Position": Pos{Line: 5-5 Pos: 86-90}
	// |                           "NamespacedName": "null"
	// |                           "Parts":
	// |                             [*name.NamePart]
	// |                               "Position": Pos{Line: 5-5 Pos: 86-90}
	// |                               "Value": "null"
	// |               "Stmt":
	// |                 [*stmt.StmtList]
	// |                   "Position": Pos{Line: 6-9 Pos: 96-135}
	// |                   "freefloating":
	// |                     "Start": "\n\t\t\t\t"
	// |                     "Stmts": "\n\t\t\t\t"
	// |                   "Stmts":
	// |                     [*stmt.Expression]
	// |                       "Position": Pos{Line: 8-8 Pos: 124-129}
	// |                       "freefloating":
	// |                         "SemiColon": ";"
	// |                         "Start": "\n\t\t\t\t\t"
	// |                         "Start": "// some comment\n"
	// |                         "Start": "\t\t\t\t\t"
	// |                       "Expr":
	// |                         [*expr.Variable]
	// |                           "Position": Pos{Line: 8-8 Pos: 124-128}
	// |                           "freefloating":
	// |                             "Dollar": "$"
	// |                           "VarName":
	// |                             [*node.Identifier]
	// |                               "Position": Pos{Line: 8-8 Pos: 124-128}
	// |                               "Value": "var"
}
