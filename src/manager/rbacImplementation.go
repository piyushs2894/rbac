package manager

import (
	"github.com/RBAC/src/model/action"
	"github.com/RBAC/src/model/resource"
	"github.com/RBAC/src/model/roleResourceMapping"
	"github.com/RBAC/src/model/userRoleMapping"
)

type Manager interface {
	CheckUserAccess(userId int64, actionType, resource string) (bool, error)
	AddOrUpdateUserRoleMapping(mapping userRoleMapping.UserRoleMapping) error
	RemoveUserRoleMapping(userId, roleId int64) error
}

// Used to generate mock modules for unit testing purpose
type ManagerModule struct {
	userRoleMappingModule     *userRoleMapping.UserRoleMappingModule
	roleResourceMappingModule *roleResourceMapping.RoleResourceMappingModule
	actionModule              *action.ActionModule
	resourceModule            *resource.ResourceModule
}

func New() *ManagerModule {
	return &ManagerModule{
		userRoleMappingModule:     userRoleMapping.NewModule(),
		roleResourceMappingModule: roleResourceMapping.NewModule(),
		actionModule:              action.NewModule(),
		resourceModule:            resource.NewModule(),
	}
}

func (module *ManagerModule) GetManager() Manager {
	return module
}

func (module *ManagerModule) CheckUserAccess(userId int64, actionType, resourceName string) (bool, error) {
	//get roles by userId
	rolesMapping, err := module.userRoleMappingModule.GetRolesByUserId(userId)
	if err != nil {
		return false, err
	}

	//get action by actionType
	action, err := module.actionModule.GetActionByType(actionType)
	if err != nil {
		return false, err
	}

	//get resource by resourceName
	resource, err := module.resourceModule.GetResourceByName(resourceName)
	if err != nil || !resource.Status {
		return false, err
	}

	// get roleId by resourceIdAndActionId
	roleId, err := module.roleResourceMappingModule.GetRoleIdByResourceIdAndActionId(resource.Id, action.Id)
	if err != nil {
		return false, err
	}

	// Check if roleId exists in rolesMapping for givenUser. If exists with Status "true", then return true
	for _, v := range rolesMapping {
		if v.RoleId == roleId && v.Status {
			return true, nil
		}
	}

	return false, nil
}

func (module *ManagerModule) AddOrUpdateUserRoleMapping(mapping userRoleMapping.UserRoleMapping) error {
	if err := module.userRoleMappingModule.InsertOrUpdateRoleMapping(mapping); err != nil {
		return err
	}

	return nil
}

func (module *ManagerModule) RemoveUserRoleMapping(userId, roleId int64) error {
	if err := module.userRoleMappingModule.RemoveUserByRoleId(userId, roleId); err != nil {
		return err
	}

	return nil
}
