package models

import "../gopkg.in/mgo.v2/bson"

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

//type SingleIssueOverview struct {
//	GithubRepo  string
//	Database string
//}
//
//
//type IssueOverview struct {
//	Report   []SingleIssueOverview
//	TotalIssueCount string
//	TotalRepoCount int32
//}

type Report struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"database"`
}