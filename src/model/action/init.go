package action

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/RBAC/src/constant"
	"github.com/RBAC/src/helper"
)

// defined map which will be initialized on app start. Private scope, can't be accessed outside this package
var actionMap map[string]Action

// Init will load action defined in file actions.csv
func Init() {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.ACTIONS_FILE_NAME)

	LoadActionsFile(fileName)
}

func LoadActionsFile(fileName string) {
	//Initialize global UserMap
	actionMap = make(map[string]Action)

	csvFile, err := helper.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadActionsFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	//Checking if userName exists or not
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records while parsing actions: ", err)
			return
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid Id while parsing actions ", err)
		}

		var row Action
		row = Action{
			Id:          id,
			ActionType:  record[1],
			Description: record[2],
		}

		actionMap[row.ActionType] = row
	}

	defer csvFile.Close()

	return
}
