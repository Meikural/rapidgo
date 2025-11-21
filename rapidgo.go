package rapidgo

import (
	"fmt"
	"net/http"
)

var routes []Route

func GET(path string) *Route {
	r := Route{
		method: "GET",
		path:   path,
		steps:  []PipelineStep{},
	}

	routes = append(routes, r)

	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}
		execute(&r, w, req)
	})

	return &r
}

func execute(r *Route, w http.ResponseWriter, req *http.Request) {
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

func Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
