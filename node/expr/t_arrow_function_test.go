package expr_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
	"gotest.tools/assert"
)

func TestArrowFunction(t *testing.T) {
	src := `<? fn() => $a;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    14,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    14,
				},
				Expr: &expr.ArrowFunction{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    13,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    13,
							},
							Value: "a",
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
}

func TestArrowFunctionReturnType(t *testing.T) {
	src := `<? fn & () : foo => $a;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    23,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    23,
				},
				Expr: &expr.ArrowFunction{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    22,
					},
					Static:        false,
					PhpDocComment: "",
					ReturnsRef:    true,
					ReturnType: &name.Name{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    16,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    16,
								},
								Value: "foo",
							},
						},
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  20,
							EndPos:    22,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
							Value: "a",
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
}
