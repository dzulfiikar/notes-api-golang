package jobs

import (
	"fmt"
	"notes-api-golang/framework/mongo/repositories"
	"time"

	mongo "notes-api-golang/framework/mongo"

	goCron "github.com/go-co-op/gocron/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteNoteJob registers a cron job to delete notes that are marked as deleted
// more than 30 days ago. It gets a scheduler instance, creates a delete note task
// function that fetches deleted notes from the database older than 30 days and deletes
// them. It schedules this task to run daily at 00:00:01 using the provided scheduler.
func DeleteNoteJob(scheduler goCron.Scheduler) {

	deleteNoteTask := goCron.NewTask(
		func() {
			fmt.Println("Executing delete note job")
			noteRepository := repositories.NewNoteRepository(mongo.Database)

			// create date 30 days ago from now
			var fromDate = time.Now().AddDate(0, 0, -30)

			result, err := noteRepository.FetchAllNotes(bson.M{"deleted": true, "deleted_at": bson.M{"$lte": fromDate}})
			if err != nil {
				fmt.Println("Error fetching notes", err)
			}

			for _, note := range result {
				_, err := noteRepository.HardDelete(note.ID)
				if err != nil {
					fmt.Println("Error deleting note", err)
				}
			}

			fmt.Println("Delete note job executed")
		},
	)

	scheduler.NewJob(goCron.DailyJob(1, goCron.NewAtTimes(goCron.NewAtTime(0, 0, 1))), deleteNoteTask)

	scheduler.Start()
}
