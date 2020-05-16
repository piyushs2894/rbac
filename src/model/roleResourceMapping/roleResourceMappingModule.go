package roleResourceMapping

type RoleResourceMappingModule struct {
	RoleResourcesMap *map[int64][]RoleResourceMapping
}

func NewModule() *RoleResourceMappingModule {
	return &RoleResourceMappingModule{
		RoleResourcesMap: &roleResourcesMap,
	}
}
