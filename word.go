package main

import (
	"errors"
	"math/rand"
	"strings"
)

type CurrentWord struct {
	Original string
	Meta     *WordMeta
	Similar  []string
}

var current *CurrentWord

func (c *CurrentWord) IsWord(w string) bool {
	return strings.ToLower(w) == strings.ToLower(c.Original)
}

type Revealed string

type WordMeta struct {
	Count        uint `json:"count"`
	LastSeenTime uint `json:"lastSeenTime"`

	// full word
	Word     string
	Revealed Revealed
}

func (r Revealed) ReturnUnrevealedArray() (res []uint) {
	for i := 0; i < len(r); i++ {
		if r[i] == '_' {
			res = append(res, uint(i))
		}
	}
	return
}

func (m *WordMeta) RevealAtIndex(idx int) {
	origB := []byte(m.Word)
	revB := []byte(m.Revealed)
	revB[idx] = origB[idx]
	m.Revealed = Revealed(revB)
}

func (m *WordMeta) Reveal() (err error) {
	r := m.Revealed
	a := r.ReturnUnrevealedArray()
	if len(a) <= 0 {
		return errors.New("nothing to reveal")
	}
	// random
	rnd := rand.Intn(len(a))
	m.RevealAtIndex(int(a[rnd]))
	return nil
}

func FindWordsWithLength(length uint32) []string {
	var res []string
	for w, _ := range availableWords {
		if uint32(len(w)) == length {
			res = append(res, w)
		}
	}
	return res
}

func FindRandomWord() string {
	word := availableWordsArray[rand.Intn(len(availableWordsArray))]
	return word
}
