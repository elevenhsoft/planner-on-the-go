package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rthornton128/goncurses"
)

const (
	Welcome = iota
	Help
	Form
	List
)

func getCurrentDay() int {
	now := time.Now()
	return (int(now.Weekday()) + 6) % 7
}

func main() {
	init_localization()

	db := DBInit()
	conn := db.OpenConn()
	TableInit(conn)

	src, err := goncurses.Init()

	if err != nil {
		log.Fatal("init:", err)
	}

	defer goncurses.End()

	loop := true
	screen := Welcome
	current_day := getCurrentDay() + 1
	selected_day := current_day

	conn = db.OpenConn()
	tasks := GetTasksFromDB(conn)
	conn = db.OpenConn()
	finished_tasks := GetFinishedFromDB(conn)

	for loop {
		switch screen {
		case Welcome:
			WelcomeScreen(src)
			char := src.GetChar()

			if char == goncurses.KEY_TAB || char == goncurses.KEY_RETURN {
				screen = Form
			}
			if char == goncurses.KEY_ESC {
				loop = false
			}

		case Help:
			HelpScreen(src)

			y, _ := src.MaxYX()
			src.Move(y-1, 0)
			src.Print(":")
			char, _ := src.GetString(1)

			if char == "q" {
				screen = List
			}

		case Form:
			var next_id int

			if len(tasks) > 0 {
				next_id = tasks[len(tasks)-1].id + 1
			} else {
				next_id = 0
			}

			task, err := InputField(src, next_id)

			if err != nil {
				screen = List
			} else {
				tasks = append(tasks, task)

				conn := db.OpenConn()
				AddToDB(conn, task, selected_day)
			}

			conn = db.OpenConn()
			tasks = GetTasksFromDB(conn)
			conn = db.OpenConn()
			finished_tasks = GetFinishedFromDB(conn)

			screen = List

		case List:
			switch selected_day {
			case 1:
				PlannerList(tasks, finished_tasks, 1, src)
			case 2:
				PlannerList(tasks, finished_tasks, 2, src)
			case 3:
				PlannerList(tasks, finished_tasks, 3, src)
			case 4:
				PlannerList(tasks, finished_tasks, 4, src)
			case 5:
				PlannerList(tasks, finished_tasks, 5, src)
			case 6:
				PlannerList(tasks, finished_tasks, 6, src)
			case 7:
				PlannerList(tasks, finished_tasks, 7, src)
			}

			y, _ := src.MaxYX()
			src.Move(y-1, 0)
			src.Print(":")
			char, _ := src.GetString(1)

			switch char {
			case "q":
				loop = false
			case "?":
				screen = Help
			case "a":
				screen = Form
			case "0":
				selected_day = current_day
			case "1":
				selected_day = 1
			case "2":
				selected_day = 2
			case "3":
				selected_day = 3
			case "4":
				selected_day = 4
			case "5":
				selected_day = 5
			case "6":
				selected_day = 6
			case "7":
				selected_day = 7
			case "x":
				src.Println("")
				mark_text, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "mark_task_info"})
				src.Print(mark_text)
				selected_task, _ := src.GetString(2)
				option, err := strconv.ParseInt(selected_task, 0, 8)

				if err != nil {
					continue
				}

				if int(option) == 0 {
					break
				}

				conn := db.OpenConn()
				TaskChangeStatus(conn, int(option))

				conn = db.OpenConn()
				tasks = GetTasksFromDB(conn)
				conn = db.OpenConn()
				finished_tasks = GetFinishedFromDB(conn)

				src.Refresh()
			case "d":
				src.Println("")
				delete_text, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "delete_task_info"})
				src.Print(delete_text)
				selected_task, _ := src.GetString(2)
				option, err := strconv.ParseInt(selected_task, 0, 8)

				if err != nil {
					continue
				}

				if int(option) == 0 {
					break
				}

				conn := db.OpenConn()
				RemoveFromDB(conn, int(option))

				conn = db.OpenConn()
				tasks = GetTasksFromDB(conn)
				conn = db.OpenConn()
				finished_tasks = GetFinishedFromDB(conn)

				src.Refresh()
			default:
				continue
			}
		}
	}

}
