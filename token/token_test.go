package token_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/comment"

	"github.com/z7zmey/php-parser/token"
)

func TestToken(t *testing.T) {
	tkn := token.NewToken([]byte(`foo`), 1, 1, 0, 3)

	c := []comment.Comment{
		comment.NewPlainComment("test comment"),
	}

	tkn.SetComments(c)

	if reflect.DeepEqual(tkn.Comments(), c) {
		t.Errorf("comments are not equal\n")
	}

	if tkn.String() != `foo` {
		t.Errorf("token value is not equal\n")
	}

	if tkn.StartLine != 1 || tkn.EndLine != 1 || tkn.StartPos != 0 || tkn.EndPos != 3 {
		t.Errorf("token position is not equal\n")
	}
}
