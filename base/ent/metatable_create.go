// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mazti/restless/base/ent/metaschema"
	"github.com/mazti/restless/base/ent/metatable"
)

// MetaTableCreate is the builder for creating a MetaTable entity.
type MetaTableCreate struct {
	config
	name       *string
	created_at *time.Time
	updated_at *time.Time
	deleted_at *time.Time
	schema     map[int]struct{}
}

// SetName sets the name field.
func (mtc *MetaTableCreate) SetName(s string) *MetaTableCreate {
	mtc.name = &s
	return mtc
}

// SetCreatedAt sets the created_at field.
func (mtc *MetaTableCreate) SetCreatedAt(t time.Time) *MetaTableCreate {
	mtc.created_at = &t
	return mtc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mtc *MetaTableCreate) SetNillableCreatedAt(t *time.Time) *MetaTableCreate {
	if t != nil {
		mtc.SetCreatedAt(*t)
	}
	return mtc
}

// SetUpdatedAt sets the updated_at field.
func (mtc *MetaTableCreate) SetUpdatedAt(t time.Time) *MetaTableCreate {
	mtc.updated_at = &t
	return mtc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mtc *MetaTableCreate) SetNillableUpdatedAt(t *time.Time) *MetaTableCreate {
	if t != nil {
		mtc.SetUpdatedAt(*t)
	}
	return mtc
}

// SetDeletedAt sets the deleted_at field.
func (mtc *MetaTableCreate) SetDeletedAt(t time.Time) *MetaTableCreate {
	mtc.deleted_at = &t
	return mtc
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (mtc *MetaTableCreate) SetNillableDeletedAt(t *time.Time) *MetaTableCreate {
	if t != nil {
		mtc.SetDeletedAt(*t)
	}
	return mtc
}

// SetSchemaID sets the schema edge to MetaSchema by id.
func (mtc *MetaTableCreate) SetSchemaID(id int) *MetaTableCreate {
	if mtc.schema == nil {
		mtc.schema = make(map[int]struct{})
	}
	mtc.schema[id] = struct{}{}
	return mtc
}

// SetNillableSchemaID sets the schema edge to MetaSchema by id if the given value is not nil.
func (mtc *MetaTableCreate) SetNillableSchemaID(id *int) *MetaTableCreate {
	if id != nil {
		mtc = mtc.SetSchemaID(*id)
	}
	return mtc
}

// SetSchema sets the schema edge to MetaSchema.
func (mtc *MetaTableCreate) SetSchema(m *MetaSchema) *MetaTableCreate {
	return mtc.SetSchemaID(m.ID)
}

// Save creates the MetaTable in the database.
func (mtc *MetaTableCreate) Save(ctx context.Context) (*MetaTable, error) {
	if mtc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if mtc.created_at == nil {
		v := metatable.DefaultCreatedAt()
		mtc.created_at = &v
	}
	if mtc.updated_at == nil {
		v := metatable.DefaultUpdatedAt()
		mtc.updated_at = &v
	}
	if len(mtc.schema) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"schema\"")
	}
	return mtc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MetaTableCreate) SaveX(ctx context.Context) *MetaTable {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mtc *MetaTableCreate) sqlSave(ctx context.Context) (*MetaTable, error) {
	var (
		mt   = &MetaTable{config: mtc.config}
		spec = &sqlgraph.CreateSpec{
			Table: metatable.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metatable.FieldID,
			},
		}
	)
	if value := mtc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: metatable.FieldName,
		})
		mt.Name = *value
	}
	if value := mtc.created_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: metatable.FieldCreatedAt,
		})
		mt.CreatedAt = *value
	}
	if value := mtc.updated_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: metatable.FieldUpdatedAt,
		})
		mt.UpdatedAt = *value
	}
	if value := mtc.deleted_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: metatable.FieldDeletedAt,
		})
		mt.DeletedAt = value
	}
	if nodes := mtc.schema; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metatable.SchemaTable,
			Columns: []string{metatable.SchemaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metaschema.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, mtc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	mt.ID = int(id)
	return mt, nil
}
