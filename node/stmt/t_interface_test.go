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

func TestInterface(t *testing.T) {
	src := `<? interface Foo {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    19,
		},
		Stmts: []node.Node{
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    19,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  13,
						EndPos:    16,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
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

func TestInterfaceExtend(t *testing.T) {
	src := `<? interface Foo extends Bar {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    31,
		},
		Stmts: []node.Node{
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    31,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  13,
						EndPos:    16,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  17,
						EndPos:    28,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  25,
								EndPos:    28,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  25,
										EndPos:    28,
									},
									Value: "Bar",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
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

func TestInterfaceExtends(t *testing.T) {
	src := `<? interface Foo extends Bar, Baz {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    36,
		},
		Stmts: []node.Node{
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    36,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  13,
						EndPos:    16,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  17,
						EndPos:    33,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  25,
								EndPos:    28,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  25,
										EndPos:    28,
									},
									Value: "Bar",
								},
							},
						},
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
									Value: "Baz",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
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
