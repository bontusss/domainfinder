package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := "b1de14abccf5d6088607941e64f14739"
	thesaurus := BigHuge{APIKEY: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonymns(word)
		if err != nil {
			log.Fatal("Failed when looking for synonyms for \"" + word + "\"", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for \"" + word + "\"")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}