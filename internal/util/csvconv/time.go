package csvconv

import "time"

func TimeExp(time *time.Time) string {
	if time == nil || time.IsZero() {
		return ""
	}
	// 2020.01.13 17:00:00
	return time.Format("2006.01.02 15:04:05")
}
