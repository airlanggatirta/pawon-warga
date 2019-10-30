package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ParsePageToken(pageToken string) (token map[string]uint64) {
	pToken := strings.Split(pageToken, "_")
	token = make(map[string]uint64)
	var date string
	var id string

	if len(pToken) != 2 {
		return
	}

	date = pToken[0]
	id = pToken[1]

	dateInt, err := strconv.ParseUint(date, 10, 64)
	if err != nil {
		return
	}

	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}

	token["date"] = dateInt
	token["id"] = idInt

	return
}

func CreatePageToken(arrayData interface{}, dataLimit int) (nextToken string) {
	voData := reflect.ValueOf(arrayData)
	if voData.Kind() != reflect.Slice {
		return
	}

	if !(voData.IsValid() && voData.Len() >= dataLimit) {
		return
	}

	lastRow := voData.Index(voData.Len() - 1)
	nextTimestamp := lastRow.FieldByName("Created").Int()
	nextID := lastRow.FieldByName("ID").Uint()
	nextToken = fmt.Sprintf("%d_%d", nextTimestamp, nextID)
	return
}
