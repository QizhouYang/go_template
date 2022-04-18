package cron

import (
	"fmt"

	"go_template/pkg/util/cron/job"
	"time"

	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

const phaseName = "cron"

type InitCronPhase struct {
	Enable bool
}

func (c *InitCronPhase) Init() error {
	//Cron = cron.New()
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	Cron = cron.New(cron.WithLocation(nyc))
	if c.Enable {
		_, err := Cron.AddJob("@hourly", job.NewRefreshHostInfo())
		if err != nil {
			return fmt.Errorf("can not add corn job: %s", err.Error())
		}

		//_, err = Cron.AddJob("@every 1m", job.NewClusterHealthCheck())
		//if err != nil {
		//	return fmt.Errorf("can not add cluster health check corn job: %s", err.Error())
		//}
		Cron.Start()
	}
	return nil
}

func (c *InitCronPhase) PhaseName() string {
	return phaseName
}
