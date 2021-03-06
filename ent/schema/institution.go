package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Institution holds the schema definition for the Institution entity.
type Institution struct {
	ent.Schema
}

// Fields of the Institution.
func (Institution) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type").Optional(),    // Hospital, Militar, SIS
		field.String("politic").Optional(), // publico o privado
	}
}

// Edges of the Institution.
func (Institution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("places", Place.Type).Ref("institutions"),
	}
}
