package scalar_test

import (
	"bytes"
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
			StartPos:  7,
			EndPos:    24,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  7,
					EndPos:    24,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  7,
						EndPos:    23,
					},
					Label: "LBL",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  11,
								EndPos:    15,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  16,
								EndPos:    19,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  16,
									EndPos:    19,
								},
								Value: "var",
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
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
			StartPos:  7,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  7,
					EndPos:    26,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  7,
						EndPos:    25,
					},
					Label: "\"LBL\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    17,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  18,
								EndPos:    21,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  18,
									EndPos:    21,
								},
								Value: "var",
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
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
			StartPos:  7,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  7,
					EndPos:    26,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  7,
						EndPos:    25,
					},
					Label: "'LBL'",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    21,
							},
							Value: "test $var",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
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
			StartPos:  7,
			EndPos:    14,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   2,
					StartPos:  7,
					EndPos:    14,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   2,
						StartPos:  7,
						EndPos:    13,
					},
					Label: "CAD",
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
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
			StartPos:  7,
			EndPos:    21,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   3,
					StartPos:  7,
					EndPos:    21,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   3,
						StartPos:  7,
						EndPos:    20,
					},
					Label: "CAD",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  11,
								EndPos:    16,
							},
							Value: "\thello",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
