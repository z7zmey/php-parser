package visitor_test

import (
	"bytes"
	"os"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func ExampleJsonDumper() {
	src := `<?php

		namespace Foo {
			class Bar {
				public function FunctionName(Type $var = null)
				{
					// some comment
					// second comment
					$var;
				}
			}
		}`

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.WithFreeFloating()
	php7parser.Parse()
	nodes := php7parser.GetRootNode()

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := &visitor.JsonDumper{
		Writer:     os.Stdout,
		NsResolver: nsResolver,
	}
	nodes.Walk(dumper)

	// Output:
	// {"type":"*node.Root","position":{"startPos":10,"endPos":166,"startLine":3,"endLine":12},"Stmts":[{"type":"*stmt.Namespace","position":{"startPos":10,"endPos":166,"startLine":3,"endLine":12},"freefloating":{"Start": [{"type":"freefloating.TokenType","value":"<?php"},{"type":"freefloating.WhiteSpaceType","value":"\n\n\t\t"}],"Stmts": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t"}]},"NamespaceName":{"type":"*name.Name","position":{"startPos":20,"endPos":22,"startLine":3,"endLine":3},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":" "}],"End": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"Parts":[{"type":"*name.NamePart","position":{"startPos":20,"endPos":22,"startLine":3,"endLine":3},"Value":"Foo"}]},"Stmts":[{"type":"*stmt.Class","position":{"startPos":29,"endPos":162,"startLine":4,"endLine":11},"namespacedName":"Foo\\Bar","freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t"}],"Name": [{"type":"freefloating.WhiteSpaceType","value":" "}],"Stmts": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t"}]},"PhpDocComment":"","ClassName":{"type":"*node.Identifier","position":{"startPos":35,"endPos":37,"startLine":4,"endLine":4},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"Value":"Bar"},"Stmts":[{"type":"*stmt.ClassMethod","position":{"startPos":45,"endPos":157,"startLine":5,"endLine":10},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t"}],"Function": [{"type":"freefloating.WhiteSpaceType","value":" "}],"ModifierList": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"PhpDocComment":"","ReturnsRef":false,"MethodName":{"type":"*node.Identifier","position":{"startPos":61,"endPos":72,"startLine":5,"endLine":5},"Value":"FunctionName"},"Modifiers":[{"type":"*node.Identifier","position":{"startPos":45,"endPos":50,"startLine":5,"endLine":5},"Value":"public"}],"Params":[{"type":"*node.Parameter","position":{"startPos":74,"endPos":89,"startLine":5,"endLine":5},"freefloating":{"Var": [{"type":"freefloating.WhiteSpaceType","value":" "}],"OptionalType": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"ByRef":false,"Variadic":false,"VariableType":{"type":"*name.Name","position":{"startPos":74,"endPos":77,"startLine":5,"endLine":5},"namespacedName":"Foo\\Type","Parts":[{"type":"*name.NamePart","position":{"startPos":74,"endPos":77,"startLine":5,"endLine":5},"Value":"Type"}]},"Variable":{"type":"*expr.Variable","position":{"startPos":79,"endPos":82,"startLine":5,"endLine":5},"freefloating":{"Dollar": [{"type":"freefloating.TokenType","value":"$"}]},"VarName":{"type":"*node.Identifier","position":{"startPos":79,"endPos":82,"startLine":5,"endLine":5},"Value":"var"}},"DefaultValue":{"type":"*expr.ConstFetch","position":{"startPos":86,"endPos":89,"startLine":5,"endLine":5},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"Constant":{"type":"*name.Name","position":{"startPos":86,"endPos":89,"startLine":5,"endLine":5},"namespacedName":"null","Parts":[{"type":"*name.NamePart","position":{"startPos":86,"endPos":89,"startLine":5,"endLine":5},"Value":"null"}]}}}],"Stmt":{"type":"*stmt.StmtList","position":{"startPos":96,"endPos":157,"startLine":6,"endLine":10},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t"}],"Stmts": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t"}]},"Stmts":[{"type":"*stmt.Expression","position":{"startPos":147,"endPos":151,"startLine":9,"endLine":9},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t\t"},{"type":"freefloating.CommentType","value":"// some comment\n"},{"type":"freefloating.WhiteSpaceType","value":"\t\t\t\t\t"},{"type":"freefloating.CommentType","value":"// second comment\n"},{"type":"freefloating.WhiteSpaceType","value":"\t\t\t\t\t"}],"SemiColon": [{"type":"freefloating.TokenType","value":";"}]},"Expr":{"type":"*expr.Variable","position":{"startPos":147,"endPos":150,"startLine":9,"endLine":9},"freefloating":{"Dollar": [{"type":"freefloating.TokenType","value":"$"}]},"VarName":{"type":"*node.Identifier","position":{"startPos":147,"endPos":150,"startLine":9,"endLine":9},"Value":"var"}}}]}}]}]}]}
}
