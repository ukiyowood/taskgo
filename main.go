package taskgo

import "fmt"


type Task struct {
	ID string
	Name string
}

func (task *Task)Print(text string) {
	fmt.Printf("[taskgo--2] %s\n", text)
}