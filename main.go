package main

import (
	"net/http"

	"github.com/RBAC/handler"
	"github.com/RBAC/src/constant"
	"github.com/RBAC/src/helper"
	"github.com/RBAC/src/manager"
	"github.com/RBAC/src/model/action"
	"github.com/RBAC/src/model/resource"
	"github.com/RBAC/src/model/roleResourceMapping"
	"github.com/RBAC/src/model/user"
	"github.com/RBAC/src/model/userRoleMapping"
)

func main() {
	constant.PARENT_DIRECTORY = helper.GetParentDirectory()

	user.Init()
	action.Init()
	resource.Init()
	roleResourceMapping.Init()
	userRoleMapping.Init()

	InitRoutes()
}

func InitRoutes() {

	//Init relevant modules or config or db connections
	managerModule := manager.New()
	web := handler.New(managerModule)

	//users
	http.HandleFunc("/api/v1/user/signUp", web.SignupHandler)
	http.HandleFunc("/api/v1/user/login", web.LoginHandler)

	http.HandleFunc("/api/v1/addOrUpdateRoletoUser", web.AddOrUpdateRoletoUserHandler)
	http.HandleFunc("/api/v1/removeRolefromUser", web.RemoveRoleFromUserHandler)
	http.HandleFunc("/api/v1/checkUserAccess", web.CheckUserAccessHandler)

	http.ListenAndServe(":8080", nil)
}
