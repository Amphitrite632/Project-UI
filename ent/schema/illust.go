package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Illust holds the schema definition for the Illust entity.
type Illust struct {
	ent.Schema
}

// Fields of the Illust.
func (Illust) Fields() []ent.Field {
	return []ent.Field{
		field.String("illustID").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("hash").
			NotEmpty().
			Immutable(),
		field.String("originalURL"),
		field.Float("heightRatio").
			Positive(),
		field.Time("createdAt").
			Default(time.Now()),
	}
}

// Edges of the Illust.
func (Illust) Edges() []ent.Edge {
	return nil
}
