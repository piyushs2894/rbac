package role

import (
	"errors"
	"reflect"
	"time"
)

/** ParentIds is taken as array to store all multilevel parents for any role.
 */

type Role struct {
	Id          int64
	ParentIds   map[int64]bool
	Name        string
	Description string
	Status      bool // To signify if this role is active or inactive
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// New role is created with status as active
func New(parentId int64, name, description string) (*Role, error) {

	lastId := (int64)(len(roleMap))

	//fetching parentIds of a input parentId
	parentIds, err := getParentIdsForId(parentId)
	if err != nil {
		return nil, err
	}

	// Preparing new list of parentIds by appending current parentId and assigning to new role
	parentIds[parentId] = true

	return &Role{
		Id:          lastId + 1,
		ParentIds:   parentIds,
		Name:        name,
		Description: description,
		Status:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func getParentIdsForId(id int64) (map[int64]bool, error) {
	if v, ok := roleMap[id]; ok {
		return v.ParentIds, nil
	}

	return nil, errors.New("Invalid role Id")
}

func getRoleIdsForParentId(id int64) []int64 {
	var roleIds []int64

	for roleId, _ := range roleMap {
		if _, ok := roleMap[roleId].ParentIds[id]; ok {
			roleIds = append(roleIds, roleId)
		}
	}

	return roleIds
}

// Validating for new role, check if parent and role name combination should be unique
func (role *Role) InsertNewRole() error {
	for _, v := range roleMap {
		if reflect.DeepEqual(role.ParentIds, v.ParentIds) && role.Name == v.Name {
			return errors.New("Dupplicate role. Role already exists")
		}
	}

	// inserting into map similar to insert in DB
	roleMap[role.Id] = *role

	return nil
}

func (role *Role) GetRoleById(id int64) error {
	for _, v := range roleMap {
		if v.Id == id {
			role = &v
		}
	}

	return errors.New("Role does not exist for given id")
}
