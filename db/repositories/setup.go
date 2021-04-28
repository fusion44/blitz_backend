package repositories

import (
	"sync"

	"github.com/fusion44/raspiblitz-backend/graph/model"
)

// SetupRepository contains all functions regarding the setup
type SetupRepository struct {
	mu                  sync.Mutex
	SetupEventObservers map[string]struct {
		ID      string
		Channel chan *model.DeviceInfo
	}
}

func NewSetupRepository() *SetupRepository {
	return &SetupRepository{SetupEventObservers: make(map[string]struct {
		ID      string
		Channel chan *model.DeviceInfo
	})}
}

func (r *SetupRepository) AddDeviceInfoObserver(id string) chan *model.DeviceInfo {
	// Make the channel
	channel := make(chan *model.DeviceInfo, 1)

	r.mu.Lock()
	r.SetupEventObservers[id] = struct {
		ID      string
		Channel chan *model.DeviceInfo
	}{ID: id, Channel: channel}
	r.mu.Unlock()

	return channel
}

func (r *SetupRepository) DeleteDeviceInfoObserver(id string) {
	r.mu.Lock()
	delete(r.SetupEventObservers, id)
	r.mu.Unlock()
}
