// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"opencensus/core/ent/district"
	"opencensus/core/ent/occurency"
	"opencensus/core/ent/province"
	"opencensus/core/ent/region"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OccurencyCreate is the builder for creating a Occurency entity.
type OccurencyCreate struct {
	config
	mutation *OccurencyMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (oc *OccurencyCreate) SetUUID(s string) *OccurencyCreate {
	oc.mutation.SetUUID(s)
	return oc
}

// SetReportedRecord sets the "reportedRecord" field.
func (oc *OccurencyCreate) SetReportedRecord(t time.Time) *OccurencyCreate {
	oc.mutation.SetReportedRecord(t)
	return oc
}

// SetResultDate sets the "resultDate" field.
func (oc *OccurencyCreate) SetResultDate(t time.Time) *OccurencyCreate {
	oc.mutation.SetResultDate(t)
	return oc
}

// SetBiologicalSex sets the "biologicalSex" field.
func (oc *OccurencyCreate) SetBiologicalSex(s string) *OccurencyCreate {
	oc.mutation.SetBiologicalSex(s)
	return oc
}

// SetAge sets the "age" field.
func (oc *OccurencyCreate) SetAge(i int) *OccurencyCreate {
	oc.mutation.SetAge(i)
	return oc
}

// SetRegionID sets the "region" edge to the Region entity by ID.
func (oc *OccurencyCreate) SetRegionID(id int) *OccurencyCreate {
	oc.mutation.SetRegionID(id)
	return oc
}

// SetNillableRegionID sets the "region" edge to the Region entity by ID if the given value is not nil.
func (oc *OccurencyCreate) SetNillableRegionID(id *int) *OccurencyCreate {
	if id != nil {
		oc = oc.SetRegionID(*id)
	}
	return oc
}

// SetRegion sets the "region" edge to the Region entity.
func (oc *OccurencyCreate) SetRegion(r *Region) *OccurencyCreate {
	return oc.SetRegionID(r.ID)
}

// SetProvinceID sets the "province" edge to the Province entity by ID.
func (oc *OccurencyCreate) SetProvinceID(id int) *OccurencyCreate {
	oc.mutation.SetProvinceID(id)
	return oc
}

// SetNillableProvinceID sets the "province" edge to the Province entity by ID if the given value is not nil.
func (oc *OccurencyCreate) SetNillableProvinceID(id *int) *OccurencyCreate {
	if id != nil {
		oc = oc.SetProvinceID(*id)
	}
	return oc
}

// SetProvince sets the "province" edge to the Province entity.
func (oc *OccurencyCreate) SetProvince(p *Province) *OccurencyCreate {
	return oc.SetProvinceID(p.ID)
}

// SetDistrictID sets the "district" edge to the District entity by ID.
func (oc *OccurencyCreate) SetDistrictID(id int) *OccurencyCreate {
	oc.mutation.SetDistrictID(id)
	return oc
}

// SetNillableDistrictID sets the "district" edge to the District entity by ID if the given value is not nil.
func (oc *OccurencyCreate) SetNillableDistrictID(id *int) *OccurencyCreate {
	if id != nil {
		oc = oc.SetDistrictID(*id)
	}
	return oc
}

// SetDistrict sets the "district" edge to the District entity.
func (oc *OccurencyCreate) SetDistrict(d *District) *OccurencyCreate {
	return oc.SetDistrictID(d.ID)
}

// Mutation returns the OccurencyMutation object of the builder.
func (oc *OccurencyCreate) Mutation() *OccurencyMutation {
	return oc.mutation
}

// Save creates the Occurency in the database.
func (oc *OccurencyCreate) Save(ctx context.Context) (*Occurency, error) {
	var (
		err  error
		node *Occurency
	)
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OccurencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OccurencyCreate) SaveX(ctx context.Context) *Occurency {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (oc *OccurencyCreate) check() error {
	if _, ok := oc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := oc.mutation.ReportedRecord(); !ok {
		return &ValidationError{Name: "reportedRecord", err: errors.New("ent: missing required field \"reportedRecord\"")}
	}
	if _, ok := oc.mutation.ResultDate(); !ok {
		return &ValidationError{Name: "resultDate", err: errors.New("ent: missing required field \"resultDate\"")}
	}
	if _, ok := oc.mutation.BiologicalSex(); !ok {
		return &ValidationError{Name: "biologicalSex", err: errors.New("ent: missing required field \"biologicalSex\"")}
	}
	if _, ok := oc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	return nil
}

func (oc *OccurencyCreate) sqlSave(ctx context.Context) (*Occurency, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OccurencyCreate) createSpec() (*Occurency, *sqlgraph.CreateSpec) {
	var (
		_node = &Occurency{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: occurency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: occurency.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occurency.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := oc.mutation.ReportedRecord(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: occurency.FieldReportedRecord,
		})
		_node.ReportedRecord = value
	}
	if value, ok := oc.mutation.ResultDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: occurency.FieldResultDate,
		})
		_node.ResultDate = value
	}
	if value, ok := oc.mutation.BiologicalSex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occurency.FieldBiologicalSex,
		})
		_node.BiologicalSex = value
	}
	if value, ok := oc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: occurency.FieldAge,
		})
		_node.Age = value
	}
	if nodes := oc.mutation.RegionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurency.RegionTable,
			Columns: []string{occurency.RegionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: region.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurency.ProvinceTable,
			Columns: []string{occurency.ProvinceColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.DistrictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurency.DistrictTable,
			Columns: []string{occurency.DistrictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: district.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OccurencyCreateBulk is the builder for creating many Occurency entities in bulk.
type OccurencyCreateBulk struct {
	config
	builders []*OccurencyCreate
}

// Save creates the Occurency entities in the database.
func (ocb *OccurencyCreateBulk) Save(ctx context.Context) ([]*Occurency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Occurency, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OccurencyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OccurencyCreateBulk) SaveX(ctx context.Context) []*Occurency {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}