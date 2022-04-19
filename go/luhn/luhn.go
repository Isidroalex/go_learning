package luhn

import (
	"regexp"
	"strings"
)

func Valid(id string) bool {

	if len(id) <= 1 {
		return false
	}
	id = strings.Replace(id, " ", "", -1)
	re := regexp.MustCompile("[^0-9]+")
	if len(re.FindAllString(id, -1)) > 0 {
		return false
	}
	return true
}
