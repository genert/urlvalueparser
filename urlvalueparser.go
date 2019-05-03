package urlvalueparser

import (
	"fmt"
	"strings"
)

func ReplacePathValues(path, replaceValue string) string {
	if path == "" || replaceValue == "" {
		return path
	}

	valueDetector := NewValueDetector()
	values := strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})

	var result []string

	for _, value := range values {
		if valueDetector.IsValue(value) {
			result = append(result, replaceValue)
		} else {
			result = append(result, value)
		}
	}

	return fmt.Sprintf("/%s", strings.Join(result, "/"))
}
