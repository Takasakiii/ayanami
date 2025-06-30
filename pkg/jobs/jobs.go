package jobs

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/go-co-op/gocron/v2"
	"time"
)

type Jobs struct {
	fileService file.Service
}

func NewJobs(fileService file.Service) *Jobs {
	return &Jobs{
		fileService: fileService,
	}
}

func (j *Jobs) Init() error {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return fmt.Errorf("jobs init newscheduler: %v", err)
	}

	_, err = scheduler.NewJob(gocron.DurationJob(time.Hour), gocron.NewTask(j.deleteExpiredFiles))
	if err != nil {
		return fmt.Errorf("jobs init deleteexpiredfiles: %v", err)
	}

	scheduler.Start()

	return nil
}
