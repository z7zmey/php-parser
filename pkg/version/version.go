package version

import (
	"errors"
	"strconv"
	"strings"
)

type Version struct {
	Major, Minor uint64
}

var (
	// ErrInvalidSemVer is returned if a version can not be parsed
	ErrInvalidSemVer = errors.New("invalid semantic version")

	// ErrUnsupportedVer is returned if a version out of supported range
	ErrUnsupportedVer = errors.New("the version is out of supported range")

	php5RangeStart = &Version{Major: 5}
	php5RangeEnd   = &Version{Major: 5, Minor: 6}

	php7RangeStart = &Version{Major: 7}
	php7RangeEnd   = &Version{Major: 7, Minor: 4}
)

func New(v string) (*Version, error) {
	// Split the parts into [0]Major, [1]Minor
	parts := strings.SplitN(v, ".", 2)
	if len(parts) != 2 {
		return nil, ErrInvalidSemVer
	}

	var ver = new(Version)
	var err error

	ver.Major, err = strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}

	ver.Minor, err = strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return ver, nil
}

func (v *Version) Validate() error {
	if !v.InRange(php5RangeStart, php5RangeEnd) && !v.InRange(php7RangeStart, php7RangeEnd) {
		return ErrUnsupportedVer
	}

	return nil
}

// Less tests if one version is less than another one
func (v *Version) Less(o *Version) bool {
	return v.Compare(o) < 0
}

// LessOrEqual tests if one version is less than another one or equal
func (v *Version) LessOrEqual(o *Version) bool {
	return v.Compare(o) <= 0
}

// Greater tests if one version is greater than another one
func (v *Version) Greater(o *Version) bool {
	return v.Compare(o) > 0
}

// GreaterOrEqual tests if one version is greater than another one or equal
func (v *Version) GreaterOrEqual(o *Version) bool {
	return v.Compare(o) >= 0
}

// GreaterOrEqual tests if one version is greater than another one or equal
func (v *Version) InRange(s, e *Version) bool {
	return v.Compare(s) >= 0 && v.Compare(e) <= 0
}

// Compare compares this version to another one. It returns -1, 0, or 1 if
// the version smaller, equal, or larger than the other version.
func (v *Version) Compare(o *Version) int {
	if d := compareSegment(v.Major, o.Major); d != 0 {
		return d
	}

	return compareSegment(v.Minor, o.Minor)
}

func compareSegment(v, o uint64) int {
	if v < o {
		return -1
	}
	if v > o {
		return 1
	}

	return 0
}
