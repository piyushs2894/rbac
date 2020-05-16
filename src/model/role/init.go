package role

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/RBAC/src/constant"
	"github.com/RBAC/src/helper"
)

// Defined map which will be initialized on app start. Private scope, can't be accessed outside this package
var roleMap map[int64]Role

// Init will load Role defined in file Roles.csv
func Init() {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.ROLES_FILE_NAME)

	LoadRolesFile(fileName)
}

func LoadRolesFile(fileName string) {
	//Initialize global UserMap
	roleMap = make(map[int64]Role)

	csvFile, err := helper.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadRolesFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	//Checking if userName exists or not
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records while parsing roles: ", err)
			return
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid ID while parsing roles", err)
		}

		var parentIds map[int64]bool
		parentIdsString := strings.Split(strings.Replace(record[1], " ", "", -1), ",")

		for _, v := range parentIdsString {
			parentId, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				fmt.Println("Invalid Parent Idwhile parsing roles", err)
			} else {
				parentIds[parentId] = true
			}
		}

		status, err := strconv.ParseBool(record[4])
		if err != nil {
			fmt.Println("Invalid Status while parsing roles", err)
		}

		var row Role
		row = Role{
			Id:          id,
			ParentIds:   parentIds,
			Name:        record[2],
			Description: record[3],
			Status:      status,
		}

		roleMap[row.Id] = row
	}

	defer csvFile.Close()

	return
}
