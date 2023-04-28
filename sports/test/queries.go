package test

const (
	clearAllDataSportsEvents = "clearSportsEvents"
	insertSportEvent         = "insertSportEvent"
)

func getSportsEventQueriesForTest() map[string]string {
	return map[string]string{
		clearAllDataSportsEvents: `DELETE FROM sports`,
		insertSportEvent: `
		INSERT OR IGNORE INTO
		sports
		(id,
		name,
		city_address, 
		num_of_participants, 
		advertised_start_time) 
		VALUES 
		(?,?,?,?,?)
		`,
	}
}
