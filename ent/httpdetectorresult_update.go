// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/schema"
)

// HTTPDetectorResultUpdate is the builder for updating HTTPDetectorResult entities.
type HTTPDetectorResultUpdate struct {
	config
	hooks    []Hook
	mutation *HTTPDetectorResultMutation
}

// Where adds a new predicate for the HTTPDetectorResultUpdate builder.
func (hdru *HTTPDetectorResultUpdate) Where(ps ...predicate.HTTPDetectorResult) *HTTPDetectorResultUpdate {
	hdru.mutation.predicates = append(hdru.mutation.predicates, ps...)
	return hdru
}

// SetTask sets the "task" field.
func (hdru *HTTPDetectorResultUpdate) SetTask(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.ResetTask()
	hdru.mutation.SetTask(i)
	return hdru
}

// AddTask adds i to the "task" field.
func (hdru *HTTPDetectorResultUpdate) AddTask(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.AddTask(i)
	return hdru
}

// SetResult sets the "result" field.
func (hdru *HTTPDetectorResultUpdate) SetResult(i int8) *HTTPDetectorResultUpdate {
	hdru.mutation.ResetResult()
	hdru.mutation.SetResult(i)
	return hdru
}

// AddResult adds i to the "result" field.
func (hdru *HTTPDetectorResultUpdate) AddResult(i int8) *HTTPDetectorResultUpdate {
	hdru.mutation.AddResult(i)
	return hdru
}

// SetMaxDuration sets the "maxDuration" field.
func (hdru *HTTPDetectorResultUpdate) SetMaxDuration(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.ResetMaxDuration()
	hdru.mutation.SetMaxDuration(i)
	return hdru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (hdru *HTTPDetectorResultUpdate) AddMaxDuration(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.AddMaxDuration(i)
	return hdru
}

// SetMessages sets the "messages" field.
func (hdru *HTTPDetectorResultUpdate) SetMessages(s []string) *HTTPDetectorResultUpdate {
	hdru.mutation.SetMessages(s)
	return hdru
}

// SetURL sets the "url" field.
func (hdru *HTTPDetectorResultUpdate) SetURL(s string) *HTTPDetectorResultUpdate {
	hdru.mutation.SetURL(s)
	return hdru
}

// SetResults sets the "results" field.
func (hdru *HTTPDetectorResultUpdate) SetResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultUpdate {
	hdru.mutation.SetResults(sdsr)
	return hdru
}

// Mutation returns the HTTPDetectorResultMutation object of the builder.
func (hdru *HTTPDetectorResultUpdate) Mutation() *HTTPDetectorResultMutation {
	return hdru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hdru *HTTPDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	hdru.defaults()
	if len(hdru.hooks) == 0 {
		if err = hdru.check(); err != nil {
			return 0, err
		}
		affected, err = hdru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hdru.check(); err != nil {
				return 0, err
			}
			hdru.mutation = mutation
			affected, err = hdru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hdru.hooks) - 1; i >= 0; i-- {
			mut = hdru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hdru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hdru *HTTPDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := hdru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hdru *HTTPDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := hdru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdru *HTTPDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := hdru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdru *HTTPDetectorResultUpdate) defaults() {
	if _, ok := hdru.mutation.UpdatedAt(); !ok {
		v := httpdetectorresult.UpdateDefaultUpdatedAt()
		hdru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdru *HTTPDetectorResultUpdate) check() error {
	if v, ok := hdru.mutation.Result(); ok {
		if err := httpdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	return nil
}

func (hdru *HTTPDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httpdetectorresult.Table,
			Columns: httpdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httpdetectorresult.FieldID,
			},
		},
	}
	if ps := hdru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hdru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httpdetectorresult.FieldUpdatedAt,
		})
	}
	if value, ok := hdru.mutation.Task(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldTask,
		})
	}
	if value, ok := hdru.mutation.AddedTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldTask,
		})
	}
	if value, ok := hdru.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: httpdetectorresult.FieldResult,
		})
	}
	if value, ok := hdru.mutation.AddedResult(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: httpdetectorresult.FieldResult,
		})
	}
	if value, ok := hdru.mutation.MaxDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := hdru.mutation.AddedMaxDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := hdru.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httpdetectorresult.FieldMessages,
		})
	}
	if value, ok := hdru.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetectorresult.FieldURL,
		})
	}
	if value, ok := hdru.mutation.Results(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httpdetectorresult.FieldResults,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hdru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httpdetectorresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HTTPDetectorResultUpdateOne is the builder for updating a single HTTPDetectorResult entity.
type HTTPDetectorResultUpdateOne struct {
	config
	hooks    []Hook
	mutation *HTTPDetectorResultMutation
}

// SetTask sets the "task" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetTask(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.ResetTask()
	hdruo.mutation.SetTask(i)
	return hdruo
}

// AddTask adds i to the "task" field.
func (hdruo *HTTPDetectorResultUpdateOne) AddTask(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AddTask(i)
	return hdruo
}

// SetResult sets the "result" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetResult(i int8) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.ResetResult()
	hdruo.mutation.SetResult(i)
	return hdruo
}

// AddResult adds i to the "result" field.
func (hdruo *HTTPDetectorResultUpdateOne) AddResult(i int8) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AddResult(i)
	return hdruo
}

// SetMaxDuration sets the "maxDuration" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetMaxDuration(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.ResetMaxDuration()
	hdruo.mutation.SetMaxDuration(i)
	return hdruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (hdruo *HTTPDetectorResultUpdateOne) AddMaxDuration(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AddMaxDuration(i)
	return hdruo
}

// SetMessages sets the "messages" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetMessages(s []string) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.SetMessages(s)
	return hdruo
}

// SetURL sets the "url" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetURL(s string) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.SetURL(s)
	return hdruo
}

// SetResults sets the "results" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.SetResults(sdsr)
	return hdruo
}

// Mutation returns the HTTPDetectorResultMutation object of the builder.
func (hdruo *HTTPDetectorResultUpdateOne) Mutation() *HTTPDetectorResultMutation {
	return hdruo.mutation
}

// Save executes the query and returns the updated HTTPDetectorResult entity.
func (hdruo *HTTPDetectorResultUpdateOne) Save(ctx context.Context) (*HTTPDetectorResult, error) {
	var (
		err  error
		node *HTTPDetectorResult
	)
	hdruo.defaults()
	if len(hdruo.hooks) == 0 {
		if err = hdruo.check(); err != nil {
			return nil, err
		}
		node, err = hdruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hdruo.check(); err != nil {
				return nil, err
			}
			hdruo.mutation = mutation
			node, err = hdruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hdruo.hooks) - 1; i >= 0; i-- {
			mut = hdruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hdruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (hdruo *HTTPDetectorResultUpdateOne) SaveX(ctx context.Context) *HTTPDetectorResult {
	node, err := hdruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hdruo *HTTPDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := hdruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdruo *HTTPDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := hdruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdruo *HTTPDetectorResultUpdateOne) defaults() {
	if _, ok := hdruo.mutation.UpdatedAt(); !ok {
		v := httpdetectorresult.UpdateDefaultUpdatedAt()
		hdruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdruo *HTTPDetectorResultUpdateOne) check() error {
	if v, ok := hdruo.mutation.Result(); ok {
		if err := httpdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	return nil
}

func (hdruo *HTTPDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *HTTPDetectorResult, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httpdetectorresult.Table,
			Columns: httpdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httpdetectorresult.FieldID,
			},
		},
	}
	id, ok := hdruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing HTTPDetectorResult.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := hdruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httpdetectorresult.FieldUpdatedAt,
		})
	}
	if value, ok := hdruo.mutation.Task(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldTask,
		})
	}
	if value, ok := hdruo.mutation.AddedTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldTask,
		})
	}
	if value, ok := hdruo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: httpdetectorresult.FieldResult,
		})
	}
	if value, ok := hdruo.mutation.AddedResult(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: httpdetectorresult.FieldResult,
		})
	}
	if value, ok := hdruo.mutation.MaxDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := hdruo.mutation.AddedMaxDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httpdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := hdruo.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httpdetectorresult.FieldMessages,
		})
	}
	if value, ok := hdruo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httpdetectorresult.FieldURL,
		})
	}
	if value, ok := hdruo.mutation.Results(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httpdetectorresult.FieldResults,
		})
	}
	_node = &HTTPDetectorResult{config: hdruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hdruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httpdetectorresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
