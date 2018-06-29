package meta_test

import (
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/meta"
)

func TestWhiteSpacePrint(t *testing.T) {
	expected := "\n    "

	w := meta.NewWhiteSpace(expected, nil)

	actual := w.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestWhiteSpaceSetGetTokenName(t *testing.T) {
	expected := meta.ArrayToken
	w := meta.NewWhiteSpace("\n    ", nil)
	w.SetTokenName(expected)

	actual := w.GetTokenName()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestWhiteSpaceGetPosition(t *testing.T) {
	expected := position.NewPosition(1, 1, 1, 1)
	q := meta.NewWhiteSpace("\n    ", expected)

	actual := q.GetPosition()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
