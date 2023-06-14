package util

import "time"

// GetYesterdayRangeTime 获取区间时间 - 昨天
func GetYesterdayRangeTime() (time.Time, time.Time) {
	now := time.Now().AddDate(0, 0, -1)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startDate, endDate
}

// GetTodayRangeTime 获取区间时间 - 今天
func GetTodayRangeTime() (time.Time, time.Time) {
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startDate, endDate
}

// GetLastMonthRangeTime 获取区间时间 - 上个月
func GetLastMonthRangeTime() (time.Time, time.Time) {
	now := time.Now().AddDate(0, -1, 0)
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetCurrentMonthRangeTime 获取区间时间 - 本月
func GetCurrentMonthRangeTime() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetCurrentYearRangeTime 获取区间时间 - 今年
func GetCurrentYearRangeTime() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(1, 0, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetLastYearRangeTime 获取区间时间 - 去年
func GetLastYearRangeTime() (time.Time, time.Time) {
	now := time.Now().AddDate(-1, 0, 0)
	firstDay := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(1, 0, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetTodayRangeDateString 获取区间时间内的日期字符串（比如：2023-05-23 2023-05-23） - 今天
func GetTodayRangeDateString() (string, string) {
	firstDay, lastDay := GetTodayRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetYesterdayRangeDateString 获取区间时间内的日期字符串（比如：2023-05-23 2023-05-23） - 昨天
func GetYesterdayRangeDateString() (string, string) {
	firstDay, lastDay := GetYesterdayRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentMonthRangeDateString 获取区间时间内的日期字符串（比如：2023-05-23 2023-05-23） - 本月
func GetCurrentMonthRangeDateString() (string, string) {
	firstDay, lastDay := GetCurrentMonthRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetLastMonthRangeDateString 获取区间时间内的日期字符串（比如：2023-05-23 2023-05-23） - 上个月
func GetLastMonthRangeDateString() (string, string) {
	firstDay, lastDay := GetLastMonthRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentYearRangeDateString 获取区间时间内的日期字符串（比如：2023-05-23 2023-05-23） - 今年
func GetCurrentYearRangeDateString() (string, string) {
	firstDay, lastDay := GetCurrentYearRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetLastYearRangeDateString 获取区间时间内的日期字符串（比如：2023-05-23 2023-05-23） - 去年
func GetLastYearRangeDateString() (string, string) {
	firstDay, lastDay := GetLastYearRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetYesterdayRangeTimeString 获取区间时间内的时间字符串（比如：2023-05-23 00:00:00 2023-05-23 23:59:59） - 昨天
func GetYesterdayRangeTimeString() (string, string) {
	firstDay, lastDay := GetYesterdayRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetTodayRangeTimeString 获取区间时间内的时间字符串（比如：2023-05-23 00:00:00 2023-05-23 23:59:59） - 今天
func GetTodayRangeTimeString() (string, string) {
	firstDay, lastDay := GetTodayRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetLastMonthRangeTimeString 获取区间时间内的时间字符串（比如：2023-05-23 00:00:00 2023-05-23 23:59:59） - 上个月
func GetLastMonthRangeTimeString() (string, string) {
	firstDay, lastDay := GetLastMonthRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetCurrentMonthRangeTimeString 获取区间时间内的时间字符串（比如：2023-05-23 00:00:00 2023-05-23 23:59:59） - 本月
func GetCurrentMonthRangeTimeString() (string, string) {
	firstDay, lastDay := GetCurrentMonthRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetLastYearRangeTimeString 获取区间时间内的时间字符串（比如：2023-05-23 00:00:00 2023-05-23 23:59:59） - 去年
func GetLastYearRangeTimeString() (string, string) {
	firstDay, lastDay := GetLastYearRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetCurrentYearRangeTimeString 获取区间时间内的时间字符串（比如：2023-05-23 00:00:00 2023-05-23 23:59:59） - 今年
func GetCurrentYearRangeTimeString() (string, string) {
	firstDay, lastDay := GetCurrentYearRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}
