// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MirrorChyan/resource-backend/internal/ent/resource"
	"github.com/MirrorChyan/resource-backend/internal/ent/storage"
	"github.com/MirrorChyan/resource-backend/internal/ent/version"
)

// VersionCreate is the builder for creating a Version entity.
type VersionCreate struct {
	config
	mutation *VersionMutation
	hooks    []Hook
}

// SetChannel sets the "channel" field.
func (vc *VersionCreate) SetChannel(v version.Channel) *VersionCreate {
	vc.mutation.SetChannel(v)
	return vc
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (vc *VersionCreate) SetNillableChannel(v *version.Channel) *VersionCreate {
	if v != nil {
		vc.SetChannel(*v)
	}
	return vc
}

// SetName sets the "name" field.
func (vc *VersionCreate) SetName(s string) *VersionCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetNumber sets the "number" field.
func (vc *VersionCreate) SetNumber(u uint64) *VersionCreate {
	vc.mutation.SetNumber(u)
	return vc
}

// SetReleaseNote sets the "release_note" field.
func (vc *VersionCreate) SetReleaseNote(s string) *VersionCreate {
	vc.mutation.SetReleaseNote(s)
	return vc
}

// SetNillableReleaseNote sets the "release_note" field if the given value is not nil.
func (vc *VersionCreate) SetNillableReleaseNote(s *string) *VersionCreate {
	if s != nil {
		vc.SetReleaseNote(*s)
	}
	return vc
}

// SetCustomData sets the "custom_data" field.
func (vc *VersionCreate) SetCustomData(s string) *VersionCreate {
	vc.mutation.SetCustomData(s)
	return vc
}

// SetNillableCustomData sets the "custom_data" field if the given value is not nil.
func (vc *VersionCreate) SetNillableCustomData(s *string) *VersionCreate {
	if s != nil {
		vc.SetCustomData(*s)
	}
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VersionCreate) SetCreatedAt(t time.Time) *VersionCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VersionCreate) SetNillableCreatedAt(t *time.Time) *VersionCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// AddStorageIDs adds the "storages" edge to the Storage entity by IDs.
func (vc *VersionCreate) AddStorageIDs(ids ...int) *VersionCreate {
	vc.mutation.AddStorageIDs(ids...)
	return vc
}

// AddStorages adds the "storages" edges to the Storage entity.
func (vc *VersionCreate) AddStorages(s ...*Storage) *VersionCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vc.AddStorageIDs(ids...)
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (vc *VersionCreate) SetResourceID(id string) *VersionCreate {
	vc.mutation.SetResourceID(id)
	return vc
}

// SetNillableResourceID sets the "resource" edge to the Resource entity by ID if the given value is not nil.
func (vc *VersionCreate) SetNillableResourceID(id *string) *VersionCreate {
	if id != nil {
		vc = vc.SetResourceID(*id)
	}
	return vc
}

// SetResource sets the "resource" edge to the Resource entity.
func (vc *VersionCreate) SetResource(r *Resource) *VersionCreate {
	return vc.SetResourceID(r.ID)
}

// Mutation returns the VersionMutation object of the builder.
func (vc *VersionCreate) Mutation() *VersionMutation {
	return vc.mutation
}

// Save creates the Version in the database.
func (vc *VersionCreate) Save(ctx context.Context) (*Version, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VersionCreate) SaveX(ctx context.Context) *Version {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VersionCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VersionCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VersionCreate) defaults() {
	if _, ok := vc.mutation.Channel(); !ok {
		v := version.DefaultChannel
		vc.mutation.SetChannel(v)
	}
	if _, ok := vc.mutation.ReleaseNote(); !ok {
		v := version.DefaultReleaseNote
		vc.mutation.SetReleaseNote(v)
	}
	if _, ok := vc.mutation.CustomData(); !ok {
		v := version.DefaultCustomData
		vc.mutation.SetCustomData(v)
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := version.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VersionCreate) check() error {
	if _, ok := vc.mutation.Channel(); !ok {
		return &ValidationError{Name: "channel", err: errors.New(`ent: missing required field "Version.channel"`)}
	}
	if v, ok := vc.mutation.Channel(); ok {
		if err := version.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`ent: validator failed for field "Version.channel": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Version.name"`)}
	}
	if v, ok := vc.mutation.Name(); ok {
		if err := version.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Version.name": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Version.number"`)}
	}
	if _, ok := vc.mutation.ReleaseNote(); !ok {
		return &ValidationError{Name: "release_note", err: errors.New(`ent: missing required field "Version.release_note"`)}
	}
	if _, ok := vc.mutation.CustomData(); !ok {
		return &ValidationError{Name: "custom_data", err: errors.New(`ent: missing required field "Version.custom_data"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Version.created_at"`)}
	}
	return nil
}

func (vc *VersionCreate) sqlSave(ctx context.Context) (*Version, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VersionCreate) createSpec() (*Version, *sqlgraph.CreateSpec) {
	var (
		_node = &Version{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(version.Table, sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.Channel(); ok {
		_spec.SetField(version.FieldChannel, field.TypeEnum, value)
		_node.Channel = value
	}
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(version.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.Number(); ok {
		_spec.SetField(version.FieldNumber, field.TypeUint64, value)
		_node.Number = value
	}
	if value, ok := vc.mutation.ReleaseNote(); ok {
		_spec.SetField(version.FieldReleaseNote, field.TypeString, value)
		_node.ReleaseNote = value
	}
	if value, ok := vc.mutation.CustomData(); ok {
		_spec.SetField(version.FieldCustomData, field.TypeString, value)
		_node.CustomData = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(version.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := vc.mutation.StoragesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   version.StoragesTable,
			Columns: []string{version.StoragesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   version.ResourceTable,
			Columns: []string{version.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.resource_versions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VersionCreateBulk is the builder for creating many Version entities in bulk.
type VersionCreateBulk struct {
	config
	err      error
	builders []*VersionCreate
}

// Save creates the Version entities in the database.
func (vcb *VersionCreateBulk) Save(ctx context.Context) ([]*Version, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Version, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VersionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VersionCreateBulk) SaveX(ctx context.Context) []*Version {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VersionCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VersionCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
