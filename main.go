package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Item keeps the struct
type Item struct {
	Title string
	Body  string
}

// database keeps a collection of item
var database []Item

// API init
type API int

// GetByName retrieves an item by its name
func (a *API) GetByName(title string, reply *Item) error {

	var getItem Item

	for _, item := range database {
		if item.Title == title {
			getItem = item
		}
	}

	*reply = getItem

	return nil
}

// CreateItem creates an item and stores it in the database
func (a *API) CreateItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

// EditItem updates an item in the database
func (a *API) EditItem(edit Item, reply *Item) error {

	var changedItem Item

	for idx, item := range database {
		if item.Title == edit.Title {
			database[idx] = Item{Title: edit.Title, Body: edit.Body}
			changedItem = database[idx]
		}
	}

	*reply = changedItem

	return nil
}

// DeleteItem deletes an item from the database
func (a *API) DeleteItem(item Item, reply *Item) error {

	var del Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}

	*reply = del

	return nil
}

// GetDB returns the DB itself
func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("listener error", err)
	}

	log.Printf("serving on port %d", 4040)

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("http serve error", err)
	}
}
