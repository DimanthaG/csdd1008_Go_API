package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// GameStat struct represents a basic video game stat entity
type GameStat struct {
	ID         int    `json:"id"`
	PlayerName string `json:"player_name"`
	Game       string `json:"game"`
	Score      string `json:"score"`
	Status     string `json:"status"` // "active" or "inactive"
}

var (
	gameStats = make(map[int]GameStat) // In-memory game stat storage
	nextID    = 1                      // Auto-incrementing game stat ID
	statsMux  sync.Mutex               // Mutex for concurrent access
)

// Handler for creating a new game stat (POST /gamestats)
func createGameStatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var newGameStat GameStat
	err := json.NewDecoder(r.Body).Decode(&newGameStat)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}

	statsMux.Lock()
	defer statsMux.Unlock()

	newGameStat.ID = nextID
	nextID++
	gameStats[newGameStat.ID] = newGameStat

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGameStat)
}

// Handler for getting a single game stat by ID (GET /gamestats/{id})
func getGameStatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/gamestats/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid game status ID", http.StatusBadRequest)
		return
	}

	statsMux.Lock()
	defer statsMux.Unlock()

	gameStat, exists := gameStats[id]
	if !exists {
		http.Error(w, "Game status not found or unavailable", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameStat)
}

// for updating a game stat by ID (PUT /gamestats/{id})
func updateGameStatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/gamestats/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid game stat ID", http.StatusBadRequest)
		return
	}

	var updatedGameStat GameStat
	err = json.NewDecoder(r.Body).Decode(&updatedGameStat)
	if err != nil {
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	statsMux.Lock()
	defer statsMux.Unlock()

	_, exists := gameStats[id]
	if !exists {
		http.Error(w, "Game stat not found", http.StatusNotFound)
		return
	}

	updatedGameStat.ID = id
	gameStats[id] = updatedGameStat

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedGameStat)
}

// for deleting a game stat by ID (DELETE /gamestats/{id})
func deleteGameStatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/gamestats/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid game stat ID", http.StatusBadRequest)
		return
	}

	statsMux.Lock()
	defer statsMux.Unlock()

	_, exists := gameStats[id]
	if !exists {
		http.Error(w, "Game stat not found", http.StatusNotFound)
		return
	}

	delete(gameStats, id)

	w.WriteHeader(http.StatusNoContent)
}

// for listing all of the present game stats (GET /gamestats)
func listGameStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statsMux.Lock()
	defer statsMux.Unlock()

	var gameStatList []GameStat
	for _, gameStat := range gameStats {
		gameStatList = append(gameStatList, gameStat)
	}

	json.NewEncoder(w).Encode(gameStatList)
}

// Main function to run the server
func main() {
	var port string
	fmt.Print("Enter the port number (eg:8080) to run the server on: ") //incase user wants to select their own one in the chance that 8080 is already in use
	fmt.Scanln(&port)

	http.HandleFunc("/gamestats", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listGameStatsHandler(w, r)
		case http.MethodPost:
			createGameStatHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/gamestats/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getGameStatHandler(w, r)
		case http.MethodPut:
			updateGameStatHandler(w, r)
		case http.MethodDelete:
			deleteGameStatHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
