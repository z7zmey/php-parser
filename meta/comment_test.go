package meta_test

import (
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/meta"
)

func TestCommentPrint(t *testing.T) {
	expected := "/** hello world */"

	comment := meta.NewComment(expected, nil)

	actual := comment.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentSetGetTokenName(t *testing.T) {
	expected := meta.ArrayToken
	c := meta.NewComment("/** hello world */", nil)
	c.SetTokenName(expected)

	actual := c.GetTokenName()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentGetPosition(t *testing.T) {
	expected := position.NewPosition(1, 1, 1, 1)
	c := meta.NewComment("/** hello world */", expected)

	actual := c.GetPosition()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
