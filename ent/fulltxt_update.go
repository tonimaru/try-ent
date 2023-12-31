// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonimaru/try-ent/ent/fulltxt"
	"github.com/tonimaru/try-ent/ent/predicate"
)

// FulltxtUpdate is the builder for updating Fulltxt entities.
type FulltxtUpdate struct {
	config
	hooks     []Hook
	mutation  *FulltxtMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FulltxtUpdate builder.
func (fu *FulltxtUpdate) Where(ps ...predicate.Fulltxt) *FulltxtUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetTxt sets the "txt" field.
func (fu *FulltxtUpdate) SetTxt(s string) *FulltxtUpdate {
	fu.mutation.SetTxt(s)
	return fu
}

// Mutation returns the FulltxtMutation object of the builder.
func (fu *FulltxtUpdate) Mutation() *FulltxtMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FulltxtUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FulltxtUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FulltxtUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FulltxtUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fu *FulltxtUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FulltxtUpdate {
	fu.modifiers = append(fu.modifiers, modifiers...)
	return fu
}

func (fu *FulltxtUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(fulltxt.Table, fulltxt.Columns, sqlgraph.NewFieldSpec(fulltxt.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Txt(); ok {
		_spec.SetField(fulltxt.FieldTxt, field.TypeString, value)
	}
	_spec.AddModifiers(fu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fulltxt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FulltxtUpdateOne is the builder for updating a single Fulltxt entity.
type FulltxtUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FulltxtMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTxt sets the "txt" field.
func (fuo *FulltxtUpdateOne) SetTxt(s string) *FulltxtUpdateOne {
	fuo.mutation.SetTxt(s)
	return fuo
}

// Mutation returns the FulltxtMutation object of the builder.
func (fuo *FulltxtUpdateOne) Mutation() *FulltxtMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FulltxtUpdate builder.
func (fuo *FulltxtUpdateOne) Where(ps ...predicate.Fulltxt) *FulltxtUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FulltxtUpdateOne) Select(field string, fields ...string) *FulltxtUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Fulltxt entity.
func (fuo *FulltxtUpdateOne) Save(ctx context.Context) (*Fulltxt, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FulltxtUpdateOne) SaveX(ctx context.Context) *Fulltxt {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FulltxtUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FulltxtUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fuo *FulltxtUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FulltxtUpdateOne {
	fuo.modifiers = append(fuo.modifiers, modifiers...)
	return fuo
}

func (fuo *FulltxtUpdateOne) sqlSave(ctx context.Context) (_node *Fulltxt, err error) {
	_spec := sqlgraph.NewUpdateSpec(fulltxt.Table, fulltxt.Columns, sqlgraph.NewFieldSpec(fulltxt.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Fulltxt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fulltxt.FieldID)
		for _, f := range fields {
			if !fulltxt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fulltxt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Txt(); ok {
		_spec.SetField(fulltxt.FieldTxt, field.TypeString, value)
	}
	_spec.AddModifiers(fuo.modifiers...)
	_node = &Fulltxt{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fulltxt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
