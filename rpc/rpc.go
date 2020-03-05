package rpcdemo

import (
	"errors"
)

// Service.Method

// DemoService ...
type DemoService struct{}

// Args ...
type Args struct {
	A, B int
}

// Div ...
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
