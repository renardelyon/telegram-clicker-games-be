package constant

type TaskStatus string

const (
	INCOMPLETE        = "incomplete"
	COMPLETE          = "complete"
	COMPLETE_NO_RECUR = "complete_no_recur"
)

func (ts *TaskStatus) IsValid() (res bool) {
	if ts == nil {
		return
	}

	switch *ts {
	case INCOMPLETE, COMPLETE, COMPLETE_NO_RECUR:
		return true

	}

	return
}
