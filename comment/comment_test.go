package comment_test

import (
	"testing"

	"github.com/z7zmey/php-parser/comment"
)

func TestCommentPrint(t *testing.T) {
	expected := "/** hello world */"

	comment := comment.NewComment(expected, nil)

	actual := comment.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentSetTokenName(t *testing.T) {
	expected := comment.ArrayToken
	c := comment.NewComment("/** hello world */", nil)
	c.SetTokenName(expected)

	actual := c.TokenName

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
