package port

import (
	"context"
	"github.com/soyoslab/soy_log_explorer/internal/global"
)

func push(args global.ESdocs, hot bool) string {
	if hot {
		global.HotRing.Push(args)
	} else {
		global.ColdRing.Push(args)
	}

	return ""
}

// Rpush is port controller
type Rpush int

// HotPush is for hot port
func (t *Rpush) HotPush(ctx context.Context, args *global.ESdocs, reply *string) error {
	*reply = push(*args, true)
	return nil
}

// ColdPush is for cold port
func (t *Rpush) ColdPush(ctx context.Context, args *global.ESdocs, reply *string) error {
	*reply = push(*args, false)
	return nil
}
