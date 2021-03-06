// Code generated by entc, DO NOT EDIT.

package hospital

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team07/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hospital applies equality check predicate on the "hospital" field. It's identical to HospitalEQ.
func Hospital(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHospital), v))
	})
}

// HospitalEQ applies the EQ predicate on the "hospital" field.
func HospitalEQ(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHospital), v))
	})
}

// HospitalNEQ applies the NEQ predicate on the "hospital" field.
func HospitalNEQ(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHospital), v))
	})
}

// HospitalIn applies the In predicate on the "hospital" field.
func HospitalIn(vs ...string) predicate.Hospital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hospital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHospital), v...))
	})
}

// HospitalNotIn applies the NotIn predicate on the "hospital" field.
func HospitalNotIn(vs ...string) predicate.Hospital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hospital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHospital), v...))
	})
}

// HospitalGT applies the GT predicate on the "hospital" field.
func HospitalGT(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHospital), v))
	})
}

// HospitalGTE applies the GTE predicate on the "hospital" field.
func HospitalGTE(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHospital), v))
	})
}

// HospitalLT applies the LT predicate on the "hospital" field.
func HospitalLT(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHospital), v))
	})
}

// HospitalLTE applies the LTE predicate on the "hospital" field.
func HospitalLTE(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHospital), v))
	})
}

// HospitalContains applies the Contains predicate on the "hospital" field.
func HospitalContains(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHospital), v))
	})
}

// HospitalHasPrefix applies the HasPrefix predicate on the "hospital" field.
func HospitalHasPrefix(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHospital), v))
	})
}

// HospitalHasSuffix applies the HasSuffix predicate on the "hospital" field.
func HospitalHasSuffix(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHospital), v))
	})
}

// HospitalEqualFold applies the EqualFold predicate on the "hospital" field.
func HospitalEqualFold(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHospital), v))
	})
}

// HospitalContainsFold applies the ContainsFold predicate on the "hospital" field.
func HospitalContainsFold(v string) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHospital), v))
	})
}

// HasReceive applies the HasEdge predicate on the "receive" edge.
func HasReceive() predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiveTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiveTable, ReceiveColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiveWith applies the HasEdge predicate on the "receive" edge with a given conditions (other predicates).
func HasReceiveWith(preds ...predicate.Transport) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiveInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiveTable, ReceiveColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSend applies the HasEdge predicate on the "send" edge.
func HasSend() predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SendTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SendTable, SendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSendWith applies the HasEdge predicate on the "send" edge with a given conditions (other predicates).
func HasSendWith(preds ...predicate.Transport) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SendInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SendTable, SendColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Hospital) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Hospital) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
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
func Not(p predicate.Hospital) predicate.Hospital {
	return predicate.Hospital(func(s *sql.Selector) {
		p(s.Not())
	})
}
