package user

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/RBAC/src/constant"
	"github.com/RBAC/src/helper"
)

// Global UserMap to be initialized on app start
var UserMap map[string]User

func Init() {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.USER_FILE_NAME)

	LoadUserFile(fileName)
}

func LoadUserFile(fileName string) {
	//Initialize global UserMap
	UserMap = make(map[string]User)

	csvFile, err := helper.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadUserFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	//Checking if userName exists or not
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records: ", err)
			return
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid User ID ", err)
		}

		var row User
		row = User{
			ID:       id,
			UserName: record[1],
			Password: record[2],
			Name:     record[3],
			Email:    record[4],
			Contact:  record[5],
		}

		UserMap[row.UserName] = row
	}

	defer csvFile.Close()

	return
}
