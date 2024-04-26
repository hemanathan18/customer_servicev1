package validations

import (
	"regexp"
	"strconv"
	// "github.com/google/libphonenumber"
)

func IsValidZipCode(zipCode int64) bool {
	zipCodeStr := strconv.Itoa(int(zipCode))
	pattern := `^\d{6}$`
	matched, _ := regexp.MatchString(pattern, zipCodeStr)
	return matched
}

func IsValidPhoneNumber(phoneNumber int64) bool {
	phoneNumberStr := strconv.Itoa(int(phoneNumber))
	pattern := `^\d{4,12}$`
	matched, _ := regexp.MatchString(pattern, phoneNumberStr)
	return matched
}

// need to alter according to the common standards
func IsValidCountryCode(countryCode int64) bool {
	countryCodeStr := strconv.Itoa(int(countryCode))
	pattern := `^\d{1,3}$`
	matched, _ := regexp.MatchString(pattern, countryCodeStr)
	return matched
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
