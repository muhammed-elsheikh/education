// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"education/ent/upload"
	"education/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UploadCreate is the builder for creating a Upload entity.
type UploadCreate struct {
	config
	mutation *UploadMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uc *UploadCreate) SetUserID(i int) *UploadCreate {
	uc.mutation.SetUserID(i)
	return uc
}

// SetName sets the "name" field.
func (uc *UploadCreate) SetName(s string) *UploadCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetUID sets the "uid" field.
func (uc *UploadCreate) SetUID(u uuid.UUID) *UploadCreate {
	uc.mutation.SetUID(u)
	return uc
}

// SetMimeType sets the "mime_type" field.
func (uc *UploadCreate) SetMimeType(s string) *UploadCreate {
	uc.mutation.SetMimeType(s)
	return uc
}

// SetSize sets the "size" field.
func (uc *UploadCreate) SetSize(i int) *UploadCreate {
	uc.mutation.SetSize(i)
	return uc
}

// SetTitle sets the "title" field.
func (uc *UploadCreate) SetTitle(s string) *UploadCreate {
	uc.mutation.SetTitle(s)
	return uc
}

// SetDescription sets the "description" field.
func (uc *UploadCreate) SetDescription(s string) *UploadCreate {
	uc.mutation.SetDescription(s)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UploadCreate) SetCreatedAt(t time.Time) *UploadCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UploadCreate) SetNillableCreatedAt(t *time.Time) *UploadCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UploadCreate) SetUpdatedAt(t time.Time) *UploadCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UploadCreate) SetNillableUpdatedAt(t *time.Time) *UploadCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetUser sets the "user" edge to the User entity.
func (uc *UploadCreate) SetUser(u *User) *UploadCreate {
	return uc.SetUserID(u.ID)
}

// Mutation returns the UploadMutation object of the builder.
func (uc *UploadCreate) Mutation() *UploadMutation {
	return uc.mutation
}

// Save creates the Upload in the database.
func (uc *UploadCreate) Save(ctx context.Context) (*Upload, error) {
	var (
		err  error
		node *Upload
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UploadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Upload)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UploadMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UploadCreate) SaveX(ctx context.Context) *Upload {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UploadCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UploadCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UploadCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := upload.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := upload.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UploadCreate) check() error {
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Upload.user_id"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Upload.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := upload.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Upload.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "Upload.uid"`)}
	}
	if _, ok := uc.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mime_type", err: errors.New(`ent: missing required field "Upload.mime_type"`)}
	}
	if _, ok := uc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Upload.size"`)}
	}
	if _, ok := uc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Upload.title"`)}
	}
	if _, ok := uc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Upload.description"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Upload.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Upload.updated_at"`)}
	}
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Upload.user"`)}
	}
	return nil
}

func (uc *UploadCreate) sqlSave(ctx context.Context) (*Upload, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UploadCreate) createSpec() (*Upload, *sqlgraph.CreateSpec) {
	var (
		_node = &Upload{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: upload.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upload.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: upload.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := uc.mutation.MimeType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldMimeType,
		})
		_node.MimeType = value
	}
	if value, ok := uc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: upload.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := uc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := uc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: upload.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: upload.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := uc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.UserTable,
			Columns: []string{upload.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UploadCreateBulk is the builder for creating many Upload entities in bulk.
type UploadCreateBulk struct {
	config
	builders []*UploadCreate
}

// Save creates the Upload entities in the database.
func (ucb *UploadCreateBulk) Save(ctx context.Context) ([]*Upload, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Upload, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UploadMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UploadCreateBulk) SaveX(ctx context.Context) []*Upload {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UploadCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UploadCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
