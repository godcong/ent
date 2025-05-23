// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package versioned

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/versioned/predicate"
	"entgo.io/ent/entc/integration/migrate/versioned/user"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (_u *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetAge sets the "age" field.
func (_u *UserUpdate) SetAge(v int32) *UserUpdate {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(v)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdate) SetNillableAge(v *int32) *UserUpdate {
	if v != nil {
		_u.SetAge(*v)
	}
	return _u
}

// AddAge adds value to the "age" field.
func (_u *UserUpdate) AddAge(v int32) *UserUpdate {
	_u.mutation.AddAge(v)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdate) SetName(v string) *UserUpdate {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdate) SetNillableName(v *string) *UserUpdate {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetAddress sets the "address" field.
func (_u *UserUpdate) SetAddress(v string) *UserUpdate {
	_u.mutation.SetAddress(v)
	return _u
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (_u *UserUpdate) SetNillableAddress(v *string) *UserUpdate {
	if v != nil {
		_u.SetAddress(*v)
	}
	return _u
}

// ClearAddress clears the value of the "address" field.
func (_u *UserUpdate) ClearAddress() *UserUpdate {
	_u.mutation.ClearAddress()
	return _u
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdate) Mutation() *UserMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *UserUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *UserUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *UserUpdate) check() error {
	if v, ok := _u.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`versioned: validator failed for field "User.name": %w`, err)}
		}
	}
	return nil
}

func (_u *UserUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := _u.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if _u.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (_u *UserUpdateOne) SetAge(v int32) *UserUpdateOne {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(v)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableAge(v *int32) *UserUpdateOne {
	if v != nil {
		_u.SetAge(*v)
	}
	return _u
}

// AddAge adds value to the "age" field.
func (_u *UserUpdateOne) AddAge(v int32) *UserUpdateOne {
	_u.mutation.AddAge(v)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdateOne) SetName(v string) *UserUpdateOne {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableName(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetAddress sets the "address" field.
func (_u *UserUpdateOne) SetAddress(v string) *UserUpdateOne {
	_u.mutation.SetAddress(v)
	return _u
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableAddress(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetAddress(*v)
	}
	return _u
}

// ClearAddress clears the value of the "address" field.
func (_u *UserUpdateOne) ClearAddress() *UserUpdateOne {
	_u.mutation.ClearAddress()
	return _u
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdateOne) Mutation() *UserMutation {
	return _u.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (_u *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated User entity.
func (_u *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *UserUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *UserUpdateOne) check() error {
	if v, ok := _u.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`versioned: validator failed for field "User.name": %w`, err)}
		}
	}
	return nil
}

func (_u *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`versioned: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("versioned: invalid field %q for query", f)}
			}
			if f != user.FieldID {
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
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := _u.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if _u.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	_node = &User{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
