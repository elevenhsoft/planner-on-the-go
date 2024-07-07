package main

import (
	"database/sql"

	"github.com/rthornton128/goncurses"
)

type Task struct {
	id      int
	task    string
	is_done bool
	day     int
}

func (t *Task) Render(w *goncurses.Window, y int, weekday int) {
	if !t.is_done && t.day == weekday {
		w.Move(y, 3)
		w.AttrOn(goncurses.A_BOLD)
		w.Printf("[ ] [%d] %q", t.id, t.task)
		w.AttrOff(goncurses.A_BOLD)
	}
}

func (t *Task) RenderDone(w *goncurses.Window, y int, weekday int) {
	if t.is_done && t.day == weekday {
		w.Move(y, 70)
		w.AttrOn(goncurses.A_DIM)
		w.Printf("[x] [%d] %q", t.id, t.task)
		w.AttrOff(goncurses.A_DIM)
	}
}

func TaskChangeStatus(conn *sql.DB, id int) {
	status := CurrentStatus(conn, id)
	UpdateStatus(conn, id, !status)
}
