package health

// Status type represents health status
type Status string

// enums health status
const (
	StatusOK          Status = "OK"
	StatusUnavailable Status = "Unavailable"
	StatusTimeout     Status = "Timeout during healthcheck"
)
