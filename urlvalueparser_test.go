package urlvalueparser_test

import (
	"github.com/genert/urlvalueparser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplacePathValues(t *testing.T) {
	tests := []struct {
		path          string
		replaceValue  string
		expectedValue string
	}{
		{
			path:          "/path/something/else",
			replaceValue:  "",
			expectedValue: "/path/something/else",
		},
		{
			path:          "/path/something/1234",
			replaceValue:  "",
			expectedValue: "/path/something/1234",
		},
		{
			path:          "/path/something/1234",
			replaceValue:  ":id",
			expectedValue: "/path/something/:id",
		},
		{
			path:          "/v1/c44d02b2-0947-4850-b751-1b777c08dcad/component/cce9de2c-5927-46a6-ba71-80efd50702f6/action",
			replaceValue:  ":id",
			expectedValue: "/v1/:id/component/:id/action",
		},
	}

	for _, tt := range tests {
		result := urlvalueparser.ReplacePathValues(tt.path, tt.replaceValue)
		assert.Equal(t, tt.expectedValue, result)
	}
}
