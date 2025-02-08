// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MirrorChyan/resource-backend/internal/ent/latestversion"
	"github.com/MirrorChyan/resource-backend/internal/ent/predicate"
	"github.com/MirrorChyan/resource-backend/internal/ent/resource"
	"github.com/MirrorChyan/resource-backend/internal/ent/version"
)

// ResourceUpdate is the builder for updating Resource entities.
type ResourceUpdate struct {
	config
	hooks    []Hook
	mutation *ResourceMutation
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ru *ResourceUpdate) Where(ps ...predicate.Resource) *ResourceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *ResourceUpdate) SetName(s string) *ResourceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableName(s *string) *ResourceUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *ResourceUpdate) SetDescription(s string) *ResourceUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableDescription(s *string) *ResourceUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ResourceUpdate) SetCreatedAt(t time.Time) *ResourceUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableCreatedAt(t *time.Time) *ResourceUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// AddVersionIDs adds the "versions" edge to the Version entity by IDs.
func (ru *ResourceUpdate) AddVersionIDs(ids ...int) *ResourceUpdate {
	ru.mutation.AddVersionIDs(ids...)
	return ru
}

// AddVersions adds the "versions" edges to the Version entity.
func (ru *ResourceUpdate) AddVersions(v ...*Version) *ResourceUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ru.AddVersionIDs(ids...)
}

// AddLatestVersionIDs adds the "latest_versions" edge to the LatestVersion entity by IDs.
func (ru *ResourceUpdate) AddLatestVersionIDs(ids ...int) *ResourceUpdate {
	ru.mutation.AddLatestVersionIDs(ids...)
	return ru
}

// AddLatestVersions adds the "latest_versions" edges to the LatestVersion entity.
func (ru *ResourceUpdate) AddLatestVersions(l ...*LatestVersion) *ResourceUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.AddLatestVersionIDs(ids...)
}

// Mutation returns the ResourceMutation object of the builder.
func (ru *ResourceUpdate) Mutation() *ResourceMutation {
	return ru.mutation
}

// ClearVersions clears all "versions" edges to the Version entity.
func (ru *ResourceUpdate) ClearVersions() *ResourceUpdate {
	ru.mutation.ClearVersions()
	return ru
}

// RemoveVersionIDs removes the "versions" edge to Version entities by IDs.
func (ru *ResourceUpdate) RemoveVersionIDs(ids ...int) *ResourceUpdate {
	ru.mutation.RemoveVersionIDs(ids...)
	return ru
}

// RemoveVersions removes "versions" edges to Version entities.
func (ru *ResourceUpdate) RemoveVersions(v ...*Version) *ResourceUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ru.RemoveVersionIDs(ids...)
}

// ClearLatestVersions clears all "latest_versions" edges to the LatestVersion entity.
func (ru *ResourceUpdate) ClearLatestVersions() *ResourceUpdate {
	ru.mutation.ClearLatestVersions()
	return ru
}

// RemoveLatestVersionIDs removes the "latest_versions" edge to LatestVersion entities by IDs.
func (ru *ResourceUpdate) RemoveLatestVersionIDs(ids ...int) *ResourceUpdate {
	ru.mutation.RemoveLatestVersionIDs(ids...)
	return ru
}

// RemoveLatestVersions removes "latest_versions" edges to LatestVersion entities.
func (ru *ResourceUpdate) RemoveLatestVersions(l ...*LatestVersion) *ResourceUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.RemoveLatestVersionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ResourceUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := resource.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Resource.name": %w`, err)}
		}
	}
	return nil
}

func (ru *ResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(resource.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(resource.FieldDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(resource.FieldCreatedAt, field.TypeTime, value)
	}
	if ru.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.VersionsTable,
			Columns: []string{resource.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedVersionsIDs(); len(nodes) > 0 && !ru.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.VersionsTable,
			Columns: []string{resource.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.VersionsTable,
			Columns: []string{resource.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.LatestVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.LatestVersionsTable,
			Columns: []string{resource.LatestVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedLatestVersionsIDs(); len(nodes) > 0 && !ru.mutation.LatestVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.LatestVersionsTable,
			Columns: []string{resource.LatestVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LatestVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.LatestVersionsTable,
			Columns: []string{resource.LatestVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ResourceUpdateOne is the builder for updating a single Resource entity.
type ResourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResourceMutation
}

// SetName sets the "name" field.
func (ruo *ResourceUpdateOne) SetName(s string) *ResourceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableName(s *string) *ResourceUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *ResourceUpdateOne) SetDescription(s string) *ResourceUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableDescription(s *string) *ResourceUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ResourceUpdateOne) SetCreatedAt(t time.Time) *ResourceUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableCreatedAt(t *time.Time) *ResourceUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// AddVersionIDs adds the "versions" edge to the Version entity by IDs.
func (ruo *ResourceUpdateOne) AddVersionIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.AddVersionIDs(ids...)
	return ruo
}

// AddVersions adds the "versions" edges to the Version entity.
func (ruo *ResourceUpdateOne) AddVersions(v ...*Version) *ResourceUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ruo.AddVersionIDs(ids...)
}

// AddLatestVersionIDs adds the "latest_versions" edge to the LatestVersion entity by IDs.
func (ruo *ResourceUpdateOne) AddLatestVersionIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.AddLatestVersionIDs(ids...)
	return ruo
}

// AddLatestVersions adds the "latest_versions" edges to the LatestVersion entity.
func (ruo *ResourceUpdateOne) AddLatestVersions(l ...*LatestVersion) *ResourceUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.AddLatestVersionIDs(ids...)
}

// Mutation returns the ResourceMutation object of the builder.
func (ruo *ResourceUpdateOne) Mutation() *ResourceMutation {
	return ruo.mutation
}

// ClearVersions clears all "versions" edges to the Version entity.
func (ruo *ResourceUpdateOne) ClearVersions() *ResourceUpdateOne {
	ruo.mutation.ClearVersions()
	return ruo
}

// RemoveVersionIDs removes the "versions" edge to Version entities by IDs.
func (ruo *ResourceUpdateOne) RemoveVersionIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.RemoveVersionIDs(ids...)
	return ruo
}

// RemoveVersions removes "versions" edges to Version entities.
func (ruo *ResourceUpdateOne) RemoveVersions(v ...*Version) *ResourceUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ruo.RemoveVersionIDs(ids...)
}

// ClearLatestVersions clears all "latest_versions" edges to the LatestVersion entity.
func (ruo *ResourceUpdateOne) ClearLatestVersions() *ResourceUpdateOne {
	ruo.mutation.ClearLatestVersions()
	return ruo
}

// RemoveLatestVersionIDs removes the "latest_versions" edge to LatestVersion entities by IDs.
func (ruo *ResourceUpdateOne) RemoveLatestVersionIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.RemoveLatestVersionIDs(ids...)
	return ruo
}

// RemoveLatestVersions removes "latest_versions" edges to LatestVersion entities.
func (ruo *ResourceUpdateOne) RemoveLatestVersions(l ...*LatestVersion) *ResourceUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.RemoveLatestVersionIDs(ids...)
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ruo *ResourceUpdateOne) Where(ps ...predicate.Resource) *ResourceUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResourceUpdateOne) Select(field string, fields ...string) *ResourceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Resource entity.
func (ruo *ResourceUpdateOne) Save(ctx context.Context) (*Resource, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourceUpdateOne) SaveX(ctx context.Context) *Resource {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ResourceUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := resource.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Resource.name": %w`, err)}
		}
	}
	return nil
}

func (ruo *ResourceUpdateOne) sqlSave(ctx context.Context) (_node *Resource, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Resource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resource.FieldID)
		for _, f := range fields {
			if !resource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != resource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(resource.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(resource.FieldDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(resource.FieldCreatedAt, field.TypeTime, value)
	}
	if ruo.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.VersionsTable,
			Columns: []string{resource.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedVersionsIDs(); len(nodes) > 0 && !ruo.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.VersionsTable,
			Columns: []string{resource.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.VersionsTable,
			Columns: []string{resource.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.LatestVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.LatestVersionsTable,
			Columns: []string{resource.LatestVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedLatestVersionsIDs(); len(nodes) > 0 && !ruo.mutation.LatestVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.LatestVersionsTable,
			Columns: []string{resource.LatestVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LatestVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.LatestVersionsTable,
			Columns: []string{resource.LatestVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Resource{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
