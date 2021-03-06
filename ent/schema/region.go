package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Region holds the schema definition for the Region entity.
type Region struct {
	ent.Schema
}

// Fields of the Region.
func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("places", Place.Type).Ref("regions"),
		edge.To("provinces", Province.Type),
	}
}
