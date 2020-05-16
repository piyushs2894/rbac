package userRoleMapping

import (
	"errors"
)

// Many to Many mapping of UserId and RoleId
type UserRoleMapping struct {
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
	Status bool  `json:"status"`
}

// New mapping is created with status as active
func New(userId, roleId, actionId int64) (*UserRoleMapping, error) {
	return &UserRoleMapping{
		UserId: userId,
		RoleId: roleId,
		Status: true,
	}, nil
}

// For new role, check if parent and role name combination should be unique
func (module *UserRoleMappingModule) InsertOrUpdateRoleMapping(userRoleMapping UserRoleMapping) error {
	// Validating if roleId and resouceIdMapping exists

	var updatedUserRoleMapping []UserRoleMapping

	if mapping, ok := (*module.UserRolesMap)[userRoleMapping.UserId]; ok {
		for _, v := range mapping {
			if v.RoleId != userRoleMapping.RoleId {
				updatedUserRoleMapping = append(updatedUserRoleMapping, v)
			}
		}
	}

	updatedUserRoleMapping = append(updatedUserRoleMapping, userRoleMapping)
	(*module.UserRolesMap)[userRoleMapping.UserId] = updatedUserRoleMapping

	return nil
}

/** RemoveUserByRoleId will remove user associatation with particular role by disabling its status to false.
 ** It won't delete that entry from DB .
 */

func (module *UserRoleMappingModule) RemoveUserByRoleId(userId, roleId int64) error {
	var updatedUserRoleMapping []UserRoleMapping

	if mapping, ok := (*module.UserRolesMap)[userId]; ok {
		for _, v := range mapping {
			if v.RoleId == roleId {
				v.Status = false
			}

			updatedUserRoleMapping = append(updatedUserRoleMapping, v)
		}
	}

	(*module.UserRolesMap)[userId] = updatedUserRoleMapping

	return nil
}

func (module *UserRoleMappingModule) GetRolesByUserId(userId int64) ([]UserRoleMapping, error) {
	v, ok := (*module.UserRolesMap)[userId]
	if !ok {
		return v, errors.New("UserId is not associated with any roles")
	}

	return v, nil
}
