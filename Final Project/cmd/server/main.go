package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"assignment-2/internal/database"
	"assignment-2/internal/handlers"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := handlers.NewRouter(db)

	go updateStatus()

	router.HandleFunc("/", handleHome)
	router.HandleFunc("/status", handleStatus)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func updateStatus() {
	for {
		status := Status{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}

		data, err := json.Marshal(status)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			return
		}

		err = ioutil.WriteFile("status.json", data, 0644)
		if err != nil {
			log.Println("Error writing JSON file:", err)
			return
		}

		time.Sleep(15 * time.Second)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("status.json")
	if err != nil {
		log.Println("Error reading JSON file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var status Status
	err = json.Unmarshal(data, &status)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
