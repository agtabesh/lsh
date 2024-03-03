package store

import (
	"fmt"

	"github.com/agtabesh/lsh/interfaces"
)

type Store string

const InMemoryStore Store = "IN_MEMORY_STORE"

func GetStore(name Store) (interfaces.Store, error) {
	if name == InMemoryStore {
		return newInMemoryStore(), nil
	} else {
		return nil, fmt.Errorf("invalid store: %s", name)
	}
}
