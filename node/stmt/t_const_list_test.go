package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestConstList(t *testing.T) {
	src := `<? const FOO = 1, BAR = 2;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.ConstList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    26,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    16,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    12,
							},
							Value: "FOO",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    16,
							},
							Value: "1",
						},
					},
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    25,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    21,
							},
							Value: "BAR",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    25,
							},
							Value: "2",
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
