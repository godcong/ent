// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/traversal/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Pets holds the value of the pets edge.
	Pets []*Pet
	// Friends holds the value of the friends edge.
	Friends []*User
	// Groups holds the value of the groups edge.
	Groups []*Group
	// Manage holds the value of the manage edge.
	Manage []*Group
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PetsWithError returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PetsWithError() ([]*Pet, error) {
	if e.loadedTypes[0] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// FriendsWithError returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsWithError() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// GroupsWithError returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsWithError() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// ManageWithError returns the Manage value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ManageWithError() ([]*Group, error) {
	if e.loadedTypes[3] {
		return e.Manage, nil
	}
	return nil, &NotLoadedError{edge: "manage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // age
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[0])
	} else if value.Valid {
		u.Age = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		u.Name = value.String
	}
	return nil
}

// QueryPets queries the pets edge of the User.
func (u *User) QueryPets() *PetQuery {
	return (&UserClient{u.config}).QueryPets(u)
}

// QueryFriends queries the friends edge of the User.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{u.config}).QueryFriends(u)
}

// QueryGroups queries the groups edge of the User.
func (u *User) QueryGroups() *GroupQuery {
	return (&UserClient{u.config}).QueryGroups(u)
}

// QueryManage queries the manage edge of the User.
func (u *User) QueryManage() *GroupQuery {
	return (&UserClient{u.config}).QueryManage(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
