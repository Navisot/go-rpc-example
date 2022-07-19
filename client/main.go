package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Item holds the item struct
type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("error dial http:", err)
	}

	a := Item{Title: "first", Body: "first description"}
	b := Item{Title: "second", Body: "second description"}
	c := Item{Title: "third", Body: "third description"}

	client.Call("API.CreateItem", a, &reply)
	client.Call("API.CreateItem", b, &reply)
	client.Call("API.CreateItem", c, &reply)

	client.Call("API.GetDB", "", &db)

	for idx, it := range db {
		fmt.Printf("Index is: %d and title is %s \n", idx, it.Title)
		fmt.Printf("Index is: %d and body is %s \n", idx, it.Body)
	}
}
