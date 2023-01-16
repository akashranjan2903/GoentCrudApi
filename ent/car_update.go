// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gocrud/ent/car"
	"github.com/gocrud/ent/predicate"
	"github.com/gocrud/ent/user"
	"github.com/google/uuid"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CarUpdate) SetUpdatedAt(t time.Time) *CarUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CarUpdate) SetDeletedAt(t time.Time) *CarUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CarUpdate) SetNillableDeletedAt(t *time.Time) *CarUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cu *CarUpdate) ClearDeletedAt() *CarUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetCreatedBy sets the "created_by" field.
func (cu *CarUpdate) SetCreatedBy(u uuid.UUID) *CarUpdate {
	cu.mutation.SetCreatedBy(u)
	return cu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cu *CarUpdate) SetNillableCreatedBy(u *uuid.UUID) *CarUpdate {
	if u != nil {
		cu.SetCreatedBy(*u)
	}
	return cu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (cu *CarUpdate) ClearCreatedBy() *CarUpdate {
	cu.mutation.ClearCreatedBy()
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *CarUpdate) SetUpdatedBy(u uuid.UUID) *CarUpdate {
	cu.mutation.SetUpdatedBy(u)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *CarUpdate) SetNillableUpdatedBy(u *uuid.UUID) *CarUpdate {
	if u != nil {
		cu.SetUpdatedBy(*u)
	}
	return cu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cu *CarUpdate) ClearUpdatedBy() *CarUpdate {
	cu.mutation.ClearUpdatedBy()
	return cu
}

// SetDeletedBy sets the "deleted_by" field.
func (cu *CarUpdate) SetDeletedBy(u uuid.UUID) *CarUpdate {
	cu.mutation.SetDeletedBy(u)
	return cu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cu *CarUpdate) SetNillableDeletedBy(u *uuid.UUID) *CarUpdate {
	if u != nil {
		cu.SetDeletedBy(*u)
	}
	return cu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (cu *CarUpdate) ClearDeletedBy() *CarUpdate {
	cu.mutation.ClearDeletedBy()
	return cu
}

// SetModel sets the "model" field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.mutation.SetModel(s)
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CarUpdate) SetUserID(id uuid.UUID) *CarUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cu *CarUpdate) SetNillableUserID(id *uuid.UUID) *CarUpdate {
	if id != nil {
		cu = cu.SetUserID(*id)
	}
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CarUpdate) SetUser(u *User) *CarUpdate {
	return cu.SetUserID(u.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CarUpdate) ClearUser() *CarUpdate {
	cu.mutation.ClearUser()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, CarMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CarUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if car.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized car.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := car.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: car.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(car.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(car.FieldDeletedAt, field.TypeTime, value)
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.ClearField(car.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.CreatedBy(); ok {
		_spec.SetField(car.FieldCreatedBy, field.TypeUUID, value)
	}
	if cu.mutation.CreatedByCleared() {
		_spec.ClearField(car.FieldCreatedBy, field.TypeUUID)
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(car.FieldUpdatedBy, field.TypeUUID, value)
	}
	if cu.mutation.UpdatedByCleared() {
		_spec.ClearField(car.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := cu.mutation.DeletedBy(); ok {
		_spec.SetField(car.FieldDeletedBy, field.TypeUUID, value)
	}
	if cu.mutation.DeletedByCleared() {
		_spec.ClearField(car.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := cu.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   car.UserTable,
			Columns: []string{car.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   car.UserTable,
			Columns: []string{car.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CarUpdateOne) SetUpdatedAt(t time.Time) *CarUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CarUpdateOne) SetDeletedAt(t time.Time) *CarUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableDeletedAt(t *time.Time) *CarUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cuo *CarUpdateOne) ClearDeletedAt() *CarUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetCreatedBy sets the "created_by" field.
func (cuo *CarUpdateOne) SetCreatedBy(u uuid.UUID) *CarUpdateOne {
	cuo.mutation.SetCreatedBy(u)
	return cuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *CarUpdateOne {
	if u != nil {
		cuo.SetCreatedBy(*u)
	}
	return cuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (cuo *CarUpdateOne) ClearCreatedBy() *CarUpdateOne {
	cuo.mutation.ClearCreatedBy()
	return cuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *CarUpdateOne) SetUpdatedBy(u uuid.UUID) *CarUpdateOne {
	cuo.mutation.SetUpdatedBy(u)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *CarUpdateOne {
	if u != nil {
		cuo.SetUpdatedBy(*u)
	}
	return cuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cuo *CarUpdateOne) ClearUpdatedBy() *CarUpdateOne {
	cuo.mutation.ClearUpdatedBy()
	return cuo
}

// SetDeletedBy sets the "deleted_by" field.
func (cuo *CarUpdateOne) SetDeletedBy(u uuid.UUID) *CarUpdateOne {
	cuo.mutation.SetDeletedBy(u)
	return cuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableDeletedBy(u *uuid.UUID) *CarUpdateOne {
	if u != nil {
		cuo.SetDeletedBy(*u)
	}
	return cuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (cuo *CarUpdateOne) ClearDeletedBy() *CarUpdateOne {
	cuo.mutation.ClearDeletedBy()
	return cuo
}

// SetModel sets the "model" field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.mutation.SetModel(s)
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CarUpdateOne) SetUserID(id uuid.UUID) *CarUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableUserID(id *uuid.UUID) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetUserID(*id)
	}
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CarUpdateOne) SetUser(u *User) *CarUpdateOne {
	return cuo.SetUserID(u.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CarUpdateOne) ClearUser() *CarUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Car, CarMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CarUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if car.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized car.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := car.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: car.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(car.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(car.FieldDeletedAt, field.TypeTime, value)
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.ClearField(car.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.CreatedBy(); ok {
		_spec.SetField(car.FieldCreatedBy, field.TypeUUID, value)
	}
	if cuo.mutation.CreatedByCleared() {
		_spec.ClearField(car.FieldCreatedBy, field.TypeUUID)
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(car.FieldUpdatedBy, field.TypeUUID, value)
	}
	if cuo.mutation.UpdatedByCleared() {
		_spec.ClearField(car.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := cuo.mutation.DeletedBy(); ok {
		_spec.SetField(car.FieldDeletedBy, field.TypeUUID, value)
	}
	if cuo.mutation.DeletedByCleared() {
		_spec.ClearField(car.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := cuo.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   car.UserTable,
			Columns: []string{car.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   car.UserTable,
			Columns: []string{car.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Car{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
