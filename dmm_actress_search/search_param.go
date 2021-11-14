package dmmapi

import (
	"fmt"
	"strconv"
	"strings"
)

type SearchParam struct {
	minBust   int64
	maxBust   int64
	minWaist  int64
	maxWaist  int64
	minHip    int64
	maxHip    int64
	minHeight int64
	maxHeight int64
	minAge    int64
	maxAge    int64
	sortType  string
	keyword   string
}

type SearchParamOption func(param *SearchParam)

func NewSearchParam(options ...SearchParamOption) *SearchParam {
	var s SearchParam

	for _, opt := range options {
		opt(&s)
	}

	return &s
}

func WithBust(min int64, max int64) SearchParamOption {
	return func(param *SearchParam) {
		if min > 0 {
			param.minBust = min
		}
		if max > 0 {
			param.maxBust = max
		}
	}
}

func WithWaist(min int64, max int64) SearchParamOption {
	return func(param *SearchParam) {
		if min > 0 {
			param.minWaist = min
		}
		if max > 0 {
			param.maxWaist = max
		}
	}
}

func WithHip(min int64, max int64) SearchParamOption {
	return func(param *SearchParam) {
		if min > 0 {
			param.minWaist = min
		}
		if max > 0 {
			param.maxWaist = max
		}
	}
}

func WithHeight(min int64, max int64) SearchParamOption {
	return func(param *SearchParam) {
		if min > 0 {
			param.minHeight = min
		}
		if max > 0 {
			param.maxHeight = max
		}
	}
}

func WithAge(min int64, max int64) SearchParamOption {
	return func(param *SearchParam) {
		if min > 0 {
			param.minAge = min
		}
		if max > 0 {
			param.maxAge = max
		}
	}
}

func WithSortType(typ string) SearchParamOption {
	return func(param *SearchParam) {
		param.sortType = typ
	}
}

func WithKeyword(keyword string) SearchParamOption {
	return func(param *SearchParam) {
		param.keyword = keyword
	}
}

func ParseNumericParam(value string) (int64, int64, error) {
	if value == "" {
		return -1, -1, nil
	}

	if strings.HasPrefix(value, ">=") {
		numStr := value[2:]
		num, err := strconv.ParseInt(numStr, 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse parameter %s: %w", value, err)
		}

		return num, -1, nil
	}

	if strings.HasPrefix(value, ">") {
		numStr := value[1:]
		num, err := strconv.ParseInt(numStr, 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse parameter %s: %w", value, err)
		}

		return num + 1, -1, nil
	}

	if strings.HasPrefix(value, "<=") {
		numStr := value[2:]
		num, err := strconv.ParseInt(numStr, 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse parameter %s: %w", value, err)
		}

		return -1, num, nil
	}

	if strings.HasPrefix(value, "<") {
		numStr := value[1:]
		num, err := strconv.ParseInt(numStr, 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse parameter %s: %w", value, err)
		}

		return -1, num - 1, nil
	}

	if strings.Contains(value, ",") {
		parts := strings.Split(value, ",")
		if len(parts) != 2 {
			return 0, 0, fmt.Errorf("invalid parameter: %s", value)
		}

		min, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid parameter: %s: %w", value, err)
		}

		max, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid parameter: %s: %w", value, err)
		}

		if min > max {
			return max, min, nil
		}

		return min, max, nil
	}

	return 0, 0, fmt.Errorf("invalid parameter: %s", value)
}

func IsValidSortType(typ string) bool {
	if strings.HasPrefix(typ, "-") {
		typ = typ[1:]
	}

	types := map[string]struct{}{
		"name":   struct{}{},
		"bust":   struct{}{},
		"waist":  struct{}{},
		"hip":    struct{}{},
		"height": struct{}{},
		"age":    struct{}{},
		"id":     struct{}{},
	}

	_, ok := types[typ]
	return ok
}
