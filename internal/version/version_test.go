package version_test

import (
	"gotest.tools/assert"
	"testing"

	"github.com/z7zmey/php-parser/internal/version"
)

func TestSmaller(t *testing.T) {
	r, err := version.Compare("7.3", "5.6")

	assert.NilError(t, err)
	assert.Equal(t, 1, r)
}

func TestGreater(t *testing.T) {
	r, err := version.Compare("5.6", "7.3")

	assert.NilError(t, err)
	assert.Equal(t, -1, r)
}

func TestEqual(t *testing.T) {
	r, err := version.Compare("7.3", "7.3")

	assert.NilError(t, err)
	assert.Equal(t, 0, r)
}
