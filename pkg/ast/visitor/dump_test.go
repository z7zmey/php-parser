package visitor_test

import (
	"bytes"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
	"testing"

	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/ast/visitor"
)

func TestDumper_root(t *testing.T) {
	o := bytes.NewBufferString("")

	p := visitor.NewDump(o).WithTokens().WithPositions()
	n := &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   2,
			StartPos:  3,
			EndPos:    4,
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		EndTkn: &token.Token{
			FreeFloating: []*token.Token{
				{
					ID:    token.T_WHITESPACE,
					Value: []byte(" "),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   2,
						StartPos:  3,
						EndPos:    4,
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `&ast.Root{
	Position: &position.Position{
		StartLine: 1,
		EndLine:   2,
		StartPos:  3,
		EndPos:    4,
	},
	Stmts: []ast.Vertex{
		&ast.StmtNop{
		},
	},
	EndTkn: &token.Token{
		FreeFloating: []*token.Token{
			{
				ID: token.T_WHITESPACE,
				Value: []byte(" "),
				Position: &position.Position{
					StartLine: 1,
					EndLine:   2,
					StartPos:  3,
					EndPos:    4,
				},
			},
		},
	},
},
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
