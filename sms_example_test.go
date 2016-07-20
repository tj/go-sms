package sms_test

import (
	env "github.com/segmentio/go-env"

	"github.com/tj/go-sms"
)

// Example showing the most basic usage and defaults.
func Example() {
	sms.Send("Hello World", env.MustGet("PHONE"))
}

// Example showing how to override the defaults unless you want to
// go all-out and use the SMS struct directly.
func Example_overrides() {
	sms.DefaultMaxPrice = 0.5
	sms.DefaultType = sms.Transactional
	sms.Send("Hello World", env.MustGet("PHONE"))
}
