package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Place holds the schema definition for the Place entity.
type Place struct {
	ent.Schema
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("kind"),
		field.String("ubigeo").Optional(),
		field.Bool("covidZone").Optional(),
		field.Float("lat").Optional(),
		field.Float("lon").Optional(),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("records", Record.Type).Ref("places"),
		edge.From("bedRecords", BedRecord.Type).Ref("places"),
		edge.From("deathRecords", DeathRecord.Type).Ref("places"),
		edge.From("infectedRecords", InfectedRecord.Type).Ref("places"),

		edge.To("institutions", Institution.Type),
		edge.To("regions", Region.Type),
	}
}
