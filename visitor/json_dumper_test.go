package visitor_test

import (
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
					$var1;
					$var2;
				}
			}
		}`

	php7parser := php7.NewParser([]byte(src))
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
	// {"type":"*node.Root","position":{"startPos":9,"endPos":179,"startLine":3,"endLine":13},"Stmts":[{"type":"*stmt.Namespace","position":{"startPos":9,"endPos":179,"startLine":3,"endLine":13},"freefloating":{"Start": [{"type":"freefloating.TokenType","value":"<?php"},{"type":"freefloating.WhiteSpaceType","value":"\n\n\t\t"}],"Stmts": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t"}]},"NamespaceName":{"type":"*name.Name","position":{"startPos":19,"endPos":22,"startLine":3,"endLine":3},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":" "}],"End": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"Parts":[{"type":"*name.NamePart","position":{"startPos":19,"endPos":22,"startLine":3,"endLine":3},"Value":"Foo"}]},"Stmts":[{"type":"*stmt.Class","position":{"startPos":28,"endPos":175,"startLine":4,"endLine":12},"namespacedName":"Foo\\Bar","freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t"}],"Name": [{"type":"freefloating.WhiteSpaceType","value":" "}],"Stmts": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t"}]},"PhpDocComment":"","ClassName":{"type":"*node.Identifier","position":{"startPos":34,"endPos":37,"startLine":4,"endLine":4},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"Value":"Bar"},"Stmts":[{"type":"*stmt.ClassMethod","position":{"startPos":44,"endPos":170,"startLine":5,"endLine":11},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t"}],"Function": [{"type":"freefloating.WhiteSpaceType","value":" "}],"ModifierList": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"PhpDocComment":"","ReturnsRef":false,"MethodName":{"type":"*node.Identifier","position":{"startPos":60,"endPos":72,"startLine":5,"endLine":5},"Value":"FunctionName"},"Modifiers":[{"type":"*node.Identifier","position":{"startPos":44,"endPos":50,"startLine":5,"endLine":5},"Value":"public"}],"Params":[{"type":"*node.Parameter","position":{"startPos":73,"endPos":89,"startLine":5,"endLine":5},"freefloating":{"Var": [{"type":"freefloating.WhiteSpaceType","value":" "}],"OptionalType": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"ByRef":false,"Variadic":false,"VariableType":{"type":"*name.Name","position":{"startPos":73,"endPos":77,"startLine":5,"endLine":5},"namespacedName":"Foo\\Type","Parts":[{"type":"*name.NamePart","position":{"startPos":73,"endPos":77,"startLine":5,"endLine":5},"Value":"Type"}]},"Variable":{"type":"*expr.Variable","position":{"startPos":78,"endPos":82,"startLine":5,"endLine":5},"freefloating":{"Dollar": [{"type":"freefloating.TokenType","value":"$"}]},"VarName":{"type":"*node.Identifier","position":{"startPos":78,"endPos":82,"startLine":5,"endLine":5},"Value":"var"}},"DefaultValue":{"type":"*expr.ConstFetch","position":{"startPos":85,"endPos":89,"startLine":5,"endLine":5},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":" "}]},"Constant":{"type":"*name.Name","position":{"startPos":85,"endPos":89,"startLine":5,"endLine":5},"namespacedName":"null","Parts":[{"type":"*name.NamePart","position":{"startPos":85,"endPos":89,"startLine":5,"endLine":5},"Value":"null"}]}}}],"Stmt":{"type":"*stmt.StmtList","position":{"startPos":95,"endPos":170,"startLine":6,"endLine":11},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t"}],"Stmts": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t"}]},"Stmts":[{"type":"*stmt.Expression","position":{"startPos":146,"endPos":152,"startLine":9,"endLine":9},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t\t"},{"type":"freefloating.CommentType","value":"// some comment\n"},{"type":"freefloating.WhiteSpaceType","value":"\t\t\t\t\t"},{"type":"freefloating.CommentType","value":"// second comment\n"},{"type":"freefloating.WhiteSpaceType","value":"\t\t\t\t\t"}],"SemiColon": [{"type":"freefloating.TokenType","value":";"}]},"Expr":{"type":"*expr.Variable","position":{"startPos":146,"endPos":151,"startLine":9,"endLine":9},"freefloating":{"Dollar": [{"type":"freefloating.TokenType","value":"$"}]},"VarName":{"type":"*node.Identifier","position":{"startPos":146,"endPos":151,"startLine":9,"endLine":9},"Value":"var1"}}},{"type":"*stmt.Expression","position":{"startPos":158,"endPos":164,"startLine":10,"endLine":10},"freefloating":{"Start": [{"type":"freefloating.WhiteSpaceType","value":"\n\t\t\t\t\t"}],"SemiColon": [{"type":"freefloating.TokenType","value":";"}]},"Expr":{"type":"*expr.Variable","position":{"startPos":158,"endPos":163,"startLine":10,"endLine":10},"freefloating":{"Dollar": [{"type":"freefloating.TokenType","value":"$"}]},"VarName":{"type":"*node.Identifier","position":{"startPos":158,"endPos":163,"startLine":10,"endLine":10},"Value":"var2"}}}]}}]}]}]}
}
