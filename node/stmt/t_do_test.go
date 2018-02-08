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

func TestDo(t *testing.T) {
	src := `<? do {} while(1);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Do{
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{},
				},
				Cond: &scalar.Lnumber{Value: "1"},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
