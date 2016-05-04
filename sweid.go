// Package sweid validates Swedish Personal Identification Number (personnummer)
// and Organisational Number (organisationsnummer).
//
// More information about these identifiers can be found at the
// Swedish Tax Agency, http://www.skatteverket.se.
package sweid

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	pinRegExp = "[0-9]{2}[0-1][0-9][0-3][0-9]{5}"
	onRegExp  = "[0-9]{2}[2-9][0-9]{7}"
)

// Validate checksum according to Luhn algorithm, https://en.wikipedia.org/wiki/Luhn_algorithm.
func luhn(s string) int {
	var t = []int{2, 1, 2, 1, 2, 1, 2, 1, 2}

	var sum int
	for i, c := range s {
		j := int(c-'0') * t[i]
		if j >= 10 {
			sum++
		}
		sum += j % 10
	}
	return 10 - sum%10
}

// validate parses and validate the number against a regular expression and then
// calculates the checksum.
func validate(num, exp string) bool {
	// Remove one -
	num = strings.Replace(num, "-", "", 1)

	// Remove century part if there are 12 characters left in identifier
	if len(num) == 12 {
		num = num[2:]
	}

	// There should be only 10 characters left by this time
	if len(num) != 10 {
		return false
	}

	// Validate against regular expression
	r, err := regexp.Compile(exp)
	if err != nil {
		return false
	}
	if !r.MatchString(num) {
		return false
	}

	// Extract the checksum and convert to int
	chksm, err := strconv.Atoi(num[9:])
	if err != nil {
		return false
	}

	// Remove the checksum before running luhn
	num = num[:9]

	// Validate checksum in identifier matches the luhn calculated one.
	return chksm == luhn(num)
}

// ValidPN validates a Swedish personal identification number (personnummer).
// Definitions can be found at http://www.skatteverket.se (search for personnummer).
func ValidPN(pn string) bool {
	return validate(pn, pinRegExp)
}

// ValidON validates a Swedish organisational number (organisationsnummer).
// Definitions can be found at http://www.skatteverket.se (search for organisationsnummer).
func ValidON(on string) bool {
	return validate(on, onRegExp)
}
