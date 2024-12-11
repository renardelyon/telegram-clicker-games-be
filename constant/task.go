package constant

type TaskStatus string
type TaskDaily string

const (
	INCOMPLETE = "incomplete"
	COMPLETE   = "complete"
)

const (
	TASK_DAILY    = "task_daily"
	TASK_ONE_TIME = "task_one_time"
)

func (ts *TaskStatus) IsValid() (res bool) {
	if ts == nil {
		return
	}

	switch *ts {
	case INCOMPLETE, COMPLETE:
		return true

	}

	return
}

func (td *TaskDaily) IsValid() (res bool) {
	if td == nil {
		return
	}

	switch *td {
	case TASK_DAILY, TASK_ONE_TIME:
		return true

	}

	return
}
