// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"opencensus/core/ent/district"
	"opencensus/core/ent/place"
	"opencensus/core/ent/province"
	"opencensus/core/ent/region"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Place is the model entity for the Place schema.
type Place struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Politic holds the value of the "politic" field.
	Politic string `json:"politic,omitempty"`
	// Ubigeo holds the value of the "ubigeo" field.
	Ubigeo string `json:"ubigeo,omitempty"`
	// CovidZone holds the value of the "covidZone" field.
	CovidZone bool `json:"covidZone,omitempty"`
	// Lat holds the value of the "lat" field.
	Lat float64 `json:"lat,omitempty"`
	// Lon holds the value of the "lon" field.
	Lon float64 `json:"lon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaceQuery when eager-loading is set.
	Edges          PlaceEdges `json:"edges"`
	place_region   *int
	place_province *int
	place_district *int
}

// PlaceEdges holds the relations/edges for other nodes in the graph.
type PlaceEdges struct {
	// Oxygenrecords holds the value of the oxygenrecords edge.
	Oxygenrecords []*OxygenRecord `json:"oxygenrecords,omitempty"`
	// BedRecords holds the value of the bedRecords edge.
	BedRecords []*BedRecord `json:"bedRecords,omitempty"`
	// DeathRecords holds the value of the deathRecords edge.
	DeathRecords []*DeathRecord `json:"deathRecords,omitempty"`
	// InfectedRecords holds the value of the infectedRecords edge.
	InfectedRecords []*InfectedRecord `json:"infectedRecords,omitempty"`
	// Region holds the value of the region edge.
	Region *Region `json:"region,omitempty"`
	// Province holds the value of the province edge.
	Province *Province `json:"province,omitempty"`
	// District holds the value of the district edge.
	District *District `json:"district,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// OxygenrecordsOrErr returns the Oxygenrecords value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) OxygenrecordsOrErr() ([]*OxygenRecord, error) {
	if e.loadedTypes[0] {
		return e.Oxygenrecords, nil
	}
	return nil, &NotLoadedError{edge: "oxygenrecords"}
}

// BedRecordsOrErr returns the BedRecords value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) BedRecordsOrErr() ([]*BedRecord, error) {
	if e.loadedTypes[1] {
		return e.BedRecords, nil
	}
	return nil, &NotLoadedError{edge: "bedRecords"}
}

// DeathRecordsOrErr returns the DeathRecords value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) DeathRecordsOrErr() ([]*DeathRecord, error) {
	if e.loadedTypes[2] {
		return e.DeathRecords, nil
	}
	return nil, &NotLoadedError{edge: "deathRecords"}
}

// InfectedRecordsOrErr returns the InfectedRecords value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) InfectedRecordsOrErr() ([]*InfectedRecord, error) {
	if e.loadedTypes[3] {
		return e.InfectedRecords, nil
	}
	return nil, &NotLoadedError{edge: "infectedRecords"}
}

// RegionOrErr returns the Region value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceEdges) RegionOrErr() (*Region, error) {
	if e.loadedTypes[4] {
		if e.Region == nil {
			// The edge region was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: region.Label}
		}
		return e.Region, nil
	}
	return nil, &NotLoadedError{edge: "region"}
}

// ProvinceOrErr returns the Province value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceEdges) ProvinceOrErr() (*Province, error) {
	if e.loadedTypes[5] {
		if e.Province == nil {
			// The edge province was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: province.Label}
		}
		return e.Province, nil
	}
	return nil, &NotLoadedError{edge: "province"}
}

// DistrictOrErr returns the District value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceEdges) DistrictOrErr() (*District, error) {
	if e.loadedTypes[6] {
		if e.District == nil {
			// The edge district was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: district.Label}
		}
		return e.District, nil
	}
	return nil, &NotLoadedError{edge: "district"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Place) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case place.FieldCovidZone:
			values[i] = &sql.NullBool{}
		case place.FieldLat, place.FieldLon:
			values[i] = &sql.NullFloat64{}
		case place.FieldID:
			values[i] = &sql.NullInt64{}
		case place.FieldKind, place.FieldName, place.FieldPolitic, place.FieldUbigeo:
			values[i] = &sql.NullString{}
		case place.ForeignKeys[0]: // place_region
			values[i] = &sql.NullInt64{}
		case place.ForeignKeys[1]: // place_province
			values[i] = &sql.NullInt64{}
		case place.ForeignKeys[2]: // place_district
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Place", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Place fields.
func (pl *Place) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case place.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case place.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				pl.Kind = value.String
			}
		case place.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case place.FieldPolitic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field politic", values[i])
			} else if value.Valid {
				pl.Politic = value.String
			}
		case place.FieldUbigeo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ubigeo", values[i])
			} else if value.Valid {
				pl.Ubigeo = value.String
			}
		case place.FieldCovidZone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field covidZone", values[i])
			} else if value.Valid {
				pl.CovidZone = value.Bool
			}
		case place.FieldLat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lat", values[i])
			} else if value.Valid {
				pl.Lat = value.Float64
			}
		case place.FieldLon:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lon", values[i])
			} else if value.Valid {
				pl.Lon = value.Float64
			}
		case place.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field place_region", value)
			} else if value.Valid {
				pl.place_region = new(int)
				*pl.place_region = int(value.Int64)
			}
		case place.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field place_province", value)
			} else if value.Valid {
				pl.place_province = new(int)
				*pl.place_province = int(value.Int64)
			}
		case place.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field place_district", value)
			} else if value.Valid {
				pl.place_district = new(int)
				*pl.place_district = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOxygenrecords queries the "oxygenrecords" edge of the Place entity.
func (pl *Place) QueryOxygenrecords() *OxygenRecordQuery {
	return (&PlaceClient{config: pl.config}).QueryOxygenrecords(pl)
}

// QueryBedRecords queries the "bedRecords" edge of the Place entity.
func (pl *Place) QueryBedRecords() *BedRecordQuery {
	return (&PlaceClient{config: pl.config}).QueryBedRecords(pl)
}

// QueryDeathRecords queries the "deathRecords" edge of the Place entity.
func (pl *Place) QueryDeathRecords() *DeathRecordQuery {
	return (&PlaceClient{config: pl.config}).QueryDeathRecords(pl)
}

// QueryInfectedRecords queries the "infectedRecords" edge of the Place entity.
func (pl *Place) QueryInfectedRecords() *InfectedRecordQuery {
	return (&PlaceClient{config: pl.config}).QueryInfectedRecords(pl)
}

// QueryRegion queries the "region" edge of the Place entity.
func (pl *Place) QueryRegion() *RegionQuery {
	return (&PlaceClient{config: pl.config}).QueryRegion(pl)
}

// QueryProvince queries the "province" edge of the Place entity.
func (pl *Place) QueryProvince() *ProvinceQuery {
	return (&PlaceClient{config: pl.config}).QueryProvince(pl)
}

// QueryDistrict queries the "district" edge of the Place entity.
func (pl *Place) QueryDistrict() *DistrictQuery {
	return (&PlaceClient{config: pl.config}).QueryDistrict(pl)
}

// Update returns a builder for updating this Place.
// Note that you need to call Place.Unwrap() before calling this method if this Place
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Place) Update() *PlaceUpdateOne {
	return (&PlaceClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Place entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Place) Unwrap() *Place {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Place is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Place) String() string {
	var builder strings.Builder
	builder.WriteString("Place(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", kind=")
	builder.WriteString(pl.Kind)
	builder.WriteString(", name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", politic=")
	builder.WriteString(pl.Politic)
	builder.WriteString(", ubigeo=")
	builder.WriteString(pl.Ubigeo)
	builder.WriteString(", covidZone=")
	builder.WriteString(fmt.Sprintf("%v", pl.CovidZone))
	builder.WriteString(", lat=")
	builder.WriteString(fmt.Sprintf("%v", pl.Lat))
	builder.WriteString(", lon=")
	builder.WriteString(fmt.Sprintf("%v", pl.Lon))
	builder.WriteByte(')')
	return builder.String()
}

// Places is a parsable slice of Place.
type Places []*Place

func (pl Places) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}
