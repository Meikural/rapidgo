# RapidGo Development Roadmap (Agile)

## Vision
Production-ready Go API framework with modular, chainable request handling pipeline.

---

## PHASE 1: Core Framework Foundation
**Goal:** Make chaining mechanism work + basic HTTP methods

### Sprint 1.1: Multi-HTTP Methods & Route Refactor
**Duration:** 1-2 hours
**Tasks:**
- [ ] Refactor global routes to Router struct
- [ ] Add POST, PUT, DELETE, PATCH methods
- [ ] Add Logic() alias to Use()
- [ ] Test all HTTP methods work
- [ ] Update example app

**Deliverable:** Users can do `rg.POST("/api").Logic(...)`

---

### Sprint 1.2: Path Parameters & Query Helpers
**Duration:** 1-2 hours
**Tasks:**
- [ ] Implement path param parsing (`:id`, `:slug`)
- [ ] Add Ctx.Param(key) - string
- [ ] Add Ctx.ParamInt(key) - int with error handling
- [ ] Add Ctx.Query(key) - string
- [ ] Add Ctx.QueryInt(key) - int with error handling
- [ ] Add Ctx.QueryDefault(key, default) - with fallback
- [ ] Test all param extraction scenarios
- [ ] Update example app with `/users/:id` route

**Deliverable:** Users can extract URL and query params easily

---

### Sprint 1.3: Request Body Parsing
**Duration:** 1 hour
**Tasks:**
- [ ] Add Ctx.BodyJSON(v any) error - parse JSON body
- [ ] Add Ctx.BodyForm() - parse form data
- [ ] Handle empty body gracefully
- [ ] Error handling for malformed JSON
- [ ] Test with POST/PUT requests

**Deliverable:** Users can parse request bodies

---

### Sprint 1.4: Response Formatting (JSON)
**Duration:** 1 hour
**Tasks:**
- [ ] Add Ctx.JSON(status int, data any) *Ctx
- [ ] Add Ctx.String(status int, text string) *Ctx
- [ ] Update execute() to properly encode responses
- [ ] Set correct Content-Type headers
- [ ] Test JSON encoding and string responses
- [ ] Update example app

**Deliverable:** Standardized response formatting

---

## PHASE 2: Chaining Mechanism (User-Defined Functions)
**Goal:** Enable `.auth().logic1().logic2()` syntax

### Sprint 2.1: Framework-Level Chaining Infrastructure
**Duration:** 2 hours
**Tasks:**
- [ ] Design how custom methods attach to Route
- [ ] Create `logic/` folder structure pattern in docs
- [ ] Add helper function/pattern for users to extend Route
- [ ] Document how to create chainable methods
- [ ] Create example: `AddNumber()`, `UpdateNumber()` in test app
- [ ] Test chaining with user-defined functions

**Deliverable:** Users can define and chain custom functions

---

### Sprint 2.2: Built-in Step: Error Handling
**Duration:** 1 hour
**Tasks:**
- [ ] Create Error() step that handles ctx.Err
- [ ] Standardize error response format (JSON)
- [ ] Handle different error types (validation, auth, system)
- [ ] Add error logging capability
- [ ] Test error flow

**Deliverable:** `.error()` step added

---

### Sprint 2.3: Built-in Step: Logging
**Duration:** 1.5 hours
**Tasks:**
- [ ] Create Logging() step
- [ ] Log request method, path, params
- [ ] Log response status, execution time
- [ ] Log errors if present
- [ ] Support different log levels (INFO, ERROR, DEBUG)
- [ ] Test logging output

**Deliverable:** `.logging()` step added

---

## PHASE 3: Security & Flow Control
**Goal:** Auth, rate limiting, caching

### Sprint 3.1: Built-in Step: Authentication & Authorization
**Duration:** 2 hours
**Tasks:**
- [ ] Create Auth(roles ...string) step
- [ ] Design auth token validation (JWT placeholder)
- [ ] Add role-based access control (RBAC)
- [ ] Extract user from context
- [ ] Return 401/403 on auth failure
- [ ] Add Ctx.GetUser() helper
- [ ] Test auth flow with different roles

**Deliverable:** `.auth("Admin", "Manager")` step added

---

### Sprint 3.2: Built-in Step: Rate Limiting
**Duration:** 1.5 hours
**Tasks:**
- [ ] Create RateLimit(count int, unit string) step
- [ ] Implement in-memory rate limiter
- [ ] Support units: "sec", "min", "hour"
- [ ] Track by IP address
- [ ] Return 429 on limit exceeded
- [ ] Add cleanup for expired limits
- [ ] Test rate limit behavior

**Deliverable:** `.ratelimit(100, "min")` step added

---

### Sprint 3.3: Built-in Step: Caching
**Duration:** 1.5 hours
**Tasks:**
- [ ] Create Cache(duration int, unit string) step
- [ ] Implement in-memory cache
- [ ] Support units: "sec", "min", "hour"
- [ ] Cache based on route + params
- [ ] Return cached result if exists
- [ ] Auto-expire old cache entries
- [ ] Add cache clear utility
- [ ] Test cache hit/miss scenarios

**Deliverable:** `.cache(5, "min")` step added

---

## PHASE 4: Global Configuration & Middleware
**Goal:** Env config, CORS, global middleware

### Sprint 4.1: Environment Configuration
**Duration:** 1 hour
**Tasks:**
- [ ] Add .env file loading
- [ ] Parse CORS_ALLOWED_ORIGINS from env
- [ ] Parse PORT, DEBUG_MODE from env
- [ ] Add config helpers
- [ ] Validate required env vars
- [ ] Document env file format
- [ ] Test env loading

**Deliverable:** Env-based config system

---

### Sprint 4.2: Built-in: CORS Handling
**Duration:** 1 hour
**Tasks:**
- [ ] Create CORS middleware step
- [ ] Read allowed origins from env (CORS_ALLOWED_ORIGINS)
- [ ] Add Access-Control-Allow-* headers
- [ ] Handle preflight OPTIONS requests
- [ ] Apply CORS to all routes
- [ ] Test CORS headers

**Deliverable:** `.cors()` built-in, configured via env

---

### Sprint 4.3: Global Middleware Support
**Duration:** 1.5 hours
**Tasks:**
- [ ] Add Router.Use(step PipelineStep) for global middleware
- [ ] Run global middleware before route-specific steps
- [ ] Support multiple global middleware
- [ ] Example: request ID injection
- [ ] Test global middleware execution order
- [ ] Document middleware pattern

**Deliverable:** Global middleware support added

---

## PHASE 5: Error Handling & Response Standardization
**Goal:** Consistent error and success responses

### Sprint 5.1: Standardized Response Format
**Duration:** 1 hour
**Tasks:**
- [ ] Design response envelope (success/error format)
- [ ] Create Ctx.Success(data any) *Ctx
- [ ] Create Ctx.Fail(status int, message string) *Ctx (enhance)
- [ ] Standardize error response structure
- [ ] Test response format consistency

**Deliverable:** Standardized JSON response format

---

### Sprint 5.2: Enhanced Error Handling
**Duration:** 1 hour
**Tasks:**
- [ ] Add error codes/types
- [ ] Add error message templates
- [ ] Improve error().step error response
- [ ] Handle panics gracefully
- [ ] Add error middleware
- [ ] Test error scenarios

**Deliverable:** Robust error handling

---

## PHASE 6: Testing & Documentation
**Goal:** Unit tests + comprehensive docs

### Sprint 6.1: Unit Tests
**Duration:** 2 hours
**Tasks:**
- [ ] Test Router creation and routes
- [ ] Test Ctx param extraction
- [ ] Test chaining execution order
- [ ] Test Auth step
- [ ] Test RateLimit step
- [ ] Test Cache step
- [ ] Test Error step
- [ ] Test Logging step
- [ ] Test response formats
- [ ] 80%+ code coverage

**Deliverable:** Comprehensive test suite

---

### Sprint 6.2: Documentation
**Duration:** 2 hours
**Tasks:**
- [ ] Write README.md with quick start
- [ ] Write docs/GETTING_STARTED.md
- [ ] Write docs/API_REFERENCE.md
- [ ] Write docs/EXAMPLES.md with real scenarios
- [ ] Document folder structure
- [ ] Document how to create custom logic functions
- [ ] Document env configuration
- [ ] Add code comments

**Deliverable:** Complete documentation

---

## PHASE 7: Example Applications
**Goal:** Real-world examples

### Sprint 7.1: Todo API Example
**Duration:** 1.5 hours
**Tasks:**
- [ ] Create examples/todo-api/ folder structure
- [ ] Write SQL queries (no ORM)
- [ ] Create logic functions (create, read, update, delete)
- [ ] Add authentication
- [ ] Add rate limiting
- [ ] Add caching
- [ ] Add logging
- [ ] Include .env example

**Deliverable:** Working todo API example

---

### Sprint 7.2: User Management API Example
**Duration:** 1.5 hours
**Tasks:**
- [ ] Create examples/user-api/ folder structure
- [ ] User registration logic
- [ ] User login logic
- [ ] User profile endpoints
- [ ] Role-based access control demo
- [ ] Auth middleware integration
- [ ] Error handling showcase

**Deliverable:** Working user management API example

---

## PHASE 8: CLI Tool (Future Planning)
**Goal:** Code generation via CLI

### Sprint 8.1: CLI Foundation
**Duration:** 2 hours
**Tasks:**
- [ ] Set up cmd/rg package
- [ ] Create rg new appname command
- [ ] Generate project scaffold
- [ ] Generate go.mod automatically
- [ ] Create example main.go
- [ ] Test scaffolding

**Deliverable:** `rg new` command working

---

### Sprint 8.2: Code Generators
**Duration:** 2 hours
**Tasks:**
- [ ] Create rg make:logic LogicName command
- [ ] Generate logic function boilerplate
- [ ] Create rg make:route Name command
- [ ] Generate route boilerplate
- [ ] Test all generators

**Deliverable:** `rg make:logic`, `rg make:route` working

---

## PHASE 9: Future Enhancements (Backlog)
- [ ] Database migration tool (Prisma-like)
- [ ] Transaction support
- [ ] Validation framework
- [ ] WebSocket support
- [ ] GraphQL support
- [ ] Metrics/monitoring
- [ ] Circuit breaker
- [ ] Request ID tracking
- [ ] Health check endpoints
- [ ] Redis caching support

---

## Development Flow Summary

**Phase 1 → Phase 2 → Phase 3 → Phase 4 → Phase 5 → Phase 6 → Phase 7 → Phase 8**

Each phase builds on the previous one. Complete phases sequentially for stable features.

**Estimated Total Time (Phases 1-7):** 20-25 hours

---

## Key Milestones

✅ **Milestone 1 (Phase 1):** Core API handling works
✅ **Milestone 2 (Phase 2):** Chaining mechanism works
✅ **Milestone 3 (Phase 3-4):** Security & config works
✅ **Milestone 4 (Phase 5-6):** Errors & docs complete
✅ **Milestone 5 (Phase 7):** Examples prove it works
⏳ **Milestone 6 (Phase 8):** CLI makes it easy

---

## Ready to Start?

Review this roadmap. Let me know if:
- Order should change
- Any sprints should be split/combined
- Anything should be added/removed
- Timeline estimates seem off

Once approved, we start Phase 1, Sprint 1.1!