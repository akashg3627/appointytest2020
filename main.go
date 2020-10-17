package main

import (
    "context"
    "fmt"
	"net/http"
	"encoding/json"
	"time"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program

type Participant struct {
	Name string
	Email string
	RSPV string
}

type Meeting struct {
	Id primitive.ObjectID 
	Title string
	Participants []Participant
	StartTime string
	EndTime string
	CreationTimestamp string
}

var client *mongo.Client

func scheduleMeeting( w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		w.Header().Set("content-type", "application/json")
		var meeting Meeting
		json.NewDecoder(r.Body).Decode(&meeting)
		collection := client.Database("test").Collection("meeting")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, meeting)
	json.NewEncoder(w).Encode(result)
	} else{
		//message = "Method not allowed"
        http.Redirect(w, r, "/", http.StatusFound)
	}
}
func meetingWithID( w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		idt := r.URL.Path[len("/meeting/"):]
		id, _:= primitive.ObjectIDFromHex(idt)
		var selectedmeeting Meeting
        collection := client.Database("test").Collection("meeting")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	
	err := collection.FindOne(ctx, Meeting{Id: id}).Decode(&selectedmeeting)
	if err != nil {
	log.Fatal(err)
	}
	json.NewEncoder(w).Encode(selectedmeeting)
		fmt.Println("meeting with id",selectedmeeting )

	}	else{
		//message = "Method not allowed"
        http.Redirect(w, r, "/", http.StatusFound)
	}
}
func meetingBtwTimeFrame( w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
fmt.Println("meeting with tym")
	}	else{
		//message = "Method not allowed"
        http.Redirect(w, r, "/", http.StatusFound)
	}
}
func meetingWithParticipantID( w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
fmt.Println("meeting with participant")
	}	else{
		//message = "Method not allowed"
        http.Redirect(w, r, "/", http.StatusFound)
	}
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://golang:12345@cluster0.f6q0p.mongodb.net/<dbname>?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
http.HandleFunc("/meeting", scheduleMeeting)
http.HandleFunc("/meeting/{id}", meetingWithID)
http.HandleFunc("/", meetingBtwTimeFrame)
http.HandleFunc("/particitant/{id}", meetingWithParticipantID)

http.ListenAndServe(":3000", nil) // set listen port

}