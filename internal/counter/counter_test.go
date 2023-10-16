package counter

import (
	"testing"
)

func TestCountingGeneral(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "john", false, false)

	expected := 12
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingLowerCased(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "john", true, false)

	expected := 0
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingUpperCased(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "John", true, false)

	expected := 12
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingWhole(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "john", false, true)

	expected := 4
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingWholeLowercased(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "john", true, true)

	expected := 0
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingWholeUppercased(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "John", true, true)

	expected := 4
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingEdgeCaseSingleLetter(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "j", false, false)

	expected := 2077
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingEdgeCaseSingleLetterWhole(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "j", false, true)

	expected := 14
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingEdgeCaseWordWithDot(t *testing.T) {
	res := CountRepository("../../testdata/corpus", "J.J", true, true)

	expected := 1
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}

func TestCountingEdgeCaseLongSearchText(t *testing.T) {
	res := CountRepository(
		"../../testdata/corpus",
		`At the end of this class , you 'll be confident in selling at any craft fair around the nation . You 'll be reeling in new customers and exceeding all of your business goals before you know it , thanks to the craft fair markets that have made Katie 's business a success ! <p> If you 're looking to take your online business to the next level , make sure to purchase this awesome class- today . Once you purchase , the class video is yours forever . You can watch at your own pace and reference it whenever you need a refresher.`,
		true,
		false)

	expected := 1
	if res != expected {
		t.Errorf("Expected word count to be %d, but got %d", expected, res)
	}
}
