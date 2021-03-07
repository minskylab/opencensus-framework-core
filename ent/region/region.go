// Code generated by entc, DO NOT EDIT.

package region

const (
	// Label holds the string label denoting the region type in the database.
	Label = "region"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgePlaces holds the string denoting the places edge name in mutations.
	EdgePlaces = "places"

	// Table holds the table name of the region in the database.
	Table = "regions"
	// PlacesTable is the table the holds the places relation/edge.
	PlacesTable = "places"
	// PlacesInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlacesInverseTable = "places"
	// PlacesColumn is the table column denoting the places relation/edge.
	PlacesColumn = "place_region"
)

// Columns holds all SQL columns for region fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
