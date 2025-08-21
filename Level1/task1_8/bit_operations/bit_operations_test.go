package bit_operations

import (
	"errors"
	"testing"
)

func TestSetBitCorrect(t *testing.T) {
	// установка уже установленного бита не должна влиять на число
	testInt := int64(5) // 0101
	ans, err := SetBit(testInt, 0)
	if err != nil {
		t.Errorf("Ошибка не ожидалась, но появилась %s", err)
	}
	if ans != testInt {
		t.Errorf("Ожидался ответ %d, но получили %d", testInt, ans)
	}

	// установка снятого бита меняет число 5 на 7
	ans, err = SetBit(testInt, 1)
	if err != nil {
		t.Errorf("Ошибка не ожидалась, но появилась %s", err)
	}
	if ans != 7 {
		t.Errorf("Ожидался ответ %d, но получили %d", 7, ans)
	}
}

func TestSetBitIncorrectPosition(t *testing.T) {
	testInt := int64(5)
	_, err := SetBit(testInt, -1)
	if !errors.Is(err, incorrectBitNumberError) {
		t.Errorf("Ожидалась ошибка incorrectBitNumberError, получили %s", err.Error())
	}

	_, err = SetBit(testInt, 32)
	if !errors.Is(err, incorrectBitNumberError) {
		t.Errorf("Ожидалась ошибка incorrectBitNumberError, получили %s", err.Error())
	}
}

func TestDistBitCorrect(t *testing.T) {
	// сброс уже установленного бита - число меняется
	testInt := int64(5) // 0101
	ans, err := DistBit(testInt, 0)
	if err != nil {
		t.Errorf("Ошибка не ожидалась, но появилась %s", err)
	}
	if ans != 4 {
		t.Errorf("Ожидался ответ %d, но получили %d", 4, ans)
	}

	// снятие 0 бита - то же число
	ans, err = DistBit(testInt, 1)
	if err != nil {
		t.Errorf("Ошибка не ожидалась, но появилась %s", err)
	}
	if ans != testInt {
		t.Errorf("Ожидался ответ %d, но получили %d", testInt, ans)
	}
}

func TestDistBitIncorrectPosition(t *testing.T) {
	testInt := int64(5)
	_, err := DistBit(testInt, -1)
	if !errors.Is(err, incorrectBitNumberError) {
		t.Errorf("Ожидалась ошибка incorrectBitNumberError, получили %s", err.Error())
	}

	_, err = DistBit(testInt, 32)
	if !errors.Is(err, incorrectBitNumberError) {
		t.Errorf("Ожидалась ошибка incorrectBitNumberError, получили %s", err.Error())
	}
}
