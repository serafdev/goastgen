package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteString(path string, data string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	bytesWritten, err := writer.WriteString(data)
	if err != nil {
		fmt.Println("Got error while writting to a file.", err)
	}
    fmt.Printf("Bytes Written: %d\n", bytesWritten)
	writer.Flush()
}
