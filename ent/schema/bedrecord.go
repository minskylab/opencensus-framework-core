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
		// TODO: Add more fields here
	}
}

// Edges of the BedRecord.
func (BedRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("bedRecords"),
	}
}
