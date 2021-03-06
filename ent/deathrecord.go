// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"opencensus/core/ent/deathrecord"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// DeathRecord is the model entity for the DeathRecord schema.
type DeathRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ReportedDate holds the value of the "reportedDate" field.
	ReportedDate time.Time `json:"reportedDate,omitempty"`
	// CollectedDate holds the value of the "collectedDate" field.
	CollectedDate time.Time `json:"collectedDate,omitempty"`
	// SinadefRegisters holds the value of the "sinadefRegisters" field.
	SinadefRegisters int `json:"sinadefRegisters,omitempty"`
	// MinsaRegisters holds the value of the "minsaRegisters" field.
	MinsaRegisters int `json:"minsaRegisters,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeathRecordQuery when eager-loading is set.
	Edges DeathRecordEdges `json:"edges"`
}

// DeathRecordEdges holds the relations/edges for other nodes in the graph.
type DeathRecordEdges struct {
	// Places holds the value of the places edge.
	Places []*Place `json:"places,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlacesOrErr returns the Places value or an error if the edge
// was not loaded in eager-loading.
func (e DeathRecordEdges) PlacesOrErr() ([]*Place, error) {
	if e.loadedTypes[0] {
		return e.Places, nil
	}
	return nil, &NotLoadedError{edge: "places"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeathRecord) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deathrecord.FieldID, deathrecord.FieldSinadefRegisters, deathrecord.FieldMinsaRegisters:
			values[i] = &sql.NullInt64{}
		case deathrecord.FieldReportedDate, deathrecord.FieldCollectedDate:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeathRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeathRecord fields.
func (dr *DeathRecord) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deathrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dr.ID = int(value.Int64)
		case deathrecord.FieldReportedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field reportedDate", values[i])
			} else if value.Valid {
				dr.ReportedDate = value.Time
			}
		case deathrecord.FieldCollectedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field collectedDate", values[i])
			} else if value.Valid {
				dr.CollectedDate = value.Time
			}
		case deathrecord.FieldSinadefRegisters:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sinadefRegisters", values[i])
			} else if value.Valid {
				dr.SinadefRegisters = int(value.Int64)
			}
		case deathrecord.FieldMinsaRegisters:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field minsaRegisters", values[i])
			} else if value.Valid {
				dr.MinsaRegisters = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPlaces queries the "places" edge of the DeathRecord entity.
func (dr *DeathRecord) QueryPlaces() *PlaceQuery {
	return (&DeathRecordClient{config: dr.config}).QueryPlaces(dr)
}

// Update returns a builder for updating this DeathRecord.
// Note that you need to call DeathRecord.Unwrap() before calling this method if this DeathRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DeathRecord) Update() *DeathRecordUpdateOne {
	return (&DeathRecordClient{config: dr.config}).UpdateOne(dr)
}

// Unwrap unwraps the DeathRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dr *DeathRecord) Unwrap() *DeathRecord {
	tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeathRecord is not a transactional entity")
	}
	dr.config.driver = tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DeathRecord) String() string {
	var builder strings.Builder
	builder.WriteString("DeathRecord(")
	builder.WriteString(fmt.Sprintf("id=%v", dr.ID))
	builder.WriteString(", reportedDate=")
	builder.WriteString(dr.ReportedDate.Format(time.ANSIC))
	builder.WriteString(", collectedDate=")
	builder.WriteString(dr.CollectedDate.Format(time.ANSIC))
	builder.WriteString(", sinadefRegisters=")
	builder.WriteString(fmt.Sprintf("%v", dr.SinadefRegisters))
	builder.WriteString(", minsaRegisters=")
	builder.WriteString(fmt.Sprintf("%v", dr.MinsaRegisters))
	builder.WriteByte(')')
	return builder.String()
}

// DeathRecords is a parsable slice of DeathRecord.
type DeathRecords []*DeathRecord

func (dr DeathRecords) config(cfg config) {
	for _i := range dr {
		dr[_i].config = cfg
	}
}