// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"opencensus/core/ent/organization"
	"opencensus/core/ent/oxygenrecord"
	"opencensus/core/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OxygenRecordUpdate is the builder for updating OxygenRecord entities.
type OxygenRecordUpdate struct {
	config
	hooks    []Hook
	mutation *OxygenRecordMutation
}

// Where adds a new predicate for the OxygenRecordUpdate builder.
func (oru *OxygenRecordUpdate) Where(ps ...predicate.OxygenRecord) *OxygenRecordUpdate {
	oru.mutation.predicates = append(oru.mutation.predicates, ps...)
	return oru
}

// SetTotalCylinders sets the "totalCylinders" field.
func (oru *OxygenRecordUpdate) SetTotalCylinders(i int) *OxygenRecordUpdate {
	oru.mutation.ResetTotalCylinders()
	oru.mutation.SetTotalCylinders(i)
	return oru
}

// AddTotalCylinders adds i to the "totalCylinders" field.
func (oru *OxygenRecordUpdate) AddTotalCylinders(i int) *OxygenRecordUpdate {
	oru.mutation.AddTotalCylinders(i)
	return oru
}

// SetTotalOwnCylinders sets the "totalOwnCylinders" field.
func (oru *OxygenRecordUpdate) SetTotalOwnCylinders(i int) *OxygenRecordUpdate {
	oru.mutation.ResetTotalOwnCylinders()
	oru.mutation.SetTotalOwnCylinders(i)
	return oru
}

// AddTotalOwnCylinders adds i to the "totalOwnCylinders" field.
func (oru *OxygenRecordUpdate) AddTotalOwnCylinders(i int) *OxygenRecordUpdate {
	oru.mutation.AddTotalOwnCylinders(i)
	return oru
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (oru *OxygenRecordUpdate) AddOrganizationIDs(ids ...int) *OxygenRecordUpdate {
	oru.mutation.AddOrganizationIDs(ids...)
	return oru
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (oru *OxygenRecordUpdate) AddOrganization(o ...*Organization) *OxygenRecordUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oru.AddOrganizationIDs(ids...)
}

// Mutation returns the OxygenRecordMutation object of the builder.
func (oru *OxygenRecordUpdate) Mutation() *OxygenRecordMutation {
	return oru.mutation
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (oru *OxygenRecordUpdate) ClearOrganization() *OxygenRecordUpdate {
	oru.mutation.ClearOrganization()
	return oru
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (oru *OxygenRecordUpdate) RemoveOrganizationIDs(ids ...int) *OxygenRecordUpdate {
	oru.mutation.RemoveOrganizationIDs(ids...)
	return oru
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (oru *OxygenRecordUpdate) RemoveOrganization(o ...*Organization) *OxygenRecordUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oru.RemoveOrganizationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oru *OxygenRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oru.hooks) == 0 {
		affected, err = oru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OxygenRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oru.mutation = mutation
			affected, err = oru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oru.hooks) - 1; i >= 0; i-- {
			mut = oru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oru *OxygenRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := oru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oru *OxygenRecordUpdate) Exec(ctx context.Context) error {
	_, err := oru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oru *OxygenRecordUpdate) ExecX(ctx context.Context) {
	if err := oru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oru *OxygenRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oxygenrecord.Table,
			Columns: oxygenrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oxygenrecord.FieldID,
			},
		},
	}
	if ps := oru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oru.mutation.TotalCylinders(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalCylinders,
		})
	}
	if value, ok := oru.mutation.AddedTotalCylinders(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalCylinders,
		})
	}
	if value, ok := oru.mutation.TotalOwnCylinders(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalOwnCylinders,
		})
	}
	if value, ok := oru.mutation.AddedTotalOwnCylinders(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalOwnCylinders,
		})
	}
	if oru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oxygenrecord.OrganizationTable,
			Columns: oxygenrecord.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oru.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !oru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oxygenrecord.OrganizationTable,
			Columns: oxygenrecord.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oru.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oxygenrecord.OrganizationTable,
			Columns: oxygenrecord.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oxygenrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OxygenRecordUpdateOne is the builder for updating a single OxygenRecord entity.
type OxygenRecordUpdateOne struct {
	config
	hooks    []Hook
	mutation *OxygenRecordMutation
}

// SetTotalCylinders sets the "totalCylinders" field.
func (oruo *OxygenRecordUpdateOne) SetTotalCylinders(i int) *OxygenRecordUpdateOne {
	oruo.mutation.ResetTotalCylinders()
	oruo.mutation.SetTotalCylinders(i)
	return oruo
}

// AddTotalCylinders adds i to the "totalCylinders" field.
func (oruo *OxygenRecordUpdateOne) AddTotalCylinders(i int) *OxygenRecordUpdateOne {
	oruo.mutation.AddTotalCylinders(i)
	return oruo
}

// SetTotalOwnCylinders sets the "totalOwnCylinders" field.
func (oruo *OxygenRecordUpdateOne) SetTotalOwnCylinders(i int) *OxygenRecordUpdateOne {
	oruo.mutation.ResetTotalOwnCylinders()
	oruo.mutation.SetTotalOwnCylinders(i)
	return oruo
}

// AddTotalOwnCylinders adds i to the "totalOwnCylinders" field.
func (oruo *OxygenRecordUpdateOne) AddTotalOwnCylinders(i int) *OxygenRecordUpdateOne {
	oruo.mutation.AddTotalOwnCylinders(i)
	return oruo
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (oruo *OxygenRecordUpdateOne) AddOrganizationIDs(ids ...int) *OxygenRecordUpdateOne {
	oruo.mutation.AddOrganizationIDs(ids...)
	return oruo
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (oruo *OxygenRecordUpdateOne) AddOrganization(o ...*Organization) *OxygenRecordUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oruo.AddOrganizationIDs(ids...)
}

// Mutation returns the OxygenRecordMutation object of the builder.
func (oruo *OxygenRecordUpdateOne) Mutation() *OxygenRecordMutation {
	return oruo.mutation
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (oruo *OxygenRecordUpdateOne) ClearOrganization() *OxygenRecordUpdateOne {
	oruo.mutation.ClearOrganization()
	return oruo
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (oruo *OxygenRecordUpdateOne) RemoveOrganizationIDs(ids ...int) *OxygenRecordUpdateOne {
	oruo.mutation.RemoveOrganizationIDs(ids...)
	return oruo
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (oruo *OxygenRecordUpdateOne) RemoveOrganization(o ...*Organization) *OxygenRecordUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oruo.RemoveOrganizationIDs(ids...)
}

// Save executes the query and returns the updated OxygenRecord entity.
func (oruo *OxygenRecordUpdateOne) Save(ctx context.Context) (*OxygenRecord, error) {
	var (
		err  error
		node *OxygenRecord
	)
	if len(oruo.hooks) == 0 {
		node, err = oruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OxygenRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oruo.mutation = mutation
			node, err = oruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oruo.hooks) - 1; i >= 0; i-- {
			mut = oruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oruo *OxygenRecordUpdateOne) SaveX(ctx context.Context) *OxygenRecord {
	node, err := oruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oruo *OxygenRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := oruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oruo *OxygenRecordUpdateOne) ExecX(ctx context.Context) {
	if err := oruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oruo *OxygenRecordUpdateOne) sqlSave(ctx context.Context) (_node *OxygenRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oxygenrecord.Table,
			Columns: oxygenrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oxygenrecord.FieldID,
			},
		},
	}
	id, ok := oruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OxygenRecord.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := oruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oruo.mutation.TotalCylinders(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalCylinders,
		})
	}
	if value, ok := oruo.mutation.AddedTotalCylinders(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalCylinders,
		})
	}
	if value, ok := oruo.mutation.TotalOwnCylinders(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalOwnCylinders,
		})
	}
	if value, ok := oruo.mutation.AddedTotalOwnCylinders(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldTotalOwnCylinders,
		})
	}
	if oruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oxygenrecord.OrganizationTable,
			Columns: oxygenrecord.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruo.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !oruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oxygenrecord.OrganizationTable,
			Columns: oxygenrecord.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oxygenrecord.OrganizationTable,
			Columns: oxygenrecord.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OxygenRecord{config: oruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oxygenrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
