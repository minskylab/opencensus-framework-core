package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("code"),
		field.String("ubigeo"),
		field.String("kind").Optional(),
		field.Bool("covidZone").Optional(),
		field.String("category").Optional(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("region", Region.Type),
		edge.To("province", Province.Type),
		edge.To("district", District.Type),
		edge.To("oxygenRecords", OxygenRecord.Type),
		edge.To("bedRecords", BedRecord.Type),
	}
}
