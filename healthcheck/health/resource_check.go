package health

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// CheckFunc is the func which executes the check
type CheckFunc func(context.Context) error

// Resource type represents resource need to check
type Resource struct {
	// Name is the name of the resource to be checked
	Name string
	// Timeout is the timeout defined for every check
	Timeout   time.Duration
	SkipOnErr bool
	// Check is the func which executes the check
	Check CheckFunc
}

// Health is the health-checks container
type Health struct {
	checks map[string]Resource
}

// NewHealthCheck ...
func NewHealthCheck(opts ...Option) (*Health, error) {
	h := &Health{
		checks: make(map[string]Resource),
	}

	for _, o := range opts {
		if err := o(h); err != nil {
			return nil, err
		}
	}

	return h, nil
}

// Register registers a check config to be performed
func (h *Health) Register(r *Resource) error {
	if r.Name == "" {
		return errors.New("health check must have a name to be registered")
	}

	if _, ok := h.checks[r.Name]; ok {
		return fmt.Errorf("health check %q is already registered", r.Name)
	}

	h.checks[r.Name] = *r

	return nil
}

// Measurement runs all registered health checks and returns summary status
func (h *Health) Measurement(ctx context.Context) *Response {
	var status = StatusOK
	var (
		mu       sync.Mutex
		failures = make(map[string]string)
	)

	var wg sync.WaitGroup
	wg.Add(len(h.checks))
	for _, c := range h.checks {
		go func(r Resource) {
			defer wg.Done()
			err := r.Check(ctx)
			if err != nil {
				status = StatusUnavailable
				mu.Lock()
				failures[r.Name] = err.Error()
				mu.Unlock()
			}
		}(c)
	}
	wg.Wait()

	return &Response{
		Status:    status,
		Timestamp: time.Now(),
		Failures:  failures,
		System: System{
			Version:       runtime.Version(),
			NumGoroutines: runtime.NumGoroutine(),
		},
		Whoiam: *getWhoiam(),
	}
}
