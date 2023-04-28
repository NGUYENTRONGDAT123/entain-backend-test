package service

import (
	"sports/db"
	"sports/proto/sports"

	"golang.org/x/net/context"
)

type Sports interface {
	// ListEvents will return a collection of sport events
	ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsReponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportsService
func NewSportsService(sportsRepo db.SportsRepo) Sports {
	return &sportsService{sportsRepo}
}

func (s *sportsService) ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsReponse, error) {
	events, err := s.sportsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sports.ListEventsReponse{Events: events}, nil
}
