package main

import (
	"fmt"
	"log"

	as "github.com/aerospike/aerospike-client-go"
)

const Namespace = "test"
const Set = "testset"

func main() {
	client, err := as.NewClient("localhost", 3000)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create key for a record
	key, err := as.NewKey(Namespace, Set, "aditya")
	if err != nil {
		log.Fatal(err)
	}

	binAge := as.NewBin("age", 25)
	binName := as.NewBin("name", "Aditya")

	client.PutBins(nil, key, binAge, binName)

	// Read a record
	record, err := client.Get(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(record.Bins)
}
