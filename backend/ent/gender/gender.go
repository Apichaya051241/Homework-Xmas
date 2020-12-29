// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"

	// Table holds the table name of the gender in the database.
	Table = "genders"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGender,
}

var (
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(string) error
)