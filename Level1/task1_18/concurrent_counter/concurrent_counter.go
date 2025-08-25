package concurrent_counter

import "sync/atomic"

type Counter struct {
	value uint64 // поле приватное чтобы не могли поменять небезопасно
}

// GetValue получить значение счетчика
func (c *Counter) GetValue() uint64 {
	return atomic.LoadUint64(&c.value)
}

// Inc увеличить значение счетчика на 1
func (c *Counter) Inc() {
	// Был выбран вариант с атомарными операциями так как он быстрее за счет обеспечения атомарности на уровне ЦП, а не блокировки горутины.
	atomic.AddUint64(&c.value, 1)
}
