package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestShellExec(t *testing.T) {
	src := "<? `cmd $a`;"

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    12,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    12,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  4,
								EndPos:    8,
							},
							Value: "cmd ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
								Value: "a",
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
