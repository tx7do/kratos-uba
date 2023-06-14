package util

import (
	"fmt"
	"testing"
)

func TestGetCurrentTimeRangeDateString(t *testing.T) {
	fmt.Println(GetTodayRangeDateString())
	fmt.Println(GetCurrentMonthRangeDateString())
	fmt.Println(GetCurrentYearRangeDateString())

	fmt.Println(GetYesterdayRangeDateString())
	fmt.Println(GetLastMonthRangeDateString())
	fmt.Println(GetLastYearRangeDateString())
}

func TestGetCurrentRangeTime(t *testing.T) {
	fmt.Println(GetTodayRangeTime())
	fmt.Println(GetCurrentMonthRangeTime())
	fmt.Println(GetCurrentYearRangeTime())

	fmt.Println(GetYesterdayRangeTime())
	fmt.Println(GetLastMonthRangeTime())
	fmt.Println(GetLastYearRangeTime())
}

func TestGetCurrentTimeRangeTimeString(t *testing.T) {
	fmt.Println(GetTodayRangeTimeString())
	fmt.Println(GetCurrentMonthRangeTimeString())
	fmt.Println(GetCurrentYearRangeTimeString())

	fmt.Println(GetYesterdayRangeTimeString())
	fmt.Println(GetLastMonthRangeTimeString())
	fmt.Println(GetLastYearRangeTimeString())
}
