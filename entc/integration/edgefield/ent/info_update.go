// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/entc/integration/edgefield/ent/info"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// InfoUpdate is the builder for updating Info entities.
type InfoUpdate struct {
	config
	hooks    []Hook
	mutation *InfoMutation
}

// Where appends a list predicates to the InfoUpdate builder.
func (_u *InfoUpdate) Where(ps ...predicate.Info) *InfoUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetContent sets the "content" field.
func (_u *InfoUpdate) SetContent(v json.RawMessage) *InfoUpdate {
	_u.mutation.SetContent(v)
	return _u
}

// AppendContent appends value to the "content" field.
func (_u *InfoUpdate) AppendContent(v json.RawMessage) *InfoUpdate {
	_u.mutation.AppendContent(v)
	return _u
}

// SetUserID sets the "user" edge to the User entity by ID.
func (_u *InfoUpdate) SetUserID(id int) *InfoUpdate {
	_u.mutation.SetUserID(id)
	return _u
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (_u *InfoUpdate) SetNillableUserID(id *int) *InfoUpdate {
	if id != nil {
		_u = _u.SetUserID(*id)
	}
	return _u
}

// SetUser sets the "user" edge to the User entity.
func (_u *InfoUpdate) SetUser(v *User) *InfoUpdate {
	return _u.SetUserID(v.ID)
}

// Mutation returns the InfoMutation object of the builder.
func (_u *InfoUpdate) Mutation() *InfoMutation {
	return _u.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (_u *InfoUpdate) ClearUser() *InfoUpdate {
	_u.mutation.ClearUser()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *InfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *InfoUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *InfoUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *InfoUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *InfoUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(info.Table, info.Columns, sqlgraph.NewFieldSpec(info.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Content(); ok {
		_spec.SetField(info.FieldContent, field.TypeJSON, value)
	}
	if value, ok := _u.mutation.AppendedContent(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, info.FieldContent, value)
		})
	}
	if _u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{info.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// InfoUpdateOne is the builder for updating a single Info entity.
type InfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InfoMutation
}

// SetContent sets the "content" field.
func (_u *InfoUpdateOne) SetContent(v json.RawMessage) *InfoUpdateOne {
	_u.mutation.SetContent(v)
	return _u
}

// AppendContent appends value to the "content" field.
func (_u *InfoUpdateOne) AppendContent(v json.RawMessage) *InfoUpdateOne {
	_u.mutation.AppendContent(v)
	return _u
}

// SetUserID sets the "user" edge to the User entity by ID.
func (_u *InfoUpdateOne) SetUserID(id int) *InfoUpdateOne {
	_u.mutation.SetUserID(id)
	return _u
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (_u *InfoUpdateOne) SetNillableUserID(id *int) *InfoUpdateOne {
	if id != nil {
		_u = _u.SetUserID(*id)
	}
	return _u
}

// SetUser sets the "user" edge to the User entity.
func (_u *InfoUpdateOne) SetUser(v *User) *InfoUpdateOne {
	return _u.SetUserID(v.ID)
}

// Mutation returns the InfoMutation object of the builder.
func (_u *InfoUpdateOne) Mutation() *InfoMutation {
	return _u.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (_u *InfoUpdateOne) ClearUser() *InfoUpdateOne {
	_u.mutation.ClearUser()
	return _u
}

// Where appends a list predicates to the InfoUpdate builder.
func (_u *InfoUpdateOne) Where(ps ...predicate.Info) *InfoUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *InfoUpdateOne) Select(field string, fields ...string) *InfoUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Info entity.
func (_u *InfoUpdateOne) Save(ctx context.Context) (*Info, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *InfoUpdateOne) SaveX(ctx context.Context) *Info {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *InfoUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *InfoUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *InfoUpdateOne) sqlSave(ctx context.Context) (_node *Info, err error) {
	_spec := sqlgraph.NewUpdateSpec(info.Table, info.Columns, sqlgraph.NewFieldSpec(info.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Info.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, info.FieldID)
		for _, f := range fields {
			if !info.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != info.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Content(); ok {
		_spec.SetField(info.FieldContent, field.TypeJSON, value)
	}
	if value, ok := _u.mutation.AppendedContent(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, info.FieldContent, value)
		})
	}
	if _u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Info{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{info.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
