package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TagChild holds the schema definition for the TagChild entity.
type TagChild struct {
	ent.Schema
}

// Fields of the TagChild.
func (TagChild) Fields() []ent.Field {
	return []ent.Field{
		field.String("parentTagID").
			NotEmpty().
			Immutable(),
		field.String("childTagID").
			NotEmpty().
			Immutable(),
		field.String("tagChildID").
			NotEmpty().
			Unique(),
	}
}

// Edges of the TagChild.
func (TagChild) Edges() []ent.Edge {
	return nil
}
