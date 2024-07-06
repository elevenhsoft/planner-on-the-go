package main

import (
	"log"
	"strconv"

	"github.com/rthornton128/goncurses"
)

const (
	Welcome = iota
	Help
	Form
	List
)

func main() {
	db := DBInit()
	conn := db.OpenConn()
	TableInit(conn)

	src, err := goncurses.Init()

	if err != nil {
		log.Fatal("init:", err)
	}

	loop := true
	screen := Welcome

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

			src.Print(":")
			char, _ := src.GetString(1)

			if char == "q" {
				screen = List
			}
		case Form:
			task, err := InputField(src, tasks[len(tasks)-1].id+1)

			if err != nil {
				screen = List
			} else {
				tasks = append(tasks, task)

				conn := db.OpenConn()
				AddToDB(conn, task)
			}

			screen = List

		case List:
			PlannerList(tasks, finished_tasks, src)

			src.Move(0, 0)
			src.Print(":")
			char, _ := src.GetString(1)

			switch char {
			case "q":
				loop = false
			case "?":
				screen = Help
			case "a":
				screen = Form
			case "x":
				src.Println("")
				src.Print("What have you done? [number] / 0 - cancel: ")
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
				src.Print("Delete? [number] / 0 - cancel: ")
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

	defer goncurses.End()
}
