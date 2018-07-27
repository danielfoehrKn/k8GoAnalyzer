package main

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

type Channel struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Action     string        `bson:"action" json:"action"`
	GithubRepo string        `bson:"githubRepo" json:"githubRepo"`
	Label      string        `bson:"label" json:"label"`
}