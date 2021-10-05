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

const connectionString = "mongodb+srv://name:password@cluster0.6bhdk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "url_shortener"
const collName = "urls"
const endpoint = "http://localhost:8080/"
var collection *mongo.Collection

func init() {
	
	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to MongoDB!")
	
	defer cancel()

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created")
}

func HandleShortenedUrl(shortenedUrl string) string {
	elemToFind := endpoint + shortenedUrl
	elemFound := searchForUrl("shortenedurl", elemToFind)
	if len(elemFound.OriginURL) > 0 {
		return elemFound.OriginURL
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

	foundElem := searchForUrl("originurl", url)

	if len(foundElem.ShortenedURL) > 0 {
		fmt.Fprintf(w, "%v", foundElem.ShortenedURL)
	} else {
		link := models.Link{
			OriginURL: url,
			ShortenedURL: endpoint + shortener.RandStringRuner(7),
		}
	
		insertOneLink(link)

		fmt.Fprintf(w, "%v", link.ShortenedURL)
	
	}	
}

func searchForUrl(key string, value string) models.Link {
	filterCursor, err := collection.Find(context.Background(), bson.M{key: value})

	if err != nil {
		log.Fatal(err)
	}

	var filteredUrl []models.Link

	if err = filterCursor.All(context.Background(), &filteredUrl); err != nil {
		log.Fatal(err)
	}

	if len(filteredUrl) > 0 {
		return filteredUrl[0]
	}

	return models.Link{}
}

func insertOneLink(link models.Link) {
	insertResult, err := collection.InsertOne(context.Background(), link)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
