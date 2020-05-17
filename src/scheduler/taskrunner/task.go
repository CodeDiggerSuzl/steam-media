package taskrunner

import (
	"errors"
	"log"
	"os"
	"stream-media/src/scheduler/dbops"
	"sync"
)

// remove the video from os
func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Error during os.Remove: %v", err)
		return err
	}
	return nil
}

// VideoClearDispatcher read from db and send all the video id to channel
func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Error during dbops.ReadVideo in taskrunner: %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished ~~")
	}
	// write all the res to data channel
	for _, id := range res {
		dc <- id
	}
	return nil
}

// VideoClearExecutor get from channel and for loop and
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			// closeure
			go func(id interface{}) {
				// can't use vid.string due to the go closeure system
				// remove from os file system
				if err := deleteVideo(id.(string)); err != nil {
					log.Printf("Error during deleteVideo in VideoClearExecutor: %v", err)
					// store the map of each id
					errMap.Store(id, err)
					return
				}
				// del form db-table record
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					log.Printf("Error during DelVideoDeletionRecord in VideoClearExecutor: %v", err)
					errMap.Store(id, err)
					return
				}
			}(vid)

		default:
			break forloop
		}
	}

	// ? for ? what
	errMap.Range(func(k, v interface{}) bool {
		// type assertion
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
