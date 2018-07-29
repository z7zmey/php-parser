package meta_test

import (
	"testing"

	"github.com/z7zmey/php-parser/meta"
)

func TestTypeString(t *testing.T) {
	c := meta.CommentType

	expected := "CommentType"
	actual := c.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestWrongTypeString(t *testing.T) {
	c := meta.Type(-1)

	expected := "Type(-1)"
	actual := c.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
