package repositories

import (
	"sync"

	"github.com/fusion44/raspiblitz-backend/graph/model"
)

// SetupRepository contains all functions regarding the setup
type SetupRepository struct {
	mu        sync.Mutex
	Observers map[string]struct {
		ID      string
		Channel chan *model.SetupInfoEvent
	}
}

func New() *SetupRepository {
	return &SetupRepository{Observers: make(map[string]struct {
		ID      string
		Channel chan *model.SetupInfoEvent
	})}
}

func (r *SetupRepository) AddObserver(id string) chan *model.SetupInfoEvent {
	// Make the channel
	channel := make(chan *model.SetupInfoEvent, 1)

	r.mu.Lock()
	r.Observers[id] = struct {
		ID      string
		Channel chan *model.SetupInfoEvent
	}{ID: id, Channel: channel}
	r.mu.Unlock()

	return channel
}

func (r *SetupRepository) DeleteObserver(id string) {
	r.mu.Lock()
	delete(r.Observers, id)
	r.mu.Unlock()
}
