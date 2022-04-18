package cron

import (
	"fmt"
	"testing"

	"github.com/robfig/cron"
)

func TestInitCronPhase_Init(t *testing.T) {
	Cron := cron.New()

	Cron.AddFunc("@every 10s", func() {
		fmt.Println("job1 run")
	})

	Cron.Start()
	Cron.AddFunc("@every 10s", func() {
		fmt.Println("job2 run")
	})
	select {}
}
