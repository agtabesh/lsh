package similarity_measure

import (
	"testing"

	"github.com/agtabesh/lsh/internal/types"
)

func TestCalculate1(t *testing.T) {
	p1 := map[types.VectorID]float64{
		"100": 5,
		"101": 5,
		"102": 5,
	}
	p2 := map[types.VectorID]float64{
		"100": 5,
		"101": 5,
		"102": 5,
	}
	sm := NewCosineSimilarityMeasure(5)
	expectedSim := 1.0
	actualSim := sm.Calculate(p1, p2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestCalculate2(t *testing.T) {
	p1 := map[types.VectorID]float64{
		"100": 1,
		"101": 5,
		"102": 2,
	}
	p2 := map[types.VectorID]float64{
		"100": 4,
		"101": 1,
		"102": 5,
	}
	sm := NewCosineSimilarityMeasure(5)
	expectedSim := 0.53526
	actualSim := sm.Calculate(p1, p2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestCalculate3(t *testing.T) {
	p1 := map[types.VectorID]float64{
		"100": 4,
		"101": 5,
		"102": 3,
	}
	p2 := map[types.VectorID]float64{
		"100": 1,
		"101": 1,
		"102": 2,
	}
	sm := NewCosineSimilarityMeasure(5)
	expectedSim := 0.86603
	actualSim := sm.Calculate(p1, p2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}

func TestCalculate4(t *testing.T) {
	p1 := map[types.VectorID]float64{
		"100": 4,
		"101": 5,
		"102": 3,
	}
	p2 := map[types.VectorID]float64{
		"100": 4,
		"103": 2,
	}
	sm := NewCosineSimilarityMeasure(5)
	expectedSim := 0.50596
	actualSim := sm.Calculate(p1, p2)
	if actualSim != expectedSim {
		t.Error("Expected similarity of ", expectedSim, ", instead got ", actualSim)
	}
}
