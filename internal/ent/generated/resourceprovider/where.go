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

package resourceprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldDescription, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldOwnerID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldContainsFold(FieldDescription, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v gidx.PrefixedID) predicate.ResourceProvider {
	return predicate.ResourceProvider(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v gidx.PrefixedID) predicate.ResourceProvider {
	vc := string(v)
	return predicate.ResourceProvider(sql.FieldContains(FieldOwnerID, vc))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v gidx.PrefixedID) predicate.ResourceProvider {
	vc := string(v)
	return predicate.ResourceProvider(sql.FieldHasPrefix(FieldOwnerID, vc))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v gidx.PrefixedID) predicate.ResourceProvider {
	vc := string(v)
	return predicate.ResourceProvider(sql.FieldHasSuffix(FieldOwnerID, vc))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v gidx.PrefixedID) predicate.ResourceProvider {
	vc := string(v)
	return predicate.ResourceProvider(sql.FieldEqualFold(FieldOwnerID, vc))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v gidx.PrefixedID) predicate.ResourceProvider {
	vc := string(v)
	return predicate.ResourceProvider(sql.FieldContainsFold(FieldOwnerID, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ResourceProvider) predicate.ResourceProvider {
	return predicate.ResourceProvider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ResourceProvider) predicate.ResourceProvider {
	return predicate.ResourceProvider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ResourceProvider) predicate.ResourceProvider {
	return predicate.ResourceProvider(func(s *sql.Selector) {
		p(s.Not())
	})
}
