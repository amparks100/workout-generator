package main

func AllTheThings() [13]Workout {
	var workouts [13]Workout

	workouts[0] = Workout{Name: "Jumping jacks", Description: "Jump alot", Reps: 20, Genre: "cardio"}
	workouts[1] = Workout{Name: "Elliptical", Description: "It's like running, but cheating", Reps: 2, Genre: "cardio"}
	workouts[2] = Workout{Name: "Treadmill", Description: "Vroom Vroom", Reps: 2, Genre: "cardio"}
	workouts[3] = Workout{Name: "Bike", Description: "A cheaper car", Reps: 10, Genre: "cardio"}
	workouts[4] = Workout{Name: "Bicep Curls", Description: "Lift those weights", Reps: 10, Genre: "arm"}
	workouts[5] = Workout{Name: "Hammer Curls", Description: "Pow Pow", Reps: 10, Genre: "arm"}
	workouts[6] = Workout{Name: "Shoulder Press", Description: "Reach for the sky!", Reps: 10, Genre: "arm"}
	workouts[7] = Workout{Name: "Upward Row", Description: "Strengthen those arms and back", Reps: 10, Genre: "arm"}
	workouts[8] = Workout{Name: "Tricep Drops", Description: "Keep your arms close to your head!", Reps: 10, Genre: "arm"}
	workouts[9] = Workout{Name: "Squat", Description: "And down and up and down and up", Reps: 15, Genre: "legs"}
	workouts[10] = Workout{Name: "Jump Squat", Description: "Feel your inner frog", Reps: 15, Genre: "legs"}
	workouts[11] = Workout{Name: "Plia", Description: "Ballerina Time!", Reps: 15, Genre: "legs"}
	workouts[12] = Workout{Name: "Lunges", Description: "Epic walking", Reps: 15, Genre: "legs"}

	return workouts
}
