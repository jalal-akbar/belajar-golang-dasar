package packagess

import (
	"fmt"
	"reflect"
	"testing"
)

type sample struct {
	Name string `required:"true" max:"10"` //Struct Tag
	Age  int
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != "" // singkat
			// v := reflect.ValueOf(data).Field(i).Interface()
			// if v == "" {
			// 	return false
			// }
		}
	}
	return true
}

func TestReflect(t *testing.T) {
	samples := sample{
		Name: "Jalal",
		Age:  26,
	}
	sampleType := reflect.TypeOf(samples)
	// Field
	fmt.Println("Sample Field for Name : ", sampleType.Field(0).Name)
	fmt.Println("Sample Field for Age : ", sampleType.Field(1).Name)
	// structTag
	fmt.Println("required : ", sampleType.Field(0).Tag.Get("required"))
	fmt.Println("max : ", sampleType.Field(0).Tag.Get("max"))
	fmt.Println("min : ", sampleType.Field(0).Tag.Get("min"))

	fmt.Println("is valid : ", isValid(samples))

}
