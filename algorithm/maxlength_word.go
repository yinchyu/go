package main

import (
	"fmt"
	"sort"
	"strings"
)

type ByAge []int

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i] > a[j] }

func secondword(s string) string {
	words := strings.Split(s, " ")
	if len(words) <= 1 {
		return s
	}
	countermap := make(map[int][]string)
	for i := range words {
		l := len(words[i])
		countermap[l] = append(countermap[l], words[i])
	}
	counter := make([]int, 0, len(countermap))
	for i, _ := range countermap {
		counter = append(counter, i)
	}
	if len(counter) < 2 {
		return ""
	}
	//sort.Ints(counter)
	sort.Sort(ByAge(counter))
	fmt.Println(counter)
	secondlength := counter[len(counter)-2]
	value := countermap[secondlength][0]
	fmt.Println(value)
	return value

}
func isCombined(word string, words []string, length int) bool {
	if length == len(word) {
		return true
	}

	for _, w := range words {
		if w == word || length+len(w) > len(word) {
			continue
		}
		if word[length:length+len(w)] == w && isCombined(word, words, length+len(w)) {
			return true
		}
	}
	return false

}
