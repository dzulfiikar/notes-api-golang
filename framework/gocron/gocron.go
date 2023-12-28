package gocron

import (
	"fmt"

	jobs "notes-api-golang/framework/gocron/jobs"

	goCron "github.com/go-co-op/gocron/v2"
)

func StartGoCronScheduler() {
	scheduler, err := goCron.NewScheduler()

	if err != nil {
		fmt.Println("Error creating scheduler", err)
		panic(err)
	}

	if err != nil {
		fmt.Println("Error creating job", err)
		panic(err)
	}

	RunJobs(scheduler)

	scheduler.Start()

	// defer func() { _ = scheduler.Shutdown() }()

	fmt.Println("GoCron scheduler started")
}

func RunJobs(scheduler goCron.Scheduler) {
	jobs.DeleteNoteJob(scheduler)
}
