package rapidgo

// Chainable is the interface that users extend with custom methods
type Chainable interface {
	Use(step PipelineStep) Chainable
	Build() *Route
}

// Internal route struct (private)
type route struct {
	method string
	path   string
	steps  []PipelineStep
}

// RouteBuilder is what users interact with and extend
type RouteBuilder struct {
	route *route
}

// Use adds a pipeline step and returns the RouteBuilder for chaining
func (rb *RouteBuilder) Use(step PipelineStep) Chainable {
	rb.route.steps = append(rb.route.steps, step)
	return rb
}

// Build converts RouteBuilder to Route (internal use)
func (rb *RouteBuilder) Build() *Route {
	return &Route{
		method: rb.route.method,
		path:   rb.route.path,
		steps:  rb.route.steps,
	}
}

// Route is the final route structure (exported for internal framework use)
type Route struct {
	method string
	path   string
	steps  []PipelineStep
}