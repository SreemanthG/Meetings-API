package main

import (
	"fmt"
	"log"
	"path"
	"net/http"
	"encoding/json"
	"context"
	"time"
	"strconv"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client

type Participant struct {
	// ID primitive.ObjectID ` bson:"_id,omitempty" json:"_id,omitempty"`
	Name string `bson:"name"  json:"name"`
	Email string `bson:"email" json:"email"`
	RSVP  string `bson:"rsvp" json:"rsvp"`
}
type Meeting struct {	
	ID primitive.ObjectID ` bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `json:"title" bson:"title" `
	Participants []Participant `json:"participants" bson:"participants"`
	Start_Time	int `json:"start_Time" bson:"start_Time"`
	End_Time int `json:"end_Time" bson:"end_Time"`
	Creation_Timestamp	time.Time `json:"creation_Timestamp" bson:"creation_Timestamp"`
}



func meetings(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":     
		
			// endstr := ( )
			startstr := r.URL.Query().Get("start")
			endstr := r.URL.Query().Get("end")
			participant:= r.URL.Query().Get("participant")

			//Route with start and end as parameters
			if((startstr != "" && endstr != "")&& participant == ""){
				start1, err :=  strconv.Atoi(startstr)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				end1, err := strconv.Atoi(endstr)
				if err != nil {
					fmt.Println(err)
				} 
				client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
				if err != nil {
					log.Fatal(err)
				}
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				err = client.Connect(ctx)
				if err != nil {
					log.Fatal(err)
				}
				defer client.Disconnect(ctx)
				var m1 []Meeting
				quickstartDatabase := client.Database("MeetingsApi")
				podcastsCollection := quickstartDatabase.Collection("meetings_test")
	
				filterCursor, err := podcastsCollection.Find(ctx, bson.D{{"start_Time", bson.D{{"$gte", start1}}},{"end_Time", bson.D{{"$lte", end1}}}})
				if err != nil {
					log.Fatal(err)
				}
				if err = filterCursor.All(ctx, &m1); err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(m1)
			} else if( participant!="" &&(startstr=="" && endstr =="")){
				fmt.Println(participant)
				client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
				if err != nil {
					log.Fatal(err)
				}
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				err = client.Connect(ctx)
				if err != nil {
					log.Fatal(err)
				}
				defer client.Disconnect(ctx)
				var m1 []Meeting
				quickstartDatabase := client.Database("MeetingsApi")
				podcastsCollection := quickstartDatabase.Collection("meetings_test")
	
				filterCursor, err := podcastsCollection.Find(ctx, bson.D{{"participants",bson.D{{"$elemMatch",bson.D{{"email",participant}}}}}})
				if err != nil {
					log.Fatal(err)
				}
				if err = filterCursor.All(ctx, &m1); err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(m1)
			} else{
				fmt.Fprintf(w, "pass the right parameters")
			}
		
			
		case "POST":
			decoder := json.NewDecoder(r.Body)
			var m Meeting
			err := decoder.Decode(&m)
			if err != nil {
				panic(err)
			}
			m.Creation_Timestamp = time.Now()
			fmt.Printf("%+v\n", m)
			
			client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				log.Fatal(err)
			}
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}
			defer client.Disconnect(ctx)
			quickstartDatabase := client.Database("MeetingsApi")
			podcastsCollection := quickstartDatabase.Collection("meetings_test")
			// episodesCollection := quickstartDatabase.Collection("episodes")
			podcastResult, err := podcastsCollection.InsertOne(ctx, m)
			if err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(podcastResult)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

}


func meeting(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":     
			// vars := mux.Vars(r)
			// varId := vars["id"]  
			// varId  := "path.Base(r.URL.Path)"
			// fmt.Println(primitive.ObjectIDFromHex(varId))
			varId,err := primitive.ObjectIDFromHex(path.Base(r.URL.Path))
			if err != nil {
					log.Fatal(err)
				}
			client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				log.Fatal(err)
			}
			var m Meeting
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}

			defer client.Disconnect(ctx)
			// var episodesFiltered []bson.M
			quickstartDatabase := client.Database("MeetingsApi")
			podcastsCollection := quickstartDatabase.Collection("meetings_test")
			fmt.Println( "Not post yet done")

			if err = podcastsCollection.FindOne(ctx,bson.M{"_id": varId}).Decode(&m); err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(m)
		case "POST":
			fmt.Fprintf(w, "Not post yet done")
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

}


func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"Homepage Endpoint Hit")
}

func handleRequests(){
	http.HandleFunc("/meeting/",meeting)
	http.HandleFunc("/",homePage)
	http.HandleFunc("/meetings",meetings)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){

	fmt.Println("Starting the application...")
	fmt.Println(time.Now())

	// ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	// client,_ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

	handleRequests()
}