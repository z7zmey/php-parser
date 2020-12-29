package version_test

import (
	"gotest.tools/assert"
	"testing"

	"github.com/z7zmey/php-parser/pkg/version"
)

func Test(t *testing.T) {
	ver, err := version.New("7.4")
	assert.NilError(t, err)

	assert.Equal(t, *ver, version.Version{
		Major: 7,
		Minor: 4,
	})
}

func TestLeadingZero(t *testing.T) {
	ver, err := version.New("07.04")
	assert.NilError(t, err)

	assert.Equal(t, *ver, version.Version{
		Major: 7,
		Minor: 4,
	})
}

func TestInRange(t *testing.T) {
	s, err := version.New("7.0")
	assert.NilError(t, err)

	e, err := version.New("7.4")
	assert.NilError(t, err)

	ver, err := version.New("7.0")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))

	ver, err = version.New("7.2")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))

	ver, err = version.New("7.4")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))
}
