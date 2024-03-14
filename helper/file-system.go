package helper

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteFile(data interface{}, filename string) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error marshalling data", err.Error())
		return
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file", err.Error())
			return
		}
		defer file.Close()
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("Error opening file", err.Error())
		return
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		fmt.Println("Error writing to file", err.Error())
		return
	}
}

func ReadFile(filename string) []byte {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("error opening file", err.Error())
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println("error getting file info", err.Error())
		return nil
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	return buffer
}
