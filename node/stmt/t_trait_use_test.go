package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTraitUse(t *testing.T) {
	src := `<? class Foo { use Bar; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
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
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestTraitsUse(t *testing.T) {
	src := `<? class Foo { use Bar, Baz; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
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
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestTraitsUseEmptyAdaptations(t *testing.T) {
	src := `<? class Foo { use Bar, Baz {} }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestTraitsUseModifier(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { one as public; } }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Ref: &stmt.TraitMethodRef{
										Method: &node.Identifier{Value: "one"},
									},
									Modifier: &node.Identifier{Value: "public"},
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
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestTraitsUseAliasModifier(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { one as public two; } }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Ref: &stmt.TraitMethodRef{
										Method: &node.Identifier{Value: "one"},
									},
									Modifier: &node.Identifier{Value: "public"},
									Alias:    &node.Identifier{Value: "two"},
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
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestTraitsUseAdaptions(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { Bar::one insteadof Baz, Quux; Baz::one as two; } }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Ref: &stmt.TraitMethodRef{
										Trait: &name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Bar"},
											},
										},
										Method: &node.Identifier{Value: "one"},
									},
									Insteadof: []node.Node{
										&name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Baz"},
											},
										},
										&name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Quux"},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Ref: &stmt.TraitMethodRef{
										Trait: &name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Baz"},
											},
										},
										Method: &node.Identifier{Value: "one"},
									},
									Alias: &node.Identifier{Value: "two"},
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
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
