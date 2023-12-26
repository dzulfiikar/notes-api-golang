package schemas

type Note struct {
	ID        string `bson:"_id,omitempty"`
	Title     string `bson:"title,omitempty"`
	Content   string `bson:"content,omitempty"`
	CreatedAt string `bson:"created_at,omitempty"`
	CreatedBy string `bson:"created_by,omitempty"`
	UpdatedAt string `bson:"updated_at,omitempty"`
	UpdatedBy string `bson:"updated_by,omitempty"`
	DeletedAt string `bson:"deleted_at,omitempty"`
	DeletedBy string `bson:"deleted_by,omitempty"`
	Deleted   bool   `bson:"deleted,omitempty"`
}
