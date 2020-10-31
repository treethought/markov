package main

import (
	"fmt"
	"strings"
)

func buildMap(corpus string) (m map[string][]string) {
	m = make(map[string][]string, 0)

	words := strings.Split(corpus, " ")
	for idx, word := range words {
		if idx == len(words)-1 {
			return
		}
		next := words[idx+1]
		_, ok := m[word]
		if !ok {
			m[word] = []string{next}
			continue
		}
		m[word] = append(m[word], next)

	}
	return m

}

func main() {

	m := buildMap("my dog is happy and he likes to run my cat is silly and he likes to eat he is also very cute")
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}

}
