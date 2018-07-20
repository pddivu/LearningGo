package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)
	for scanner.Scan() {
		word := scanner.Text()
		n := hashbucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	for i := 0; i < 12; i++ {
		fmt.Println(i, "-", len(buckets[i]))
	}
	//fmt.Println(buckets)
	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))
}
func hashbucket(word string, bucket int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % bucket
}
