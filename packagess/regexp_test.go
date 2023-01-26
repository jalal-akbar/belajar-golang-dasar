package packagess

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegExp(t *testing.T) {
	str := `e([a-z])o`
	var regex *regexp.Regexp = regexp.MustCompile(str)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eDo"))
	fmt.Println(regex.MatchString("eki"))

	search := regex.FindAllString("eko edo eki eyo", 1)
	fmt.Println(search)
}
