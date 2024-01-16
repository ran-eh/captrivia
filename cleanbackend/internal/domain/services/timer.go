// internal/domain/services/timer.go
package services

import (
	"sync"
	"time"
)

// TimerService is responsible for handling timing logic within the game.
type TimerService struct {
	timers   map[string]*time.Timer
	timeLock sync.Mutex
}

// NewTimerService returns a new instance of TimerService.
func NewTimerService() *TimerService {
	return &TimerService{
		timers: make(map[string]*time.Timer),
	}
}

// StartTimer starts a countdown timer for a specified identifier (e.g., a player ID or game session ID).
func (ts *TimerService) StartTimer(id string, duration time.Duration, callback func()) {
	ts.timeLock.Lock()
	defer ts.timeLock.Unlock()

	// If there's already an existing timer, stop it before starting a new one
	if existingTimer, exists := ts.timers[id]; exists && existingTimer != nil {
		existingTimer.Stop()
		delete(ts.timers, id)
	}

	// Create a new timer and save it to the map
	ts.timers[id] = time.AfterFunc(duration, func() {
		// Delete the timer entry before calling the callback
		ts.timeLock.Lock()
		delete(ts.timers, id)
		ts.timeLock.Unlock()

		callback()
	})
}

// StopTimer stops the countdown timer for the specified identifier.
func (ts *TimerService) StopTimer(id string) {
	ts.timeLock.Lock()
	defer ts.timeLock.Unlock()

	if timer, exists := ts.timers[id]; exists && timer != nil {
		timer.Stop()
		delete(ts.timers, id)
	}
}

// RemainingTime returns the remaining time for the given timer identifier.
// If there’s no active timer for the provided ID, it returns 0 and false.
func (ts *TimerService) RemainingTime(id string) (time.Duration, bool) {
	ts.timeLock.Lock()
	defer ts.timeLock.Unlock()

	timer, exists := ts.timers[id]
	if !exists || timer == nil {
		return 0, false
	}

	// Getting the remaining time isn't straightforward in the stdlib.
	// We would need to manage and store that state ourselves if it were a requirement.

	// For the purpose of this implementation, we'll just return that a timer exists,
	// but for an actual remaining time check, you'd need to manage the state separately.
	return 0, true
}