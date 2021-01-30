package main

import (
	"fmt"
	"log"
	"time"
)

var running bool

func TickWordReveal() {
	// running by default
	running = true

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if !running {
				continue
			}
			if current != nil {
				if err := current.Meta.Reveal(); err != nil {
					log.Println("error revealing next character:", err)
					// next word
					NextWord()
				} else {
					log.Println("-> revealed")
				}
				PrintWord()
			}
			break
		}
	}
}

func ReadUserInput() {
	for {
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			log.Fatalln("Error reading input:", err)
			return
		}

		// stop revealing
		if input == "/stop" {
			log.Println("👉 Stopping reveal ticker ...")
			if running {
				running = false
				log.Println("  ✅ stopped")
			} else {
				log.Println("  ❌ not running")
			}
		}

		// start revealing
		if input == "/start" {
			log.Println("👉 Starting reveal ticker ...")
			if !running {
				running = true
				log.Println("  ✅ started")
			} else {
				log.Println("  ❌ not stopped")
			}
		}

		if current == nil {
			continue
		}

		if current.IsWord(input) {
			// Next word
			NextWord()
			log.Println("That was ✨correct✨!")
		} else {
			log.Println("nope!")
		}
	}
}
