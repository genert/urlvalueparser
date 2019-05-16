package urlvalueparser

import (
	"github.com/google/uuid"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type (
	valueDetector struct{}

	ValueDetector interface {
		IsValue(value string) bool
	}
)

const (
	Email       string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	Date        string = `/^(\d{2}|\d{4})\-\d\d\-\d\d$/`
	Ethereum    string = "^0x[0-9a-fA-F]{40}$"
	JWTToken    string = `^([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`
	Base64      string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	Latitude    string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude   string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	Semver      string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	Hexadecimal string = "^[0-9a-fA-F]+$"
	Hexcolor    string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RGBcolor    string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	DataURI     string = "^data:.+\\/(.+);base64$"
)

var (
	rxEmail       = regexp.MustCompile(Email)
	rxDate        = regexp.MustCompile(Date)
	rxEthereum    = regexp.MustCompile(Ethereum)
	rxJWTToken    = regexp.MustCompile(JWTToken)
	rxBase64      = regexp.MustCompile(Base64)
	rxSemver      = regexp.MustCompile(Semver)
	rxLatitude    = regexp.MustCompile(Latitude)
	rxLongitude   = regexp.MustCompile(Longitude)
	rxHexadecimal = regexp.MustCompile(Hexadecimal)
	rxHexcolor    = regexp.MustCompile(Hexcolor)
	rxRGBcolor    = regexp.MustCompile(RGBcolor)
	rxDataURI     = regexp.MustCompile(DataURI)
)

func NewValueDetector() ValueDetector {
	return &valueDetector{}
}

func (detector *valueDetector) isEmail(value string) bool {
	return rxEmail.MatchString(value)
}

func (detector *valueDetector) isSemVer(value string) bool {
	return rxSemver.MatchString(value)
}

func (detector *valueDetector) isDate(value string) bool {
	r := regexp.MustCompile(`/^(\d{2}|\d{4})\-\d\d\-\d\d$/`)
	return r.MatchString(value)
}

func (detector *valueDetector) isUUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}

func (detector *valueDetector) isNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}

	return false
}

func (detector *valueDetector) isJWTToken(value string) bool {
	return rxJWTToken.MatchString(value)
}

func (detector *valueDetector) isEthereumAddress(value string) bool {
	return rxEthereum.MatchString(value)
}

func (detector *valueDetector) isBase64(value string) bool {
	return rxBase64.MatchString(value)
}

// IsIP checks if a string is either IP version 4 or 6.
func (detector *valueDetector) isIP(str string) bool {
	return net.ParseIP(str) != nil
}

// IsDataURI checks if a string is base64 encoded data URI such as an image
func (detector *valueDetector) isDataURI(str string) bool {
	dataURI := strings.Split(str, ",")

	if !rxDataURI.MatchString(dataURI[0]) {
		return false
	}

	return detector.isBase64(dataURI[1])
}

// IsValue checks if value is recognized by value detectors. Empty string returns true.
func (detector *valueDetector) IsValue(value string) bool {
	if value == "" {
		return true
	}

	if detector.isNumber(value) {
		return true
	}

	if detector.isUUID(value) {
		return true
	}

	if detector.isDataURI(value) {
		return true
	}

	if detector.isIP(value) {
		return true
	}

	if detector.isEmail(value) {
		return true
	}

	if detector.isSemVer(value) {
		return true
	}

	if detector.isEthereumAddress(value) {
		return true
	}

	if detector.isJWTToken(value) {
		return true
	}

	return false
}
