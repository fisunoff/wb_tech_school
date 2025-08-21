package bit_operations

import "errors"

var incorrectBitNumberError = errors.New("порядковый номер бита должен быть от 0 до 31")

func SetBit(oldValue int64, bitNumber int) (int64, error) {
	if bitNumber < 0 || bitNumber > 31 {
		return 0, incorrectBitNumberError
	}
	mask := int64(1) << bitNumber
	return oldValue | mask, nil
}

func DistBit(oldValue int64, bitNumber int) (int64, error) {
	if bitNumber < 0 || bitNumber > 31 {
		return 0, incorrectBitNumberError
	}
	mask := ^(int64(1) << bitNumber)
	return oldValue & mask, nil
}
