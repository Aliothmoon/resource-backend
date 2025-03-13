// Code generated by ent, DO NOT EDIT.

package resource

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the resource type in the database.
	Label = "resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdateType holds the string denoting the update_type field in the database.
	FieldUpdateType = "update_type"
	// EdgeVersions holds the string denoting the versions edge name in mutations.
	EdgeVersions = "versions"
	// Table holds the table name of the resource in the database.
	Table = "resources"
	// VersionsTable is the table that holds the versions relation/edge.
	VersionsTable = "versions"
	// VersionsInverseTable is the table name for the Version entity.
	// It exists in this package in order to avoid circular dependency with the "version" package.
	VersionsInverseTable = "versions"
	// VersionsColumn is the table column denoting the versions relation/edge.
	VersionsColumn = "resource_versions"
)

// Columns holds all SQL columns for resource fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCreatedAt,
	FieldUpdateType,
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

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdateType holds the default value on creation for the "update_type" field.
	DefaultUpdateType string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Resource queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdateType orders the results by the update_type field.
func ByUpdateType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateType, opts...).ToFunc()
}

// ByVersionsCount orders the results by versions count.
func ByVersionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVersionsStep(), opts...)
	}
}

// ByVersions orders the results by versions terms.
func ByVersions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVersionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVersionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VersionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VersionsTable, VersionsColumn),
	)
}
