package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Eochs/cyoa"
)

func main() {
	port := flag.Int("port", 3001, "the port to start CYOA on")
	filename := flag.String("file", "gopher.json", "JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using hte story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
