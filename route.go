package rapidgo

type Route struct {
	method string
	path   string
	steps  []PipelineStep
}

func (r *Route) Use(step PipelineStep) *Route {
	r.steps = append(r.steps, step)
	return r
}
