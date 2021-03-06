// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet/app/ent/doctor"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DoctorCreate is the builder for creating a Doctor entity.
type DoctorCreate struct {
	config
	mutation *DoctorMutation
	hooks    []Hook
}

// SetName sets the name field.
func (dc *DoctorCreate) SetName(i int) *DoctorCreate {
	dc.mutation.SetName(i)
	return dc
}

// Mutation returns the DoctorMutation object of the builder.
func (dc *DoctorCreate) Mutation() *DoctorMutation {
	return dc.mutation
}

// Save creates the Doctor in the database.
func (dc *DoctorCreate) Save(ctx context.Context) (*Doctor, error) {
	if _, ok := dc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := doctor.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Doctor
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DoctorCreate) SaveX(ctx context.Context) *Doctor {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DoctorCreate) sqlSave(ctx context.Context) (*Doctor, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DoctorCreate) createSpec() (*Doctor, *sqlgraph.CreateSpec) {
	var (
		d     = &Doctor{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doctor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldName,
		})
		d.Name = value
	}
	return d, _spec
}
