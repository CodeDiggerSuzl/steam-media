package dbops

import "log"

// AddVideoDeleteRecord add delete record
func AddVideoDeleteRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_record(video_id) VALUES (?)")
	defer stmtIns.Close()
	if err != nil {
		log.Printf("Error during dbConn.Prepare: %v", err)
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("Error during stmtIns.Exec: %v", err)
		return err
	}
	return nil
}
