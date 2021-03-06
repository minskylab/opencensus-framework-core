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
		field.Int("busyCovidBeds"),
		field.Int("availableCovidBeds"),
		field.Int("totalCovidBeds"),

		field.Int("busyChildrenUCIBeds"),
		field.Int("availableChildrenUCIBeds"),
		field.Int("totalChildrenUCIBeds"),

		field.Int("busyAdultUCIBeds"),
		field.Int("availableAdultUCIBeds"),
		field.Int("totalAdultUCIBeds"),

		field.Int("busyVentilatorBeds"),
		field.Int("availableVentilatorBeds"),
		field.Int("totalVentilatorBeds"),

		// TODO: Add more fields here.
	}
}

// Edges of the BedRecord.
func (BedRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("records", Record.Type).Ref("bedRecords"),
		edge.To("places", Place.Type),
	}
}
