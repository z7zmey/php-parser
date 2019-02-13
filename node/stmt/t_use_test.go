package stmt_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestSimpleUse(t *testing.T) {
	src := `<? use Foo;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    11,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    11,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    10,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
									Value: "Foo",
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

func TestUseFullyQualified(t *testing.T) {
	src := `<? use \Foo;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    12,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    12,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    11,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    11,
									},
									Value: "Foo",
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

func TestUseFullyQualifiedAlias(t *testing.T) {
	src := `<? use \Foo as Bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    19,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    19,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  9,
							EndPos:    18,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  9,
								EndPos:    11,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  9,
										EndPos:    11,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
							Value: "Bar",
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

func TestUseList(t *testing.T) {
	src := `<? use Foo, Bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    16,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    16,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    10,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    15,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
									Value: "Bar",
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

func TestUseListAlias(t *testing.T) {
	src := `<? use Foo, Bar as Baz;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    23,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    23,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    10,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    22,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
							Value: "Baz",
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

func TestUseListFunctionType(t *testing.T) {
	src := `<? use function Foo, \Bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    26,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    26,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    15,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  17,
							EndPos:    19,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    19,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    19,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  23,
							EndPos:    25,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  23,
								EndPos:    25,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  23,
										EndPos:    25,
									},
									Value: "Bar",
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

func TestUseListFunctionTypeAliases(t *testing.T) {
	src := `<? use function Foo as foo, \Bar as bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    40,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    40,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    15,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  17,
							EndPos:    26,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    19,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    19,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    26,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  30,
							EndPos:    39,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  30,
								EndPos:    32,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  30,
										EndPos:    32,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  37,
								EndPos:    39,
							},
							Value: "bar",
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

func TestUseListConstType(t *testing.T) {
	src := `<? use const Foo, \Bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    23,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    23,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    12,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    16,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    16,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    16,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  20,
							EndPos:    22,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  20,
								EndPos:    22,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  20,
										EndPos:    22,
									},
									Value: "Bar",
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

func TestUseListConstTypeAliases(t *testing.T) {
	src := `<? use const Foo as foo, \Bar as bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    37,
		},
		Stmts: []node.Node{
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    37,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    12,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    23,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    16,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    16,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  21,
								EndPos:    23,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  27,
							EndPos:    36,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  27,
								EndPos:    29,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  27,
										EndPos:    29,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  34,
								EndPos:    36,
							},
							Value: "bar",
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

func TestGroupUse(t *testing.T) {
	src := `<? use Foo\{Bar, Baz};`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    22,
		},
		Stmts: []node.Node{
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    22,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    10,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    15,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    20,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
									Value: "Baz",
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
}

func TestGroupUseAlias(t *testing.T) {
	src := `<? use Foo\{Bar, Baz as quux};`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    30,
		},
		Stmts: []node.Node{
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    30,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    10,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  13,
							EndPos:    15,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    15,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  13,
										EndPos:    15,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  18,
							EndPos:    28,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  18,
								EndPos:    20,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
									Value: "Baz",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  25,
								EndPos:    28,
							},
							Value: "quux",
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
}

func TestFunctionGroupUse(t *testing.T) {
	src := `<? use function Foo\{Bar, Baz};`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    31,
		},
		Stmts: []node.Node{
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    31,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    15,
					},
					Value: "function",
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  17,
						EndPos:    19,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    19,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  22,
							EndPos:    24,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  22,
								EndPos:    24,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  22,
										EndPos:    24,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  27,
							EndPos:    29,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  27,
								EndPos:    29,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  27,
										EndPos:    29,
									},
									Value: "Baz",
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
}

func TestConstGroupUse(t *testing.T) {
	src := `<? use const Foo\{Bar, Baz};`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    28,
		},
		Stmts: []node.Node{
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    28,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    12,
					},
					Value: "const",
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  14,
						EndPos:    16,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    16,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    21,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    21,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    21,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  24,
							EndPos:    26,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    26,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  24,
										EndPos:    26,
									},
									Value: "Baz",
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
}

func TestMixedGroupUse(t *testing.T) {
	src := `<? use Foo\{const Bar, function Baz};`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    37,
		},
		Stmts: []node.Node{
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    37,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  8,
						EndPos:    10,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    21,
						},
						UseType: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    17,
							},
							Value: "const",
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    21,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  19,
										EndPos:    21,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  33,
							EndPos:    35,
						},
						UseType: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  24,
								EndPos:    31,
							},
							Value: "function",
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  33,
								EndPos:    35,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  33,
										EndPos:    35,
									},
									Value: "Baz",
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
}
