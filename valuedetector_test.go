package urlvalueparser_test

import (
	"github.com/genert/urlvalueparser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewValueDetector_IsValue_ReturnsFalse(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	tests := []struct {
		value string
	}{
		{
			value: "path",
		},
		{
			value: "credentials",
		},
		{
			value: "userId",
		},
		{
			value: "VmFsaWQgc3RyaW5nCg=",
		},
		{
			value: "ZZZZABCDEFABCDE",
		},
		{
			value: "eyJhbGciOiJIUzI-NiIsInR5cCI6__pXVCJ9.e30.Et9HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf5=BADD",
		},
	}

	for _, tt := range tests {
		result := valueDetector.IsValue(tt.value)

		assert.False(t, result, "expected value to be false", tt.value)
	}
}

func TestNewValueDetector_IsValueInteger(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("23")

	assert.True(t, result)
}

func TestNewValueDetector_IsValueUUID(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("64999b0b-fb4b-4ee5-b014-2d0f37ff824e")

	assert.True(t, result)
}

func TestNewValueDetector_IsValueUUIDUpperCase(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("FA029CA5-8885-4A77-AfA3-B79EFFA7f559")

	assert.True(t, result)
}

func TestNewValueDetector_IsValue_JWTToken(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ")

	assert.True(t, result)
}

func TestNewValueDetector_IsValue_Email(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("random@email.com")

	assert.True(t, result)
}

func TestNewValueDetector_IsValue_DataURI(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("data:text/vnd-example+xyz;foo=bar;base64,R0lGODdh")

	assert.True(t, result)
}

func TestNewValueDetector_IsValue_IPV4(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("80.235.87.232")

	assert.True(t, result)
}

func TestNewValueDetector_IsValue_IPV6(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("0:0:0:0:0:ffff:50eb:57e8")

	assert.True(t, result)
}

func TestNewValueDetector_IsValue_SemVer(t *testing.T) {
	valueDetector := urlvalueparser.NewValueDetector()
	result := valueDetector.IsValue("2.0.0")

	assert.True(t, result)
}
