package some_old_service

import "fmt"

// Имитация старого сервиса.
// Будем считать что мы обязаны им пользоваться, но это неудобно, потому что для отправки сообщения нужно сделать 3 вызова.

type Message struct {
	number string // номер телефона
	text   string // текст сообщения
}

// По легенде здесь какой-то сложный код (или не наш код), который нельзя переписать под новый интерфейс.

func (m *Message) Send() {
	fmt.Printf("Отправка сообщения на номер %s, текст: %s\n", m.number, m.text)
}

func (m *Message) SetNumber(number string) {
	m.number = number
}

func (m *Message) SetText(text string) {
	m.text = text
}
