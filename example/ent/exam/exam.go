// Code generated by ent, DO NOT EDIT.

package exam

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the exam type in the database.
	Label = "exam"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// EdgeCourses holds the string denoting the courses edge name in mutations.
	EdgeCourses = "courses"
	// EdgeStudents holds the string denoting the students edge name in mutations.
	EdgeStudents = "students"
	// Table holds the table name of the exam in the database.
	Table = "exams"
	// CoursesTable is the table that holds the courses relation/edge.
	CoursesTable = "exams"
	// CoursesInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CoursesInverseTable = "courses"
	// CoursesColumn is the table column denoting the courses relation/edge.
	CoursesColumn = "exam_courses"
	// StudentsTable is the table that holds the students relation/edge.
	StudentsTable = "exams"
	// StudentsInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentsInverseTable = "students"
	// StudentsColumn is the table column denoting the students relation/edge.
	StudentsColumn = "student_exams"
)

// Columns holds all SQL columns for exam fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldScore,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "exams"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_courses",
	"student_exams",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Exam queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByCoursesField orders the results by courses field.
func ByCoursesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCoursesStep(), sql.OrderByField(field, opts...))
	}
}

// ByStudentsField orders the results by students field.
func ByStudentsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentsStep(), sql.OrderByField(field, opts...))
	}
}
func newCoursesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CoursesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CoursesTable, CoursesColumn),
	)
}
func newStudentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StudentsTable, StudentsColumn),
	)
}
