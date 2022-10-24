// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package billproduct

import (
	"entgo.io/contrib/entgql/internal/todopulid/ent/predicate"
	"entgo.io/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Sku applies equality check predicate on the "sku" field. It's identical to SkuEQ.
func Sku(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSku), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BillProduct {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BillProduct {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SkuEQ applies the EQ predicate on the "sku" field.
func SkuEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSku), v))
	})
}

// SkuNEQ applies the NEQ predicate on the "sku" field.
func SkuNEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSku), v))
	})
}

// SkuIn applies the In predicate on the "sku" field.
func SkuIn(vs ...string) predicate.BillProduct {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSku), v...))
	})
}

// SkuNotIn applies the NotIn predicate on the "sku" field.
func SkuNotIn(vs ...string) predicate.BillProduct {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSku), v...))
	})
}

// SkuGT applies the GT predicate on the "sku" field.
func SkuGT(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSku), v))
	})
}

// SkuGTE applies the GTE predicate on the "sku" field.
func SkuGTE(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSku), v))
	})
}

// SkuLT applies the LT predicate on the "sku" field.
func SkuLT(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSku), v))
	})
}

// SkuLTE applies the LTE predicate on the "sku" field.
func SkuLTE(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSku), v))
	})
}

// SkuContains applies the Contains predicate on the "sku" field.
func SkuContains(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSku), v))
	})
}

// SkuHasPrefix applies the HasPrefix predicate on the "sku" field.
func SkuHasPrefix(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSku), v))
	})
}

// SkuHasSuffix applies the HasSuffix predicate on the "sku" field.
func SkuHasSuffix(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSku), v))
	})
}

// SkuEqualFold applies the EqualFold predicate on the "sku" field.
func SkuEqualFold(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSku), v))
	})
}

// SkuContainsFold applies the ContainsFold predicate on the "sku" field.
func SkuContainsFold(v string) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSku), v))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...uint64) predicate.BillProduct {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...uint64) predicate.BillProduct {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v uint64) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillProduct) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillProduct) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
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
func Not(p predicate.BillProduct) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		p(s.Not())
	})
}
