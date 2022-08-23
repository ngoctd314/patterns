package health

// System informations
type System struct {
	Version       string `json:"version"`
	NumGoroutines int    `json:"num_goroutines"`
}
