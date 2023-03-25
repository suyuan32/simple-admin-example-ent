// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-example-api/ent/course"
	"github.com/suyuan32/simple-admin-example-api/ent/exam"
	"github.com/suyuan32/simple-admin-example-api/ent/student"
)

// Exam is the model entity for the Exam schema.
type Exam struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Score holds the value of the "score" field.
	Score uint8 `json:"score,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExamQuery when eager-loading is set.
	Edges         ExamEdges `json:"edges"`
	exam_courses  *uint64
	student_exams *uint64
}

// ExamEdges holds the relations/edges for other nodes in the graph.
type ExamEdges struct {
	// Courses holds the value of the courses edge.
	Courses *Course `json:"courses,omitempty"`
	// Students holds the value of the students edge.
	Students *Student `json:"students,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CoursesOrErr returns the Courses value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExamEdges) CoursesOrErr() (*Course, error) {
	if e.loadedTypes[0] {
		if e.Courses == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Courses, nil
	}
	return nil, &NotLoadedError{edge: "courses"}
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExamEdges) StudentsOrErr() (*Student, error) {
	if e.loadedTypes[1] {
		if e.Students == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exam) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exam.FieldID, exam.FieldScore:
			values[i] = new(sql.NullInt64)
		case exam.FieldCreatedAt, exam.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case exam.ForeignKeys[0]: // exam_courses
			values[i] = new(sql.NullInt64)
		case exam.ForeignKeys[1]: // student_exams
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Exam", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exam fields.
func (e *Exam) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exam.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = uint64(value.Int64)
		case exam.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case exam.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case exam.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				e.Score = uint8(value.Int64)
			}
		case exam.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_courses", value)
			} else if value.Valid {
				e.exam_courses = new(uint64)
				*e.exam_courses = uint64(value.Int64)
			}
		case exam.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field student_exams", value)
			} else if value.Valid {
				e.student_exams = new(uint64)
				*e.student_exams = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryCourses queries the "courses" edge of the Exam entity.
func (e *Exam) QueryCourses() *CourseQuery {
	return NewExamClient(e.config).QueryCourses(e)
}

// QueryStudents queries the "students" edge of the Exam entity.
func (e *Exam) QueryStudents() *StudentQuery {
	return NewExamClient(e.config).QueryStudents(e)
}

// Update returns a builder for updating this Exam.
// Note that you need to call Exam.Unwrap() before calling this method if this Exam
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exam) Update() *ExamUpdateOne {
	return NewExamClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Exam entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exam) Unwrap() *Exam {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exam is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exam) String() string {
	var builder strings.Builder
	builder.WriteString("Exam(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", e.Score))
	builder.WriteByte(')')
	return builder.String()
}

// Exams is a parsable slice of Exam.
type Exams []*Exam