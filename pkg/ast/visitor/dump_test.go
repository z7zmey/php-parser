package visitor_test

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/ast/traverser"
	"github.com/z7zmey/php-parser/pkg/ast/visitor"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
	"os"
)

func ExampleDump() {
	stxTree := &ast.Root{
		Node: ast.Node{
			Tokens: token.Collection{
				token.Start: []token.Token{
					{
						ID:    token.T_WHITESPACE,
						Value: []byte(" "),
					},
				},
			},
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    1,
			},
		},
		Stmts: []ast.Vertex{
			&ast.Identifier{},
			&ast.Parameter{
				Variadic: true,
				Var:      &ast.ExprVariable{},
			},
			&ast.StmtInlineHtml{
				Value: []byte("foo"),
			},
		},
	}

	traverser.NewDFS(visitor.NewDump(os.Stdout)).Traverse(stxTree)

	//output:
	//&ast.Root{
	//	Node: ast.Node{
	//		Tokens: token.Collection{
	//			token.Start: []token.Token{
	//				{
	//					ID:    token.T_WHITESPACE,
	//					Value: []byte(" "),
	//				},
	//			},
	//		},
	//		Position: &position.Position{
	//			StartLine: 1,
	//			EndLine:   1,
	//			StartPos:  0,
	//			EndPos:    1,
	//		},
	//	},
	//	Stmts: []ast.Vertex{
	//		&ast.Identifier{
	//			Value: []byte(""),
	//		},
	//		&ast.Parameter{
	//			Variadic: true,
	//			Var: &ast.ExprVariable{
	//			},
	//		},
	//		&ast.StmtInlineHtml{
	//			Value: []byte("foo"),
	//		},
	//	},
	//}
}
