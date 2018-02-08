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

func TestDeclare(t *testing.T) {
	src := `<? declare(ticks=1);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.Nop{},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestDeclareStmts(t *testing.T) {
	src := `<? declare(ticks=1) {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
