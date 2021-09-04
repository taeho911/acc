package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Acc struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Index    int                `bson:"index,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Location string             `bson:"location,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Memo     string             `bson:"memo,omitempty"`
	Alias    []string           `bson:"alias,omitempty"`
	Deleted  bool               `bson:"deleted"`
}
