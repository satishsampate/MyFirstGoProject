package main

import (
	"com/tutorials/helpers"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("First go program")

	var myStudnet helpers.Student

	myStudnet.Name = " name sirname"
	myStudnet.RollNumber = 45
	fmt.Printf("Stuend name is %s and his roll number is %d", myStudnet.Name, myStudnet.RollNumber)

	// read json and convert it to go
	// read go struct and convert it to the json.

	myJson := `[
		{"full_name":"blessing james",
		"email":"blessing@gmail.com",
		"gender":"Male",
		"status":"active"}
		,
		{"full_name":"matt john",
		"email":"matt@gmail.com",
		"gender":"Male",
		"status":"active"}
		,
		{"full_name":"john peace",
		"email":"peace@gmail.com",
		"gender":"Midgard",
		"status":"active"}]`

	type Person struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Gender   string `json:"gender"`
		Status   string `json:"status"`
	}

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	fmt.Printf("unmarshalled : %v", unmarshalled)

	// connecting to the redis from go
	// https://tutorialedge.net/golang/go-redis-tutorial/

	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// we can call set with a `Key` and a `Value`.
	err = client.Set("name", "Elliot", 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}
	client.Close()

	// connect go to the aerospike database.

}
