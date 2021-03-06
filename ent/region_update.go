// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"opencensus/core/ent/organization"
	"opencensus/core/ent/predicate"
	"opencensus/core/ent/province"
	"opencensus/core/ent/region"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RegionUpdate is the builder for updating Region entities.
type RegionUpdate struct {
	config
	hooks    []Hook
	mutation *RegionMutation
}

// Where adds a new predicate for the RegionUpdate builder.
func (ru *RegionUpdate) Where(ps ...predicate.Region) *RegionUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RegionUpdate) SetName(s string) *RegionUpdate {
	ru.mutation.SetName(s)
	return ru
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (ru *RegionUpdate) AddOrganizationIDs(ids ...int) *RegionUpdate {
	ru.mutation.AddOrganizationIDs(ids...)
	return ru
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (ru *RegionUpdate) AddOrganization(o ...*Organization) *RegionUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ru.AddOrganizationIDs(ids...)
}

// AddProvinceIDs adds the "province" edge to the Province entity by IDs.
func (ru *RegionUpdate) AddProvinceIDs(ids ...int) *RegionUpdate {
	ru.mutation.AddProvinceIDs(ids...)
	return ru
}

// AddProvince adds the "province" edges to the Province entity.
func (ru *RegionUpdate) AddProvince(p ...*Province) *RegionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddProvinceIDs(ids...)
}

// Mutation returns the RegionMutation object of the builder.
func (ru *RegionUpdate) Mutation() *RegionMutation {
	return ru.mutation
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (ru *RegionUpdate) ClearOrganization() *RegionUpdate {
	ru.mutation.ClearOrganization()
	return ru
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (ru *RegionUpdate) RemoveOrganizationIDs(ids ...int) *RegionUpdate {
	ru.mutation.RemoveOrganizationIDs(ids...)
	return ru
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (ru *RegionUpdate) RemoveOrganization(o ...*Organization) *RegionUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ru.RemoveOrganizationIDs(ids...)
}

// ClearProvince clears all "province" edges to the Province entity.
func (ru *RegionUpdate) ClearProvince() *RegionUpdate {
	ru.mutation.ClearProvince()
	return ru
}

// RemoveProvinceIDs removes the "province" edge to Province entities by IDs.
func (ru *RegionUpdate) RemoveProvinceIDs(ids ...int) *RegionUpdate {
	ru.mutation.RemoveProvinceIDs(ids...)
	return ru
}

// RemoveProvince removes "province" edges to Province entities.
func (ru *RegionUpdate) RemoveProvince(p ...*Province) *RegionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemoveProvinceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RegionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RegionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RegionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RegionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RegionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   region.Table,
			Columns: region.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: region.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldName,
		})
	}
	if ru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   region.OrganizationTable,
			Columns: region.OrganizationPrimaryKey,
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
	if nodes := ru.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !ru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   region.OrganizationTable,
			Columns: region.OrganizationPrimaryKey,
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
	if nodes := ru.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   region.OrganizationTable,
			Columns: region.OrganizationPrimaryKey,
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
	if ru.mutation.ProvinceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   region.ProvinceTable,
			Columns: region.ProvincePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedProvinceIDs(); len(nodes) > 0 && !ru.mutation.ProvinceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   region.ProvinceTable,
			Columns: region.ProvincePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   region.ProvinceTable,
			Columns: region.ProvincePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{region.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RegionUpdateOne is the builder for updating a single Region entity.
type RegionUpdateOne struct {
	config
	hooks    []Hook
	mutation *RegionMutation
}

// SetName sets the "name" field.
func (ruo *RegionUpdateOne) SetName(s string) *RegionUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (ruo *RegionUpdateOne) AddOrganizationIDs(ids ...int) *RegionUpdateOne {
	ruo.mutation.AddOrganizationIDs(ids...)
	return ruo
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (ruo *RegionUpdateOne) AddOrganization(o ...*Organization) *RegionUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ruo.AddOrganizationIDs(ids...)
}

// AddProvinceIDs adds the "province" edge to the Province entity by IDs.
func (ruo *RegionUpdateOne) AddProvinceIDs(ids ...int) *RegionUpdateOne {
	ruo.mutation.AddProvinceIDs(ids...)
	return ruo
}

// AddProvince adds the "province" edges to the Province entity.
func (ruo *RegionUpdateOne) AddProvince(p ...*Province) *RegionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddProvinceIDs(ids...)
}

// Mutation returns the RegionMutation object of the builder.
func (ruo *RegionUpdateOne) Mutation() *RegionMutation {
	return ruo.mutation
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (ruo *RegionUpdateOne) ClearOrganization() *RegionUpdateOne {
	ruo.mutation.ClearOrganization()
	return ruo
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (ruo *RegionUpdateOne) RemoveOrganizationIDs(ids ...int) *RegionUpdateOne {
	ruo.mutation.RemoveOrganizationIDs(ids...)
	return ruo
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (ruo *RegionUpdateOne) RemoveOrganization(o ...*Organization) *RegionUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ruo.RemoveOrganizationIDs(ids...)
}

// ClearProvince clears all "province" edges to the Province entity.
func (ruo *RegionUpdateOne) ClearProvince() *RegionUpdateOne {
	ruo.mutation.ClearProvince()
	return ruo
}

// RemoveProvinceIDs removes the "province" edge to Province entities by IDs.
func (ruo *RegionUpdateOne) RemoveProvinceIDs(ids ...int) *RegionUpdateOne {
	ruo.mutation.RemoveProvinceIDs(ids...)
	return ruo
}

// RemoveProvince removes "province" edges to Province entities.
func (ruo *RegionUpdateOne) RemoveProvince(p ...*Province) *RegionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemoveProvinceIDs(ids...)
}

// Save executes the query and returns the updated Region entity.
func (ruo *RegionUpdateOne) Save(ctx context.Context) (*Region, error) {
	var (
		err  error
		node *Region
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RegionUpdateOne) SaveX(ctx context.Context) *Region {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RegionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RegionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RegionUpdateOne) sqlSave(ctx context.Context) (_node *Region, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   region.Table,
			Columns: region.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: region.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Region.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldName,
		})
	}
	if ruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   region.OrganizationTable,
			Columns: region.OrganizationPrimaryKey,
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
	if nodes := ruo.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !ruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   region.OrganizationTable,
			Columns: region.OrganizationPrimaryKey,
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
	if nodes := ruo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   region.OrganizationTable,
			Columns: region.OrganizationPrimaryKey,
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
	if ruo.mutation.ProvinceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   region.ProvinceTable,
			Columns: region.ProvincePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedProvinceIDs(); len(nodes) > 0 && !ruo.mutation.ProvinceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   region.ProvinceTable,
			Columns: region.ProvincePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   region.ProvinceTable,
			Columns: region.ProvincePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Region{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{region.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
