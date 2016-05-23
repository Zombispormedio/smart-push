package utils

import (
    "reflect"
)

func Contains(array []interface{}, elem interface{})bool{
   var contains bool

	s := reflect.ValueOf(array)

	for i := 0; i < s.Len(); i++ {
		value:=s.Index(i).Interface()
        
        if value==elem{
            contains=true
            break;
        }
		
	}
    return contains
}