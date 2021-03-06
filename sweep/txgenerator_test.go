package sweep

import (
	"testing"

	"github.com/qtumproject/qtumsuite/wire"
	"github.com/qtumproject/qlnd/input"
)

var (
	witnessTypes = []input.WitnessType{
		input.CommitmentTimeLock,
		input.HtlcAcceptedSuccessSecondLevel,
		input.HtlcOfferedRemoteTimeout,
		input.WitnessKeyHash,
	}
	expectedWeight  = int64(1459)
	expectedSummary = "1 CommitmentTimeLock, 1 " +
		"HtlcAcceptedSuccessSecondLevel, 1 HtlcOfferedRemoteTimeout, " +
		"1 WitnessKeyHash"
)

// TestWeightEstimate tests that the estimated weight and number of CSVs/CLTVs
// used is correct for a transaction that uses inputs with the witness types
// defined in witnessTypes.
func TestWeightEstimate(t *testing.T) {
	t.Parallel()

	var inputs []input.Input
	for _, witnessType := range witnessTypes {
		inputs = append(inputs, input.NewBaseInput(
			&wire.OutPoint{}, witnessType,
			&input.SignDescriptor{}, 0,
		))
	}

	_, weight := getWeightEstimate(inputs)
	if weight != expectedWeight {
		t.Fatalf("unexpected weight. expected %d but got %d.",
			expectedWeight, weight)
	}
	summary := inputTypeSummary(inputs)
	if summary != expectedSummary {
		t.Fatalf("unexpected summary. expected %s but got %s.",
			expectedSummary, summary)
	}
}
