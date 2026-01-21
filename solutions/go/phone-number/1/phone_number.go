package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func Number(phoneNumber string) (string, error) {
	r := regexp.MustCompile(`[0-9]+`)
	parsed := r.FindAllString(phoneNumber, -1)
	result := strings.Join(parsed, "")

	numLength := utf8.RuneCountInString(result)

	//Check the length
	if numLength < 10 || numLength > 11 {
		return "", errors.New("Number must be 10 or 11 digits.")
	}

	areaCode := ""
	exchangeCode := ""
	subscriberNumber := ""

	if numLength == 11 {
		if result[0] != '1' {
			return "", errors.New("Invalid country code.")
		}

		areaCode = string(result[1:4])
		exchangeCode = string(result[4:7])
		subscriberNumber = string(result[7:])
	} else {
		areaCode = string(result[0:3])
		exchangeCode = string(result[3:6])
		subscriberNumber = string(result[6:])
	}

	if areaCode[0] < '2' {
		return "", errors.New("Invalid area code.")
	}

	if exchangeCode[0] < '2' {
		return "", errors.New("Invalid exchange code.")
	}

	return fmt.Sprintf("%s%s%s", areaCode, exchangeCode, subscriberNumber), nil
}

func AreaCode(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return cleanNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", cleanNumber[0:3], cleanNumber[3:6], cleanNumber[6:]), nil
}
