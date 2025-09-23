package sleeper

import "time"

func Sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}
