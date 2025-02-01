// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/MirrorChyan/resource-backend/internal/ent/resource"
)

// Resource is the model entity for the Resource schema.
type Resource struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceQuery when eager-loading is set.
	Edges        ResourceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ResourceEdges holds the relations/edges for other nodes in the graph.
type ResourceEdges struct {
	// Versions holds the value of the versions edge.
	Versions []*Version `json:"versions,omitempty"`
	// LatestVersions holds the value of the latest_versions edge.
	LatestVersions []*LatestVersion `json:"latest_versions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VersionsOrErr returns the Versions value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) VersionsOrErr() ([]*Version, error) {
	if e.loadedTypes[0] {
		return e.Versions, nil
	}
	return nil, &NotLoadedError{edge: "versions"}
}

// LatestVersionsOrErr returns the LatestVersions value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) LatestVersionsOrErr() ([]*LatestVersion, error) {
	if e.loadedTypes[1] {
		return e.LatestVersions, nil
	}
	return nil, &NotLoadedError{edge: "latest_versions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resource.FieldID, resource.FieldName, resource.FieldDescription:
			values[i] = new(sql.NullString)
		case resource.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Resource fields.
func (r *Resource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resource.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case resource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case resource.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case resource.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Resource.
// This includes values selected through modifiers, order, etc.
func (r *Resource) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryVersions queries the "versions" edge of the Resource entity.
func (r *Resource) QueryVersions() *VersionQuery {
	return NewResourceClient(r.config).QueryVersions(r)
}

// QueryLatestVersions queries the "latest_versions" edge of the Resource entity.
func (r *Resource) QueryLatestVersions() *LatestVersionQuery {
	return NewResourceClient(r.config).QueryLatestVersions(r)
}

// Update returns a builder for updating this Resource.
// Note that you need to call Resource.Unwrap() before calling this method if this Resource
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Resource) Update() *ResourceUpdateOne {
	return NewResourceClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Resource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Resource) Unwrap() *Resource {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Resource is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Resource) String() string {
	var builder strings.Builder
	builder.WriteString("Resource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Resources is a parsable slice of Resource.
type Resources []*Resource
