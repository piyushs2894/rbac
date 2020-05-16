package resource

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
var resourceMap map[string]Resource

// Init will load Resource defined in file Resources.csv
func Init() {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.RESOURCES_FILE_NAME)

	LoadResourcesFile(fileName)
}

func LoadResourcesFile(fileName string) {
	//Initialize global UserMap
	resourceMap = make(map[string]Resource)

	csvFile, err := helper.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadResourcesFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	//Checking if userName exists or not
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records while parsing resources: ", err)
			return
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid ID while parsing resources ", err)
		}

		status, err := strconv.ParseBool(record[2])
		if err != nil {
			fmt.Println("Invalid Status while parsing resources ", err)
		}

		var row Resource
		row = Resource{
			Id:     id,
			Name:   record[1],
			Status: status,
		}

		resourceMap[row.Name] = row
	}

	defer csvFile.Close()

	return
}
