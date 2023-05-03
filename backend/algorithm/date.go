package algorithm

import (
	"fmt"
	"regexp"
	"time"
)

// DateToDay converts a date string to a day string.
// @params date: date string in the format of "dd/mm/yyyy"
// @return day: day string in the format of "$day_of_a_week[i]", "" if error
func DateToDay(date string) string {
	day_of_a_week_indonesia := []string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}

	t_date1, err1 := time.Parse("02/01/2006", date)
	t_date2, err2 := time.Parse("2/01/2006", date)
	t_date3, err3 := time.Parse("02/1/2006", date)
	t_date4, err4 := time.Parse("2/1/2006", date)

	if err1 != nil && err2 != nil && err3 != nil && err4 != nil {
		fmt.Println("Error: DateToDay: Invalid date format")
		return ""
	}

	var t_date time.Time
	if err1 == nil {
		t_date = t_date1
	} else if err2 == nil {
		t_date = t_date2
	} else if err3 == nil {
		t_date = t_date3
	} else if err4 == nil {
		t_date = t_date4
	}

	return day_of_a_week_indonesia[t_date.Weekday()]
}

// ExtractDates extract date from a string.
// @params text: string to be extracted
// @return date: date string in the format of "dd/mm/yyyy", may be empty or multiple
func ExtractDates(text string) []string {
	Lower(&text) // not case sensitive
	Trim(&text)  // remove whitespace so it's easier to check
	regex_pattern := `(\d{1,2})[\/\-\\](\d{1,2})[\/\-\\](\d{4})`

	re := regexp.MustCompile(regex_pattern)
	return re.FindAllString(text, -1)
}
