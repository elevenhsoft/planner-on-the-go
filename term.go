package main

import (
	"errors"

	"github.com/rthornton128/goncurses"
)

const WelcomeHeader = "Planner on the Go"
const NewTaskHeader = "Add new Task"
const ListHeader = "Your tasks to-do today!"
const DoneTasksHeader = "You have done this!"
const HelpHeader = "Planner on the Go - Help"

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

func PlannerList(tasks []Task, finished []Task, w *goncurses.Window) {
	w.Clear()
	w.Move(5, 3)
	w.Print(ListHeader)
	w.Move(5, 70)
	w.Print(DoneTasksHeader)
	w.Println("")
	w.Println("")

	y, _ := w.CursorYX()
	for n, task := range tasks {
		task.Render(w, y+n)
	}

	for n, task := range finished {
		task.RenderDone(w, y+n)
	}

	w.Refresh()
}

func HelpScreen(w *goncurses.Window) {
	w.Clear()
	header_len := len(HelpHeader)
	_, x := w.MaxYX()
	w.Move(3, x/2-header_len/2)
	w.Print(HelpHeader)
	w.Println("")
	w.Println("")
	w.Println("List of possible actions, type the letter and press Enter.")
	w.Println("Command line starts with : and waits for your input.")
	w.Println("")
	w.Println(":? - Show help page")
	w.Println(":a - Add new task")
	w.Println(":x - Change task status")
	w.Println(":d - Delete task")
	w.Println(":q - Quit app/help page")
	w.Println("")

	w.Refresh()
}
