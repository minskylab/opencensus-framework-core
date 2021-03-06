// Code generated by entc, DO NOT EDIT.

package infectedrecord

import (
	"opencensus/core/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
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
func IDGT(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ReportedDate applies equality check predicate on the "reportedDate" field. It's identical to ReportedDateEQ.
func ReportedDate(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReportedDate), v))
	})
}

// CollectedDate applies equality check predicate on the "collectedDate" field. It's identical to CollectedDateEQ.
func CollectedDate(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCollectedDate), v))
	})
}

// PcrTotalTests applies equality check predicate on the "pcrTotalTests" field. It's identical to PcrTotalTestsEQ.
func PcrTotalTests(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPcrTotalTests), v))
	})
}

// PrTotalTests applies equality check predicate on the "prTotalTests" field. It's identical to PrTotalTestsEQ.
func PrTotalTests(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrTotalTests), v))
	})
}

// AgTotalTests applies equality check predicate on the "agTotalTests" field. It's identical to AgTotalTestsEQ.
func AgTotalTests(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgTotalTests), v))
	})
}

// PcrPositiveTests applies equality check predicate on the "pcrPositiveTests" field. It's identical to PcrPositiveTestsEQ.
func PcrPositiveTests(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPcrPositiveTests), v))
	})
}

// PrPositiveTests applies equality check predicate on the "prPositiveTests" field. It's identical to PrPositiveTestsEQ.
func PrPositiveTests(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrPositiveTests), v))
	})
}

// AgPositiveTests applies equality check predicate on the "agPositiveTests" field. It's identical to AgPositiveTestsEQ.
func AgPositiveTests(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgPositiveTests), v))
	})
}

// ReportedDateEQ applies the EQ predicate on the "reportedDate" field.
func ReportedDateEQ(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReportedDate), v))
	})
}

// ReportedDateNEQ applies the NEQ predicate on the "reportedDate" field.
func ReportedDateNEQ(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReportedDate), v))
	})
}

// ReportedDateIn applies the In predicate on the "reportedDate" field.
func ReportedDateIn(vs ...time.Time) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReportedDate), v...))
	})
}

// ReportedDateNotIn applies the NotIn predicate on the "reportedDate" field.
func ReportedDateNotIn(vs ...time.Time) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReportedDate), v...))
	})
}

// ReportedDateGT applies the GT predicate on the "reportedDate" field.
func ReportedDateGT(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReportedDate), v))
	})
}

// ReportedDateGTE applies the GTE predicate on the "reportedDate" field.
func ReportedDateGTE(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReportedDate), v))
	})
}

// ReportedDateLT applies the LT predicate on the "reportedDate" field.
func ReportedDateLT(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReportedDate), v))
	})
}

// ReportedDateLTE applies the LTE predicate on the "reportedDate" field.
func ReportedDateLTE(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReportedDate), v))
	})
}

// CollectedDateEQ applies the EQ predicate on the "collectedDate" field.
func CollectedDateEQ(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCollectedDate), v))
	})
}

// CollectedDateNEQ applies the NEQ predicate on the "collectedDate" field.
func CollectedDateNEQ(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCollectedDate), v))
	})
}

// CollectedDateIn applies the In predicate on the "collectedDate" field.
func CollectedDateIn(vs ...time.Time) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCollectedDate), v...))
	})
}

// CollectedDateNotIn applies the NotIn predicate on the "collectedDate" field.
func CollectedDateNotIn(vs ...time.Time) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCollectedDate), v...))
	})
}

// CollectedDateGT applies the GT predicate on the "collectedDate" field.
func CollectedDateGT(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCollectedDate), v))
	})
}

// CollectedDateGTE applies the GTE predicate on the "collectedDate" field.
func CollectedDateGTE(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCollectedDate), v))
	})
}

// CollectedDateLT applies the LT predicate on the "collectedDate" field.
func CollectedDateLT(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCollectedDate), v))
	})
}

// CollectedDateLTE applies the LTE predicate on the "collectedDate" field.
func CollectedDateLTE(v time.Time) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCollectedDate), v))
	})
}

// PcrTotalTestsEQ applies the EQ predicate on the "pcrTotalTests" field.
func PcrTotalTestsEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPcrTotalTests), v))
	})
}

// PcrTotalTestsNEQ applies the NEQ predicate on the "pcrTotalTests" field.
func PcrTotalTestsNEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPcrTotalTests), v))
	})
}

// PcrTotalTestsIn applies the In predicate on the "pcrTotalTests" field.
func PcrTotalTestsIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPcrTotalTests), v...))
	})
}

// PcrTotalTestsNotIn applies the NotIn predicate on the "pcrTotalTests" field.
func PcrTotalTestsNotIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPcrTotalTests), v...))
	})
}

// PcrTotalTestsGT applies the GT predicate on the "pcrTotalTests" field.
func PcrTotalTestsGT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPcrTotalTests), v))
	})
}

// PcrTotalTestsGTE applies the GTE predicate on the "pcrTotalTests" field.
func PcrTotalTestsGTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPcrTotalTests), v))
	})
}

// PcrTotalTestsLT applies the LT predicate on the "pcrTotalTests" field.
func PcrTotalTestsLT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPcrTotalTests), v))
	})
}

// PcrTotalTestsLTE applies the LTE predicate on the "pcrTotalTests" field.
func PcrTotalTestsLTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPcrTotalTests), v))
	})
}

// PrTotalTestsEQ applies the EQ predicate on the "prTotalTests" field.
func PrTotalTestsEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrTotalTests), v))
	})
}

// PrTotalTestsNEQ applies the NEQ predicate on the "prTotalTests" field.
func PrTotalTestsNEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrTotalTests), v))
	})
}

// PrTotalTestsIn applies the In predicate on the "prTotalTests" field.
func PrTotalTestsIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrTotalTests), v...))
	})
}

// PrTotalTestsNotIn applies the NotIn predicate on the "prTotalTests" field.
func PrTotalTestsNotIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrTotalTests), v...))
	})
}

// PrTotalTestsGT applies the GT predicate on the "prTotalTests" field.
func PrTotalTestsGT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrTotalTests), v))
	})
}

// PrTotalTestsGTE applies the GTE predicate on the "prTotalTests" field.
func PrTotalTestsGTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrTotalTests), v))
	})
}

// PrTotalTestsLT applies the LT predicate on the "prTotalTests" field.
func PrTotalTestsLT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrTotalTests), v))
	})
}

// PrTotalTestsLTE applies the LTE predicate on the "prTotalTests" field.
func PrTotalTestsLTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrTotalTests), v))
	})
}

// AgTotalTestsEQ applies the EQ predicate on the "agTotalTests" field.
func AgTotalTestsEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgTotalTests), v))
	})
}

// AgTotalTestsNEQ applies the NEQ predicate on the "agTotalTests" field.
func AgTotalTestsNEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgTotalTests), v))
	})
}

// AgTotalTestsIn applies the In predicate on the "agTotalTests" field.
func AgTotalTestsIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAgTotalTests), v...))
	})
}

// AgTotalTestsNotIn applies the NotIn predicate on the "agTotalTests" field.
func AgTotalTestsNotIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAgTotalTests), v...))
	})
}

// AgTotalTestsGT applies the GT predicate on the "agTotalTests" field.
func AgTotalTestsGT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgTotalTests), v))
	})
}

// AgTotalTestsGTE applies the GTE predicate on the "agTotalTests" field.
func AgTotalTestsGTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgTotalTests), v))
	})
}

// AgTotalTestsLT applies the LT predicate on the "agTotalTests" field.
func AgTotalTestsLT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgTotalTests), v))
	})
}

// AgTotalTestsLTE applies the LTE predicate on the "agTotalTests" field.
func AgTotalTestsLTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgTotalTests), v))
	})
}

// PcrPositiveTestsEQ applies the EQ predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPcrPositiveTests), v))
	})
}

// PcrPositiveTestsNEQ applies the NEQ predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsNEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPcrPositiveTests), v))
	})
}

// PcrPositiveTestsIn applies the In predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPcrPositiveTests), v...))
	})
}

// PcrPositiveTestsNotIn applies the NotIn predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsNotIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPcrPositiveTests), v...))
	})
}

// PcrPositiveTestsGT applies the GT predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsGT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPcrPositiveTests), v))
	})
}

// PcrPositiveTestsGTE applies the GTE predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsGTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPcrPositiveTests), v))
	})
}

// PcrPositiveTestsLT applies the LT predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsLT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPcrPositiveTests), v))
	})
}

// PcrPositiveTestsLTE applies the LTE predicate on the "pcrPositiveTests" field.
func PcrPositiveTestsLTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPcrPositiveTests), v))
	})
}

// PrPositiveTestsEQ applies the EQ predicate on the "prPositiveTests" field.
func PrPositiveTestsEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrPositiveTests), v))
	})
}

// PrPositiveTestsNEQ applies the NEQ predicate on the "prPositiveTests" field.
func PrPositiveTestsNEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrPositiveTests), v))
	})
}

// PrPositiveTestsIn applies the In predicate on the "prPositiveTests" field.
func PrPositiveTestsIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrPositiveTests), v...))
	})
}

// PrPositiveTestsNotIn applies the NotIn predicate on the "prPositiveTests" field.
func PrPositiveTestsNotIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrPositiveTests), v...))
	})
}

// PrPositiveTestsGT applies the GT predicate on the "prPositiveTests" field.
func PrPositiveTestsGT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrPositiveTests), v))
	})
}

// PrPositiveTestsGTE applies the GTE predicate on the "prPositiveTests" field.
func PrPositiveTestsGTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrPositiveTests), v))
	})
}

// PrPositiveTestsLT applies the LT predicate on the "prPositiveTests" field.
func PrPositiveTestsLT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrPositiveTests), v))
	})
}

// PrPositiveTestsLTE applies the LTE predicate on the "prPositiveTests" field.
func PrPositiveTestsLTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrPositiveTests), v))
	})
}

// AgPositiveTestsEQ applies the EQ predicate on the "agPositiveTests" field.
func AgPositiveTestsEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgPositiveTests), v))
	})
}

// AgPositiveTestsNEQ applies the NEQ predicate on the "agPositiveTests" field.
func AgPositiveTestsNEQ(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgPositiveTests), v))
	})
}

// AgPositiveTestsIn applies the In predicate on the "agPositiveTests" field.
func AgPositiveTestsIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAgPositiveTests), v...))
	})
}

// AgPositiveTestsNotIn applies the NotIn predicate on the "agPositiveTests" field.
func AgPositiveTestsNotIn(vs ...int) predicate.InfectedRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InfectedRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAgPositiveTests), v...))
	})
}

// AgPositiveTestsGT applies the GT predicate on the "agPositiveTests" field.
func AgPositiveTestsGT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgPositiveTests), v))
	})
}

// AgPositiveTestsGTE applies the GTE predicate on the "agPositiveTests" field.
func AgPositiveTestsGTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgPositiveTests), v))
	})
}

// AgPositiveTestsLT applies the LT predicate on the "agPositiveTests" field.
func AgPositiveTestsLT(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgPositiveTests), v))
	})
}

// AgPositiveTestsLTE applies the LTE predicate on the "agPositiveTests" field.
func AgPositiveTestsLTE(v int) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgPositiveTests), v))
	})
}

// HasPlaces applies the HasEdge predicate on the "places" edge.
func HasPlaces() predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlacesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PlacesTable, PlacesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlacesWith applies the HasEdge predicate on the "places" edge with a given conditions (other predicates).
func HasPlacesWith(preds ...predicate.Place) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlacesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PlacesTable, PlacesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InfectedRecord) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InfectedRecord) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
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
func Not(p predicate.InfectedRecord) predicate.InfectedRecord {
	return predicate.InfectedRecord(func(s *sql.Selector) {
		p(s.Not())
	})
}