// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/carservice"
	"github.com/team07/app/ent/jobposition"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/transport"
	"github.com/team07/app/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the password field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetJobpositionID sets the jobposition edge to JobPosition by id.
func (uu *UserUpdate) SetJobpositionID(id int) *UserUpdate {
	uu.mutation.SetJobpositionID(id)
	return uu
}

// SetNillableJobpositionID sets the jobposition edge to JobPosition by id if the given value is not nil.
func (uu *UserUpdate) SetNillableJobpositionID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetJobpositionID(*id)
	}
	return uu
}

// SetJobposition sets the jobposition edge to JobPosition.
func (uu *UserUpdate) SetJobposition(j *JobPosition) *UserUpdate {
	return uu.SetJobpositionID(j.ID)
}

// AddUserofIDs adds the userof edge to Ambulance by ids.
func (uu *UserUpdate) AddUserofIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserofIDs(ids...)
	return uu
}

// AddUserof adds the userof edges to Ambulance.
func (uu *UserUpdate) AddUserof(a ...*Ambulance) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddUserofIDs(ids...)
}

// AddUseridIDs adds the userid edge to Carservice by ids.
func (uu *UserUpdate) AddUseridIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUseridIDs(ids...)
	return uu
}

// AddUserid adds the userid edges to Carservice.
func (uu *UserUpdate) AddUserid(c ...*Carservice) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddUseridIDs(ids...)
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (uu *UserUpdate) AddCarinspectionIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCarinspectionIDs(ids...)
	return uu
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (uu *UserUpdate) AddCarinspections(c ...*CarInspection) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarinspectionIDs(ids...)
}

// AddCarrepairrecordIDs adds the carrepairrecords edge to CarRepairrecord by ids.
func (uu *UserUpdate) AddCarrepairrecordIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCarrepairrecordIDs(ids...)
	return uu
}

// AddCarrepairrecords adds the carrepairrecords edges to CarRepairrecord.
func (uu *UserUpdate) AddCarrepairrecords(c ...*CarRepairrecord) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarrepairrecordIDs(ids...)
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (uu *UserUpdate) AddCarcheckinoutIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCarcheckinoutIDs(ids...)
	return uu
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (uu *UserUpdate) AddCarcheckinout(c ...*CarCheckInOut) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarcheckinoutIDs(ids...)
}

// AddUserIDs adds the user edge to Transport by ids.
func (uu *UserUpdate) AddUserIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserIDs(ids...)
	return uu
}

// AddUser adds the user edges to Transport.
func (uu *UserUpdate) AddUser(t ...*Transport) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddUserIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearJobposition clears the jobposition edge to JobPosition.
func (uu *UserUpdate) ClearJobposition() *UserUpdate {
	uu.mutation.ClearJobposition()
	return uu
}

// RemoveUserofIDs removes the userof edge to Ambulance by ids.
func (uu *UserUpdate) RemoveUserofIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserofIDs(ids...)
	return uu
}

// RemoveUserof removes userof edges to Ambulance.
func (uu *UserUpdate) RemoveUserof(a ...*Ambulance) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveUserofIDs(ids...)
}

// RemoveUseridIDs removes the userid edge to Carservice by ids.
func (uu *UserUpdate) RemoveUseridIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUseridIDs(ids...)
	return uu
}

// RemoveUserid removes userid edges to Carservice.
func (uu *UserUpdate) RemoveUserid(c ...*Carservice) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveUseridIDs(ids...)
}

// RemoveCarinspectionIDs removes the carinspections edge to CarInspection by ids.
func (uu *UserUpdate) RemoveCarinspectionIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCarinspectionIDs(ids...)
	return uu
}

// RemoveCarinspections removes carinspections edges to CarInspection.
func (uu *UserUpdate) RemoveCarinspections(c ...*CarInspection) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarinspectionIDs(ids...)
}

// RemoveCarrepairrecordIDs removes the carrepairrecords edge to CarRepairrecord by ids.
func (uu *UserUpdate) RemoveCarrepairrecordIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCarrepairrecordIDs(ids...)
	return uu
}

// RemoveCarrepairrecords removes carrepairrecords edges to CarRepairrecord.
func (uu *UserUpdate) RemoveCarrepairrecords(c ...*CarRepairrecord) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarrepairrecordIDs(ids...)
}

// RemoveCarcheckinoutIDs removes the carcheckinout edge to CarCheckInOut by ids.
func (uu *UserUpdate) RemoveCarcheckinoutIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCarcheckinoutIDs(ids...)
	return uu
}

// RemoveCarcheckinout removes carcheckinout edges to CarCheckInOut.
func (uu *UserUpdate) RemoveCarcheckinout(c ...*CarCheckInOut) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarcheckinoutIDs(ids...)
}

// RemoveUserIDs removes the user edge to Transport by ids.
func (uu *UserUpdate) RemoveUserIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserIDs(ids...)
	return uu
}

// RemoveUser removes user edges to Transport.
func (uu *UserUpdate) RemoveUser(t ...*Transport) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if uu.mutation.JobpositionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.JobpositionTable,
			Columns: []string{user.JobpositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jobposition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.JobpositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.JobpositionTable,
			Columns: []string{user.JobpositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jobposition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedUserofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserofTable,
			Columns: []string{user.UserofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserofTable,
			Columns: []string{user.UserofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedUseridIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UseridTable,
			Columns: []string{user.UseridColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UseridIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UseridTable,
			Columns: []string{user.UseridColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedCarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarinspectionsTable,
			Columns: []string{user.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarinspectionsTable,
			Columns: []string{user.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedCarrepairrecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarrepairrecordsTable,
			Columns: []string{user.CarrepairrecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrepairrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CarrepairrecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarrepairrecordsTable,
			Columns: []string{user.CarrepairrecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrepairrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedCarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarcheckinoutTable,
			Columns: []string{user.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarcheckinoutTable,
			Columns: []string{user.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTable,
			Columns: []string{user.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTable,
			Columns: []string{user.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the password field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetJobpositionID sets the jobposition edge to JobPosition by id.
func (uuo *UserUpdateOne) SetJobpositionID(id int) *UserUpdateOne {
	uuo.mutation.SetJobpositionID(id)
	return uuo
}

// SetNillableJobpositionID sets the jobposition edge to JobPosition by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableJobpositionID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetJobpositionID(*id)
	}
	return uuo
}

// SetJobposition sets the jobposition edge to JobPosition.
func (uuo *UserUpdateOne) SetJobposition(j *JobPosition) *UserUpdateOne {
	return uuo.SetJobpositionID(j.ID)
}

// AddUserofIDs adds the userof edge to Ambulance by ids.
func (uuo *UserUpdateOne) AddUserofIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserofIDs(ids...)
	return uuo
}

// AddUserof adds the userof edges to Ambulance.
func (uuo *UserUpdateOne) AddUserof(a ...*Ambulance) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddUserofIDs(ids...)
}

// AddUseridIDs adds the userid edge to Carservice by ids.
func (uuo *UserUpdateOne) AddUseridIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUseridIDs(ids...)
	return uuo
}

// AddUserid adds the userid edges to Carservice.
func (uuo *UserUpdateOne) AddUserid(c ...*Carservice) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddUseridIDs(ids...)
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (uuo *UserUpdateOne) AddCarinspectionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCarinspectionIDs(ids...)
	return uuo
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (uuo *UserUpdateOne) AddCarinspections(c ...*CarInspection) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarinspectionIDs(ids...)
}

// AddCarrepairrecordIDs adds the carrepairrecords edge to CarRepairrecord by ids.
func (uuo *UserUpdateOne) AddCarrepairrecordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCarrepairrecordIDs(ids...)
	return uuo
}

// AddCarrepairrecords adds the carrepairrecords edges to CarRepairrecord.
func (uuo *UserUpdateOne) AddCarrepairrecords(c ...*CarRepairrecord) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarrepairrecordIDs(ids...)
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (uuo *UserUpdateOne) AddCarcheckinoutIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCarcheckinoutIDs(ids...)
	return uuo
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (uuo *UserUpdateOne) AddCarcheckinout(c ...*CarCheckInOut) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarcheckinoutIDs(ids...)
}

// AddUserIDs adds the user edge to Transport by ids.
func (uuo *UserUpdateOne) AddUserIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserIDs(ids...)
	return uuo
}

// AddUser adds the user edges to Transport.
func (uuo *UserUpdateOne) AddUser(t ...*Transport) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddUserIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearJobposition clears the jobposition edge to JobPosition.
func (uuo *UserUpdateOne) ClearJobposition() *UserUpdateOne {
	uuo.mutation.ClearJobposition()
	return uuo
}

// RemoveUserofIDs removes the userof edge to Ambulance by ids.
func (uuo *UserUpdateOne) RemoveUserofIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserofIDs(ids...)
	return uuo
}

// RemoveUserof removes userof edges to Ambulance.
func (uuo *UserUpdateOne) RemoveUserof(a ...*Ambulance) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveUserofIDs(ids...)
}

// RemoveUseridIDs removes the userid edge to Carservice by ids.
func (uuo *UserUpdateOne) RemoveUseridIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUseridIDs(ids...)
	return uuo
}

// RemoveUserid removes userid edges to Carservice.
func (uuo *UserUpdateOne) RemoveUserid(c ...*Carservice) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveUseridIDs(ids...)
}

// RemoveCarinspectionIDs removes the carinspections edge to CarInspection by ids.
func (uuo *UserUpdateOne) RemoveCarinspectionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCarinspectionIDs(ids...)
	return uuo
}

// RemoveCarinspections removes carinspections edges to CarInspection.
func (uuo *UserUpdateOne) RemoveCarinspections(c ...*CarInspection) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarinspectionIDs(ids...)
}

// RemoveCarrepairrecordIDs removes the carrepairrecords edge to CarRepairrecord by ids.
func (uuo *UserUpdateOne) RemoveCarrepairrecordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCarrepairrecordIDs(ids...)
	return uuo
}

// RemoveCarrepairrecords removes carrepairrecords edges to CarRepairrecord.
func (uuo *UserUpdateOne) RemoveCarrepairrecords(c ...*CarRepairrecord) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarrepairrecordIDs(ids...)
}

// RemoveCarcheckinoutIDs removes the carcheckinout edge to CarCheckInOut by ids.
func (uuo *UserUpdateOne) RemoveCarcheckinoutIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCarcheckinoutIDs(ids...)
	return uuo
}

// RemoveCarcheckinout removes carcheckinout edges to CarCheckInOut.
func (uuo *UserUpdateOne) RemoveCarcheckinout(c ...*CarCheckInOut) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarcheckinoutIDs(ids...)
}

// RemoveUserIDs removes the user edge to Transport by ids.
func (uuo *UserUpdateOne) RemoveUserIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserIDs(ids...)
	return uuo
}

// RemoveUser removes user edges to Transport.
func (uuo *UserUpdateOne) RemoveUser(t ...*Transport) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if uuo.mutation.JobpositionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.JobpositionTable,
			Columns: []string{user.JobpositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jobposition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.JobpositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.JobpositionTable,
			Columns: []string{user.JobpositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jobposition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedUserofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserofTable,
			Columns: []string{user.UserofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserofTable,
			Columns: []string{user.UserofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedUseridIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UseridTable,
			Columns: []string{user.UseridColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UseridIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UseridTable,
			Columns: []string{user.UseridColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedCarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarinspectionsTable,
			Columns: []string{user.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarinspectionsTable,
			Columns: []string{user.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedCarrepairrecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarrepairrecordsTable,
			Columns: []string{user.CarrepairrecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrepairrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CarrepairrecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarrepairrecordsTable,
			Columns: []string{user.CarrepairrecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrepairrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedCarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarcheckinoutTable,
			Columns: []string{user.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarcheckinoutTable,
			Columns: []string{user.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTable,
			Columns: []string{user.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTable,
			Columns: []string{user.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
