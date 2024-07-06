package main

import (
	"database/sql"

	"github.com/rthornton128/goncurses"
)

type Task struct {
	id      int
	task    string
	is_done bool
}

func (t *Task) Render(w *goncurses.Window, y int) {
	if !t.is_done {
		w.Move(y, 3)
		w.AttrOn(goncurses.A_BOLD)
		w.Printf("[ ] [%d] %q", t.id, t.task)
		w.AttrOff(goncurses.A_BOLD)
	}
}

func (t *Task) RenderDone(w *goncurses.Window, y int) {
	if t.is_done {
		w.Move(y, 70)
		w.AttrOn(goncurses.A_UNDERLINE)
		w.Printf("[x] [%d] %q", t.id, t.task)
		w.AttrOff(goncurses.A_UNDERLINE)
	}
}

func TaskChangeStatus(conn *sql.DB, id int) {
	status := CurrentStatus(conn, id)
	UpdateStatus(conn, id, !status)
}
