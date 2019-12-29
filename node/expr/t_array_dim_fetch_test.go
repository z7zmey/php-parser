package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
)

func TestArrayDimFetch(t *testing.T) {
	src := `<? $a[1];`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    9,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    9,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    8,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
							Value: "a",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    7,
						},
						Value: "1",
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

func TestArrayDimFetchNested(t *testing.T) {
	src := `<? $a[1][2];`

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
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
								Value: "a",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    7,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    10,
						},
						Value: "2",
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
