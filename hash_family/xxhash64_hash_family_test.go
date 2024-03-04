package hash_family

import (
	"testing"

	"github.com/agtabesh/lsh/types"
)

func TestHash(t *testing.T) {
	count := 3
	text := "sample text"
	hf := NewXXHASH64HashFamily(count)
	hashes := hf.Hash(text)

	expectedCount := count
	actualCount := len(hashes)
	if actualCount != expectedCount {
		t.Error("Expected hash values count: ", expectedCount, ", got: ", actualCount)
	}

	expectedValue := types.SignatureEntry(10216336028491398715)
	actualValue := hashes[0]
	if actualValue != expectedValue {
		t.Error("Expected hash value: ", expectedValue, ", got: ", actualValue)
	}

	expectedValue = types.SignatureEntry(7095479961646886861)
	actualValue = hashes[1]
	if actualValue != expectedValue {
		t.Error("Expected hash value: ", expectedValue, ", got: ", actualValue)
	}

	expectedValue = types.SignatureEntry(5501795655278852606)
	actualValue = hashes[2]
	if actualValue != expectedValue {
		t.Error("Expected hash value: ", expectedValue, ", got: ", actualValue)
	}
}

func TestMinHash(t *testing.T) {
	count := 3
	vector := types.Vector{"a": 1, "b": 1, "c": 1}
	hf := NewXXHASH64HashFamily(count)
	hashes := hf.MinHash(vector)

	expectedCount := count
	actualCount := len(hashes)
	if actualCount != expectedCount {
		t.Error("Expected hash values count: ", expectedCount, ", got: ", actualCount)
	}

	expectedValue := types.SignatureEntry(8666379929374662555)
	actualValue := hashes[0]
	if actualValue != expectedValue {
		t.Error("Expected hash value: ", expectedValue, ", got: ", actualValue)
	}

	expectedValue = types.SignatureEntry(6429003490305337916)
	actualValue = hashes[1]
	if actualValue != expectedValue {
		t.Error("Expected hash value: ", expectedValue, ", got: ", actualValue)
	}

	expectedValue = types.SignatureEntry(815288398222543995)
	actualValue = hashes[2]
	if actualValue != expectedValue {
		t.Error("Expected hash value: ", expectedValue, ", got: ", actualValue)
	}
}
