package version

import (
	"errors"
	"strconv"
	"strings"
)

type version struct {
	major int
	minor int
}

func Compare(a string, b string) (int, error) {
	first, err := parse(a)
	if err != nil {
		return 0, err
	}

	second, err := parse(b)
	if err != nil {
		return 0, err
	}

	if first.major < second.major {
		return -1, nil
	}

	if first.major > second.major {
		return 1, nil
	}

	if first.minor < second.minor {
		return -1, nil
	}

	if first.minor > second.minor {
		return 1, nil
	}

	return 0, nil
}

func parse(v string) (version, error) {
	i := strings.Index(v, ".")
	if i == -1 {
		return version{}, errors.New("version must contain major and minor parts")
	}

	major, err := strconv.Atoi(v[:i])
	if err != nil {
		return version{}, err
	}

	minor, err := strconv.Atoi(v[i+1:])
	if err != nil {
		return version{}, err
	}

	return version{major, minor}, nil
}
