package persistenceMem

import "github.com/chenyf/mqttapi/vlplugin/vlpersistence"

type system struct {
	status *dbStatus
}

func (s *system) GetInfo() (*vlpersistence.SystemState, error) {
	state := &vlpersistence.SystemState{}

	return state, nil
}

func (s *system) SetInfo(state *vlpersistence.SystemState) error {
	return nil
}
