package resource

import (
	"errors"
	"time"
)

type Resource struct {
	Id        int64
	Name      string
	Status    bool // To signify if this resource is active or inactive
	CreatedAt time.Time
	UpdatedAt time.Time
}

// New resource is created with status as active. And Resource name should be unique
func New(name string) *Resource {

	lastId := (int64)(len(resourceMap))

	return &Resource{
		Id:        lastId + 1,
		Name:      name,
		Status:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (module *ResourceModule) InsertNewResource(resource Resource) error {
	//validate if resource exists or not
	if _, err := module.GetResourceByName(resource.Name); err == nil {
		return errors.New("Resource already exists")
	}

	// inserting into map, refer to insert in DB
	(*module.ResourceMap)[resource.Name] = resource

	return nil
}

//empty Resource object is passed as receiver here
func (module *ResourceModule) GetResourceByName(name string) (*Resource, error) {
	v, ok := (*module.ResourceMap)[name]
	if ok {
		return &v, nil
	}

	return nil, errors.New("Resource does not exist ")
}
