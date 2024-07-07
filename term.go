package main

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rthornton128/goncurses"
)

const WelcomeHeader = "Planner on the Go"

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
	new_task_text, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "new_task_header"})
	w.Move(3, 3)
	w.Println(new_task_text)
	w.Println("")

	todo_text, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "todo_prompt"})
	w.MovePrint(5, 3, todo_text)
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

func PlannerList(tasks []Task, finished []Task, weekday int, w *goncurses.Window) {
	w.Clear()

	w.AttrOn(goncurses.A_BOLD)

	day := localized_week()

	switch weekday {
	case 1:
		w.MovePrint(2, 3, day[0])
	case 2:
		w.MovePrint(2, 3, day[1])
	case 3:
		w.MovePrint(2, 3, day[2])
	case 4:
		w.MovePrint(2, 3, day[3])
	case 5:
		w.MovePrint(2, 3, day[4])
	case 6:
		w.MovePrint(2, 3, day[5])
	case 7:
		w.MovePrint(2, 3, day[6])
	}
	w.AttrOff(goncurses.A_BOLD)

	w.AttrOn(goncurses.A_UNDERLINE)
	w.Move(5, 3)
	list_header, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "list_header"})
	w.Print(list_header)
	w.Move(5, 70)
	done_header, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "done_header"})
	w.Print(done_header)
	w.AttrOff(goncurses.A_UNDERLINE)
	w.Println("")
	w.Println("")

	y, _ := w.CursorYX()
	task_counter := 0
	for _, task := range tasks {
		if task.day == weekday {
			task_counter++
		}
		task.Render(w, y+task_counter, weekday)
	}

	task_counter = 0
	for _, task := range finished {
		if task.day == weekday {
			task_counter++
		}
		task.RenderDone(w, y+task_counter, weekday)
	}

	w.Refresh()
}

func HelpScreen(w *goncurses.Window) {
	w.Clear()
	help_header, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "help_header"})
	w.Move(3, 3)
	w.Print(help_header)
	w.Println("")
	w.Println("")
	help_info, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "help_info"})
	help_cli_info, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "help_cli_info"})
	w.Println(help_info)
	w.Println(help_cli_info)
	w.Println("")
	show_help_page, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "show_help_page"})
	add_new_task_help, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "add_new_task_help"})
	change_task_status_help, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "change_task_help"})
	delete_task_help, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "delete_task_help"})
	quit_app_help, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "quit_app_help"})
	w.Println(show_help_page)
	w.Println(change_task_status_help)
	w.Println(add_new_task_help)
	w.Println(delete_task_help)
	w.Println(quit_app_help)
	w.Println("")

	w.Refresh()
}
