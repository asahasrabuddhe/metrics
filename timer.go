package metrics

import (
	"time"
	"fmt"
)

var (
	start time.Time
	elapsed time.Duration
)

func Start() {
	start = time.Now()
}

func End() {
	elapsed = time.Since(start)
	fmt.Println("Execution finished in ", elapsed)
}



