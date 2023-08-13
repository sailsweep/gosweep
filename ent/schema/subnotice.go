package schema

import "entgo.io/ent"

// SubNotice holds the schema definition for the SubNotice entity.
type SubNotice struct {
	ent.Schema
}

// Fields of the SubNotice.
func (SubNotice) Fields() []ent.Field {
	return nil
}

// Edges of the SubNotice.
func (SubNotice) Edges() []ent.Edge {
	return nil
}
