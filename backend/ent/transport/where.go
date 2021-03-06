// Code generated by entc, DO NOT EDIT.

package transport

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team07/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Symptom applies equality check predicate on the "symptom" field. It's identical to SymptomEQ.
func Symptom(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// Drugallergy applies equality check predicate on the "drugallergy" field. It's identical to DrugallergyEQ.
func Drugallergy(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugallergy), v))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// SymptomEQ applies the EQ predicate on the "symptom" field.
func SymptomEQ(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// SymptomNEQ applies the NEQ predicate on the "symptom" field.
func SymptomNEQ(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymptom), v))
	})
}

// SymptomIn applies the In predicate on the "symptom" field.
func SymptomIn(vs ...string) predicate.Transport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymptom), v...))
	})
}

// SymptomNotIn applies the NotIn predicate on the "symptom" field.
func SymptomNotIn(vs ...string) predicate.Transport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymptom), v...))
	})
}

// SymptomGT applies the GT predicate on the "symptom" field.
func SymptomGT(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymptom), v))
	})
}

// SymptomGTE applies the GTE predicate on the "symptom" field.
func SymptomGTE(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymptom), v))
	})
}

// SymptomLT applies the LT predicate on the "symptom" field.
func SymptomLT(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymptom), v))
	})
}

// SymptomLTE applies the LTE predicate on the "symptom" field.
func SymptomLTE(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymptom), v))
	})
}

// SymptomContains applies the Contains predicate on the "symptom" field.
func SymptomContains(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymptom), v))
	})
}

// SymptomHasPrefix applies the HasPrefix predicate on the "symptom" field.
func SymptomHasPrefix(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymptom), v))
	})
}

// SymptomHasSuffix applies the HasSuffix predicate on the "symptom" field.
func SymptomHasSuffix(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymptom), v))
	})
}

// SymptomEqualFold applies the EqualFold predicate on the "symptom" field.
func SymptomEqualFold(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymptom), v))
	})
}

// SymptomContainsFold applies the ContainsFold predicate on the "symptom" field.
func SymptomContainsFold(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymptom), v))
	})
}

// DrugallergyEQ applies the EQ predicate on the "drugallergy" field.
func DrugallergyEQ(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyNEQ applies the NEQ predicate on the "drugallergy" field.
func DrugallergyNEQ(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyIn applies the In predicate on the "drugallergy" field.
func DrugallergyIn(vs ...string) predicate.Transport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDrugallergy), v...))
	})
}

// DrugallergyNotIn applies the NotIn predicate on the "drugallergy" field.
func DrugallergyNotIn(vs ...string) predicate.Transport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDrugallergy), v...))
	})
}

// DrugallergyGT applies the GT predicate on the "drugallergy" field.
func DrugallergyGT(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyGTE applies the GTE predicate on the "drugallergy" field.
func DrugallergyGTE(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyLT applies the LT predicate on the "drugallergy" field.
func DrugallergyLT(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyLTE applies the LTE predicate on the "drugallergy" field.
func DrugallergyLTE(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyContains applies the Contains predicate on the "drugallergy" field.
func DrugallergyContains(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyHasPrefix applies the HasPrefix predicate on the "drugallergy" field.
func DrugallergyHasPrefix(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyHasSuffix applies the HasSuffix predicate on the "drugallergy" field.
func DrugallergyHasSuffix(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyEqualFold applies the EqualFold predicate on the "drugallergy" field.
func DrugallergyEqualFold(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDrugallergy), v))
	})
}

// DrugallergyContainsFold applies the ContainsFold predicate on the "drugallergy" field.
func DrugallergyContainsFold(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDrugallergy), v))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Transport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Transport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// HasSend applies the HasEdge predicate on the "send" edge.
func HasSend() predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SendTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SendTable, SendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSendWith applies the HasEdge predicate on the "send" edge with a given conditions (other predicates).
func HasSendWith(preds ...predicate.Hospital) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SendInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SendTable, SendColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceive applies the HasEdge predicate on the "receive" edge.
func HasReceive() predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiveTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReceiveTable, ReceiveColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiveWith applies the HasEdge predicate on the "receive" edge with a given conditions (other predicates).
func HasReceiveWith(preds ...predicate.Hospital) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiveInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReceiveTable, ReceiveColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAmbulance applies the HasEdge predicate on the "ambulance" edge.
func HasAmbulance() predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AmbulanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AmbulanceTable, AmbulanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAmbulanceWith applies the HasEdge predicate on the "ambulance" edge with a given conditions (other predicates).
func HasAmbulanceWith(preds ...predicate.Ambulance) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AmbulanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AmbulanceTable, AmbulanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Transport) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Transport) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
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
func Not(p predicate.Transport) predicate.Transport {
	return predicate.Transport(func(s *sql.Selector) {
		p(s.Not())
	})
}
