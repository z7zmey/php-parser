package scanner_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/comment"

	"github.com/z7zmey/php-parser/scanner"
)

func TestToken(t *testing.T) {
	pos := position.NewPosition(1, 1, 0, 3)
	tkn := scanner.NewToken(`foo`, pos)

	c := []*comment.Comment{
		comment.NewComment("test comment", nil),
	}

	tkn.SetComments(c)

	if !reflect.DeepEqual(tkn.Comments(), c) {
		t.Errorf("comments are not equal\n")
	}

	if tkn.String() != `foo` {
		t.Errorf("token value is not equal\n")
	}

	if tkn.Position() != pos {
		t.Errorf("token position is not equal\n")
	}
}
