package parser

import (
	"dayCounter/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	testcases := []struct{
		desc string
		input string
		date1 *entity.TimeRequest
		date2 *entity.TimeRequest
	}{
		{
			desc:  "a",
			input: "2/6/1989 to 4/6/1989",
			date1: &entity.TimeRequest{
				Year:  1989,
				Month: 6,
				Day:   2,
			},
			date2: &entity.TimeRequest{
				Year:  1989,
				Month: 6,
				Day:   4,
			},
		},
		{
			desc:  "b",
			input: "2/6/2009 to 4/6/2009",
			date1: &entity.TimeRequest{
				Year:  2009,
				Month: 6,
				Day:   2,
			},
			date2: &entity.TimeRequest{
				Year:  2009,
				Month: 6,
				Day:   4,
			},
		},
	}

	for _, tc := range testcases{
		t.Run(tc.desc, func(t *testing.T) {
			date1 , date2 , err := Parser.ParseInput(tc.input)
			if err != nil {
				return
			}

			assert.Equal(t, date1,tc.date1)
			assert.Equal(t, date2,tc.date2)
		})
	}
}