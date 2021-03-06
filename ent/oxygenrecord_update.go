// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"opencensus/core/ent/oxygenrecord"
	"opencensus/core/ent/place"
	"opencensus/core/ent/predicate"
	"time"

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

// SetReportedDate sets the "reportedDate" field.
func (oru *OxygenRecordUpdate) SetReportedDate(t time.Time) *OxygenRecordUpdate {
	oru.mutation.SetReportedDate(t)
	return oru
}

// SetCollectedDate sets the "collectedDate" field.
func (oru *OxygenRecordUpdate) SetCollectedDate(t time.Time) *OxygenRecordUpdate {
	oru.mutation.SetCollectedDate(t)
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

// SetDailyProduction sets the "dailyProduction" field.
func (oru *OxygenRecordUpdate) SetDailyProduction(i int) *OxygenRecordUpdate {
	oru.mutation.ResetDailyProduction()
	oru.mutation.SetDailyProduction(i)
	return oru
}

// AddDailyProduction adds i to the "dailyProduction" field.
func (oru *OxygenRecordUpdate) AddDailyProduction(i int) *OxygenRecordUpdate {
	oru.mutation.AddDailyProduction(i)
	return oru
}

// SetMaxDailyProduction sets the "maxDailyProduction" field.
func (oru *OxygenRecordUpdate) SetMaxDailyProduction(i int) *OxygenRecordUpdate {
	oru.mutation.ResetMaxDailyProduction()
	oru.mutation.SetMaxDailyProduction(i)
	return oru
}

// AddMaxDailyProduction adds i to the "maxDailyProduction" field.
func (oru *OxygenRecordUpdate) AddMaxDailyProduction(i int) *OxygenRecordUpdate {
	oru.mutation.AddMaxDailyProduction(i)
	return oru
}

// SetDailyConsumption sets the "dailyConsumption" field.
func (oru *OxygenRecordUpdate) SetDailyConsumption(i int) *OxygenRecordUpdate {
	oru.mutation.ResetDailyConsumption()
	oru.mutation.SetDailyConsumption(i)
	return oru
}

// AddDailyConsumption adds i to the "dailyConsumption" field.
func (oru *OxygenRecordUpdate) AddDailyConsumption(i int) *OxygenRecordUpdate {
	oru.mutation.AddDailyConsumption(i)
	return oru
}

// SetMainSourceKind sets the "mainSourceKind" field.
func (oru *OxygenRecordUpdate) SetMainSourceKind(s string) *OxygenRecordUpdate {
	oru.mutation.SetMainSourceKind(s)
	return oru
}

// AddPlaceIDs adds the "places" edge to the Place entity by IDs.
func (oru *OxygenRecordUpdate) AddPlaceIDs(ids ...int) *OxygenRecordUpdate {
	oru.mutation.AddPlaceIDs(ids...)
	return oru
}

// AddPlaces adds the "places" edges to the Place entity.
func (oru *OxygenRecordUpdate) AddPlaces(p ...*Place) *OxygenRecordUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oru.AddPlaceIDs(ids...)
}

// Mutation returns the OxygenRecordMutation object of the builder.
func (oru *OxygenRecordUpdate) Mutation() *OxygenRecordMutation {
	return oru.mutation
}

// ClearPlaces clears all "places" edges to the Place entity.
func (oru *OxygenRecordUpdate) ClearPlaces() *OxygenRecordUpdate {
	oru.mutation.ClearPlaces()
	return oru
}

// RemovePlaceIDs removes the "places" edge to Place entities by IDs.
func (oru *OxygenRecordUpdate) RemovePlaceIDs(ids ...int) *OxygenRecordUpdate {
	oru.mutation.RemovePlaceIDs(ids...)
	return oru
}

// RemovePlaces removes "places" edges to Place entities.
func (oru *OxygenRecordUpdate) RemovePlaces(p ...*Place) *OxygenRecordUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oru.RemovePlaceIDs(ids...)
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
	if value, ok := oru.mutation.ReportedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oxygenrecord.FieldReportedDate,
		})
	}
	if value, ok := oru.mutation.CollectedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oxygenrecord.FieldCollectedDate,
		})
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
	if value, ok := oru.mutation.DailyProduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyProduction,
		})
	}
	if value, ok := oru.mutation.AddedDailyProduction(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyProduction,
		})
	}
	if value, ok := oru.mutation.MaxDailyProduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldMaxDailyProduction,
		})
	}
	if value, ok := oru.mutation.AddedMaxDailyProduction(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldMaxDailyProduction,
		})
	}
	if value, ok := oru.mutation.DailyConsumption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyConsumption,
		})
	}
	if value, ok := oru.mutation.AddedDailyConsumption(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyConsumption,
		})
	}
	if value, ok := oru.mutation.MainSourceKind(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oxygenrecord.FieldMainSourceKind,
		})
	}
	if oru.mutation.PlacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   oxygenrecord.PlacesTable,
			Columns: oxygenrecord.PlacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: place.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oru.mutation.RemovedPlacesIDs(); len(nodes) > 0 && !oru.mutation.PlacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   oxygenrecord.PlacesTable,
			Columns: oxygenrecord.PlacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: place.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oru.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   oxygenrecord.PlacesTable,
			Columns: oxygenrecord.PlacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: place.FieldID,
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

// SetReportedDate sets the "reportedDate" field.
func (oruo *OxygenRecordUpdateOne) SetReportedDate(t time.Time) *OxygenRecordUpdateOne {
	oruo.mutation.SetReportedDate(t)
	return oruo
}

// SetCollectedDate sets the "collectedDate" field.
func (oruo *OxygenRecordUpdateOne) SetCollectedDate(t time.Time) *OxygenRecordUpdateOne {
	oruo.mutation.SetCollectedDate(t)
	return oruo
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

// SetDailyProduction sets the "dailyProduction" field.
func (oruo *OxygenRecordUpdateOne) SetDailyProduction(i int) *OxygenRecordUpdateOne {
	oruo.mutation.ResetDailyProduction()
	oruo.mutation.SetDailyProduction(i)
	return oruo
}

// AddDailyProduction adds i to the "dailyProduction" field.
func (oruo *OxygenRecordUpdateOne) AddDailyProduction(i int) *OxygenRecordUpdateOne {
	oruo.mutation.AddDailyProduction(i)
	return oruo
}

// SetMaxDailyProduction sets the "maxDailyProduction" field.
func (oruo *OxygenRecordUpdateOne) SetMaxDailyProduction(i int) *OxygenRecordUpdateOne {
	oruo.mutation.ResetMaxDailyProduction()
	oruo.mutation.SetMaxDailyProduction(i)
	return oruo
}

// AddMaxDailyProduction adds i to the "maxDailyProduction" field.
func (oruo *OxygenRecordUpdateOne) AddMaxDailyProduction(i int) *OxygenRecordUpdateOne {
	oruo.mutation.AddMaxDailyProduction(i)
	return oruo
}

// SetDailyConsumption sets the "dailyConsumption" field.
func (oruo *OxygenRecordUpdateOne) SetDailyConsumption(i int) *OxygenRecordUpdateOne {
	oruo.mutation.ResetDailyConsumption()
	oruo.mutation.SetDailyConsumption(i)
	return oruo
}

// AddDailyConsumption adds i to the "dailyConsumption" field.
func (oruo *OxygenRecordUpdateOne) AddDailyConsumption(i int) *OxygenRecordUpdateOne {
	oruo.mutation.AddDailyConsumption(i)
	return oruo
}

// SetMainSourceKind sets the "mainSourceKind" field.
func (oruo *OxygenRecordUpdateOne) SetMainSourceKind(s string) *OxygenRecordUpdateOne {
	oruo.mutation.SetMainSourceKind(s)
	return oruo
}

// AddPlaceIDs adds the "places" edge to the Place entity by IDs.
func (oruo *OxygenRecordUpdateOne) AddPlaceIDs(ids ...int) *OxygenRecordUpdateOne {
	oruo.mutation.AddPlaceIDs(ids...)
	return oruo
}

// AddPlaces adds the "places" edges to the Place entity.
func (oruo *OxygenRecordUpdateOne) AddPlaces(p ...*Place) *OxygenRecordUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oruo.AddPlaceIDs(ids...)
}

// Mutation returns the OxygenRecordMutation object of the builder.
func (oruo *OxygenRecordUpdateOne) Mutation() *OxygenRecordMutation {
	return oruo.mutation
}

// ClearPlaces clears all "places" edges to the Place entity.
func (oruo *OxygenRecordUpdateOne) ClearPlaces() *OxygenRecordUpdateOne {
	oruo.mutation.ClearPlaces()
	return oruo
}

// RemovePlaceIDs removes the "places" edge to Place entities by IDs.
func (oruo *OxygenRecordUpdateOne) RemovePlaceIDs(ids ...int) *OxygenRecordUpdateOne {
	oruo.mutation.RemovePlaceIDs(ids...)
	return oruo
}

// RemovePlaces removes "places" edges to Place entities.
func (oruo *OxygenRecordUpdateOne) RemovePlaces(p ...*Place) *OxygenRecordUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oruo.RemovePlaceIDs(ids...)
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
	if value, ok := oruo.mutation.ReportedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oxygenrecord.FieldReportedDate,
		})
	}
	if value, ok := oruo.mutation.CollectedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oxygenrecord.FieldCollectedDate,
		})
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
	if value, ok := oruo.mutation.DailyProduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyProduction,
		})
	}
	if value, ok := oruo.mutation.AddedDailyProduction(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyProduction,
		})
	}
	if value, ok := oruo.mutation.MaxDailyProduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldMaxDailyProduction,
		})
	}
	if value, ok := oruo.mutation.AddedMaxDailyProduction(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldMaxDailyProduction,
		})
	}
	if value, ok := oruo.mutation.DailyConsumption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyConsumption,
		})
	}
	if value, ok := oruo.mutation.AddedDailyConsumption(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oxygenrecord.FieldDailyConsumption,
		})
	}
	if value, ok := oruo.mutation.MainSourceKind(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oxygenrecord.FieldMainSourceKind,
		})
	}
	if oruo.mutation.PlacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   oxygenrecord.PlacesTable,
			Columns: oxygenrecord.PlacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: place.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruo.mutation.RemovedPlacesIDs(); len(nodes) > 0 && !oruo.mutation.PlacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   oxygenrecord.PlacesTable,
			Columns: oxygenrecord.PlacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: place.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruo.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   oxygenrecord.PlacesTable,
			Columns: oxygenrecord.PlacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: place.FieldID,
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
