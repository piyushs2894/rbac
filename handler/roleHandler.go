package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RBAC/src/model/userRoleMapping"
)

func (web *Web) AddOrUpdateRoletoUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	r.Body.Close()
	var userRoleMappingRequest userRoleMapping.UserRoleMapping

	if err = json.Unmarshal(body, &userRoleMappingRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userRoleMappingRequest.Status = true

	manager := web.managerModule.GetManager()

	if err = manager.AddOrUpdateUserRoleMapping(userRoleMappingRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Add role to User is Success")))
}

func (web *Web) RemoveRoleFromUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	r.Body.Close()
	var userRoleMappingRequest userRoleMapping.UserRoleMapping

	if err = json.Unmarshal(body, &userRoleMappingRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	manager := web.managerModule.GetManager()

	if err = manager.RemoveUserRoleMapping(userRoleMappingRequest.UserId, userRoleMappingRequest.RoleId); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Remove role from User is Success")))
}

func (web *Web) CheckUserAccessHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var userId int64
	var actionType, resourceName string

	if r.FormValue("user_id") != "" {
		if userId, err = strconv.ParseInt(r.FormValue("user_id"), 10, 64); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if r.FormValue("action_type") != "" {
		actionType = r.FormValue("action_type")
	}

	if r.FormValue("resource_name") != "" {
		resourceName = r.FormValue("resource_name")
	}

	manager := web.managerModule.GetManager()

	isAuthorized, err := manager.CheckUserAccess(userId, actionType, resourceName)
	if !isAuthorized || err != nil {
		http.Error(w, errors.New("Unauthorized Access for this resource").Error(), 401)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("User is Authorized for accessing this resource")))
}
