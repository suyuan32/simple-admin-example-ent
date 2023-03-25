// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/suyuan32/simple-admin-example-api/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/suyuan32/simple-admin-example-api/ent/course"
	"github.com/suyuan32/simple-admin-example-api/ent/exam"
	"github.com/suyuan32/simple-admin-example-api/ent/student"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Course is the client for interacting with the Course builders.
	Course *CourseClient
	// Exam is the client for interacting with the Exam builders.
	Exam *ExamClient
	// Student is the client for interacting with the Student builders.
	Student *StudentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Course = NewCourseClient(c.config)
	c.Exam = NewExamClient(c.config)
	c.Student = NewStudentClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Course:  NewCourseClient(cfg),
		Exam:    NewExamClient(cfg),
		Student: NewStudentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Course:  NewCourseClient(cfg),
		Exam:    NewExamClient(cfg),
		Student: NewStudentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Course.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Course.Use(hooks...)
	c.Exam.Use(hooks...)
	c.Student.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Course.Intercept(interceptors...)
	c.Exam.Intercept(interceptors...)
	c.Student.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CourseMutation:
		return c.Course.mutate(ctx, m)
	case *ExamMutation:
		return c.Exam.mutate(ctx, m)
	case *StudentMutation:
		return c.Student.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CourseClient is a client for the Course schema.
type CourseClient struct {
	config
}

// NewCourseClient returns a client for the Course from the given config.
func NewCourseClient(c config) *CourseClient {
	return &CourseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `course.Hooks(f(g(h())))`.
func (c *CourseClient) Use(hooks ...Hook) {
	c.hooks.Course = append(c.hooks.Course, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `course.Intercept(f(g(h())))`.
func (c *CourseClient) Intercept(interceptors ...Interceptor) {
	c.inters.Course = append(c.inters.Course, interceptors...)
}

// Create returns a builder for creating a Course entity.
func (c *CourseClient) Create() *CourseCreate {
	mutation := newCourseMutation(c.config, OpCreate)
	return &CourseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Course entities.
func (c *CourseClient) CreateBulk(builders ...*CourseCreate) *CourseCreateBulk {
	return &CourseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Course.
func (c *CourseClient) Update() *CourseUpdate {
	mutation := newCourseMutation(c.config, OpUpdate)
	return &CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CourseClient) UpdateOne(co *Course) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourse(co))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CourseClient) UpdateOneID(id uint64) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourseID(id))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Course.
func (c *CourseClient) Delete() *CourseDelete {
	mutation := newCourseMutation(c.config, OpDelete)
	return &CourseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CourseClient) DeleteOne(co *Course) *CourseDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CourseClient) DeleteOneID(id uint64) *CourseDeleteOne {
	builder := c.Delete().Where(course.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CourseDeleteOne{builder}
}

// Query returns a query builder for Course.
func (c *CourseClient) Query() *CourseQuery {
	return &CourseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCourse},
		inters: c.Interceptors(),
	}
}

// Get returns a Course entity by its id.
func (c *CourseClient) Get(ctx context.Context, id uint64) (*Course, error) {
	return c.Query().Where(course.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CourseClient) GetX(ctx context.Context, id uint64) *Course {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStudents queries the students edge of a Course.
func (c *CourseClient) QueryStudents(co *Course) *StudentQuery {
	query := (&StudentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, course.StudentsTable, course.StudentsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryExams queries the exams edge of a Course.
func (c *CourseClient) QueryExams(co *Course) *ExamQuery {
	query := (&ExamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(exam.Table, exam.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, course.ExamsTable, course.ExamsColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CourseClient) Hooks() []Hook {
	return c.hooks.Course
}

// Interceptors returns the client interceptors.
func (c *CourseClient) Interceptors() []Interceptor {
	return c.inters.Course
}

func (c *CourseClient) mutate(ctx context.Context, m *CourseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CourseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CourseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Course mutation op: %q", m.Op())
	}
}

// ExamClient is a client for the Exam schema.
type ExamClient struct {
	config
}

// NewExamClient returns a client for the Exam from the given config.
func NewExamClient(c config) *ExamClient {
	return &ExamClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `exam.Hooks(f(g(h())))`.
func (c *ExamClient) Use(hooks ...Hook) {
	c.hooks.Exam = append(c.hooks.Exam, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `exam.Intercept(f(g(h())))`.
func (c *ExamClient) Intercept(interceptors ...Interceptor) {
	c.inters.Exam = append(c.inters.Exam, interceptors...)
}

// Create returns a builder for creating a Exam entity.
func (c *ExamClient) Create() *ExamCreate {
	mutation := newExamMutation(c.config, OpCreate)
	return &ExamCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Exam entities.
func (c *ExamClient) CreateBulk(builders ...*ExamCreate) *ExamCreateBulk {
	return &ExamCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Exam.
func (c *ExamClient) Update() *ExamUpdate {
	mutation := newExamMutation(c.config, OpUpdate)
	return &ExamUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExamClient) UpdateOne(e *Exam) *ExamUpdateOne {
	mutation := newExamMutation(c.config, OpUpdateOne, withExam(e))
	return &ExamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExamClient) UpdateOneID(id uint64) *ExamUpdateOne {
	mutation := newExamMutation(c.config, OpUpdateOne, withExamID(id))
	return &ExamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Exam.
func (c *ExamClient) Delete() *ExamDelete {
	mutation := newExamMutation(c.config, OpDelete)
	return &ExamDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ExamClient) DeleteOne(e *Exam) *ExamDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ExamClient) DeleteOneID(id uint64) *ExamDeleteOne {
	builder := c.Delete().Where(exam.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExamDeleteOne{builder}
}

// Query returns a query builder for Exam.
func (c *ExamClient) Query() *ExamQuery {
	return &ExamQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeExam},
		inters: c.Interceptors(),
	}
}

// Get returns a Exam entity by its id.
func (c *ExamClient) Get(ctx context.Context, id uint64) (*Exam, error) {
	return c.Query().Where(exam.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExamClient) GetX(ctx context.Context, id uint64) *Exam {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCourses queries the courses edge of a Exam.
func (c *ExamClient) QueryCourses(e *Exam) *CourseQuery {
	query := (&CourseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(exam.Table, exam.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, exam.CoursesTable, exam.CoursesColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudents queries the students edge of a Exam.
func (c *ExamClient) QueryStudents(e *Exam) *StudentQuery {
	query := (&StudentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(exam.Table, exam.FieldID, id),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, exam.StudentsTable, exam.StudentsColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ExamClient) Hooks() []Hook {
	return c.hooks.Exam
}

// Interceptors returns the client interceptors.
func (c *ExamClient) Interceptors() []Interceptor {
	return c.inters.Exam
}

func (c *ExamClient) mutate(ctx context.Context, m *ExamMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ExamCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ExamUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ExamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ExamDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Exam mutation op: %q", m.Op())
	}
}

// StudentClient is a client for the Student schema.
type StudentClient struct {
	config
}

// NewStudentClient returns a client for the Student from the given config.
func NewStudentClient(c config) *StudentClient {
	return &StudentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `student.Hooks(f(g(h())))`.
func (c *StudentClient) Use(hooks ...Hook) {
	c.hooks.Student = append(c.hooks.Student, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `student.Intercept(f(g(h())))`.
func (c *StudentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Student = append(c.inters.Student, interceptors...)
}

// Create returns a builder for creating a Student entity.
func (c *StudentClient) Create() *StudentCreate {
	mutation := newStudentMutation(c.config, OpCreate)
	return &StudentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Student entities.
func (c *StudentClient) CreateBulk(builders ...*StudentCreate) *StudentCreateBulk {
	return &StudentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Student.
func (c *StudentClient) Update() *StudentUpdate {
	mutation := newStudentMutation(c.config, OpUpdate)
	return &StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudentClient) UpdateOne(s *Student) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudent(s))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudentClient) UpdateOneID(id uint64) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudentID(id))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Student.
func (c *StudentClient) Delete() *StudentDelete {
	mutation := newStudentMutation(c.config, OpDelete)
	return &StudentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StudentClient) DeleteOne(s *Student) *StudentDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StudentClient) DeleteOneID(id uint64) *StudentDeleteOne {
	builder := c.Delete().Where(student.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudentDeleteOne{builder}
}

// Query returns a query builder for Student.
func (c *StudentClient) Query() *StudentQuery {
	return &StudentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStudent},
		inters: c.Interceptors(),
	}
}

// Get returns a Student entity by its id.
func (c *StudentClient) Get(ctx context.Context, id uint64) (*Student, error) {
	return c.Query().Where(student.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudentClient) GetX(ctx context.Context, id uint64) *Student {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCourses queries the courses edge of a Student.
func (c *StudentClient) QueryCourses(s *Student) *CourseQuery {
	query := (&CourseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, student.CoursesTable, student.CoursesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryExams queries the exams edge of a Student.
func (c *StudentClient) QueryExams(s *Student) *ExamQuery {
	query := (&ExamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, id),
			sqlgraph.To(exam.Table, exam.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, student.ExamsTable, student.ExamsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudentClient) Hooks() []Hook {
	return c.hooks.Student
}

// Interceptors returns the client interceptors.
func (c *StudentClient) Interceptors() []Interceptor {
	return c.inters.Student
}

func (c *StudentClient) mutate(ctx context.Context, m *StudentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StudentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StudentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Student mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Course, Exam, Student []ent.Hook
	}
	inters struct {
		Course, Exam, Student []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
