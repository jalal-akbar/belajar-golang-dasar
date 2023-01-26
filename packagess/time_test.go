package packagess

import (
	"fmt"
	"testing"
	"time"
)

func TestTiime(t *testing.T) {
	// Now return the current local time
	now := time.Now()
	fmt.Println("Day : ", now.Day())
	fmt.Println("Month : ", now.Month())
	fmt.Println("Year : ", now.Year())
	fmt.Println("Day : ", now.Day())
	fmt.Println("Location : ", now.Location())

	utc := time.Date(2010, 10, 10, 10, 10, 10, 10, time.UTC)
	hour, min, sec := now.Clock()
	fmt.Println("Utc : ", utc)
	fmt.Printf("Hour : %v\n", hour)
	fmt.Printf("Minutes : %v\n", min)
	fmt.Printf("Second : %v\n", sec)

	// 2006-01-02-  T15:04:05Z07:00
	var layout = "2006-01-02"
	parsing, _ := time.Parse(layout, "1996-10-02")
	fmt.Println("Parsing : ", parsing)

}
