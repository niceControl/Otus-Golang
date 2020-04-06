package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	//"log"
)


func Unpack(str string)(string,  error){
	sl:= strings.Split(str,"")
	var result string
	onlyDigit := true
	for i:=0; i < len(str);i++{
		b,err:= strconv.Atoi(sl[i])
		if err != nil {
			result += sl[i]
			onlyDigit = false
		} else {
		if i == 0{
			continue
		}
		result += strings.Repeat(sl[i-1], b-1)
		}
	}
	if onlyDigit {
		fmt.Println("Incorrect input")
		os.Exit(0)
	}
	return result, nil
}
func main() {
	str := "sssss5"
	fmt.Println(Unpack(str))

}
