package main

import (
	"log"

	env "github.com/joho/godotenv"
	"github.com/markkizz/scheduling-my-boring-stuff/cronjob"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal("[Error]: Can not load environment variables.")
	}

	cronjob.Run()

}
