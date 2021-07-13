package port

import (
	"context"

	"github.com/soyoslab/soy_log_explorer/internal/global"
)

func push(data []byte, hot bool) {
	if hot {
		global.HotRing.Push(string(data))
	} else {
		global.ColdRing.Push(string(data))
	}
}

// Rpush is port controller
type Rpush int

// HotPush is for hot port
func (t *Rpush) HotPush(ctx context.Context, args *[]byte, reply *string) error {
	push(*args, true)
	return nil
}

// ColdPush is for cold port
func (t *Rpush) ColdPush(ctx context.Context, args *[]byte, reply *string) error {
	push(*args, false)
	return nil
}
