package php7_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/internal/parser/php7"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/internal/tree"
	"github.com/z7zmey/php-parser/pkg/ast"
)

func TestHaltCompiller(t *testing.T) {
	src := `<? __halt_compiler();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 21, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtHaltCompiler,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{3, 21, 1, 1},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{0, 2, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{2, 3, 1, 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{20, 21, 1, 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNamespaceName(t *testing.T) {
	src := `<? namespace Foo\Bar;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 21, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtNamespace,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{3, 21, 1, 1},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{0, 2, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{2, 3, 1, 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{20, 21, 1, 1},
					},
				},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeNameName,
						Group:    ast.NodeGroupNamespaceName,
						Position: ast.Position{13, 20, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{13, 16, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{12, 13, 1, 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{17, 20, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_NS_SEPARATOR,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{16, 17, 1, 1},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNamespaceNameGroup(t *testing.T) {
	src := `<? namespace Foo\Bar {}`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 23, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtNamespace,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{3, 23, 1, 1},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{0, 2, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{2, 3, 1, 1},
					},
				},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeNameName,
						Group:    ast.NodeGroupNamespaceName,
						Position: ast.Position{13, 20, 1, 1},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupEnd,
								Position: ast.Position{20, 21, 1, 1},
							},
						},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{13, 16, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{12, 13, 1, 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{17, 20, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_NS_SEPARATOR,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{16, 17, 1, 1},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNamespaceGroup(t *testing.T) {
	src := `<? namespace {}`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtNamespace,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{3, 15, 1, 1},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{0, 2, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{2, 3, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupNamespace,
						Position: ast.Position{12, 13, 1, 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMixedGroupUse(t *testing.T) {
	src := `<? use \foo\bar\{function foo, bar, const baz,};`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 48, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtGroupUse,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{3, 48, 1, 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeNameName,
						Group:    ast.NodeGroupPrefix,
						Position: ast.Position{8, 15, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{8, 11, 1, 1},
							},
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{12, 15, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_NS_SEPARATOR,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{11, 12, 1, 1},
									},
								},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{17, 29, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupUseType,
								Position: ast.Position{17, 25, 1, 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{26, 29, 1, 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{26, 29, 1, 1},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Position: ast.Position{25, 26, 1, 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{31, 34, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{31, 34, 1, 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{31, 34, 1, 1},
										Tokens: []ast.Token{
											{
												Type:     ',',
												Group:    ast.TokenGroupStart,
												Position: ast.Position{29, 30, 1, 1},
											},
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupStart,
												Position: ast.Position{30, 31, 1, 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{36, 45, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupUseType,
								Position: ast.Position{36, 41, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     ',',
										Group:    ast.TokenGroupStart,
										Position: ast.Position{34, 35, 1, 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{35, 36, 1, 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{42, 45, 1, 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{42, 45, 1, 1},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupStart,
												Position: ast.Position{41, 42, 1, 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupUse,
						Position: ast.Position{6, 7, 1, 1},
					},
					{
						Type:     scanner.T_NS_SEPARATOR,
						Group:    ast.TokenGroupUse,
						Position: ast.Position{7, 8, 1, 1},
					},
					{
						Type:     ',',
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{45, 46, 1, 1},
					},
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{0, 2, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{2, 3, 1, 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{47, 48, 1, 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMixedGroupUse2(t *testing.T) {
	src := `<? use foo\bar\{function foo, bar, const baz,};`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 47, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtGroupUse,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{3, 47, 1, 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeNameName,
						Group:    ast.NodeGroupPrefix,
						Position: ast.Position{7, 14, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{7, 10, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{6, 7, 1, 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{11, 14, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_NS_SEPARATOR,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{10, 11, 1, 1},
									},
								},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{16, 28, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupUseType,
								Position: ast.Position{16, 24, 1, 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{25, 28, 1, 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{25, 28, 1, 1},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Position: ast.Position{24, 25, 1, 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{30, 33, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{30, 33, 1, 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{30, 33, 1, 1},
										Tokens: []ast.Token{
											{
												Type:     ',',
												Group:    ast.TokenGroupStart,
												Position: ast.Position{28, 29, 1, 1},
											},
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupStart,
												Position: ast.Position{29, 30, 1, 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{35, 44, 1, 1},
						Tokens:   []ast.Token{},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupUseType,
								Position: ast.Position{35, 40, 1, 1},
								Tokens: []ast.Token{
									{
										Type:     ',',
										Group:    ast.TokenGroupStart,
										Position: ast.Position{33, 34, 1, 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{34, 35, 1, 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{41, 44, 1, 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{41, 44, 1, 1},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupStart,
												Position: ast.Position{40, 41, 1, 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ',',
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{44, 45, 1, 1},
					},
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{0, 2, 1, 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{2, 3, 1, 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{46, 47, 1, 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestGroupUse(t *testing.T) {
	src := `<? use function \foo\bar\{foo as bar, baz,};`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 44, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtGroupUse,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 44, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupUseType,
						Position: ast.Position{PS: 7, PE: 15, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStart,
								Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{PS: 26, PE: 36, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 26, PE: 29, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 26, PE: 29, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupAlias,
								Position: ast.Position{PS: 33, PE: 36, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{PS: 38, PE: 41, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 38, PE: 41, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 38, PE: 41, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{
												Type:     ',',
												Group:    ast.TokenGroupStart,
												Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
											},
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupStart,
												Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
					{
						Type:     ast.NodeTypeNameName,
						Group:    ast.NodeGroupPrefix,
						Position: ast.Position{PS: 17, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 21, PE: 24, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_NS_SEPARATOR,
										Group:    ast.TokenGroupStart,
										Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupUse,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_NS_SEPARATOR,
						Group:    ast.TokenGroupUse,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ',',
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_OPEN_TAG,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{PE: 2, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStart,
						Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestGroupUse2(t *testing.T) {
	src := `<? use const foo\bar\{foo, bar as baz,};`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 40, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtGroupUse,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 40, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupUseType,
						Position: ast.Position{PS: 7, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{PS: 22, PE: 25, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 22, PE: 25, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 22, PE: 25, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUseList,
						Position: ast.Position{PS: 27, PE: 37, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 27, PE: 30, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 27, PE: 30, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: ',', Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupAlias,
								Position: ast.Position{PS: 34, PE: 37, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{},
					},
					{
						Type:     ast.NodeTypeNameName,
						Group:    ast.NodeGroupPrefix,
						Position: ast.Position{PS: 13, PE: 20, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeNameNamePart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_NS_SEPARATOR, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ',',
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 39, PE: 40, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestUse(t *testing.T) {
	src := `<? use \foo\bar as baz, bar\foo;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 32, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtUseList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 32, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUses,
						Position: ast.Position{PS: 8, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 8, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 8, PE: 11, LS: 1, LE: 1},
									},
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 12, PE: 15, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_NS_SEPARATOR, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupAlias,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}},
							{
								Type:     scanner.T_NS_SEPARATOR,
								Group:    ast.TokenGroupSlash,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUses,
						Position: ast.Position{PS: 24, PE: 31, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 24, PE: 31, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: ',', Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
										},
									},
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 28, PE: 31, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_NS_SEPARATOR, Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestUse2(t *testing.T) {
	src := `<? use const foo\bar, \bar\foo as baz;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 38, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtUseList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 38, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupUseType,
						Position: ast.Position{PS: 7, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUses,
						Position: ast.Position{PS: 13, PE: 20, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 13, PE: 20, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_NS_SEPARATOR, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
					{
						Type:     ast.NodeTypeStmtUse,
						Group:    ast.NodeGroupUses,
						Position: ast.Position{PS: 23, PE: 37, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupUse,
								Position: ast.Position{PS: 23, PE: 30, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 23, PE: 26, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 27, PE: 30, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_NS_SEPARATOR, Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupAlias,
								Position: ast.Position{PS: 34, PE: 37, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1}},
							{
								Type:     scanner.T_NS_SEPARATOR,
								Group:    ast.TokenGroupSlash,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestConstant(t *testing.T) {
	src := `<? const foo = $a, bar = $b;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 28, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtConstList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 28, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtConstant,
						Group:    ast.NodeGroupConsts,
						Position: ast.Position{PS: 9, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupConstantName,
								Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtConstant,
						Group:    ast.NodeGroupConsts,
						Position: ast.Position{PS: 19, PE: 27, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupConstantName,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 25, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 25, PE: 27, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestSimpleVariable(t *testing.T) {
	src := `<? $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 6, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestSimpleVariable2(t *testing.T) {
	src := `<? ${$b};`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{Type: '$', Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1}},
							{Type: '{', Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}},
							{
								Type:     '}',
								Group:    ast.TokenGroupEnd,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestSimpleVariable3(t *testing.T) {
	src := `<? $$c;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 7, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{Type: '$', Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInnerStmtList(t *testing.T) {
	src := `<? {$c;}`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 8, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtStmtList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 4, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestList(t *testing.T) {
	src := `<? list() = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
								Children: []ast.Node{},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestList2(t *testing.T) {
	src := `<? list(,) = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:  ast.NodeTypeExprArrayItem,
										Group: ast.NodeGroupItems,
									},
									{
										Type:  ast.NodeTypeExprArrayItem,
										Group: ast.NodeGroupItems,
										Tokens: []ast.Token{
											{
												Type:     ',',
												Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestList3(t *testing.T) {
	src := `<? list(,$a,,$b,) = $c;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 23, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:  ast.NodeTypeExprArrayItem,
										Group: ast.NodeGroupItems,
									},
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
														Tokens: []ast.Token{
															{
																Type:     ',',
																Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
															},
														},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{},
									},
									{
										Type:  ast.NodeTypeExprArrayItem,
										Group: ast.NodeGroupItems,
										Tokens: []ast.Token{
											{
												Type:     ',',
												Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
														Tokens: []ast.Token{
															{
																Type:     ',',
																Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
															},
														},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{},
									},
									{
										Type:  ast.NodeTypeExprArrayItem,
										Group: ast.NodeGroupItems,
										Tokens: []ast.Token{
											{
												Type:     ',',
												Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestArrayItemExpression(t *testing.T) {
	src := `<? list($a) = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
													},
												},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestArrayItemExpressionPair(t *testing.T) {
	src := `<? list($a => $b) = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 23, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 8, PE: 16, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupKey,
												Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
													},
												},
											},
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupExpr,
												Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestArrayItemReference(t *testing.T) {
	src := `<? list(&$a) = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 18, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 8, PE: 11, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprReference,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 8, PE: 11, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeExprVariable,
														Group:    ast.NodeGroupVar,
														Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupVarName,
																Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
															},
														},
													},
												},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestArrayItemReferencePair(t *testing.T) {
	src := `<? list($a => &$b) = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 8, PE: 17, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupKey,
												Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
													},
												},
											},
											{
												Type:     ast.NodeTypeExprReference,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 14, PE: 17, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeExprVariable,
														Group:    ast.NodeGroupVar,
														Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupVarName,
																Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
															},
														},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupExpr,
												Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestArrayItemList(t *testing.T) {
	src := `<? list(list($a)) = $a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 23, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 8, PE: 16, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprList,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 8, PE: 16, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeExprArrayItem,
														Group:    ast.NodeGroupItems,
														Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeExprVariable,
																Group:    ast.NodeGroupVal,
																Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
																Children: []ast.Node{
																	{
																		Type:     ast.NodeTypeIdentifier,
																		Group:    ast.NodeGroupVarName,
																		Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestArrayItemListPair(t *testing.T) {
	src := `<? list($a => list($b)) = $c;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 29, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 29, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 28, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayItem,
										Group:    ast.NodeGroupItems,
										Position: ast.Position{PS: 8, PE: 22, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupKey,
												Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
													},
												},
											},
											{
												Type:     ast.NodeTypeExprList,
												Group:    ast.NodeGroupVal,
												Position: ast.Position{PS: 14, PE: 22, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeExprArrayItem,
														Group:    ast.NodeGroupItems,
														Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeExprVariable,
																Group:    ast.NodeGroupVal,
																Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
																Children: []ast.Node{
																	{
																		Type:     ast.NodeTypeIdentifier,
																		Group:    ast.NodeGroupVarName,
																		Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
																	},
																},
															},
														},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupExpr,
												Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceVar(t *testing.T) {
	src := `<? $a[];`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 8, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprArrayDimFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     '[',
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
							{
								Type:     ']',
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceConstant(t *testing.T) {
	src := `<? foo [ $a ] ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprArrayDimFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprConstFetch,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameName,
										Group:    ast.NodeGroupConstant,
										Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameNamePart,
												Group:    ast.NodeGroupParts,
												Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
												Tokens: []ast.Token{
													{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
													{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
												},
											},
										},
										Tokens: []ast.Token{},
									},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDim,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
							},
							{
								Type:     '[',
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
							{
								Type:     ']',
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceConstantClassFetch(t *testing.T) {
	src := `<? foo :: bar [ ] ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 19, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprArrayDimFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprClassConstFetch,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameName,
										Group:    ast.NodeGroupClass,
										Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameNamePart,
												Group:    ast.NodeGroupParts,
												Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
												Tokens: []ast.Token{
													{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
													{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
												},
											},
										},
										Tokens: []ast.Token{},
									},
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupConstantName,
										Position: ast.Position{PS: 10, PE: 13, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
							},
							{
								Type:     '[',
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
							{
								Type:     ']',
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceConstantVariableClassFetch(t *testing.T) {
	src := `<? $a :: foo [ ] ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 18, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprArrayDimFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprClassConstFetch,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupClass,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{
											{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
										},
									},
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupConstantName,
										Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
							{
								Type:     '[',
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
							},
							{
								Type:     ']',
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceExprDim(t *testing.T) {
	src := `<? ($a){$b};`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprArrayDimFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
									{Type: '(', Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1}},
									{
										Type:     ')',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDim,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     '{',
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
							{
								Type:     '}',
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceArrayDim(t *testing.T) {
	src := `<? array()[$b];`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprArrayDimFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArray,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDim,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     '[',
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
							},
							{
								Type:     ']',
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceShortArrayMethodCallEmptyArgumentList(t *testing.T) {
	src := `<? []->methodName();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprMethodCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprShortArray,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupMethod,
								Position: ast.Position{PS: 7, PE: 17, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 17, PE: 19, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceScalarStringExprMethodCallArgumentList(t *testing.T) {
	src := `<? "foo"->{$b}($a, ...$b, );`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 28, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 28, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprMethodCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeScalarString,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupMethod,
								Position: ast.Position{PS: 10, PE: 14, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{},
									},
								},
								Tokens: []ast.Token{
									{Type: '{', Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}},
									{
										Type:     '}',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 14, PE: 24, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeArgument,
										Group:    ast.NodeGroupArguments,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{},
									},
									{
										Type:     ast.NodeTypeArgument,
										Flag:     ast.NodeFlagVariadic,
										Group:    ast.NodeGroupArguments,
										Position: ast.Position{PS: 19, PE: 24, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 22, PE: 24, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 22, PE: 24, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     ',',
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceFetchMethod(t *testing.T) {
	src := `<? $a->$b;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprPropertyFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupProperty,
								Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStaticPropertyFetch(t *testing.T) {
	src := `<? static::$a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprStaticPropertyFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupProperty,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClassStaticPropertyFetch(t *testing.T) {
	src := `<? foo::$a;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 11, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprStaticPropertyFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupProperty,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferenceStaticPropertyFetch(t *testing.T) {
	src := `<? $a::$b;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprStaticPropertyFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupProperty,
								Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStaticCallIdentifier(t *testing.T) {
	src := `<? \foo::function();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprStaticCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameFullyQualified,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 4, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupCall,
								Position: ast.Position{PS: 9, PE: 17, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 17, PE: 19, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStaticCallExpr(t *testing.T) {
	src := `<? namespace\foo::{$a}();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 25, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 25, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprStaticCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameRelative,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupCall,
								Position: ast.Position{PS: 18, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{},
									},
								},
								Tokens: []ast.Token{
									{Type: '{', Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
									{
										Type:     '}',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 22, PE: 24, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStaticCallVar(t *testing.T) {
	src := `<? $a::$b();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprStaticCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupCall,
								Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionCallName(t *testing.T) {
	src := `<? foo();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprFunctionCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupFunction,
								Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionCallVar(t *testing.T) {
	src := `<? $a();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 8, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprFunctionCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupFunction,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionCallExpr(t *testing.T) {
	src := `<? ($a)();`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprFunctionCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupFunction,
								Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
									{Type: '(', Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1}},
									{
										Type:     ')',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionCallScalar(t *testing.T) {
	src := `<? 'foo'($a);`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprFunctionCall,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeScalarString,
								Group:    ast.NodeGroupFunction,
								Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 8, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeArgument,
										Group:    ast.NodeGroupArguments,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionInInnerStatement(t *testing.T) {
	src := `<? {function foo() {$a;}}`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 25, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtStmtList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 25, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtFunction,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 4, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupFunctionName,
								Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 20, PE: 23, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupReturnType,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionReturnTypedReference(t *testing.T) {
	src := `<? function &foo(): bar {}`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 26, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtFunction,
				Flag:     ast.NodeFlagRef,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 26, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupFunctionName,
						Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
						Tokens:   []ast.Token{},
					},
					{
						Type:     ast.NodeTypeStmtReturnType,
						Group:    ast.NodeGroupReturnType,
						Position: ast.Position{PS: 18, PE: 23, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 20, PE: 23, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 20, PE: 23, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupFunction,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupReturnType,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionParamsAndNullable(t *testing.T) {
	src := `<? function foo(?array & ...$a, & ...$b, ...$c, $d): ?callable {}`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 65, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtFunction,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 65, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupFunctionName,
						Position: ast.Position{PS: 12, PE: 15, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtReturnType,
						Group:    ast.NodeGroupReturnType,
						Position: ast.Position{PS: 51, PE: 62, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNullable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 53, PE: 62, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 54, PE: 62, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 52, PE: 53, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{},
					},
					{
						Type:     ast.NodeTypeParameter,
						Flag:     ast.NodeFlagRef | ast.NodeFlagVariadic,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 16, PE: 30, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNullable,
								Group:    ast.NodeGroupVarType,
								Position: ast.Position{PS: 16, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 17, PE: 22, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 28, PE: 30, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupOptionalType,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupAmpersand,
								Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeParameter,
						Flag:     ast.NodeFlagRef | ast.NodeFlagVariadic,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 32, PE: 39, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 37, PE: 39, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: ',', Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupAmpersand,
								Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeParameter,
						Flag:     ast.NodeFlagVariadic,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 41, PE: 46, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 44, PE: 46, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: ',', Position: ast.Position{PS: 39, PE: 40, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1}},
						},
					},
					{
						Type:     ast.NodeTypeParameter,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 48, PE: 50, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 48, PE: 50, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: ',', Position: ast.Position{PS: 46, PE: 47, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupReturnType,
						Position: ast.Position{PS: 62, PE: 63, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFunctionParamsWithDefaultValue(t *testing.T) {
	src := `<? function foo ( array & ... $a = $v, & ... $b = $v, ... $c = $v, $d = $v ) : array { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 88, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtFunction,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 88, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupFunctionName,
						Position: ast.Position{PS: 12, PE: 15, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtReturnType,
						Group:    ast.NodeGroupReturnType,
						Position: ast.Position{PS: 77, PE: 84, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 79, PE: 84, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 78, PE: 79, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 76, PE: 77, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeParameter,
						Flag:     ast.NodeFlagRef | ast.NodeFlagVariadic,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 18, PE: 37, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarType,
								Position: ast.Position{PS: 18, PE: 23, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDefaultValue,
								Position: ast.Position{PS: 35, PE: 37, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 35, PE: 37, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupOptionalType,
								Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupAmpersand,
								Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVariadic,
								Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeParameter,
						Flag:     ast.NodeFlagRef | ast.NodeFlagVariadic,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 39, PE: 52, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 45, PE: 47, LS: 1, LE: 1},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDefaultValue,
								Position: ast.Position{PS: 50, PE: 52, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 50, PE: 52, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 38, PE: 39, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupAmpersand,
								Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVariadic,
								Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeParameter,
						Flag:     ast.NodeFlagVariadic,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 54, PE: 65, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 58, PE: 60, LS: 1, LE: 1},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDefaultValue,
								Position: ast.Position{PS: 63, PE: 65, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 63, PE: 65, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 52, PE: 53, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 62, PE: 63, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 53, PE: 54, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVariadic,
								Position: ast.Position{PS: 57, PE: 58, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 60, PE: 61, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeParameter,
						Group:    ast.NodeGroupParams,
						Position: ast.Position{PS: 67, PE: 74, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 67, PE: 69, LS: 1, LE: 1},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupDefaultValue,
								Position: ast.Position{PS: 72, PE: 74, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 72, PE: 74, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 65, PE: 66, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 71, PE: 72, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 66, PE: 67, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 69, PE: 70, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupParamList,
						Position: ast.Position{PS: 74, PE: 75, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupReturnType,
						Position: ast.Position{PS: 84, PE: 85, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 86, PE: 87, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClass(t *testing.T) {
	src := `<? class foo { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClassInStmtList(t *testing.T) {
	src := `<? { class foo { } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtStmtList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtClass,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 5, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupClassName,
								Position: ast.Position{PS: 11, PE: 14, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClassExtendsImplements(t *testing.T) {
	src := `<? class foo extends bar implements baz , quz { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 49, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 49, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassExtends,
						Group:    ast.NodeGroupExtends,
						Position: ast.Position{PS: 13, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupClassName,
								Position: ast.Position{PS: 21, PE: 24, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 21, PE: 24, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassImplements,
						Group:    ast.NodeGroupImplements,
						Position: ast.Position{PS: 25, PE: 45, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupInterfaceNames,
								Position: ast.Position{PS: 36, PE: 39, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 36, PE: 39, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupInterfaceNames,
								Position: ast.Position{PS: 42, PE: 45, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 42, PE: 45, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 39, PE: 40, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClassModifiers(t *testing.T) {
	src := `<? final abstract class foo { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 31, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 31, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupModifiers,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupModifiers,
						Position: ast.Position{PS: 9, PE: 17, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupModifierList,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClassModifiersExtendsImplements(t *testing.T) {
	src := `<? abstract class foo extends bar implements baz , quz { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 58, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 58, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupModifiers,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 18, PE: 21, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassExtends,
						Group:    ast.NodeGroupExtends,
						Position: ast.Position{PS: 22, PE: 33, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupClassName,
								Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassImplements,
						Group:    ast.NodeGroupImplements,
						Position: ast.Position{PS: 34, PE: 54, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupInterfaceNames,
								Position: ast.Position{PS: 45, PE: 48, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 45, PE: 48, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupInterfaceNames,
								Position: ast.Position{PS: 51, PE: 54, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 51, PE: 54, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 48, PE: 49, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 50, PE: 51, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupModifierList,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 54, PE: 55, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 56, PE: 57, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTrait(t *testing.T) {
	src := `<? trait foo { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtTrait,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupTraitName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitInStmtList(t *testing.T) {
	src := `<? { trait foo { } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtStmtList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtTrait,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 5, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupTraitName,
								Position: ast.Position{PS: 11, PE: 14, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInterface(t *testing.T) {
	src := `<? interface foo { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtInterface,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInterfaceInStmtList(t *testing.T) {
	src := `<? { interface foo { } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtStmtList,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtInterface,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 5, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupClassName,
								Position: ast.Position{PS: 15, PE: 18, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInterfaceExtends(t *testing.T) {
	src := `<? interface foo extends bar, baz { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 37, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtInterface,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 37, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtInterfaceExtends,
						Group:    ast.NodeGroupExtends,
						Position: ast.Position{PS: 17, PE: 33, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupInterfaceNames,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupInterfaceNames,
								Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: ',', Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestProperties(t *testing.T) {
	src := `<? class foo { var $a ; public $b = $c ; static $d , $e ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 59, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 59, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtPropertyList,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 23, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupModifiers,
								Position: ast.Position{PS: 15, PE: 18, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtProperty,
								Group:    ast.NodeGroupProperties,
								Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupPropertyList,
								Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtPropertyList,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 24, PE: 40, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupModifiers,
								Position: ast.Position{PS: 24, PE: 30, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtProperty,
								Group:    ast.NodeGroupProperties,
								Position: ast.Position{PS: 31, PE: 38, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 36, PE: 38, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 36, PE: 38, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupPropertyList,
								Position: ast.Position{PS: 38, PE: 39, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 39, PE: 40, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtPropertyList,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 41, PE: 57, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupModifiers,
								Position: ast.Position{PS: 41, PE: 47, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtProperty,
								Group:    ast.NodeGroupProperties,
								Position: ast.Position{PS: 48, PE: 50, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 48, PE: 50, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 48, PE: 50, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtProperty,
								Group:    ast.NodeGroupProperties,
								Position: ast.Position{PS: 53, PE: 55, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 53, PE: 55, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 53, PE: 55, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 50, PE: 51, LS: 1, LE: 1}},
									{Type: ',', Position: ast.Position{PS: 51, PE: 52, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 52, PE: 53, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupPropertyList,
								Position: ast.Position{PS: 55, PE: 56, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 56, PE: 57, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 57, PE: 58, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClassConstants(t *testing.T) {
	src := `<? class foo { const foo = $a ; protected const bar = $b, baz = $c ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 70, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 70, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassConstList,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 31, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtConstant,
								Group:    ast.NodeGroupConsts,
								Position: ast.Position{PS: 21, PE: 29, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupConstantName,
										Position: ast.Position{PS: 21, PE: 24, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 27, PE: 29, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 27, PE: 29, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupConstList,
								Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtClassConstList,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 32, PE: 68, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupModifiers,
								Position: ast.Position{PS: 32, PE: 41, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtConstant,
								Group:    ast.NodeGroupConsts,
								Position: ast.Position{PS: 48, PE: 56, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupConstantName,
										Position: ast.Position{PS: 48, PE: 51, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 54, PE: 56, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 54, PE: 56, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 53, PE: 54, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 51, PE: 52, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeStmtConstant,
								Group:    ast.NodeGroupConsts,
								Position: ast.Position{PS: 58, PE: 66, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupConstantName,
										Position: ast.Position{PS: 58, PE: 61, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 64, PE: 66, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 64, PE: 66, LS: 1, LE: 1},
												Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 56, PE: 57, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 63, PE: 64, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 57, PE: 58, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 61, PE: 62, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupModifierList,
								Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupConstList,
								Position: ast.Position{PS: 66, PE: 67, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 67, PE: 68, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 68, PE: 69, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAbstractMethod(t *testing.T) {
	src := `<? class foo { function foo () ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 34, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 34, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassMethod,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 32, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupMethodName,
								Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtNop,
								Group:    ast.NodeGroupStmt,
								Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMethod(t *testing.T) {
	src := `<? class foo { private static function & foo () : bar { $a ; } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 64, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 64, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtClassMethod,
						Flag:     ast.NodeFlagRef,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 62, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupModifiers,
								Position: ast.Position{PS: 15, PE: 22, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupModifiers,
								Position: ast.Position{PS: 23, PE: 29, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupMethodName,
								Position: ast.Position{PS: 41, PE: 44, LS: 1, LE: 1},
								Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtReturnType,
								Position: ast.Position{PS: 48, PE: 53, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameName,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 50, PE: 53, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameNamePart,
												Group:    ast.NodeGroupParts,
												Position: ast.Position{PS: 50, PE: 53, LS: 1, LE: 1},
												Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtStmtList,
								Group:    ast.NodeGroupStmt,
								Position: ast.Position{PS: 54, PE: 62, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeStmtExpression,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 56, PE: 60, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 56, PE: 58, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 56, PE: 58, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 55, PE: 56, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 58, PE: 59, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 59, PE: 60, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 53, PE: 54, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStmts,
										Position: ast.Position{PS: 60, PE: 61, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupModifierList,
								Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupFunction,
								Position: ast.Position{PS: 38, PE: 39, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 62, PE: 63, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUseNop(t *testing.T) {
	src := `<? class foo { use foo ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 26, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 26, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtNop,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUseEmptyAdaptations(t *testing.T) {
	src := `<? class foo { use foo , bar { } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 34, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 34, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 32, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtTraitAdaptationList,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 29, PE: 32, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupAdaptationList,
										Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUsePrecedence(t *testing.T) {
	src := `<? class foo { use foo , bar { foo :: bar insteadof baz , quz ; } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 67, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 67, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 65, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtTraitAdaptationList,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 29, PE: 65, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeWrapper,
										Group:    ast.NodeGroupAdaptations,
										Position: ast.Position{PS: 31, PE: 63, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeStmtTraitUsePrecedence,
												Group:    ast.NodeGroupStmt,
												Position: ast.Position{PS: 31, PE: 61, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeStmtTraitMethodRef,
														Group:    ast.NodeGroupRef,
														Position: ast.Position{PS: 31, PE: 41, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeNameName,
																Group:    ast.NodeGroupTrait,
																Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
																Children: []ast.Node{
																	{
																		Type:     ast.NodeTypeNameNamePart,
																		Group:    ast.NodeGroupParts,
																		Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
																		Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
																	},
																},
																Tokens: []ast.Token{},
															},
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupMethod,
																Position: ast.Position{PS: 38, PE: 41, LS: 1, LE: 1},
																Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1}}},
															},
														},
														Tokens: []ast.Token{
															{
																Type:     scanner.T_WHITESPACE,
																Group:    ast.TokenGroupName,
																Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
															},
														},
													},
													{
														Type:     ast.NodeTypeNameName,
														Group:    ast.NodeGroupInsteadof,
														Position: ast.Position{PS: 52, PE: 55, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeNameNamePart,
																Group:    ast.NodeGroupParts,
																Position: ast.Position{PS: 52, PE: 55, LS: 1, LE: 1},
																Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 51, PE: 52, LS: 1, LE: 1}}},
															},
														},
														Tokens: []ast.Token{},
													},
													{
														Type:     ast.NodeTypeNameName,
														Group:    ast.NodeGroupInsteadof,
														Position: ast.Position{PS: 58, PE: 61, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeNameNamePart,
																Group:    ast.NodeGroupParts,
																Position: ast.Position{PS: 58, PE: 61, LS: 1, LE: 1},
																Tokens: []ast.Token{
																	{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 55, PE: 56, LS: 1, LE: 1}},
																	{Type: ',', Position: ast.Position{PS: 56, PE: 57, LS: 1, LE: 1}},
																	{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 57, PE: 58, LS: 1, LE: 1}},
																},
															},
														},
														Tokens: []ast.Token{},
													},
												},
												Tokens: []ast.Token{
													{
														Type:     scanner.T_WHITESPACE,
														Group:    ast.TokenGroupRef,
														Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1},
													},
												},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 61, PE: 62, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 62, PE: 63, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupAdaptationList,
										Position: ast.Position{PS: 63, PE: 64, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 65, PE: 66, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUseAlias(t *testing.T) {
	src := `<? class foo { use foo , bar { foo :: bar as baz ; } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 54, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 54, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 52, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtTraitAdaptationList,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 29, PE: 52, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeWrapper,
										Group:    ast.NodeGroupAdaptations,
										Position: ast.Position{PS: 31, PE: 50, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeStmtTraitUseAlias,
												Group:    ast.NodeGroupStmt,
												Position: ast.Position{PS: 31, PE: 48, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeStmtTraitMethodRef,
														Group:    ast.NodeGroupRef,
														Position: ast.Position{PS: 31, PE: 41, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeNameName,
																Group:    ast.NodeGroupTrait,
																Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
																Children: []ast.Node{
																	{
																		Type:     ast.NodeTypeNameNamePart,
																		Group:    ast.NodeGroupParts,
																		Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
																		Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
																	},
																},
																Tokens: []ast.Token{},
															},
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupMethod,
																Position: ast.Position{PS: 38, PE: 41, LS: 1, LE: 1},
																Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1}}},
															},
														},
														Tokens: []ast.Token{
															{
																Type:     scanner.T_WHITESPACE,
																Group:    ast.TokenGroupName,
																Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
															},
														},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupAlias,
														Position: ast.Position{PS: 45, PE: 48, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{
													{
														Type:     scanner.T_WHITESPACE,
														Group:    ast.TokenGroupRef,
														Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1},
													},
												},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 48, PE: 49, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupAdaptationList,
										Position: ast.Position{PS: 50, PE: 51, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 52, PE: 53, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUseAliasReserved(t *testing.T) {
	src := `<? class foo { use foo , bar { foo as function ; } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 52, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 52, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 50, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtTraitAdaptationList,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 29, PE: 50, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeWrapper,
										Group:    ast.NodeGroupAdaptations,
										Position: ast.Position{PS: 31, PE: 48, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeStmtTraitUseAlias,
												Group:    ast.NodeGroupStmt,
												Position: ast.Position{PS: 31, PE: 46, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeStmtTraitMethodRef,
														Group:    ast.NodeGroupRef,
														Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupMethod,
																Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
															},
														},
														Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupAlias,
														Position: ast.Position{PS: 38, PE: 46, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{
													{
														Type:     scanner.T_WHITESPACE,
														Group:    ast.TokenGroupRef,
														Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
													},
												},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 46, PE: 47, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupAdaptationList,
										Position: ast.Position{PS: 48, PE: 49, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 50, PE: 51, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUseAliasModifierOnly(t *testing.T) {
	src := `<? class foo { use foo , bar { foo as final ; } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 49, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 49, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 47, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtTraitAdaptationList,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 29, PE: 47, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeWrapper,
										Group:    ast.NodeGroupAdaptations,
										Position: ast.Position{PS: 31, PE: 45, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeStmtTraitUseAlias,
												Group:    ast.NodeGroupStmt,
												Position: ast.Position{PS: 31, PE: 43, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeStmtTraitMethodRef,
														Group:    ast.NodeGroupRef,
														Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupMethod,
																Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
															},
														},
														Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupModifier,
														Position: ast.Position{PS: 38, PE: 43, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{
													{
														Type:     scanner.T_WHITESPACE,
														Group:    ast.TokenGroupRef,
														Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
													},
												},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupAdaptationList,
										Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTraitUseAliasModifier(t *testing.T) {
	src := `<? class foo { use foo , bar { foo as public bar ; foo as abstract bar ; } }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 76, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtClass,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 76, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupClassName,
						Position: ast.Position{PS: 9, PE: 12, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtTraitUse,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 15, PE: 74, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTraits,
								Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeStmtTraitAdaptationList,
								Group:    ast.NodeGroupTraitAdaptationList,
								Position: ast.Position{PS: 29, PE: 74, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeWrapper,
										Group:    ast.NodeGroupAdaptations,
										Position: ast.Position{PS: 31, PE: 50, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeStmtTraitUseAlias,
												Group:    ast.NodeGroupStmt,
												Position: ast.Position{PS: 31, PE: 48, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeStmtTraitMethodRef,
														Group:    ast.NodeGroupRef,
														Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupMethod,
																Position: ast.Position{PS: 31, PE: 34, LS: 1, LE: 1},
															},
														},
														Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupModifier,
														Position: ast.Position{PS: 38, PE: 44, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1}}},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupAlias,
														Position: ast.Position{PS: 45, PE: 48, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{
													{
														Type:     scanner.T_WHITESPACE,
														Group:    ast.TokenGroupRef,
														Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
													},
												},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 48, PE: 49, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeWrapper,
										Group:    ast.NodeGroupAdaptations,
										Position: ast.Position{PS: 51, PE: 72, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeStmtTraitUseAlias,
												Group:    ast.NodeGroupStmt,
												Position: ast.Position{PS: 51, PE: 70, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeStmtTraitMethodRef,
														Group:    ast.NodeGroupRef,
														Position: ast.Position{PS: 51, PE: 54, LS: 1, LE: 1},
														Children: []ast.Node{
															{
																Type:     ast.NodeTypeIdentifier,
																Group:    ast.NodeGroupMethod,
																Position: ast.Position{PS: 51, PE: 54, LS: 1, LE: 1},
															},
														},
														Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 50, PE: 51, LS: 1, LE: 1}}},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupModifier,
														Position: ast.Position{PS: 58, PE: 66, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 57, PE: 58, LS: 1, LE: 1}}},
													},
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupAlias,
														Position: ast.Position{PS: 67, PE: 70, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 66, PE: 67, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{
													{
														Type:     scanner.T_WHITESPACE,
														Group:    ast.TokenGroupRef,
														Position: ast.Position{PS: 54, PE: 55, LS: 1, LE: 1},
													},
												},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 70, PE: 71, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 71, PE: 72, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupAdaptationList,
										Position: ast.Position{PS: 72, PE: 73, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupName,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 74, PE: 75, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestShortList(t *testing.T) {
	src := `<? [ ] = $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprShortList,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArrayPairList,
										Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClone(t *testing.T) {
	src := `<? clone $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprClone,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssign(t *testing.T) {
	src := `<? $a = $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignAssign,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignReference(t *testing.T) {
	src := `<? $a = & $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignReference,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupEqual,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignPlus(t *testing.T) {
	src := `<? $a += $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignPlus,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignMinus(t *testing.T) {
	src := `<? $a -= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignMinus,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignMul(t *testing.T) {
	src := `<? $a *= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignMul,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignPow(t *testing.T) {
	src := `<? $a **= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignPow,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignDiv(t *testing.T) {
	src := `<? $a /= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignDiv,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignConcat(t *testing.T) {
	src := `<? $a .= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignConcat,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignMod(t *testing.T) {
	src := `<? $a %= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignMod,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignBitwiseAnd(t *testing.T) {
	src := `<? $a &= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignBitwiseAnd,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignBitwiseOr(t *testing.T) {
	src := `<? $a |= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignBitwiseOr,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignBitwiseXor(t *testing.T) {
	src := `<? $a ^= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignBitwiseXor,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignShiftLeft(t *testing.T) {
	src := `<? $a <<= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignShiftLeft,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAssignShiftRight(t *testing.T) {
	src := `<? $a >>= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeAssignShiftRight,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPostInc(t *testing.T) {
	src := `<? $a ++ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprPostInc,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPostDec(t *testing.T) {
	src := `<? $a -- ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprPostDec,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPreInc(t *testing.T) {
	src := `<? ++ $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprPreInc,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPreDec(t *testing.T) {
	src := `<? -- $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprPreDec,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryBooleanOr(t *testing.T) {
	src := `<? $a || $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryBooleanOr,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryBooleanAnd(t *testing.T) {
	src := `<? $a && $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryBooleanAnd,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryLogicalOr(t *testing.T) {
	src := `<? $a or $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryLogicalOr,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryLogicalAnd(t *testing.T) {
	src := `<? $a and $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryLogicalAnd,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryLogicalXor(t *testing.T) {
	src := `<? $a xor $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryLogicalXor,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryBitwiseOr(t *testing.T) {
	src := `<? $a | $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryBitwiseOr,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryBitwiseAnd(t *testing.T) {
	src := `<? $a & $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryBitwiseAnd,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryBitwiseXor(t *testing.T) {
	src := `<? $a ^ $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryBitwiseXor,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryConcat(t *testing.T) {
	src := `<? $a . $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryConcat,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryPlus(t *testing.T) {
	src := `<? $a + $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryPlus,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryMinus(t *testing.T) {
	src := `<? $a - $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryMinus,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryMul(t *testing.T) {
	src := `<? $a * $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryMul,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryPow(t *testing.T) {
	src := `<? $a ** $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryPow,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryDiv(t *testing.T) {
	src := `<? $a / $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryDiv,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryMod(t *testing.T) {
	src := `<? $a % $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryMod,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryShiftLeft(t *testing.T) {
	src := `<? $a << $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryShiftLeft,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryShiftRight(t *testing.T) {
	src := `<? $a >> $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryShiftRight,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestUnaryPlus(t *testing.T) {
	src := `<? + $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprUnaryPlus,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestUnaryMinus(t *testing.T) {
	src := `<? - $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprUnaryMinus,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBooleanNot(t *testing.T) {
	src := `<? ! $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprBooleanNot,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBitwiseNot(t *testing.T) {
	src := `<? ~ $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprBitwiseNot,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryIdentical(t *testing.T) {
	src := `<? $a === $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryIdentical,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryNotIdentical(t *testing.T) {
	src := `<? $a !== $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryNotIdentical,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryEqual(t *testing.T) {
	src := `<? $a == $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryEqual,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryNotEqual(t *testing.T) {
	src := `<? $a != $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryNotEqual,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_IS_NOT_EQUAL,
								Group:    ast.TokenGroupEqual,
								Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryNotEqual2(t *testing.T) {
	src := `<? $a <> $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryNotEqual,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_IS_NOT_EQUAL,
								Group:    ast.TokenGroupEqual,
								Position: ast.Position{PS: 6, PE: 8, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinarySmaller(t *testing.T) {
	src := `<? $a < $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinarySmaller,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinarySmallerOrEqual(t *testing.T) {
	src := `<? $a <= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinarySmallerOrEqual,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryGreater(t *testing.T) {
	src := `<? $a > $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryGreater,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryGreaterOrEqual(t *testing.T) {
	src := `<? $a >= $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryGreaterOrEqual,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinarySpaceship(t *testing.T) {
	src := `<? $a <=> $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinarySpaceship,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBinaryCoalesce(t *testing.T) {
	src := `<? $a ?? $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeBinaryCoalesce,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupLeft,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupRight,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastInt(t *testing.T) {
	src := `<? ( int ) $a ; ( integer ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 32, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastInt,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_INT_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 16, PE: 32, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastInt,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 16, PE: 30, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 28, PE: 30, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 28, PE: 30, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}},
							{
								Type:     scanner.T_INT_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 16, PE: 27, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastDouble(t *testing.T) {
	src := `<? ( real ) $a ; ( double ) $a ; ( float ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 47, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastDouble,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_DOUBLE_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 17, PE: 32, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastDouble,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 17, PE: 30, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 28, PE: 30, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 28, PE: 30, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
							{
								Type:     scanner.T_DOUBLE_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 17, PE: 27, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
				},
			},
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 33, PE: 47, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastDouble,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 33, PE: 45, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 43, PE: 45, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 43, PE: 45, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 42, PE: 43, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1}},
							{
								Type:     scanner.T_DOUBLE_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 33, PE: 42, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 46, PE: 47, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastString(t *testing.T) {
	src := `<? ( string ) $a ; ( binary ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 34, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastString,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_STRING_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 19, PE: 34, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastString,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 19, PE: 32, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
							{
								Type:     scanner.T_STRING_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 19, PE: 29, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastArray(t *testing.T) {
	src := `<? ( array ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastArray,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_ARRAY_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastObject(t *testing.T) {
	src := `<? ( object ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 18, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastObject,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_OBJECT_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastBoolean(t *testing.T) {
	src := `<? ( bool ) $a ; ( boolean ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 33, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastBool,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_BOOL_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 17, PE: 33, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastBool,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 17, PE: 31, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 29, PE: 31, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 29, PE: 31, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
							{
								Type:     scanner.T_BOOL_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 17, PE: 28, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestCastUnset(t *testing.T) {
	src := `<? ( unset ) $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeCastUnset,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_UNSET_CAST,
								Group:    ast.TokenGroupCast,
								Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDie(t *testing.T) {
	src := `<? die ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 8, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprExit,
						Flag:     ast.NodeFlagAltSyntax,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExitNoExpr(t *testing.T) {
	src := `<? exit ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprExit,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 8, PE: 11, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}},
									{Type: '(', Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
									{
										Type:     ')',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestExitExpr(t *testing.T) {
	src := `<? exit ( $a ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprExit,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 8, PE: 14, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}},
									{Type: '(', Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
									},
									{
										Type:     ')',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestErrorSuppress(t *testing.T) {
	src := `<? @ $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 9, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprErrorSuppress,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPrint(t *testing.T) {
	src := `<? print $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprPrint,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestYield(t *testing.T) {
	src := `<? yield ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprYield,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestYieldVal(t *testing.T) {
	src := `<? yield $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprYield,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVal,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestYieldKeyVal(t *testing.T) {
	src := `<? yield $a => $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 19, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprYield,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupKey,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVal,
								Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestYieldFrom(t *testing.T) {
	src := `<? yield  from $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 19, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprYieldFrom,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{Type: scanner.T_YIELD_FROM, Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestWrappedExpr(t *testing.T) {
	src := `<? ( $a ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 11, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeWrapper,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{Type: '(', Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupEnd,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
							{
								Type:     ')',
								Group:    ast.TokenGroupEnd,
								Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTernaryFull(t *testing.T) {
	src := `<? $a ? $b : $c ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprTernary,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupCond,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupIfTrue,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupIfFalse,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCond,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupTrue,
								Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTernary(t *testing.T) {
	src := `<? $a ? : $c ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprTernary,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupCond,
								Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupIfFalse,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCond,
								Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupTrue,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestIsset(t *testing.T) {
	src := `<? isset ( $a , $b , ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprIsset,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVars,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVars,
								Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
										Tokens: []ast.Token{
											// TODO: this tokens must be at parent node
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupIsset,
								Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVarList,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
							{
								Type:     ',',
								Group:    ast.TokenGroupVarList,
								Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEmpty(t *testing.T) {
	src := `<? empty ( $a ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprEmpty,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupEmpty,
								Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInclude(t *testing.T) {
	src := `<? include $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprInclude,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestIncludeOnce(t *testing.T) {
	src := `<? include_once $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprIncludeOnce,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEval(t *testing.T) {
	src := `<? eval ( $a ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprEval,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupEval,
								Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestRequire(t *testing.T) {
	src := `<? require $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprRequire,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestRequireOnce(t *testing.T) {
	src := `<? require_once $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprRequireOnce,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNew(t *testing.T) {
	src := `<? new foo ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 7, PE: 10, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewVar(t *testing.T) {
	src := `<? new $a ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 10, PE: 13, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewVarDim(t *testing.T) {
	src := `<? new $a [ ] ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 19, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 14, PE: 17, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewVarDimExpr(t *testing.T) {
	src := `<? new $a [ $b ] ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 22, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewVarDim2(t *testing.T) {
	src := `<? new $a { $b } ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 22, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
									{
										Type:     '{',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
									},
									{
										Type:     '}',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewPropertyFetch(t *testing.T) {
	src := `<? new $a -> foo ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 22, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprPropertyFetch,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 16, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupClass,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupProperty,
										Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewStaticPropertyFetch(t *testing.T) {
	src := `<? new foo :: $bar ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprStaticPropertyFetch,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 18, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameName,
										Group:    ast.NodeGroupClass,
										Position: ast.Position{PS: 7, PE: 10, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameNamePart,
												Group:    ast.NodeGroupParts,
												Position: ast.Position{PS: 7, PE: 10, LS: 1, LE: 1},
												Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupProperty,
										Position: ast.Position{PS: 14, PE: 18, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 14, PE: 18, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 19, PE: 22, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewVarStaticPropertyFetch(t *testing.T) {
	src := `<? new $a :: $bar ( ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 23, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 21, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprStaticPropertyFetch,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupClass,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupProperty,
										Position: ast.Position{PS: 13, PE: 17, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 13, PE: 17, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeArgumentList,
								Group:    ast.NodeGroupArgumentList,
								Position: ast.Position{PS: 18, PE: 21, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupArgumentList,
										Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewSimpleAnonymousClass(t *testing.T) {
	src := `<? new class { } ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 18, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtClass,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 16, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStmts,
										Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNewFullAnonymousClass(t *testing.T) {
	src := `<? new class ( ) extends bar implements baz , quz { } ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 55, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 55, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprNew,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 53, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtClass,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 7, PE: 53, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeArgumentList,
										Group:    ast.NodeGroupArgumentList,
										Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}},
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupArgumentList,
												Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeStmtClassExtends,
										Group:    ast.NodeGroupExtends,
										Position: ast.Position{PS: 17, PE: 28, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameName,
												Group:    ast.NodeGroupClassName,
												Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeNameNamePart,
														Group:    ast.NodeGroupParts,
														Position: ast.Position{PS: 25, PE: 28, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeStmtClassImplements,
										Group:    ast.NodeGroupImplements,
										Position: ast.Position{PS: 29, PE: 49, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameName,
												Group:    ast.NodeGroupInterfaceNames,
												Position: ast.Position{PS: 40, PE: 43, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeNameNamePart,
														Group:    ast.NodeGroupParts,
														Position: ast.Position{PS: 40, PE: 43, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 39, PE: 40, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{},
											},
											{
												Type:     ast.NodeTypeNameName,
												Group:    ast.NodeGroupInterfaceNames,
												Position: ast.Position{PS: 46, PE: 49, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeNameNamePart,
														Group:    ast.NodeGroupParts,
														Position: ast.Position{PS: 46, PE: 49, LS: 1, LE: 1},
														Tokens: []ast.Token{
															{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1}},
															{Type: ',', Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1}},
															{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1}},
														},
													},
												},
												Tokens: []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupName,
										Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupStmts,
										Position: ast.Position{PS: 51, PE: 52, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 53, PE: 54, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 54, PE: 55, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInstanceOf(t *testing.T) {
	src := `<? new $a instanceof foo ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 26, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 26, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprInstanceOf,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprNew,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 3, PE: 9, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupClass,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
								},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupClass,
								Position: ast.Position{PS: 21, PE: 24, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 21, PE: 24, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestClosure(t *testing.T) {
	src := `<? function ( ) { } ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 21, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 21, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprClosure,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupFunction,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupParameterList,
								Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupReturnType,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStaticClosure(t *testing.T) {
	src := `<? static function ( ) { } ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 28, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 28, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprClosure,
						Flag:     ast.NodeFlagStatic,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 26, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStatic,
								Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupFunction,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupParameterList,
								Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupReturnType,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFullClosure(t *testing.T) {
	src := `<? function & ( ) use ( $a, & $b ) : foo { } ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 46, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 46, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprClosure,
						Flag:     ast.NodeFlagRef,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 44, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprClosureUse,
								Group:    ast.NodeGroupClosureUse,
								Position: ast.Position{PS: 18, PE: 34, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 24, PE: 26, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 24, PE: 26, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeExprReference,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 28, PE: 32, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVar,
												Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupUse,
										Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupLexicalVarList,
										Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeStmtReturnType,
								Group:    ast.NodeGroupReturnType,
								Position: ast.Position{PS: 35, PE: 40, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameName,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 37, PE: 40, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameNamePart,
												Group:    ast.NodeGroupParts,
												Position: ast.Position{PS: 37, PE: 40, LS: 1, LE: 1},
												Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupFunction,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupAmpersand,
								Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupParameterList,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupReturnType,
								Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 42, PE: 43, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFullStaticClosure(t *testing.T) {
	src := `<? static function & ( ) use ( $a, & $b ) : foo { } ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 53, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 53, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprClosure,
						Flag:     ast.NodeFlagStatic | ast.NodeFlagRef,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 51, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprClosureUse,
								Group:    ast.NodeGroupClosureUse,
								Position: ast.Position{PS: 25, PE: 41, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeExprReference,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 35, PE: 39, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVar,
												Position: ast.Position{PS: 37, PE: 39, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 37, PE: 39, LS: 1, LE: 1},
														Tokens:   []ast.Token{{Type: ',', Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1}}},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupUse,
										Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupLexicalVarList,
										Position: ast.Position{PS: 39, PE: 40, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeStmtReturnType,
								Group:    ast.NodeGroupReturnType,
								Position: ast.Position{PS: 42, PE: 47, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameName,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 44, PE: 47, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeNameNamePart,
												Group:    ast.NodeGroupParts,
												Position: ast.Position{PS: 44, PE: 47, LS: 1, LE: 1},
												Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStatic,
								Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupFunction,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupAmpersand,
								Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupParameterList,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupReturnType,
								Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 49, PE: 50, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 51, PE: 52, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 52, PE: 53, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestLNumber(t *testing.T) {
	src := `<? 1 ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 6, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarLnumber,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDNumber(t *testing.T) {
	src := `<? .1 ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 7, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarDnumber,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantLine(t *testing.T) {
	src := `<? __LINE__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantFile(t *testing.T) {
	src := `<? __FILE__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantDir(t *testing.T) {
	src := `<? __DIR__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 12, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantTrait(t *testing.T) {
	src := `<? __TRAIT__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantMethod(t *testing.T) {
	src := `<? __METHOD__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 15, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantFunction(t *testing.T) {
	src := `<? __FUNCTION__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 15, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantNamespace(t *testing.T) {
	src := `<? __NAMESPACE__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 18, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestMagicConstantClass(t *testing.T) {
	src := `<? __CLASS__ ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Tokens:   []ast.Token{},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarMagicConstant,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 12, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEmptyHeredoc(t *testing.T) {
	src := `<? <<<EOT
EOT;
`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 2},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 2},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarHeredoc,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 2},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 2, LE: 2},
					},
				},
			},
		},
		Tokens: []ast.Token{
			{
				Type:     scanner.T_WHITESPACE,
				Group:    ast.TokenGroupEnd,
				Position: ast.Position{PS: 14, PE: 15, LS: 2, LE: 2},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestHeredoc(t *testing.T) {
	src := `<? <<<EOT
	some text
EOT;
`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 25, 1, 3},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 25, LS: 1, LE: 3},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarHeredoc,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 3},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 10, PE: 21, LS: 2, LE: 2},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 24, PE: 25, LS: 3, LE: 3},
					},
				},
			},
		},
		Tokens: []ast.Token{
			{
				Type:     scanner.T_WHITESPACE,
				Group:    ast.TokenGroupEnd,
				Position: ast.Position{PS: 25, PE: 26, LS: 3, LE: 3},
			},
		},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDereferencableScalar(t *testing.T) {
	src := `<? [ ] ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 8, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprShortArray,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupArrayPairList,
								Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestConstantScalar(t *testing.T) {
	src := `<? FOO ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 8, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 8, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprConstFetch,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupConstant,
								Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVar(t *testing.T) {
	src := `<? "string $a string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 23, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 21, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 11, LS: 1, LE: 1},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 13, PE: 20, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarStringOffset(t *testing.T) {
	src := `<? "$a[FOO] string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 21, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 21, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeScalarString,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 7, PE: 10, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 11, PE: 18, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarStringOffset2(t *testing.T) {
	src := `<? "$a[9223372036854775808] string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 37, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 37, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 35, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeScalarString,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 7, PE: 26, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 27, PE: 34, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarStringOffset3(t *testing.T) {
	src := `<? "$a[-9223372036854775808] string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 38, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 38, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 36, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeScalarString,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 7, PE: 27, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 28, PE: 35, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarNumStringOffset(t *testing.T) {
	src := `<? "$a[9223372036854775807] string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 37, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 37, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 35, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeScalarLnumber,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 7, PE: 26, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 27, PE: 34, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarUnaryMinusOffset(t *testing.T) {
	src := `<? "$a[-9223372036854775807] string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 38, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 38, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 36, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeExprUnaryMinus,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 7, PE: 27, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeScalarLnumber,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 8, PE: 27, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 28, PE: 35, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarOffsetVar(t *testing.T) {
	src := `<? "$a[$b] string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayDimFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 10, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupDim,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     '[',
										Group:    ast.TokenGroupVar,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
									},
									{
										Type:     ']',
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 10, PE: 17, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarFetchProperty(t *testing.T) {
	src := `<? "$a->foo string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 21, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 21, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprPropertyFetch,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVar,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
											},
										},
									},
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupProperty,
										Position: ast.Position{PS: 8, PE: 11, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 11, PE: 18, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarDollarOpenCurlBracesExpr(t *testing.T) {
	src := `<? "${ $a } string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 21, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 21, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 7, PE: 9, LS: 1, LE: 1},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1}}},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_DOLLAR_OPEN_CURLY_BRACES,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
									},
									{
										Type:     '}',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 11, PE: 18, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarDollarOpenCurlBracesStringVarName(t *testing.T) {
	src := `<? "${a} string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 18, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 18, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 8, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_DOLLAR_OPEN_CURLY_BRACES,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
									},
									{
										Type:     '}',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 8, PE: 15, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarDollarOpenCurlBracesStringVarNameFetch(t *testing.T) {
	src := `<? "${a[ $b ] } string" ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 25, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 25, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarEncapsed,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 23, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprArrayDimFetch,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 6, PE: 13, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupVar,
												Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{},
											},
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupDim,
												Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     '[',
												Group:    ast.TokenGroupVar,
												Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1},
											},
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupExpr,
												Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
											},
											{
												Type:     ']',
												Group:    ast.TokenGroupExpr,
												Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_DOLLAR_OPEN_CURLY_BRACES,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
									},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
									},
									{
										Type:     '}',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 15, PE: 22, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedStringVarCurlyOpen(t *testing.T) {
	src := `<? <<<"EOT"
{$a } string"
EOT;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 30, 1, 3},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 30, LS: 1, LE: 3},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeScalarHeredoc,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 29, LS: 1, LE: 3},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeWrapper,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 12, PE: 17, LS: 2, LE: 2},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 13, PE: 15, LS: 2, LE: 2},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 13, PE: 15, LS: 2, LE: 2},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_CURLY_OPEN, Position: ast.Position{PS: 12, PE: 13, LS: 2, LE: 2}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 15, PE: 16, LS: 2, LE: 2},
									},
									{
										Type:     '}',
										Group:    ast.TokenGroupEnd,
										Position: ast.Position{PS: 16, PE: 17, LS: 2, LE: 2},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 17, PE: 26, LS: 2, LE: 2},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 29, PE: 30, LS: 3, LE: 3},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEmptyBackticksExpr(t *testing.T) {
	src := "<? `` ;"

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 7, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprShellExec,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 5, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBackticksExpr(t *testing.T) {
	src := "<? `foo` ;"

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 10, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 10, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprShellExec,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 7, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEncapsedBackticksExpr(t *testing.T) {
	src := "<? `$a foo $b` ;"

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprShellExec,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 4, PE: 6, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeScalarEncapsedStringPart,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 6, PE: 11, LS: 1, LE: 1},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupParts,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestIf(t *testing.T) {
	src := `<? if ( $a ) $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtIf,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 13, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIf,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestIfElseIF(t *testing.T) {
	src := `<? if ( $a ) $b ; elseif ( $c ) $d ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 36, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtIf,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 36, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 13, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtElseIf,
						Group:    ast.NodeGroupElseIf,
						Position: ast.Position{PS: 18, PE: 36, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupCond,
								Position: ast.Position{PS: 27, PE: 29, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 27, PE: 29, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmt,
								Position: ast.Position{PS: 32, PE: 36, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 32, PE: 34, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 32, PE: 34, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupElseIf,
								Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIf,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestIfElse(t *testing.T) {
	src := `<? if ( $a ) $b ; else $c ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 27, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtIf,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 27, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 13, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtElse,
						Group:    ast.NodeGroupElse,
						Position: ast.Position{PS: 18, PE: 27, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmt,
								Position: ast.Position{PS: 23, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 23, PE: 25, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 23, PE: 25, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIf,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltIf(t *testing.T) {
	src := `<? if ( $a ) : $b ; $c ; endif ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 32, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtAltIf,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 32, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 15, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 15, PE: 19, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 20, PE: 24, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
									},
								},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIf,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltIfElseIf(t *testing.T) {
	src := `<? if ( $a ) : $b ; elseif ( $c ) : $d ; endif ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 48, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtAltIf,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 48, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 15, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 15, PE: 19, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
									},
								},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtAltElseIf,
						Flag:     ast.NodeFlagAltSyntax,
						Group:    ast.NodeGroupElseIf,
						Position: ast.Position{PS: 20, PE: 40, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupElseIf,
								Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupExpr,
								Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCond,
								Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIf,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 46, PE: 47, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltIfElse(t *testing.T) {
	src := `<? if ( $a ) : $b ; else : $d ; endif ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 39, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtAltIf,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 39, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 15, PE: 19, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 15, PE: 19, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
									},
								},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtAltElse,
						Flag:     ast.NodeFlagAltSyntax,
						Group:    ast.NodeGroupElse,
						Position: ast.Position{PS: 20, PE: 31, LS: 1, LE: 1},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupElse,
								Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIf,
						Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 37, PE: 38, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 38, PE: 39, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestWhile(t *testing.T) {
	src := `<? while ( $a ) { $a ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtWhile,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 16, PE: 24, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 18, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 18, PE: 20, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 18, PE: 20, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupWhile,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltWhile(t *testing.T) {
	src := `<? while ( $a ) : $a ; endwhile ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 33, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtWhile,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 33, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 18, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 18, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 18, PE: 20, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 18, PE: 20, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
									},
								},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupWhile,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDoWhile(t *testing.T) {
	src := `<? do { $a } while ( $b ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 27, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtDo,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 27, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 6, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 5, PE: 6, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupWhile,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestFor(t *testing.T) {
	src := `<? for ( $a ; ; $b ) { $c ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 29, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtFor,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 29, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupInit,
						Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupLoop,
						Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 21, PE: 29, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 23, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 23, PE: 25, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 23, PE: 25, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupFor,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupInitExpr,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCondExpr,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIncExpr,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltFor(t *testing.T) {
	src := `<? for ( $a ; $b , $c ; $d ) : $e ; endfor ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 44, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtAltFor,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 44, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupInit,
						Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 14, PE: 16, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
									{Type: ',', Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupLoop,
						Position: ast.Position{PS: 24, PE: 26, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 24, PE: 26, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 31, PE: 35, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 31, PE: 35, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
									},
								},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 42, PE: 43, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupFor,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupInitExpr,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCondExpr,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupIncExpr,
						Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEmptySwitch(t *testing.T) {
	src := `<? switch ( $a ) { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 20, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtSwitch,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 20, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtCaseList,
						Group:    ast.NodeGroupCaseList,
						Position: ast.Position{PS: 17, PE: 20, LS: 1, LE: 1},
						Children: []ast.Node{},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCaseListEnd,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSwitch,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestSwitch(t *testing.T) {
	src := `<? switch ( $a ) { ; case $a : $b ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 37, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtSwitch,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 37, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtCaseList,
						Group:    ast.NodeGroupCaseList,
						Position: ast.Position{PS: 17, PE: 37, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtCase,
								Group:    ast.NodeGroupCases,
								Position: ast.Position{PS: 21, PE: 35, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeStmtExpression,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 31, PE: 35, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
									},
									{
										Type:     ':',
										Group:    ast.TokenGroupCaseSeparator,
										Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCaseListStart,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupCaseListStart,
								Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCaseListEnd,
								Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSwitch,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltSwitch(t *testing.T) {
	src := `<? switch ( $a ) : case $a : $b ; endswitch ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 45, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtSwitch,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 45, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtCaseList,
						Group:    ast.NodeGroupCaseList,
						Position: ast.Position{PS: 19, PE: 33, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtCase,
								Group:    ast.NodeGroupCases,
								Position: ast.Position{PS: 19, PE: 33, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 24, PE: 26, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 24, PE: 26, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
									},
									{
										Type:     ast.NodeTypeStmtExpression,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 29, PE: 33, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 29, PE: 31, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 29, PE: 31, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupExpr,
										Position: ast.Position{PS: 26, PE: 27, LS: 1, LE: 1},
									},
									{
										Type:     ':',
										Group:    ast.TokenGroupCaseSeparator,
										Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCaseListEnd,
								Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSwitch,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltSwitchDefault(t *testing.T) {
	src := `<? switch ( $a ) : ; default ; $b ; endswitch ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 47, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtSwitch,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 47, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupCond,
						Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtCaseList,
						Group:    ast.NodeGroupCaseList,
						Position: ast.Position{PS: 21, PE: 35, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtDefault,
								Group:    ast.NodeGroupCases,
								Position: ast.Position{PS: 21, PE: 35, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeStmtExpression,
										Group:    ast.NodeGroupStmts,
										Position: ast.Position{PS: 31, PE: 35, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeExprVariable,
												Group:    ast.NodeGroupExpr,
												Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
												Children: []ast.Node{
													{
														Type:     ast.NodeTypeIdentifier,
														Group:    ast.NodeGroupVarName,
														Position: ast.Position{PS: 31, PE: 33, LS: 1, LE: 1},
														Tokens:   []ast.Token{},
													},
												},
												Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1}}},
											},
										},
										Tokens: []ast.Token{
											{
												Type:     scanner.T_WHITESPACE,
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
											},
											{
												Type:     ';',
												Group:    ast.TokenGroupSemiColon,
												Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
											},
										},
									},
								},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}},
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupDefault,
										Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupCaseSeparator,
										Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCaseListStart,
								Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupCaseListStart,
								Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCaseListEnd,
								Position: ast.Position{PS: 35, PE: 36, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 46, PE: 47, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSwitch,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestBreak(t *testing.T) {
	src := `<? break $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtBreak,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestContinue(t *testing.T) {
	src := `<? continue $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 16, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtContinue,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 16, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 12, PE: 14, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestReturn(t *testing.T) {
	src := `<? return $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 14, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtReturn,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 14, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestGlobal(t *testing.T) {
	src := `<? global $a , $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 19, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtGlobal,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 19, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupVars,
						Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupVars,
						Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}},
									{Type: ',', Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVarList,
						Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStatic(t *testing.T) {
	src := `<? static $a , $b = $c ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtStatic,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtStaticVar,
						Group:    ast.NodeGroupVars,
						Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 10, PE: 12, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStaticVar,
						Group:    ast.NodeGroupVars,
						Position: ast.Position{PS: 15, PE: 22, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 15, PE: 17, LS: 1, LE: 1},
									},
								},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 20, PE: 22, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}},
											{Type: ',', Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 17, PE: 18, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVarList,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestEcho(t *testing.T) {
	src := `<? echo $a , $b ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 17, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtEcho,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 17, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExprs,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExprs,
						Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}},
									{Type: ',', Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_ECHO,
						Group:    ast.TokenGroupEcho,
						Position: ast.Position{PS: 3, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestInlineHtml(t *testing.T) {
	src := ` foo <? $a ?> bar `

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{0, 18, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtInlineHtml,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PE: 5, LS: 1, LE: 1},
			},
			{
				Type:     ast.NodeTypeStmtExpression,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 8, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 8, PE: 10, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_OPEN_TAG, Position: ast.Position{PS: 5, PE: 7, LS: 1, LE: 1}},
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
					},
				},
			},
			{
				Type:     ast.NodeTypeStmtInlineHtml,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 13, PE: 18, LS: 1, LE: 1},
				Tokens:   []ast.Token{},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestUnset(t *testing.T) {
	src := `<? unset ( $a , $b , ) ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 24, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtUnset,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 24, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupVars,
						Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 11, PE: 13, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupVars,
						Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 16, PE: 18, LS: 1, LE: 1},
								Tokens: []ast.Token{
									{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1}},
									{Type: ',', Position: ast.Position{PS: 14, PE: 15, LS: 1, LE: 1}},
								},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupUnset,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVarList,
						Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1},
					},
					{
						Type:     ',',
						Group:    ast.TokenGroupVarList,
						Position: ast.Position{PS: 19, PE: 20, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVarList,
						Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCloseParenthesisToken,
						Position: ast.Position{PS: 22, PE: 23, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestNop(t *testing.T) {
	src := `<? ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 4, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtNop,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestForeach1(t *testing.T) {
	src := `<? foreach ( $a as $b => $c ) { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 33, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtForeach,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 33, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupKey,
						Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupVar,
						Position: ast.Position{PS: 25, PE: 27, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 25, PE: 27, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 24, PE: 25, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
						Children: []ast.Node{},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 31, PE: 32, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupForeach,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupKey,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVar,
						Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestForeach2(t *testing.T) {
	src := `<? foreach ( $a as & $b ) { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 29, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtForeach,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 29, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprReference,
						Group:    ast.NodeGroupVar,
						Position: ast.Position{PS: 19, PE: 23, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
										Tokens:   []ast.Token{},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 26, PE: 29, LS: 1, LE: 1},
						Children: []ast.Node{},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupForeach,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVar,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestForeach3(t *testing.T) {
	src := `<? foreach ( $a as list ( $b ) ) { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 36, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtForeach,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 36, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprList,
						Group:    ast.NodeGroupVar,
						Position: ast.Position{PS: 19, PE: 30, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayItem,
								Group:    ast.NodeGroupItems,
								Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVal,
										Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupList,
								Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupArrayPairList,
								Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 33, PE: 36, LS: 1, LE: 1},
						Children: []ast.Node{},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupForeach,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVar,
						Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestForeach4(t *testing.T) {
	src := `<? foreach ( $a as [ $b ] ) : $c ; endforeach ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 47, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtAltForeach,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 47, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 13, PE: 15, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}}},
					},
					{
						Type:     ast.NodeTypeExprShortList,
						Group:    ast.NodeGroupVar,
						Position: ast.Position{PS: 19, PE: 25, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprArrayItem,
								Group:    ast.NodeGroupItems,
								Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupVal,
										Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 21, PE: 23, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 20, PE: 21, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupArrayPairList,
								Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 30, PE: 34, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 30, PE: 34, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 30, PE: 32, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 32, PE: 33, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1},
									},
								},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 34, PE: 35, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 46, PE: 47, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupForeach,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupVar,
						Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestDeclare(t *testing.T) {
	src := `<? declare ( FOO = $a ) { }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 27, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtDeclare,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 27, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtConstant,
						Group:    ast.NodeGroupConsts,
						Position: ast.Position{PS: 13, PE: 21, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupConstantName,
								Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
						Children: []ast.Node{},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupDeclare,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupConstList,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltDeclare(t *testing.T) {
	src := `<? declare ( FOO = $a ) : $b ; enddeclare ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 43, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtDeclare,
				Flag:     ast.NodeFlagAltSyntax,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 43, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtConstant,
						Group:    ast.NodeGroupConsts,
						Position: ast.Position{PS: 13, PE: 21, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupConstantName,
								Position: ast.Position{PS: 13, PE: 16, LS: 1, LE: 1},
								Tokens:   []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 19, PE: 21, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 18, PE: 19, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupName,
								Position: ast.Position{PS: 16, PE: 17, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtStmtList,
						Group:    ast.NodeGroupStmt,
						Position: ast.Position{PS: 26, PE: 30, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 26, PE: 30, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 26, PE: 28, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 25, PE: 26, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1},
									},
								},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupCond,
						Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 30, PE: 31, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupAltEnd,
						Position: ast.Position{PS: 41, PE: 42, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 42, PE: 43, LS: 1, LE: 1},
					},
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupDeclare,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupConstList,
						Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTry(t *testing.T) {
	src := `<? try { $a ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtTry,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 9, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupTry,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTryCatch(t *testing.T) {
	src := `<? try { $a ; } catch ( foo | bar $b ) { $c ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 47, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtTry,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 47, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 9, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtCatch,
						Group:    ast.NodeGroupCatches,
						Position: ast.Position{PS: 16, PE: 47, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTypes,
								Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTypes,
								Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1}},
											{Type: '|', Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 34, PE: 36, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 34, PE: 36, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 41, PE: 45, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 41, PE: 43, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 41, PE: 43, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCatch,
								Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCond,
								Position: ast.Position{PS: 38, PE: 39, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupTry,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestTryCatchFinally(t *testing.T) {
	src := `<? try { $a ; } catch ( foo | bar $b ) { $c ; } finally { $d ; }`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 64, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtTry,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 64, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeStmtExpression,
						Group:    ast.NodeGroupStmts,
						Position: ast.Position{PS: 9, PE: 13, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupExpr,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
							},
						},
						Tokens: []ast.Token{
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
							},
							{
								Type:     ';',
								Group:    ast.TokenGroupSemiColon,
								Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtCatch,
						Group:    ast.NodeGroupCatches,
						Position: ast.Position{PS: 16, PE: 47, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTypes,
								Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 24, PE: 27, LS: 1, LE: 1},
										Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 23, PE: 24, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeNameName,
								Group:    ast.NodeGroupTypes,
								Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeNameNamePart,
										Group:    ast.NodeGroupParts,
										Position: ast.Position{PS: 30, PE: 33, LS: 1, LE: 1},
										Tokens: []ast.Token{
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 27, PE: 28, LS: 1, LE: 1}},
											{Type: '|', Position: ast.Position{PS: 28, PE: 29, LS: 1, LE: 1}},
											{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 29, PE: 30, LS: 1, LE: 1}},
										},
									},
								},
								Tokens: []ast.Token{},
							},
							{
								Type:     ast.NodeTypeExprVariable,
								Group:    ast.NodeGroupVar,
								Position: ast.Position{PS: 34, PE: 36, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeIdentifier,
										Group:    ast.NodeGroupVarName,
										Position: ast.Position{PS: 34, PE: 36, LS: 1, LE: 1},
									},
								},
								Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 33, PE: 34, LS: 1, LE: 1}}},
							},
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 41, PE: 45, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 41, PE: 43, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 41, PE: 43, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 40, PE: 41, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 43, PE: 44, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 44, PE: 45, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 15, PE: 16, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCatch,
								Position: ast.Position{PS: 21, PE: 22, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupVar,
								Position: ast.Position{PS: 36, PE: 37, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupCond,
								Position: ast.Position{PS: 38, PE: 39, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 45, PE: 46, LS: 1, LE: 1},
							},
						},
					},
					{
						Type:     ast.NodeTypeStmtFinally,
						Group:    ast.NodeGroupFinally,
						Position: ast.Position{PS: 56, PE: 64, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeStmtExpression,
								Group:    ast.NodeGroupStmts,
								Position: ast.Position{PS: 58, PE: 62, LS: 1, LE: 1},
								Children: []ast.Node{
									{
										Type:     ast.NodeTypeExprVariable,
										Group:    ast.NodeGroupExpr,
										Position: ast.Position{PS: 58, PE: 60, LS: 1, LE: 1},
										Children: []ast.Node{
											{
												Type:     ast.NodeTypeIdentifier,
												Group:    ast.NodeGroupVarName,
												Position: ast.Position{PS: 58, PE: 60, LS: 1, LE: 1},
												Tokens:   []ast.Token{},
											},
										},
										Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 57, PE: 58, LS: 1, LE: 1}}},
									},
								},
								Tokens: []ast.Token{
									{
										Type:     scanner.T_WHITESPACE,
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 60, PE: 61, LS: 1, LE: 1},
									},
									{
										Type:     ';',
										Group:    ast.TokenGroupSemiColon,
										Position: ast.Position{PS: 61, PE: 62, LS: 1, LE: 1},
									},
								},
							},
						},
						Tokens: []ast.Token{
							{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 47, PE: 48, LS: 1, LE: 1}},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupFinally,
								Position: ast.Position{PS: 55, PE: 56, LS: 1, LE: 1},
							},
							{
								Type:     scanner.T_WHITESPACE,
								Group:    ast.TokenGroupStmts,
								Position: ast.Position{PS: 62, PE: 63, LS: 1, LE: 1},
							},
						},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupTry,
						Position: ast.Position{PS: 6, PE: 7, LS: 1, LE: 1},
					},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupStmts,
						Position: ast.Position{PS: 13, PE: 14, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestThrow(t *testing.T) {
	src := `<? throw $a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 13, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtThrow,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 13, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeExprVariable,
						Group:    ast.NodeGroupExpr,
						Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
						Children: []ast.Node{
							{
								Type:     ast.NodeTypeIdentifier,
								Group:    ast.NodeGroupVarName,
								Position: ast.Position{PS: 9, PE: 11, LS: 1, LE: 1},
							},
						},
						Tokens: []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupExpr,
						Position: ast.Position{PS: 11, PE: 12, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 12, PE: 13, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestGoto(t *testing.T) {
	src := `<? goto a ;`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 11, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtGoto,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 11, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupLabel,
						Position: ast.Position{PS: 8, PE: 9, LS: 1, LE: 1},
						Tokens:   []ast.Token{{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 7, PE: 8, LS: 1, LE: 1}}},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupLabel,
						Position: ast.Position{PS: 9, PE: 10, LS: 1, LE: 1},
					},
					{
						Type:     ';',
						Group:    ast.TokenGroupSemiColon,
						Position: ast.Position{PS: 10, PE: 11, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestLabel(t *testing.T) {
	src := `<? a :`

	expected := &ast.Node{
		Type:     ast.NodeTypeRoot,
		Position: ast.Position{3, 6, 1, 1},
		Children: []ast.Node{
			{
				Type:     ast.NodeTypeStmtLabel,
				Group:    ast.NodeGroupStmts,
				Position: ast.Position{PS: 3, PE: 6, LS: 1, LE: 1},
				Children: []ast.Node{
					{
						Type:     ast.NodeTypeIdentifier,
						Group:    ast.NodeGroupLabelName,
						Position: ast.Position{PS: 3, PE: 4, LS: 1, LE: 1},
					},
				},
				Tokens: []ast.Token{
					{Type: scanner.T_OPEN_TAG, Position: ast.Position{PE: 2, LS: 1, LE: 1}},
					{Type: scanner.T_WHITESPACE, Position: ast.Position{PS: 2, PE: 3, LS: 1, LE: 1}},
					{
						Type:     scanner.T_WHITESPACE,
						Group:    ast.TokenGroupLabel,
						Position: ast.Position{PS: 4, PE: 5, LS: 1, LE: 1},
					},
				},
			},
		},
		Tokens: []ast.Token{},
	}

	php7parser := php7.NewParser()
	a := &tree.Tree{}

	php7parser.Parse([]byte(src), a)

	actual := a.RootNode()
	assert.DeepEqual(t, expected, actual)
}
