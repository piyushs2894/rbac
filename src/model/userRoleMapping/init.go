package userRoleMapping

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
var userRolesMap map[int64][]UserRoleMapping

// Init will load roleResourceMapping defined in file roleResourceMapping.csv
func Init() {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.USER_ROLE_MAPPING_FILE_NAME)

	LoadUserRolesMappingFile(fileName)
}

func LoadUserRolesMappingFile(fileName string) {
	userRolesMap = make(map[int64][]UserRoleMapping)

	csvFile, err := helper.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadUserRolesMappingFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records: ", err)
			return
		}

		userId, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid User ID while parsing user role mapping", err)
		}

		roleId, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid Role ID while parsing user role mapping", err)
		}

		status, err := strconv.ParseBool(record[2])
		if err != nil {
			fmt.Println("Invalid Status while parsing user role mapping", err)
		}

		userRolesMap[userId] = append(userRolesMap[userId], UserRoleMapping{
			UserId: userId,
			RoleId: roleId,
			Status: status,
		})
	}

	defer csvFile.Close()

	return
}
