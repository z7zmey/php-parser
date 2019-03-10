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

func TestDeclare(t *testing.T) {
	src := `<? declare(ticks=1);`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    20,
		},
		Stmts: []node.Node{
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    20,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    18,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    16,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    18,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.Nop{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  19,
						EndPos:    20,
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

func TestDeclareStmts(t *testing.T) {
	src := `<? declare(ticks=1, strict_types=1) {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    38,
		},
		Stmts: []node.Node{
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    38,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    18,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    16,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    18,
							},
							Value: "1",
						},
					},
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  20,
							EndPos:    34,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    32,
							},
							Value: "strict_types",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  33,
								EndPos:    34,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  36,
						EndPos:    38,
					},
					Stmts: []node.Node{},
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

func TestAltDeclare(t *testing.T) {
	src := `<? declare(ticks=1): enddeclare;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    32,
		},
		Stmts: []node.Node{
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    32,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    18,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    16,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    18,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
				Alt: true,
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
