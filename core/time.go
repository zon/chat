package core

import "time"

func FormatTime(t time.Time) string {
	now := time.Now()
	elapsed := now.Sub(t)
	if elapsed < 24 * time.Hour {
	    return t.Format("Today at 3:04PM")
	} else if elapsed < 48 * time.Hour {
		return t.Format("Yesterday at 3:04PM")
	} else {
		return t.Format("1/02/06, 3:04PM")
	}
}
