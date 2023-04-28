package db

const (
	sportEventsList = "list"
)

func getSportEventQueries() map[string]string {
	return map[string]string{
		sportEventsList: `
			SELECT 
				id, 
				name, 
				city_address, 
				num_of_participants, 
				advertised_start_time
			FROM sports
		`,
	}
}
