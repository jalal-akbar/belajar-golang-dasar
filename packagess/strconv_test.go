package packagess

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {
	// ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println("ParseInt : ", number)
	} else {
		fmt.Println(err.Error())
	}
	// ParseBool returns the boolean value represented by the string.
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("ParseBool : ", boolean)
	} else {
		fmt.Println(err.Error())
	}
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	convToInt, _ := strconv.Atoi("42")
	fmt.Println("Atoi: ", convToInt)
}

// Bilangan desimal (atau basis 10) adalah sistem penomoran yang menggunakan 10 simbol, yaitu 0 sampai 9.
// Bilangan biner (atau basis 2) adalah sistem penomoran yang hanya menggunakan 2 simbol, yaitu 0 dan 1. Ini adalah dasar angka yang digunakan dalam komputer.
// Bilangan oktal (atau basis 8) adalah sistem penomoran yang menggunakan 8 simbol, yaitu 0 sampai 7.
// Bilangan heksadesimal (atau basis 16) adalah sistem penomoran yang menggunakan 16 simbol, yaitu 0 sampai 9 dan A sampai F.
