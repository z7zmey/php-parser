package expr_test

import (
	"bytes"
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

func TestArray(t *testing.T) {
	src := `<? array();`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    11,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    11,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    10,
					},
					Items: []node.Node{},
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

func TestArrayItem(t *testing.T) {
	src := `<? array(1);`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    12,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    12,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    11,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    10,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    10,
								},
								Value: "1",
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

func TestArrayItems(t *testing.T) {
	src := `<? array(1=>1, &$b,);`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    21,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    21,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    20,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    13,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    10,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  13,
									EndPos:    13,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    18,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  17,
											EndPos:    18,
										},
										Value: "b",
									},
								},
							},
						},
						&expr.ArrayItem{},
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
