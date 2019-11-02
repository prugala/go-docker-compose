package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env.local") //Overwrite env values on a local envirement

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println(os.Getenv("HELLO"))
}
