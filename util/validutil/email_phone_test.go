package validutil_test

import (
	"testing"

	"github.com/gitkeng/ihttp/util/validutil"
	"github.com/magiconair/properties/assert"
)

func TestValidEmailAddress(t *testing.T) {
	dataTests := []struct {
		EmailAddress string
		Expect       error
	}{
		{
			EmailAddress: "gitkeng@mail.com",
			Expect:       nil,
		},
		{
			EmailAddress: "gitkeng@gec.co",
			Expect:       validutil.ErrEmailMxRecordNotFound,
		},
		{
			EmailAddress: "gitkenggec.co",
			Expect:       validutil.ErrEmailInvalidFormat,
		},
		{
			EmailAddress: "somchai@gmail.com",
			Expect:       nil,
		},
		{
			EmailAddress: "somchai@samui.com",
			Expect:       validutil.ErrEmailMxRecordNotFound,
		},
	}

	for _, data := range dataTests {
		result := validutil.ValidEmailAddress(data.EmailAddress)
		t.Logf("email: %s, result: %v", data.EmailAddress, result)
		assert.Equal(t, result, data.Expect)
	}

}

func TestIsValidMobileNumber(t *testing.T) {
	dataTests := []struct {
		CountryCode  string
		MobileNumber string
		Expect       bool
	}{
		{
			CountryCode:  "TH",
			MobileNumber: "027016161",
			Expect:       false,
		},
		{
			CountryCode:  "TH",
			MobileNumber: "0877177177",
			Expect:       true,
		},
		{
			CountryCode:  "TH",
			MobileNumber: "877177177",
			Expect:       true,
		},
		{
			CountryCode:  "TH",
			MobileNumber: "0677177177",
			Expect:       true,
		},
		{
			CountryCode:  "TH",
			MobileNumber: "08771771779",
			Expect:       false,
		},
		{
			CountryCode:  "TH",
			MobileNumber: "0",
			Expect:       false,
		},
		{
			CountryCode:  "TH",
			MobileNumber: "06512345a7",
			Expect:       false,
		},
	}

	for _, data := range dataTests {
		result := validutil.IsValidMobileNumber(data.CountryCode, data.MobileNumber)
		t.Logf("mobile: %s, result: %t", data.MobileNumber, result)
		assert.Equal(t, result, data.Expect)

	}

}
func TestIsValidPhoneNumber(t *testing.T) {
	dataTests := []struct {
		CountryCode string
		PhoneNumber string
		Expect      bool
	}{
		{
			CountryCode: "TH",
			PhoneNumber: "027016161",
			Expect:      true,
		},
		{
			CountryCode: "TH",
			PhoneNumber: "0877177177",
			Expect:      true,
		},
		{
			CountryCode: "TH",
			PhoneNumber: "877177177",
			Expect:      true,
		},
		{
			CountryCode: "TH",
			PhoneNumber: "0677177177",
			Expect:      true,
		},
		{
			CountryCode: "TH",
			PhoneNumber: "08771771779",
			Expect:      false,
		},
		{
			CountryCode: "TH",
			PhoneNumber: "01333443",
			Expect:      false,
		},
		{
			CountryCode: "TH",
			PhoneNumber: "02123d9911",
			Expect:      false,
		},
	}

	for _, data := range dataTests {
		result := validutil.IsValidPhoneNumber(data.CountryCode, data.PhoneNumber)
		t.Logf("phoneno: %s, result: %t", data.PhoneNumber, result)
		assert.Equal(t, result, data.Expect)
	}

}
