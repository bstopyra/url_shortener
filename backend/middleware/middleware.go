package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ogierhaq/url_shortener/backend/models"
	"github.com/ogierhaq/url_shortener/backend/shortener"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb+srv://admin:IgnRGGIlVxNGeKMv@cluster0.6bhdk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

// Database Name
const dbName = "url_shortener"

// Collection Name
const collName = "urls"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {
	
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Connected to MongoDB!")

	fmt.Println("Collection instance created")
}

func HandleShortenedUrl(shortenedUrl string) string {

	toFind := "http://localhost:8080/" + shortenedUrl

	filterCursor, err := collection.Find(context.Background(), bson.M{"shortenedurl": toFind})

	if err != nil {
		log.Fatal(err)
	}
	var filteredUrl []models.Link

	if err = filterCursor.All(context.Background(), &filteredUrl); err != nil {
		log.Fatal(err)
	}

	if len(filteredUrl) > 0 {
		return filteredUrl[0].OriginURL
	}

	return "http://localhost:3000"
}

func PostLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	
	var url string;

	for key := range r.PostForm {
		url = key
	}

	if len(checkUrl(url)) > 0 {
		fmt.Fprintf(w, "%v", checkUrl(url))
	} else {
		link := models.Link{
			OriginURL: url,
			ShortenedURL: "http://localhost:8080/" + shortener.RandStringRuner(7),
		}
	
		fmt.Fprintf(w, "%v", link.ShortenedURL)
	
		insertOneLink(link)
	}	
}

func insertOneLink(link models.Link) {
	insertResult, err := collection.InsertOne(context.Background(), link)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func checkUrl(url string) string{
	filterCursor, err := collection.Find(context.Background(), bson.M{"originurl": url})

	if err != nil {
		log.Fatal(err)
	}
	var filteredUrl []models.Link
	
	if err = filterCursor.All(context.Background(), &filteredUrl); err != nil {
		log.Fatal(err)
	}

	if len(filteredUrl) > 0 {
		return filteredUrl[0].ShortenedURL
	}

	return ""
}