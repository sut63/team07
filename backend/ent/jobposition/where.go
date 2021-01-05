// Code generated by entc, DO NOT EDIT.

package jobposition

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team07/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
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
func IDGT(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PositionName applies equality check predicate on the "position_name" field. It's identical to PositionNameEQ.
func PositionName(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPositionName), v))
	})
}

// PositionNameEQ applies the EQ predicate on the "position_name" field.
func PositionNameEQ(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPositionName), v))
	})
}

// PositionNameNEQ applies the NEQ predicate on the "position_name" field.
func PositionNameNEQ(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPositionName), v))
	})
}

// PositionNameIn applies the In predicate on the "position_name" field.
func PositionNameIn(vs ...string) predicate.JobPosition {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobPosition(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPositionName), v...))
	})
}

// PositionNameNotIn applies the NotIn predicate on the "position_name" field.
func PositionNameNotIn(vs ...string) predicate.JobPosition {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobPosition(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPositionName), v...))
	})
}

// PositionNameGT applies the GT predicate on the "position_name" field.
func PositionNameGT(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPositionName), v))
	})
}

// PositionNameGTE applies the GTE predicate on the "position_name" field.
func PositionNameGTE(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPositionName), v))
	})
}

// PositionNameLT applies the LT predicate on the "position_name" field.
func PositionNameLT(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPositionName), v))
	})
}

// PositionNameLTE applies the LTE predicate on the "position_name" field.
func PositionNameLTE(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPositionName), v))
	})
}

// PositionNameContains applies the Contains predicate on the "position_name" field.
func PositionNameContains(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPositionName), v))
	})
}

// PositionNameHasPrefix applies the HasPrefix predicate on the "position_name" field.
func PositionNameHasPrefix(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPositionName), v))
	})
}

// PositionNameHasSuffix applies the HasSuffix predicate on the "position_name" field.
func PositionNameHasSuffix(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPositionName), v))
	})
}

// PositionNameEqualFold applies the EqualFold predicate on the "position_name" field.
func PositionNameEqualFold(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPositionName), v))
	})
}

// PositionNameContainsFold applies the ContainsFold predicate on the "position_name" field.
func PositionNameContainsFold(v string) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPositionName), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInspectionresults applies the HasEdge predicate on the "inspectionresults" edge.
func HasInspectionresults() predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InspectionresultsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InspectionresultsTable, InspectionresultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInspectionresultsWith applies the HasEdge predicate on the "inspectionresults" edge with a given conditions (other predicates).
func HasInspectionresultsWith(preds ...predicate.InspectionResult) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InspectionresultsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InspectionresultsTable, InspectionresultsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.JobPosition) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.JobPosition) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
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
func Not(p predicate.JobPosition) predicate.JobPosition {
	return predicate.JobPosition(func(s *sql.Selector) {
		p(s.Not())
	})
}
