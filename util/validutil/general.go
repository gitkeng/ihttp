package validutil

func IsValidLength(value string, min, max int) bool {
	if len(value) >= min && len(value) <= max {
		return true
	}
	return false
}
