// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/http"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/schema"
)

// HTTPUpdate is the builder for updating HTTP entities.
type HTTPUpdate struct {
	config
	hooks    []Hook
	mutation *HTTPMutation
}

// Where adds a new predicate for the HTTPUpdate builder.
func (hu *HTTPUpdate) Where(ps ...predicate.HTTP) *HTTPUpdate {
	hu.mutation.predicates = append(hu.mutation.predicates, ps...)
	return hu
}

// SetStatus sets the "status" field.
func (hu *HTTPUpdate) SetStatus(s schema.Status) *HTTPUpdate {
	hu.mutation.ResetStatus()
	hu.mutation.SetStatus(s)
	return hu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hu *HTTPUpdate) SetNillableStatus(s *schema.Status) *HTTPUpdate {
	if s != nil {
		hu.SetStatus(*s)
	}
	return hu
}

// AddStatus adds s to the "status" field.
func (hu *HTTPUpdate) AddStatus(s schema.Status) *HTTPUpdate {
	hu.mutation.AddStatus(s)
	return hu
}

// SetName sets the "name" field.
func (hu *HTTPUpdate) SetName(s string) *HTTPUpdate {
	hu.mutation.SetName(s)
	return hu
}

// SetOwner sets the "owner" field.
func (hu *HTTPUpdate) SetOwner(s string) *HTTPUpdate {
	hu.mutation.SetOwner(s)
	return hu
}

// SetDescription sets the "description" field.
func (hu *HTTPUpdate) SetDescription(s string) *HTTPUpdate {
	hu.mutation.SetDescription(s)
	return hu
}

// SetReceivers sets the "receivers" field.
func (hu *HTTPUpdate) SetReceivers(s []string) *HTTPUpdate {
	hu.mutation.SetReceivers(s)
	return hu
}

// SetIps sets the "ips" field.
func (hu *HTTPUpdate) SetIps(s []string) *HTTPUpdate {
	hu.mutation.SetIps(s)
	return hu
}

// SetURL sets the "url" field.
func (hu *HTTPUpdate) SetURL(s string) *HTTPUpdate {
	hu.mutation.SetURL(s)
	return hu
}

// SetTimeout sets the "timeout" field.
func (hu *HTTPUpdate) SetTimeout(s string) *HTTPUpdate {
	hu.mutation.SetTimeout(s)
	return hu
}

// Mutation returns the HTTPMutation object of the builder.
func (hu *HTTPUpdate) Mutation() *HTTPMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HTTPUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	hu.defaults()
	if len(hu.hooks) == 0 {
		if err = hu.check(); err != nil {
			return 0, err
		}
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hu.check(); err != nil {
				return 0, err
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HTTPUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HTTPUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HTTPUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HTTPUpdate) defaults() {
	if _, ok := hu.mutation.UpdatedAt(); !ok {
		v := http.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HTTPUpdate) check() error {
	if v, ok := hu.mutation.Status(); ok {
		if err := http.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := hu.mutation.Name(); ok {
		if err := http.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := hu.mutation.Owner(); ok {
		if err := http.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if v, ok := hu.mutation.URL(); ok {
		if err := http.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if v, ok := hu.mutation.Timeout(); ok {
		if err := http.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	return nil
}

func (hu *HTTPUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   http.Table,
			Columns: http.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: http.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: http.FieldUpdatedAt,
		})
	}
	if value, ok := hu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: http.FieldStatus,
		})
	}
	if value, ok := hu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: http.FieldStatus,
		})
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldName,
		})
	}
	if value, ok := hu.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldOwner,
		})
	}
	if value, ok := hu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldDescription,
		})
	}
	if value, ok := hu.mutation.Receivers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: http.FieldReceivers,
		})
	}
	if value, ok := hu.mutation.Ips(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: http.FieldIps,
		})
	}
	if value, ok := hu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldURL,
		})
	}
	if value, ok := hu.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldTimeout,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{http.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HTTPUpdateOne is the builder for updating a single HTTP entity.
type HTTPUpdateOne struct {
	config
	hooks    []Hook
	mutation *HTTPMutation
}

// SetStatus sets the "status" field.
func (huo *HTTPUpdateOne) SetStatus(s schema.Status) *HTTPUpdateOne {
	huo.mutation.ResetStatus()
	huo.mutation.SetStatus(s)
	return huo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (huo *HTTPUpdateOne) SetNillableStatus(s *schema.Status) *HTTPUpdateOne {
	if s != nil {
		huo.SetStatus(*s)
	}
	return huo
}

// AddStatus adds s to the "status" field.
func (huo *HTTPUpdateOne) AddStatus(s schema.Status) *HTTPUpdateOne {
	huo.mutation.AddStatus(s)
	return huo
}

// SetName sets the "name" field.
func (huo *HTTPUpdateOne) SetName(s string) *HTTPUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// SetOwner sets the "owner" field.
func (huo *HTTPUpdateOne) SetOwner(s string) *HTTPUpdateOne {
	huo.mutation.SetOwner(s)
	return huo
}

// SetDescription sets the "description" field.
func (huo *HTTPUpdateOne) SetDescription(s string) *HTTPUpdateOne {
	huo.mutation.SetDescription(s)
	return huo
}

// SetReceivers sets the "receivers" field.
func (huo *HTTPUpdateOne) SetReceivers(s []string) *HTTPUpdateOne {
	huo.mutation.SetReceivers(s)
	return huo
}

// SetIps sets the "ips" field.
func (huo *HTTPUpdateOne) SetIps(s []string) *HTTPUpdateOne {
	huo.mutation.SetIps(s)
	return huo
}

// SetURL sets the "url" field.
func (huo *HTTPUpdateOne) SetURL(s string) *HTTPUpdateOne {
	huo.mutation.SetURL(s)
	return huo
}

// SetTimeout sets the "timeout" field.
func (huo *HTTPUpdateOne) SetTimeout(s string) *HTTPUpdateOne {
	huo.mutation.SetTimeout(s)
	return huo
}

// Mutation returns the HTTPMutation object of the builder.
func (huo *HTTPUpdateOne) Mutation() *HTTPMutation {
	return huo.mutation
}

// Save executes the query and returns the updated HTTP entity.
func (huo *HTTPUpdateOne) Save(ctx context.Context) (*HTTP, error) {
	var (
		err  error
		node *HTTP
	)
	huo.defaults()
	if len(huo.hooks) == 0 {
		if err = huo.check(); err != nil {
			return nil, err
		}
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = huo.check(); err != nil {
				return nil, err
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HTTPUpdateOne) SaveX(ctx context.Context) *HTTP {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HTTPUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HTTPUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HTTPUpdateOne) defaults() {
	if _, ok := huo.mutation.UpdatedAt(); !ok {
		v := http.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HTTPUpdateOne) check() error {
	if v, ok := huo.mutation.Status(); ok {
		if err := http.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := huo.mutation.Name(); ok {
		if err := http.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := huo.mutation.Owner(); ok {
		if err := http.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if v, ok := huo.mutation.URL(); ok {
		if err := http.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if v, ok := huo.mutation.Timeout(); ok {
		if err := http.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	return nil
}

func (huo *HTTPUpdateOne) sqlSave(ctx context.Context) (_node *HTTP, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   http.Table,
			Columns: http.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: http.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing HTTP.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: http.FieldUpdatedAt,
		})
	}
	if value, ok := huo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: http.FieldStatus,
		})
	}
	if value, ok := huo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: http.FieldStatus,
		})
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldName,
		})
	}
	if value, ok := huo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldOwner,
		})
	}
	if value, ok := huo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldDescription,
		})
	}
	if value, ok := huo.mutation.Receivers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: http.FieldReceivers,
		})
	}
	if value, ok := huo.mutation.Ips(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: http.FieldIps,
		})
	}
	if value, ok := huo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldURL,
		})
	}
	if value, ok := huo.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: http.FieldTimeout,
		})
	}
	_node = &HTTP{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{http.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}