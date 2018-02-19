package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestClassConstList(t *testing.T) {
	src := `<? class foo{ public const FOO = 1, BAR = 2; }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
						},
						Consts: []node.Node{
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "FOO"},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "BAR"},
								Expr:          &scalar.Lnumber{Value: "2"},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestClassConstListWithoutModifiers(t *testing.T) {
	src := `<? class foo{ const FOO = 1, BAR = 2; }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Consts: []node.Node{
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "FOO"},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "BAR"},
								Expr:          &scalar.Lnumber{Value: "2"},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
