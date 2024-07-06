package main

import (
	"errors"

	"github.com/rthornton128/goncurses"
)

const WelcomeHeader = "Planner on the Go"
const NewTaskHeader = "Add new Task"
const ListHeader = "Your tasks to-do today!"
const DoneTasksHeader = "You have done this!"

func WelcomeScreen(w *goncurses.Window) {
	w.Clear()
	header_len := len(WelcomeHeader)
	_, x := w.MaxYX()
	w.Move(3, x/2-header_len/2)
	w.Print(WelcomeHeader)
	w.Refresh()
}

func InputField(w *goncurses.Window, next_id int) (Task, error) {
	w.Clear()
	header_len := len(NewTaskHeader)
	_, x := w.MaxYX()
	w.Move(3, x/2-header_len/2)
	w.Println(NewTaskHeader)
	w.Println("")

	w.MovePrint(4, 10, "To Do: ")
	w.Refresh()

	todo, _ := w.GetString(50)

	if len(todo) > 0 {
		return Task{
			id:      next_id,
			task:    todo,
			is_done: false,
		}, nil
	}

	return Task{}, errors.New("Empty task")
}

func PlannerList(list []Task, w *goncurses.Window) {
	w.Clear()
	header_len := len(ListHeader)
	_, x := w.MaxYX()
	w.Move(3, x/2-header_len/2)
	w.Println(ListHeader)
	w.Println()

	for _, task := range list {
		task.Render(task.id, w)
	}

	w.Println("")
	w.Println("")
	header_len = len(DoneTasksHeader)
	w.Move(5+len(list), x/2-header_len/2)
	w.Println(DoneTasksHeader)
	w.Println()

	for _, task := range list {
		task.RenderDone(task.id, w)
	}

	w.Refresh()
}
