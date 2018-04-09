package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTrait(t *testing.T) {
	src := `<? trait Foo {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Trait{
				PhpDocComment: "",
				TraitName:     &node.Identifier{Value: "Foo"},
				Stmts:         []node.Node{},
			},
		},
	}

	actual, _, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
