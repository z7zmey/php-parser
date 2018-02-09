
package stmt_test

import (
	"github.com/z7zmey/php-parser/node/name"
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestInterface(t *testing.T) {
	src := `<? interface Foo {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestInterfaceExtend(t *testing.T) {
	src := `<? interface Foo extends Bar {}`

	expected := &stmt.StmtList{
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestInterfaceExtends(t *testing.T) {
	src := `<? interface Foo extends Bar, Baz {}`

	expected := &stmt.StmtList{
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

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
