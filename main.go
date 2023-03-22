package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	
	//inicia la app de fiber
	app := fiber.New()

	//conexion con db mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb"))

	if err != nil {
		panic(err)
	}
	coll := client.Database("gomongodb").Collection("users")
	coll.InsertOne(context.TODO(), bson.D{{Key: "name", Value: "seba",}})


	//cors para abilitar la conexion desde el front
	app.Use(cors.New())

	//rutas
	app.Static("/", "./client/dist")
	
	app.Get("/users", func(c *fiber.Ctx)error {
		return c.JSON(&fiber.Map{"data": "users desde el backend",
		})
	})

	app.Post("/users", func(c *fiber.Ctx)error {
		
		return c.JSON(&fiber.Map{"data": "Creando usuario",
		})
	})

	//puerto
	app.Listen(":3000")
	fmt.Println("Server on port 3000")

}			

//go run .		