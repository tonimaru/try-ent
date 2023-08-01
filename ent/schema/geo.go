package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/tonimaru/try-ent/pkg/geo"
)

// Geo holds the schema definition for the Geo entity.
type Geo struct {
	ent.Schema
}

// Fields of the Geo.
func (Geo) Fields() []ent.Field {
	return []ent.Field{
		field.Other("point", &geo.Point{}).
			SchemaType(geo.PointSchemaType()),
	}
}

// Edges of the Geo.
func (Geo) Edges() []ent.Edge {
	return nil
}
