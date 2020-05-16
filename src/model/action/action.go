package action

import "errors"

//Assuming all actions are independent and unique.
type Action struct {
	Id          int64
	ActionType  string
	Description string
}

// For simplicity lastId is taken as len(actionMap), if DB is used then this will be auto incremented.
func New(actionType, description string) *Action {

	lastId := (int64)(len(actionMap))

	return &Action{
		Id:          lastId + 1,
		ActionType:  actionType,
		Description: description,
	}
}

func (module *ActionModule) InsertNewAction(action Action) error {
	// check if this action is existing already
	if _, ok := (*module.ActionMap)[action.ActionType]; ok {
		return errors.New("Duplicate Action. Action already exists in database")
	}

	// inserting into map, refer to insert in DB
	(*module.ActionMap)[action.ActionType] = action

	return nil
}

//empty action object is passed as receiver here
func (module *ActionModule) GetActionByType(actionType string) (Action, error) {
	if v, ok := (*module.ActionMap)[actionType]; ok {
		return v, nil
	}

	return Action{}, errors.New("Invalid Action Type")
}
