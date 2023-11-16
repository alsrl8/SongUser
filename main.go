package main

import (
	"SongUser/web"
	"log"
)

func main() {
	router := web.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
