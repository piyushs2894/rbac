package resource

type ResourceModule struct {
	ResourceMap *map[string]Resource
}

func NewModule() *ResourceModule {
	return &ResourceModule{
		ResourceMap: &resourceMap,
	}
}
