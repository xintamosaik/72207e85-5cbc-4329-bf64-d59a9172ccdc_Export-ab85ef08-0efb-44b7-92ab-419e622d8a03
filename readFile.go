package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Specify the file path
	filePath := "/path/to/your/file.txt"

	// Read the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the file content
	fmt.Println(string(content))
}
