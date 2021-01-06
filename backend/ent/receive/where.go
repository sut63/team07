// Code generated by entc, DO NOT EDIT.

package receive

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team07/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sendname applies equality check predicate on the "sendname" field. It's identical to SendnameEQ.
func Sendname(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSendname), v))
	})
}

// SendnameEQ applies the EQ predicate on the "sendname" field.
func SendnameEQ(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSendname), v))
	})
}

// SendnameNEQ applies the NEQ predicate on the "sendname" field.
func SendnameNEQ(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSendname), v))
	})
}

// SendnameIn applies the In predicate on the "sendname" field.
func SendnameIn(vs ...string) predicate.Receive {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receive(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSendname), v...))
	})
}

// SendnameNotIn applies the NotIn predicate on the "sendname" field.
func SendnameNotIn(vs ...string) predicate.Receive {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receive(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSendname), v...))
	})
}

// SendnameGT applies the GT predicate on the "sendname" field.
func SendnameGT(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSendname), v))
	})
}

// SendnameGTE applies the GTE predicate on the "sendname" field.
func SendnameGTE(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSendname), v))
	})
}

// SendnameLT applies the LT predicate on the "sendname" field.
func SendnameLT(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSendname), v))
	})
}

// SendnameLTE applies the LTE predicate on the "sendname" field.
func SendnameLTE(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSendname), v))
	})
}

// SendnameContains applies the Contains predicate on the "sendname" field.
func SendnameContains(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSendname), v))
	})
}

// SendnameHasPrefix applies the HasPrefix predicate on the "sendname" field.
func SendnameHasPrefix(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSendname), v))
	})
}

// SendnameHasSuffix applies the HasSuffix predicate on the "sendname" field.
func SendnameHasSuffix(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSendname), v))
	})
}

// SendnameEqualFold applies the EqualFold predicate on the "sendname" field.
func SendnameEqualFold(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSendname), v))
	})
}

// SendnameContainsFold applies the ContainsFold predicate on the "sendname" field.
func SendnameContainsFold(v string) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSendname), v))
	})
}

// HasReceiveid applies the HasEdge predicate on the "receiveid" edge.
func HasReceiveid() predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiveidTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiveidTable, ReceiveidColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiveidWith applies the HasEdge predicate on the "receiveid" edge with a given conditions (other predicates).
func HasReceiveidWith(preds ...predicate.Transport) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiveidInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiveidTable, ReceiveidColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Receive) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Receive) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
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
func Not(p predicate.Receive) predicate.Receive {
	return predicate.Receive(func(s *sql.Selector) {
		p(s.Not())
	})
}
