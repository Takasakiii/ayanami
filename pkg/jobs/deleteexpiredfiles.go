package jobs

import (
	"log"
)

func (j *Jobs) deleteExpiredFiles() {
	log.Println("[INFO] Start deleting expired files")

	err := j.fileService.DeleteExpired()
	if err != nil {
		log.Printf("[ERROR] Failed to delete expired files: %v", err)
		return
	}

	log.Println("[INFO] Finish deleting expired files")
}
