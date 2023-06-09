// Code generated by ent, DO NOT EDIT.

package course

import (
	"time"
)

const (
	// Label holds the string label denoting the course type in the database.
	Label = "course"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// EdgeStudents holds the string denoting the students edge name in mutations.
	EdgeStudents = "students"
	// EdgeExams holds the string denoting the exams edge name in mutations.
	EdgeExams = "exams"
	// Table holds the table name of the course in the database.
	Table = "courses"
	// StudentsTable is the table that holds the students relation/edge. The primary key declared below.
	StudentsTable = "student_courses"
	// StudentsInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentsInverseTable = "students"
	// ExamsTable is the table that holds the exams relation/edge.
	ExamsTable = "exams"
	// ExamsInverseTable is the table name for the Exam entity.
	// It exists in this package in order to avoid circular dependency with the "exam" package.
	ExamsInverseTable = "exams"
	// ExamsColumn is the table column denoting the exams relation/edge.
	ExamsColumn = "exam_courses"
)

// Columns holds all SQL columns for course fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCode,
}

var (
	// StudentsPrimaryKey and StudentsColumn2 are the table columns denoting the
	// primary key for the students relation (M2M).
	StudentsPrimaryKey = []string{"student_id", "course_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
