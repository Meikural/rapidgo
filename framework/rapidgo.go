package rapidgo

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes []Route
	mux    *http.ServeMux
}

func New() *Router {
	return &Router{
		routes: []Route{},
		mux:    http.NewServeMux(),
	}
}

func (rg *Router) GET(path string) Chainable {
	rb := &RouteBuilder{
		route: &route{
			method: "GET",
			path:   path,
			steps:  []PipelineStep{},
		},
	}

	// Register route handler
	rg.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		builtRoute := rb.Build()
		rg.execute(builtRoute, w, req)
	})

	return rb
}

func (rg *Router) POST(path string) Chainable {
	rb := &RouteBuilder{
		route: &route{
			method: "POST",
			path:   path,
			steps:  []PipelineStep{},
		},
	}

	// Register route handler
	rg.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		builtRoute := rb.Build()
		rg.execute(builtRoute, w, req)
	})

	return rb
}

func (rg *Router) PUT(path string) Chainable {
	rb := &RouteBuilder{
		route: &route{
			method: "PUT",
			path:   path,
			steps:  []PipelineStep{},
		},
	}

	// Register route handler
	rg.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "PUT" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		builtRoute := rb.Build()
		rg.execute(builtRoute, w, req)
	})

	return rb
}

func (rg *Router) DELETE(path string) Chainable {
	rb := &RouteBuilder{
		route: &route{
			method: "DELETE",
			path:   path,
			steps:  []PipelineStep{},
		},
	}

	// Register route handler
	rg.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "DELETE" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		builtRoute := rb.Build()
		rg.execute(builtRoute, w, req)
	})

	return rb
}

func (rg *Router) PATCH(path string) Chainable {
	rb := &RouteBuilder{
		route: &route{
			method: "PATCH",
			path:   path,
			steps:  []PipelineStep{},
		},
	}

	// Register route handler
	rg.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "PATCH" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		builtRoute := rb.Build()
		rg.execute(builtRoute, w, req)
	})

	return rb
}

func (rg *Router) execute(r *Route, w http.ResponseWriter, req *http.Request) {
	ctx := newCtx(w, req)

	for _, step := range r.steps {
		ctx = step(ctx)
		if ctx.stop {
			break
		}
	}

	if ctx.Err != nil {
		http.Error(w, ctx.Err.Error(), ctx.Status)
		return
	}

	fmt.Fprintf(w, "%v", ctx.Result)
}

func (rg *Router) Start(addr string) error {
	return http.ListenAndServe(addr, rg.mux)
}