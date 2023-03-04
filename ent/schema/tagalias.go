package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TagAlias holds the schema definition for the TagAlias entity.
type TagAlias struct {
	ent.Schema
}

// Fields of the TagAlias.
func (TagAlias) Fields() []ent.Field {
	return []ent.Field{
		field.String("parentTagID").
			NotEmpty().
			Immutable(),
		field.String("alias").
			NotEmpty(),
		field.String("tagAliasID").
			NotEmpty().
			Unique(),
	}
}

// Edges of the TagAlias.
func (TagAlias) Edges() []ent.Edge {
	return nil
}
