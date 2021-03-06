// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"opencensus/core/ent/oxygenrecord"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// OxygenRecord is the model entity for the OxygenRecord schema.
type OxygenRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TotalCylinders holds the value of the "totalCylinders" field.
	TotalCylinders int `json:"totalCylinders,omitempty"`
	// TotalOwnCylinders holds the value of the "totalOwnCylinders" field.
	TotalOwnCylinders int `json:"totalOwnCylinders,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OxygenRecordQuery when eager-loading is set.
	Edges OxygenRecordEdges `json:"edges"`
}

// OxygenRecordEdges holds the relations/edges for other nodes in the graph.
type OxygenRecordEdges struct {
	// Organization holds the value of the organization edge.
	Organization []*Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading.
func (e OxygenRecordEdges) OrganizationOrErr() ([]*Organization, error) {
	if e.loadedTypes[0] {
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OxygenRecord) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case oxygenrecord.FieldID, oxygenrecord.FieldTotalCylinders, oxygenrecord.FieldTotalOwnCylinders:
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type OxygenRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OxygenRecord fields.
func (or *OxygenRecord) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oxygenrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			or.ID = int(value.Int64)
		case oxygenrecord.FieldTotalCylinders:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field totalCylinders", values[i])
			} else if value.Valid {
				or.TotalCylinders = int(value.Int64)
			}
		case oxygenrecord.FieldTotalOwnCylinders:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field totalOwnCylinders", values[i])
			} else if value.Valid {
				or.TotalOwnCylinders = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrganization queries the "organization" edge of the OxygenRecord entity.
func (or *OxygenRecord) QueryOrganization() *OrganizationQuery {
	return (&OxygenRecordClient{config: or.config}).QueryOrganization(or)
}

// Update returns a builder for updating this OxygenRecord.
// Note that you need to call OxygenRecord.Unwrap() before calling this method if this OxygenRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OxygenRecord) Update() *OxygenRecordUpdateOne {
	return (&OxygenRecordClient{config: or.config}).UpdateOne(or)
}

// Unwrap unwraps the OxygenRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OxygenRecord) Unwrap() *OxygenRecord {
	tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OxygenRecord is not a transactional entity")
	}
	or.config.driver = tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OxygenRecord) String() string {
	var builder strings.Builder
	builder.WriteString("OxygenRecord(")
	builder.WriteString(fmt.Sprintf("id=%v", or.ID))
	builder.WriteString(", totalCylinders=")
	builder.WriteString(fmt.Sprintf("%v", or.TotalCylinders))
	builder.WriteString(", totalOwnCylinders=")
	builder.WriteString(fmt.Sprintf("%v", or.TotalOwnCylinders))
	builder.WriteByte(')')
	return builder.String()
}

// OxygenRecords is a parsable slice of OxygenRecord.
type OxygenRecords []*OxygenRecord

func (or OxygenRecords) config(cfg config) {
	for _i := range or {
		or[_i].config = cfg
	}
}
