package common

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

var (
	bm = bluemonday.StrictPolicy()
)

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateTimestampRange(month int64, year int64) (dateStart int64, dateEnd int64, err error) {
	layout := "2006-01-02 15:04:05"

	startMonth := fmt.Sprintf("%02d", month)
	endMonth := fmt.Sprintf("%02d", month+1)

	start := fmt.Sprintf("%d-%s-01 00:00:00", year, startMonth)
	end := fmt.Sprintf("%d-%s-01 00:00:00", year, endMonth)

	t, err := time.Parse(layout, start)
	if err != nil {
		return 0, 0, err
	}
	dateStart = t.Unix()

	t2, err := time.Parse(layout, end)
	if err != nil {
		return 0, 0, err
	}
	dateEnd = t2.Unix()

	return
}

func StripHTML(content string) (result string) {
	result = bm.Sanitize(content)

	return
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func MatchesAllowedPhoneNumberPrefix(phoneNumber string) (matched bool, err error) {
	matched, err = regexp.MatchString("^(\\+62|0|62)\\d{10,16}$", phoneNumber)
	return
}
