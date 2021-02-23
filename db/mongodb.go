package db

import (
	"context"
	"fmt"
	"go-todo/models"
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
const USERCOLLECTION = "users"
const USERTASKCOLLECTION = "users_tasks"

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

func InsertTask(task models.Task) interface{} {
	//var result models.Task
	fmt.Println("InsertTask", task)
	result, err := db.Collection(COLLECTNAME).InsertOne(context.Background(), task)

	fmt.Println("INSERTED TASK", result)

	if err != nil {
		fmt.Println("Errs", err)
		log.Fatal(err)
	}

	return result.InsertedID
}

func GetTasksByUser(userId string) []models.Task {
	var results []models.UsersTask
	var tasks []models.Task
	//var task models.Task
	var elem models.UsersTask
	cur, err := db.Collection(USERTASKCOLLECTION).Find(context.Background(), bson.M{"userid": userId})
	if err != nil {
		fmt.Println("Errs", err)
		log.Fatal(err)
	}
	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Err", err)
			log.Fatal(err)
		}
		//results = append(results, elem.TaskId)
		fmt.Println("TASK ID", elem.TaskId)
	}
	fmt.Println("TASK BY USER", results)

	/*filterCur := db.Collection(COLLECTNAME).Find(context.Background(), results)
	for filterCur.Next(context.Background()) {
		err := filterCur.Decode(&task)
		if err != nil {
			fmt.Println("Err", err)
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("Erros", err)
		log.Fatal(err)
	}
	cur.Close(context.Background())*/

	return tasks
}

func GetTasks() []models.Task {
	//cur, err := db.Collection(COLLECTNAME).Find(context.Background(), bson.D{{}})
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
	idPrimitive, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	result, err := db.Collection(COLLECTNAME).UpdateOne(
		context.Background(),
		bson.M{"_id": idPrimitive},
		bson.D{
			{"$set", bson.D{{"iscomplete", true}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
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

func InsertUser(user models.User) {
	fmt.Println("Create User", user)
	_, err := db.Collection(USERCOLLECTION).InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println("Errs", err)
		log.Fatal(err)
	}
}

func GetUser(user models.User) models.User {
	var result models.User
	if err := db.Collection(USERCOLLECTION).FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("CURE", result)

	return result
}

func AssignTask(usertask models.UsersTask) {
	_, err := db.Collection(USERTASKCOLLECTION).InsertOne(context.Background(), usertask)
	if err != nil {
		fmt.Println("Errs", err)
		log.Fatal(err)
	}
}
