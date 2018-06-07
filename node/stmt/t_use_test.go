package stmt_test

import (
	"bytes"
	"github.com/z7zmey/php-parser/node/name"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestSimpleUse(t *testing.T) {
	src := `<? use Foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
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

func TestUseFullyQualified(t *testing.T) {
	src := `<? use \Foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
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

func TestUseFullyQualifiedAlias(t *testing.T) {
	src := `<? use \Foo as Bar;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Alias: &node.Identifier{Value: "Bar"},
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

func TestUseList(t *testing.T) {
	src := `<? use Foo, Bar;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
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

func TestUseListAlias(t *testing.T) {
	src := `<? use Foo, Bar as Baz;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						Alias: &node.Identifier{Value: "Baz"},
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

func TestUseListFunctionType(t *testing.T) {
	src := `<? use function Foo, \Bar;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				UseType: &node.Identifier{Value: "function"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
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

func TestUseListFunctionTypeAliases(t *testing.T) {
	src := `<? use function Foo as foo, \Bar as bar;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				UseType: &node.Identifier{Value: "function"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Alias: &node.Identifier{Value: "foo"},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						Alias: &node.Identifier{Value: "bar"},
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

func TestUseListConstType(t *testing.T) {
	src := `<? use const Foo, \Bar;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				UseType: &node.Identifier{Value: "const"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
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

func TestUseListConstTypeAliases(t *testing.T) {
	src := `<? use const Foo as foo, \Bar as bar;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.UseList{
				UseType: &node.Identifier{Value: "const"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Alias: &node.Identifier{Value: "foo"},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						Alias: &node.Identifier{Value: "bar"},
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

func TestGroupUse(t *testing.T) {
	src := `<? use Foo\{Bar, Baz};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.GroupUse{
				Prefix: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Baz"},
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
}

func TestGroupUseAlias(t *testing.T) {
	src := `<? use Foo\{Bar, Baz as quux};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.GroupUse{
				Prefix: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Baz"},
							},
						},
						Alias: &node.Identifier{Value: "quux"},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestFunctionGroupUse(t *testing.T) {
	src := `<? use function Foo\{Bar, Baz};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.GroupUse{
				UseType: &node.Identifier{Value: "function"},
				Prefix: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Baz"},
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
}

func TestConstGroupUse(t *testing.T) {
	src := `<? use const Foo\{Bar, Baz};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.GroupUse{
				UseType: &node.Identifier{Value: "const"},
				Prefix: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Baz"},
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
}

func TestMixedGroupUse(t *testing.T) {
	src := `<? use Foo\{const Bar, function Baz};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.GroupUse{
				Prefix: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						UseType: &node.Identifier{Value: "const"},
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
					&stmt.Use{
						UseType: &node.Identifier{Value: "function"},
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Baz"},
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
}
