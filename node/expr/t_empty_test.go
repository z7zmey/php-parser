package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestEmpty(t *testing.T) {
	src := `<? empty($a);`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    13,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    13,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    11,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
							Value: "a",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src))
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
