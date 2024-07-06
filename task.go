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

func (t *Task) Render(idx int, w *goncurses.Window) {
	if !t.is_done {
		w.AttrOn(goncurses.A_BOLD)
		w.Printf("[ ] [%d] %q\n", idx, t.task)
		w.AttrOff(goncurses.A_BOLD)
	}
}

func (t *Task) RenderDone(idx int, w *goncurses.Window) {
	if t.is_done {
		w.AttrOn(goncurses.A_UNDERLINE)
		w.Printf("[x] [%d] %q\n", idx, t.task)
		w.AttrOff(goncurses.A_UNDERLINE)
	}
}

func (t *Task) Mark(conn *sql.DB, id int) {
	t.is_done = !t.is_done

	UpdateStatus(conn, id, t.is_done)
}
