package main

type InMemoryStore struct {
	scores map[string]int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		scores: map[string]int{},
	}
}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryStore) RecordWin(name string) {
	i.scores[name]++
}
