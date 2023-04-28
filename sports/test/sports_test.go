package test

import (
	"context"
	"database/sql"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/protobuf/types/known/timestamppb"
	"sports/db"
	"sports/proto/sports"
	"sports/service"
)

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

func NewTestSportDB() (*sql.DB, error) {
	// Open a database connection using a test database driver.
	sportsDB, err := sql.Open("sqlite3", "./sports_test.db")
	if err != nil {
		return nil, err
	}

	// Initialize the test database by running SQL scripts.
	_, err = sportsDB.Exec(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, name TEXT, city_address TEXT, num_of_participants INTEGER, advertised_start_time DATETIME)`)
	if err != nil {
		return nil, err
	}

	return sportsDB, nil
}

func InsertNewSportsEvent(sportsEvent *sports.Event, r *sql.DB, t *testing.T) {
	ts, err := ptypes.Timestamp(sportsEvent.AdvertisedStartTime)
	if err != nil {
		t.Fatalf("Failed to convert timestamp: %v", err)
	}
	_, err = r.Exec(getSportsEventQueriesForTest()[insertSportEvent],
		&sportsEvent.Id,
		&sportsEvent.Name,
		&sportsEvent.CityAddress,
		&sportsEvent.NumOfParticipants,
		ts,
	)

	if err != nil {
		t.Fatalf("Failed to insert race record: %v", err)
	}
}

func TestListEvents_Default(t *testing.T) {
	// Set up a test database with for testing
	sportsDB, err := NewTestSportDB()
	defer sportsDB.Close()

	//clear the data
	sportsDB.Exec(getSportsEventQueriesForTest()[clearAllDataSportsEvents])

	// Set up a new SportsService with the test database
	sportsRepo := db.NewSportsRepo(sportsDB)
	sportsService := service.NewSportsService(sportsRepo)

	timeTest1, err := time.Parse(time.RFC3339, "1992-04-05T00:00:00Z")
	timeTest2, err := time.Parse(time.RFC3339, "4452-04-05T00:00:00Z")
	timeTest3, err := time.Parse(time.RFC3339, "2004-04-05T00:00:00Z")
	// Insert a race record into the races table
	InsertNewSportsEvent(&sports.Event{
		Id:                  1,
		Name:                "Horse Racing",
		CityAddress:         "Davismouth",
		NumOfParticipants:   826,
		AdvertisedStartTime: timestamppb.New(timeTest1),
	}, sportsDB, t)
	InsertNewSportsEvent(&sports.Event{
		Id:                  2,
		Name:                "Human Racing",
		CityAddress:         "Brisbane",
		NumOfParticipants:   123,
		AdvertisedStartTime: timestamppb.New(timeTest2),
	}, sportsDB, t)
	InsertNewSportsEvent(&sports.Event{
		Id:                  3,
		Name:                "Turtle Racing",
		CityAddress:         "Manchester",
		NumOfParticipants:   456,
		AdvertisedStartTime: timestamppb.New(timeTest3),
	}, sportsDB, t)

	// Set up a new context and request for the ListRaces RPC
	ctx := context.Background()
	req := &sports.ListEventsRequest{}

	// Call the ListRaces RPC
	resp, err := sportsService.ListEvents(ctx, req)
	if err != nil {
		t.Fatalf("Failed to list races: %v", err)
	}

	// Check that the response contains the expected data
	expectedEvents := []*sports.Event{
		&sports.Event{
			Id:                  1,
			Name:                "Horse Racing",
			CityAddress:         "Davismouth",
			NumOfParticipants:   826,
			AdvertisedStartTime: timestamppb.New(timeTest1),
		},
		&sports.Event{
			Id:                  2,
			Name:                "Human Racing",
			CityAddress:         "Brisbane",
			NumOfParticipants:   123,
			AdvertisedStartTime: timestamppb.New(timeTest2),
		},
		&sports.Event{
			Id:                  3,
			Name:                "Turtle Racing",
			CityAddress:         "Manchester",
			NumOfParticipants:   456,
			AdvertisedStartTime: timestamppb.New(timeTest3),
		},
	}
	if !reflect.DeepEqual(resp.Events, expectedEvents) {
		t.Errorf("Response did not match expected value. Got %v, expected %v", resp.Events, expectedEvents)
	}
}

func TestListEvents_WithIdsFilter(t *testing.T) {
	// Set up a test database with for testing
	sportsDB, err := NewTestSportDB()
	defer sportsDB.Close()

	//clear the data
	sportsDB.Exec(getSportsEventQueriesForTest()[clearAllDataSportsEvents])

	// Set up a new SportsService with the test database
	sportsRepo := db.NewSportsRepo(sportsDB)
	sportsService := service.NewSportsService(sportsRepo)

	timeTest1, err := time.Parse(time.RFC3339, "1992-04-05T00:00:00Z")
	timeTest2, err := time.Parse(time.RFC3339, "4452-04-05T00:00:00Z")
	timeTest3, err := time.Parse(time.RFC3339, "2004-04-05T00:00:00Z")
	// Insert a race record into the races table
	InsertNewSportsEvent(&sports.Event{
		Id:                  1,
		Name:                "Horse Racing",
		CityAddress:         "Davismouth",
		NumOfParticipants:   826,
		AdvertisedStartTime: timestamppb.New(timeTest1),
	}, sportsDB, t)
	InsertNewSportsEvent(&sports.Event{
		Id:                  2,
		Name:                "Human Racing",
		CityAddress:         "Brisbane",
		NumOfParticipants:   123,
		AdvertisedStartTime: timestamppb.New(timeTest2),
	}, sportsDB, t)
	InsertNewSportsEvent(&sports.Event{
		Id:                  3,
		Name:                "Turtle Racing",
		CityAddress:         "Manchester",
		NumOfParticipants:   456,
		AdvertisedStartTime: timestamppb.New(timeTest3),
	}, sportsDB, t)

	// Set up a new context and request for the ListRaces RPC
	ctx := context.Background()
	ids := []int64{1, 2}
	req := &sports.ListEventsRequest{
		Filter: &sports.ListEventsRequestFilter{
			Ids: ids,
		},
	}

	// Call the ListRaces RPC
	resp, err := sportsService.ListEvents(ctx, req)
	if err != nil {
		t.Fatalf("Failed to list races: %v", err)
	}

	// Check that the response contains the expected data
	expectedEvents := []*sports.Event{
		&sports.Event{
			Id:                  1,
			Name:                "Horse Racing",
			CityAddress:         "Davismouth",
			NumOfParticipants:   826,
			AdvertisedStartTime: timestamppb.New(timeTest1),
		},
		&sports.Event{
			Id:                  2,
			Name:                "Human Racing",
			CityAddress:         "Brisbane",
			NumOfParticipants:   123,
			AdvertisedStartTime: timestamppb.New(timeTest2),
		},
	}
	if !reflect.DeepEqual(resp.Events, expectedEvents) {
		t.Errorf("Response did not match expected value. Got %v, expected %v", resp.Events, expectedEvents)
	}
}

func TestListEvents_WithIdsFilterButOneIdNotExisted(t *testing.T) {
	// Set up a test database with for testing
	sportsDB, err := NewTestSportDB()
	defer sportsDB.Close()

	//clear the data
	sportsDB.Exec(getSportsEventQueriesForTest()[clearAllDataSportsEvents])

	// Set up a new SportsService with the test database
	sportsRepo := db.NewSportsRepo(sportsDB)
	sportsService := service.NewSportsService(sportsRepo)

	timeTest1, err := time.Parse(time.RFC3339, "1992-04-05T00:00:00Z")
	timeTest2, err := time.Parse(time.RFC3339, "4452-04-05T00:00:00Z")
	timeTest3, err := time.Parse(time.RFC3339, "2004-04-05T00:00:00Z")
	// Insert a race record into the races table
	InsertNewSportsEvent(&sports.Event{
		Id:                  1,
		Name:                "Horse Racing",
		CityAddress:         "Davismouth",
		NumOfParticipants:   826,
		AdvertisedStartTime: timestamppb.New(timeTest1),
	}, sportsDB, t)
	InsertNewSportsEvent(&sports.Event{
		Id:                  2,
		Name:                "Human Racing",
		CityAddress:         "Brisbane",
		NumOfParticipants:   123,
		AdvertisedStartTime: timestamppb.New(timeTest2),
	}, sportsDB, t)
	InsertNewSportsEvent(&sports.Event{
		Id:                  3,
		Name:                "Turtle Racing",
		CityAddress:         "Manchester",
		NumOfParticipants:   456,
		AdvertisedStartTime: timestamppb.New(timeTest3),
	}, sportsDB, t)

	// Set up a new context and request for the ListRaces RPC
	ctx := context.Background()
	ids := []int64{1, 99}
	req := &sports.ListEventsRequest{
		Filter: &sports.ListEventsRequestFilter{
			Ids: ids,
		},
	}

	// Call the ListRaces RPC
	resp, err := sportsService.ListEvents(ctx, req)
	if err != nil {
		t.Fatalf("Failed to list races: %v", err)
	}

	// Check that the response contains the expected data
	expectedEvents := []*sports.Event{
		&sports.Event{
			Id:                  1,
			Name:                "Horse Racing",
			CityAddress:         "Davismouth",
			NumOfParticipants:   826,
			AdvertisedStartTime: timestamppb.New(timeTest1),
		},
	}
	if !reflect.DeepEqual(resp.Events, expectedEvents) {
		t.Errorf("Response did not match expected value. Got %v, expected %v", resp.Events, expectedEvents)
	}
}
