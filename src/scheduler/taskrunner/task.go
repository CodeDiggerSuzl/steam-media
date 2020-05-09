package taskrunner

import (
	"errors"
	"log"
	"os"
	"stream-media/src/scheduler/dbops"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Error during os.Remove: %v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Error during dbops.ReadVideo in taskrunner: %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			// closeure
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					log.Printf("Error during deleteVideo in VideoClearExecutor: %v", err)
					errMap.Store(id, err)
					return
				}

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

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
