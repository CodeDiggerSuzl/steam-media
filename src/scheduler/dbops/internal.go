package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// work flow
// api save video id to mysql
// dispatcher query form mysql get video id and send to datachannel
// executor -> datachannel -> video id -> delete videos

// ReadVideoDeletionRecord page query video delete record
func ReadVideoDeletionRecord(count int) ([]string, error) {
	var ids []string
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_record LIMIT ?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("Error during db.Prepare: %v", err)
		return ids, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Error during Query: %v", err)
		return ids, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error during rows.Scan: %v", err)
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// DelVideoDeletionRecord delete delete record by video id
func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_record where video_id  = ?")
	defer stmtDel.Close()
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Error during delete video records %v", err)
		return err
	}
	return nil
}
