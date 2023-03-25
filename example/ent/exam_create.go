// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-example-api/ent/course"
	"github.com/suyuan32/simple-admin-example-api/ent/exam"
	"github.com/suyuan32/simple-admin-example-api/ent/student"
)

// ExamCreate is the builder for creating a Exam entity.
type ExamCreate struct {
	config
	mutation *ExamMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ec *ExamCreate) SetCreatedAt(t time.Time) *ExamCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *ExamCreate) SetNillableCreatedAt(t *time.Time) *ExamCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *ExamCreate) SetUpdatedAt(t time.Time) *ExamCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *ExamCreate) SetNillableUpdatedAt(t *time.Time) *ExamCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetScore sets the "score" field.
func (ec *ExamCreate) SetScore(u uint8) *ExamCreate {
	ec.mutation.SetScore(u)
	return ec
}

// SetID sets the "id" field.
func (ec *ExamCreate) SetID(u uint64) *ExamCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetCoursesID sets the "courses" edge to the Course entity by ID.
func (ec *ExamCreate) SetCoursesID(id uint64) *ExamCreate {
	ec.mutation.SetCoursesID(id)
	return ec
}

// SetNillableCoursesID sets the "courses" edge to the Course entity by ID if the given value is not nil.
func (ec *ExamCreate) SetNillableCoursesID(id *uint64) *ExamCreate {
	if id != nil {
		ec = ec.SetCoursesID(*id)
	}
	return ec
}

// SetCourses sets the "courses" edge to the Course entity.
func (ec *ExamCreate) SetCourses(c *Course) *ExamCreate {
	return ec.SetCoursesID(c.ID)
}

// SetStudentsID sets the "students" edge to the Student entity by ID.
func (ec *ExamCreate) SetStudentsID(id uint64) *ExamCreate {
	ec.mutation.SetStudentsID(id)
	return ec
}

// SetNillableStudentsID sets the "students" edge to the Student entity by ID if the given value is not nil.
func (ec *ExamCreate) SetNillableStudentsID(id *uint64) *ExamCreate {
	if id != nil {
		ec = ec.SetStudentsID(*id)
	}
	return ec
}

// SetStudents sets the "students" edge to the Student entity.
func (ec *ExamCreate) SetStudents(s *Student) *ExamCreate {
	return ec.SetStudentsID(s.ID)
}

// Mutation returns the ExamMutation object of the builder.
func (ec *ExamCreate) Mutation() *ExamMutation {
	return ec.mutation
}

// Save creates the Exam in the database.
func (ec *ExamCreate) Save(ctx context.Context) (*Exam, error) {
	ec.defaults()
	return withHooks[*Exam, ExamMutation](ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExamCreate) SaveX(ctx context.Context) *Exam {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *ExamCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *ExamCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *ExamCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := exam.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := exam.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExamCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Exam.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Exam.updated_at"`)}
	}
	if _, ok := ec.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Exam.score"`)}
	}
	return nil
}

func (ec *ExamCreate) sqlSave(ctx context.Context) (*Exam, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *ExamCreate) createSpec() (*Exam, *sqlgraph.CreateSpec) {
	var (
		_node = &Exam{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(exam.Table, sqlgraph.NewFieldSpec(exam.FieldID, field.TypeUint64))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(exam.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(exam.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.Score(); ok {
		_spec.SetField(exam.FieldScore, field.TypeUint8, value)
		_node.Score = value
	}
	if nodes := ec.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   exam.CoursesTable,
			Columns: []string{exam.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_courses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exam.StudentsTable,
			Columns: []string{exam.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.student_exams = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExamCreateBulk is the builder for creating many Exam entities in bulk.
type ExamCreateBulk struct {
	config
	builders []*ExamCreate
}

// Save creates the Exam entities in the database.
func (ecb *ExamCreateBulk) Save(ctx context.Context) ([]*Exam, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Exam, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExamMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *ExamCreateBulk) SaveX(ctx context.Context) []*Exam {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *ExamCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *ExamCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}