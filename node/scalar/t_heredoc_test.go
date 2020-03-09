package scalar_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestHeredocSimpleLabel(t *testing.T) {
	src := `<? <<<LBL
test $var
LBL;
`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   3,
			StartPos:  3,
			EndPos:    24,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  3,
					EndPos:    24,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    23,
					},
					Label: "<<<LBL\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  10,
								EndPos:    15,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  15,
								EndPos:    19,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  15,
									EndPos:    19,
								},
								Value: "var",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  19,
								EndPos:    20,
							},
							Value: "\n",
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

func TestSimpleHeredocLabel(t *testing.T) {
	src := `<? <<<"LBL"
test $var
LBL;
`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   3,
			StartPos:  3,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  3,
					EndPos:    26,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    25,
					},
					Label: "<<<\"LBL\"\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  12,
								EndPos:    17,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  17,
								EndPos:    21,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  17,
									EndPos:    21,
								},
								Value: "var",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  21,
								EndPos:    22,
							},
							Value: "\n",
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

func TestSimpleNowdocLabel(t *testing.T) {
	src := `<? <<<'LBL'
test $var
LBL;
`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   3,
			StartPos:  3,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  3,
					EndPos:    26,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    25,
					},
					Label: "<<<'LBL'\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  12,
								EndPos:    22,
							},
							Value: "test $var\n",
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

func TestEmptyHeredoc(t *testing.T) {
	src := `<? <<<CAD
CAD;
`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   2,
			StartPos:  3,
			EndPos:    14,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   2,
					StartPos:  3,
					EndPos:    14,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   2,
						StartPos:  3,
						EndPos:    13,
					},
					Label: "<<<CAD\n",
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

func TestHeredocScalarString(t *testing.T) {
	src := `<? <<<CAD
	hello
CAD;
`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   3,
			StartPos:  3,
			EndPos:    21,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  3,
					EndPos:    21,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  3,
						EndPos:    20,
					},
					Label: "<<<CAD\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  10,
								EndPos:    17,
							},
							Value: "\thello\n",
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
