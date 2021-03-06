package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InfectedRecord holds the schema definition for the InfectedRecord entity.
type InfectedRecord struct {
	ent.Schema
}

// Fields of the InfectedRecord.
func (InfectedRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pcrTotalTests"),
		field.Int("prTotalTests"),
		field.Int("agTotalTests"),

		field.Int("pcrPositiveTests"),
		field.Int("prPositiveTests"),
		field.Int("agPositiveTests"),

		// TODO: Add more fields here
	}
}

// Edges of the InfectedRecord.
func (InfectedRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("records", Record.Type).Ref("infectedRecords"),
		edge.To("places", Place.Type),
	}
}
