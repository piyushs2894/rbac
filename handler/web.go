package handler

import "github.com/RBAC/src/manager"

// Web module can be used to pass config, db connections and other fields
type Web struct {
	managerModule *manager.ManagerModule
}

func New(module *manager.ManagerModule) *Web {
	return &Web{
		managerModule: module,
	}
}
