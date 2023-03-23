package todo

import "time"

type item struct {
	task        string
	done        bool
	assigned    time.Time
	completedAt time.Time
}

type allTasks []item

func (t *allTasks) Add(task string) {

	todo := item{
		task:        task,
		done:        false,
		assigned:    time.Now(),
		completedAt: time.Time{},
	}

	*t = append(*t, todo)
}
