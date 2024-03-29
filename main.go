package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sxc/cali/api"
	"github.com/sxc/cali/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const dburi = "mongodb://localhost:27017"
const dburi = "mongodb://root:password@localhost:27017/"
const dbname = "cali"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{
		FirstName: "John",
		LastName:  "Appleseed",
	}

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var john types.User

	if err := coll.FindOne(ctx, bson.M{}).Decode(&john); err != nil {
		log.Fatal(err)
	}

	fmt.Println(john)

	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUserByID)
	app.Listen(*listenAddr)
}
