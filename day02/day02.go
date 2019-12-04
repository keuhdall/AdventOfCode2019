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
	for i := 0; i < len(arr); i+=4 {
		switch arr[i] {
			case 1:
				arr[arr[i+3]] = arr[arr[i+1]] + arr[arr[i+2]]
			case 2:
				arr[arr[i+3]] = arr[arr[i+1]] * arr[arr[i+2]]
			case 99:
				fmt.Println("Final state of the array : ")
				fmt.Println(arr)
				return
			default:
				fmt.Println("Unexpected opcode, exiting")
				os.Exit(1)
		}
	}
}

func convertArray(arr []string) []int {
	var intValues []int
	for _, i := range arr {
		val, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		intValues = append(intValues, val)
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
