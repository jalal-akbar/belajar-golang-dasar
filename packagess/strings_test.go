package packagess

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	var (
		join []string = []string{"jalal", "muh", "akbar"}
	)
	// Clone returns a fresh copy of s.
	fmt.Println("Clone : ", strings.Clone("akbar"))
	// Compare returns an integer comparing two strings lexicographically.
	// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	fmt.Println("Compare : ", strings.Compare("j", "k"))
	// HasPrefix tests whether the string s begins with prefix.
	fmt.Println("Has Prefix : ", strings.HasPrefix("jalaluddin muh akbar", "akbar"))
	// HasSuffix tests whether the string s ends with suffix.
	fmt.Println("Has Suffix : ", strings.HasSuffix("jalaluddin muh akbar", "akbar"))
	// Join concatenates the elements of its first argument to create a single string.
	// The separator string sep is placed between elements in the resulting string.
	fmt.Println("Join : ", strings.Join(join, "+"))
	// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
	fmt.Println("Split : ", strings.Split("jalaluddin muh akbar", " "))
	// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
	fmt.Println("Trim : ", strings.Trim("++++Jalaluddin Muh Akbar+++++", "+"))
	// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
	fmt.Println("ReplaceAll : ", strings.ReplaceAll("jalal muh jalal muh", "muh", "akbar"))
}
