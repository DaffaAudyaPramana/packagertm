package packagertm

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accounts struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama       string             `bson:"name,omitempty" json:"name,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Sosial     string             `bson:"sosial,omitempty" json:"sosial,omitempty"`
	Perusahaan string             `bson:"perusahaan,omitempty" json:"perusahaan,omitempty"`
}
