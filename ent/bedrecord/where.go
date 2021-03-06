// Code generated by entc, DO NOT EDIT.

package bedrecord

import (
	"opencensus/core/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BusyCovidBeds applies equality check predicate on the "busyCovidBeds" field. It's identical to BusyCovidBedsEQ.
func BusyCovidBeds(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBusyCovidBeds), v))
	})
}

// AvailableCovidBeds applies equality check predicate on the "availableCovidBeds" field. It's identical to AvailableCovidBedsEQ.
func AvailableCovidBeds(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvailableCovidBeds), v))
	})
}

// BusyCovidBedsEQ applies the EQ predicate on the "busyCovidBeds" field.
func BusyCovidBedsEQ(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBusyCovidBeds), v))
	})
}

// BusyCovidBedsNEQ applies the NEQ predicate on the "busyCovidBeds" field.
func BusyCovidBedsNEQ(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBusyCovidBeds), v))
	})
}

// BusyCovidBedsIn applies the In predicate on the "busyCovidBeds" field.
func BusyCovidBedsIn(vs ...int) predicate.BedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBusyCovidBeds), v...))
	})
}

// BusyCovidBedsNotIn applies the NotIn predicate on the "busyCovidBeds" field.
func BusyCovidBedsNotIn(vs ...int) predicate.BedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBusyCovidBeds), v...))
	})
}

// BusyCovidBedsGT applies the GT predicate on the "busyCovidBeds" field.
func BusyCovidBedsGT(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBusyCovidBeds), v))
	})
}

// BusyCovidBedsGTE applies the GTE predicate on the "busyCovidBeds" field.
func BusyCovidBedsGTE(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBusyCovidBeds), v))
	})
}

// BusyCovidBedsLT applies the LT predicate on the "busyCovidBeds" field.
func BusyCovidBedsLT(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBusyCovidBeds), v))
	})
}

// BusyCovidBedsLTE applies the LTE predicate on the "busyCovidBeds" field.
func BusyCovidBedsLTE(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBusyCovidBeds), v))
	})
}

// AvailableCovidBedsEQ applies the EQ predicate on the "availableCovidBeds" field.
func AvailableCovidBedsEQ(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvailableCovidBeds), v))
	})
}

// AvailableCovidBedsNEQ applies the NEQ predicate on the "availableCovidBeds" field.
func AvailableCovidBedsNEQ(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvailableCovidBeds), v))
	})
}

// AvailableCovidBedsIn applies the In predicate on the "availableCovidBeds" field.
func AvailableCovidBedsIn(vs ...int) predicate.BedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvailableCovidBeds), v...))
	})
}

// AvailableCovidBedsNotIn applies the NotIn predicate on the "availableCovidBeds" field.
func AvailableCovidBedsNotIn(vs ...int) predicate.BedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvailableCovidBeds), v...))
	})
}

// AvailableCovidBedsGT applies the GT predicate on the "availableCovidBeds" field.
func AvailableCovidBedsGT(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvailableCovidBeds), v))
	})
}

// AvailableCovidBedsGTE applies the GTE predicate on the "availableCovidBeds" field.
func AvailableCovidBedsGTE(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvailableCovidBeds), v))
	})
}

// AvailableCovidBedsLT applies the LT predicate on the "availableCovidBeds" field.
func AvailableCovidBedsLT(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvailableCovidBeds), v))
	})
}

// AvailableCovidBedsLTE applies the LTE predicate on the "availableCovidBeds" field.
func AvailableCovidBedsLTE(v int) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvailableCovidBeds), v))
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrganizationTable, OrganizationPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrganizationTable, OrganizationPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BedRecord) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BedRecord) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BedRecord) predicate.BedRecord {
	return predicate.BedRecord(func(s *sql.Selector) {
		p(s.Not())
	})
}
