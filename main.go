package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Child struct {
	value  string
	weight int
}

type State struct {
	value    string
	children map[string]int
}

type Chain struct {
	states map[string]State
}

func cleanString(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, " ")

}

func (c *Chain) FromString(s string) {
	log.Info("Building states from string")

	clean := cleanString(s)
	words := strings.Split(clean, " ")

	for idx, word := range words {
		if idx == len(words)-1 {
			return
		}

		state, ok := c.states[word]
		if !ok {
			state = *new(State)
			state.value = word
			state.children = make(map[string]int)
		}

		nxt := words[idx+1]
		_, ok = state.children[nxt]
		if !ok {
			state.children[nxt] = 0
		} else {
			state.children[nxt]++
		}

		log.Debug("State: %s, children: %v", state.value, state.children)
		c.states[word] = state

	}

}

func weightedSelection(items map[string]int) string {
	rand.Seed(time.Now().Unix())
	flat := []string{}
	for k, w := range items {
		n := 0
		for n <= w {
			flat = append(flat, k)
			n++
		}
	}

	i := rand.Int() % len(flat)
	res := flat[i]
	log.Debugf("Selected %s", res)
	return res

}

func (c *Chain) Generate() {
	log.Info("Generating text")
	length := 20
	curWord := ""
	for _, state := range c.states {
		curWord = state.value
		break
	}

	log.Debugf("Initial word: %s", curWord)

	result := []string{curWord}

	i := 0
	for i < length {
		cur := c.states[curWord]
		log.Debugf("Current: %s", cur.value)

		for len(cur.children) == 0 {
			log.Debugf("No children selecting new word")
			for _, state := range c.states {
				cur = state
				log.Debugf("Current: %s", cur.value)
			}
		}

		log.Debugf("Selecting from %v", cur.children)
		nxt := weightedSelection(cur.children)
		result = append(result, nxt)
		curWord = nxt
		cur = c.states[nxt]
		log.Debugf("Moved state to %v", cur)
		log.Debug("")
		i++
	}

	log.Info(strings.Join(result, " "))

}

func main() {
	log.SetLevel(log.InfoLevel)

	c := new(Chain)
	c.states = make(map[string]State)

	data, err := ioutil.ReadFile("resources/shakespeare.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.FromString(string(data))
	c.Generate()

}
