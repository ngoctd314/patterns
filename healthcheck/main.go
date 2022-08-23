package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ngoctd314/patterns/healthcheck/health"
)

func main() {
	h, _ := health.NewHealthCheck(
		health.WithChecks(
			&health.Resource{
				Name:      "aerospike",
				Timeout:   0,
				SkipOnErr: false,
				Check: func(context.Context) error {
					return errors.New("fail during aerospike")
				},
			}))
	h.Register(&health.Resource{
		Name:      "mysql",
		Timeout:   0,
		SkipOnErr: false,
		Check: func(ctx context.Context) error {
			return errors.New("failed during connect mysql")
		},
	})

	r := h.Measurement(context.TODO())
	data, _ := json.Marshal(r)
	fmt.Println(string(data))
}
