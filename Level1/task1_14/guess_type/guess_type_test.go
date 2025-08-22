package guess_type

import "testing"

func TestGuessType(t *testing.T) {
	intAns := GuessType(int(1))
	if intAns != Int {
		t.Errorf("Вместо %s определился тип %s", Int, intAns)
	}

	stringAns := GuessType("abc")
	if stringAns != String {
		t.Errorf("Вместо %s определился тип %s", String, intAns)
	}

	boolAns := GuessType(true)
	if boolAns != Bool {
		t.Errorf("Вместо %s определился тип %s", Bool, intAns)
	}

	chanIntAns := GuessType(make(chan int))
	if chanIntAns != ChanInt {
		t.Errorf("Вместо %s определился тип %s", chanIntAns, intAns)
	}

	chanStringAns := GuessType(make(chan string))
	if chanStringAns != ChanString {
		t.Errorf("Вместо %s определился тип %s", ChanString, chanStringAns)
	}

	chanBoolAns := GuessType(make(chan bool))
	if chanBoolAns != ChanBool {
		t.Errorf("Вместо %s определился тип %s", ChanBool, chanBoolAns)
	}

	chanAns := GuessType(make(chan float64))
	if chanAns != ChanAnother {
		t.Errorf("Вместо %s определился тип %s", ChanAnother, chanAns)
	}

	undefinedAns := GuessType(struct{}{})
	if undefinedAns != Undefined {
		t.Errorf("Вместо %s определился тип %s", Undefined, undefinedAns)
	}
}
