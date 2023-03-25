package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Exam struct {
	ent.Schema
}

func (Exam) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("score"),
	}
}

func (Exam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Exam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("courses", Course.Type).Unique(),
		edge.From("students", Student.Type).Ref("exams").Unique(),
	}
}
