package visitor_test

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/ast/traverser"
	"github.com/z7zmey/php-parser/pkg/ast/visitor"
	"os"
)

func ExampleDump() {
	stxTree := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.Identifier{},
			&ast.Parameter{
				Variadic: true,
				Var:      &ast.ExprVariable{},
			},
			&ast.StmtInlineHtml{
				Value: "foo",
			},
		},
	}

	traverser.NewDFS(visitor.NewDump(os.Stdout)).Traverse(stxTree)

	//output:
	//&ast.Root{
	//	Stmts: []ast.Vertex{
	//		&ast.Identifier{
	//			Value: "",
	//		},
	//		&ast.Parameter{
	//			Variadic: true,
	//			Var: &ast.ExprVariable{
	//			},
	//		},
	//		&ast.StmtInlineHtml{
	//			Value: "foo",
	//		},
	//	},
	//}
}
