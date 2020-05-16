package userRoleMapping

type UserRoleMappingModule struct {
	UserRolesMap *map[int64][]UserRoleMapping
}

func NewModule() *UserRoleMappingModule {
	return &UserRoleMappingModule{
		UserRolesMap: &userRolesMap,
	}
}
