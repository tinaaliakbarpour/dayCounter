package main

import (
	"bufio"
	"dayCounter/entity"
	"dayCounter/parser"
	"fmt"
	"log"
	"os"
	"dayCounter/time"
)

func main() {
	fmt.Println("please enter 2 dates for starting calculation then hit Enter : for example <2/6/1989 to 5/6/1990>")
	//reading data from user input <it should be like 2/6/1983 to 22/6/1983> and after pressing ENTER the calculation will start
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Println("error in read from input: ", err)
		return
	}

	date1 , date2 ,err := parser.Parser.ParseInput(input)

	if err != nil{
		log.Println(err)
		return
	}

	// the dates will be filled in these two structs
	start := &entity.TimeRequest{}
	end := &entity.TimeRequest{}

	// detect which date is before another one
	if time.Calculator.IsBefore(*date1,*date2){
		start = date1
		end = date2
	}else {
		start = date2
		end = date1
	}

	fmt.Println(time.Calculator.CountDays(*start,*end) - 1)

}


