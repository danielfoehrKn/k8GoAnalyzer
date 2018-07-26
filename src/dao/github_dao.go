package dao

import (
"log"
. "../models"
"../gopkg.in/mgo.v2"
"../gopkg.in/mgo.v2/bson"
	"fmt"
)

type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	//Db name: github , collection: channels
	COLLECTION_MOVIES = "movies"
	COLLECTION_CHANNELS = "channels"
)

// Establish a connection to database
func (m *DAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to db successfull")
	db = session.DB(m.Database)
}

// Find list of movies
func (m *DAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION_MOVIES).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *DAO) FindAllChannels() ([]Channel, error) {
	var channels []Channel
	err := db.C(COLLECTION_CHANNELS).Find(bson.M{}).All(&channels)
	return channels, err
}

func (m *DAO) GenerateReport() ([]bson.M ,error) {
// see https://stackoverflow.com/questions/26062658/mongodb-aggregation-in-golang
	o1 := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"repo": "$githubRepo",
				"action": "$action",
				"label": "$label",
			},
			"count": bson.M{
				"$sum": 1,
			},
		},
	}
	// ability to chain a lot of commands like filter, match etc.
	operations := []bson.M{o1}

	pipe := db.C(COLLECTION_CHANNELS).Pipe(operations)

	// Run the queries and capture the results
	results := []bson.M{}
	err1 := pipe.All(&results)

	if err1 != nil {
		fmt.Printf("ERROR : %s\n", err1.Error())
		return nil, err1
	} else {
		//fmt.Printf("Id : %s, Size: %sn", results[0]["_id"], results[0]["count"])
		fmt.Printf("Results : %s", results)
		return results, nil
	}
}

// Find a movie by its id
func (m *DAO) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION_MOVIES).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *DAO) FindChannelById(id string) (Channel, error) {
	var channel Channel
	err := db.C(COLLECTION_CHANNELS).FindId(bson.ObjectIdHex(id)).One(&channel)
	return channel, err
}

// Insert a movie into database
func (m *DAO) Insert(movie Movie) error {
	err := db.C(COLLECTION_MOVIES).Insert(&movie)
	return err
}

func (m *DAO) InsertChannel(channel Channel) error {
	err := db.C(COLLECTION_CHANNELS).Insert(&channel)
	return err
}

// Delete an existing movie
func (m *DAO) Delete(movie Movie) error {
	err := db.C(COLLECTION_MOVIES).Remove(&movie)
	return err
}

// Update an existing movie
func (m *DAO) Update(movie Movie) error {
	err := db.C(COLLECTION_MOVIES).UpdateId(movie.ID, &movie)
	return err
}
