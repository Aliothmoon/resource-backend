// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MirrorChyan/resource-backend/internal/ent/predicate"
	"github.com/MirrorChyan/resource-backend/internal/ent/version"
)

// VersionDelete is the builder for deleting a Version entity.
type VersionDelete struct {
	config
	hooks    []Hook
	mutation *VersionMutation
}

// Where appends a list predicates to the VersionDelete builder.
func (vd *VersionDelete) Where(ps ...predicate.Version) *VersionDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VersionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vd.sqlExec, vd.mutation, vd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VersionDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VersionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(version.Table, sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt))
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vd.mutation.done = true
	return affected, err
}

// VersionDeleteOne is the builder for deleting a single Version entity.
type VersionDeleteOne struct {
	vd *VersionDelete
}

// Where appends a list predicates to the VersionDelete builder.
func (vdo *VersionDeleteOne) Where(ps ...predicate.Version) *VersionDeleteOne {
	vdo.vd.mutation.Where(ps...)
	return vdo
}

// Exec executes the deletion query.
func (vdo *VersionDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{version.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VersionDeleteOne) ExecX(ctx context.Context) {
	if err := vdo.Exec(ctx); err != nil {
		panic(err)
	}
}