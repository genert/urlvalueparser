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
		{
			path:          "/something/new/0x323b5d4c32345ced77393b3530b1eed0f346429d/another",
			replaceValue:  ":id",
			expectedValue: "/something/new/:id/another",
		},
		{
			path:          "/something/new/0xXYZb5d4c32345ced77393b3530b1eed0f346429d/another",
			replaceValue:  ":id",
			expectedValue: "/something/new/0xXYZb5d4c32345ced77393b3530b1eed0f346429d/another",
		},
		{
			path:          "/something/1.0.2/21.235.231.232/another/callme@random.org/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ",
			replaceValue:  ":id",
			expectedValue: "/something/:id/:id/another/:id/:id",
		},
	}

	for _, tt := range tests {
		result := urlvalueparser.ReplacePathValues(tt.path, tt.replaceValue)
		assert.Equal(t, tt.expectedValue, result)
	}
}
