// Code generated by entc, DO NOT EDIT.

package disease

const (
	// Label holds the string label denoting the disease type in the database.
	Label = "disease"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisease holds the string denoting the disease field in the database.
	FieldDisease = "disease"

	// Table holds the table name of the disease in the database.
	Table = "diseases"
)

// Columns holds all SQL columns for disease fields.
var Columns = []string{
	FieldID,
	FieldDisease,
}

var (
	// DiseaseValidator is a validator for the "disease" field. It is called by the builders before save.
	DiseaseValidator func(string) error
)
