package time

import (
	"dayCounter/entity"
)

//CountDays will count days between 2 dates
func (t timePackage) CountDays(start, end entity.TimeRequest) int {
	days := 0
	for t.IsBefore(start,end){
		days++
		t.IncreaseDateByOneDay(&start)
	}
	return days
}

//IsBefore will check if the first date is before second date
func (t timePackage) IsBefore(date1 entity.TimeRequest, date2 entity.TimeRequest) bool {
	if date1.Year < date2.Year{
		return true
	}
	if date1.Year == date2.Year {
		if date1.Month < date2.Month {
			return true
		}
		if date1.Month == date2.Month {
			if date1.Day < date2.Day {
				return true
			}
			return false
		}
		return false
	}
	return false
}


//IsLeapYear will check whether the year is leap year or not
func (t timePackage) IsLeapYear(year int) bool {
	return (year%4 == 0) && ((year%100 != 0) || (year%400 == 0))
}

//IsLastDayInMonth will check whether the day is the last day of month or not
func (t timePackage) IsLastDayInMonth(request entity.TimeRequest) bool {
	leapYear := false
	if t.IsLeapYear(request.Year) {
		leapYear = true
	}

	return (request.Month == 1 && request.Day == 31) ||
		(leapYear && request.Month == 2 && request.Day == 29) || (!leapYear && request.Month == 2 && request.Day == 28) ||
		(request.Month == 3 && request.Day == 31) ||
		(request.Month == 4 && request.Day == 30) ||
		(request.Month == 5 && request.Day == 31) ||
		(request.Month == 6 && request.Day == 30) ||
		(request.Month == 7 && request.Day == 31) ||
		(request.Month == 8 && request.Day == 31) ||
		(request.Month == 9 && request.Day == 30) ||
		(request.Month == 10 && request.Day == 31) ||
		(request.Month == 11 && request.Day == 30) ||
		(request.Month == 12 && request.Day == 31)

}

//IncreaseDateByOneDay will increase date one by one
func (t timePackage) IncreaseDateByOneDay(request *entity.TimeRequest) {

	if t.IsLastDayInMonth(*request) {
		if t.IsLastMonthInYear(request.Month) {
			request.Month = 1
			request.Day= 1
			request.Year++
		} else {
			request.Day = 1
			request.Month++
		}
	} else {
		request.Day++
	}
}

//IsLastMonthInYear checks whether it's the last month of the year
func (t timePackage) IsLastMonthInYear(month int) bool {
	if month == 12{
		return true
	}
	return false
}
