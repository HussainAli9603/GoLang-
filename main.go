package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/mux"
)

var client *mongo.Client

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Time      primitive.DateTime `bson:"time"`
	Symbol string             `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Open   string             `json:"open,omitempty" bson:"open,omitempty"`
	High   string             `json:"high,omitempty" bson:"high,omitempty"`
	Low    string             `json:"low,omitempty" bson:"low,omitempty"`
	Close  string             `json:"close,omitempty" bson:"close,omitempty"`
	Volume string             `json:"volume,omitempty" bson:"volume,omitempty"`
}



func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database("InterviewTest").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}



func main() {
	fmt.Println("Server is running on PORT :: 9000")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://todolist:usman098@cluster0.metrh.mongodb.net/InterviewTest?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println(client)
	router := mux.NewRouter()
	router.HandleFunc("/v1/api/data", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/v1/api/get-all-data", GetPeopleEndpoint).Methods("GET")
	http.ListenAndServe(":9000", router)
}