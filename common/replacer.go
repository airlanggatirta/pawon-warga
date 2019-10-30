package common

import (
	"fmt"
	"strings"
)

func ReplacePhoneNumberMask(phoneNumber string) string {
	num := phoneNumber
	num = strings.TrimLeft(num, "0")
	num = strings.TrimLeft(num, "+62")

	if num != phoneNumber {
		num = fmt.Sprintf("62%s", num)
	}

	return num
}
