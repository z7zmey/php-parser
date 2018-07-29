package meta_test

import (
	"testing"

	"github.com/z7zmey/php-parser/meta"
)

func TestCommentPrint(t *testing.T) {
	expected := "/** hello world */"

	comment := meta.Data{
		Value: expected,
	}

	actual := comment.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
