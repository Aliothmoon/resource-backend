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
	"github.com/MirrorChyan/resource-backend/internal/ent/predicate"
	"github.com/MirrorChyan/resource-backend/internal/ent/storage"
	"github.com/MirrorChyan/resource-backend/internal/ent/version"
)

// StorageUpdate is the builder for updating Storage entities.
type StorageUpdate struct {
	config
	hooks    []Hook
	mutation *StorageMutation
}

// Where appends a list predicates to the StorageUpdate builder.
func (su *StorageUpdate) Where(ps ...predicate.Storage) *StorageUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateType sets the "update_type" field.
func (su *StorageUpdate) SetUpdateType(st storage.UpdateType) *StorageUpdate {
	su.mutation.SetUpdateType(st)
	return su
}

// SetNillableUpdateType sets the "update_type" field if the given value is not nil.
func (su *StorageUpdate) SetNillableUpdateType(st *storage.UpdateType) *StorageUpdate {
	if st != nil {
		su.SetUpdateType(*st)
	}
	return su
}

// SetOs sets the "os" field.
func (su *StorageUpdate) SetOs(s string) *StorageUpdate {
	su.mutation.SetOs(s)
	return su
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (su *StorageUpdate) SetNillableOs(s *string) *StorageUpdate {
	if s != nil {
		su.SetOs(*s)
	}
	return su
}

// ClearOs clears the value of the "os" field.
func (su *StorageUpdate) ClearOs() *StorageUpdate {
	su.mutation.ClearOs()
	return su
}

// SetArch sets the "arch" field.
func (su *StorageUpdate) SetArch(s string) *StorageUpdate {
	su.mutation.SetArch(s)
	return su
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (su *StorageUpdate) SetNillableArch(s *string) *StorageUpdate {
	if s != nil {
		su.SetArch(*s)
	}
	return su
}

// ClearArch clears the value of the "arch" field.
func (su *StorageUpdate) ClearArch() *StorageUpdate {
	su.mutation.ClearArch()
	return su
}

// SetPackagePath sets the "package_path" field.
func (su *StorageUpdate) SetPackagePath(s string) *StorageUpdate {
	su.mutation.SetPackagePath(s)
	return su
}

// SetNillablePackagePath sets the "package_path" field if the given value is not nil.
func (su *StorageUpdate) SetNillablePackagePath(s *string) *StorageUpdate {
	if s != nil {
		su.SetPackagePath(*s)
	}
	return su
}

// SetResourcePath sets the "resource_path" field.
func (su *StorageUpdate) SetResourcePath(s string) *StorageUpdate {
	su.mutation.SetResourcePath(s)
	return su
}

// SetNillableResourcePath sets the "resource_path" field if the given value is not nil.
func (su *StorageUpdate) SetNillableResourcePath(s *string) *StorageUpdate {
	if s != nil {
		su.SetResourcePath(*s)
	}
	return su
}

// ClearResourcePath clears the value of the "resource_path" field.
func (su *StorageUpdate) ClearResourcePath() *StorageUpdate {
	su.mutation.ClearResourcePath()
	return su
}

// SetFileHashes sets the "file_hashes" field.
func (su *StorageUpdate) SetFileHashes(m map[string]string) *StorageUpdate {
	su.mutation.SetFileHashes(m)
	return su
}

// ClearFileHashes clears the value of the "file_hashes" field.
func (su *StorageUpdate) ClearFileHashes() *StorageUpdate {
	su.mutation.ClearFileHashes()
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *StorageUpdate) SetCreatedAt(t time.Time) *StorageUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *StorageUpdate) SetNillableCreatedAt(t *time.Time) *StorageUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetVersionID sets the "version" edge to the Version entity by ID.
func (su *StorageUpdate) SetVersionID(id int) *StorageUpdate {
	su.mutation.SetVersionID(id)
	return su
}

// SetVersion sets the "version" edge to the Version entity.
func (su *StorageUpdate) SetVersion(v *Version) *StorageUpdate {
	return su.SetVersionID(v.ID)
}

// SetOldVersionID sets the "old_version" edge to the Version entity by ID.
func (su *StorageUpdate) SetOldVersionID(id int) *StorageUpdate {
	su.mutation.SetOldVersionID(id)
	return su
}

// SetNillableOldVersionID sets the "old_version" edge to the Version entity by ID if the given value is not nil.
func (su *StorageUpdate) SetNillableOldVersionID(id *int) *StorageUpdate {
	if id != nil {
		su = su.SetOldVersionID(*id)
	}
	return su
}

// SetOldVersion sets the "old_version" edge to the Version entity.
func (su *StorageUpdate) SetOldVersion(v *Version) *StorageUpdate {
	return su.SetOldVersionID(v.ID)
}

// Mutation returns the StorageMutation object of the builder.
func (su *StorageUpdate) Mutation() *StorageMutation {
	return su.mutation
}

// ClearVersion clears the "version" edge to the Version entity.
func (su *StorageUpdate) ClearVersion() *StorageUpdate {
	su.mutation.ClearVersion()
	return su
}

// ClearOldVersion clears the "old_version" edge to the Version entity.
func (su *StorageUpdate) ClearOldVersion() *StorageUpdate {
	su.mutation.ClearOldVersion()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StorageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StorageUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StorageUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StorageUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StorageUpdate) check() error {
	if v, ok := su.mutation.UpdateType(); ok {
		if err := storage.UpdateTypeValidator(v); err != nil {
			return &ValidationError{Name: "update_type", err: fmt.Errorf(`ent: validator failed for field "Storage.update_type": %w`, err)}
		}
	}
	if v, ok := su.mutation.PackagePath(); ok {
		if err := storage.PackagePathValidator(v); err != nil {
			return &ValidationError{Name: "package_path", err: fmt.Errorf(`ent: validator failed for field "Storage.package_path": %w`, err)}
		}
	}
	if su.mutation.VersionCleared() && len(su.mutation.VersionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Storage.version"`)
	}
	return nil
}

func (su *StorageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(storage.Table, storage.Columns, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateType(); ok {
		_spec.SetField(storage.FieldUpdateType, field.TypeEnum, value)
	}
	if value, ok := su.mutation.Os(); ok {
		_spec.SetField(storage.FieldOs, field.TypeString, value)
	}
	if su.mutation.OsCleared() {
		_spec.ClearField(storage.FieldOs, field.TypeString)
	}
	if value, ok := su.mutation.Arch(); ok {
		_spec.SetField(storage.FieldArch, field.TypeString, value)
	}
	if su.mutation.ArchCleared() {
		_spec.ClearField(storage.FieldArch, field.TypeString)
	}
	if value, ok := su.mutation.PackagePath(); ok {
		_spec.SetField(storage.FieldPackagePath, field.TypeString, value)
	}
	if value, ok := su.mutation.ResourcePath(); ok {
		_spec.SetField(storage.FieldResourcePath, field.TypeString, value)
	}
	if su.mutation.ResourcePathCleared() {
		_spec.ClearField(storage.FieldResourcePath, field.TypeString)
	}
	if value, ok := su.mutation.FileHashes(); ok {
		_spec.SetField(storage.FieldFileHashes, field.TypeJSON, value)
	}
	if su.mutation.FileHashesCleared() {
		_spec.ClearField(storage.FieldFileHashes, field.TypeJSON)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(storage.FieldCreatedAt, field.TypeTime, value)
	}
	if su.mutation.VersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   storage.VersionTable,
			Columns: []string{storage.VersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.VersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   storage.VersionTable,
			Columns: []string{storage.VersionColumn},
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
	if su.mutation.OldVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   storage.OldVersionTable,
			Columns: []string{storage.OldVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OldVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   storage.OldVersionTable,
			Columns: []string{storage.OldVersionColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StorageUpdateOne is the builder for updating a single Storage entity.
type StorageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StorageMutation
}

// SetUpdateType sets the "update_type" field.
func (suo *StorageUpdateOne) SetUpdateType(st storage.UpdateType) *StorageUpdateOne {
	suo.mutation.SetUpdateType(st)
	return suo
}

// SetNillableUpdateType sets the "update_type" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableUpdateType(st *storage.UpdateType) *StorageUpdateOne {
	if st != nil {
		suo.SetUpdateType(*st)
	}
	return suo
}

// SetOs sets the "os" field.
func (suo *StorageUpdateOne) SetOs(s string) *StorageUpdateOne {
	suo.mutation.SetOs(s)
	return suo
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableOs(s *string) *StorageUpdateOne {
	if s != nil {
		suo.SetOs(*s)
	}
	return suo
}

// ClearOs clears the value of the "os" field.
func (suo *StorageUpdateOne) ClearOs() *StorageUpdateOne {
	suo.mutation.ClearOs()
	return suo
}

// SetArch sets the "arch" field.
func (suo *StorageUpdateOne) SetArch(s string) *StorageUpdateOne {
	suo.mutation.SetArch(s)
	return suo
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableArch(s *string) *StorageUpdateOne {
	if s != nil {
		suo.SetArch(*s)
	}
	return suo
}

// ClearArch clears the value of the "arch" field.
func (suo *StorageUpdateOne) ClearArch() *StorageUpdateOne {
	suo.mutation.ClearArch()
	return suo
}

// SetPackagePath sets the "package_path" field.
func (suo *StorageUpdateOne) SetPackagePath(s string) *StorageUpdateOne {
	suo.mutation.SetPackagePath(s)
	return suo
}

// SetNillablePackagePath sets the "package_path" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillablePackagePath(s *string) *StorageUpdateOne {
	if s != nil {
		suo.SetPackagePath(*s)
	}
	return suo
}

// SetResourcePath sets the "resource_path" field.
func (suo *StorageUpdateOne) SetResourcePath(s string) *StorageUpdateOne {
	suo.mutation.SetResourcePath(s)
	return suo
}

// SetNillableResourcePath sets the "resource_path" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableResourcePath(s *string) *StorageUpdateOne {
	if s != nil {
		suo.SetResourcePath(*s)
	}
	return suo
}

// ClearResourcePath clears the value of the "resource_path" field.
func (suo *StorageUpdateOne) ClearResourcePath() *StorageUpdateOne {
	suo.mutation.ClearResourcePath()
	return suo
}

// SetFileHashes sets the "file_hashes" field.
func (suo *StorageUpdateOne) SetFileHashes(m map[string]string) *StorageUpdateOne {
	suo.mutation.SetFileHashes(m)
	return suo
}

// ClearFileHashes clears the value of the "file_hashes" field.
func (suo *StorageUpdateOne) ClearFileHashes() *StorageUpdateOne {
	suo.mutation.ClearFileHashes()
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *StorageUpdateOne) SetCreatedAt(t time.Time) *StorageUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableCreatedAt(t *time.Time) *StorageUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetVersionID sets the "version" edge to the Version entity by ID.
func (suo *StorageUpdateOne) SetVersionID(id int) *StorageUpdateOne {
	suo.mutation.SetVersionID(id)
	return suo
}

// SetVersion sets the "version" edge to the Version entity.
func (suo *StorageUpdateOne) SetVersion(v *Version) *StorageUpdateOne {
	return suo.SetVersionID(v.ID)
}

// SetOldVersionID sets the "old_version" edge to the Version entity by ID.
func (suo *StorageUpdateOne) SetOldVersionID(id int) *StorageUpdateOne {
	suo.mutation.SetOldVersionID(id)
	return suo
}

// SetNillableOldVersionID sets the "old_version" edge to the Version entity by ID if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableOldVersionID(id *int) *StorageUpdateOne {
	if id != nil {
		suo = suo.SetOldVersionID(*id)
	}
	return suo
}

// SetOldVersion sets the "old_version" edge to the Version entity.
func (suo *StorageUpdateOne) SetOldVersion(v *Version) *StorageUpdateOne {
	return suo.SetOldVersionID(v.ID)
}

// Mutation returns the StorageMutation object of the builder.
func (suo *StorageUpdateOne) Mutation() *StorageMutation {
	return suo.mutation
}

// ClearVersion clears the "version" edge to the Version entity.
func (suo *StorageUpdateOne) ClearVersion() *StorageUpdateOne {
	suo.mutation.ClearVersion()
	return suo
}

// ClearOldVersion clears the "old_version" edge to the Version entity.
func (suo *StorageUpdateOne) ClearOldVersion() *StorageUpdateOne {
	suo.mutation.ClearOldVersion()
	return suo
}

// Where appends a list predicates to the StorageUpdate builder.
func (suo *StorageUpdateOne) Where(ps ...predicate.Storage) *StorageUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StorageUpdateOne) Select(field string, fields ...string) *StorageUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Storage entity.
func (suo *StorageUpdateOne) Save(ctx context.Context) (*Storage, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StorageUpdateOne) SaveX(ctx context.Context) *Storage {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StorageUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StorageUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StorageUpdateOne) check() error {
	if v, ok := suo.mutation.UpdateType(); ok {
		if err := storage.UpdateTypeValidator(v); err != nil {
			return &ValidationError{Name: "update_type", err: fmt.Errorf(`ent: validator failed for field "Storage.update_type": %w`, err)}
		}
	}
	if v, ok := suo.mutation.PackagePath(); ok {
		if err := storage.PackagePathValidator(v); err != nil {
			return &ValidationError{Name: "package_path", err: fmt.Errorf(`ent: validator failed for field "Storage.package_path": %w`, err)}
		}
	}
	if suo.mutation.VersionCleared() && len(suo.mutation.VersionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Storage.version"`)
	}
	return nil
}

func (suo *StorageUpdateOne) sqlSave(ctx context.Context) (_node *Storage, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(storage.Table, storage.Columns, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Storage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storage.FieldID)
		for _, f := range fields {
			if !storage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateType(); ok {
		_spec.SetField(storage.FieldUpdateType, field.TypeEnum, value)
	}
	if value, ok := suo.mutation.Os(); ok {
		_spec.SetField(storage.FieldOs, field.TypeString, value)
	}
	if suo.mutation.OsCleared() {
		_spec.ClearField(storage.FieldOs, field.TypeString)
	}
	if value, ok := suo.mutation.Arch(); ok {
		_spec.SetField(storage.FieldArch, field.TypeString, value)
	}
	if suo.mutation.ArchCleared() {
		_spec.ClearField(storage.FieldArch, field.TypeString)
	}
	if value, ok := suo.mutation.PackagePath(); ok {
		_spec.SetField(storage.FieldPackagePath, field.TypeString, value)
	}
	if value, ok := suo.mutation.ResourcePath(); ok {
		_spec.SetField(storage.FieldResourcePath, field.TypeString, value)
	}
	if suo.mutation.ResourcePathCleared() {
		_spec.ClearField(storage.FieldResourcePath, field.TypeString)
	}
	if value, ok := suo.mutation.FileHashes(); ok {
		_spec.SetField(storage.FieldFileHashes, field.TypeJSON, value)
	}
	if suo.mutation.FileHashesCleared() {
		_spec.ClearField(storage.FieldFileHashes, field.TypeJSON)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(storage.FieldCreatedAt, field.TypeTime, value)
	}
	if suo.mutation.VersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   storage.VersionTable,
			Columns: []string{storage.VersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.VersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   storage.VersionTable,
			Columns: []string{storage.VersionColumn},
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
	if suo.mutation.OldVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   storage.OldVersionTable,
			Columns: []string{storage.OldVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OldVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   storage.OldVersionTable,
			Columns: []string{storage.OldVersionColumn},
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
	_node = &Storage{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
