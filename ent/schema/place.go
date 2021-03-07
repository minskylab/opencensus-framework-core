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
		field.String("kind"),
		field.String("name"),
		field.String("politic").Optional(), // private or public
		field.String("ubigeo").Optional(),
		field.Bool("covidZone").Optional(),
		field.Float("lat").Optional(),
		field.Float("lon").Optional(),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("oxygenrecords", OxygenRecord.Type).Ref("places"),
		edge.From("bedRecords", BedRecord.Type).Ref("places"),
		edge.From("deathRecords", DeathRecord.Type).Ref("places"),
		edge.From("infectedRecords", InfectedRecord.Type).Ref("places"),

		edge.To("region", Region.Type).Unique(),
		edge.To("province", Province.Type).Unique(),
		edge.To("district", District.Type).Unique(),
	}
}
