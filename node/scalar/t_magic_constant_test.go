package scalar_test

import (
	"bytes"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
)

func TestMagicConstant(t *testing.T) {
	src := `<? __DIR__;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__DIR__"},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
