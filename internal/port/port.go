package port

import (
	"context"
	"errors"

	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
)

// Rpush is port controller
type Rpush int

// HotPush is for hot port
func (t *Rpush) HotPush(ctx context.Context, args *esdocs.ESdocs, reply *string) error {
	err := global.HotRing.Push(*args)
	if err != nil {
		return errors.New("hotport is full")
	}
	return nil
}

// ColdPush is for cold port
func (t *Rpush) ColdPush(ctx context.Context, args *[]byte, reply *string) error {
	err := global.ColdRing.Push(string(*args))
	if err != nil {
		return errors.New("coldport is full")
	}
	return nil
}
