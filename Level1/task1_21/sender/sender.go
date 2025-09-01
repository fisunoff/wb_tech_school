package sender

// Под данный интерфейс подходит адаптер

type NotificationSender interface {
	SendMessage(number string, message string)
}
