package errors_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/position"
)

func TestConstructor(t *testing.T) {
	pos := position.NewPosition(1, 2, 3, 4)

	actual := errors.NewError("message", pos)

	expected := &errors.Error{
		Msg: "message",
		Pos: pos,
	}

	assert.DeepEqual(t, expected, actual)
}

func TestPrint(t *testing.T) {
	pos := position.NewPosition(1, 2, 3, 4)

	Error := errors.NewError("message", pos)

	actual := Error.String()

	expected := "message at line 1"

	assert.DeepEqual(t, expected, actual)
}

func TestPrintWithotPos(t *testing.T) {
	Error := errors.NewError("message", nil)

	actual := Error.String()

	expected := "message"

	assert.DeepEqual(t, expected, actual)
}
