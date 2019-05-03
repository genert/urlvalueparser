package urlvalueparser

import (
	"regexp"
	"strconv"
)

type valueDetector struct {
	minHexLength int
}

type ValueDetector interface {
	IsValue(value string) bool
}

func NewValueDetector() ValueDetector {
	return &valueDetector{}
}

func (detector *valueDetector) isDate(value string) bool {
	r := regexp.MustCompile(`/^(\d{2}|\d{4})\-\d\d\-\d\d$/`)
	return r.MatchString(value)
}

func (detector *valueDetector) isUUID(value string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(value)
}

func (detector *valueDetector) isNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}

	return false
}

func (detector *valueDetector) isJWTToken(value string) bool {
	r := regexp.MustCompile(`^([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`)
	return r.MatchString(value)
}

func (detector *valueDetector) IsValue(value string) bool {
	if detector.isNumber(value) {
		return true
	}

	if detector.isUUID(value) {
		return true
	}

	if detector.isJWTToken(value) {
		return true
	}

	return false
}
