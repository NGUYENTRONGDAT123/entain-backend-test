package test

import (
	"context"
	"database/sql"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"git.neds.sh/matty/entain/racing/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

func NewTestDB() (*sql.DB, error) {
	// Open a database connection using a test database driver.
	racingDB, err := sql.Open("sqlite3", "./race_test.db")
	if err != nil {
		return nil, err
	}

	// Initialize the test database by running SQL scripts.
	_, err = racingDB.Exec(`CREATE TABLE IF NOT EXISTS races (id INTEGER PRIMARY KEY, meeting_id INTEGER, name TEXT, number INTEGER, visible INTEGER, advertised_start_time DATETIME)`)
	if err != nil {
		return nil, err
	}

	return racingDB, nil
}

func InsertNewRace(race *racing.Race, r *sql.DB, t *testing.T) {
	ts, err := ptypes.Timestamp(race.AdvertisedStartTime)
	if err != nil {
		t.Fatalf("Failed to convert timestamp: %v", err)
	}
	_, err = r.Exec(getRaceQueriesForTest()[insertNewRace],
		&race.Id,
		&race.MeetingId,
		&race.Name,
		&race.Number,
		&race.Visible,
		ts,
	)

	if err != nil {
		t.Fatalf("Failed to insert race record: %v", err)
	}
}

func TestListRace_Default(t *testing.T) {
	// Set up a test database with for testing
	racingDB, err := NewTestDB()
	defer racingDB.Close()

	//clear the data
	racingDB.Exec(getRaceQueriesForTest()[clearAllDataRace])

	// Set up a new RacingService with the test database
	racesRepo := db.NewRacesRepo(racingDB)
	racingService := service.NewRacingService(racesRepo)

	timeTest, err := time.Parse(time.RFC3339, "1992-04-05T00:00:00Z")
	// Insert a race record into the races table
	InsertNewRace(&racing.Race{
		Id:                  1,
		MeetingId:           1,
		Name:                "Test Race 1",
		Number:              1,
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(timeTest),
	}, racingDB, t)
	InsertNewRace(&racing.Race{
		Id:                  2,
		MeetingId:           2,
		Name:                "Test Race 2",
		Number:              2,
		Visible:             false,
		AdvertisedStartTime: timestamppb.New(timeTest),
	}, racingDB, t)
	InsertNewRace(&racing.Race{
		Id:                  3,
		MeetingId:           3,
		Name:                "Test Race 3",
		Number:              3,
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(timeTest),
	}, racingDB, t)

	// Set up a new context and request for the ListRaces RPC
	ctx := context.Background()
	req := &racing.ListRacesRequest{}

	// Call the ListRaces RPC
	resp, err := racingService.ListRaces(ctx, req)
	if err != nil {
		t.Fatalf("Failed to list races: %v", err)
	}

	// Check that the response contains the expected data
	expectedRaces := []*racing.Race{
		&racing.Race{
			Id:                  1,
			MeetingId:           1,
			Name:                "Test Race 1",
			Number:              1,
			Visible:             true,
			AdvertisedStartTime: timestamppb.New(timeTest),
		},
		&racing.Race{
			Id:                  2,
			MeetingId:           2,
			Name:                "Test Race 2",
			Number:              2,
			Visible:             false,
			AdvertisedStartTime: timestamppb.New(timeTest),
		},
		&racing.Race{
			Id:                  3,
			MeetingId:           3,
			Name:                "Test Race 3",
			Number:              3,
			Visible:             true,
			AdvertisedStartTime: timestamppb.New(timeTest),
		},
	}
	if !reflect.DeepEqual(resp.Races, expectedRaces) {
		t.Errorf("Response did not match expected value. Got %v, expected %v", resp.Races, expectedRaces)
	}
}

func TestListRace_VisibleFilter(t *testing.T) {
	// Set up a test database with for testing
	racingDB, err := NewTestDB()
	defer racingDB.Close()

	//clear the data
	racingDB.Exec(getRaceQueriesForTest()[clearAllDataRace])

	// Set up a new RacingService with the test database
	racesRepo := db.NewRacesRepo(racingDB)
	racingService := service.NewRacingService(racesRepo)

	timeTest, err := time.Parse(time.RFC3339, "1992-04-05T00:00:00Z")
	// Insert a race record into the races table
	InsertNewRace(&racing.Race{
		Id:                  1,
		MeetingId:           1,
		Name:                "Test Race 1",
		Number:              1,
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(timeTest),
	}, racingDB, t)
	InsertNewRace(&racing.Race{
		Id:                  2,
		MeetingId:           2,
		Name:                "Test Race 2",
		Number:              2,
		Visible:             false,
		AdvertisedStartTime: timestamppb.New(timeTest),
	}, racingDB, t)
	InsertNewRace(&racing.Race{
		Id:                  3,
		MeetingId:           3,
		Name:                "Test Race 3",
		Number:              3,
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(timeTest),
	}, racingDB, t)

	// Set up a new context and request for the ListRaces RPC
	ctx := context.Background()
	visible := true
	req := &racing.ListRacesRequest{
		Filter: &racing.ListRacesRequestFilter{
			Visible: &visible,
		}}

	// Call the ListRaces RPC
	resp, err := racingService.ListRaces(ctx, req)
	if err != nil {
		t.Fatalf("Failed to list races: %v", err)
	}

	// Check that the response contains the expected data
	expectedRaces := []*racing.Race{
		&racing.Race{
			Id:                  1,
			MeetingId:           1,
			Name:                "Test Race 1",
			Number:              1,
			Visible:             true,
			AdvertisedStartTime: timestamppb.New(timeTest),
		},
		&racing.Race{
			Id:                  3,
			MeetingId:           3,
			Name:                "Test Race 3",
			Number:              3,
			Visible:             true,
			AdvertisedStartTime: timestamppb.New(timeTest),
		},
	}
	if !reflect.DeepEqual(resp.Races, expectedRaces) {
		t.Errorf("Response did not match expected value. Got %v, expected %v", resp.Races, expectedRaces)
	}
}
