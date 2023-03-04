package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// IllustTag holds the schema definition for the IllustTag entity.
type IllustTag struct {
	ent.Schema
}

// Fields of the IllustTag.
func (IllustTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("parentIllustID").
			NotEmpty().
			Immutable(),
		field.String("prentTagID").
			NotEmpty().
			Immutable(),
		field.String("illustTagID").
			NotEmpty().
			Unique().
			Immutable(),
	}
}

// Edges of the IllustTag.
func (IllustTag) Edges() []ent.Edge {
	return nil
}
