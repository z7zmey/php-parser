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

func TestSimpleClass(t *testing.T) {
	src := `<? class foo{ }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    15,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    15,
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
				Stmts: []node.Node{},
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

func TestAbstractClass(t *testing.T) {
	src := `<? abstract class foo{ }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    24,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    24,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  18,
						EndPos:    21,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    11,
						},
						Value: "abstract",
					},
				},
				Stmts: []node.Node{},
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

func TestClassExtends(t *testing.T) {
	src := `<? final class foo extends bar { }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    34,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    34,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  15,
						EndPos:    18,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
						Value: "final",
					},
				},
				Extends: &stmt.ClassExtends{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  19,
						EndPos:    30,
					},
					ClassName: &name.Name{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  27,
							EndPos:    30,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  27,
									EndPos:    30,
								},
								Value: "bar",
							},
						},
					},
				},
				Stmts: []node.Node{},
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

func TestClassImplement(t *testing.T) {
	src := `<? final class foo implements bar { }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    37,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    37,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  15,
						EndPos:    18,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  19,
						EndPos:    33,
					},
					InterfaceNames: []node.Node{
						&name.Name{
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
									Value: "bar",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
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

func TestClassImplements(t *testing.T) {
	src := `<? final class foo implements bar, baz { }`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    42,
		},
		Stmts: []node.Node{
			&stmt.Class{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    42,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  15,
						EndPos:    18,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    8,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  19,
						EndPos:    38,
					},
					InterfaceNames: []node.Node{
						&name.Name{
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
									Value: "bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  35,
								EndPos:    38,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  35,
										EndPos:    38,
									},
									Value: "baz",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
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

func TestAnonimousClass(t *testing.T) {
	src := `<? new class() extends foo implements bar, baz { };`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    51,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    51,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    50,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    50,
						},
						PhpDocComment: "",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
						},
						Extends: &stmt.ClassExtends{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    26,
							},
							ClassName: &name.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  23,
									EndPos:    26,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  23,
											EndPos:    26,
										},
										Value: "foo",
									},
								},
							},
						},
						Implements: &stmt.ClassImplements{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  27,
								EndPos:    46,
							},
							InterfaceNames: []node.Node{
								&name.Name{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  38,
										EndPos:    41,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  38,
												EndPos:    41,
											},
											Value: "bar",
										},
									},
								},
								&name.Name{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  43,
										EndPos:    46,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  43,
												EndPos:    46,
											},
											Value: "baz",
										},
									},
								},
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
