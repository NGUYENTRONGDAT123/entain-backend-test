package test

const (
	clearAllDataRace = "clearRace"
	insertNewRace    = "insertNewRace"
)

func getRaceQueriesForTest() map[string]string {
	return map[string]string{
		clearAllDataRace: `DELETE FROM races`,
		insertNewRace: `
		INSERT OR IGNORE INTO
		races
		(id,
		meeting_id,
		name,
		number,
		visible,
		advertised_start_time)
		VALUES
		(?,?,?,?,?,?)`,
	}
}
