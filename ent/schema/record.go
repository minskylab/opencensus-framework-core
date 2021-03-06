package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Time("reportedDate"),
		field.Time("collectedDate"),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("oxygenRecords", OxygenRecord.Type),
		edge.To("bedRecords", BedRecord.Type),
		edge.To("deathRecords", DeathRecord.Type),
		edge.To("infectedRecords", InfectedRecord.Type),
	}
}
