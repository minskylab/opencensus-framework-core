package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeathRecord holds the schema definition for the DeathRecord entity.
type DeathRecord struct {
	ent.Schema
}

// Fields of the DeathRecord.
func (DeathRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sinadefRegisters"),
		field.Int("minsaRegisters"),
		// TODO: Add more fields here
	}
}

// Edges of the DeathRecord.
func (DeathRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("records", Record.Type).Ref("deathRecords"),
		edge.To("places", Place.Type),
	}
}
