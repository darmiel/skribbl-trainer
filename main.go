package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

var availableWords = make(map[string]*WordMeta)
var availableWordsArray []string

func main() {
	rand.Seed(time.Now().Unix())

	if err := json.Unmarshal([]byte(GermanDataJson), &availableWords); err != nil {
		log.Println("Error parsing:", err)
		return
	}

	// initial word map creation
	for word, meta := range availableWords {
		meta.Word = word
		meta.Revealed = ""
		for i := 0; i < len(word); i++ {
			// always reveal "-" and " "
			if word[i] == '-' {
				meta.Revealed += "-"
			} else if word[i] == ' ' {
				meta.Revealed += " "
			} else {
				meta.Revealed += "_"
			}
		}
		log.Println(word, meta.Revealed)
	}

	availableWordsArray = make([]string, len(availableWords))
	var i int
	for word, _ := range availableWords {
		availableWordsArray[i] = word
		i++
	}

	// next word
	NextWord()
	PrintWord()

	go TickWordReveal()
	go ReadUserInput()

	sig := make(chan bool)
	<-sig
}

func NextWord() {
	rand.Seed(int64(time.Now().Nanosecond()))

	// find random word
	word := FindRandomWord()
	if word == "" {
		log.Println("WARN: word was empty:", word)
		return
	}

	meta, ok := availableWords[word]
	if !ok {
		log.Println("WARN: meta not found for word", word)
		return
	}

	// get first 10 "similar" words
	similarLengthWords := FindWordsWithLength(uint32(len(word)))
	rand.Shuffle(len(similarLengthWords), func(i, j int) {
		similarLengthWords[i], similarLengthWords[j] = similarLengthWords[j], similarLengthWords[i]
	})
	if len(similarLengthWords) > 9 {
		similarLengthWords = similarLengthWords[:9]
	} else {
		similarLengthWords = similarLengthWords[:]
	}
	similarLengthWords = append(similarLengthWords, word)
	rand.Shuffle(len(similarLengthWords), func(i, j int) {
		similarLengthWords[i], similarLengthWords[j] = similarLengthWords[j], similarLengthWords[i]
	})

	current = &CurrentWord{
		Original: word,
		Meta:     meta,
		Similar:  similarLengthWords,
	}
}
