package main

import (
	"log"
	"strconv"

	"github.com/rthornton128/goncurses"
)

const (
	Welcome = iota
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

	for loop {
		switch screen {
		case Welcome:
			{
				WelcomeScreen(src)
				char := src.GetChar()

				if char == goncurses.KEY_TAB || char == goncurses.KEY_RETURN {
					screen = Form
				}
				if char == goncurses.KEY_ESC {
					loop = false
				}

			}
		case Form:
			{
				task, err := InputField(src, tasks[len(tasks)-1].id+1)

				if err != nil {
					screen = List
				} else {
					tasks = append(tasks, task)

					conn := db.OpenConn()
					AddToDB(conn, task)
				}

				screen = List

			}
		case List:
			{
				PlannerList(tasks, src)

				src.Print(":")
				char, _ := src.GetString(1)

				switch char {
				case "q":
					loop = false
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
					tasks[option-1].Mark(conn, int(option))

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

					src.Refresh()
				default:
					continue
				}
			}
		}
	}

	defer goncurses.End()
}
