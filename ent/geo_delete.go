// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonimaru/try-ent/ent/predicate"

	entgeo "github.com/tonimaru/try-ent/ent/geo"
)

// GeoDelete is the builder for deleting a Geo entity.
type GeoDelete struct {
	config
	hooks    []Hook
	mutation *GeoMutation
}

// Where appends a list predicates to the GeoDelete builder.
func (gd *GeoDelete) Where(ps ...predicate.Geo) *GeoDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GeoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gd.sqlExec, gd.mutation, gd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GeoDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GeoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(entgeo.Table, sqlgraph.NewFieldSpec(entgeo.FieldID, field.TypeInt))
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gd.mutation.done = true
	return affected, err
}

// GeoDeleteOne is the builder for deleting a single Geo entity.
type GeoDeleteOne struct {
	gd *GeoDelete
}

// Where appends a list predicates to the GeoDelete builder.
func (gdo *GeoDeleteOne) Where(ps ...predicate.Geo) *GeoDeleteOne {
	gdo.gd.mutation.Where(ps...)
	return gdo
}

// Exec executes the deletion query.
func (gdo *GeoDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{entgeo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GeoDeleteOne) ExecX(ctx context.Context) {
	if err := gdo.Exec(ctx); err != nil {
		panic(err)
	}
}