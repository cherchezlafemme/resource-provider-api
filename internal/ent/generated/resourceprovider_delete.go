// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/predicate"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/resourceprovider"
)

// ResourceProviderDelete is the builder for deleting a ResourceProvider entity.
type ResourceProviderDelete struct {
	config
	hooks    []Hook
	mutation *ResourceProviderMutation
}

// Where appends a list predicates to the ResourceProviderDelete builder.
func (rpd *ResourceProviderDelete) Where(ps ...predicate.ResourceProvider) *ResourceProviderDelete {
	rpd.mutation.Where(ps...)
	return rpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rpd *ResourceProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rpd.sqlExec, rpd.mutation, rpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rpd *ResourceProviderDelete) ExecX(ctx context.Context) int {
	n, err := rpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rpd *ResourceProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(resourceprovider.Table, sqlgraph.NewFieldSpec(resourceprovider.FieldID, field.TypeString))
	if ps := rpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rpd.mutation.done = true
	return affected, err
}

// ResourceProviderDeleteOne is the builder for deleting a single ResourceProvider entity.
type ResourceProviderDeleteOne struct {
	rpd *ResourceProviderDelete
}

// Where appends a list predicates to the ResourceProviderDelete builder.
func (rpdo *ResourceProviderDeleteOne) Where(ps ...predicate.ResourceProvider) *ResourceProviderDeleteOne {
	rpdo.rpd.mutation.Where(ps...)
	return rpdo
}

// Exec executes the deletion query.
func (rpdo *ResourceProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := rpdo.rpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{resourceprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rpdo *ResourceProviderDeleteOne) ExecX(ctx context.Context) {
	if err := rpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
