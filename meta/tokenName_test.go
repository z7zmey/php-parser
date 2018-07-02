package meta_test

import (
	"testing"

	"github.com/z7zmey/php-parser/meta"
)

func TestTokenNameString(t *testing.T) {
	c := meta.NsSeparatorToken

	expected := "NsSeparatorToken"
	actual := c.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestWrongTokenNameString(t *testing.T) {
	c := meta.TokenName(-1)

	expected := "TokenName(-1)"
	actual := c.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
