package parser_test

import (
	"testing"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
)

func TestComments(t *testing.T) {
	n := node.NewIdentifier("test")

	commentGroup := []*comment.Comment{
		comment.NewComment("/** hello world */", nil),
		comment.NewComment("// hello world", nil),
	}

	comments := parser.Comments{}
	comments.AddComments(n, commentGroup)

	if comments[n][0].String() != "/** hello world */" {
		t.Errorf("expected and actual are not equal\n")
	}
	if comments[n][1].String() != "// hello world" {
		t.Errorf("expected and actual are not equal\n")
	}
}
