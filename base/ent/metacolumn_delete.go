// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mazti/restless/base/ent/metacolumn"
	"github.com/mazti/restless/base/ent/predicate"
)

// MetaColumnDelete is the builder for deleting a MetaColumn entity.
type MetaColumnDelete struct {
	config
	predicates []predicate.MetaColumn
}

// Where adds a new predicate to the delete builder.
func (mcd *MetaColumnDelete) Where(ps ...predicate.MetaColumn) *MetaColumnDelete {
	mcd.predicates = append(mcd.predicates, ps...)
	return mcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcd *MetaColumnDelete) Exec(ctx context.Context) (int, error) {
	return mcd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (mcd *MetaColumnDelete) ExecX(ctx context.Context) int {
	n, err := mcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcd *MetaColumnDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: metacolumn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metacolumn.FieldID,
			},
		},
	}
	if ps := mcd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mcd.driver, spec)
}

// MetaColumnDeleteOne is the builder for deleting a single MetaColumn entity.
type MetaColumnDeleteOne struct {
	mcd *MetaColumnDelete
}

// Exec executes the deletion query.
func (mcdo *MetaColumnDeleteOne) Exec(ctx context.Context) error {
	n, err := mcdo.mcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{metacolumn.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcdo *MetaColumnDeleteOne) ExecX(ctx context.Context) {
	mcdo.mcd.ExecX(ctx)
}
