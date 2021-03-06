// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/team07/app/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AmbulanceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AmbulanceQueryRuleFunc func(context.Context, *ent.AmbulanceQuery) error

// EvalQuery return f(ctx, q).
func (f AmbulanceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AmbulanceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AmbulanceQuery", q)
}

// The AmbulanceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AmbulanceMutationRuleFunc func(context.Context, *ent.AmbulanceMutation) error

// EvalMutation calls f(ctx, m).
func (f AmbulanceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AmbulanceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AmbulanceMutation", m)
}

// The CarCheckInOutQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarCheckInOutQueryRuleFunc func(context.Context, *ent.CarCheckInOutQuery) error

// EvalQuery return f(ctx, q).
func (f CarCheckInOutQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarCheckInOutQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarCheckInOutQuery", q)
}

// The CarCheckInOutMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarCheckInOutMutationRuleFunc func(context.Context, *ent.CarCheckInOutMutation) error

// EvalMutation calls f(ctx, m).
func (f CarCheckInOutMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarCheckInOutMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarCheckInOutMutation", m)
}

// The CarInspectionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarInspectionQueryRuleFunc func(context.Context, *ent.CarInspectionQuery) error

// EvalQuery return f(ctx, q).
func (f CarInspectionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarInspectionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarInspectionQuery", q)
}

// The CarInspectionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarInspectionMutationRuleFunc func(context.Context, *ent.CarInspectionMutation) error

// EvalMutation calls f(ctx, m).
func (f CarInspectionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarInspectionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarInspectionMutation", m)
}

// The CarRepairrecordQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarRepairrecordQueryRuleFunc func(context.Context, *ent.CarRepairrecordQuery) error

// EvalQuery return f(ctx, q).
func (f CarRepairrecordQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarRepairrecordQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarRepairrecordQuery", q)
}

// The CarRepairrecordMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarRepairrecordMutationRuleFunc func(context.Context, *ent.CarRepairrecordMutation) error

// EvalMutation calls f(ctx, m).
func (f CarRepairrecordMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarRepairrecordMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarRepairrecordMutation", m)
}

// The CarbrandQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarbrandQueryRuleFunc func(context.Context, *ent.CarbrandQuery) error

// EvalQuery return f(ctx, q).
func (f CarbrandQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarbrandQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarbrandQuery", q)
}

// The CarbrandMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarbrandMutationRuleFunc func(context.Context, *ent.CarbrandMutation) error

// EvalMutation calls f(ctx, m).
func (f CarbrandMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarbrandMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarbrandMutation", m)
}

// The CarregisterQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarregisterQueryRuleFunc func(context.Context, *ent.CarregisterQuery) error

// EvalQuery return f(ctx, q).
func (f CarregisterQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarregisterQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarregisterQuery", q)
}

// The CarregisterMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarregisterMutationRuleFunc func(context.Context, *ent.CarregisterMutation) error

// EvalMutation calls f(ctx, m).
func (f CarregisterMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarregisterMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarregisterMutation", m)
}

// The CarserviceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarserviceQueryRuleFunc func(context.Context, *ent.CarserviceQuery) error

// EvalQuery return f(ctx, q).
func (f CarserviceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarserviceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarserviceQuery", q)
}

// The CarserviceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarserviceMutationRuleFunc func(context.Context, *ent.CarserviceMutation) error

// EvalMutation calls f(ctx, m).
func (f CarserviceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarserviceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarserviceMutation", m)
}

// The DeliverQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DeliverQueryRuleFunc func(context.Context, *ent.DeliverQuery) error

// EvalQuery return f(ctx, q).
func (f DeliverQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeliverQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DeliverQuery", q)
}

// The DeliverMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DeliverMutationRuleFunc func(context.Context, *ent.DeliverMutation) error

// EvalMutation calls f(ctx, m).
func (f DeliverMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DeliverMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DeliverMutation", m)
}

// The DistanceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DistanceQueryRuleFunc func(context.Context, *ent.DistanceQuery) error

// EvalQuery return f(ctx, q).
func (f DistanceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DistanceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DistanceQuery", q)
}

// The DistanceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DistanceMutationRuleFunc func(context.Context, *ent.DistanceMutation) error

// EvalMutation calls f(ctx, m).
func (f DistanceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DistanceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DistanceMutation", m)
}

// The HospitalQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type HospitalQueryRuleFunc func(context.Context, *ent.HospitalQuery) error

// EvalQuery return f(ctx, q).
func (f HospitalQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.HospitalQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.HospitalQuery", q)
}

// The HospitalMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type HospitalMutationRuleFunc func(context.Context, *ent.HospitalMutation) error

// EvalMutation calls f(ctx, m).
func (f HospitalMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.HospitalMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.HospitalMutation", m)
}

// The InspectionResultQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InspectionResultQueryRuleFunc func(context.Context, *ent.InspectionResultQuery) error

// EvalQuery return f(ctx, q).
func (f InspectionResultQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InspectionResultQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InspectionResultQuery", q)
}

// The InspectionResultMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InspectionResultMutationRuleFunc func(context.Context, *ent.InspectionResultMutation) error

// EvalMutation calls f(ctx, m).
func (f InspectionResultMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InspectionResultMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InspectionResultMutation", m)
}

// The InsuranceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InsuranceQueryRuleFunc func(context.Context, *ent.InsuranceQuery) error

// EvalQuery return f(ctx, q).
func (f InsuranceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InsuranceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InsuranceQuery", q)
}

// The InsuranceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InsuranceMutationRuleFunc func(context.Context, *ent.InsuranceMutation) error

// EvalMutation calls f(ctx, m).
func (f InsuranceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InsuranceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InsuranceMutation", m)
}

// The JobPositionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type JobPositionQueryRuleFunc func(context.Context, *ent.JobPositionQuery) error

// EvalQuery return f(ctx, q).
func (f JobPositionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.JobPositionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.JobPositionQuery", q)
}

// The JobPositionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type JobPositionMutationRuleFunc func(context.Context, *ent.JobPositionMutation) error

// EvalMutation calls f(ctx, m).
func (f JobPositionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.JobPositionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.JobPositionMutation", m)
}

// The PurposeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PurposeQueryRuleFunc func(context.Context, *ent.PurposeQuery) error

// EvalQuery return f(ctx, q).
func (f PurposeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PurposeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PurposeQuery", q)
}

// The PurposeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PurposeMutationRuleFunc func(context.Context, *ent.PurposeMutation) error

// EvalMutation calls f(ctx, m).
func (f PurposeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PurposeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PurposeMutation", m)
}

// The RepairingQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RepairingQueryRuleFunc func(context.Context, *ent.RepairingQuery) error

// EvalQuery return f(ctx, q).
func (f RepairingQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RepairingQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RepairingQuery", q)
}

// The RepairingMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RepairingMutationRuleFunc func(context.Context, *ent.RepairingMutation) error

// EvalMutation calls f(ctx, m).
func (f RepairingMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RepairingMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RepairingMutation", m)
}

// The TransportQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TransportQueryRuleFunc func(context.Context, *ent.TransportQuery) error

// EvalQuery return f(ctx, q).
func (f TransportQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TransportQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TransportQuery", q)
}

// The TransportMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TransportMutationRuleFunc func(context.Context, *ent.TransportMutation) error

// EvalMutation calls f(ctx, m).
func (f TransportMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TransportMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TransportMutation", m)
}

// The UrgentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UrgentQueryRuleFunc func(context.Context, *ent.UrgentQuery) error

// EvalQuery return f(ctx, q).
func (f UrgentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UrgentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UrgentQuery", q)
}

// The UrgentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UrgentMutationRuleFunc func(context.Context, *ent.UrgentMutation) error

// EvalMutation calls f(ctx, m).
func (f UrgentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UrgentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UrgentMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}
