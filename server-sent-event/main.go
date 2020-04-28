package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var counter int

type (
	Currency float64
	Item     struct {
		Name     string   `json:"name,omitempty"`
		Quantity int      `json:"quantity,omitempty"`
		Price    Currency `json:"price,omitempty"`
	}
	Store struct {
		Items map[string]Item `json:"items,omitempty"`
	}
	Dashboard struct {
		Users         uint   `json:"users,omitempty"`
		UsersLoggedIn uint   `json:"users_logged_in,omitempty"`
		Inventory     *Store `json:"inventory,omitempty"`
	}
)

var dashboard chan *Dashboard

func main() {
	dashboard = make(chan *Dashboard)
	go updateDashboard()
	// Register static files handle '/index.html' ->serves-> client/index.html
	http.Handle("/", http.FileServer(http.Dir("client")))
	// Register RESTful  handler for 'sse/dashboard'
	http.HandleFunc("/sse/dashboard", dashboardHandler)
	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(<-dashboard)
	fmt.Fprintf(w, "data: %v\n\n", buf.String())
	fmt.Printf("data: %v\n", buf.String())
}

func updateDashboard() {
	for {
		inv := updateInventory()
		db := &Dashboard{
			Users:         uint(rand.Uint32()),
			UsersLoggedIn: uint(rand.Uint32() % 200),
			Inventory:     inv,
		}
		dashboard <- db
	}
}

func updateInventory() *Store {
	inv := &Store{}
	inv.Items = make(map[string]Item)

	a := Item{Name: "Books", Price: 33.59, Quantity: int(rand.Int31() % 53)}
	inv.Items["book"] = a

	a = Item{Name: "Bicyles", Price: 190.89, Quantity: int(rand.Int31() % 232)}
	inv.Items["bicycle"] = a

	a = Item{Name: "Water Bottles", Price: 10.02, Quantity: int(rand.Int31() % 93)}
	inv.Items["wbottle"] = a
	return inv
}
