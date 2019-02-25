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

func TestSimpleVar(t *testing.T) {
	src := `<? "test $var";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    15,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    15,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    14,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    13,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    13,
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

func TestSimpleVarOneChar(t *testing.T) {
	src := `<? "test $a";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    13,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    13,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    12,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    11,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    11,
								},
								Value: "a",
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

func TestSimpleVarEndsEcapsed(t *testing.T) {
	src := `<? "test $var\"";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    17,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    17,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    16,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    13,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    13,
								},
								Value: "var",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    15,
							},
							Value: "\\\"",
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

func TestStringVarCurveOpen(t *testing.T) {
	src := `<? "=$a{$b}";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    13,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    13,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    12,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    5,
							},
							Value: "=",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    7,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  6,
									EndPos:    7,
								},
								Value: "a",
							},
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    10,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    10,
								},
								Value: "b",
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

func TestSimpleVarPropertyFetch(t *testing.T) {
	src := `<? "test $foo->bar()";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    22,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    22,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    21,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    18,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  10,
									EndPos:    13,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  10,
										EndPos:    13,
									},
									Value: "foo",
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    18,
								},
								Value: "bar",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    20,
							},
							Value: "()",
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

func TestDollarOpenCurlyBraces(t *testing.T) {
	src := `<? "test ${foo}";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    17,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    17,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    16,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    15,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
								Value: "foo",
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

func TestDollarOpenCurlyBracesDimNumber(t *testing.T) {
	src := `<? "test ${foo[0]}";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    20,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    20,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    19,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    18,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    14,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    14,
									},
									Value: "foo",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    16,
								},
								Value: "0",
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

func TestCurlyOpenMethodCall(t *testing.T) {
	src := `<? "test {$foo->bar()}";`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    24,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    24,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    23,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    9,
							},
							Value: "test ",
						},
						&expr.MethodCall{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    21,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  11,
									EndPos:    14,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  11,
										EndPos:    14,
									},
									Value: "foo",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    19,
								},
								Value: "bar",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  20,
									EndPos:    21,
								},
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
