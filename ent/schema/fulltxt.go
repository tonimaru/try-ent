package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Fulltxt holds the schema definition for the Fulltxt entity.
type Fulltxt struct {
	ent.Schema
}

// Fields of the Fulltxt.
func (Fulltxt) Fields() []ent.Field {
	return []ent.Field{
		field.Text("txt"),
	}
}

// Edges of the Fulltxt.
func (Fulltxt) Edges() []ent.Edge {
	return nil
}

// Indexes returns the indexes of the schema.
func (f Fulltxt) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("txt").Annotations(entsql.IndexType("FULLTEXT")),
		// Indexing using the ngram full-text parser #2911
		// https://github.com/ent/ent/issues/2911
	}
}
