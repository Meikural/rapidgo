package rapidgo

import (
	"fmt"
	"net/http"
)

type Ctx struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Store  map[string]any
	Result any
	Status int
	Err    error
	stop   bool
}

func newCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		Req:    r,
		Res:    w,
		Store:  map[string]any{},
		Status: 200,
	}
}

func (c *Ctx) Set(key string, val any) {
	c.Store[key] = val
}

func (c *Ctx) Get(key string) any {
	return c.Store[key]
}

func (c *Ctx) Fail(status int, msg string) *Ctx {
	c.Status = status
	c.Err = fmt.Errorf(msg)
	c.stop = true
	return c
}

func (c *Ctx) Stop() {
	c.stop = true
}
