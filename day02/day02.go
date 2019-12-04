package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const filename = "input"

func parseArray(arr []int) []int {
	for i := 0; i < len(arr); i+=4 {
		switch arr[i] {
			case 1:
				arr[arr[i+3]] = arr[arr[i+1]] + arr[arr[i+2]]
			case 2:
				arr[arr[i+3]] = arr[arr[i+1]] * arr[arr[i+2]]
			case 99:
				return arr
			default:
				fmt.Println("Unexpected opcode, exiting")
				os.Exit(1)
		}
	}
	return []int{0}
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

func searchGoal(arr []int, goal int) {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			var mem = make([]int,len(arr))
			copy(mem,arr)
			mem[1] = i
			mem[2] = j
			if parseArray(mem)[0] == goal {
				fmt.Printf("Goal is %d %d\n", i,j)
				return
			}
			mem = arr
		}
	}
}

func main() {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	arr := convertArray(strings.Split(string(f), ","))
	searchGoal(arr, 19690720)
}
