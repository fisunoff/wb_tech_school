package adapter

import "task1_21/some_old_service"

// В адаптере мы всё это можем послать за раз.

type Adapter struct {
	smsSender *some_old_service.Message
}

func NewAdapter(smsSender *some_old_service.Message) *Adapter {
	return &Adapter{smsSender: smsSender}
}

func (adapter *Adapter) SendMessage(number string, message string) {
	adapter.smsSender.SetNumber(number)
	adapter.smsSender.SetText(message)
	adapter.smsSender.Send()
}
