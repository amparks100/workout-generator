package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WELCOME")
}

func AllWorkoutIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(AllTheThings())
}

func ArmWorkoutIndex(w http.ResponseWriter, r *http.Request) {
	var result []Workout
	for _, workout := range AllTheThings() {
		if workout.Genre == "arm" {
			log.Print("found a thing")
			result = append(result, workout)
		}
	}
	json.NewEncoder(w).Encode(result)
}

func GenerateIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CreateAWorkout(4))
}

func CreateAWorkout(amount int) []Workout {
	var workouts []Workout
	total := AllTheThings()
	for i := 0; i < amount; i++ {
		workouts = append(workouts, total[rand.Intn(len(total))])
	}

	return workouts
}
