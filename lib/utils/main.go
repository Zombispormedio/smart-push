package utils

import (
	"reflect"
	"strings"

	"strconv"
)

func Contains(array []interface{}, elem interface{}) bool {
	var contains bool

	s := reflect.ValueOf(array)

	for i := 0; i < s.Len(); i++ {
		value := s.Index(i).Interface()

		if value == elem {
			contains = true
			break
		}

	}
	return contains
}

type Timestampkey struct {
	Key       string
	Timestamp int64
	Index     int
}

func GetMaxTimestampKey(keys []string) *Timestampkey {
	var result string
	var timestamp int64
	var index int
	max := &Timestampkey{}

	for i, v := range keys {
		elems := strings.Split(v, ":")
		str := elems[2]

		if str != "" {

			t, _ := strconv.ParseInt(str, 10, 64)

			if t > timestamp {
				timestamp = t
				result = v
				index = i
			}

		}

	}

	max.Index = index
	max.Timestamp = timestamp
	max.Key = result

	return max
}

func GetMinTimestampKey(keys []string) *Timestampkey {
	var result string
	var timestamp int64
	var index int
	min := &Timestampkey{}

	for i, v := range keys {
		elems := strings.Split(v, ":")
		str := elems[2]

		if str != "" {

			t, _ := strconv.ParseInt(str, 10, 64)

			if timestamp == 0 {
				timestamp = t
				result = v
				index = i
			} else {
				if t < timestamp {
					timestamp = t
					result = v
					index = i
				}
			}

		}

	}

	min.Index = index
	min.Timestamp = timestamp
	min.Key = result

	return min
}
