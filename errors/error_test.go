package errors_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/scanner"

	"github.com/kylelemons/godebug/pretty"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		diff := pretty.Compare(expected, actual)

		if diff != "" {
			t.Errorf("diff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("expected and actual are not equal\n")
		}

	}
}

func TestConstructor(t *testing.T) {
	token := scanner.Token{
		Value:     "test",
		StartLine: 1,
		EndLine:   2,
		StartPos:  3,
		EndPos:    4,
	}

	actual := errors.NewError("message", token)

	expected := &errors.Error{
		Msg: "message",
		Pos: position.Position{
			StartLine: 1,
			EndLine:   2,
			StartPos:  3,
			EndPos:    4,
		},
	}

	assertEqual(t, expected, actual)
}

func TestPrint(t *testing.T) {
	token := scanner.Token{
		Value:     "test",
		StartLine: 1,
		EndLine:   2,
		StartPos:  3,
		EndPos:    4,
	}

	Error := errors.NewError("message", token)

	actual := Error.String()

	expected := "message at line 1"

	assertEqual(t, expected, actual)
}
