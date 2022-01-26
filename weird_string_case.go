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
	i, arr := 0, []string{}
	for k, val := range strings.Split(str, "") {
		if i%2 == 0 {
			arr = append(arr, strings.ToUpper(val))
		} else {
			arr = append(arr, strings.ToLower(val))
		}
		if str[k] == 32 {
			i = 0
		} else {
			i++
		}
	}
	return strings.Join(arr, "")
}
