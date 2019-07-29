# workout-generator
small Go http service to generate workouts
## How to run

`go install`

`workout-generator` - this will start the http service on port `8080`

## Routes

`/` - Landing page
`/all` - list all workouts in repo
`/arm` - list all of the arm workouts in repo
`/generate` - generates a workout consisting of 4 workouts from the repo
