package schema

import "entgo.io/ent"

// NoticeChannel holds the schema definition for the NoticeChannel entity.
type NoticeChannel struct {
	ent.Schema
}

// Fields of the NoticeChannel.
func (NoticeChannel) Fields() []ent.Field {
	return nil
}

// Edges of the NoticeChannel.
func (NoticeChannel) Edges() []ent.Edge {
	return nil
}
