package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportsRepo) seed() error {
	sportsChoices := []string{"Horse Racing", "Car Racing", "Dog Racing", "Human Racing", "Bike Racing"}

	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, name TEXT, city_address TEXT, num_of_participants INTEGER, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO sports(id, name, city_address, num_of_participants, advertised_start_time) VALUES (?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.RandomChoice(sportsChoices),
				faker.Address().City(),
				faker.Number().Between(0, 1000),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
