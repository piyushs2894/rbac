package roleResourceMapping

import (
	"errors"
)

// Many to Many mapping of RoleId and ResourceId
type RoleResourceMapping struct {
	RoleId     int64
	ResourceId int64
	ActionId   int64
	Status     bool
}

// New mapping is created with status as active
func New(roleId, resourceId, actionId int64) (*RoleResourceMapping, error) {
	return &RoleResourceMapping{
		RoleId:     roleId,
		ResourceId: resourceId,
		ActionId:   actionId,
		Status:     true,
	}, nil
}

// For new role-resource-action_type combination, validate if it exists in DB. It should be unique
func (module *RoleResourceMappingModule) InsertNewRoleResourceMapping(roleResourceMapping RoleResourceMapping) error {
	// Validating if roleId and resouceIdMapping exists
	if mapping, ok := (*module.RoleResourcesMap)[roleResourceMapping.ResourceId]; ok {
		for _, v := range mapping {
			if v.RoleId == roleResourceMapping.RoleId && v.ActionId == roleResourceMapping.ActionId {
				return errors.New("Role Id - Resource-Action mapping already exists")
			}
		}
	}

	// inserting into map similar to insert in DB
	(*module.RoleResourcesMap)[roleResourceMapping.ResourceId] = append((*module.RoleResourcesMap)[roleResourceMapping.ResourceId], RoleResourceMapping{
		RoleId:     roleResourceMapping.RoleId,
		ResourceId: roleResourceMapping.ResourceId,
		ActionId:   roleResourceMapping.ActionId,
		Status:     roleResourceMapping.Status,
	})

	return nil
}

func (module *RoleResourceMappingModule) GetRoleIdByResourceIdAndActionId(resourceId, actionId int64) (int64, error) {
	if mapping, ok := (*module.RoleResourcesMap)[resourceId]; ok {
		for _, v := range mapping {
			if v.ActionId == actionId {
				return v.RoleId, nil
			}
		}
	}
	return 0, errors.New("No role associated with give resource and actionType")
}
