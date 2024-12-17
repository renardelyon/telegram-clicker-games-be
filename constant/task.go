package constant

type TaskStatus string
type TaskDaily string
type TaskDesc string

const (
	INCOMPLETE = "incomplete"
	COMPLETE   = "complete"
)

const (
	TASK_DAILY    = "task_daily"
	TASK_ONE_TIME = "task_one_time"
)

const (
	WATCH_ADS          = "watch_ads"
	SUBSCRIBE_TELEGRAM = "subscribe_telegram"
	DAILY_CHECKIN      = "daily_checkin"
	FOLLOW_TIKTOK      = "follow_tiktok"
	FOLLOW_TWITTER     = "follow_twitter"
	INVITE_FRIENDS     = "invite_friends"
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

func (td *TaskDesc) IsValid() (res bool) {
	if td == nil {
		return
	}

	switch *td {
	case WATCH_ADS,
		SUBSCRIBE_TELEGRAM,
		DAILY_CHECKIN,
		FOLLOW_TIKTOK,
		FOLLOW_TWITTER,
		INVITE_FRIENDS:
		return true

	}

	return
}
