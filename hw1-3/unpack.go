package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrIncorrectInput = errors.New("incorrect input, string contains only digits")

func Unpack(str string)(string,  error){
	sl:= strings.Split(str,"")
	var result string
	onlyDigit := true
	var builder strings.Builder
	for i:=0; i < len(str);i++{
		b,err:= strconv.Atoi(sl[i])
		if err != nil {
			builder.WriteString(sl[i])
			onlyDigit = false
		} else {
		if i == 0{
			builder.WriteString(sl[i])
			continue
		}
		for x:=0; x<b-1; x++ {
			builder.WriteString(sl[i-1])
		}

		}
		result = builder.String()
	}
	if onlyDigit {
		return "Something went wrong:", ErrIncorrectInput

	}
	return result, nil
}
func main() {
	str := "4d5f7asd7sdd7qee"
	fmt.Println(Unpack(str))

}
