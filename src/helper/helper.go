package helper

import (
	"encoding/csv"
	"fmt"
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file %s. Error: %+v", fileName, err)
	}

	return file, nil
}

func WriteFile(writer *csv.Writer, record []string) error {
	if err := writer.Write(record); err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func GetParentDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed in getting current working directory ", err)
		return wd
	}

	return wd
}
