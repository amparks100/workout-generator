package main

type Workout struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Reps        int    `json:"reps"`
	Genre       string `json:genre`
}

type Workouts []Workout
