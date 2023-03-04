package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("tagID").
			NotEmpty().
			Unique().
			Immutable(),
		field.Enum("type").
			Values(
				"WORK",
				"CHARACTER",
				"AUTHOR",
			),
		field.String("officialName").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
