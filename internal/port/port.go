package port

import (
	"context"
	"encoding/json"
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"log"
)

// args: 압축된 struct
func push(args string, hot bool) string {
	var raw map[string]interface{}

	err := json.Unmarshal([]byte(args), &raw)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	out, err := json.Marshal(raw)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	if hot {
		global.HotRing.Push(string(out))
	} else {
		global.ColdRing.Push(string(out))
	}

	return string(out)
}

// Rpush is port controller
type Rpush int

// HotPush is for hot port
func (t *Rpush) HotPush(ctx context.Context, args *string, reply *string) error {
	*reply = push(*args, true)
	return nil
}

// ColdPush is for cold port
func (t *Rpush) ColdPush(ctx context.Context, args *string, reply *string) error {
	*reply = push(*args, false)
	return nil
}
