// Code generated by entc, DO NOT EDIT.

package place

const (
	// Label holds the string label denoting the place type in the database.
	Label = "place"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPolitic holds the string denoting the politic field in the database.
	FieldPolitic = "politic"
	// FieldUbigeo holds the string denoting the ubigeo field in the database.
	FieldUbigeo = "ubigeo"
	// FieldCovidZone holds the string denoting the covidzone field in the database.
	FieldCovidZone = "covid_zone"
	// FieldLat holds the string denoting the lat field in the database.
	FieldLat = "lat"
	// FieldLon holds the string denoting the lon field in the database.
	FieldLon = "lon"

	// EdgeOxygenrecords holds the string denoting the oxygenrecords edge name in mutations.
	EdgeOxygenrecords = "oxygenrecords"
	// EdgeBedRecords holds the string denoting the bedrecords edge name in mutations.
	EdgeBedRecords = "bedRecords"
	// EdgeDeathRecords holds the string denoting the deathrecords edge name in mutations.
	EdgeDeathRecords = "deathRecords"
	// EdgeInfectedRecords holds the string denoting the infectedrecords edge name in mutations.
	EdgeInfectedRecords = "infectedRecords"
	// EdgeRegions holds the string denoting the regions edge name in mutations.
	EdgeRegions = "regions"
	// EdgeProvinces holds the string denoting the provinces edge name in mutations.
	EdgeProvinces = "provinces"
	// EdgeDistricts holds the string denoting the districts edge name in mutations.
	EdgeDistricts = "districts"

	// Table holds the table name of the place in the database.
	Table = "places"
	// OxygenrecordsTable is the table the holds the oxygenrecords relation/edge. The primary key declared below.
	OxygenrecordsTable = "oxygen_record_places"
	// OxygenrecordsInverseTable is the table name for the OxygenRecord entity.
	// It exists in this package in order to avoid circular dependency with the "oxygenrecord" package.
	OxygenrecordsInverseTable = "oxygen_records"
	// BedRecordsTable is the table the holds the bedRecords relation/edge. The primary key declared below.
	BedRecordsTable = "bed_record_places"
	// BedRecordsInverseTable is the table name for the BedRecord entity.
	// It exists in this package in order to avoid circular dependency with the "bedrecord" package.
	BedRecordsInverseTable = "bed_records"
	// DeathRecordsTable is the table the holds the deathRecords relation/edge. The primary key declared below.
	DeathRecordsTable = "death_record_places"
	// DeathRecordsInverseTable is the table name for the DeathRecord entity.
	// It exists in this package in order to avoid circular dependency with the "deathrecord" package.
	DeathRecordsInverseTable = "death_records"
	// InfectedRecordsTable is the table the holds the infectedRecords relation/edge. The primary key declared below.
	InfectedRecordsTable = "infected_record_places"
	// InfectedRecordsInverseTable is the table name for the InfectedRecord entity.
	// It exists in this package in order to avoid circular dependency with the "infectedrecord" package.
	InfectedRecordsInverseTable = "infected_records"
	// RegionsTable is the table the holds the regions relation/edge. The primary key declared below.
	RegionsTable = "place_regions"
	// RegionsInverseTable is the table name for the Region entity.
	// It exists in this package in order to avoid circular dependency with the "region" package.
	RegionsInverseTable = "regions"
	// ProvincesTable is the table the holds the provinces relation/edge. The primary key declared below.
	ProvincesTable = "place_provinces"
	// ProvincesInverseTable is the table name for the Province entity.
	// It exists in this package in order to avoid circular dependency with the "province" package.
	ProvincesInverseTable = "provinces"
	// DistrictsTable is the table the holds the districts relation/edge. The primary key declared below.
	DistrictsTable = "place_districts"
	// DistrictsInverseTable is the table name for the District entity.
	// It exists in this package in order to avoid circular dependency with the "district" package.
	DistrictsInverseTable = "districts"
)

// Columns holds all SQL columns for place fields.
var Columns = []string{
	FieldID,
	FieldKind,
	FieldName,
	FieldPolitic,
	FieldUbigeo,
	FieldCovidZone,
	FieldLat,
	FieldLon,
}

var (
	// OxygenrecordsPrimaryKey and OxygenrecordsColumn2 are the table columns denoting the
	// primary key for the oxygenrecords relation (M2M).
	OxygenrecordsPrimaryKey = []string{"oxygen_record_id", "place_id"}
	// BedRecordsPrimaryKey and BedRecordsColumn2 are the table columns denoting the
	// primary key for the bedRecords relation (M2M).
	BedRecordsPrimaryKey = []string{"bed_record_id", "place_id"}
	// DeathRecordsPrimaryKey and DeathRecordsColumn2 are the table columns denoting the
	// primary key for the deathRecords relation (M2M).
	DeathRecordsPrimaryKey = []string{"death_record_id", "place_id"}
	// InfectedRecordsPrimaryKey and InfectedRecordsColumn2 are the table columns denoting the
	// primary key for the infectedRecords relation (M2M).
	InfectedRecordsPrimaryKey = []string{"infected_record_id", "place_id"}
	// RegionsPrimaryKey and RegionsColumn2 are the table columns denoting the
	// primary key for the regions relation (M2M).
	RegionsPrimaryKey = []string{"place_id", "region_id"}
	// ProvincesPrimaryKey and ProvincesColumn2 are the table columns denoting the
	// primary key for the provinces relation (M2M).
	ProvincesPrimaryKey = []string{"place_id", "province_id"}
	// DistrictsPrimaryKey and DistrictsColumn2 are the table columns denoting the
	// primary key for the districts relation (M2M).
	DistrictsPrimaryKey = []string{"place_id", "district_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
