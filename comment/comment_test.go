package comment_test

import (
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

func TestComments(t *testing.T) {
	n := node.NewIdentifier("test")

	commentGroup := []*comment.Comment{
		comment.NewComment("/** hello world */", nil),
		comment.NewComment("// hello world", nil),
	}

	comments := comment.Comments{}
	comments.AddComments(n, commentGroup)

	if comments[n][0].String() != "/** hello world */" {
		t.Errorf("expected and actual are not equal\n")
	}
	if comments[n][1].String() != "// hello world" {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentPos(t *testing.T) {
	expected := position.NewPosition(0, 0, 0, 0)

	comment := comment.NewComment("/** hello world */", expected)

	actual := comment.Position()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
