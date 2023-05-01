package algorithm

import (
	"time"
	"fmt"
	"regexp"
)

// DateToDay converts a date string to a day string.
// @params date: date string in the format of "dd/mm/yyyy"
// @return day: day string in the format of "$day_of_a_week[i]", "" if error
func DateToDay(date string) string {
	day_of_a_week_indonesia := []string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}

	t_date, err := time.Parse("02/01/2006", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	return day_of_a_week_indonesia[t_date.Weekday()]
}

// ExtractDates extract dates from a string.
// @params text: string to be extracted
// @return date: date string in the format of "dd/mm/yyyy", may be empty or multiple
func ExtractDates(text string) []string {
	Lower(&text) // not case sensitive
	regex_pattern := `(\d{1,2})[\/\-\\](\d{1,2})[\/\-\\](\d{4})`

	re := regexp.MustCompile(regex_pattern)
	return re.FindAllString(text, -1)
}