package guess_type

import "reflect"

const Int = "int"
const String = "string"
const Bool = "bool"
const ChanInt = "chan int"
const ChanString = "chan string"
const ChanBool = "chan bool"
const ChanAnother = "chan"
const Undefined = "undefined" // всё остальное, что не нужно определять по ТЗ

func GuessType(v any) string { // any более современный вариант
	switch v.(type) {
	case int:
		return Int
	case string:
		return String
	case bool:
		return Bool
	case chan int:
		return ChanInt
	case chan string:
		return ChanString
	case chan bool:
		return ChanBool
	}
	if reflect.TypeOf(v).Kind() == reflect.Chan { // chan в целом через .(type) определить не получается
		return ChanAnother
	}
	return Undefined
}
