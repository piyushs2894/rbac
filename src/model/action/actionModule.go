package action

type ActionModule struct {
	ActionMap *map[string]Action
}

func NewModule() *ActionModule {
	return &ActionModule{
		ActionMap: &actionMap,
	}
}
