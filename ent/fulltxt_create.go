// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonimaru/try-ent/ent/fulltxt"
)

// FulltxtCreate is the builder for creating a Fulltxt entity.
type FulltxtCreate struct {
	config
	mutation *FulltxtMutation
	hooks    []Hook
}

// SetTxt sets the "txt" field.
func (fc *FulltxtCreate) SetTxt(s string) *FulltxtCreate {
	fc.mutation.SetTxt(s)
	return fc
}

// Mutation returns the FulltxtMutation object of the builder.
func (fc *FulltxtCreate) Mutation() *FulltxtMutation {
	return fc.mutation
}

// Save creates the Fulltxt in the database.
func (fc *FulltxtCreate) Save(ctx context.Context) (*Fulltxt, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FulltxtCreate) SaveX(ctx context.Context) *Fulltxt {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FulltxtCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FulltxtCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FulltxtCreate) check() error {
	if _, ok := fc.mutation.Txt(); !ok {
		return &ValidationError{Name: "txt", err: errors.New(`ent: missing required field "Fulltxt.txt"`)}
	}
	return nil
}

func (fc *FulltxtCreate) sqlSave(ctx context.Context) (*Fulltxt, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FulltxtCreate) createSpec() (*Fulltxt, *sqlgraph.CreateSpec) {
	var (
		_node = &Fulltxt{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(fulltxt.Table, sqlgraph.NewFieldSpec(fulltxt.FieldID, field.TypeInt))
	)
	if value, ok := fc.mutation.Txt(); ok {
		_spec.SetField(fulltxt.FieldTxt, field.TypeString, value)
		_node.Txt = value
	}
	return _node, _spec
}

// FulltxtCreateBulk is the builder for creating many Fulltxt entities in bulk.
type FulltxtCreateBulk struct {
	config
	builders []*FulltxtCreate
}

// Save creates the Fulltxt entities in the database.
func (fcb *FulltxtCreateBulk) Save(ctx context.Context) ([]*Fulltxt, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fulltxt, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FulltxtMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FulltxtCreateBulk) SaveX(ctx context.Context) []*Fulltxt {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FulltxtCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FulltxtCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
