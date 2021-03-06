// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/schema"
)

// HTTPDetectorCreate is the builder for creating a HTTPDetector entity.
type HTTPDetectorCreate struct {
	config
	mutation *HTTPDetectorMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hdc *HTTPDetectorCreate) SetCreatedAt(t time.Time) *HTTPDetectorCreate {
	hdc.mutation.SetCreatedAt(t)
	return hdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hdc *HTTPDetectorCreate) SetNillableCreatedAt(t *time.Time) *HTTPDetectorCreate {
	if t != nil {
		hdc.SetCreatedAt(*t)
	}
	return hdc
}

// SetUpdatedAt sets the "updated_at" field.
func (hdc *HTTPDetectorCreate) SetUpdatedAt(t time.Time) *HTTPDetectorCreate {
	hdc.mutation.SetUpdatedAt(t)
	return hdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hdc *HTTPDetectorCreate) SetNillableUpdatedAt(t *time.Time) *HTTPDetectorCreate {
	if t != nil {
		hdc.SetUpdatedAt(*t)
	}
	return hdc
}

// SetStatus sets the "status" field.
func (hdc *HTTPDetectorCreate) SetStatus(s schema.Status) *HTTPDetectorCreate {
	hdc.mutation.SetStatus(s)
	return hdc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hdc *HTTPDetectorCreate) SetNillableStatus(s *schema.Status) *HTTPDetectorCreate {
	if s != nil {
		hdc.SetStatus(*s)
	}
	return hdc
}

// SetName sets the "name" field.
func (hdc *HTTPDetectorCreate) SetName(s string) *HTTPDetectorCreate {
	hdc.mutation.SetName(s)
	return hdc
}

// SetOwner sets the "owner" field.
func (hdc *HTTPDetectorCreate) SetOwner(s string) *HTTPDetectorCreate {
	hdc.mutation.SetOwner(s)
	return hdc
}

// SetDescription sets the "description" field.
func (hdc *HTTPDetectorCreate) SetDescription(s string) *HTTPDetectorCreate {
	hdc.mutation.SetDescription(s)
	return hdc
}

// SetReceivers sets the "receivers" field.
func (hdc *HTTPDetectorCreate) SetReceivers(s []string) *HTTPDetectorCreate {
	hdc.mutation.SetReceivers(s)
	return hdc
}

// SetTimeout sets the "timeout" field.
func (hdc *HTTPDetectorCreate) SetTimeout(s string) *HTTPDetectorCreate {
	hdc.mutation.SetTimeout(s)
	return hdc
}

// SetIps sets the "ips" field.
func (hdc *HTTPDetectorCreate) SetIps(s []string) *HTTPDetectorCreate {
	hdc.mutation.SetIps(s)
	return hdc
}

// SetURL sets the "url" field.
func (hdc *HTTPDetectorCreate) SetURL(s string) *HTTPDetectorCreate {
	hdc.mutation.SetURL(s)
	return hdc
}

// Mutation returns the HTTPDetectorMutation object of the builder.
func (hdc *HTTPDetectorCreate) Mutation() *HTTPDetectorMutation {
	return hdc.mutation
}

// Save creates the HTTPDetector in the database.
func (hdc *HTTPDetectorCreate) Save(ctx context.Context) (*HTTPDetector, error) {
	var (
		err  error
		node *HTTPDetector
	)
	hdc.defaults()
	if len(hdc.hooks) == 0 {
		if err = hdc.check(); err != nil {
			return nil, err
		}
		node, err = hdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hdc.check(); err != nil {
				return nil, err
			}
			hdc.mutation = mutation
			node, err = hdc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hdc.hooks) - 1; i >= 0; i-- {
			mut = hdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hdc *HTTPDetectorCreate) SaveX(ctx context.Context) *HTTPDetector {
	v, err := hdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (hdc *HTTPDetectorCreate) defaults() {
	if _, ok := hdc.mutation.CreatedAt(); !ok {
		v := httpdetector.DefaultCreatedAt()
		hdc.mutation.SetCreatedAt(v)
	}
	if _, ok := hdc.mutation.UpdatedAt(); !ok {
		v := httpdetector.DefaultUpdatedAt()
		hdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := hdc.mutation.Status(); !ok {
		v := httpdetector.DefaultStatus
		hdc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdc *HTTPDetectorCreate) check() error {
	if _, ok := hdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := hdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := hdc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := hdc.mutation.Status(); ok {
		if err := httpdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := hdc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := hdc.mutation.Name(); ok {
		if err := httpdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := hdc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("ent: missing required field \"owner\"")}
	}
	if v, ok := hdc.mutation.Owner(); ok {
		if err := httpdetector.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if _, ok := hdc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := hdc.mutation.Receivers(); !ok {
		return &ValidationError{Name: "receivers", err: errors.New("ent: missing required field \"receivers\"")}
	}
	if _, ok := hdc.mutation.Timeout(); !ok {
		return &ValidationError{Name: "timeout", err: errors.New("ent: missing required field \"timeout\"")}
	}
	if v, ok := hdc.mutation.Timeout(); ok {
		if err := httpdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	if _, ok := hdc.mutation.Ips(); !ok {
		return &ValidationError{Name: "ips", err: errors.New("ent: missing required field \"ips\"")}
	}
	if _, ok := hdc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if v, ok := hdc.mutation.URL(); ok {
		if err := httpdetector.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	return nil
}

func (hdc *HTTPDetectorCreate) sqlSave(ctx context.Context) (*HTTPDetector, error) {
	_node, _spec := hdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hdc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hdc *HTTPDetectorCreate) createSpec() (*HTTPDetector, *sqlgraph.CreateSpec) {
	var (
		_node = &HTTPDetector{config: hdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: httpdetector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httpdetector.FieldID,
			},
		}
	)
	if value, ok := hdc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httpdetector.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := hdc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httpdetector.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := hdc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: httpdetector.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := hdc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetector.FieldName,
		})
		_node.Name = value
	}
	if value, ok := hdc.mutation.Owner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetector.FieldOwner,
		})
		_node.Owner = value
	}
	if value, ok := hdc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetector.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := hdc.mutation.Receivers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httpdetector.FieldReceivers,
		})
		_node.Receivers = value
	}
	if value, ok := hdc.mutation.Timeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetector.FieldTimeout,
		})
		_node.Timeout = value
	}
	if value, ok := hdc.mutation.Ips(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httpdetector.FieldIps,
		})
		_node.Ips = value
	}
	if value, ok := hdc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetector.FieldURL,
		})
		_node.URL = value
	}
	return _node, _spec
}

// HTTPDetectorCreateBulk is the builder for creating many HTTPDetector entities in bulk.
type HTTPDetectorCreateBulk struct {
	config
	builders []*HTTPDetectorCreate
}

// Save creates the HTTPDetector entities in the database.
func (hdcb *HTTPDetectorCreateBulk) Save(ctx context.Context) ([]*HTTPDetector, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hdcb.builders))
	nodes := make([]*HTTPDetector, len(hdcb.builders))
	mutators := make([]Mutator, len(hdcb.builders))
	for i := range hdcb.builders {
		func(i int, root context.Context) {
			builder := hdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HTTPDetectorMutation)
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
					_, err = mutators[i+1].Mutate(root, hdcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hdcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hdcb *HTTPDetectorCreateBulk) SaveX(ctx context.Context) []*HTTPDetector {
	v, err := hdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
