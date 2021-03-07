package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Occurency holds the schema definition for the Occurency entity.
type Occurency struct {
	ent.Schema
}

// Fields of the Occurency.
func (Occurency) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.Time("reportedRecord"),
		field.Time("resultDate"),
		field.String("biologicalSex"),
		field.Int("age"),
	}
}

// Edges of the Occurency.
func (Occurency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("region", Region.Type).Unique(),
		edge.To("province", Province.Type).Unique(),
		edge.To("district", District.Type).Unique(),
	}
}
