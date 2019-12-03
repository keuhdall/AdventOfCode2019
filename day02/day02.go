package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const filename = "input"

func parseArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 1:
			arr[i+3] = arr[i+1] + arr[i+2]
			break
		case 2:
			arr[i+3] = arr[i+1] * arr[i+2]
			break
		case 99:
			fmt.Printf("Final value at index 0 : %d\n", arr[0])
			return
		default:
			continue
		}
	}
}

func convertArray(arr []string) []int {
	var intValues []int
	for i := 0; i < len(arr); i++ {
		val, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		intValues = append(intValues, i, val)
	}
	return intValues
}

func main() {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	arr := convertArray(strings.Split(string(f), ","))
	parseArray(arr)
}
