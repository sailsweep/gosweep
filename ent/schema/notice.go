package schema

import "entgo.io/ent"

// Notice holds the schema definition for the Notice entity.
type Notice struct {
	ent.Schema
}

// Fields of the Notice.
func (Notice) Fields() []ent.Field {
	return nil
}

// Edges of the Notice.
func (Notice) Edges() []ent.Edge {
	return nil
}
