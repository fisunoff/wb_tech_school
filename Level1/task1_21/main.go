package main

import (
	"task1_21/adapter"
	"task1_21/some_old_service"
)

func main() {
	oldService := &some_old_service.Message{}
	adapterObject := adapter.NewAdapter(oldService)
	adapterObject.SendMessage("88005553535", "Лучше позвонить, чем у кого-то занимать.")
}
