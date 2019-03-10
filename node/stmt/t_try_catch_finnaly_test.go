package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTry(t *testing.T) {
	src := `<? 
		try {}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   -1,
			StartPos:  6,
			EndPos:    -1,
		},
		Stmts: []node.Node{
			&stmt.Try{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   -1,
					StartPos:  6,
					EndPos:    -1,
				},
				Stmts:   []node.Node{},
				Catches: []node.Node{},
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

func TestTryCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  6,
			EndPos:    36,
		},
		Stmts: []node.Node{
			&stmt.Try{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    36,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    36,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  20,
									EndPos:    29,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  20,
											EndPos:    29,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  30,
								EndPos:    32,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    32,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
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

func TestPhp7TryCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception|RuntimeException $e) {}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  6,
			EndPos:    53,
		},
		Stmts: []node.Node{
			&stmt.Try{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    53,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    53,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  20,
									EndPos:    29,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  20,
											EndPos:    29,
										},
										Value: "Exception",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    46,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  30,
											EndPos:    46,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  47,
								EndPos:    49,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  47,
									EndPos:    49,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTryCatchCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {} catch (RuntimeException $e) {}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  6,
			EndPos:    67,
		},
		Stmts: []node.Node{
			&stmt.Try{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    67,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    36,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  20,
									EndPos:    29,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  20,
											EndPos:    29,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  30,
								EndPos:    32,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    32,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  37,
							EndPos:    67,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  44,
									EndPos:    60,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  44,
											EndPos:    60,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  61,
								EndPos:    63,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  61,
									EndPos:    63,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
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

func TestTryCatchFinally(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {} finally {}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  6,
			EndPos:    47,
		},
		Stmts: []node.Node{
			&stmt.Try{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    47,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    36,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  20,
									EndPos:    29,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  20,
											EndPos:    29,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  30,
								EndPos:    32,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  30,
									EndPos:    32,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  37,
						EndPos:    47,
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

func TestTryCatchCatchCatch(t *testing.T) {
	src := `<? try {} catch (Exception $e) {} catch (\RuntimeException $e) {} catch (namespace\AdditionException $e) {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    107,
		},
		Stmts: []node.Node{
			&stmt.Try{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    107,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    33,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  17,
									EndPos:    26,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  17,
											EndPos:    26,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  27,
								EndPos:    29,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  27,
									EndPos:    29,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  34,
							EndPos:    65,
						},
						Types: []node.Node{
							&name.FullyQualified{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  41,
									EndPos:    58,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  42,
											EndPos:    58,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  59,
								EndPos:    61,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  59,
									EndPos:    61,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  66,
							EndPos:    107,
						},
						Types: []node.Node{
							&name.Relative{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  73,
									EndPos:    100,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  83,
											EndPos:    100,
										},
										Value: "AdditionException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  101,
								EndPos:    103,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  101,
									EndPos:    103,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
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
