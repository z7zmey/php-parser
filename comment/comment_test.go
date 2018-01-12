package comment_test

import (
	"testing"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

func TestComments(t *testing.T) {
	n := node.NewIdentifier("test")

	commentGroup := []comment.Comment{
		comment.NewDocComment("/** hello world */"),
		comment.NewPlainComment("// hello world"),
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
