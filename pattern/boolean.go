package pattern

import (
	"fmt"
	"strconv"
)

type Boolean bool

func (b Boolean) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func (b Boolean) Match(s string, _ map[string]string) (map[string]string, error) {
	x, err := strconv.ParseBool(s)
	if err != nil {
		return nil, fmt.Errorf("expected %s but matched value %s could not be interpreted as a boolean", b, s)
	}

	if bool(b) != x {
		return nil, fmt.Errorf("expected %s but matched value %s", b, s)
	}

	return nil, nil
}