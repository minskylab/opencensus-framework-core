// Code generated by entc, DO NOT EDIT.

package organization

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldUbigeo holds the string denoting the ubigeo field in the database.
	FieldUbigeo = "ubigeo"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldCovidZone holds the string denoting the covidzone field in the database.
	FieldCovidZone = "covid_zone"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"

	// EdgeRegion holds the string denoting the region edge name in mutations.
	EdgeRegion = "region"
	// EdgeProvince holds the string denoting the province edge name in mutations.
	EdgeProvince = "province"
	// EdgeDistrict holds the string denoting the district edge name in mutations.
	EdgeDistrict = "district"
	// EdgeOxygenRecords holds the string denoting the oxygenrecords edge name in mutations.
	EdgeOxygenRecords = "oxygenRecords"
	// EdgeBedRecords holds the string denoting the bedrecords edge name in mutations.
	EdgeBedRecords = "bedRecords"

	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// RegionTable is the table the holds the region relation/edge. The primary key declared below.
	RegionTable = "organization_region"
	// RegionInverseTable is the table name for the Region entity.
	// It exists in this package in order to avoid circular dependency with the "region" package.
	RegionInverseTable = "regions"
	// ProvinceTable is the table the holds the province relation/edge. The primary key declared below.
	ProvinceTable = "organization_province"
	// ProvinceInverseTable is the table name for the Province entity.
	// It exists in this package in order to avoid circular dependency with the "province" package.
	ProvinceInverseTable = "provinces"
	// DistrictTable is the table the holds the district relation/edge. The primary key declared below.
	DistrictTable = "organization_district"
	// DistrictInverseTable is the table name for the District entity.
	// It exists in this package in order to avoid circular dependency with the "district" package.
	DistrictInverseTable = "districts"
	// OxygenRecordsTable is the table the holds the oxygenRecords relation/edge. The primary key declared below.
	OxygenRecordsTable = "organization_oxygenRecords"
	// OxygenRecordsInverseTable is the table name for the OxygenRecord entity.
	// It exists in this package in order to avoid circular dependency with the "oxygenrecord" package.
	OxygenRecordsInverseTable = "oxygen_records"
	// BedRecordsTable is the table the holds the bedRecords relation/edge. The primary key declared below.
	BedRecordsTable = "organization_bedRecords"
	// BedRecordsInverseTable is the table name for the BedRecord entity.
	// It exists in this package in order to avoid circular dependency with the "bedrecord" package.
	BedRecordsInverseTable = "bed_records"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCode,
	FieldUbigeo,
	FieldKind,
	FieldCovidZone,
	FieldCategory,
}

var (
	// RegionPrimaryKey and RegionColumn2 are the table columns denoting the
	// primary key for the region relation (M2M).
	RegionPrimaryKey = []string{"organization_id", "region_id"}
	// ProvincePrimaryKey and ProvinceColumn2 are the table columns denoting the
	// primary key for the province relation (M2M).
	ProvincePrimaryKey = []string{"organization_id", "province_id"}
	// DistrictPrimaryKey and DistrictColumn2 are the table columns denoting the
	// primary key for the district relation (M2M).
	DistrictPrimaryKey = []string{"organization_id", "district_id"}
	// OxygenRecordsPrimaryKey and OxygenRecordsColumn2 are the table columns denoting the
	// primary key for the oxygenRecords relation (M2M).
	OxygenRecordsPrimaryKey = []string{"organization_id", "oxygen_record_id"}
	// BedRecordsPrimaryKey and BedRecordsColumn2 are the table columns denoting the
	// primary key for the bedRecords relation (M2M).
	BedRecordsPrimaryKey = []string{"organization_id", "bed_record_id"}
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
