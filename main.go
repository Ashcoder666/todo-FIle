package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// arr := []string{}
	fmt.Println("Welcome to gofer todo CLI")
	// if this file exists = get all data to array else add new file
	filepath := "todobackup.txt"

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}

	defer file.Close()

	for {
		fmt.Println("To see Todo list :Type TODO | To add new Todo :Type ADD | ctrl+c to exit")

		var todoOrAdd string
		fmt.Scan(&todoOrAdd)
		if todoOrAdd == "TODO" {

			data, readErr := os.ReadFile(filepath)

			if readErr != nil {
				fmt.Printf("Error reading from the file: %v\n", readErr)
				return
			}

			// Convert the byte slice to a string for printing
			fileContent := string(data)
			arrcontent := strings.Split(fileContent, ",")
			// fmt.Println(arrcontent)

			for index, todo := range arrcontent {
				if index != len(arrcontent)-1 {

					fmt.Printf("%v: %v \n", index, todo)
				}

			}
		} else if todoOrAdd == "ADD" {
			println("Enter your todo here")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newTodo := scanner.Text()
			// arr = append(arr, newTodo)
			newTodoWithNewline := newTodo + ","

			data := []byte(newTodoWithNewline)
			_, writeErr := file.Write(data)
			if writeErr != nil {
				fmt.Printf("Error writing to the file: %v\n", writeErr)
				return
			}
		} else {
			println("Invalid Entry")
		}
		// arrayString := strings.Join(arr, " ")

	}

}
