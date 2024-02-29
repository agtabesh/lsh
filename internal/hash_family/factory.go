package hash_family

import (
	"fmt"

	"github.com/agtabesh/lsh/internal/interfaces"
)

type HashFamily string

const XXHash64 HashFamily = "XXHash64"

func GetHashFamily(name HashFamily, count int) (interfaces.HashFamily, error) {
	if name == XXHash64 {
		return newXXHASH64HashFamily(count), nil
	} else {
		return nil, fmt.Errorf("invalid hash family: %s", name)
	}
}
