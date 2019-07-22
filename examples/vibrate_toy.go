package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sextech/lovense"
)

func main() {
	remote := lovense.NewRemote()
	toys, err := remote.Discover()

	if err != nil {
		log.Fatal(err)
	}

	var toy *lovense.Toy

	// Search for first connected toy
	for _, t := range toys {
		if t.Status == lovense.Connected {
			toy = t
			break
		}
	}

	if toy == nil {
		log.Fatal("no connected toy found")
	}

	fmt.Printf("Using toy : %s\n", toy.Name)

	toy.Vibrate(lovense.AllVibrator, 5)
	time.Sleep(2 * time.Second)
	toy.Vibrate(lovense.AllVibrator, 0)
}
