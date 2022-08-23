package health

import (
	"sync"
	"time"
)

// Response type represent healthcheck response
type Response struct {
	Status    Status            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Failures  map[string]string `json:"failures,omitempty"`
	System    System            `json:"system"`
	Whoiam    whoiam            `json:"whoiam"`
}

type whoiam struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var (
	whoiamInstance *whoiam
	whoiamMu       sync.Mutex
)

func getWhoiam() *whoiam {
	if whoiamInstance == nil {
		whoiamMu.Lock()
		defer whoiamMu.Unlock()
		if whoiamInstance == nil {
			whoiamInstance = &whoiam{
				Name:    "Healcheck name",
				Version: "v1",
			}
		}
	}
	return whoiamInstance
}
