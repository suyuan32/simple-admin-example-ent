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

// CourseUpdate is the builder for updating Course entities.
type CourseUpdate struct {
	config
	hooks    []Hook
	mutation *CourseMutation
}

// Where appends a list predicates to the CourseUpdate builder.
func (cu *CourseUpdate) Where(ps ...predicate.Course) *CourseUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CourseUpdate) SetUpdatedAt(t time.Time) *CourseUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetName sets the "name" field.
func (cu *CourseUpdate) SetName(s string) *CourseUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetCode sets the "code" field.
func (cu *CourseUpdate) SetCode(i int8) *CourseUpdate {
	cu.mutation.ResetCode()
	cu.mutation.SetCode(i)
	return cu
}

// AddCode adds i to the "code" field.
func (cu *CourseUpdate) AddCode(i int8) *CourseUpdate {
	cu.mutation.AddCode(i)
	return cu
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (cu *CourseUpdate) AddStudentIDs(ids ...uint64) *CourseUpdate {
	cu.mutation.AddStudentIDs(ids...)
	return cu
}

// AddStudents adds the "students" edges to the Student entity.
func (cu *CourseUpdate) AddStudents(s ...*Student) *CourseUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddStudentIDs(ids...)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (cu *CourseUpdate) AddExamIDs(ids ...uint64) *CourseUpdate {
	cu.mutation.AddExamIDs(ids...)
	return cu
}

// AddExams adds the "exams" edges to the Exam entity.
func (cu *CourseUpdate) AddExams(e ...*Exam) *CourseUpdate {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddExamIDs(ids...)
}

// Mutation returns the CourseMutation object of the builder.
func (cu *CourseUpdate) Mutation() *CourseMutation {
	return cu.mutation
}

// ClearStudents clears all "students" edges to the Student entity.
func (cu *CourseUpdate) ClearStudents() *CourseUpdate {
	cu.mutation.ClearStudents()
	return cu
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (cu *CourseUpdate) RemoveStudentIDs(ids ...uint64) *CourseUpdate {
	cu.mutation.RemoveStudentIDs(ids...)
	return cu
}

// RemoveStudents removes "students" edges to Student entities.
func (cu *CourseUpdate) RemoveStudents(s ...*Student) *CourseUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveStudentIDs(ids...)
}

// ClearExams clears all "exams" edges to the Exam entity.
func (cu *CourseUpdate) ClearExams() *CourseUpdate {
	cu.mutation.ClearExams()
	return cu
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (cu *CourseUpdate) RemoveExamIDs(ids ...uint64) *CourseUpdate {
	cu.mutation.RemoveExamIDs(ids...)
	return cu
}

// RemoveExams removes "exams" edges to Exam entities.
func (cu *CourseUpdate) RemoveExams(e ...*Exam) *CourseUpdate {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveExamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CourseUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, CourseMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CourseUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CourseUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CourseUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CourseUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := course.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(course.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(course.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Code(); ok {
		_spec.SetField(course.FieldCode, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedCode(); ok {
		_spec.AddField(course.FieldCode, field.TypeInt8, value)
	}
	if cu.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   course.StudentsTable,
			Columns: course.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !cu.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   course.StudentsTable,
			Columns: course.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   course.StudentsTable,
			Columns: course.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   course.ExamsTable,
			Columns: []string{course.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedExamsIDs(); len(nodes) > 0 && !cu.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   course.ExamsTable,
			Columns: []string{course.ExamsColumn},
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
	if nodes := cu.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   course.ExamsTable,
			Columns: []string{course.ExamsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CourseUpdateOne is the builder for updating a single Course entity.
type CourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CourseUpdateOne) SetUpdatedAt(t time.Time) *CourseUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CourseUpdateOne) SetName(s string) *CourseUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetCode sets the "code" field.
func (cuo *CourseUpdateOne) SetCode(i int8) *CourseUpdateOne {
	cuo.mutation.ResetCode()
	cuo.mutation.SetCode(i)
	return cuo
}

// AddCode adds i to the "code" field.
func (cuo *CourseUpdateOne) AddCode(i int8) *CourseUpdateOne {
	cuo.mutation.AddCode(i)
	return cuo
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (cuo *CourseUpdateOne) AddStudentIDs(ids ...uint64) *CourseUpdateOne {
	cuo.mutation.AddStudentIDs(ids...)
	return cuo
}

// AddStudents adds the "students" edges to the Student entity.
func (cuo *CourseUpdateOne) AddStudents(s ...*Student) *CourseUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddStudentIDs(ids...)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (cuo *CourseUpdateOne) AddExamIDs(ids ...uint64) *CourseUpdateOne {
	cuo.mutation.AddExamIDs(ids...)
	return cuo
}

// AddExams adds the "exams" edges to the Exam entity.
func (cuo *CourseUpdateOne) AddExams(e ...*Exam) *CourseUpdateOne {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddExamIDs(ids...)
}

// Mutation returns the CourseMutation object of the builder.
func (cuo *CourseUpdateOne) Mutation() *CourseMutation {
	return cuo.mutation
}

// ClearStudents clears all "students" edges to the Student entity.
func (cuo *CourseUpdateOne) ClearStudents() *CourseUpdateOne {
	cuo.mutation.ClearStudents()
	return cuo
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (cuo *CourseUpdateOne) RemoveStudentIDs(ids ...uint64) *CourseUpdateOne {
	cuo.mutation.RemoveStudentIDs(ids...)
	return cuo
}

// RemoveStudents removes "students" edges to Student entities.
func (cuo *CourseUpdateOne) RemoveStudents(s ...*Student) *CourseUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveStudentIDs(ids...)
}

// ClearExams clears all "exams" edges to the Exam entity.
func (cuo *CourseUpdateOne) ClearExams() *CourseUpdateOne {
	cuo.mutation.ClearExams()
	return cuo
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (cuo *CourseUpdateOne) RemoveExamIDs(ids ...uint64) *CourseUpdateOne {
	cuo.mutation.RemoveExamIDs(ids...)
	return cuo
}

// RemoveExams removes "exams" edges to Exam entities.
func (cuo *CourseUpdateOne) RemoveExams(e ...*Exam) *CourseUpdateOne {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveExamIDs(ids...)
}

// Where appends a list predicates to the CourseUpdate builder.
func (cuo *CourseUpdateOne) Where(ps ...predicate.Course) *CourseUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CourseUpdateOne) Select(field string, fields ...string) *CourseUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Course entity.
func (cuo *CourseUpdateOne) Save(ctx context.Context) (*Course, error) {
	cuo.defaults()
	return withHooks[*Course, CourseMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CourseUpdateOne) SaveX(ctx context.Context) *Course {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CourseUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CourseUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CourseUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := course.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CourseUpdateOne) sqlSave(ctx context.Context) (_node *Course, err error) {
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Course.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, course.FieldID)
		for _, f := range fields {
			if !course.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != course.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(course.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(course.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Code(); ok {
		_spec.SetField(course.FieldCode, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedCode(); ok {
		_spec.AddField(course.FieldCode, field.TypeInt8, value)
	}
	if cuo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   course.StudentsTable,
			Columns: course.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !cuo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   course.StudentsTable,
			Columns: course.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   course.StudentsTable,
			Columns: course.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   course.ExamsTable,
			Columns: []string{course.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedExamsIDs(); len(nodes) > 0 && !cuo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   course.ExamsTable,
			Columns: []string{course.ExamsColumn},
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
	if nodes := cuo.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   course.ExamsTable,
			Columns: []string{course.ExamsColumn},
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
	_node = &Course{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
