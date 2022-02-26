package parser

import (
	"dayCounter/entity"
	"errors"
	"strconv"
	"strings"
)

//ParseInput will parse user input and fill the time request
func (parser) ParseInput(input string) (*entity.TimeRequest, *entity.TimeRequest, error) {

	if !strings.Contains(input," to "){
		return nil,nil,errors.New("the Input format is not correct")
	}

	dates := strings.Split(input," to ")

	//remove the \n from second date
	dates[1] = strings.Replace(dates[1],"\n","",-1)

	var response []entity.TimeRequest

	for _, date := range dates{
		if !strings.Contains(date,"/"){
			return nil,nil,errors.New("the Input format is not correct ")
		}
		yearMonthDay := strings.Split(date,"/")
		year , err := strconv.Atoi(yearMonthDay[2])
		if err != nil {
			return nil, nil, err
		}
		month , err := strconv.Atoi(yearMonthDay[1])
		if err != nil {
			return nil, nil, err
		}
		day , err := strconv.Atoi(yearMonthDay[0])
		if err != nil {
			return nil, nil, err
		}

		response = append(response,entity.TimeRequest{
			Year:  year,
			Month: month,
			Day:   day,
		})
	}

	return &response[0], &response[1], nil
}

