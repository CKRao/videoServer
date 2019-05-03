package dbops

import "log"

func AddVideoDeletionRecord(vid string) error {
	insql := "INSERT INTO video_del_server (video_id) VALUES (?)"
	stmt, err := dbConn.Prepare(insql)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error : ", err)
		return err
	}

	_, err = stmt.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord Exec error :", err)
		return err
	}

	defer stmt.Close()
	return nil
}
