package schema

import "entgo.io/ent"

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		//TODO: 書く
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
