package similarity_measure

import (
	"testing"

	"github.com/agtabesh/lsh/internal/types"
)

func TestMeasure1(t *testing.T) {
	s1 := types.Signature{5, 5, 5, 4}
	s2 := types.Signature{5, 5, 5, 4}
	dm := newHammingSimilarity()
	expectedSim := 1.0
	actualSim := dm.Measure(s1, s2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestMeasure2(t *testing.T) {
	s1 := types.Signature{1, 5, 2}
	s2 := types.Signature{4, 1, 5}
	sm := newHammingSimilarity()
	expectedSim := 0.0
	actualSim := sm.Measure(s1, s2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestMeasure3(t *testing.T) {
	s1 := types.Signature{4, 5, 3, 6}
	s2 := types.Signature{1, 1, 3, 7}
	sm := newHammingSimilarity()
	expectedSim := 0.25
	actualSim := sm.Measure(s1, s2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestMeasure4(t *testing.T) {
	s1 := types.Signature{1, 5, 3, 6}
	s2 := types.Signature{1, 4, 3, 7}
	sm := newHammingSimilarity()
	expectedSim := 0.5
	actualSim := sm.Measure(s1, s2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestMeasure5(t *testing.T) {
	s1 := types.Signature{1, 5, 3, 6}
	s2 := types.Signature{1, 4, 3, 6}
	sm := newHammingSimilarity()
	expectedSim := 0.75
	actualSim := sm.Measure(s1, s2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}
