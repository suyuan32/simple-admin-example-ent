package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Course struct {
	ent.Schema
}

func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int8("code"),
	}
}

func (Course) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("students", Student.Type).Ref("courses"),
		edge.From("exams", Exam.Type).Ref("courses"),
	}
}
