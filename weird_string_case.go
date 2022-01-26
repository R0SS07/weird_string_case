package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toWeirdCase("String"))            // "StRiNg"
	fmt.Println(toWeirdCase("Weird string case")) // "WeIrD StRiNg CaSe"

}
func toWeirdCase(str string) string {
	arr := strings.Split(str, "")
	res := []string{}
	index := 0
	for k, value := range arr {
		if index%2.0 == 0 {
			res = append(res, strings.ToUpper(value))
		} else {
			res = append(res, strings.ToLower(value))
		}
		if arr[k] == " " {
			index = 0
		} else {
			index++
		}
	}
	return strings.Join(res, "")
}
