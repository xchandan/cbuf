package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("dmesg.log", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	cb := NewContextBuffer(20)

	for idx := 1; sc.Scan(); idx++ {
		line := fmt.Sprintf("%d: %s", idx, sc.Text())
		cb.Add(line)
		fmt.Println(line)
	}
	fmt.Println("==================")

	for _, line := range cb.Get() {
		fmt.Println(line)
	}
}
