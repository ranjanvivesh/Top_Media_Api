package model

import(
     "go.mongodb.org/mongo-driver/bson/primitive"

)

type Cluster struct {
	Id 	      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	MovieName string `json:"moviename"`
	Director  string `json:"director"`
	Year	  string `json:"year"`
	PG_Rating string `json:"pg_rating"`
	Watched	  bool   `json:"watched`	
}