// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-example-api/ent/course"
	"github.com/suyuan32/simple-admin-example-api/ent/exam"
	"github.com/suyuan32/simple-admin-example-api/ent/predicate"
	"github.com/suyuan32/simple-admin-example-api/ent/student"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StudentUpdate) SetUpdatedAt(t time.Time) *StudentUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetAge sets the "age" field.
func (su *StudentUpdate) SetAge(i int8) *StudentUpdate {
	su.mutation.ResetAge()
	su.mutation.SetAge(i)
	return su
}

// AddAge adds i to the "age" field.
func (su *StudentUpdate) AddAge(i int8) *StudentUpdate {
	su.mutation.AddAge(i)
	return su
}

// SetBirthAt sets the "birth_at" field.
func (su *StudentUpdate) SetBirthAt(t time.Time) *StudentUpdate {
	su.mutation.SetBirthAt(t)
	return su
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (su *StudentUpdate) AddCourseIDs(ids ...uint64) *StudentUpdate {
	su.mutation.AddCourseIDs(ids...)
	return su
}

// AddCourses adds the "courses" edges to the Course entity.
func (su *StudentUpdate) AddCourses(c ...*Course) *StudentUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCourseIDs(ids...)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (su *StudentUpdate) AddExamIDs(ids ...uint64) *StudentUpdate {
	su.mutation.AddExamIDs(ids...)
	return su
}

// AddExams adds the "exams" edges to the Exam entity.
func (su *StudentUpdate) AddExams(e ...*Exam) *StudentUpdate {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddExamIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (su *StudentUpdate) ClearCourses() *StudentUpdate {
	su.mutation.ClearCourses()
	return su
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (su *StudentUpdate) RemoveCourseIDs(ids ...uint64) *StudentUpdate {
	su.mutation.RemoveCourseIDs(ids...)
	return su
}

// RemoveCourses removes "courses" edges to Course entities.
func (su *StudentUpdate) RemoveCourses(c ...*Course) *StudentUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCourseIDs(ids...)
}

// ClearExams clears all "exams" edges to the Exam entity.
func (su *StudentUpdate) ClearExams() *StudentUpdate {
	su.mutation.ClearExams()
	return su
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (su *StudentUpdate) RemoveExamIDs(ids ...uint64) *StudentUpdate {
	su.mutation.RemoveExamIDs(ids...)
	return su
}

// RemoveExams removes "exams" edges to Exam entities.
func (su *StudentUpdate) RemoveExams(e ...*Exam) *StudentUpdate {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveExamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StudentUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AddedAge(); ok {
		_spec.AddField(student.FieldAge, field.TypeInt8, value)
	}
	if value, ok := su.mutation.BirthAt(); ok {
		_spec.SetField(student.FieldBirthAt, field.TypeTime, value)
	}
	if su.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !su.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ExamsTable,
			Columns: []string{student.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedExamsIDs(); len(nodes) > 0 && !su.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ExamsTable,
			Columns: []string{student.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ExamsTable,
			Columns: []string{student.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StudentUpdateOne) SetUpdatedAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetAge sets the "age" field.
func (suo *StudentUpdateOne) SetAge(i int8) *StudentUpdateOne {
	suo.mutation.ResetAge()
	suo.mutation.SetAge(i)
	return suo
}

// AddAge adds i to the "age" field.
func (suo *StudentUpdateOne) AddAge(i int8) *StudentUpdateOne {
	suo.mutation.AddAge(i)
	return suo
}

// SetBirthAt sets the "birth_at" field.
func (suo *StudentUpdateOne) SetBirthAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetBirthAt(t)
	return suo
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (suo *StudentUpdateOne) AddCourseIDs(ids ...uint64) *StudentUpdateOne {
	suo.mutation.AddCourseIDs(ids...)
	return suo
}

// AddCourses adds the "courses" edges to the Course entity.
func (suo *StudentUpdateOne) AddCourses(c ...*Course) *StudentUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCourseIDs(ids...)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (suo *StudentUpdateOne) AddExamIDs(ids ...uint64) *StudentUpdateOne {
	suo.mutation.AddExamIDs(ids...)
	return suo
}

// AddExams adds the "exams" edges to the Exam entity.
func (suo *StudentUpdateOne) AddExams(e ...*Exam) *StudentUpdateOne {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddExamIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (suo *StudentUpdateOne) ClearCourses() *StudentUpdateOne {
	suo.mutation.ClearCourses()
	return suo
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (suo *StudentUpdateOne) RemoveCourseIDs(ids ...uint64) *StudentUpdateOne {
	suo.mutation.RemoveCourseIDs(ids...)
	return suo
}

// RemoveCourses removes "courses" edges to Course entities.
func (suo *StudentUpdateOne) RemoveCourses(c ...*Course) *StudentUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCourseIDs(ids...)
}

// ClearExams clears all "exams" edges to the Exam entity.
func (suo *StudentUpdateOne) ClearExams() *StudentUpdateOne {
	suo.mutation.ClearExams()
	return suo
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (suo *StudentUpdateOne) RemoveExamIDs(ids ...uint64) *StudentUpdateOne {
	suo.mutation.RemoveExamIDs(ids...)
	return suo
}

// RemoveExams removes "exams" edges to Exam entities.
func (suo *StudentUpdateOne) RemoveExams(e ...*Exam) *StudentUpdateOne {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveExamIDs(ids...)
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StudentUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AddedAge(); ok {
		_spec.AddField(student.FieldAge, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.BirthAt(); ok {
		_spec.SetField(student.FieldBirthAt, field.TypeTime, value)
	}
	if suo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !suo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ExamsTable,
			Columns: []string{student.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedExamsIDs(); len(nodes) > 0 && !suo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ExamsTable,
			Columns: []string{student.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ExamsTable,
			Columns: []string{student.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
