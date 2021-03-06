package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BedRecord holds the schema definition for the BedRecord entity.
type BedRecord struct {
	ent.Schema
}

// Fields of the BedRecord.
func (BedRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Time("reportedDate"),
		field.Time("collectedDate"),

		field.Int("busyBeds"),
		field.Int("availableBeds"),
		field.Int("totalBeds"),

		field.String("kindBed"), // UCI, CovidBed, NoCovidBed, Ventilator
		field.String("kindAge"),

		// TODO: Add more fields here.
	}
}

// Edges of the BedRecord.
func (BedRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("places", Place.Type),
	}
}
