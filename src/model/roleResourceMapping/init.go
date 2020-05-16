package roleResourceMapping

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
var roleResourcesMap map[int64][]RoleResourceMapping

// Init will load roleResourceMapping defined in file roleResourceMapping.csv
func Init() {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.ROLE_RESOURCE_MAPPING_FILE_NAME)

	LoadRolesFile(fileName)
}

func LoadRolesFile(fileName string) {
	roleResourcesMap = make(map[int64][]RoleResourceMapping)

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
			fmt.Println("Error in reading csv records: ", err)
			return
		}

		roleId, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid Role ID while parsing roleResourceMapping", err)
		}

		resourceId, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid Resource ID while parsing roleResourceMapping", err)
		}

		actionId, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			fmt.Println("Invalid Action ID while parsing roleResourceMapping", err)
		}

		status, err := strconv.ParseBool(record[3])
		if err != nil {
			fmt.Println("Invalid Status while parsing roleResourceMapping", err)
		}

		roleResourcesMap[resourceId] = append(roleResourcesMap[resourceId], RoleResourceMapping{
			RoleId:     roleId,
			ResourceId: resourceId,
			ActionId:   actionId,
			Status:     status,
		})

	}

	defer csvFile.Close()

	return
}
