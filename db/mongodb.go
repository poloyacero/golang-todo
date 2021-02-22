package db

import (
	"context"
	"fmt"
	"go-server/models"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection domain
const MONGOCONNECTION = "mongodb://localhost:27017"

// DB name
const DBNAME = "todos"

// collection name
const COLLECTNAME = "tasks"

var db *mongo.Database

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(MONGOCONNECTION)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(DBNAME)
	fmt.Println("Connected to MongoDB!")
}

func InsertTask(task models.Task) {
	fmt.Println("InsertTask", task)
	_, err := db.Collection(COLLECTNAME).InsertOne(context.Background(), task)
	if err != nil {
		fmt.Println("Errs", err)
		log.Fatal(err)
	}
}

func GetTasks() []models.Task {
	cur, err := db.Collection(COLLECTNAME).Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println("Errs", err)
		log.Fatal(err)
	}
	var elements []models.Task
	var elem models.Task
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Err", err)
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("Erros", err)
		log.Fatal(err)
	}
	cur.Close(context.Background())
	fmt.Println("elemetns", elements)
	return elements
}

func GetTask(task models.Task) {
	fmt.Println("GetTask", task)
}

func UpdateTask(task models.Task, taskId string) {

}

func DeleteTask(taskId string) {
	fmt.Println("DEL", taskId)
	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	fmt.Println("DEL", idPrimitive)
	res, err := db.Collection(COLLECTNAME).DeleteOne(context.Background(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
	}
	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	} else {
		// Print the results of the DeleteOne() method
		fmt.Println("DeleteOne Result:", res)

		// *mongo.DeleteResult object returned by API call
		fmt.Println("DeleteOne TYPE:", reflect.TypeOf(res))
	}

}
