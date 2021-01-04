// Code generated by entc, DO NOT EDIT.

package repairing

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team07/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Repairpart applies equality check predicate on the "repairpart" field. It's identical to RepairpartEQ.
func Repairpart(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepairpart), v))
	})
}

// RepairpartEQ applies the EQ predicate on the "repairpart" field.
func RepairpartEQ(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepairpart), v))
	})
}

// RepairpartNEQ applies the NEQ predicate on the "repairpart" field.
func RepairpartNEQ(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRepairpart), v))
	})
}

// RepairpartIn applies the In predicate on the "repairpart" field.
func RepairpartIn(vs ...string) predicate.Repairing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRepairpart), v...))
	})
}

// RepairpartNotIn applies the NotIn predicate on the "repairpart" field.
func RepairpartNotIn(vs ...string) predicate.Repairing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRepairpart), v...))
	})
}

// RepairpartGT applies the GT predicate on the "repairpart" field.
func RepairpartGT(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRepairpart), v))
	})
}

// RepairpartGTE applies the GTE predicate on the "repairpart" field.
func RepairpartGTE(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRepairpart), v))
	})
}

// RepairpartLT applies the LT predicate on the "repairpart" field.
func RepairpartLT(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRepairpart), v))
	})
}

// RepairpartLTE applies the LTE predicate on the "repairpart" field.
func RepairpartLTE(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRepairpart), v))
	})
}

// RepairpartContains applies the Contains predicate on the "repairpart" field.
func RepairpartContains(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRepairpart), v))
	})
}

// RepairpartHasPrefix applies the HasPrefix predicate on the "repairpart" field.
func RepairpartHasPrefix(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRepairpart), v))
	})
}

// RepairpartHasSuffix applies the HasSuffix predicate on the "repairpart" field.
func RepairpartHasSuffix(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRepairpart), v))
	})
}

// RepairpartEqualFold applies the EqualFold predicate on the "repairpart" field.
func RepairpartEqualFold(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRepairpart), v))
	})
}

// RepairpartContainsFold applies the ContainsFold predicate on the "repairpart" field.
func RepairpartContainsFold(v string) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRepairpart), v))
	})
}

// HasRepairs applies the HasEdge predicate on the "repairs" edge.
func HasRepairs() predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepairsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepairsTable, RepairsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepairsWith applies the HasEdge predicate on the "repairs" edge with a given conditions (other predicates).
func HasRepairsWith(preds ...predicate.CarRepairrecord) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepairsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepairsTable, RepairsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Repairing) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Repairing) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
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
func Not(p predicate.Repairing) predicate.Repairing {
	return predicate.Repairing(func(s *sql.Selector) {
		p(s.Not())
	})
}
