package main

// Importing Libraries
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
	"sync"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client
var lock sync.Mutex	
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


// Meetings Controller
func meetings(w http.ResponseWriter, r *http.Request){
	lock.Lock()
    defer lock.Unlock()
	switch r.Method {
		case "GET":     
			// Storing the Queries params in variables
			startStr := r.URL.Query().Get("start")
			endStr := r.URL.Query().Get("end")
			participant:= r.URL.Query().Get("participant")
			offset:= r.URL.Query().Get("offset")
			limit := r.URL.Query().Get("limit")
			//For start and end query params
			if((startStr != "" && endStr != "")&& participant == ""){
				start, err :=  strconv.Atoi(startStr)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				end, err := strconv.Atoi(endStr)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)

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
				


				var meetings []Meeting
				meetingApiDatabase := client.Database("MeetingsApi")
				meetingsCollection := meetingApiDatabase.Collection("meetings_test")
				filterCursor, err := meetingsCollection.Find(ctx, bson.D{{"start_Time", bson.D{{"$gte", start}}},{"end_Time", bson.D{{"$lte", end}}}})
				if err != nil {
					log.Fatal(err)
				}
				if err = filterCursor.All(ctx, &meetings); err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(meetings[9])


			// For participant params
			} else if( participant!="" &&(startStr=="" && endStr =="")){
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
				opts := options.Find()
				if(offset!=""&&limit!=""){
					skip,err := strconv.Atoi(offset)
					if err != nil {
						log.Fatal(err)
					}
					limit,err := strconv.Atoi(limit)
					if err != nil {
						log.Fatal(err)
					}
					opts.SetLimit(int64(limit))
					opts.SetSkip(int64(skip))
				}
			
					// opts := options.FindOptions{
					// 	Skip: i(skip),
					// 	Limit: (limit),
					//   }
				defer client.Disconnect(ctx)
				var meetings []Meeting
				meetingApiDatabase := client.Database("MeetingsApi")
				meetingsCollection := meetingApiDatabase.Collection("meetings_test")
	
				filterCursor, err := meetingsCollection.Find(ctx, bson.D{{"participants",bson.D{{"$elemMatch",bson.D{{"email",participant}}}}}},opts)
				if err != nil {
					log.Fatal(err)
				}
				if err = filterCursor.All(ctx, &meetings); err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(meetings)
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
			meetingApiDatabase := client.Database("MeetingsApi")
			meetingsCollection := meetingApiDatabase.Collection("meetings_test")
			var m1 []Meeting
			noCLashing :=0
				for i, s := range m.Participants {
					fmt.Println(i, s)
					filterCursor, err := meetingsCollection.Find(ctx, bson.D{{"start_Time", bson.D{{"$gte", m.Start_Time }}},{"end_Time", bson.D{{"$lte",  m.End_Time }}},{"participants",bson.D{{"$elemMatch",bson.D{{"email",s.Email},{"rsvp","yes"}}}}}})
					if err != nil {
						log.Fatal(err)
					}
					if err = filterCursor.All(ctx, &m1); err != nil {
						log.Fatal(err)
					}
					fmt.Println(m1)
					if(m1!=nil){
						fmt.Fprintf(w,"One of the participants with email "+s.Email+" timings are clashing")
						noCLashing=1
						break;
					}
				}
			if(noCLashing==0){
				meetingResult, err := meetingsCollection.InsertOne(ctx, m)
				if err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(meetingResult)

			
			}
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

}

// Meeting Controller
func meeting(w http.ResponseWriter, r *http.Request){
	lock.Lock()
    defer lock.Unlock()
	switch r.Method {
		case "GET":     
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
			meetingApiDatabase := client.Database("MeetingsApi")
			meetingsCollection := meetingApiDatabase.Collection("meetings_test")
			fmt.Println( "Not post yet done")
			if err = meetingsCollection.FindOne(ctx,bson.M{"_id": varId}).Decode(&m); err != nil {
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
	fmt.Fprintf(w,"Welcome to the MeetingsAPI")
}

// Routing
func handleRequests(){
	http.HandleFunc("/meeting/",meeting)
	http.HandleFunc("/",homePage)
	http.HandleFunc("/meetings",meetings)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	
	fmt.Println("Starting the application...")
	fmt.Println(time.Now())
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