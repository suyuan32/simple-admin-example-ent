package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int8("age"),
		field.Time("birth_at"),
	}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("courses", Course.Type),
		edge.To("exams", Exam.Type),
	}
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "students"},
	}
}
