package time

import (
	"dayCounter/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimePackage_IsLastMonthInYear(t *testing.T) {

	testcases := []struct {
		desc string
		month int
		result bool
	}{
		{
			desc:   "a",
			month:  1,
			result: false,
		},
		{
			desc:   "b",
			month:  12,
			result:  true,
		},
	}

	for _, tc := range testcases{
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, Calculator.IsLastMonthInYear(tc.month),tc.result)
		})
	}
}

func TestTimePackage_IncreaseDateByOneDay(t *testing.T) {
	testcases := []struct{
		desc string
		request *entity.TimeRequest
		result * entity.TimeRequest
	}{
		{
			desc:    "a",
			request:  & entity.TimeRequest{
				Year:  1989,
				Month: 1,
				Day:   3,
			},
			result: &entity.TimeRequest{
				Year:  1989,
				Month: 1,
				Day:   4,
			},
		},
		{
			desc:    "b",
			request:  &entity.TimeRequest{
				Year:  1989,
				Month: 1,
				Day:   31,
			},
			result: &entity.TimeRequest{
				Year:  1989,
				Month: 2,
				Day:   1,
			},
		},
	}

	for _ , tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			Calculator.IncreaseDateByOneDay(tc.request)
			assert.Equal(t, tc.request,tc.result)
		})
	}
}

func TestTimePackage_IsLastDayInMonth(t *testing.T) {

	testcases := []struct{
		desc string
		request entity.TimeRequest
		result bool
	}{
		{
			desc:    "a",
			request: entity.TimeRequest{
				Year:  1989,
				Month: 1,
				Day:   31,
			},
			result:  true,
		},
		{
			desc:    "b",
			request: entity.TimeRequest{
				Year:  1989,
				Month: 9,
				Day:   1,
			},
			result:  false,
		},
	}

	for _, tc := range testcases{
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, Calculator.IsLastDayInMonth(tc.request),tc.result)
		})
	}
}

func TestTimePackage_IsLeapYear(t *testing.T) {
	testcase := []struct{
		desc string
		year int
		result bool
	}{
		{
			desc:   "a",
			year:   1989,
			result: false,
		},
		{
			desc:   "b",
			year:   2020,
			result: true,
		},
	}

	for _ , tc := range testcase{
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, Calculator.IsLeapYear(tc.year),tc.result)
		})
	}
}

func TestTimePackage_IsBefore(t *testing.T) {
	testcases := []struct{
		desc string
		date1 entity.TimeRequest
		date2 entity.TimeRequest
		result bool
	}{
		{
			desc:   "a",
			date1:  entity.TimeRequest{
				Year:  1989,
				Month: 10,
				Day:   1,
			},
			date2:  entity.TimeRequest{
				Year:  1989,
				Month: 9,
				Day:   30,
			},
			result: false,
		},
		{
			desc:   "b",
			date1:  entity.TimeRequest{
				Year:  1989,
				Month: 10,
				Day:   1,
			},
			date2:  entity.TimeRequest{
				Year:  1989,
				Month: 11,
				Day:   30,
			},
			result: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, Calculator.IsBefore(tc.date1,tc.date2),tc.result)
		})
	}
}

func TestTimePackage_CountDays(t *testing.T) {

	testcase := []struct{
		desc string
		start entity.TimeRequest
		end entity.TimeRequest
		result int
	}{
		{
			desc:   "a",
			start:  entity.TimeRequest{
				Year:  1989,
				Month: 1,
				Day:   5,
			},
			end:    entity.TimeRequest{
				Year:  1989,
				Month: 1,
				Day:   19,
			},
			result: 14,
		},
	}

	for _, tc := range testcase{
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, Calculator.CountDays(tc.start,tc.end),tc.result)
		})
	}
}