package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//open file
	f, err := os.Open("book.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	// the frequency of the word

	words, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main: %s", err)
	}

	//display the word frequencies
	//for wor

	pairs := sortWordFrequency(words)

	// Print the sorted frequencies

	for _, pair := range pairs {
		fmt.Printf("%s - %d\n", pair.Key, pair.Value)
	}

	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("error with maxWord: %s\n", err)
	}
	fmt.Printf("%#v has a frequency %d\n", w, n)

}

func freq(r io.Reader) (map[string]int, error) {
	//create map to store word frequencies
	wordFreq := make(map[string]int)

	// create a scanner to read the file

	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	//read each word and update the word frequencies
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err

	}
	return wordFreq, nil

}

// for sorting frequency of words
type Pair struct {
	Key   string
	Value int
}

// implemet the len, lees and swapmetods on Pair list
// to satisfy the sort.Interface interface
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value // sort in decending order
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFrequency(m map[string]int) PairList {
	//convert map in to pair list

	pairs := make(PairList, len(m))
	i := 0
	for key, value := range m {
		pairs[i] = Pair{key, value}
		i++
	}
	//sort the pair list

	// sort.Sort(pairs)
	return pairs

}

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}
	maxW := "" //word with max frequency
	maxF := 0  //max frequency of that word

	for k, v := range m {
		if v > maxF {
			maxW = k
			maxF = v
		}
	}
	return maxW, maxF, nil
}
