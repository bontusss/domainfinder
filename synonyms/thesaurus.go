package main

type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}