package time

import "dayCounter/entity"

var (
	Calculator timeInterface = &timePackage{}

)
type timeInterface interface {
	IsBefore(entity.TimeRequest, entity.TimeRequest)bool
	IsLeapYear(int)bool
	IsLastDayInMonth(entity.TimeRequest)bool
	IncreaseDateByOneDay(*entity.TimeRequest)
	CountDays(entity.TimeRequest, entity.TimeRequest)int
	IsLastMonthInYear(int)bool
}

type timePackage struct {}















