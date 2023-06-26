package ihttp

import "github.com/gitkeng/ihttp/util/stringutil"

// ToErrors convert error to Errors or new Errors
func ToErrors(err error, code string, message string, field map[string]any) Errors {

	if err == nil {
		return nil
	}

	if errors, ok := err.(*Errors); ok {
		return *errors
	} else {
		errResp := make(Errors, 0)
		if stringutil.IsEmptyString(message) {
			message = err.Error()
		}

		errResp = append(errResp, Error{
			Code:    code,
			Message: message,
			Fields:  field,
		})
		return errResp
	}
}
