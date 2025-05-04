package helpers

import "regexp"

func IsValidUesrname(username string) bool {
	for _, val := range username {
		if (val < 'a' || val > 'z') && (val < 'A' || val > 'Z') && (val < '0' || val > '9') && val != '_' {
			return false
		}
	}
	return true
}

func IsvalidName(name string) bool {
	for _, val := range name {
		if (val < 'a' || val > 'z') && (val < 'A' || val > 'Z') {
			return false
		}
	}
	return true
}

func IsValidEmail(email string) bool {
	reg, err := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(email))
}
