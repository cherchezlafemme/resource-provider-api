// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/resourceprovider"
	"go.infratographer.com/x/gidx"
)

// ResourceProviderCreate is the builder for creating a ResourceProvider entity.
type ResourceProviderCreate struct {
	config
	mutation *ResourceProviderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rpc *ResourceProviderCreate) SetCreatedAt(t time.Time) *ResourceProviderCreate {
	rpc.mutation.SetCreatedAt(t)
	return rpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rpc *ResourceProviderCreate) SetNillableCreatedAt(t *time.Time) *ResourceProviderCreate {
	if t != nil {
		rpc.SetCreatedAt(*t)
	}
	return rpc
}

// SetUpdatedAt sets the "updated_at" field.
func (rpc *ResourceProviderCreate) SetUpdatedAt(t time.Time) *ResourceProviderCreate {
	rpc.mutation.SetUpdatedAt(t)
	return rpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rpc *ResourceProviderCreate) SetNillableUpdatedAt(t *time.Time) *ResourceProviderCreate {
	if t != nil {
		rpc.SetUpdatedAt(*t)
	}
	return rpc
}

// SetName sets the "name" field.
func (rpc *ResourceProviderCreate) SetName(s string) *ResourceProviderCreate {
	rpc.mutation.SetName(s)
	return rpc
}

// SetDescription sets the "description" field.
func (rpc *ResourceProviderCreate) SetDescription(s string) *ResourceProviderCreate {
	rpc.mutation.SetDescription(s)
	return rpc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rpc *ResourceProviderCreate) SetNillableDescription(s *string) *ResourceProviderCreate {
	if s != nil {
		rpc.SetDescription(*s)
	}
	return rpc
}

// SetOrganizationalUnitID sets the "organizational_unit_id" field.
func (rpc *ResourceProviderCreate) SetOrganizationalUnitID(gi gidx.PrefixedID) *ResourceProviderCreate {
	rpc.mutation.SetOrganizationalUnitID(gi)
	return rpc
}

// SetID sets the "id" field.
func (rpc *ResourceProviderCreate) SetID(gi gidx.PrefixedID) *ResourceProviderCreate {
	rpc.mutation.SetID(gi)
	return rpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rpc *ResourceProviderCreate) SetNillableID(gi *gidx.PrefixedID) *ResourceProviderCreate {
	if gi != nil {
		rpc.SetID(*gi)
	}
	return rpc
}

// Mutation returns the ResourceProviderMutation object of the builder.
func (rpc *ResourceProviderCreate) Mutation() *ResourceProviderMutation {
	return rpc.mutation
}

// Save creates the ResourceProvider in the database.
func (rpc *ResourceProviderCreate) Save(ctx context.Context) (*ResourceProvider, error) {
	rpc.defaults()
	return withHooks(ctx, rpc.sqlSave, rpc.mutation, rpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rpc *ResourceProviderCreate) SaveX(ctx context.Context) *ResourceProvider {
	v, err := rpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpc *ResourceProviderCreate) Exec(ctx context.Context) error {
	_, err := rpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpc *ResourceProviderCreate) ExecX(ctx context.Context) {
	if err := rpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpc *ResourceProviderCreate) defaults() {
	if _, ok := rpc.mutation.CreatedAt(); !ok {
		v := resourceprovider.DefaultCreatedAt()
		rpc.mutation.SetCreatedAt(v)
	}
	if _, ok := rpc.mutation.UpdatedAt(); !ok {
		v := resourceprovider.DefaultUpdatedAt()
		rpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rpc.mutation.ID(); !ok {
		v := resourceprovider.DefaultID()
		rpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpc *ResourceProviderCreate) check() error {
	if _, ok := rpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "ResourceProvider.created_at"`)}
	}
	if _, ok := rpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "ResourceProvider.updated_at"`)}
	}
	if _, ok := rpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "ResourceProvider.name"`)}
	}
	if v, ok := rpc.mutation.Name(); ok {
		if err := resourceprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "ResourceProvider.name": %w`, err)}
		}
	}
	if _, ok := rpc.mutation.OrganizationalUnitID(); !ok {
		return &ValidationError{Name: "organizational_unit_id", err: errors.New(`generated: missing required field "ResourceProvider.organizational_unit_id"`)}
	}
	return nil
}

func (rpc *ResourceProviderCreate) sqlSave(ctx context.Context) (*ResourceProvider, error) {
	if err := rpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rpc.mutation.id = &_node.ID
	rpc.mutation.done = true
	return _node, nil
}

func (rpc *ResourceProviderCreate) createSpec() (*ResourceProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &ResourceProvider{config: rpc.config}
		_spec = sqlgraph.NewCreateSpec(resourceprovider.Table, sqlgraph.NewFieldSpec(resourceprovider.FieldID, field.TypeString))
	)
	if id, ok := rpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rpc.mutation.CreatedAt(); ok {
		_spec.SetField(resourceprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rpc.mutation.UpdatedAt(); ok {
		_spec.SetField(resourceprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rpc.mutation.Name(); ok {
		_spec.SetField(resourceprovider.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rpc.mutation.Description(); ok {
		_spec.SetField(resourceprovider.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rpc.mutation.OrganizationalUnitID(); ok {
		_spec.SetField(resourceprovider.FieldOrganizationalUnitID, field.TypeString, value)
		_node.OrganizationalUnitID = value
	}
	return _node, _spec
}

// ResourceProviderCreateBulk is the builder for creating many ResourceProvider entities in bulk.
type ResourceProviderCreateBulk struct {
	config
	builders []*ResourceProviderCreate
}

// Save creates the ResourceProvider entities in the database.
func (rpcb *ResourceProviderCreateBulk) Save(ctx context.Context) ([]*ResourceProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rpcb.builders))
	nodes := make([]*ResourceProvider, len(rpcb.builders))
	mutators := make([]Mutator, len(rpcb.builders))
	for i := range rpcb.builders {
		func(i int, root context.Context) {
			builder := rpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, rpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, rpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rpcb *ResourceProviderCreateBulk) SaveX(ctx context.Context) []*ResourceProvider {
	v, err := rpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpcb *ResourceProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := rpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpcb *ResourceProviderCreateBulk) ExecX(ctx context.Context) {
	if err := rpcb.Exec(ctx); err != nil {
		panic(err)
	}
}