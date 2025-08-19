package main

import (
	"fmt"
	"strings"
)

type Human struct {
	name string
	age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Привет! Меня зовут %s, мой возраст: %d.\n", h.name, h.age)
}

func (h *Human) GetFullInfo() string {
	return fmt.Sprintf("Имя: %s, возраст: %d", h.name, h.age)
}

type Action struct {
	Human
	actions []string // может ещё делать какие-то действия
}

func (a *Action) getActionsString() string {
	return strings.Join(a.actions, "; ")
}

func (a *Action) PrintActions() {
	fmt.Printf("%s может выполнять следующие действия: %s.\n", a.name, a.getActionsString())
}

func (a *Action) GetFullInfo() string {
	return a.Human.GetFullInfo() + fmt.Sprintf(", Возможные действия: %s", a.getActionsString())
}

func main() {
	anton := Human{"Антон", 20}
	actions := []string{"Писать код", "Работать работу"}
	antonWorker := Action{anton, actions}
	antonWorker.SayHello()
	antonWorker.PrintActions()
	fmt.Println(anton.GetFullInfo())
	fmt.Println(antonWorker.GetFullInfo())
}
