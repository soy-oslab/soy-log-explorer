package port

import (
	"context"
	"errors"
	"fmt"

	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/internal/rest"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
	"github.com/soyoslab/soy_log_explorer/pkg/signal"
)

// Rpush is port controller
type Rpush int

// HotPush is for hot port
func (t *Rpush) HotPush(ctx context.Context, args *esdocs.ESdocs, reply *string) error {
	fmt.Println("Hot", args.Index, args.Docs)
	go rest.ESPush(*args)
	return nil
}

// ColdPush is for cold port
func (t *Rpush) ColdPush(ctx context.Context, args *[]byte, reply *string) error {
	err := global.ColdRing.Push(string(*args))
	signal.Signal()
	if err != nil {
		return errors.New("coldport is full")
	}
	return nil
}
