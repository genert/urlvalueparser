package urlvalueparser

import (
	"fmt"
	"strings"
)

var valueDetect = NewValueDetector()

func ReplacePathValues(path, replaceValue string) string {
	if path == "" || replaceValue == "" {
		return path
	}

	values := strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})

	var result []string

	for _, value := range values {
		if valueDetect.IsValue(value) {
			result = append(result, replaceValue)
		} else {
			result = append(result, value)
		}
	}

	return fmt.Sprintf("/%s", strings.Join(result, "/"))
}
