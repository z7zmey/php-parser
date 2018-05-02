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

func TestInterface(t *testing.T) {
	src := `<? interface Foo {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Stmts:         []node.Node{},
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

func TestInterfaceExtend(t *testing.T) {
	src := `<? interface Foo extends Bar {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Extends: []node.Node{
					&name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Bar"},
						},
					},
				},
				Stmts: []node.Node{},
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

func TestInterfaceExtends(t *testing.T) {
	src := `<? interface Foo extends Bar, Baz {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Extends: []node.Node{
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
				Stmts: []node.Node{},
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
