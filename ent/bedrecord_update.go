// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"opencensus/core/ent/bedrecord"
	"opencensus/core/ent/organization"
	"opencensus/core/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BedRecordUpdate is the builder for updating BedRecord entities.
type BedRecordUpdate struct {
	config
	hooks    []Hook
	mutation *BedRecordMutation
}

// Where adds a new predicate for the BedRecordUpdate builder.
func (bru *BedRecordUpdate) Where(ps ...predicate.BedRecord) *BedRecordUpdate {
	bru.mutation.predicates = append(bru.mutation.predicates, ps...)
	return bru
}

// SetBusyCovidBeds sets the "busyCovidBeds" field.
func (bru *BedRecordUpdate) SetBusyCovidBeds(i int) *BedRecordUpdate {
	bru.mutation.ResetBusyCovidBeds()
	bru.mutation.SetBusyCovidBeds(i)
	return bru
}

// AddBusyCovidBeds adds i to the "busyCovidBeds" field.
func (bru *BedRecordUpdate) AddBusyCovidBeds(i int) *BedRecordUpdate {
	bru.mutation.AddBusyCovidBeds(i)
	return bru
}

// SetAvailableCovidBeds sets the "availableCovidBeds" field.
func (bru *BedRecordUpdate) SetAvailableCovidBeds(i int) *BedRecordUpdate {
	bru.mutation.ResetAvailableCovidBeds()
	bru.mutation.SetAvailableCovidBeds(i)
	return bru
}

// AddAvailableCovidBeds adds i to the "availableCovidBeds" field.
func (bru *BedRecordUpdate) AddAvailableCovidBeds(i int) *BedRecordUpdate {
	bru.mutation.AddAvailableCovidBeds(i)
	return bru
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (bru *BedRecordUpdate) AddOrganizationIDs(ids ...int) *BedRecordUpdate {
	bru.mutation.AddOrganizationIDs(ids...)
	return bru
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (bru *BedRecordUpdate) AddOrganization(o ...*Organization) *BedRecordUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return bru.AddOrganizationIDs(ids...)
}

// Mutation returns the BedRecordMutation object of the builder.
func (bru *BedRecordUpdate) Mutation() *BedRecordMutation {
	return bru.mutation
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (bru *BedRecordUpdate) ClearOrganization() *BedRecordUpdate {
	bru.mutation.ClearOrganization()
	return bru
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (bru *BedRecordUpdate) RemoveOrganizationIDs(ids ...int) *BedRecordUpdate {
	bru.mutation.RemoveOrganizationIDs(ids...)
	return bru
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (bru *BedRecordUpdate) RemoveOrganization(o ...*Organization) *BedRecordUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return bru.RemoveOrganizationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bru *BedRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bru.hooks) == 0 {
		affected, err = bru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BedRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bru.mutation = mutation
			affected, err = bru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bru.hooks) - 1; i >= 0; i-- {
			mut = bru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bru *BedRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := bru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bru *BedRecordUpdate) Exec(ctx context.Context) error {
	_, err := bru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bru *BedRecordUpdate) ExecX(ctx context.Context) {
	if err := bru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bru *BedRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bedrecord.Table,
			Columns: bedrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bedrecord.FieldID,
			},
		},
	}
	if ps := bru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bru.mutation.BusyCovidBeds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldBusyCovidBeds,
		})
	}
	if value, ok := bru.mutation.AddedBusyCovidBeds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldBusyCovidBeds,
		})
	}
	if value, ok := bru.mutation.AvailableCovidBeds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldAvailableCovidBeds,
		})
	}
	if value, ok := bru.mutation.AddedAvailableCovidBeds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldAvailableCovidBeds,
		})
	}
	if bru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bedrecord.OrganizationTable,
			Columns: bedrecord.OrganizationPrimaryKey,
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
	if nodes := bru.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !bru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bedrecord.OrganizationTable,
			Columns: bedrecord.OrganizationPrimaryKey,
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
	if nodes := bru.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bedrecord.OrganizationTable,
			Columns: bedrecord.OrganizationPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, bru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bedrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BedRecordUpdateOne is the builder for updating a single BedRecord entity.
type BedRecordUpdateOne struct {
	config
	hooks    []Hook
	mutation *BedRecordMutation
}

// SetBusyCovidBeds sets the "busyCovidBeds" field.
func (bruo *BedRecordUpdateOne) SetBusyCovidBeds(i int) *BedRecordUpdateOne {
	bruo.mutation.ResetBusyCovidBeds()
	bruo.mutation.SetBusyCovidBeds(i)
	return bruo
}

// AddBusyCovidBeds adds i to the "busyCovidBeds" field.
func (bruo *BedRecordUpdateOne) AddBusyCovidBeds(i int) *BedRecordUpdateOne {
	bruo.mutation.AddBusyCovidBeds(i)
	return bruo
}

// SetAvailableCovidBeds sets the "availableCovidBeds" field.
func (bruo *BedRecordUpdateOne) SetAvailableCovidBeds(i int) *BedRecordUpdateOne {
	bruo.mutation.ResetAvailableCovidBeds()
	bruo.mutation.SetAvailableCovidBeds(i)
	return bruo
}

// AddAvailableCovidBeds adds i to the "availableCovidBeds" field.
func (bruo *BedRecordUpdateOne) AddAvailableCovidBeds(i int) *BedRecordUpdateOne {
	bruo.mutation.AddAvailableCovidBeds(i)
	return bruo
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (bruo *BedRecordUpdateOne) AddOrganizationIDs(ids ...int) *BedRecordUpdateOne {
	bruo.mutation.AddOrganizationIDs(ids...)
	return bruo
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (bruo *BedRecordUpdateOne) AddOrganization(o ...*Organization) *BedRecordUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return bruo.AddOrganizationIDs(ids...)
}

// Mutation returns the BedRecordMutation object of the builder.
func (bruo *BedRecordUpdateOne) Mutation() *BedRecordMutation {
	return bruo.mutation
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (bruo *BedRecordUpdateOne) ClearOrganization() *BedRecordUpdateOne {
	bruo.mutation.ClearOrganization()
	return bruo
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (bruo *BedRecordUpdateOne) RemoveOrganizationIDs(ids ...int) *BedRecordUpdateOne {
	bruo.mutation.RemoveOrganizationIDs(ids...)
	return bruo
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (bruo *BedRecordUpdateOne) RemoveOrganization(o ...*Organization) *BedRecordUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return bruo.RemoveOrganizationIDs(ids...)
}

// Save executes the query and returns the updated BedRecord entity.
func (bruo *BedRecordUpdateOne) Save(ctx context.Context) (*BedRecord, error) {
	var (
		err  error
		node *BedRecord
	)
	if len(bruo.hooks) == 0 {
		node, err = bruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BedRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bruo.mutation = mutation
			node, err = bruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bruo.hooks) - 1; i >= 0; i-- {
			mut = bruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bruo *BedRecordUpdateOne) SaveX(ctx context.Context) *BedRecord {
	node, err := bruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bruo *BedRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := bruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bruo *BedRecordUpdateOne) ExecX(ctx context.Context) {
	if err := bruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bruo *BedRecordUpdateOne) sqlSave(ctx context.Context) (_node *BedRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bedrecord.Table,
			Columns: bedrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bedrecord.FieldID,
			},
		},
	}
	id, ok := bruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BedRecord.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := bruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bruo.mutation.BusyCovidBeds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldBusyCovidBeds,
		})
	}
	if value, ok := bruo.mutation.AddedBusyCovidBeds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldBusyCovidBeds,
		})
	}
	if value, ok := bruo.mutation.AvailableCovidBeds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldAvailableCovidBeds,
		})
	}
	if value, ok := bruo.mutation.AddedAvailableCovidBeds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bedrecord.FieldAvailableCovidBeds,
		})
	}
	if bruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bedrecord.OrganizationTable,
			Columns: bedrecord.OrganizationPrimaryKey,
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
	if nodes := bruo.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !bruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bedrecord.OrganizationTable,
			Columns: bedrecord.OrganizationPrimaryKey,
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
	if nodes := bruo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bedrecord.OrganizationTable,
			Columns: bedrecord.OrganizationPrimaryKey,
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
	_node = &BedRecord{config: bruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bedrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
