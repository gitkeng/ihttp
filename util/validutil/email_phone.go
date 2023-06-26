package validutil

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/gitkeng/ihttp/util/phoneutil"
	"github.com/gitkeng/ihttp/util/stringutil"
)

var (
	ErrEmailMxRecordNotFound = errors.New("not found email mx record")
	ErrEmailInvalidFormat    = errors.New("email is invalid format")
)

func IsValidEmailAddress(emailAddress string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(emailAddress) < 3 && len(emailAddress) > 254 {
		return false
	}
	return emailRegex.MatchString(emailAddress)
}

func ValidEmailAddress(emailAddress string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(emailAddress) < 3 && len(emailAddress) > 254 {
		return ErrEmailInvalidFormat
	}
	if !emailRegex.MatchString(emailAddress) {
		return ErrEmailInvalidFormat
	}
	parts := strings.Split(emailAddress, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return ErrEmailMxRecordNotFound
	}
	return nil
}

func IsValidPhoneNumber(countryCode, phoneNumber string) bool {
	iso := phoneutil.GetISO3166ByCountry(countryCode)
	if stringutil.IsEmptyString(iso.CountryCode) {
		return false
	}
	for _, phoneLength := range iso.PhoneNumberLengths {
		regex := fmt.Sprintf("^([0-9]){%d,%d}$", phoneLength, phoneLength+1)
		phoneRegex := regexp.MustCompile(regex)
		if phoneRegex.MatchString(phoneNumber) {
			return true
		}
	}
	return false

}

func IsValidMobileNumber(countryCode, mobile string) bool {
	iso := phoneutil.GetISO3166ByCountry(countryCode)
	if stringutil.IsEmptyString(iso.CountryCode) {
		return false
	}
	for _, phoneLength := range iso.PhoneNumberLengths {
		regex := fmt.Sprintf("^([0-9]){%d,%d}$", phoneLength, phoneLength+1)
		phoneRegex := regexp.MustCompile(regex)
		prefix := ""
		if phoneRegex.MatchString(mobile) {
			if len(mobile) == phoneLength+1 {
				prefix = mobile[1:2]
			} else {
				prefix = mobile[:1]
			}
			for _, mobilePrefix := range iso.MobileBeginWith {
				if prefix == mobilePrefix {
					return true
				}
			}

		}
	}
	return false
}
