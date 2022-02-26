package parser

import "dayCounter/entity"

var Parser parserInterface = &parser{}

type parserInterface interface {
	ParseInput(string)(*entity.TimeRequest, *entity.TimeRequest, error)
}

type parser struct {}




