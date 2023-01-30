package utils

import (
	"regexp"
)

type Regexp regexp.Regexp

func DetectCarrier(trackingCode string) string {
	carrier := ""
	uspsRegex := regexp.MustCompile(`^[0-9]{22}$`)
	upsRegex := regexp.MustCompile(`^[1-9][0-9]{14,18}$`)
	fedexRegex := regexp.MustCompile(`^[0-9]{12}$`)

	if uspsRegex.MatchString(trackingCode) {
		carrier = "USPS"
	} else if upsRegex.MatchString(trackingCode) {
		carrier = "UPS"
	} else if fedexRegex.MatchString(trackingCode) {
		carrier = "Fedex"
	} else {
		carrier = "Unsupported"
	}
	return carrier
}
