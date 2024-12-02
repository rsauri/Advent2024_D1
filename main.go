package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getData(filename string) ([]int, error) {
	// Open the List 1 file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the list file: ", err)
		return nil, err
	}
	defer file.Close()

	// Create a slice to store the values
	var values []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert the line to an integer
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting line to integer:", err)
			return nil, err
		}
		// Add the integer to the slice
		values = append(values, num)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Sort the slice
	sort.Ints(values)

	// Print the sorted values
	return values, nil

}

func subtractAndAbs(a, b int) int {
	result := a - b
	if result < 0 {
		return -result
	}
	return result
}

func main() {
	var list1, list2 []int
	var err error
	var distance int = 0

	//Get First list
	list1, err = getData("Inputs/List1.txt")
	if err != nil {
		fmt.Println("Error processing List 1: ", err)
		return
	}

	//Get Second List
	list2, err = getData("Inputs/List2.txt")
	if err != nil {
		fmt.Println("Error processing List 2: ", err)
		return
	}

	//Loop thru each list and calculate the distance
	for i := 0; i < 1000; i++ {
		distance += subtractAndAbs(list1[i], list2[i])
	}

	fmt.Println("Distance: ", distance)
}
