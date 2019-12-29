package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestProperty(t *testing.T) {
	src := `<? class foo {var $a;}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    22,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    22,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    21,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    17,
								},
								Value: "var",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    20,
										},
										Value: "a",
									},
								},
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

func TestProperties(t *testing.T) {
	src := `<? class foo {public static $a, $b = 1;}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    40,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    40,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    39,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    20,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  21,
									EndPos:    27,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    30,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  28,
										EndPos:    30,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  28,
											EndPos:    30,
										},
										Value: "a",
									},
								},
							},
							&stmt.Property{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  32,
									EndPos:    38,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  32,
										EndPos:    34,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  32,
											EndPos:    34,
										},
										Value: "b",
									},
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  37,
										EndPos:    38,
									},
									Value: "1",
								},
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

func TestProperties2(t *testing.T) {
	src := `<? class foo {public static $a = 1, $b;}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    40,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    40,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    39,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    20,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  21,
									EndPos:    27,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    34,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  28,
										EndPos:    30,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  28,
											EndPos:    30,
										},
										Value: "a",
									},
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
							&stmt.Property{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  36,
									EndPos:    38,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  36,
										EndPos:    38,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  36,
											EndPos:    38,
										},
										Value: "b",
									},
								},
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

func TestPropertyType(t *testing.T) {
	src := `<? class foo {var bar $a;}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    26,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    25,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    17,
								},
								Value: "var",
							},
						},
						Type: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    21,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    21,
									},
									Value: "bar",
								},
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  22,
									EndPos:    24,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  22,
											EndPos:    24,
										},
										Value: "a",
									},
								},
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
}
