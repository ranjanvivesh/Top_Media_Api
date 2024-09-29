package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/ranjanvivesh/topmedia/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString = "mongodb+srv://vivesh:Vivesh4153@cluster.bi0vy.mongodb.net/?retryWrites=true&w=majority&appName=Cluster"
const dbName = "RankList"
const colName = "EntList"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The MongoDb Connection Successfull")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection Instance is ready")

}

func addOneTitle(title model.Cluster) {
	insert, err := collection.InsertOne(context.Background(), title)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Id of the inserted Title  :", insert.InsertedID)
}

func updateOneTitle(titleID string) {
	id, _ := primitive.ObjectIDFromHex(titleID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

func deleteOneTitle(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}

	deletedCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DeleteCount:", deletedCount)
}

func deleteAllTitles() {
	filter := bson.M{}
	deletedCount, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DeleteCount:", deletedCount)

}

func getAllTitles() []primitive.M {
	filter := bson.D{{}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var titles []primitive.M

	for cur.Next(context.Background()) {
		var title bson.M
		err := cur.Decode(&title)
		if err != nil {
			log.Fatal(err)
		}
		titles = append(titles, title)
	}
	defer cur.Close(context.Background())

	return titles
}

func GetAllTitles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	alltitles := getAllTitles()
	json.NewEncoder(w).Encode(alltitles)
}

func CreateTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var title model.Cluster
	_ = json.NewDecoder(r.Body).Decode(&title)
	addOneTitle(title)
	json.NewEncoder(w).Encode(title)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneTitle(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteATitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneTitle(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllTitles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllTitles()

}
