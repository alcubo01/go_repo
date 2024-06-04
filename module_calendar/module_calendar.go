package module_calendar

import (
	"fmt"
	"time"

	"github.com/alcubo01/go_repo/module_golors"
)

var colorWeekend = module_golors.Rgb(255, 0, 0)

// Calendar represents a calendar.
type Calendar struct {
	year          int
	month         time.Month
	printableData []rune
}

func addStringToRunes(runes []rune, newStr string) []rune {
	arr := []rune(newStr)
	for i := 0; i < len(arr); i++ {
		runes = append(runes, arr[i])
	}
	return runes
}

func addRunes(runes []rune, newRunes []rune) []rune {
	for i := 0; i < len(newRunes); i++ {
		runes = append(runes, newRunes[i])
	}
	return runes
}

// NewCalendar creates a new Calendar instance.
func NewCalendar(year int, month time.Month) *Calendar {
	// Crear un objeto de calendario
	calendar := &Calendar{
		year:          year,
		month:         month,
		printableData: []rune{'\n'},
	}

	NextMonthFirstDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	MonthFirstDay := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
	totalDays := int(NextMonthFirstDay.Sub(MonthFirstDay).Hours() / 24)

	firstWeekday := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()
	_, week := MonthFirstDay.ISOWeek()

	calendar.printableData = addStringToRunes(calendar.printableData, fmt.Sprintf("%s:\n\n", month))

	for day := 1; day <= totalDays; day++ {
		weekday := day + int(firstWeekday)

		if day == 1 || (weekday-1)%7 == 0 {
			calendar.printableData = addStringToRunes(calendar.printableData, fmt.Sprintf("week:%2d ", week))
			week = week + 1
		}

		if day == 1 {
			for i := 0; i < int(firstWeekday); i++ {
				calendar.printableData = addStringToRunes(calendar.printableData, "    ")
			}
		}

		if (weekday%7 == 0) || ((weekday+1)%7 == 0) {
			calendar.printableData = addStringToRunes(
				calendar.printableData, module_golors.ColoredText(
					colorWeekend, fmt.Sprintf(" %2d ", day)))
			if weekday%7 == 0 {
				calendar.printableData = addStringToRunes(calendar.printableData, "\n")
			}
		} else {
			calendar.printableData = addStringToRunes(calendar.printableData, fmt.Sprintf(" %2d ", day))
		}
	}

	calendar.printableData = addStringToRunes(calendar.printableData, "\n")
	return calendar
}

// Print prints the Calendar.
func GenerateCalendar(year int, month int, code int) string {

	if code == -1 {
		calendar := NewCalendar(year, time.Month(month))
		return string(calendar.printableData)
	} else if code == -3 {
		result := []rune{}
		for i := -1; i < 2; i++ {
			calendar := NewCalendar(year, time.Month(month+i))
			result = addRunes(result, calendar.printableData)
		}
		return string(result)
	}
	return "Invalid arguments provided"
}
