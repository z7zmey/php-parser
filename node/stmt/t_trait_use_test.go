package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTraitUse(t *testing.T) {
	src := `<? class Foo { use Bar; }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    25,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    25,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    23,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
										Value: "Bar",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.Nop{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  22,
								EndPos:    23,
							},
						},
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

func TestTraitsUse(t *testing.T) {
	src := `<? class Foo { use Bar, Baz; }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    30,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    30,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    28,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    27,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  24,
											EndPos:    27,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.Nop{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  27,
								EndPos:    28,
							},
						},
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

func TestTraitsUseEmptyAdaptations(t *testing.T) {
	src := `<? class Foo { use Bar, Baz {} }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    32,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    32,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    30,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    27,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  24,
											EndPos:    27,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  28,
								EndPos:    30,
							},
						},
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

func TestTraitsUseModifier(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { one as public; } }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    48,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    48,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    46,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    27,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  24,
											EndPos:    27,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  28,
								EndPos:    46,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  30,
										EndPos:    43,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    33,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    33,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  37,
											EndPos:    43,
										},
										Value: "public",
									},
								},
							},
						},
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

func TestTraitsUseAliasModifier(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { one as public two; } }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    52,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    52,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    50,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    27,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  24,
											EndPos:    27,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  28,
								EndPos:    50,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  30,
										EndPos:    47,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    33,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    33,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  37,
											EndPos:    43,
										},
										Value: "public",
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  44,
											EndPos:    47,
										},
										Value: "two",
									},
								},
							},
						},
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

func TestTraitsUseAdaptions(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { Bar::one insteadof Baz, Quux; Baz::one as two; } }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    80,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    80,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    78,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  19,
									EndPos:    22,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  19,
											EndPos:    22,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  24,
									EndPos:    27,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  24,
											EndPos:    27,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  28,
								EndPos:    78,
							},
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  30,
										EndPos:    58,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  30,
											EndPos:    38,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  30,
												EndPos:    33,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  30,
														EndPos:    33,
													},
													Value: "Bar",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  35,
												EndPos:    38,
											},
											Value: "one",
										},
									},
									Insteadof: []node.Node{
										&name.Name{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  49,
												EndPos:    52,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  49,
														EndPos:    52,
													},
													Value: "Baz",
												},
											},
										},
										&name.Name{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  54,
												EndPos:    58,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  54,
														EndPos:    58,
													},
													Value: "Quux",
												},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  60,
										EndPos:    75,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  60,
											EndPos:    68,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  60,
												EndPos:    63,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  60,
														EndPos:    63,
													},
													Value: "Baz",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  65,
												EndPos:    68,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  72,
											EndPos:    75,
										},
										Value: "two",
									},
								},
							},
						},
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
