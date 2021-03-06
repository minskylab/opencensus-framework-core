package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OxygenRecord holds the schema definition for the OxygenRecord entity.
type OxygenRecord struct {
	ent.Schema
}

// Fields of the OxygenRecord.
func (OxygenRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("totalCylinders"),
		field.Int("totalOwnCylinders"),

		field.Int("dailyProduction"),
		field.Int("maxDailyProduction"),
		field.Int("dailyConsumption"),

		field.String("mainSourceKind"), // Plant, Tank criog, isotank, generator, network

		// TODO: Add more fields here
	}
}

// Edges of the OxygenRecord.
func (OxygenRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("records", Record.Type).Ref("oxygenRecords"),
		edge.To("places", Place.Type),
	}
}
