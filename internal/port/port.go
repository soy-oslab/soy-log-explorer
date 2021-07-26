package port

import (
	"context"

	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
)

// Rpush is port controller
type Rpush int

// HotPush is for hot port
func (t *Rpush) HotPush(ctx context.Context, args *esdocs.ESdocs, reply *string) error {
	global.HotRing.Push(*args)
	return nil
}

// ColdPush is for cold port
func (t *Rpush) ColdPush(ctx context.Context, args *[]byte, reply *string) error {
	global.ColdRing.Push(string(*args))
	return nil
}
